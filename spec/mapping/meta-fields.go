// Hand port of https://github.com/elastic/elasticsearch-specification/blob/76e25d34bff1060e300c95f4be468ef88e4f3465/specification/_types/mapping/meta-fields.ts
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

import "github.com/ngicks/und/undefinedable"

type FieldMapping struct {
	FullName string              `json:"full_name"`
	Mapping  map[string]Property `json:"mapping"`
}

type AllField struct {
	Analyzer                 string `json:"analyzer"`
	Enabled                  bool   `json:"enabled"`
	OmitNorms                bool   `json:"omit_norms"`
	SearchAnalyzer           string `json:"search_analyzer"`
	Similarity               string `json:"similarity"`
	Store                    bool   `json:"store"`
	StoreTermVectorOffsets   bool   `json:"store_term_vector_offsets"`
	StoreTermVectorPayloads  bool   `json:"store_term_vector_payloads"`
	StoreTermVectorPositions bool   `json:"store_term_vector_positions"`
	StoreTermVectors         bool   `json:"store_term_vectors"`
}

type FieldNamesField struct {
	Enabled bool `json:"enabled"`
}

type IndexField struct {
	Enabled bool `json:"enabled"`
}

type RoutingField struct {
	Required bool `json:"required"`
}

type SizeField struct {
	Enabled bool `json:"enabled"`
}

type SourceField struct {
	Compress          undefinedable.Undefinedable[bool]            `json:"compress"`
	CompressThreshold undefinedable.Undefinedable[string]          `json:"compress_threshold"`
	Enabled           undefinedable.Undefinedable[bool]            `json:"enabled"`
	Excludes          undefinedable.Undefinedable[[]string]        `json:"excludes"`
	Includes          undefinedable.Undefinedable[[]string]        `json:"includes"`
	Mode              undefinedable.Undefinedable[SourceFieldMode] `json:"mode"`
}

type SourceFieldMode string

const (
	Disabled SourceFieldMode = "disabled"
	Stored   SourceFieldMode = "stored"
	/**
	 *  Instead of storing source documents on disk exactly as you send them,
	 *  Elasticsearch can reconstruct source content on the fly upon retrieval.
	 */
	Synthetic SourceFieldMode = "synthetic"
)
