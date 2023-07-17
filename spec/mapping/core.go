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
	"github.com/ngicks/und/elastic"
	"github.com/ngicks/und/undefinedable"
)

type CorePropertyBase struct {
	PropertyBase
	CopyTo     elastic.Elastic[string]             `json:"copy_to"`
	Similarity undefinedable.Undefinedable[string] `json:"similarity"`
	Store      undefinedable.Undefinedable[bool]   `json:"store"`
}

type DocValuesPropertyBase struct {
	CorePropertyBase
	DocValues undefinedable.Undefinedable[bool] `json:"doc_values"`
}

type BinaryProperty struct {
	DocValuesPropertyBase
	Type EsType `json:"type"`
}

type BooleanProperty struct {
	DocValuesPropertyBase
	Boost     undefinedable.Undefinedable[float64]                  `json:"boost"`
	Fielddata undefinedable.Undefinedable[indices.NumericFielddata] `json:"fielddata"`
	Index     undefinedable.Undefinedable[bool]                     `json:"index"`
	NullValue undefinedable.Undefinedable[bool]                     `json:"null_value"`
	Type      EsType                                                `json:"type"`
}

type DateProperty struct {
	DocValuesPropertyBase

	Boost undefinedable.Undefinedable[float64] `json:"boost"`

	Fielddata undefinedable.Undefinedable[indices.NumericFielddata] `json:"fielddata"`

	Format          undefinedable.Undefinedable[string] `json:"format"`
	IgnoreMalformed undefinedable.Undefinedable[bool]   `json:"ignore_malformed"`
	Index           undefinedable.Undefinedable[bool]   `json:"index"`
	NullValue       undefinedable.Undefinedable[any]    `json:"null_value"` // date string formatted as ones specified in format field or epoch millis int
	PrecisionStep   undefinedable.Undefinedable[int]    `json:"precision_step"`
	Locale          undefinedable.Undefinedable[string] `json:"locale"`
	Type            EsType                              `json:"type"`
}

type DateNanosProperty struct {
	DocValuesPropertyBase

	Boost           undefinedable.Undefinedable[float64] `json:"boost"`
	Format          undefinedable.Undefinedable[string]  `json:"format"`
	IgnoreMalformed undefinedable.Undefinedable[bool]    `json:"ignore_malformed"`
	Index           undefinedable.Undefinedable[bool]    `json:"index"`
	NullValue       undefinedable.Undefinedable[any]     `json:"null_value"` // date string formatted as ones specified in format field or epoch millis int
	PrecisionStep   undefinedable.Undefinedable[int]     `json:"precision_step"`
	Type            EsType                               `json:"type"`
}

type JoinProperty struct {
	PropertyBase

	Relations           undefinedable.Undefinedable[map[string]elastic.Elastic[string]] `json:"relations"`
	EagerGlobalOrdinals undefinedable.Undefinedable[bool]                               `json:"eager_global_ordinals"`
	Type                EsType                                                          `json:"type"`
}

type KeywordProperty struct {
	DocValuesPropertyBase

	Boost                    undefinedable.Undefinedable[float64]      `json:"boost"`
	EagerGlobalOrdinals      undefinedable.Undefinedable[bool]         `json:"eager_global_ordinals"`
	Index                    undefinedable.Undefinedable[bool]         `json:"index"`
	IndexOptions             undefinedable.Undefinedable[IndexOptions] `json:"index_options"`
	Normalizer               undefinedable.Undefinedable[string]       `json:"normalizer"`
	Norms                    undefinedable.Undefinedable[bool]         `json:"norms"`
	NullValue                undefinedable.Undefinedable[string]       `json:"null_value"`
	SplitQueriesOnWhitespace undefinedable.Undefinedable[bool]         `json:"split_queries_on_whitespace"`
	/**
	 * For internal use by Elastic only. Marks the field as a time series dimension. Defaults to false.
	 * @availability stack stability=experimental
	 * @availability serverless stability=experimental
	 */
	TimeSeriesDimension undefinedable.Undefinedable[bool] `json:"time_series_dimension"`
	Type                EsType                            `json:"type"`
}

