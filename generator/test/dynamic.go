package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	gentypehelper "github.com/ngicks/estype/gentypehelper"
	elastic "github.com/ngicks/und/elastic"
	serde "github.com/ngicks/und/serde"
	"io"
	"sort"
)

type Dynamic struct {
	NestedInherit DynamicNestedInheritNested `json:"nested_inherit"`
	NestedRuntime DynamicNestedRuntimeNested `json:"nested_runtime"`
	NestedStrict  DynamicNestedStrictNested  `json:"nested_strict"`
	ObjectFalse   DynamicObjectFalseObject   `json:"object_false"`
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d Dynamic) ToRaw() DynamicRaw {
	return DynamicRaw{
		NestedInherit: gentypehelper.MapPlainToRawElastic[DynamicNestedInheritNestedRaw](d.NestedInherit),
		NestedRuntime: gentypehelper.MapPlainToRawElastic[DynamicNestedRuntimeNestedRaw](d.NestedRuntime),
		NestedStrict:  gentypehelper.MapPlainToRawElastic[DynamicNestedStrictNestedRaw](d.NestedStrict),
		ObjectFalse:   gentypehelper.MapPlainToRawElastic[DynamicObjectFalseObjectRaw](d.ObjectFalse),
	}
}

type DynamicRaw struct {
	NestedInherit elastic.Elastic[DynamicNestedInheritNestedRaw] `json:"nested_inherit"`
	NestedRuntime elastic.Elastic[DynamicNestedRuntimeNestedRaw] `json:"nested_runtime"`
	NestedStrict  elastic.Elastic[DynamicNestedStrictNestedRaw]  `json:"nested_strict"`
	ObjectFalse   elastic.Elastic[DynamicObjectFalseObjectRaw]   `json:"object_false"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d DynamicRaw) ToPlain() Dynamic {
	return Dynamic{
		NestedInherit: gentypehelper.MapElasticToPlainSingle[DynamicNestedInheritNested](d.NestedInherit),
		NestedRuntime: gentypehelper.MapElasticToPlainSingle[DynamicNestedRuntimeNested](d.NestedRuntime),
		NestedStrict:  gentypehelper.MapElasticToPlainSingle[DynamicNestedStrictNested](d.NestedStrict),
		ObjectFalse:   gentypehelper.MapElasticToPlainSingle[DynamicObjectFalseObject](d.ObjectFalse),
	}
}

type DynamicNestedInheritNested struct {
	Age  int32                          `json:"age"`
	Name DynamicNestedInheritNameObject `json:"name"`
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d DynamicNestedInheritNested) ToRaw() DynamicNestedInheritNestedRaw {
	return DynamicNestedInheritNestedRaw{
		Age:  gentypehelper.MapSingleValueToElastic[int32](d.Age),
		Name: gentypehelper.MapPlainToRawElastic[DynamicNestedInheritNameObjectRaw](d.Name),
	}
}

type DynamicNestedInheritNestedRaw struct {
	Age  elastic.Elastic[int32]                             `json:"age"`
	Name elastic.Elastic[DynamicNestedInheritNameObjectRaw] `json:"name"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d DynamicNestedInheritNestedRaw) ToPlain() DynamicNestedInheritNested {
	return DynamicNestedInheritNested{
		Age:  d.Age.ValueSingle(),
		Name: gentypehelper.MapElasticToPlainSingle[DynamicNestedInheritNameObject](d.Name),
	}
}

type DynamicNestedInheritNameObject struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d DynamicNestedInheritNameObject) ToRaw() DynamicNestedInheritNameObjectRaw {
	return DynamicNestedInheritNameObjectRaw{
		First: gentypehelper.MapSingleValueToElastic[string](d.First),
		Last:  gentypehelper.MapSingleValueToElastic[string](d.Last),
	}
}

