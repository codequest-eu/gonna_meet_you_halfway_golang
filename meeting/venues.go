package meeting

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"encoding/json"

	"github.com/codequest-eu/gonna_meet_you_halfway_golang/models"
)

var (
	fsClientID     = os.Getenv("FOURSQUARE_CLIENT_ID")
	fsClientSecret = os.Getenv("FOURSQUARE_CLIENT_SECRET")
)

func AskForVenues(middlePoint models.Position) (*[]models.Venue, error) {
	url := buildURL(middlePoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(res.Body)
	var venueResp map[string]interface{}
	if err := json.Unmarshal(body, &venueResp); err != nil {
		return nil, err
	}
	return getVenues(venueResp), nil
}

func buildURL(middlePoint models.Position) string {
	pointParam := fmt.Sprintf("%v,%v", middlePoint.Latitude, middlePoint.Longitude)
	return "https://api.foursquare.com/v2/venues/search?ll=" + pointParam + "&client_id=" + fsClientID + "&client_secret=" + fsClientSecret + "&v=20161020&m=foursquare&llAcc=500&query=caffee"
}

func getVenues(response map[string]interface{}) *[]models.Venue {
	var venues []models.Venue

	respDic := response["response"].(map[string]interface{})
	venuesArr := respDic["venues"].([]interface{})

	for _, v := range venuesArr {
		venue := v.(map[string]interface{})
		loc := venue["location"].(map[string]interface{})

		l := models.Position{
			Longitude: loc["lng"].(float64),
			Latitude:  loc["lat"].(float64),
		}
		venues = append(venues, models.Venue{
			Name:     venue["name"].(string),
			Location: l,
		})
	}
	return &venues
}
