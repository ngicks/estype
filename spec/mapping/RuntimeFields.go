// Hand port of https://github.com/elastic/elasticsearch-specification/blob/76e25d34bff1060e300c95f4be468ef88e4f3465/specification/_types/mapping/RuntimeFields.ts
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
	esspec "github.com/ngicks/estype/spec"
	runtimefieldtype "github.com/ngicks/estype/spec/mapping/runtime_field_type"
	"github.com/ngicks/und/sliceund"
)

type RuntimeFields map[string]RuntimeField

type RuntimeField struct {
	/** For type `lookup` */
	FetchFields sliceund.Und[[]RuntimeFieldFetchFields] `json:"fetch_fields,omitempty"`
	Format      sliceund.Und[string]                    `json:"format,omitempty"`
	/** For type `lookup` */
	InputField sliceund.Und[string] `json:"input_field,omitempty"`
	/** For type `lookup` */
	TargetField sliceund.Und[string] `json:"target_field,omitempty"`
	/** For type `lookup` */
	TargetIndex sliceund.Und[string]              `json:"target_index,omitempty"`
	Script      sliceund.Und[esspec.Script]       `json:"script,omitempty"`
	Type        runtimefieldtype.RuntimeFieldType `json:"type"`
}

/** @shortcut_property field */
type RuntimeFieldFetchFields struct {
	Field  string               `json:"field"`
	Format sliceund.Und[string] `json:"format,omitempty"`
}