type DynamicNestedInheritNameObjectRaw struct {
	First elastic.Elastic[string] `json:"first"`
	Last  elastic.Elastic[string] `json:"last"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d DynamicNestedInheritNameObjectRaw) ToPlain() DynamicNestedInheritNameObject {
	return DynamicNestedInheritNameObject{
		First: d.First.ValueSingle(),
		Last:  d.Last.ValueSingle(),
	}
}

type DynamicNestedRuntimeNested struct {
	Age                   int32                          `json:"age"`
	Name                  DynamicNestedRuntimeNameObject `json:"name"`
	AdditionalProperties_ map[string]any
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d DynamicNestedRuntimeNested) ToRaw() DynamicNestedRuntimeNestedRaw {
	return DynamicNestedRuntimeNestedRaw{
		Age:                   gentypehelper.MapSingleValueToElastic[int32](d.Age),
		Name:                  gentypehelper.MapPlainToRawElastic[DynamicNestedRuntimeNameObjectRaw](d.Name),
		AdditionalProperties_: d.AdditionalProperties_,
	}
}

// MarshalJSON implements json.Marshaler
// so that both known fields and additional properties are marshaled into a same JSON object.
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d DynamicNestedRuntimeNested) MarshalJSON() ([]byte, error) {
	buf := gentypehelper.GetBuf()
	defer gentypehelper.PutBuf(buf)
	var (
		bin []byte
		err error
	)
	buf.WriteByte('{')
	buf.WriteString("\"age\":")
	bin, err = json.Marshal(d.Age)
	if err != nil {
		return nil, err
	}
	buf.Write(bin)
	buf.WriteByte(',')
	buf.WriteString("\"name\":")
	bin, err = json.Marshal(d.Name)
	if err != nil {
		return nil, err
	}
	buf.Write(bin)
	buf.WriteByte(',')
	keys := make([]string, 0, len(d.AdditionalProperties_))
	for k := range d.AdditionalProperties_ {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		bin, err = json.Marshal(d.AdditionalProperties_[key])
		if err != nil {
			return nil, err
		}
		buf.WriteByte('"')
		json.HTMLEscape(buf, []byte(key))
		buf.WriteString("\":")
		buf.Write(bin)
		buf.WriteByte(',')
	}
	if buf.Len() > 1 {
		buf.Truncate(buf.Len() - 1)
	}
	buf.WriteByte('}')
	return append([]byte{}, buf.Bytes()...), nil
}

// UnmarshalJSON implements json.Unmarshaler
// to add the special handling rule where
// additional fields in the input JSON object are stored into the AdditionalProperties_ field
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d *DynamicNestedRuntimeNested) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	token, err := dec.Token()
	if err != nil {
		return err
	}
	if token != json.Delim('{') {
		return fmt.Errorf("unknown token. Assuming the input is a JSON object, but received wrong delim = %s", token)
	}
	firstWriteToAdditionalProp := true
	for {
		token, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		switch token {
		case "age":
			err = dec.Decode(&d.Age)
		case "name":
			err = dec.Decode(&d.Name)
		default:
			if key, ok := token.(string); ok {
				var o any
				err = dec.Decode(&o)
				if err != nil {
					return err
				}
				// map re-initialization is deferred until at least a successful decode.
				if firstWriteToAdditionalProp {
					firstWriteToAdditionalProp = false
					d.AdditionalProperties_ = make(map[string]any)
				}
				d.AdditionalProperties_[key] = o
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type DynamicNestedRuntimeNestedRaw struct {
	Age                   elastic.Elastic[int32]                             `json:"age"`
	Name                  elastic.Elastic[DynamicNestedRuntimeNameObjectRaw] `json:"name"`
	AdditionalProperties_ map[string]any
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d DynamicNestedRuntimeNestedRaw) ToPlain() DynamicNestedRuntimeNested {
	return DynamicNestedRuntimeNested{
		Age:                   d.Age.ValueSingle(),
		Name:                  gentypehelper.MapElasticToPlainSingle[DynamicNestedRuntimeNameObject](d.Name),
		AdditionalProperties_: d.AdditionalProperties_,
	}
}

// MarshalJSON implements json.Marshaler
// so that both known fields and additional properties are marshaled into a same JSON object.
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d DynamicNestedRuntimeNestedRaw) MarshalJSON() ([]byte, error) {
	buf := gentypehelper.GetBuf()
	defer gentypehelper.PutBuf(buf)
	var (
		bin []byte
		err error
	)
	buf.WriteByte('{')
	buf.WriteString("\"age\":")
	bin, err = serde.Marshal(d.Age)
	if err != nil {
		return nil, err
	}
	buf.Write(bin)
	buf.WriteByte(',')
	buf.WriteString("\"name\":")
	bin, err = serde.Marshal(d.Name)
	if err != nil {
		return nil, err
	}
	buf.Write(bin)
	buf.WriteByte(',')
	keys := make([]string, 0, len(d.AdditionalProperties_))
	for k := range d.AdditionalProperties_ {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		bin, err = serde.Marshal(d.AdditionalProperties_[key])
		if err != nil {
			return nil, err
		}
		buf.WriteByte('"')
		json.HTMLEscape(buf, []byte(key))
		buf.WriteString("\":")
		buf.Write(bin)
		buf.WriteByte(',')
	}
	if buf.Len() > 1 {
		buf.Truncate(buf.Len() - 1)
	}
	buf.WriteByte('}')
	return append([]byte{}, buf.Bytes()...), nil
}

// UnmarshalJSON implements json.Unmarshaler
// to add the special handling rule where
// additional fields in the input JSON object are stored into the AdditionalProperties_ field
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d *DynamicNestedRuntimeNestedRaw) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	token, err := dec.Token()
	if err != nil {
		return err
	}
	if token != json.Delim('{') {
		return fmt.Errorf("unknown token. Assuming the input is a JSON object, but received wrong delim = %s", token)
	}
	firstWriteToAdditionalProp := true
	for {
		token, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		switch token {
		case "age":
			err = dec.Decode(&d.Age)
		case "name":
			err = dec.Decode(&d.Name)
		default:
			if key, ok := token.(string); ok {
				var o any
				err = dec.Decode(&o)
				if err != nil {
					return err
				}
				// map re-initialization is deferred until at least a successful decode.
				if firstWriteToAdditionalProp {
					firstWriteToAdditionalProp = false
					d.AdditionalProperties_ = make(map[string]any)
				}
				d.AdditionalProperties_[key] = o
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type DynamicNestedRuntimeNameObject struct {
	First                 string `json:"first"`
	Last                  string `json:"last"`
	AdditionalProperties_ map[string]any
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d DynamicNestedRuntimeNameObject) ToRaw() DynamicNestedRuntimeNameObjectRaw {
	return DynamicNestedRuntimeNameObjectRaw{
		First:                 gentypehelper.MapSingleValueToElastic[string](d.First),
		Last:                  gentypehelper.MapSingleValueToElastic[string](d.Last),
		AdditionalProperties_: d.AdditionalProperties_,
	}
}

// MarshalJSON implements json.Marshaler
// so that both known fields and additional properties are marshaled into a same JSON object.
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d DynamicNestedRuntimeNameObject) MarshalJSON() ([]byte, error) {
	buf := gentypehelper.GetBuf()
	defer gentypehelper.PutBuf(buf)
	var (
		bin []byte
		err error
	)
	buf.WriteByte('{')
	buf.WriteString("\"first\":")
	bin, err = json.Marshal(d.First)
	if err != nil {
		return nil, err
	}
	buf.Write(bin)
	buf.WriteByte(',')
	buf.WriteString("\"last\":")
	bin, err = json.Marshal(d.Last)
	if err != nil {
		return nil, err
	}
	buf.Write(bin)
	buf.WriteByte(',')
	keys := make([]string, 0, len(d.AdditionalProperties_))
	for k := range d.AdditionalProperties_ {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		bin, err = json.Marshal(d.AdditionalProperties_[key])
		if err != nil {
			return nil, err
		}
		buf.WriteByte('"')
		json.HTMLEscape(buf, []byte(key))
		buf.WriteString("\":")
		buf.Write(bin)
		buf.WriteByte(',')
	}
	if buf.Len() > 1 {
		buf.Truncate(buf.Len() - 1)
	}
	buf.WriteByte('}')
	return append([]byte{}, buf.Bytes()...), nil
}

// UnmarshalJSON implements json.Unmarshaler
// to add the special handling rule where
// additional fields in the input JSON object are stored into the AdditionalProperties_ field
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d *DynamicNestedRuntimeNameObject) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	token, err := dec.Token()
	if err != nil {
		return err
	}
	if token != json.Delim('{') {
		return fmt.Errorf("unknown token. Assuming the input is a JSON object, but received wrong delim = %s", token)
	}
	firstWriteToAdditionalProp := true
	for {
		token, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		switch token {
		case "first":
			err = dec.Decode(&d.First)
		case "last":
			err = dec.Decode(&d.Last)
		default:
			if key, ok := token.(string); ok {
				var o any
				err = dec.Decode(&o)
				if err != nil {
					return err
				}
				// map re-initialization is deferred until at least a successful decode.
				if firstWriteToAdditionalProp {
					firstWriteToAdditionalProp = false
					d.AdditionalProperties_ = make(map[string]any)
				}
				d.AdditionalProperties_[key] = o
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type DynamicNestedRuntimeNameObjectRaw struct {
	First                 elastic.Elastic[string] `json:"first"`
	Last                  elastic.Elastic[string] `json:"last"`
	AdditionalProperties_ map[string]any
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d DynamicNestedRuntimeNameObjectRaw) ToPlain() DynamicNestedRuntimeNameObject {
	return DynamicNestedRuntimeNameObject{
		First:                 d.First.ValueSingle(),
		Last:                  d.Last.ValueSingle(),
		AdditionalProperties_: d.AdditionalProperties_,
	}
}

// MarshalJSON implements json.Marshaler
// so that both known fields and additional properties are marshaled into a same JSON object.
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d DynamicNestedRuntimeNameObjectRaw) MarshalJSON() ([]byte, error) {
	buf := gentypehelper.GetBuf()
	defer gentypehelper.PutBuf(buf)
	var (
		bin []byte
		err error
	)
	buf.WriteByte('{')
	buf.WriteString("\"first\":")
	bin, err = serde.Marshal(d.First)
	if err != nil {
		return nil, err
	}
	buf.Write(bin)
	buf.WriteByte(',')
	buf.WriteString("\"last\":")
	bin, err = serde.Marshal(d.Last)
	if err != nil {
		return nil, err
	}
	buf.Write(bin)
	buf.WriteByte(',')
	keys := make([]string, 0, len(d.AdditionalProperties_))
	for k := range d.AdditionalProperties_ {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		bin, err = serde.Marshal(d.AdditionalProperties_[key])
		if err != nil {
			return nil, err
		}
		buf.WriteByte('"')
		json.HTMLEscape(buf, []byte(key))
		buf.WriteString("\":")
		buf.Write(bin)
		buf.WriteByte(',')
	}
	if buf.Len() > 1 {
		buf.Truncate(buf.Len() - 1)
	}
	buf.WriteByte('}')
	return append([]byte{}, buf.Bytes()...), nil
}

// UnmarshalJSON implements json.Unmarshaler
// to add the special handling rule where
// additional fields in the input JSON object are stored into the AdditionalProperties_ field
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d *DynamicNestedRuntimeNameObjectRaw) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	token, err := dec.Token()
	if err != nil {
		return err
	}
	if token != json.Delim('{') {
		return fmt.Errorf("unknown token. Assuming the input is a JSON object, but received wrong delim = %s", token)
	}
	firstWriteToAdditionalProp := true
	for {
		token, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		switch token {
		case "first":
			err = dec.Decode(&d.First)
		case "last":
			err = dec.Decode(&d.Last)
		default:
			if key, ok := token.(string); ok {
				var o any
				err = dec.Decode(&o)
				if err != nil {
					return err
				}
				// map re-initialization is deferred until at least a successful decode.
				if firstWriteToAdditionalProp {
					firstWriteToAdditionalProp = false
					d.AdditionalProperties_ = make(map[string]any)
				}
				d.AdditionalProperties_[key] = o
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type DynamicNestedStrictNested struct {
	Age  int32                         `json:"age"`
	Name DynamicNestedStrictNameObject `json:"name"`
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d DynamicNestedStrictNested) ToRaw() DynamicNestedStrictNestedRaw {
	return DynamicNestedStrictNestedRaw{
		Age:  gentypehelper.MapSingleValueToElastic[int32](d.Age),
		Name: gentypehelper.MapPlainToRawElastic[DynamicNestedStrictNameObjectRaw](d.Name),
	}
}

type DynamicNestedStrictNestedRaw struct {
	Age  elastic.Elastic[int32]                            `json:"age"`
	Name elastic.Elastic[DynamicNestedStrictNameObjectRaw] `json:"name"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d DynamicNestedStrictNestedRaw) ToPlain() DynamicNestedStrictNested {
	return DynamicNestedStrictNested{
		Age:  d.Age.ValueSingle(),
		Name: gentypehelper.MapElasticToPlainSingle[DynamicNestedStrictNameObject](d.Name),
	}
}

