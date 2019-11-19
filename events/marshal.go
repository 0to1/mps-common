package events

import "encoding/json"

func Marshal(msg interface{}) []byte {
	data, err := json.Marshal(msg)
	if err != nil {
		return nil
	}

	return data
}
