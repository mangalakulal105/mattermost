// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package sharedchannel

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mattermost/mattermost-server/v5/mlog"
	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/services/remotecluster"
)

func (scs *Service) onReceiveSyncMessage(msg model.RemoteClusterMsg, rc *model.RemoteCluster, response remotecluster.Response) error {
	if msg.Topic != TopicSync {
		return fmt.Errorf("wrong topic, expected `%s`, got `%s`", TopicSync, msg.Topic)
	}

	if len(msg.Payload) == 0 {
		return errors.New("empty sync message")
	}

	var syncMessages []syncMsg

	if err := json.Unmarshal(msg.Payload, &syncMessages); err != nil {
		return fmt.Errorf("invalid sync message: %w", err)
	}

	scs.server.GetLogger().Log(mlog.LvlSharedChannelServiceDebug, "Batch of sync messages received",
		mlog.String("remote", rc.DisplayName),
		mlog.Int("sync_msg_count", len(syncMessages)),
	)

	return scs.processSyncMessages(syncMessages, rc, response)
}

func (scs *Service) processSyncMessages(syncMessages []syncMsg, rc *model.RemoteCluster, response remotecluster.Response) error {
	var channel *model.Channel
	postErrors := make([]string, 0)
	usersSyncd := make([]string, 0)
	var lastSyncAt int64
	var err error

	chanToTeamMap := make(map[string]*model.Team)

	for _, sm := range syncMessages {

		scs.server.GetLogger().Log(mlog.LvlSharedChannelServiceDebug, "Sync msg received",
			mlog.String("post_id", sm.PostId),
			mlog.String("channel_id", sm.ChannelId),
			mlog.Int("reaction_count", len(sm.Reactions)),
			mlog.Int("user_count", len(sm.Users)),
			mlog.Bool("has_post", sm.Post != nil),
		)

		if channel == nil {
			if channel, err = scs.server.GetStore().Channel().Get(sm.ChannelId, true); err != nil {
				// if the channel doesn't exist then none of these sync messages are going to work.
				return fmt.Errorf("channel not found processing sync messages: %w", err)
			}
		}

		if sm.Post != nil {
			if _, ok := chanToTeamMap[sm.Post.ChannelId]; !ok {
				team, err2 := scs.server.GetStore().Channel().GetTeamForChannel(sm.Post.ChannelId)
				if err2 != nil {
					scs.server.GetLogger().Log(mlog.LvlSharedChannelServiceError, "Error getting Team for Channel",
						mlog.String("ChannelId", sm.Post.ChannelId),
						mlog.String("PostId", sm.Post.Id),
						mlog.Err(err2),
					)
					postErrors = append(postErrors, sm.Post.Id)
					continue
				}
				chanToTeamMap[sm.Post.ChannelId] = team
			}
			sm.Post.Message = scs.processPermalinkFromRemote(sm.Post, chanToTeamMap[sm.Post.ChannelId])
		}

		// add/update users first
		for _, user := range sm.Users {
			if userSaved, err := scs.upsertSyncUser(user, rc); err != nil {
				scs.server.GetLogger().Log(mlog.LvlSharedChannelServiceError, "Error upserting sync user",
					mlog.String("post_id", sm.PostId),
					mlog.String("channel_id", sm.ChannelId),
					mlog.String("user_id", user.Id),
					mlog.Err(err))
			} else {
				usersSyncd = append(usersSyncd, userSaved.Id)
				scs.server.GetLogger().Log(mlog.LvlSharedChannelServiceDebug, "User added via sync",
					mlog.String("channel_id", sm.ChannelId),
					mlog.String("user_id", user.Id))
			}
		}

		// add post (may be nil if only reactions changed)
		if sm.Post != nil {
			rpost, err := scs.upsertSyncPost(sm.Post, channel, rc)
			if err != nil {
				postErrors = append(postErrors, sm.Post.Id)
				scs.server.GetLogger().Log(mlog.LvlSharedChannelServiceError, "Error upserting sync post",
					mlog.String("post_id", sm.Post.Id),
					mlog.String("channel_id", sm.Post.ChannelId),
					mlog.Err(err))
			} else if lastSyncAt < rpost.UpdateAt {
				lastSyncAt = rpost.UpdateAt
			}
		}

		// add/remove reactions
		for _, reaction := range sm.Reactions {
			if _, err := scs.upsertSyncReaction(reaction, rc); err != nil {
				scs.server.GetLogger().Log(mlog.LvlSharedChannelServiceError, "Error creating/deleting sync reaction",
					mlog.String("user_id", reaction.UserId),
					mlog.String("post_id", reaction.PostId),
					mlog.String("emoji", reaction.EmojiName),
					mlog.Int64("delete_at", reaction.DeleteAt),
					mlog.Err(err))
			} else {
				scs.server.GetLogger().Log(mlog.LvlSharedChannelServiceDebug, "Reaction upserted via sync",
					mlog.String("user_id", reaction.UserId),
					mlog.String("post_id", reaction.PostId),
					mlog.String("emoji", reaction.EmojiName),
					mlog.Int64("delete_at", reaction.DeleteAt))

				if lastSyncAt < reaction.UpdateAt {
					lastSyncAt = reaction.UpdateAt
				}
			}
		}
	}

	response[ResponseLastUpdateAt] = lastSyncAt // might be zero
	response[ResponsePostErrors] = postErrors   // might be empty
	response[ResponseUsersSynced] = usersSyncd  // might be empty

	return nil
}

