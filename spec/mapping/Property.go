// Hand port of https://github.com/elastic/elasticsearch-specification/blob/76e25d34bff1060e300c95f4be468ef88e4f3465/specification/_types/mapping/Property.ts
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
	"encoding/json"
	"fmt"

	"github.com/ngicks/und/sliceund"
)

// Property is enum for mapping field type.
// type is not necessary to be set. If absent property is considered object.
//
// | BinaryProperty
// | BooleanProperty
// | DynamicProperty
// | JoinProperty
// | KeywordProperty
// | MatchOnlyTextProperty
// | PercolatorProperty
// | RankFeatureProperty
// | RankFeaturesProperty
// | SearchAsYouTypeProperty
// | TextProperty
// | VersionProperty
// | WildcardProperty
// | DateNanosProperty
// | DateProperty
// | AggregateMetricDoubleProperty
// | DenseVectorProperty
// | FlattenedProperty
// | NestedProperty
// | ObjectProperty
// | CompletionProperty
// | ConstantKeywordProperty
// | FieldAliasProperty
// | HistogramProperty
// | IpProperty
// | Murmur3HashProperty
// | TokenCountProperty
// | GeoPointProperty
// | GeoShapeProperty
// | PointProperty
// | ShapeProperty
// | ByteNumberProperty
// | DoubleNumberProperty
// | FloatNumberProperty
// | HalfFloatNumberProperty
// | IntegerNumberProperty
// | LongNumberProperty
// | ScaledFloatNumberProperty
// | ShortNumberProperty
// | UnsignedLongNumberProperty
// | DateRangeProperty
// | DoubleRangeProperty
// | FloatRangeProperty
// | IntegerRangeProperty
// | IpRangeProperty
// | LongRangeProperty
type Property struct {
	Val any
}

func (p Property) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Val)
}

func (p *Property) UnmarshalJSON(data []byte) error {
	type prop struct {
		Type sliceund.Und[EsType] `json:"type,omitempty"`
	}
	var inner prop

	err := json.Unmarshal(data, &inner)
	if err != nil {
		return err
	}

	if inner.Type.IsUndefined() {
		var oo ObjectProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
		return nil
	}

	switch inner.Type.Value() {
	default:
		return fmt.Errorf("unknown type = %s", inner.Type.Value())
	case Binary:
		var oo BinaryProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case Boolean:
		var oo BooleanProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case Dynamic:
		var oo DynamicProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case Join:
		var oo JoinProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case Keyword:
		var oo KeywordProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case MatchOnlyText:
		var oo MatchOnlyTextProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case Percolator:
		var oo PercolatorProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case RankFeature:
		var oo RankFeatureProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case RankFeatures:
		var oo RankFeaturesProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case SearchAsYouType:
		var oo SearchAsYouTypeProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case Text:
		var oo TextProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case Version:
		var oo VersionProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case Wildcard:
		var oo WildcardProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case DateNanos:
		var oo DateNanosProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case Date:
		var oo DateProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case AggregateMetricDouble:
		var oo AggregateMetricDoubleProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case DenseVector:
		var oo DenseVectorProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case Flattened:
		var oo FlattenedProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case Nested:
		var oo NestedProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case Object:
		var oo ObjectProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case Completion:
		var oo CompletionProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case ConstantKeyword:
		var oo ConstantKeywordProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case FieldAlias:
		var oo FieldAliasProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case Histogram:
		var oo HistogramProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case Ip:
		var oo IpProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case Murmur3Hash:
		var oo Murmur3HashProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case TokenCount:
		var oo TokenCountProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case GeoPoint:
		var oo GeoPointProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case GeoShape:
		var oo GeoShapeProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case Point:
		var oo PointProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case Shape:
		var oo ShapeProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case ByteNumber:
		var oo ByteNumberProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case DoubleNumber:
		var oo DoubleNumberProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case FloatNumber:
		var oo FloatNumberProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case HalfFloatNumber:
		var oo HalfFloatNumberProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case IntegerNumber:
		var oo IntegerNumberProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case LongNumber:
		var oo LongNumberProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case ScaledFloatNumber:
		var oo ScaledFloatNumberProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case ShortNumber:
		var oo ShortNumberProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case UnsignedLongNumber:
		var oo UnsignedLongNumberProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case DateRange:
		var oo DateRangeProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case DoubleRange:
		var oo DoubleRangeProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case FloatRange:
		var oo FloatRangeProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case IntegerRange:
		var oo IntegerRangeProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case IpRange:
		var oo IpRangeProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	case LongRange:
		var oo LongRangeProperty
		err := json.Unmarshal(data, &oo)
		if err != nil {
			return err
		}
		p.Val = oo
	}
	return nil
}

type PropertyBase struct {
	Meta        sliceund.Und[map[string]string]   `json:"meta,omitempty,omitempty"`
	Properties  sliceund.Und[map[string]Property] `json:"properties,omitempty,omitempty"`
	IgnoreAbove sliceund.Und[int]                 `json:"ignore_above,omitempty,omitempty"`
	Dynamic     sliceund.Und[DynamicMapping]      `json:"dynamic,omitempty,omitempty"`
	Fields      sliceund.Und[map[string]Property] `json:"fields,omitempty,omitempty"`
}
