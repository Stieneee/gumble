package gumble

import (
	"github.com/stieneee/gumble/gumble/proto/MumbleProto"
)

// AccessTokens are additional passwords that can be provided to the server to
// gain access to restricted channels.
type AccessTokens []string

func (a AccessTokens) writeMessage(client *Client) error {
	packet := MumbleProto.Authenticate{
		Tokens: a,
	}
	return client.Conn.WriteProto(&packet)
}
