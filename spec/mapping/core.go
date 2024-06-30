// Hand port of https://github.com/elastic/elasticsearch-specification/blob/76e25d34bff1060e300c95f4be468ef88e4f3465/specification/_types/mapping/core.ts
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
	"github.com/ngicks/estype/spec"
	"github.com/ngicks/estype/spec/indices"
	timeseriesmetrictype "github.com/ngicks/estype/spec/mapping/TimeSeriesMetricType"
	"github.com/ngicks/und/sliceund"
	"github.com/ngicks/und/sliceund/elastic"
)

type CorePropertyBase struct {
	PropertyBase
	CopyTo     elastic.Elastic[string] `json:"copy_to,omitempty"`
	Similarity sliceund.Und[string]    `json:"similarity,omitempty"`
	Store      sliceund.Und[bool]      `json:"store,omitempty"`
}

type DocValuesPropertyBase struct {
	CorePropertyBase
	DocValues sliceund.Und[bool] `json:"doc_values,omitempty"`
}

type BinaryProperty struct {
	DocValuesPropertyBase
	Type EsType `json:"type"`
}

type BooleanProperty struct {
	DocValuesPropertyBase
	Boost     sliceund.Und[float64]                  `json:"boost,omitempty"`
	Fielddata sliceund.Und[indices.NumericFielddata] `json:"fielddata,omitempty"`
	Index     sliceund.Und[bool]                     `json:"index,omitempty"`
	NullValue sliceund.Und[bool]                     `json:"null_value,omitempty"`
	Type      EsType                                 `json:"type"`
}

type DateProperty struct {
	DocValuesPropertyBase

	Boost sliceund.Und[float64] `json:"boost,omitempty"`

	Fielddata sliceund.Und[indices.NumericFielddata] `json:"fielddata,omitempty"`

	Format          sliceund.Und[string] `json:"format,omitempty"`
	IgnoreMalformed sliceund.Und[bool]   `json:"ignore_malformed,omitempty"`
	Index           sliceund.Und[bool]   `json:"index,omitempty"`
	NullValue       sliceund.Und[any]    `json:"null_value,omitempty"` // date string formatted as ones specified in format field or epoch millis int
	PrecisionStep   sliceund.Und[int]    `json:"precision_step,omitempty"`
	Locale          sliceund.Und[string] `json:"locale,omitempty"`
	Type            EsType               `json:"type"`
}

type DateNanosProperty struct {
	DocValuesPropertyBase

	Boost           sliceund.Und[float64] `json:"boost,omitempty"`
	Format          sliceund.Und[string]  `json:"format,omitempty"`
	IgnoreMalformed sliceund.Und[bool]    `json:"ignore_malformed,omitempty"`
	Index           sliceund.Und[bool]    `json:"index,omitempty"`
	NullValue       sliceund.Und[any]     `json:"null_value,omitempty"` // date string formatted as ones specified in format field or epoch millis int
	PrecisionStep   sliceund.Und[int]     `json:"precision_step,omitempty"`
	Type            EsType                `json:"type"`
}

type JoinProperty struct {
	PropertyBase

	Relations           sliceund.Und[map[string]elastic.Elastic[string]] `json:"relations,omitempty,omitempty"`
	EagerGlobalOrdinals sliceund.Und[bool]                               `json:"eager_global_ordinals,omitempty"`
	Type                EsType                                           `json:"type"`
}

type KeywordProperty struct {
	DocValuesPropertyBase

	Boost                    sliceund.Und[float64]      `json:"boost,omitempty"`
	EagerGlobalOrdinals      sliceund.Und[bool]         `json:"eager_global_ordinals,omitempty"`
	Index                    sliceund.Und[bool]         `json:"index,omitempty"`
	IndexOptions             sliceund.Und[IndexOptions] `json:"index_options,omitempty"`
	Normalizer               sliceund.Und[string]       `json:"normalizer,omitempty"`
	Norms                    sliceund.Und[bool]         `json:"norms,omitempty"`
	NullValue                sliceund.Und[string]       `json:"null_value,omitempty"`
	SplitQueriesOnWhitespace sliceund.Und[bool]         `json:"split_queries_on_whitespace,omitempty"`
	/**
	 * For internal use by Elastic only. Marks the field as a time series dimension. Defaults to false.
	 * @availability stack stability=experimental
	 * @availability serverless stability=experimental
	 */
	TimeSeriesDimension sliceund.Und[bool] `json:"time_series_dimension,omitempty"`
	Type                EsType             `json:"type"`
}