type DynamicNestedStrictNameObject struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d DynamicNestedStrictNameObject) ToRaw() DynamicNestedStrictNameObjectRaw {
	return DynamicNestedStrictNameObjectRaw{
		First: gentypehelper.MapSingleValueToElastic[string](d.First),
		Last:  gentypehelper.MapSingleValueToElastic[string](d.Last),
	}
}

type DynamicNestedStrictNameObjectRaw struct {
	First elastic.Elastic[string] `json:"first"`
	Last  elastic.Elastic[string] `json:"last"`
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d DynamicNestedStrictNameObjectRaw) ToPlain() DynamicNestedStrictNameObject {
	return DynamicNestedStrictNameObject{
		First: d.First.ValueSingle(),
		Last:  d.Last.ValueSingle(),
	}
}

type DynamicObjectFalseObject struct {
	Age                   int32                        `json:"age"`
	Name                  DynamicObjectFalseNameObject `json:"name"`
	AdditionalProperties_ map[string]any
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d DynamicObjectFalseObject) ToRaw() DynamicObjectFalseObjectRaw {
	return DynamicObjectFalseObjectRaw{
		Age:                   gentypehelper.MapSingleValueToElastic[int32](d.Age),
		Name:                  gentypehelper.MapPlainToRawElastic[DynamicObjectFalseNameObjectRaw](d.Name),
		AdditionalProperties_: d.AdditionalProperties_,
	}
}