type NumberPropertyBase struct {
	DocValuesPropertyBase

	Boost           undefinedable.Undefinedable[float64]       `json:"boost"`
	Coerce          undefinedable.Undefinedable[bool]          `json:"coerce"`
	IgnoreMalformed undefinedable.Undefinedable[bool]          `json:"ignore_malformed"`
	Index           undefinedable.Undefinedable[bool]          `json:"index"`
	OnScriptError   undefinedable.Undefinedable[OnScriptError] `json:"on_script_error"`
	Script          undefinedable.Undefinedable[spec.Script]   `json:"script"`
	/**
	 * For internal use by Elastic only. Marks the field as a time series dimension. Defaults to false.
	 * @availability stack stability=experimental
	 * @availability serverless stability=experimental
	 */
	TimeSeriesMetric undefinedable.Undefinedable[timeseriesmetrictype.TimeSeriesMetricType] `json:"time_series_metric"`
	/**
	 * For internal use by Elastic only. Marks the field as a time series dimension. Defaults to false.
	 * @server_default false
	 * @availability stack stability=experimental
	 * @availability serverless stability=experimental
	 */
	TimeSeriesDimension undefinedable.Undefinedable[bool] `json:"time_series_dimension"`
}

type OnScriptError string

const (
	Fail     OnScriptError = "fail"
	Continue OnScriptError = "continue"
)

type FloatNumberProperty struct {
	NumberPropertyBase

	Type      EsType                               `json:"type"`
	NullValue undefinedable.Undefinedable[float32] `json:"null_value"`
}

type HalfFloatNumberProperty struct {
	NumberPropertyBase

	Type      EsType                               `json:"type"`
	NullValue undefinedable.Undefinedable[float32] `json:"null_value"` // TODO: use float16 package?
}

type DoubleNumberProperty struct {
	NumberPropertyBase

	Type      EsType                               `json:"type"`
	NullValue undefinedable.Undefinedable[float64] `json:"null_value"`
}

type IntegerNumberProperty struct {
	NumberPropertyBase

	Type      EsType                             `json:"type"`
	NullValue undefinedable.Undefinedable[int32] `json:"null_value"`
}

type LongNumberProperty struct {
	NumberPropertyBase

	Type      EsType                             `json:"type"`
	NullValue undefinedable.Undefinedable[int64] `json:"null_value"`
}

type ShortNumberProperty struct {
	NumberPropertyBase

	Type      EsType                             `json:"type"`
	NullValue undefinedable.Undefinedable[int16] `json:"null_value"`
}

type ByteNumberProperty struct {
	NumberPropertyBase

	Type      EsType                            `json:"type"`
	NullValue undefinedable.Undefinedable[int8] `json:"null_value"`
}

type UnsignedLongNumberProperty struct {
	NumberPropertyBase

	Type      EsType                              `json:"type"`
	NullValue undefinedable.Undefinedable[uint64] `json:"null_value"`
}

type ScaledFloatNumberProperty struct {
	NumberPropertyBase

	Type          EsType                               `json:"type"`
	NullValue     undefinedable.Undefinedable[float64] `json:"null_value"`
	ScalingFactor undefinedable.Undefinedable[float64] `json:"scaling_factor"`
}

type PercolatorProperty struct {
	PropertyBase

	Type EsType `json:"type"`
}

type RankFeatureProperty struct {
	PropertyBase

	PositiveScoreImpact undefinedable.Undefinedable[bool] `json:"positive_score_impact"`
	Type                EsType                            `json:"type"`
}

type RankFeaturesProperty struct {
	PropertyBase

	Type EsType `json:"type"`
}

