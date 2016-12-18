package meeting

import (
	"fmt"
	"os"

	"github.com/codequest-eu/gonna_meet_you_halfway_golang/models"
	"github.com/kellydunn/golang-geo"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
)

var gdAPIKey = os.Getenv("GOOGLE_DIRECTIONS_KEY")

type GService struct {
	client *maps.Client
}

func NewGService() (*GService, error) {
	o := maps.WithAPIKey(gdAPIKey)
	client, err := maps.NewClient(o)
	if err != nil {
		return nil, err
	}
	return &GService{client}, nil
}

func (gs *GService) CalculateMiddlePoint(locationA models.Position, locationB models.Position) (*models.Position, error) {
	pointParamA := fmt.Sprintf("%v,%v", locationA.Latitude, locationA.Longitude)
	pointParamB := fmt.Sprintf("%v,%v", locationB.Latitude, locationB.Longitude)
	r := &maps.DirectionsRequest{
		Origin:       pointParamA,
		Destination:  pointParamB,
		Alternatives: false,
		Mode:         maps.TravelModeWalking,
	}
	routes, _, err := gs.client.Directions(context.Background(), r)
	if err != nil {
		return nil, err
	}

	totalDistance := calculateTotalDistance(routes)
	middlePoint := evaluateMiddlePoint(routes, totalDistance)
	return &middlePoint, nil
}

func (gs *GService) AskForPlaces(middlePoint models.Position) (*[]models.Venue, error) {
	var venues []models.Venue
	placesTypes := []maps.PlaceType{maps.PlaceTypeCafe, maps.PlaceTypeBar, maps.PlaceTypeRestaurant}

	for _, t := range placesTypes {
		req := maps.RadarSearchRequest {
			Location: &maps.LatLng{middlePoint.Latitude, middlePoint.Longitude},
			Radius: 200,
			Type: t,
		}
		res, err := gs.client.RadarSearch(context.Background(), &req)
		if err != nil {
			return nil, err
		}
		for _, p := range res.Results {
			pdReq := maps.PlaceDetailsRequest{p.PlaceID, ""}
			d, err := gs.client.PlaceDetails(context.Background(), &pdReq)
			if err != nil {
				return nil, err
			}
			venue := models.Venue{
				d.Name,
				"",
				models.Position{
					d.Geometry.Location.Lat,
					d.Geometry.Location.Lng,
				},
			}
			venues = append(venues, venue)
		}
	}
	return &venues, nil
}

func calculateTotalDistance(routes []maps.Route) int {
	var totalDistance int
	route := routes[0]
	leg := route.Legs[0]
	for _, step := range leg.Steps {
		totalDistance += step.Distance.Meters
	}
	return totalDistance
}

func evaluateMiddlePoint(routes []maps.Route, totalDistance int) models.Position {
	var middlePoint models.Position
	var middlePointDistance int
	route := routes[0]
	leg := route.Legs[0]
	for _, step := range leg.Steps {
		middlePointDistance += step.Distance.Meters
		if middlePointDistance >= totalDistance/2 {
			endPoint := geo.NewPoint(step.EndLocation.Lat, step.EndLocation.Lng)
			startPoint := geo.NewPoint(step.StartLocation.Lat, step.StartLocation.Lng)
			bearing := endPoint.BearingTo(startPoint)
			distanceInM := float64(middlePointDistance - totalDistance/2)
			distanceInKm := distanceInM / 1000
			newMiddlePoint := endPoint.PointAtDistanceAndBearing(distanceInKm, bearing)
			middlePoint = models.Position{
				Longitude: newMiddlePoint.Lng(),
				Latitude:  newMiddlePoint.Lat(),
			}
			break
		}
	}
	return middlePoint
}
