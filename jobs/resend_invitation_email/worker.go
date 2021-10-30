// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.
package resend_invitation_email

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/mattermost/mattermost-server/v6/app"
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/shared/mlog"
)

const TwentyFourHoursInMillis int64 = 86400000
const FourtyEightHoursInMillis int64 = 172800000
const SeventyTwoHoursInMillis int64 = 259200000

type ResendInvitationEmailWorker struct {
	name    string
	stop    chan bool
	stopped chan bool
	jobs    chan model.Job
	App     *app.App
}

func (rse *ResendInvitationEmailJobInterfaceImpl) MakeWorker() model.Worker {
	worker := ResendInvitationEmailWorker{
		name:    ResendInvitationEmailJob,
		stop:    make(chan bool, 1),
		stopped: make(chan bool, 1),
		jobs:    make(chan model.Job),
		App:     rse.App,
	}
	return &worker
}

func (rseworker *ResendInvitationEmailWorker) Run() {
	mlog.Debug("Worker started", mlog.String("worker", rseworker.name))

	defer func() {
		mlog.Debug("Worker finished", mlog.String("worker", rseworker.name))
		rseworker.stopped <- true
	}()

	for {
		select {
		case <-rseworker.stop:
			mlog.Debug("Worker received stop signal", mlog.String("worker", rseworker.name))
			return
		case job := <-rseworker.jobs:
			mlog.Debug("Worker received a new candidate job.", mlog.String("worker", rseworker.name))
			rseworker.DoJob(&job)
		}
	}
}

func (rseworker *ResendInvitationEmailWorker) Stop() {
	mlog.Debug("Worker stopping", mlog.String("worker", rseworker.name))
	rseworker.stop <- true
	<-rseworker.stopped
}

func (rseworker *ResendInvitationEmailWorker) JobChannel() chan<- model.Job {
	return rseworker.jobs
}

func (rseworker *ResendInvitationEmailWorker) DoJob(job *model.Job) {
	resendInviteEmailIntervalFlag := rseworker.App.Config().FeatureFlags.ResendInviteEmailInterval

	switch resendInviteEmailIntervalFlag {
	case "48":
		rseworker.DoJob_24_48(job)
	case "72":
		rseworker.DoJob_24_72(job)
	default:
		rseworker.DoJob_24(job)
	}
}

func (rseworker *ResendInvitationEmailWorker) DoJob_24(job *model.Job) {
	elapsedTimeSinceSchedule, DurationInMillis_24, _, _ := rseworker.GetDurations(job)
	if elapsedTimeSinceSchedule > DurationInMillis_24 {
		rseworker.ResendEmails(job)
		rseworker.TearDown(job)
	}
}

func (rseworker *ResendInvitationEmailWorker) DoJob_24_48(job *model.Job) {
	elapsedTimeSinceSchedule, DurationInMillis_24, DurationInMillis_48, _ := rseworker.GetDurations(job)
	systemValue, _ := rseworker.App.Srv().Store.System().GetByName(model.OverUserLimitForgivenCount)
	if (elapsedTimeSinceSchedule > DurationInMillis_24) && (elapsedTimeSinceSchedule < DurationInMillis_48) && (systemValue.Value == "0") {
		rseworker.ResendEmails(job)
		rseworker.setNumResendEmailSent("1")
	} else if elapsedTimeSinceSchedule > DurationInMillis_48 {
		rseworker.ResendEmails(job)
		rseworker.TearDown(job)
	}
}

func (rseworker *ResendInvitationEmailWorker) DoJob_24_72(job *model.Job) {
	elapsedTimeSinceSchedule, DurationInMillis_24, _, DurationInMillis_72 := rseworker.GetDurations(job)
	systemValue, _ := rseworker.App.Srv().Store.System().GetByName(model.OverUserLimitForgivenCount)
	if (elapsedTimeSinceSchedule > DurationInMillis_24) && (elapsedTimeSinceSchedule < DurationInMillis_72) && (systemValue.Value == "0") {
		rseworker.ResendEmails(job)
		rseworker.setNumResendEmailSent("1")
	} else if elapsedTimeSinceSchedule > DurationInMillis_72 {
		rseworker.ResendEmails(job)
		rseworker.TearDown(job)
	}
}

func (rseworker *ResendInvitationEmailWorker) setJobSuccess(job *model.Job) {
	if err := rseworker.App.Srv().Jobs.SetJobSuccess(job); err != nil {
		mlog.Error("Worker: Failed to set success for job", mlog.String("worker", rseworker.name), mlog.String("job_id", job.Id), mlog.String("error", err.Error()))
		rseworker.setJobError(job, err)
	}
}