// MarshalJSON implements json.Marshaler
// so that both known fields and additional properties are marshaled into a same JSON object.
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d DynamicObjectFalseObject) MarshalJSON() ([]byte, error) {
	buf := gentypehelper.GetBuf()
	defer gentypehelper.PutBuf(buf)
	var (
		bin []byte
		err error
	)
	buf.WriteByte('{')
	buf.WriteString("\"age\":")
	bin, err = json.Marshal(d.Age)
	if err != nil {
		return nil, err
	}
	buf.Write(bin)
	buf.WriteByte(',')
	buf.WriteString("\"name\":")
	bin, err = json.Marshal(d.Name)
	if err != nil {
		return nil, err
	}
	buf.Write(bin)
	buf.WriteByte(',')
	keys := make([]string, 0, len(d.AdditionalProperties_))
	for k := range d.AdditionalProperties_ {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		bin, err = json.Marshal(d.AdditionalProperties_[key])
		if err != nil {
			return nil, err
		}
		buf.WriteByte('"')
		json.HTMLEscape(buf, []byte(key))
		buf.WriteString("\":")
		buf.Write(bin)
		buf.WriteByte(',')
	}
	if buf.Len() > 1 {
		buf.Truncate(buf.Len() - 1)
	}
	buf.WriteByte('}')
	return append([]byte{}, buf.Bytes()...), nil
}

