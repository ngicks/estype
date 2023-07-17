package mapping

// Enum for Elasticsearch mapping field data types.
type EsType string

const (
	// common
	Binary          EsType = "binary"
	Boolean         EsType = "boolean"
	Dynamic         EsType = "dynamic"
	Join            EsType = "join"
	Keyword         EsType = "keyword"
	MatchOnlyText   EsType = "match_only_text"
	Percolator      EsType = "percolator"
	RankFeature     EsType = "rank_feature"
	RankFeatures    EsType = "rank_features"
	SearchAsYouType EsType = "search_as_you_type"
	Text            EsType = "text"
	Version         EsType = "version"
	Wildcard        EsType = "wildcard"
	// dates
	DateNanos EsType = "date_nanos"
	Date      EsType = "date"
	// complex
	AggregateMetricDouble EsType = "aggregate_metric_double"
	DenseVector           EsType = "dense_vector"
	Flattened             EsType = "flattened"
	Nested                EsType = "nested"
	Object                EsType = "object"
	// structured
	Completion      EsType = "completion"
	ConstantKeyword EsType = "constant_keyword"
	FieldAlias      EsType = "alias"
	Histogram       EsType = "histogram"
	Ip              EsType = "ip"
	Murmur3Hash     EsType = "murmur3"
	TokenCount      EsType = "token_count"
	// spatial
	GeoPoint EsType = "geo_point"
	GeoShape EsType = "geo_shape"
	Point    EsType = "point"
	Shape    EsType = "shape"
	// numbers
	ByteNumber         EsType = "byte"
	DoubleNumber       EsType = "double"
	FloatNumber        EsType = "float"
	HalfFloatNumber    EsType = "half_float"
	IntegerNumber      EsType = "integer"
	LongNumber         EsType = "long"
	ScaledFloatNumber  EsType = "scaled_float"
	ShortNumber        EsType = "short"
	UnsignedLongNumber EsType = "unsigned_long"
	// ranges
	DateRange    EsType = "date_range"
	DoubleRange  EsType = "double_range"
	FloatRange   EsType = "float_range"
	IntegerRange EsType = "integer_range"
	IpRange      EsType = "ip_range"
	LongRange    EsType = "long_range"
)
