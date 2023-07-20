# estype

- A code generator that generates Go structs from
  [Elasticsearch](https://www.elastic.co/guide/en/elastic-stack/index.html)
  mapping.json.
- Types for the Elasticsearch
  [field data types](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/mapping-types.html),
  such like a date string to a time.Time, or a geo point type that can be
  unmarshaled from every 6 variants of representations.

## Usage

```
# ./genestype --help
Usage of ./genestype:
  -c string
        path to config file.
        see definition of github.com/ngicks/estype/generator.GeneratorOption.
  -m string
        path to mapping.json.
        You can use one that can be fetched from '<index_name>/_mapping',
        or one that you've sent when creating index.
  -o string
        [optional]
        path to output generated code.
        More than 2 distinct mappings should not be generated to the same directory
        because it possibly creates helper functions / types depending on the config and the mapping.
        defaults to stdout. (default "--")
  -p string
        package name of generated code.
```

see [./generator/test/testdata](./generator/test/testdata) for example mappings
and options. see example of generated in [./generator/test](./generator/test)

You can also use the `generator` module directly.

```go
var optionFile, mappingFile io.Reader
var outputFile io.Writer

var generateOpt generator.GeneratorOption
if err := json.NewDecoder(optionFile).Decode(&generateOpt); err != nil {
  panic(err)
}

bin, err := io.ReadAll(mappingFile)
if err != nil {
  panic(err)
}
generateOpt.Mapping, err = eshelper.GetMapping(bin)
if err != nil {
  panic(err)
}

generateOpt.GenerateTypeName = generator.ChainFieldName

f := jen.NewFilePath(*packagePath)
generateOpt.NewContext(f).Gen()

if err := f.Render(outputFile); err != nil {
  panic(err)
}
```

## Rationale

The Elasticsearch is complex enough to sometimes confuse people when they are
trying to create / consume JSON documents it stores, needing a code generator
and helper types.

For an overview of the Elasticsearch, see
[here](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/elasticsearch-intro.html)
and
[here](https://www.elastic.co/guide/en/elasticsearch/reference/8.4/documents-indices.html).

The Elasticsearch is a popular distributed data store which analyzes the JSON
documents clients have stored, and provides great search functionality over
them.

While it can operate on schema-less way, it also allows users to set mapping on
indices, through which you can define and optionally fixate the shape of JSON
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
  in mapping. If `"format"` is blank or includes `"epoch_millis"` or
  `"epoch_second"` the field also accepts JSON number as milliseconds or seconds
  since the epoch respectively.
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
type is not done automatically. The user code must explicitly handle them.

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
- The estime package is a collection of functions that helps the code generator
  to parse
  [the DateTimeFormatter formats](https://docs.oracle.com/javase/8/docs/api/java/time/format/DateTimeFormatter.html)
  and generate time.Time-based types that understand all possible date format
  defined in the mapping.
  - `DateTimeFormatter` defines optional section tokens, `[` and `]`. The estime
    package internally use the `optionalstring` sub package to break down
    optional section to all possible strings. For example, `ab[c[d]]` into
    `abcd`, `abc` and `ab`.
  - Since Go defines its specific tokens for time layout. 2006 or 06 for year,
    15 for hours and so on. The estime package converts Java time layout tokens
    into Go std layout tokens if and only if Go has counterparts for them.
    - It drops support for number of tokens including `G(era)`,
      `Q/q(quarter-of-year)`, `w(week-of-week-based-year)` and
      `W(week-of-month)`.
  - With this package, the code generator generates time.Time-based types.

## The Code generator

The code generator finally generates Go struct for easier generation /
consumption of JSON documents stored in a (cluster of) Elasticsearch
instance(s).

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
