package geoquery

import (
	"fmt"
	"testing"

	"github.com/Dzikri7/GeoQuery/helper"
)

func TestUpdateGetData(t *testing.T) {
	mconn := helper.SetConnection("MONGOSTRING", "geojson")
	datagedung := helper.GetAllBangunanLineString(mconn, "petapedia")
	fmt.Println(datagedung)
}
