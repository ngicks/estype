package eshelper

import (
	"encoding/json"
	"fmt"

	"github.com/ngicks/estype/spec/indices/indexstate"
	"github.com/ngicks/estype/spec/mapping"
)

func GetMapping(bin []byte) (mapping.TypeMapping, error) {
	var err error

	mappingIndex := map[string]indexstate.IndexState{}
	err = json.Unmarshal(bin, &mappingIndex)
	if err == nil && len(mappingIndex) > 0 {
		var indexState indexstate.IndexState
		for _, v := range mappingIndex {
			indexState = v
			break
		}
		if len(indexState.Mappings.Value().Properties.Value()) > 0 {
			return indexState.Mappings.Value(), nil
		}
	}

	var indexState indexstate.IndexState
	err = json.Unmarshal(bin, &indexState)
	if err == nil && len(indexState.Mappings.Value().Properties.Value()) > 0 {
		return indexState.Mappings.Value(), nil
	}

	var typeMapping mapping.TypeMapping
	err = json.Unmarshal(bin, &typeMapping)
	if err == nil && len(typeMapping.Properties.Value()) > 0 {
		return typeMapping, nil
	}
	return mapping.TypeMapping{}, fmt.Errorf(
		"unknown mapping format. "+
			"The input must be one of data fetched from ${ELASTICSEARCH_URL}/${INDEX_NAME}/_mapping, IndexState or Mapping. %s",
		bin,
	)
}
