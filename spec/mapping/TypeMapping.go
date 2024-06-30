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

import "github.com/ngicks/und/sliceund"

type TypeMapping struct {
	AllField           sliceund.Und[AllField]                     `json:"all_field,omitempty"`
	DateDetection      sliceund.Und[bool]                         `json:"date_detection,omitempty"`
	Dynamic            sliceund.Und[DynamicMapping]               `json:"dynamic,omitempty"`
	DynamicDateFormats sliceund.Und[[]string]                     `json:"dynamic_date_formats,omitempty"`
	DynamicTemplates   sliceund.Und[[]map[string]DynamicTemplate] `json:"dynamic_templates,omitempty"`
	FieldNames_        sliceund.Und[FieldNamesField]              `json:"_field_names,omitempty"`
	IndexField         sliceund.Und[IndexField]                   `json:"index_field,omitempty"`
	/** @doc_id mapping-meta-field */
	Meta             sliceund.Und[map[string]any]          `json:"_meta,omitempty"`
	NumericDetection sliceund.Und[bool]                    `json:"numeric_detection,omitempty"`
	Properties       sliceund.Und[map[string]Property]     `json:"properties,omitempty"`
	Routing          sliceund.Und[RoutingField]            `json:"_routing,omitempty"`
	Size             sliceund.Und[SizeField]               `json:"_size,omitempty"`
	Source           sliceund.Und[SourceField]             `json:"_source,omitempty"`
	Runtime          sliceund.Und[map[string]RuntimeField] `json:"runtime,omitempty"`
	Enabled          sliceund.Und[bool]                    `json:"enabled,omitempty"`
	/** @since 7.16.0 */
	DataStreamTimestamp sliceund.Und[DataStreamTimestamp] `json:"_data_stream_timestamp,omitempty"`
}

type DataStreamTimestamp struct {
	Enabled bool `json:"enabled"`
}
