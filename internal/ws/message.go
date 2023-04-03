package ws

import "encoding/json"

type MessageFormat struct {
	Token      string // JWT authorization token of the current logged in user
	Data       any    // Data sent by the client
	ActionType string // Action to be performed by the server on the incoming data
}

/*
 * Message from the /ws endpoint is received as a string, (stringified JSON)
 * 1. It is Unmarshalled to the message format (MessageFormat), as given above
 * 2. If the the Unmarshalled object does not implement the same structure, it is discarded
 * 3. The token is verified, to check if the user is logged in
 * 4. Action type if the action to be performed by the server
 * 5. If the action type does not match any of the predefined action types, it is discarded
 * 6. The data is then sent to the appropriate channel, based on the Action type
 * 7. Return proper response to the client on the same channel
 */

func (m *MessageFormat) UnmarshalJSON(b []byte) error {
	// Unmarshal the message
	var msg MessageFormat
	err := json.Unmarshal(b, &msg)
	if err != nil {
		return err
	}
	return nil
}
