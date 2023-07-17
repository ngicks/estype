// Hand port of https://github.com/elastic/elasticsearch-specification/blob/76e25d34bff1060e300c95f4be468ef88e4f3465/specification/_types/mapping/TypeMapping.ts
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

type TypeMapping struct {
	AllField           undefinedable.Undefinedable[AllField]                     `json:"all_field"`
	DateDetection      undefinedable.Undefinedable[bool]                         `json:"date_detection"`
	Dynamic            undefinedable.Undefinedable[DynamicMapping]               `json:"dynamic"`
	DynamicDateFormats undefinedable.Undefinedable[[]string]                     `json:"dynamic_date_formats"`
	DynamicTemplates   undefinedable.Undefinedable[[]map[string]DynamicTemplate] `json:"dynamic_templates"`
	FieldNames_        undefinedable.Undefinedable[FieldNamesField]              `json:"_field_names"`
	IndexField         undefinedable.Undefinedable[IndexField]                   `json:"index_field"`
	/** @doc_id mapping-meta-field */
	Meta             undefinedable.Undefinedable[map[string]any]          `json:"_meta"`
	NumericDetection undefinedable.Undefinedable[bool]                    `json:"numeric_detection"`
	Properties       undefinedable.Undefinedable[map[string]Property]     `json:"properties"`
	Routing          undefinedable.Undefinedable[RoutingField]            `json:"_routing"`
	Size             undefinedable.Undefinedable[SizeField]               `json:"_size"`
	Source           undefinedable.Undefinedable[SourceField]             `json:"_source"`
	Runtime          undefinedable.Undefinedable[map[string]RuntimeField] `json:"runtime"`
	Enabled          undefinedable.Undefinedable[bool]                    `json:"enabled"`
	/** @since 7.16.0 */
	DataStreamTimestamp undefinedable.Undefinedable[DataStreamTimestamp] `json:"_data_stream_timestamp"`
}

type DataStreamTimestamp struct {
	Enabled bool `json:"enabled"`
}
