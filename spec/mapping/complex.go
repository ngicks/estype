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
	"github.com/ngicks/und/sliceund"
)

type FlattenedProperty struct {
	PropertyBase
	Boost                    sliceund.Und[float64]      `json:"boost,omitempty"`
	DepthLimit               sliceund.Und[int]          `json:"depth_limit,omitempty"`
	DocValues                sliceund.Und[bool]         `json:"doc_values,omitempty"`
	EagerGlobalOrdinals      sliceund.Und[bool]         `json:"eager_global_ordinals,omitempty"`
	Index                    sliceund.Und[bool]         `json:"index,omitempty"`
	IndexOptions             sliceund.Und[IndexOptions] `json:"index_options,omitempty"`
	NullValue                sliceund.Und[string]       `json:"null_value,omitempty"`
	Similarity               sliceund.Und[string]       `json:"similarity,omitempty"`
	SplitQueriesOnWhitespace sliceund.Und[bool]         `json:"split_queries_on_whitespace,omitempty"`
	Type                     EsType                     `json:"type"`
}

type NestedProperty struct {
	CorePropertyBase
	Enabled         sliceund.Und[bool] `json:"enabled,omitempty"`
	IncludeInParent sliceund.Und[bool] `json:"include_in_parent,omitempty"`
	IncludeInRoot   sliceund.Und[bool] `json:"include_in_root,omitempty"`
	Type            EsType             `json:"type"`
}

type ObjectProperty struct {
	CorePropertyBase
	Enabled sliceund.Und[bool]   `json:"enabled,omitempty"`
	Type    sliceund.Und[EsType] `json:"type,omitempty"`
}

type DenseVectorProperty struct {
	PropertyBase
	Type         EsType                                `json:"type"`
	Dims         int                                   `json:"dims"`
	Similarity   sliceund.Und[string]                  `json:"similarity,omitempty"`
	Index        sliceund.Und[bool]                    `json:"index,omitempty"`
	IndexOptions sliceund.Und[DenseVectorIndexOptions] `json:"index_options,omitempty"`
}

type AggregateMetricDoubleProperty struct {
	PropertyBase
	Type             EsType                                                  `json:"type"`
	DefaultMetric    string                                                  `json:"default_metric"`
	Metrics          []string                                                `json:"metrics"`
	TimeSeriesMetric sliceund.Und[timeseriesmetrictype.TimeSeriesMetricType] `json:"time_series_metric,omitempty"`
}