func (scs *Service) upsertSyncUser(user *model.User, rc *model.RemoteCluster) (*model.User, error) {
	var err error
	var userSaved *model.User

	user.RemoteId = model.NewString(rc.RemoteId)

	// does the user already exist?
	euser, err := scs.server.GetStore().User().Get(user.Id)
	if err != nil {
		if _, ok := err.(errNotFound); !ok {
			return nil, fmt.Errorf("error checking sync user: %w", err)
		}
	}

	if euser == nil {
		if userSaved, err = scs.server.GetStore().User().Save(user); err != nil {
			if _, ok := err.(errConflict); !ok {
				return nil, fmt.Errorf("error inserting sync user: %w", err)
			}
			// probably a username or email collision
			// TODO: handle collision by modifying username/email (MM-32133)
			return nil, fmt.Errorf("username or email collision inserting sync user: %w", err)
		}
	} else {
		patch := &model.UserPatch{
			Nickname:  &user.Nickname,
			FirstName: &user.FirstName,
			LastName:  &user.LastName,
			Position:  &user.Position,
			Locale:    &user.Locale,
			Timezone:  user.Timezone,
		}
		euser.Patch(patch)
		userUpdated, err := scs.server.GetStore().User().Update(euser, false)
		if err != nil {
			return nil, fmt.Errorf("error updating sync user: %w", err)
		}
		userSaved = userUpdated.New
	}
	return userSaved, nil
}

func (scs *Service) upsertSyncPost(post *model.Post, channel *model.Channel, rc *model.RemoteCluster) (*model.Post, error) {
	var appErr *model.AppError

	post.RemoteId = model.NewString(rc.RemoteId)

	rpost, err := scs.server.GetStore().Post().GetSingle(post.Id)
	if err != nil {
		if _, ok := err.(errNotFound); !ok {
			return nil, fmt.Errorf("error checking sync post: %w", err)
		}
	}

	if rpost == nil {
		// post doesn't exist; create new one
		rpost, appErr = scs.app.CreatePost(post, channel, true, true)
		scs.server.GetLogger().Log(mlog.LvlSharedChannelServiceDebug, "Created sync post",
			mlog.String("post_id", post.Id),
			mlog.String("channel_id", post.ChannelId))
	} else {
		// update post
		rpost, appErr = scs.app.UpdatePost(post, false)
		scs.server.GetLogger().Log(mlog.LvlSharedChannelServiceDebug, "Updated sync post",
			mlog.String("post_id", post.Id),
			mlog.String("channel_id", post.ChannelId))
	}

	var rerr error
	if appErr != nil {
		rerr = errors.New(appErr.Error())
	}
	return rpost, rerr
}

func (scs *Service) upsertSyncReaction(reaction *model.Reaction, rc *model.RemoteCluster) (*model.Reaction, error) {
	savedReaction := reaction
	var appErr *model.AppError

	if reaction.DeleteAt == 0 {
		savedReaction, appErr = scs.app.SaveReactionForPost(reaction)
	} else {
		appErr = scs.app.DeleteReactionForPost(reaction)
	}

	var err error
	if appErr != nil {
		err = errors.New(appErr.Error())
	}
	return savedReaction, err
}
