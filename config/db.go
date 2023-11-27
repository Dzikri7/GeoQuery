package config

import "github.com/Dzikri7/GeoQuery/helper"

var Mongocon = helper.SetConnection(Mongostring, "geojson")
