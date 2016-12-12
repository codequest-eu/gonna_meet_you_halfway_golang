package models

type StartData struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	OtherEmail string `json:"otherEmail"`
}

type NewMeeting struct {
	Identifier string `json:"meetingIdentifier"`
	Topics     Topics `json:"topics"`
}

type Topics struct {
	SuggestionsTopicName     string `json:"suggestionsTopicName"`
	MyLocationTopicName      string `json:"myLocationTopicName"`
	OtherLocationTopicName   string `json:"otherLocationTopicName"`
	MeetingLocationTopicName string `json:"meetingLocationTopicName"`
}