type NumberPropertyBase struct {
	DocValuesPropertyBase

	Boost           sliceund.Und[float64]       `json:"boost,omitempty"`
	Coerce          sliceund.Und[bool]          `json:"coerce,omitempty"`
	IgnoreMalformed sliceund.Und[bool]          `json:"ignore_malformed,omitempty"`
	Index           sliceund.Und[bool]          `json:"index,omitempty"`
	OnScriptError   sliceund.Und[OnScriptError] `json:"on_script_error,omitempty"`
	Script          sliceund.Und[spec.Script]   `json:"script,omitempty"`
	/**
	 * For internal use by Elastic only. Marks the field as a time series dimension. Defaults to false.
	 * @availability stack stability=experimental
	 * @availability serverless stability=experimental
	 */
	TimeSeriesMetric sliceund.Und[timeseriesmetrictype.TimeSeriesMetricType] `json:"time_series_metric,omitempty"`
	/**
	 * For internal use by Elastic only. Marks the field as a time series dimension. Defaults to false.
	 * @server_default false
	 * @availability stack stability=experimental
	 * @availability serverless stability=experimental
	 */
	TimeSeriesDimension sliceund.Und[bool] `json:"time_series_dimension,omitempty"`
}

type OnScriptError string

const (
	Fail     OnScriptError = "fail"
	Continue OnScriptError = "continue"
)

type FloatNumberProperty struct {
	NumberPropertyBase

	Type      EsType                `json:"type"`
	NullValue sliceund.Und[float32] `json:"null_value,omitempty"`
}

type HalfFloatNumberProperty struct {
	NumberPropertyBase

	Type      EsType                `json:"type"`
	NullValue sliceund.Und[float32] `json:"null_value,omitempty"` // TODO: use float16 package?
}

type DoubleNumberProperty struct {
	NumberPropertyBase

	Type      EsType                `json:"type"`
	NullValue sliceund.Und[float64] `json:"null_value,omitempty"`
}

type IntegerNumberProperty struct {
	NumberPropertyBase

	Type      EsType              `json:"type"`
	NullValue sliceund.Und[int32] `json:"null_value,omitempty"`
}

type LongNumberProperty struct {
	NumberPropertyBase

	Type      EsType              `json:"type"`
	NullValue sliceund.Und[int64] `json:"null_value,omitempty"`
}

type ShortNumberProperty struct {
	NumberPropertyBase

	Type      EsType              `json:"type"`
	NullValue sliceund.Und[int16] `json:"null_value,omitempty"`
}

type ByteNumberProperty struct {
	NumberPropertyBase

	Type      EsType             `json:"type"`
	NullValue sliceund.Und[int8] `json:"null_value,omitempty"`
}

type UnsignedLongNumberProperty struct {
	NumberPropertyBase

	Type      EsType               `json:"type"`
	NullValue sliceund.Und[uint64] `json:"null_value,omitempty"`
}

type ScaledFloatNumberProperty struct {
	NumberPropertyBase

	Type          EsType                `json:"type"`
	NullValue     sliceund.Und[float64] `json:"null_value,omitempty"`
	ScalingFactor sliceund.Und[float64] `json:"scaling_factor,omitempty"`
}

type PercolatorProperty struct {
	PropertyBase

	Type EsType `json:"type"`
}

type RankFeatureProperty struct {
	PropertyBase

	PositiveScoreImpact sliceund.Und[bool] `json:"positive_score_impact,omitempty"`
	Type                EsType             `json:"type"`
}

type RankFeaturesProperty struct {
	PropertyBase

	Type EsType `json:"type"`
}

type SearchAsYouTypeProperty struct {
	CorePropertyBase

	Analyzer            sliceund.Und[string]           `json:"analyzer,omitempty"`
	Index               sliceund.Und[bool]             `json:"index,omitempty"`
	IndexOptions        sliceund.Und[IndexOptions]     `json:"index_options,omitempty"`
	MaxShingleSize      sliceund.Und[int]              `json:"max_shingle_size,omitempty"`
	Norms               sliceund.Und[bool]             `json:"norms,omitempty"`
	SearchAnalyzer      sliceund.Und[string]           `json:"search_analyzer,omitempty"`
	SearchQuoteAnalyzer sliceund.Und[string]           `json:"search_quote_analyzer,omitempty"`
	TermVector          sliceund.Und[TermVectorOption] `json:"term_vector,omitempty"`
	Type                EsType                         `json:"type"`
}

// MatchOnlyTextProperty is an example of a property which does not derive from PropertyBase.
// We have checked and this property does not support all properties of the base type.
// In a future iteration we may remodel properties and identify truly common properties that should form
// a base type that can be considered a common ancestor for all properties. Some clients will generate
// a synthetic version of this today.

/**
 * A variant of text that trades scoring and efficiency of positional queries for space efficiency. This field
 * effectively stores data the same way as a text field that only indexes documents (index_options: docs) and
 * disables norms (norms: false). Term queries perform as fast if not faster as on text fields, however queries
 * that need positions such as the match_phrase query perform slower as they need to look at the _source document
 * to verify whether a phrase matches. All queries return constant scores that are equal to 1.0.
 */
