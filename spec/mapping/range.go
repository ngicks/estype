// Hand port of https://github.com/elastic/elasticsearch-specification/blob/76e25d34bff1060e300c95f4be468ef88e4f3465/specification/_types/mapping/range.ts
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

type RangePropertyBase struct {
	DocValuesPropertyBase
	Boost  sliceund.Und[float64] `json:" boost,omitempty"`
	Coerce sliceund.Und[bool]    `json:" coerce,omitempty"`
	Index  sliceund.Und[bool]    `json:" index,omitempty"`
}

type DateRangeProperty struct {
	RangePropertyBase
	Format sliceund.Und[string] `json:" format,omitempty"`
	Type   EsType               `json:"type"`
}

type DoubleRangeProperty struct {
	RangePropertyBase
	Type EsType `json:"type"`
}

type FloatRangeProperty struct {
	RangePropertyBase
	Type EsType `json:"type"`
}

type IntegerRangeProperty struct {
	RangePropertyBase
	Type EsType `json:"type"`
}

type IpRangeProperty struct {
	RangePropertyBase
	Type EsType `json:"type"`
}

type LongRangeProperty struct {
	RangePropertyBase
	Type EsType `json:"type"`
}
