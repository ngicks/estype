package mapping

import (
	"fmt"
)

func IsObjectLike(prop Property) bool {
	switch GetTypeName(prop) {
	case Object, Nested:
		return true
	default:
		return false
	}
}

func IsObject(prop Property) bool {
	switch GetTypeName(prop) {
	case Object:
		return true
	default:
		return false
	}
}

func IsNested(prop Property) bool {
	switch GetTypeName(prop) {
	case Nested:
		return true
	default:
		return false
	}
}

func GetTypeName(prop Property) EsType {
	switch prop.Val.(type) {
	case ObjectProperty:
		return Object
	case AggregateMetricDoubleProperty:
		return AggregateMetricDouble
	case FieldAliasProperty:
		return FieldAlias
	case BinaryProperty:
		return Binary
	case BooleanProperty:
		return Boolean
	case CompletionProperty:
		return Completion
	case DateProperty:
		return Date
	case DateNanosProperty:
		return DateNanos
	case DenseVectorProperty:
		return DenseVector
	case FlattenedProperty:
		return Flattened
	case GeoPointProperty:
		return GeoPoint
	case GeoShapeProperty:
		return GeoShape
	case HistogramProperty:
		return Histogram
	case IpProperty:
		return Ip
	case JoinProperty:
		return Join
	case NestedProperty:
		return Nested
	case PercolatorProperty:
		return Percolator
	case PointProperty:
		return Point
	case RankFeatureProperty:
		return RankFeature
	case RankFeaturesProperty:
		return RankFeatures
	case SearchAsYouTypeProperty:
		return SearchAsYouType
	case ShapeProperty:
		return Shape
	case TokenCountProperty:
		return TokenCount
	case VersionProperty:
		return Version
	case KeywordProperty:
		return Keyword
	case ConstantKeywordProperty:
		return ConstantKeyword
	case WildcardProperty:
		return Wildcard
	case TextProperty:
		return Text
	case LongNumberProperty:
		return LongNumber
	case IntegerNumberProperty:
		return IntegerNumber
	case ShortNumberProperty:
		return ShortNumber
	case ByteNumberProperty:
		return ByteNumber
	case DoubleNumberProperty:
		return DoubleNumber
	case FloatNumberProperty:
		return FloatNumber
	case HalfFloatNumberProperty:
		return HalfFloatNumber
	case UnsignedLongNumberProperty:
		return UnsignedLongNumber
	case ScaledFloatNumberProperty:
		return ScaledFloatNumber
	case IntegerRangeProperty:
		return IntegerRange
	case FloatRangeProperty:
		return FloatRange
	case LongRangeProperty:
		return LongRange
	case DoubleRangeProperty:
		return DoubleRange
	case DateRangeProperty:
		return DateRange
	case IpRangeProperty:
		return IpRange
	}
	panic(fmt.Errorf("unknown property type = %T", prop))
}