// UnmarshalJSON implements json.Unmarshaler
// to add the special handling rule where
// additional fields in the input JSON object are stored into the AdditionalProperties_ field
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d *DynamicObjectFalseObject) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	token, err := dec.Token()
	if err != nil {
		return err
	}
	if token != json.Delim('{') {
		return fmt.Errorf("unknown token. Assuming the input is a JSON object, but received wrong delim = %s", token)
	}
	firstWriteToAdditionalProp := true
	for {
		token, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		switch token {
		case "age":
			err = dec.Decode(&d.Age)
		case "name":
			err = dec.Decode(&d.Name)
		default:
			if key, ok := token.(string); ok {
				var o any
				err = dec.Decode(&o)
				if err != nil {
					return err
				}
				// map re-initialization is deferred until at least a successful decode.
				if firstWriteToAdditionalProp {
					firstWriteToAdditionalProp = false
					d.AdditionalProperties_ = make(map[string]any)
				}
				d.AdditionalProperties_[key] = o
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type DynamicObjectFalseObjectRaw struct {
	Age                   elastic.Elastic[int32]                           `json:"age"`
	Name                  elastic.Elastic[DynamicObjectFalseNameObjectRaw] `json:"name"`
	AdditionalProperties_ map[string]any
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d DynamicObjectFalseObjectRaw) ToPlain() DynamicObjectFalseObject {
	return DynamicObjectFalseObject{
		Age:                   d.Age.ValueSingle(),
		Name:                  gentypehelper.MapElasticToPlainSingle[DynamicObjectFalseNameObject](d.Name),
		AdditionalProperties_: d.AdditionalProperties_,
	}
}

// MarshalJSON implements json.Marshaler
// so that both known fields and additional properties are marshaled into a same JSON object.
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d DynamicObjectFalseObjectRaw) MarshalJSON() ([]byte, error) {
	buf := gentypehelper.GetBuf()
	defer gentypehelper.PutBuf(buf)
	var (
		bin []byte
		err error
	)
	buf.WriteByte('{')
	buf.WriteString("\"age\":")
	bin, err = serde.Marshal(d.Age)
	if err != nil {
		return nil, err
	}
	buf.Write(bin)
	buf.WriteByte(',')
	buf.WriteString("\"name\":")
	bin, err = serde.Marshal(d.Name)
	if err != nil {
		return nil, err
	}
	buf.Write(bin)
	buf.WriteByte(',')
	keys := make([]string, 0, len(d.AdditionalProperties_))
	for k := range d.AdditionalProperties_ {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		bin, err = serde.Marshal(d.AdditionalProperties_[key])
		if err != nil {
			return nil, err
		}
		buf.WriteByte('"')
		json.HTMLEscape(buf, []byte(key))
		buf.WriteString("\":")
		buf.Write(bin)
		buf.WriteByte(',')
	}
	if buf.Len() > 1 {
		buf.Truncate(buf.Len() - 1)
	}
	buf.WriteByte('}')
	return append([]byte{}, buf.Bytes()...), nil
}

// UnmarshalJSON implements json.Unmarshaler
// to add the special handling rule where
// additional fields in the input JSON object are stored into the AdditionalProperties_ field
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d *DynamicObjectFalseObjectRaw) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	token, err := dec.Token()
	if err != nil {
		return err
	}
	if token != json.Delim('{') {
		return fmt.Errorf("unknown token. Assuming the input is a JSON object, but received wrong delim = %s", token)
	}
	firstWriteToAdditionalProp := true
	for {
		token, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		switch token {
		case "age":
			err = dec.Decode(&d.Age)
		case "name":
			err = dec.Decode(&d.Name)
		default:
			if key, ok := token.(string); ok {
				var o any
				err = dec.Decode(&o)
				if err != nil {
					return err
				}
				// map re-initialization is deferred until at least a successful decode.
				if firstWriteToAdditionalProp {
					firstWriteToAdditionalProp = false
					d.AdditionalProperties_ = make(map[string]any)
				}
				d.AdditionalProperties_[key] = o
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type DynamicObjectFalseNameObject struct {
	First                 string `json:"first"`
	Last                  string `json:"last"`
	AdditionalProperties_ map[string]any
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d DynamicObjectFalseNameObject) ToRaw() DynamicObjectFalseNameObjectRaw {
	return DynamicObjectFalseNameObjectRaw{
		First:                 gentypehelper.MapSingleValueToElastic[string](d.First),
		Last:                  gentypehelper.MapSingleValueToElastic[string](d.Last),
		AdditionalProperties_: d.AdditionalProperties_,
	}
}

// MarshalJSON implements json.Marshaler
// so that both known fields and additional properties are marshaled into a same JSON object.
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d DynamicObjectFalseNameObject) MarshalJSON() ([]byte, error) {
	buf := gentypehelper.GetBuf()
	defer gentypehelper.PutBuf(buf)
	var (
		bin []byte
		err error
	)
	buf.WriteByte('{')
	buf.WriteString("\"first\":")
	bin, err = json.Marshal(d.First)
	if err != nil {
		return nil, err
	}
	buf.Write(bin)
	buf.WriteByte(',')
	buf.WriteString("\"last\":")
	bin, err = json.Marshal(d.Last)
	if err != nil {
		return nil, err
	}
	buf.Write(bin)
	buf.WriteByte(',')
	keys := make([]string, 0, len(d.AdditionalProperties_))
	for k := range d.AdditionalProperties_ {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		bin, err = json.Marshal(d.AdditionalProperties_[key])
		if err != nil {
			return nil, err
		}
		buf.WriteByte('"')
		json.HTMLEscape(buf, []byte(key))
		buf.WriteString("\":")
		buf.Write(bin)
		buf.WriteByte(',')
	}
	if buf.Len() > 1 {
		buf.Truncate(buf.Len() - 1)
	}
	buf.WriteByte('}')
	return append([]byte{}, buf.Bytes()...), nil
}

// UnmarshalJSON implements json.Unmarshaler
// to add the special handling rule where
// additional fields in the input JSON object are stored into the AdditionalProperties_ field
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d *DynamicObjectFalseNameObject) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	token, err := dec.Token()
	if err != nil {
		return err
	}
	if token != json.Delim('{') {
		return fmt.Errorf("unknown token. Assuming the input is a JSON object, but received wrong delim = %s", token)
	}
	firstWriteToAdditionalProp := true
	for {
		token, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		switch token {
		case "first":
			err = dec.Decode(&d.First)
		case "last":
			err = dec.Decode(&d.Last)
		default:
			if key, ok := token.(string); ok {
				var o any
				err = dec.Decode(&o)
				if err != nil {
					return err
				}
				// map re-initialization is deferred until at least a successful decode.
				if firstWriteToAdditionalProp {
					firstWriteToAdditionalProp = false
					d.AdditionalProperties_ = make(map[string]any)
				}
				d.AdditionalProperties_[key] = o
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type DynamicObjectFalseNameObjectRaw struct {
	First                 elastic.Elastic[string] `json:"first"`
	Last                  elastic.Elastic[string] `json:"last"`
	AdditionalProperties_ map[string]any
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d DynamicObjectFalseNameObjectRaw) ToPlain() DynamicObjectFalseNameObject {
	return DynamicObjectFalseNameObject{
		First:                 d.First.ValueSingle(),
		Last:                  d.Last.ValueSingle(),
		AdditionalProperties_: d.AdditionalProperties_,
	}
}

// MarshalJSON implements json.Marshaler
// so that both known fields and additional properties are marshaled into a same JSON object.
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d DynamicObjectFalseNameObjectRaw) MarshalJSON() ([]byte, error) {
	buf := gentypehelper.GetBuf()
	defer gentypehelper.PutBuf(buf)
	var (
		bin []byte
		err error
	)
	buf.WriteByte('{')
	buf.WriteString("\"first\":")
	bin, err = serde.Marshal(d.First)
	if err != nil {
		return nil, err
	}
	buf.Write(bin)
	buf.WriteByte(',')
	buf.WriteString("\"last\":")
	bin, err = serde.Marshal(d.Last)
	if err != nil {
		return nil, err
	}
	buf.Write(bin)
	buf.WriteByte(',')
	keys := make([]string, 0, len(d.AdditionalProperties_))
	for k := range d.AdditionalProperties_ {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		bin, err = serde.Marshal(d.AdditionalProperties_[key])
		if err != nil {
			return nil, err
		}
		buf.WriteByte('"')
		json.HTMLEscape(buf, []byte(key))
		buf.WriteString("\":")
		buf.Write(bin)
		buf.WriteByte(',')
	}
	if buf.Len() > 1 {
		buf.Truncate(buf.Len() - 1)
	}
	buf.WriteByte('}')
	return append([]byte{}, buf.Bytes()...), nil
}

// UnmarshalJSON implements json.Unmarshaler
// to add the special handling rule where
// additional fields in the input JSON object are stored into the AdditionalProperties_ field
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d *DynamicObjectFalseNameObjectRaw) UnmarshalJSON(data []byte) error {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	token, err := dec.Token()
	if err != nil {
		return err
	}
	if token != json.Delim('{') {
		return fmt.Errorf("unknown token. Assuming the input is a JSON object, but received wrong delim = %s", token)
	}
	firstWriteToAdditionalProp := true
	for {
		token, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		switch token {
		case "first":
			err = dec.Decode(&d.First)
		case "last":
			err = dec.Decode(&d.Last)
		default:
			if key, ok := token.(string); ok {
				var o any
				err = dec.Decode(&o)
				if err != nil {
					return err
				}
				// map re-initialization is deferred until at least a successful decode.
				if firstWriteToAdditionalProp {
					firstWriteToAdditionalProp = false
					d.AdditionalProperties_ = make(map[string]any)
				}
				d.AdditionalProperties_[key] = o
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}
