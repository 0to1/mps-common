package events

import "encoding/json"

const TopicSMessage string = "go.micro.topic.smessage"

type SMessage struct {
	Index          uint16
	Magic1         uint16
	Magic2         uint16
	Magic3         uint16
	Ts             uint16
	OrderStatus    uint16
	CarrierNo      uint16
	CarrierStatus  uint16
	CarrierStation uint16
}

func (s SMessage) Marshal() []byte {
	data, err := json.Marshal(s)
	if err != nil {
		return nil
	}

	return data
}

func (s *SMessage) Unmarshal(data []byte) *SMessage {
	err := json.Unmarshal(data, s)
	if err != nil {
		return nil
	}

	return s
}

func (s SMessage) String() string {
	return "SMessage"
}

func MarshalSMessage(msg SMessage) []byte {
	data, err := json.Marshal(msg)
	if err != nil {
		return nil
	}

	return data
}

func UnmarshalSMessage(data []byte) *SMessage {
	var msg SMessage
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return nil
	}

	return &msg
}
