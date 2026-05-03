package manager

type ActionCode int

const (
	Ping ActionCode = iota
	JoinRoom
	LeaveRoom
	SendMessage
	ReceiveMessage
)

func (k ActionCode) String() string {
	return actionByCode[k]
}

var (
	pingAction           = "Ping"
	joinRoomAction       = "JoinRoom"
	leaveRoomAction      = "LeaveRoom"
	sendMessageAction    = "SendMessage"
	receiveMessageAction = "ReceiveMessage"
)

var actionByCode = []string{
	Ping:           pingAction,
	JoinRoom:       joinRoomAction,
	LeaveRoom:      leaveRoomAction,
	SendMessage:    sendMessageAction,
	ReceiveMessage: receiveMessageAction,
}
