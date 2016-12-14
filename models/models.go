package models

type InviteData struct {
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	OtherEmail string   `json:"otherEmail"`
	Location   Position `json:"position"`
}

type AcceptData struct {
	Name     string   `json:"name"`
	Location Position `json:"position"`
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

type MeetingSuggestion struct {
	LocationA Position `json:"locationA"`
	LocationB Position `json:"locationB"`
	Center    Position `json:"center"`
	Venues    []Venue  `json:"venues"`
}

func (ms *MeetingSuggestion) SetLocationB(position Position) {
	ms.LocationB = position
}

func (ms *MeetingSuggestion) SetCenter(center Position) {
	ms.Center = center
}

func (ms *MeetingSuggestion) SetVenues(venues []Venue) {
	ms.Venues = venues
}

type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Venue struct {
	Name     string   `json:"name"`
	Location Position `json:"position"`
}
