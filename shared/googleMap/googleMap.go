package googleMap

import (
	"os"

	"googlemaps.github.io/maps"
)

var (
	Client *maps.Client
)

func SetupClient() {
	var err error
	Client, err = maps.NewClient(maps.WithAPIKey("AIzaSyCmua_JtLFnNux2uKsi1sACWNm_qrSxlBo"))
	if err != nil {
		os.Exit(0)
	}
}
