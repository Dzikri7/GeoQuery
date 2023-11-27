package gcf

import (
	geoquery "github.com/Dzikri7/GeoQuery"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("Init", geoquery.PostGeoIntersects)
}