type SearchAsYouTypeProperty struct {
	CorePropertyBase

	Analyzer            undefinedable.Undefinedable[string]           `json:"analyzer"`
	Index               undefinedable.Undefinedable[bool]             `json:"index"`
	IndexOptions        undefinedable.Undefinedable[IndexOptions]     `json:"index_options"`
	MaxShingleSize      undefinedable.Undefinedable[int]              `json:"max_shingle_size"`
	Norms               undefinedable.Undefinedable[bool]             `json:"norms"`
	SearchAnalyzer      undefinedable.Undefinedable[string]           `json:"search_analyzer"`
	SearchQuoteAnalyzer undefinedable.Undefinedable[string]           `json:"search_quote_analyzer"`
	TermVector          undefinedable.Undefinedable[TermVectorOption] `json:"term_vector"`
	Type                EsType                                        `json:"type"`
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
	Fields undefinedable.Undefinedable[map[string]Property] `json:"fields"`
	/**
	 * Metadata about the field.
	 * @doc_id mapping-meta-field
	 */
	Meta undefinedable.Undefinedable[map[string]string] `json:"meta"`
	/**
	 * Allows you to copy the values of multiple fields into a group
	 * field, which can then be queried as a single field.
	 */
	CopyTo elastic.Elastic[string] `json:"copy_to"`
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

	Analyzer                 undefinedable.Undefinedable[string]                           `json:"analyzer"`
	Boost                    undefinedable.Undefinedable[float64]                          `json:"boost"`
	EagerGlobalOrdinals      undefinedable.Undefinedable[bool]                             `json:"eager_global_ordinals"`
	Fielddata                undefinedable.Undefinedable[bool]                             `json:"fielddata"`
	FielddataFrequencyFilter undefinedable.Undefinedable[indices.FielddataFrequencyFilter] `json:"fielddata_frequency_filter"`
	Index                    undefinedable.Undefinedable[bool]                             `json:"index"`
	IndexOptions             undefinedable.Undefinedable[IndexOptions]                     `json:"index_options"`
	IndexPhrases             undefinedable.Undefinedable[bool]                             `json:"index_phrases"`
	IndexPrefixes            undefinedable.Undefinedable[TextIndexPrefixes]                `json:"index_prefixes"`
	Norms                    undefinedable.Undefinedable[bool]                             `json:"norms"`
	PositionIncrementGap     undefinedable.Undefinedable[int]                              `json:"position_increment_gap"`
	SearchAnalyzer           undefinedable.Undefinedable[string]                           `json:"search_analyzer"`
	SearchQuoteAnalyzer      undefinedable.Undefinedable[string]                           `json:"search_quote_analyzer"`
	TermVector               undefinedable.Undefinedable[TermVectorOption]                 `json:"term_vector"`
	Type                     EsType                                                        `json:"type"`
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
	NullValue undefinedable.Undefinedable[string] `json:"null_value"`
}

type DynamicProperty struct {
	DocValuesPropertyBase

	Type EsType `json:"type"`

	Enabled   undefinedable.Undefinedable[bool]    `json:"enabled"`
	NullValue undefinedable.Undefinedable[any]     `json:"null_value"` // long | double | string | boolean | null | UserDefinedValue
	Boost     undefinedable.Undefinedable[float64] `json:"boost"`

	// NumberPropertyBase & long, double
	Coerce           undefinedable.Undefinedable[bool]                                      `json:"coerce"`
	Script           undefinedable.Undefinedable[spec.Script]                               `json:"script"`
	OnScriptError    undefinedable.Undefinedable[OnScriptError]                             `json:"on_script_error"`
	IgnoreMalformed  undefinedable.Undefinedable[bool]                                      `json:"ignore_malformed"`
	TimeSeriesMetric undefinedable.Undefinedable[timeseriesmetrictype.TimeSeriesMetricType] `json:"time_series_metric"`

	// string
	Analyzer             undefinedable.Undefinedable[string]            `json:"analyzer"`
	EagerGlobalOrdinals  undefinedable.Undefinedable[bool]              `json:"eager_global_ordinals"`
	Index                undefinedable.Undefinedable[bool]              `json:"index"`
	IndexOptions         undefinedable.Undefinedable[IndexOptions]      `json:"index_options"`
	IndexPhrases         undefinedable.Undefinedable[bool]              `json:"index_phrases"`
	IndexPrefixes        undefinedable.Undefinedable[TextIndexPrefixes] `json:"index_prefixes"`
	Norms                undefinedable.Undefinedable[bool]              `json:"norms"`
	PositionIncrementGap undefinedable.Undefinedable[int]               `json:"position_increment_gap"`
	SearchAnalyzer       undefinedable.Undefinedable[string]            `json:"search_analyzer"`
	SearchQuoteAnalyzer  undefinedable.Undefinedable[string]            `json:"search_quote_analyzer"`
	TermVector           undefinedable.Undefinedable[TermVectorOption]  `json:"term_vector"`

	// date
	Format        undefinedable.Undefinedable[string] `json:"format"`
	PrecisionStep undefinedable.Undefinedable[int]    `json:"precision_step"`
	Locale        undefinedable.Undefinedable[string] `json:"locale"`
}
