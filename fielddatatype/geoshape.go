package fielddatatype

import (
	"encoding/json"
	"fmt"

	"github.com/go-spatial/geom"
	"github.com/go-spatial/geom/encoding/geojson"
	"github.com/go-spatial/geom/encoding/wkt"
)

// TODO: test against elasticsearch 8.4
type GeoShape struct {
	// avoid embedding this because it could confuse user when re-defined type is made in user code.
	Geometry geom.Geometry
}

func (g *GeoShape) UnmarshalJSON(data []byte) error {
	if len(data) < 3 {
		return fmt.Errorf("too short: input must longer than 2 chars but %d", len(data))
	}
	switch data[0] {
	case '{':
		var geo geojson.Geometry
		err := geo.UnmarshalJSON(data)
		if err != nil {
			return err
		}
		g.Geometry = geo.Geometry
		return nil
	case '"':
		geo, err := wkt.DecodeBytes(data[1 : len(data)-1])
		if err != nil {
			return err
		}
		g.Geometry = geo
		return nil
	}

	return fmt.Errorf(
		"unknown type: must be geojson or geohash string literal, but was %s",
		string(data),
	)
}

func (g GeoShape) MarshalJSON() ([]byte, error) {
	return json.Marshal(geojson.Geometry{Geometry: g.Geometry})
}
