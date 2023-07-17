// Hand port of https://github.com/elastic/elasticsearch-specification/blob/76e25d34bff1060e300c95f4be468ef88e4f3465/specification/_types/mapping/geo.ts
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
	"github.com/ngicks/und/undefinedable"
)

type GeoPointProperty struct {
	DocValuesPropertyBase
	IgnoreMalformed undefinedable.Undefinedable[bool] `json:" ignore_malformed"`
	IgnoreZValue    undefinedable.Undefinedable[bool] `json:" ignore_z_value"`
	NullValue       undefinedable.Undefinedable[any]  `json:" null_value"`
	Type            EsType                            `json:"type"`
}

type GeoOrientation string

const (
	/** @aliases RIGHT, counterclockwise, ccw */
	Right GeoOrientation = "right"
	/** @aliases LEFT, clockwise, cw */
	Left GeoOrientation = "left"
)

/**
 * The `geo_shape` data type facilitates the indexing of and searching with arbitrary geo shapes such as rectangles
 * and polygons.
 *
 * @doc_id geo-shape
 */
type GeoShapeProperty struct {
	DocValuesPropertyBase
	Coerce          undefinedable.Undefinedable[bool]           `json:" coerce"`
	IgnoreMalformed undefinedable.Undefinedable[bool]           `json:" ignore_malformed"`
	IgnoreZValue    undefinedable.Undefinedable[bool]           `json:" ignore_z_value"`
	Orientation     undefinedable.Undefinedable[GeoOrientation] `json:" orientation"`
	Strategy        undefinedable.Undefinedable[GeoStrategy]    `json:" strategy"`
	Type            EsType                                      `json:"type"`
}

type GeoStrategy string

const (
	Recursive GeoStrategy = "recursive"
	Term      GeoStrategy = "term"
)

type GeoTree string

const (
	Geohash  GeoTree = "geohash"
	Quadtree GeoTree = "quadtree"
)

type PointProperty struct {
	DocValuesPropertyBase
	IgnoreMalformed undefinedable.Undefinedable[bool]   `json:" ignore_malformed"`
	IgnoreZValue    undefinedable.Undefinedable[bool]   `json:" ignore_z_value"`
	NullValue       undefinedable.Undefinedable[string] `json:" null_value"`
	Type            EsType                              `json:"type"`
}

/**
 * The `shape` data type facilitates the indexing of and searching with arbitrary `x, y` cartesian shapes such as
 * rectangles and polygons.
 *
 * @doc_id shape
 */
type ShapeProperty struct {
	DocValuesPropertyBase
	Coerce          undefinedable.Undefinedable[bool]           `json:" coerce"`
	IgnoreMalformed undefinedable.Undefinedable[bool]           `json:" ignore_malformed"`
	IgnoreZValue    undefinedable.Undefinedable[bool]           `json:" ignore_z_value"`
	Orientation     undefinedable.Undefinedable[GeoOrientation] `json:" orientation"`
	Type            EsType                                      `json:"type"`
}
