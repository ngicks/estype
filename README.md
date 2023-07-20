# estype

- A code generator that generates Go structs from
  [Elasticsearch](https://www.elastic.co/guide/en/elastic-stack/index.html)
  mapping.json.
- Types for the Elasticsearch
  [field data types](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/mapping-types.html),
  such like a date string to a time.Time, or a geo point type that can be
  unmarshaled from every 6 variants of representations.

## Rationale

The Elasticsearch is complex enough to confuse people when they are generating /
consuming JSON documents it stores, needing a code generator and helper types.

For an overview of the Elasticsearch, see
[here](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/elasticsearch-intro.html)
and
[here](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/documents-indices.html).

The Elasticsearch is a popular distributed data store which analyzes the JSON
documents clients have stored, and provides great search functionality over
them.

While it can operate on schema-less way, it also allows users to set mapping on
indices, through which you can define and optionally fixate data format of JSON
documents partially or fully.

Surprisingly, all fields of input JSON objects are allowed to be any of
[`undefined | null | T | (null | T)[]`](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/array.html).
More surprisingly it also accepts nested array of T (`T[][]`) like
`[1, 2, 3, [4, 5]]`, which will behave as if it were flattened (i.e.
`[1, 2, 3, 4, 5]`) in the search context.

Moreover, some of field data types need special handling.

- [Boolean](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/boolean.html)
  accepts `false`, `"false"`, `""` as false values. `true`, `"true"` as true
  values.
- [Date](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/date.html)
  and
  [Date Nanoseconds](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/date_nanos.html)
  vary their stringified data format on basis of corresponding `"format"` field
  in mapping.
- [Geo Point](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/geo-point.html#geo-point)
  and
  [Geo Shape](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/geo-shape.html)
  has multiple representations. Both support [GeoJSON](https://geojson.org/) and
  [Well-Known Text](https://docs.ogc.org/is/12-063r5/12-063r5.html). For
  historical reasons, GeoPoint supports additionally 4 more formats, 6 different
  formats in total.
- [Object](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/object.html)
  and
  [Nested](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/nested.html)
  define sub objects of input JSON documents. Objects are hierarchical in
  nature.

Generating and/or consuming documents sometimes are confusing:

- You can store `{"kwd":"foo"}` for an index set up with mapping
  `{"properties":{"kwd":{"type":"keyword"}}}`.
- At same time it also could be `{"kwd":["foo","bar"]}`.
- As your service grows, you might find that storing multiple values to `"kwd"`
  is more suitable for your use case.
- Or otherwise, it could happen when more than 2 different apps are writing to
  the same index.
- At least those 2 variants are legal for the mapping.

## Problem: This does not fit Go's conventional approach.

Conventional Go way is to use statically defined structures.

Go has a native unmarshaling method to parse serialized JSON documents to
structs. However unmarshaling multiple types of data source into a single Go
type is not done automatically, needing efforts of user code.

## Helper types

The module fielddatatype hosts helper types for those types marshal / unmarshal
into multiple formats.

- AggregateMetricDouble* kind types are all variants of the
  aggregated_metric_double field data type.
  - The code generator selects an appropriate one for a mapping.
- Boolean, BooleanStr is a bool based type that is unmarshaled from `true` /
  `false` / `"true"` / `"false"` / `""` and marshal into boolean literal or
  string of `"true"` | `"false"` respectively.
- GeoPoint can be unmarshaled from all 6 possible formats and marshals into 1
  selected format, `{"lat":123,"lon":456}`
- GeoShape accepts 2 possible formats as source. It marshal into GeoJSON format
  delegating behavior to `github.com/go-spatial/geom`.
- The estime package is a set of functions that makes Go std time package
  understand
  [the DateTimeFormatter formats](https://docs.oracle.com/javase/8/docs/api/java/time/format/DateTimeFormatter.html).
  - It flattens optional parts of the format into any possible patterns then
    converts Java time tokens into Go equivalents.
  - It drops support for number of tokens, for which Go has no counterparts,
    including `G(era)`, `Q/q(quarter-of-year)`, `w(week-of-week-based-year)` and
    `W(week-of-month)`.
  - The code generator does this transformation.

## The Code generator

The code generator finally generates Go struct to generate / consume JSON
documents which is stored in a (cluster of) Elasticsearch instance(s).

- It generates plain and raw types.
- Raw types accept all possible field data.
  - Namely, `undefined | null | T | (null| T)[]`
- Plain types are the application defined type.
  - When invoking the code generator, the user can pass configuration file.
  - The configuration file defines which field are to store single value (`T`),
    multiple values(`[]T`), required(`T`or `[]T`) or optional(`*T` or `*[]T`).
- It generates date format on basis of its mapping definition.
  - If format for the field is empty or a single built-in format, it uses
    pre-generated types to reduce generated code.

Future updates may optimize behavior.
