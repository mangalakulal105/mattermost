package platform

import "github.com/mattermost/mattermost/server/public/model"

type BroadcastHook struct {
	HasChanges func(msg *model.WebSocketEvent, webConn *WebConn, args map[string]any) bool
	Process    func(msg *model.WebSocketEvent, webConn *WebConn, args map[string]any) *model.WebSocketEvent
}

func (h *Hub) runBroadcastHooks(msg *model.WebSocketEvent, webConn *WebConn) *model.WebSocketEvent {
	hookIDs := msg.GetBroadcast().BroadcastHooks

	if len(hookIDs) == 0 {
		return msg
	}

	// Check first if any hooks want to make changes to the event
	hasChanges := false

	for i, hookID := range hookIDs {
		hook := h.broadcastHooks[hookID]
		args := msg.GetBroadcast().BroadcastHookArgs[i]
		if hook == nil {
			continue
		}

		if hook.HasChanges(msg, webConn, args) {
			hasChanges = true
			break
		}
	}

	if !hasChanges {
		return msg
	}

	// Copy the event and remove any precomputed JSON since one or more hooks wants to make changes to it
	msg = msg.Copy()
	msg.RemovePrecomputedJSON()

	for i, hookID := range hookIDs {
		hook := h.broadcastHooks[hookID]
		args := msg.GetBroadcast().BroadcastHookArgs[i]
		if hook == nil {
			continue
		}

		hook.Process(msg, webConn, args)
	}

	return msg
}