func (rseworker *ResendInvitationEmailWorker) setJobError(job *model.Job, appError *model.AppError) {
	if err := rseworker.App.Srv().Jobs.SetJobError(job, appError); err != nil {
		mlog.Error("Worker: Failed to set job error", mlog.String("worker", rseworker.name), mlog.String("job_id", job.Id), mlog.String("error", err.Error()))
	}
}

func (rseworker *ResendInvitationEmailWorker) cleanEmailData(emailStringData string) ([]string, error) {
	// emailStringData looks like this ["user1@gmail.com","user2@gmail.com"]
	emails := []string{}
	err := json.Unmarshal([]byte(emailStringData), &emails)
	if err != nil {
		return nil, err
	}

	return emails, nil
}

func (rseworker *ResendInvitationEmailWorker) removeAlreadyJoined(teamID string, emailList []string) []string {
	var notJoinedYet []string
	for _, email := range emailList {
		// check if the user with this email is on the system already
		user, appErr := rseworker.App.GetUserByEmail(email)
		if appErr != nil {
			notJoinedYet = append(notJoinedYet, email)
			continue
		}
		// now we check if they are part of the team already
		userID := []string{user.Id}
		members, appErr := rseworker.App.GetTeamMembersByIds(teamID, userID, nil)
		if len(members) == 0 || appErr != nil {
			notJoinedYet = append(notJoinedYet, email)
		}
	}

	return notJoinedYet
}

func (rseworker *ResendInvitationEmailWorker) setNumResendEmailSent(num string) {
	sysVar := &model.System{Name: model.NumberOfInviteEmailsSent, Value: num}
	if err := rseworker.App.Srv().Store.System().SaveOrUpdate(sysVar); err != nil {
		mlog.Error("Unable to save NUMBER_OF_INVITE_EMAIL_SENT", mlog.String("worker", rseworker.name), mlog.Err(err))
	}
}

func (rseworker *ResendInvitationEmailWorker) GetDurations(job *model.Job) (int64, int64, int64, int64) {
	scheduledAt, _ := strconv.ParseInt(job.Data["scheduledAt"], 10, 64)
	now := model.GetMillis()

	elapsedTimeSinceSchedule := now - scheduledAt

	duration_24 := os.Getenv("MM_RESEND_INVITATION_EMAIL_JOB_DURATION")
	DurationInMillis_24, parseError := strconv.ParseInt(duration_24, 10, 64)
	if parseError != nil {
		// default to 24 hours
		DurationInMillis_24 = TwentyFourHoursInMillis
	}

	duration_48 := os.Getenv("MM_RESEND_INVITATION_EMAIL_JOB_DURATION_48")
	DurationInMillis_48, parseError := strconv.ParseInt(duration_48, 10, 64)
	if parseError != nil {
		// default to 48 hours
		DurationInMillis_48 = FourtyEightHoursInMillis
	}

	duration_72 := os.Getenv("MM_RESEND_INVITATION_EMAIL_JOB_DURATION_72")
	DurationInMillis_72, parseError := strconv.ParseInt(duration_72, 10, 64)
	if parseError != nil {
		// default to 72 hours
		DurationInMillis_72 = SeventyTwoHoursInMillis
	}

	return elapsedTimeSinceSchedule, DurationInMillis_24, DurationInMillis_48, DurationInMillis_72

}

func (rseworker *ResendInvitationEmailWorker) TearDown(job *model.Job) {
	rseworker.App.Srv().Store.System().PermanentDeleteByName(model.NumberOfInviteEmailsSent)
	rseworker.setJobSuccess(job)
}

func (rseworker *ResendInvitationEmailWorker) ResendEmails(job *model.Job) {
	teamID := job.Data["teamID"]
	emailListData := job.Data["emailList"]

	emailList, err := rseworker.cleanEmailData(emailListData)
	if err != nil {
		appErr := model.NewAppError("worker: "+rseworker.name, "job_id: "+job.Id, nil, err.Error(), http.StatusInternalServerError)
		mlog.Error("Worker: Failed to clean emails string data", mlog.String("worker", rseworker.name), mlog.String("job_id", job.Id), mlog.String("error", appErr.Error()))
		rseworker.setJobError(job, appErr)
	}

	emailList = rseworker.removeAlreadyJoined(teamID, emailList)

	_, appErr := rseworker.App.InviteNewUsersToTeamGracefully(emailList, teamID, job.Data["senderID"])
	if appErr != nil {
		mlog.Error("Worker: Failed to send emails", mlog.String("worker", rseworker.name), mlog.String("job_id", job.Id), mlog.String("error", appErr.Error()))
		rseworker.setJobError(job, appErr)
	}
}
