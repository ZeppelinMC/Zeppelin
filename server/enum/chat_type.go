package enum

const (
	ChatTypeRegular int32 = iota
	ChatTypeEmoteCommand
	ChatTypeMsgCommandIncoming
	ChatTypeMsgCommandOutgoing
	ChatTypeSayCommand
	ChatTypeTeamMsgCommandIncoming
	ChatTypeTeamMsgCommandOutgoing
)
