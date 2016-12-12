package storage

import (
	"io/ioutil"

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