type MatchOnlyTextProperty struct {
	Type EsType `json:"type"`
	/**
	 * Multi-fields allow the same string value to be indexed in multiple ways for different purposes, such as one
	 * field for search and a multi-field for sorting and aggregations, or the same string value analyzed by different analyzers.
	 * @doc_id multi-fields
	 */
	Fields sliceund.Und[map[string]Property] `json:"fields,omitempty"`
	/**
	 * Metadata about the field.
	 * @doc_id mapping-meta-field
	 */
	Meta sliceund.Und[map[string]string] `json:"meta,omitempty"`
	/**
	 * Allows you to copy the values of multiple fields into a group
	 * field, which can then be queried as a single field.
	 */
	CopyTo elastic.Elastic[string] `json:"copy_to,omitempty"`
}

type IndexOptions string

const (
	Docs      IndexOptions = "docs"
	Freqs     IndexOptions = "freqs"
	Positions IndexOptions = "positions"
	Offsets   IndexOptions = "offsets"
)

type TextIndexPrefixes struct {
	MaxChars int `json:"max_chars"`
	MinChars int `json:"min_chars"`
}

type TextProperty struct {
	CorePropertyBase

	Analyzer                 sliceund.Und[string]                           `json:"analyzer,omitempty"`
	Boost                    sliceund.Und[float64]                          `json:"boost,omitempty"`
	EagerGlobalOrdinals      sliceund.Und[bool]                             `json:"eager_global_ordinals,omitempty"`
	Fielddata                sliceund.Und[bool]                             `json:"fielddata,omitempty"`
	FielddataFrequencyFilter sliceund.Und[indices.FielddataFrequencyFilter] `json:"fielddata_frequency_filter,omitempty"`
	Index                    sliceund.Und[bool]                             `json:"index,omitempty"`
	IndexOptions             sliceund.Und[IndexOptions]                     `json:"index_options,omitempty"`
	IndexPhrases             sliceund.Und[bool]                             `json:"index_phrases,omitempty"`
	IndexPrefixes            sliceund.Und[TextIndexPrefixes]                `json:"index_prefixes,omitempty"`
	Norms                    sliceund.Und[bool]                             `json:"norms,omitempty"`
	PositionIncrementGap     sliceund.Und[int]                              `json:"position_increment_gap,omitempty"`
	SearchAnalyzer           sliceund.Und[string]                           `json:"search_analyzer,omitempty"`
	SearchQuoteAnalyzer      sliceund.Und[string]                           `json:"search_quote_analyzer,omitempty"`
	TermVector               sliceund.Und[TermVectorOption]                 `json:"term_vector,omitempty"`
	Type                     EsType                                         `json:"type"`
}

type VersionProperty struct {
	DocValuesPropertyBase

	Type EsType `json:"type"`
}

type WildcardProperty struct {
	DocValuesPropertyBase

	Type EsType `json:"type"`
	/**
	 * @availability stack since=7.15.0
	 * @availability serverless
	 */
	NullValue sliceund.Und[string] `json:"null_value,omitempty"`
}

type DynamicProperty struct {
	DocValuesPropertyBase

	Type EsType `json:"type"`

	Enabled   sliceund.Und[bool]    `json:"enabled,omitempty"`
	NullValue sliceund.Und[any]     `json:"null_value,omitempty"` // long | double | string | boolean | null | UserDefinedValue
	Boost     sliceund.Und[float64] `json:"boost,omitempty"`

	// NumberPropertyBase & long, double
	Coerce           sliceund.Und[bool]                                      `json:"coerce,omitempty"`
	Script           sliceund.Und[spec.Script]                               `json:"script,omitempty"`
	OnScriptError    sliceund.Und[OnScriptError]                             `json:"on_script_error,omitempty"`
	IgnoreMalformed  sliceund.Und[bool]                                      `json:"ignore_malformed,omitempty"`
	TimeSeriesMetric sliceund.Und[timeseriesmetrictype.TimeSeriesMetricType] `json:"time_series_metric,omitempty"`

	// string
	Analyzer             sliceund.Und[string]            `json:"analyzer,omitempty"`
	EagerGlobalOrdinals  sliceund.Und[bool]              `json:"eager_global_ordinals,omitempty"`
	Index                sliceund.Und[bool]              `json:"index,omitempty"`
	IndexOptions         sliceund.Und[IndexOptions]      `json:"index_options,omitempty"`
	IndexPhrases         sliceund.Und[bool]              `json:"index_phrases,omitempty"`
	IndexPrefixes        sliceund.Und[TextIndexPrefixes] `json:"index_prefixes,omitempty"`
	Norms                sliceund.Und[bool]              `json:"norms,omitempty"`
	PositionIncrementGap sliceund.Und[int]               `json:"position_increment_gap,omitempty"`
	SearchAnalyzer       sliceund.Und[string]            `json:"search_analyzer,omitempty"`
	SearchQuoteAnalyzer  sliceund.Und[string]            `json:"search_quote_analyzer,omitempty"`
	TermVector           sliceund.Und[TermVectorOption]  `json:"term_vector,omitempty"`

	// date
	Format        sliceund.Und[string] `json:"format,omitempty"`
	PrecisionStep sliceund.Und[int]    `json:"precision_step,omitempty"`
	Locale        sliceund.Und[string] `json:"locale,omitempty"`
}
