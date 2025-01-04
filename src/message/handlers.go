package message

import pb_shared "mist-io/src/protos/frontend/v1"

func (wsc WsConnection) UpdateStoredJwtToken(message *pb_shared.Input_UpdateJwtToken) {
	wsc.JwtToken = message.UpdateJwtToken.Access
}
