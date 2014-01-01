package foursquare

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	EXPLORE_URL = "https://api.foursquare.com/v2/venues/explore?"
)

type Api struct {
	clientId     string
	clientSecret string
}

func NewApi(clientId string, clientSecret string) *Api {
	api := new(Api)
	api.clientId = clientId
	api.clientSecret = clientSecret
	return api
}

func (self *Api) Explore(req ExploreRequest) (resp ExploreResponse) {
	params := self.clientParams()
	if req.LatLng != "" {
		params.Set("ll", req.LatLng)
	} else {
		params.Set("near", req.Near)
	}

	r, _ := http.Get(EXPLORE_URL + params.Encode())
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)
	dec.Decode(&resp)
	return resp
}

func (self *Api) clientParams() url.Values {
	params := url.Values{}
	params.Set("client_id", self.clientId)
	params.Set("client_secret", self.clientSecret)
	return params
}
