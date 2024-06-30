// Hand port of https://github.com/elastic/elasticsearch-specification/blob/76e25d34bff1060e300c95f4be468ef88e4f3465/specification/_types/mapping/specialized.ts
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
	mappingdef "github.com/ngicks/estype/spec"
	"github.com/ngicks/und/sliceund"
)

type CompletionProperty struct {
	DocValuesPropertyBase

	Analyzer                   sliceund.Und[string]           `json:"analyzer,omitempty"`
	Contexts                   sliceund.Und[[]SuggestContext] `json:"contexts,omitempty"`
	MaxInputLength             sliceund.Und[int]              `json:"max_input_length,omitempty"`
	PreservePositionIncrements sliceund.Und[bool]             `json:"preserve_position_increments,omitempty"`
	PreserveSeparators         sliceund.Und[bool]             `json:"preserve_separators,omitempty"`
	SearchAnalyzer             sliceund.Und[string]           `json:"search_analyzer,omitempty"`
	Type                       EsType                         `json:"type"`
}

type SuggestContext struct {
	Name      string               `json:"name"`
	Path      sliceund.Und[string] `json:"path,omitempty"`
	Type      string               `json:"type"`
	Precision sliceund.Und[any]    `json:"precision,omitempty"` // int or string
}

type ConstantKeywordProperty struct {
	PropertyBase

	Value sliceund.Und[any] `json:"value,omitempty"`
	Type  EsType            `json:"type"`
}

type FieldAliasProperty struct {
	PropertyBase

	Path sliceund.Und[string] `json:"path,omitempty"`
	Type EsType               `json:"type"`
}

type HistogramProperty struct {
	PropertyBase

	IgnoreMalformed sliceund.Und[bool] `json:"ignore_malformed,omitempty"`
	Type            EsType             `json:"type"`
}

type IpProperty struct {
	DocValuesPropertyBase

	Boost           sliceund.Und[float64]           `json:"boost,omitempty"`
	Index           sliceund.Und[bool]              `json:"index,omitempty"`
	IgnoreMalformed sliceund.Und[bool]              `json:"ignore_malformed,omitempty"`
	NullValue       sliceund.Und[string]            `json:"null_value,omitempty"`
	OnScriptError   sliceund.Und[OnScriptError]     `json:"on_script_error,omitempty"`
	Script          sliceund.Und[mappingdef.Script] `json:"script,omitempty"`
	/**
	 * For internal use by Elastic only. Marks the field as a time series dimension. Defaults to false.
	 * @availability stack stability=experimental
	 * @availability serverless stability=experimental
	 */
	TimeSeriesDimension sliceund.Und[bool] `json:"time_series_dimension,omitempty"`
	Type                EsType             `json:"type"`
}

type Murmur3HashProperty struct {
	DocValuesPropertyBase

	Type EsType `json:"type"`
}

type TokenCountProperty struct {
	DocValuesPropertyBase
	Analyzer                 sliceund.Und[string]  `json:"analyzer,omitempty"`
	Boost                    sliceund.Und[float64] `json:"boost,omitempty"`
	Index                    sliceund.Und[bool]    `json:"index,omitempty"`
	NullValue                sliceund.Und[float64] `json:"null_value,omitempty"`
	EnablePositionIncrements sliceund.Und[bool]    `json:"enable_position_increments,omitempty"`
	Type                     EsType                `json:"type"`
}
