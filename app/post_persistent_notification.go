// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package app

import (
	"context"
	"net/http"
	"time"

	"github.com/mattermost/mattermost-server/v6/app/request"
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/shared/mlog"
	"github.com/mattermost/mattermost-server/v6/store"
	"github.com/pkg/errors"
)

// DeletePersistentNotificationsPost stops persistent notifications, if mentioned user(except post owner) reacts, reply or ack on the post.
// Post-owner can only delete the original post to stop the notifications, in which case "checkMentionedUser" must be "false" and "mentionedUserID" can be empty.
func (a *App) DeletePersistentNotificationsPost(c request.CTX, post *model.Post, mentionedUserID string, checkMentionedUser bool) *model.AppError {
	if !a.IsPersistentNotificationsEnabled() {
		c.Logger().Debug("DeletePersistentNotificationsPost: Persistent Notification feature is not enabled.")
		return nil
	}

	if posts, _, err := a.Srv().Store().PostPersistentNotification().Get(model.GetPersistentNotificationsPostsParams{PostID: post.Id}); err != nil {
		return model.NewAppError("DeletePersistentNotificationsPost", "app.post_priority.delete_persistent_notification_post.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
	} else if len(posts) == 0 {
		// Either the notification post already deleted or was never a notification post
		return nil
	}

	if !*a.Config().ServiceSettings.AllowPersistentNotificationsForGuests {
		user, nErr := a.Srv().Store().User().Get(context.Background(), c.Session().UserId)
		if nErr != nil {
			var nfErr *store.ErrNotFound
			switch {
			case errors.As(nErr, &nfErr):
				return model.NewAppError("DeletePersistentNotificationsPost", MissingAccountError, nil, "", http.StatusNotFound).Wrap(nErr)
			default:
				return model.NewAppError("DeletePersistentNotificationsPost", "app.user.get.app_error", nil, "", http.StatusInternalServerError).Wrap(nErr)
			}
		}
		if user.IsGuest() {
			c.Logger().Debug("DeletePersistentNotificationsPost: Persistent Notification feature is not enabled for guests.")
			return nil
		}
	}

	// if checkMentionedUser is "false" that would mean user is already authorized to delete the persistent-notification.
	isUserMentioned := !checkMentionedUser

	// post owner is not allowed to stop the persistent notifications via ack, reply or reaction.
	if checkMentionedUser && mentionedUserID != post.UserId {
		if err := a.forEachPersistentNotificationPost([]*model.Post{post}, func(_ *model.Post, _ *model.Channel, _ *model.Team, mentions *ExplicitMentions, _ model.UserMap, _ map[string]map[string]model.StringMap) error {
			if mentions.isUserMentioned(mentionedUserID) {
				isUserMentioned = true
			}
			return nil
		}); err != nil {
			return model.NewAppError("DeletePersistentNotificationsPost", "app.post_priority.delete_persistent_notification_post.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		}
	}

	if isUserMentioned {
		if err := a.Srv().Store().PostPersistentNotification().Delete([]string{post.Id}); err != nil {
			return model.NewAppError("DeletePersistentNotificationsPost", "app.post_priority.delete_persistent_notification_post.app_error", nil, "", http.StatusInternalServerError).Wrap(err)
		}
	}

	return nil
}

func (a *App) SendPersistentNotifications() error {
	notificationInterval := time.Duration(*a.Config().ServiceSettings.PersistentNotificationInterval) * time.Minute
	notificationMaxCount := int16(*a.Config().ServiceSettings.PersistentNotificationMaxCount)

	// fetch posts for which first notificationInterval duration has passed
	maxCreateAt := time.Now().Add(-notificationInterval).UnixMilli()

	pagination := model.CursorPagination{
		Direction: "down",
		PerPage:   500,
	}

	// Pagination loop
	for {
		notificationPosts, hasNext, err := a.Srv().Store().PostPersistentNotification().Get(model.GetPersistentNotificationsPostsParams{
			MaxCreateAt:   maxCreateAt,
			MaxLastSentAt: maxCreateAt,
			Pagination:    pagination,
		})
		if err != nil {
			return errors.Wrap(err, "failed to get posts for persistent notifications")
		}

		// No posts available at the moment for persistent notifications
		if len(notificationPosts) == 0 {
			return nil
		}
		pagination.FromID = notificationPosts[len(notificationPosts)-1].PostId
		pagination.FromCreateAt = notificationPosts[len(notificationPosts)-1].CreateAt

		notificationPostsMap := make(map[string]*model.PostPersistentNotifications, len(notificationPosts))
		postIds := make([]string, 0, len(notificationPosts))
		for _, p := range notificationPosts {
			postIds = append(postIds, p.PostId)
			notificationPostsMap[p.PostId] = p
		}
		posts, err := a.Srv().Store().Post().GetPostsByIds(postIds)
		if err != nil {
			return errors.Wrap(err, "failed to get posts by IDs")
		}

		var expiredPosts []*model.Post
		var validPosts []*model.Post
		for _, p := range posts {
			if notificationPostsMap[p.Id].SentCount >= notificationMaxCount {
				expiredPosts = append(expiredPosts, p)
			} else {
				validPosts = append(validPosts, p)
			}
		}

		// Delete expired notifications posts
		expiredPostsIds := make([]string, 0, len(expiredPosts))
		for _, p := range expiredPosts {
			expiredPostsIds = append(expiredPostsIds, p.Id)
		}
		if err := a.Srv().Store().PostPersistentNotification().Delete(expiredPostsIds); err != nil {
			return errors.Wrapf(err, "failed to delete expired notifications: %v", expiredPostsIds)
		}

		// Send notifications to validPosts
		if err := a.forEachPersistentNotificationPost(validPosts, a.sendPersistentNotifications); err != nil {
			return err
		}

		// Update last activity for valid notifications posts
		validPostsIds := make([]string, 0, len(validPosts))
		for _, p := range validPosts {
			validPostsIds = append(validPostsIds, p.Id)
		}
		if err := a.Srv().Store().PostPersistentNotification().UpdateLastActivity(validPostsIds); err != nil {
			return errors.Wrapf(err, "failed to update lastSentAt for valid notifications: %v", validPostsIds)
		}

		// Break pagination loop
		if !hasNext {
			break
		}
	}
	return nil
}

func (a *App) forEachPersistentNotificationPost(posts []*model.Post, fn func(post *model.Post, channel *model.Channel, team *model.Team, mentions *ExplicitMentions, profileMap model.UserMap, channelNotifyProps map[string]map[string]model.StringMap) error) error {
	channelIds := make(model.StringSet)
	for _, p := range posts {
		channelIds.Add(p.ChannelId)
	}
	channels, err := a.Srv().Store().Channel().GetChannelsByIds(channelIds.Val(), false)
	if err != nil {
		return errors.Wrap(err, "failed to get channels by IDs")
	}
	channelsMap := make(map[string]*model.Channel, len(channels))
	for _, c := range channels {
		channelsMap[c.Id] = c
	}

	teamIds := make(model.StringSet)
	for _, c := range channels {
		if c.TeamId != "" {
			teamIds.Add(c.TeamId)
		}
	}
	teams := make([]*model.Team, 0, len(teamIds))
	if len(teamIds) > 0 {
		teams, err = a.Srv().Store().Team().GetMany(teamIds.Val())
		if err != nil {
			return errors.Wrap(err, "failed to get teams by IDs")
		}
	}
	teamsMap := make(map[string]*model.Team, len(teams))
	for _, t := range teams {
		teamsMap[t.Id] = t
	}

	channelGroupMap := make(map[string]map[string]*model.Group, len(channelsMap))
	channelProfileMap := make(map[string]model.UserMap, len(channelsMap))
	channelKeywords := make(map[string]map[string][]string, len(channelsMap))
	channelNotifyProps := make(map[string]map[string]model.StringMap, len(channelsMap))
	for _, c := range channelsMap {
		if c.Type != model.ChannelTypeDirect {
			groups, err := a.getGroupsAllowedForReferenceInChannel(c, teamsMap[c.TeamId])
			if err != nil {
				return errors.Wrap(err, "failed to get groups for channels")
			}
			channelGroupMap[c.Id] = make(map[string]*model.Group, len(groups))
			for k, v := range groups {
				channelGroupMap[c.Id][k] = v
			}
			props, err := a.Srv().Store().Channel().GetAllChannelMembersNotifyPropsForChannel(c.Id, true)
			if err != nil {
				return errors.Wrap(err, "failed to get channel notify props")
			}
			channelNotifyProps[c.Id] = props
		}

		profileMap, err := a.Srv().Store().User().GetAllProfilesInChannel(context.Background(), c.Id, true)
		if err != nil {
			return errors.Wrapf(err, "failed to get profiles for channel %s", c.Id)
		}

		channelKeywords[c.Id] = make(map[string][]string, len(profileMap))
		validProfileMap := make(map[string]*model.User, len(profileMap))
		for k, v := range profileMap {
			if v.IsBot {
				continue
			}
			validProfileMap[k] = v
			channelKeywords[c.Id]["@"+v.Username] = []string{k}
		}
		channelProfileMap[c.Id] = validProfileMap
	}

	for _, post := range posts {
		channel := channelsMap[post.ChannelId]
		team := teamsMap[channel.TeamId]
		if channel.IsGroupOrDirect() {
			team = &model.Team{}
		}
		profileMap := channelProfileMap[channel.Id]

		mentions := &ExplicitMentions{}
		if channel.Type == model.ChannelTypeDirect {
			otherUserId := channel.GetOtherUserIdForDM(post.UserId)
			if _, ok := profileMap[otherUserId]; ok {
				mentions.addMention(otherUserId, DMMention)
			}
		} else {
			keywords := channelKeywords[channel.Id]
			mentions = getExplicitMentions(post, keywords, channelGroupMap[channel.Id])
			for _, group := range mentions.GroupMentions {
				_, err := a.insertGroupMentions(group, channel, profileMap, mentions)
				if err != nil {
					return errors.Wrapf(err, "failed to include mentions from group - %s for channel - %s", group.Id, channel.Id)
				}
			}
		}

		if err := fn(post, channel, team, mentions, profileMap, channelNotifyProps); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) sendPersistentNotifications(post *model.Post, channel *model.Channel, team *model.Team, mentions *ExplicitMentions, profileMap model.UserMap, channelNotifyProps map[string]map[string]model.StringMap) error {
	mentionedUsersList := make(model.StringArray, 0, len(mentions.Mentions))
	for id := range mentions.Mentions {
		// Don't send notification to post owner
		if id != post.UserId {
			mentionedUsersList = append(mentionedUsersList, id)
		}
	}

	sender := profileMap[post.UserId]
	notification := &PostNotification{
		Post:       post,
		Channel:    channel,
		ProfileMap: profileMap,
		Sender:     sender,
	}

	// Check for channel-wide mentions in channels that have too many members for those to work
	if int64(len(mentionedUsersList)) > *a.Config().TeamSettings.MaxNotificationsPerChannel {
		return errors.Errorf("mentioned users: %d are more than allowed users: %d", len(mentionedUsersList), *a.Config().TeamSettings.MaxNotificationsPerChannel)
	}

	if a.canSendPushNotifications() {
		for _, userID := range mentionedUsersList {
			user := profileMap[userID]
			if user == nil {
				continue
			}

			status, err := a.GetStatus(userID)
			if err != nil {
				mlog.Warn("Unable to fetch online status", mlog.String("user_id", userID), mlog.Err(err))
				status = &model.Status{UserId: userID, Status: model.StatusOffline, Manual: false, LastActivityAt: 0, ActiveChannel: ""}
			}

			if ShouldSendPushNotification(profileMap[userID], channelNotifyProps[channel.Id][userID], true, status, post) {
				a.sendPushNotification(
					notification,
					user,
					true,
					false,
					"",
				)
			} else {
				// register that a notification was not sent
				a.NotificationsLog().Debug("Persistent Notification not sent",
					mlog.String("ackId", ""),
					mlog.String("type", model.PushTypeMessage),
					mlog.String("userId", userID),
					mlog.String("postId", post.Id),
					mlog.String("status", model.PushNotSent),
				)
			}
		}
	}

	desktopUsers := make([]string, 0, len(mentionedUsersList))
	for _, id := range mentionedUsersList {
		user := profileMap[id]
		if user == nil {
			continue
		}

		if user.NotifyProps[model.DesktopNotifyProp] != model.UserNotifyNone && a.persistentNotificationsAllowedForStatus(id) {
			desktopUsers = append(desktopUsers, id)
		}
	}

	if len(desktopUsers) != 0 {
		post = a.PreparePostForClient(request.EmptyContext(a.Log()), post, false, false, true)
		postJSON, jsonErr := post.ToJSON()
		if jsonErr != nil {
			return errors.Wrapf(jsonErr, "failed to encode post to JSON")
		}

		for _, u := range desktopUsers {
			message := model.NewWebSocketEvent(model.WebsocketEventPersistentNotificationTriggered, team.Id, post.ChannelId, u, nil, "")

			message.Add("post", postJSON)
			message.Add("channel_type", channel.Type)
			message.Add("channel_display_name", notification.GetChannelName(model.ShowUsername, ""))
			message.Add("channel_name", channel.Name)
			message.Add("sender_name", notification.GetSenderName(model.ShowUsername, *a.Config().ServiceSettings.EnablePostUsernameOverride))
			message.Add("team_id", team.Id)

			if len(post.FileIds) != 0 {
				message.Add("otherFile", "true")

				infos, err := a.Srv().Store().FileInfo().GetForPost(post.Id, false, false, true)
				if err != nil {
					mlog.Warn("Unable to get fileInfo for push notifications.", mlog.String("post_id", post.Id), mlog.Err(err))
				}

				for _, info := range infos {
					if info.IsImage() {
						message.Add("image", "true")
						break
					}
				}
			}

			message.Add("mentions", model.ArrayToJSON(desktopUsers))
			a.Publish(message)
		}
	}

	return nil
}

func (a *App) persistentNotificationsAllowedForStatus(userID string) bool {
	var status *model.Status
	var err *model.AppError
	if status, err = a.GetStatus(userID); err != nil {
		status = &model.Status{UserId: userID, Status: model.StatusOffline, Manual: false, LastActivityAt: 0, ActiveChannel: ""}
	}

	return status.Status != model.StatusDnd && status.Status != model.StatusOutOfOffice
}

func (a *App) IsPersistentNotificationsEnabled() bool {
	return a.IsPostPriorityEnabled() && *a.Config().ServiceSettings.AllowPersistentNotifications
}
