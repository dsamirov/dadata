package dadata // import "gopkg.in/webdeskltd/dadata.v2"

const constBaseSuggestURL = "https://suggestions.dadata.ru/suggestions/api/4_1/rs/"

var baseSuggestURL = constBaseSuggestURL

// GeoIPResponse response for GeoIP
type GeoIPResponse struct {
	Location *ResponseAddress `json:"location"`
}
