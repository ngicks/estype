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
	"github.com/ngicks/und/undefinedable"
)

type CompletionProperty struct {
	DocValuesPropertyBase

	Analyzer                   undefinedable.Undefinedable[string]           `json:"analyzer"`
	Contexts                   undefinedable.Undefinedable[[]SuggestContext] `json:"contexts"`
	MaxInputLength             undefinedable.Undefinedable[int]              `json:"max_input_length"`
	PreservePositionIncrements undefinedable.Undefinedable[bool]             `json:"preserve_position_increments"`
	PreserveSeparators         undefinedable.Undefinedable[bool]             `json:"preserve_separators"`
	SearchAnalyzer             undefinedable.Undefinedable[string]           `json:"search_analyzer"`
	Type                       EsType                                        `json:"type"`
}

type SuggestContext struct {
	Name      string                              `json:"name"`
	Path      undefinedable.Undefinedable[string] `json:"path"`
	Type      string                              `json:"type"`
	Precision undefinedable.Undefinedable[any]    `json:"precision"` // int or string
}

type ConstantKeywordProperty struct {
	PropertyBase

	Value undefinedable.Undefinedable[any] `json:"value"`
	Type  EsType                           `json:"type"`
}

type FieldAliasProperty struct {
	PropertyBase

	Path undefinedable.Undefinedable[string] `json:"path"`
	Type EsType                              `json:"type"`
}

type HistogramProperty struct {
	PropertyBase

	IgnoreMalformed undefinedable.Undefinedable[bool] `json:"ignore_malformed"`
	Type            EsType                            `json:"type"`
}

type IpProperty struct {
	DocValuesPropertyBase

	Boost           undefinedable.Undefinedable[float64]           `json:"boost"`
	Index           undefinedable.Undefinedable[bool]              `json:"index"`
	IgnoreMalformed undefinedable.Undefinedable[bool]              `json:"ignore_malformed"`
	NullValue       undefinedable.Undefinedable[string]            `json:"null_value"`
	OnScriptError   undefinedable.Undefinedable[OnScriptError]     `json:"on_script_error"`
	Script          undefinedable.Undefinedable[mappingdef.Script] `json:"script"`
	/**
	 * For internal use by Elastic only. Marks the field as a time series dimension. Defaults to false.
	 * @availability stack stability=experimental
	 * @availability serverless stability=experimental
	 */
	TimeSeriesDimension undefinedable.Undefinedable[bool] `json:"time_series_dimension"`
	Type                EsType                            `json:"type"`
}

type Murmur3HashProperty struct {
	DocValuesPropertyBase

	Type EsType `json:"type"`
}

type TokenCountProperty struct {
	DocValuesPropertyBase
	Analyzer                 undefinedable.Undefinedable[string]  `json:"analyzer"`
	Boost                    undefinedable.Undefinedable[float64] `json:"boost"`
	Index                    undefinedable.Undefinedable[bool]    `json:"index"`
	NullValue                undefinedable.Undefinedable[float64] `json:"null_value"`
	EnablePositionIncrements undefinedable.Undefinedable[bool]    `json:"enable_position_increments"`
	Type                     EsType                               `json:"type"`
}
