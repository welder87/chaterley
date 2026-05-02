package manager

import "encoding/json"

type MessageDTO struct {
	Type    ActionCode      `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type BaseMessagePayload struct {
	AuthorID string `json:"authorId"`
	RoomID   string `json:"roomId"`
}

type BaseMessagePayloadWithContent struct {
	BaseMessagePayload
	Content string `json:"content"`
}

type SentMessagePayload struct {
	BaseMessagePayloadWithContent
}

type ReceivedMessagePayload struct {
	BaseMessagePayloadWithContent
}

type LeaveRoomMessagePayload struct {
	BaseMessagePayload
}

type JoinRoomMessagePayload struct {
	BaseMessagePayload
}
