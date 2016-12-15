package storage

import (
	"io/ioutil"

	"github.com/codequest-eu/gonna_meet_you_halfway_golang/models"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type googleStorage struct {
	client *datastore.Client
}

func NewGoogleStorage(keyPath string, projectID string) (Store, error) {
	data, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}
	conf, err := google.JWTConfigFromJSON(
		data,
		datastore.ScopeDatastore,
	)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	client, err := datastore.NewClient(
		ctx,
		projectID,
		option.WithTokenSource(conf.TokenSource(ctx)),
	)
	if err != nil {
		return nil, err
	}
	return &googleStorage{client}, nil
}

func (gs *googleStorage) GetTopics(meetingID string) (models.Topics, error) {
	var topics models.Topics
	return topics, gs.client.Get(context.Background(), keyFor("Topics", meetingID, nil), &topics)
}

func (gs *googleStorage) SaveTopics(meetingID string, topics models.Topics) error {
	_, err := gs.client.Put(context.Background(), keyFor("Topics", meetingID, nil), &topics)
	return err
}

func (gs *googleStorage) GetMeetingSuggestion(meetingID string) (models.MeetingSuggestion, error) {
	var meetingSuggestion models.MeetingSuggestion
	return meetingSuggestion, gs.client.Get(context.Background(), keyFor("MeetingSuggestion", meetingID, nil), &meetingSuggestion)
}

func (gs *googleStorage) SaveMeetingSuggestion(meetingID string, meetingSuggestion models.MeetingSuggestion) error {
	_, err := gs.client.Put(context.Background(), keyFor("MeetingSuggestion", meetingID, nil), &meetingSuggestion)
	return err
}

func (gs *googleStorage) GetPlaceSuggestion(placeSuggestionID string) (models.PlaceSuggestion, error) {
	var placeSuggestion []models.PlaceSuggestion
	q := datastore.NewQuery("PlaceSuggestion").Filter("PlaceIdentifier=", placeSuggestionID).Limit(1)
	_, err := gs.client.GetAll(context.Background(), q, &placeSuggestion)
	return placeSuggestion[0], err
}

func (gs *googleStorage) SavePlaceSuggestion(meetingID string, placeSuggestion models.PlaceSuggestion) error {
	_, err := gs.client.Put(context.Background(), keyFor("PlaceSuggestion", meetingID, nil), &placeSuggestion)
	return err
}

func keyFor(kind string, hashID string, parent *datastore.Key) *datastore.Key {
	return datastore.NameKey(
		kind,   // kind
		hashID, // name
		parent, // parent
	)
}
