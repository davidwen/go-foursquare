package foursquare

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

type ClientConfig struct {
	ClientId     string
	ClientSecret string
}

var CONFIG ClientConfig

func TestExploreNear(t *testing.T) {
	api := getApi()
	resp := api.Explore(ExploreRequest{Near: "New York"})
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.GetVenues())
}

func TestExploreLatLng(t *testing.T) {
	api := getApi()
	resp := api.Explore(ExploreRequest{LatLng: "40.7,-74"})
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.GetVenues())
}

func getApi() *Api {
	config := getClientConfig()
	return NewApi(config.ClientId, config.ClientSecret)
}

func getClientConfig() ClientConfig {
	if CONFIG.ClientId == "" {
		r, _ := ioutil.ReadFile("api_test.cfg")
		json.Unmarshal(r, &CONFIG)
	}
	return CONFIG
}
