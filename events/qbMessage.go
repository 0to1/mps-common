package events

import "encoding/json"

const TopicQBMessage = "go.micro.topic.qbmessage"

type QBMessage struct {
	TsID       uint16
	Priority   uint16
	Code       uint16
	Ikey       uint16
	Parameters []uint16
}

func (qb QBMessage) Marshal() []byte {
	data, err := json.Marshal(qb)
	if err != nil {
		return nil
	}

	return data
}

func (qb *QBMessage) Unmarshal(data []byte) *QBMessage {
	err := json.Unmarshal(data, qb)
	if err != nil {
		return nil
	}

	return qb
}

func (qb QBMessage) String() string {
	return "QBMessage"
}

func MarshalQBMessage(msg QBMessage) []byte {
	data, err := json.Marshal(msg)
	if err != nil {
		return nil
	}

	return data
}

func UnmarshalQBMessage(data []byte) *QBMessage {
	var msg QBMessage
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil
	}

	return &msg
}
