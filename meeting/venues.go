package meeting

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/codequest-eu/gonna_meet_you_halfway_golang/models"
)

var (
	gpAPIKey = os.Getenv("GOOGLE_DIRECTIONS_KEY")
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
	return "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=" + pointParam + "&radius=500&types=cafe|bar|restaurant&key=" + gpAPIKey
}

func getVenues(response map[string]interface{}) *[]models.Venue {
	var venues []models.Venue

	respDic := response["results"].([]interface{})
	for _, v := range respDic {
		venue := v.(map[string]interface{})
		g := venue["geometry"].(map[string]interface{})
		l := g["location"].(map[string]interface{})

		pos := models.Position{
			Longitude: l["lng"].(float64),
			Latitude:  l["lat"].(float64),
		}
		venues = append(venues, models.Venue{
			Name:     venue["name"].(string),
			Location: pos,
		})
	}
	return &venues
}
