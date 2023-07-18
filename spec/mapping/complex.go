// https://github.com/elastic/elasticsearch-specification/blob/76e25d34bff1060e300c95f4be468ef88e4f3465/specification/_types/mapping/complex.ts
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
	timeseriesmetrictype "github.com/ngicks/estype/spec/mapping/TimeSeriesMetricType"
	"github.com/ngicks/und/undefinedable"
)

type FlattenedProperty struct {
	PropertyBase
	Boost                    undefinedable.Undefinedable[float64]      `json:"boost"`
	DepthLimit               undefinedable.Undefinedable[int]          `json:"depth_limit"`
	DocValues                undefinedable.Undefinedable[bool]         `json:"doc_values"`
	EagerGlobalOrdinals      undefinedable.Undefinedable[bool]         `json:"eager_global_ordinals"`
	Index                    undefinedable.Undefinedable[bool]         `json:"index"`
	IndexOptions             undefinedable.Undefinedable[IndexOptions] `json:"index_options"`
	NullValue                undefinedable.Undefinedable[string]       `json:"null_value"`
	Similarity               undefinedable.Undefinedable[string]       `json:"similarity"`
	SplitQueriesOnWhitespace undefinedable.Undefinedable[bool]         `json:"split_queries_on_whitespace"`
	Type                     EsType                                    `json:"type"`
}

type NestedProperty struct {
	CorePropertyBase
	Enabled         undefinedable.Undefinedable[bool] `json:"enabled"`
	IncludeInParent undefinedable.Undefinedable[bool] `json:"include_in_parent"`
	IncludeInRoot   undefinedable.Undefinedable[bool] `json:"include_in_root"`
	Type            EsType                            `json:"type"`
}

type ObjectProperty struct {
	CorePropertyBase
	Enabled undefinedable.Undefinedable[bool]   `json:"enabled"`
	Type    undefinedable.Undefinedable[EsType] `json:"type"`
}

type DenseVectorProperty struct {
	PropertyBase
	Type         EsType                                               `json:"type"`
	Dims         int                                                  `json:"dims"`
	Similarity   undefinedable.Undefinedable[string]                  `json:"similarity"`
	Index        undefinedable.Undefinedable[bool]                    `json:"index"`
	IndexOptions undefinedable.Undefinedable[DenseVectorIndexOptions] `json:"index_options"`
}

type AggregateMetricDoubleProperty struct {
	PropertyBase
	Type             EsType                                                                 `json:"type"`
	DefaultMetric    string                                                                 `json:"default_metric"`
	Metrics          []string                                                               `json:"metrics"`
	TimeSeriesMetric undefinedable.Undefinedable[timeseriesmetrictype.TimeSeriesMetricType] `json:"time_series_metric"`
}
