package foursquare

// https://developer.foursquare.com/docs/venues/explore
type ExploreRequest struct {
	LatLng string
	Near   string
}

type ExploreResponse struct {
	Response ExploreBody
}

type ExploreBody struct {
	Groups []ExploreGroup
}

type ExploreGroup struct {
	Items []ExploreItem
}

type ExploreItem struct {
	Venue Venue
}

func (resp *ExploreResponse) GetVenues() (venues []Venue) {
	for _, group := range resp.Response.Groups {
		for _, item := range group.Items {
			venues = append(venues, item.Venue)
		}
	}
	return venues
}
