package model

type FollowRequest struct {
	SenderEmail   string `json:"senderEmail"`
	ReceiverEmail string `json:"receiverEmail"`
	//TimeOfRequestSending time.Time `json:"timeOfRequestSending"`
	//TimeOfRequestAnswer  time.Time `json:"timeOfRequestAnswer"`
}
