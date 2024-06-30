// Hand port of https://github.com/elastic/elasticsearch-specification/blob/76e25d34bff1060e300c95f4be468ef88e4f3465/specification/_types/mapping/dynamic-template.ts
/*
 * Licensed to Elasticsearch B.V. under one or more contributor
 * license agreements. See the NOTICE file distributed with
 * this work for additional information regarding copyright
 * ownership. Elasticsearch B.V. licenses this file to you under
 * the Apache License, Version 2.0 (the "License"); you may
 * not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package mapping

import (
	"fmt"

	"github.com/ngicks/und/sliceund"
)

type DynamicTemplate struct {
	Mapping          sliceund.Und[Property]  `json:"mapping,omitempty"`
	Match            sliceund.Und[string]    `json:"match,omitempty"`
	MatchMappingType sliceund.Und[string]    `json:"match_mapping_type,omitempty"`
	MatchPattern     sliceund.Und[MatchType] `json:"match_pattern,omitempty"`
	PathMatch        sliceund.Und[string]    `json:"path_match,omitempty"`
	PathUnmatch      sliceund.Und[string]    `json:"path_unmatch,omitempty"`
	Unmatch          sliceund.Und[string]    `json:"unmatch,omitempty"`
}

type MatchType string

const (
	Simple MatchType = "simple"
	Regex  MatchType = "regex"
)

type DynamicMapping string

func (m DynamicMapping) MarshalJSON() ([]byte, error) {
	switch m {
	case Strict, Runtime:
		return []byte("\"" + m + "\""), nil
	case True, False:
		return []byte(m), nil
	}
	return nil, fmt.Errorf("DynamicMapping with unknown value: %s", m)
}

func (m *DynamicMapping) UnmarshalJSON(data []byte) error {
	switch data[0] {
	default:
		return fmt.Errorf("unknown input: %s", data)
	case 't':
		if string(data) == "true" {
			*m = True
		}
	case 'f':
		if string(data) == "false" {
			*m = False
		}
	case '"':
		switch string(data[1 : len(data)-1]) {
		case string(Strict):
			*m = Strict
		case string(Runtime):
			*m = Runtime
		case string(True):
			*m = True
		case string(False):
			*m = False
		}
	}
	return nil
}

const (
	Strict  DynamicMapping = "strict"
	Runtime DynamicMapping = "runtime"
	True    DynamicMapping = "true"
	False   DynamicMapping = "false"
)
