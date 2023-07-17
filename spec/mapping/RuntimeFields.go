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
	"github.com/ngicks/und/undefinedable"
)

type RuntimeFields map[string]RuntimeField

type RuntimeField struct {
	/** For type `lookup` */
	FetchFields undefinedable.Undefinedable[[]RuntimeFieldFetchFields] `json:"fetch_fields"`
	Format      undefinedable.Undefinedable[string]                    `json:"format"`
	/** For type `lookup` */
	InputField undefinedable.Undefinedable[string] `json:"input_field"`
	/** For type `lookup` */
	TargetField undefinedable.Undefinedable[string] `json:"target_field"`
	/** For type `lookup` */
	TargetIndex undefinedable.Undefinedable[string]        `json:"target_index"`
	Script      undefinedable.Undefinedable[esspec.Script] `json:"script"`
	Type        runtimefieldtype.RuntimeFieldType          `json:"type"`
}

/** @shortcut_property field */
type RuntimeFieldFetchFields struct {
	Field  string                              `json:"field"`
	Format undefinedable.Undefinedable[string] `json:"format"`
}
