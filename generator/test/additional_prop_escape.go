package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	gentypehelper "github.com/ngicks/estype/gentypehelper"
	elastic "github.com/ngicks/und/sliceund/elastic"
	"io"
	"reflect"
	"sort"
)

type AdditionalPropEscape struct {
	U003chmu003e          AdditionalPropEscapeU003chmu003eObject `json:"<hm>"`
	U2728                 AdditionalPropEscapeU2728Object        `json:"✨"`
	AdditionalProperties_ map[string]any
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AdditionalPropEscape) ToRaw() AdditionalPropEscapeRaw {
	return AdditionalPropEscapeRaw{
		U003chmu003e:          gentypehelper.MapPlainToRawElastic[AdditionalPropEscapeU003chmu003eObjectRaw](d.U003chmu003e),
		U2728:                 gentypehelper.MapPlainToRawElastic[AdditionalPropEscapeU2728ObjectRaw](d.U2728),
		AdditionalProperties_: d.AdditionalProperties_,
	}
}

// MarshalJSON implements json.Marshaler
// so that both known fields and additional properties are marshaled into a same JSON object.
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d AdditionalPropEscape) MarshalJSON() ([]byte, error) {
	buf := gentypehelper.GetBuf()
	defer gentypehelper.PutBuf(buf)
	var (
		bin []byte
		err error
	)
	buf.WriteByte('{')
	buf.WriteString("\"\\u003chm\\u003e\":")
	bin, err = json.Marshal(d.U003chmu003e)
	if err != nil {
		return nil, err
	}
	buf.Write(bin)
	buf.WriteByte(',')
	buf.WriteString("\"✨\":")
	bin, err = json.Marshal(d.U2728)
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
func (d *AdditionalPropEscape) UnmarshalJSON(data []byte) error {
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
		case "<hm>":
			err = dec.Decode(&d.U003chmu003e)
		case "✨":
			err = dec.Decode(&d.U2728)
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

type AdditionalPropEscapeRaw struct {
	U003chmu003e          elastic.Elastic[AdditionalPropEscapeU003chmu003eObjectRaw] `json:"<hm>,omitempty"`
	U2728                 elastic.Elastic[AdditionalPropEscapeU2728ObjectRaw]        `json:"✨,omitempty"`
	AdditionalProperties_ map[string]any
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AdditionalPropEscapeRaw) ToPlain() AdditionalPropEscape {
	return AdditionalPropEscape{
		U003chmu003e:          gentypehelper.MapElasticToPlainSingle[AdditionalPropEscapeU003chmu003eObject](d.U003chmu003e),
		U2728:                 gentypehelper.MapElasticToPlainSingle[AdditionalPropEscapeU2728Object](d.U2728),
		AdditionalProperties_: d.AdditionalProperties_,
	}
}

// MarshalJSON implements json.Marshaler
// so that both known fields and additional properties are marshaled into a same JSON object.
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d AdditionalPropEscapeRaw) MarshalJSON() ([]byte, error) {
	buf := gentypehelper.GetBuf()
	defer gentypehelper.PutBuf(buf)
	var (
		bin []byte
		err error
	)
	buf.WriteByte('{')
	// This field is tagged with ",omitempty".
	if !reflect.ValueOf(d.U003chmu003e).IsZero() {
		buf.WriteString("\"\\u003chm\\u003e\":")
		bin, err = json.Marshal(d.U003chmu003e)
		if err != nil {
			return nil, err
		}
		buf.Write(bin)
		buf.WriteByte(',')
	}
	// This field is tagged with ",omitempty".
	if !reflect.ValueOf(d.U2728).IsZero() {
		buf.WriteString("\"✨\":")
		bin, err = json.Marshal(d.U2728)
		if err != nil {
			return nil, err
		}
		buf.Write(bin)
		buf.WriteByte(',')
	}
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
func (d *AdditionalPropEscapeRaw) UnmarshalJSON(data []byte) error {
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
		case "<hm>":
			err = dec.Decode(&d.U003chmu003e)
		case "✨":
			err = dec.Decode(&d.U2728)
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

type AdditionalPropEscapeU003chmu003eObject struct {
	U0026mahu0026         string    `json:"&mah&"`
	FooBar__              *struct{} `json:"__foo_bar,omitempty"`
	AdditionalProperties_ map[string]any
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AdditionalPropEscapeU003chmu003eObject) ToRaw() AdditionalPropEscapeU003chmu003eObjectRaw {
	return AdditionalPropEscapeU003chmu003eObjectRaw{
		U0026mahu0026:         gentypehelper.MapSingleValueToElastic[string](d.U0026mahu0026),
		AdditionalProperties_: d.AdditionalProperties_,
	}
}

// MarshalJSON implements json.Marshaler
// so that both known fields and additional properties are marshaled into a same JSON object.
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d AdditionalPropEscapeU003chmu003eObject) MarshalJSON() ([]byte, error) {
	buf := gentypehelper.GetBuf()
	defer gentypehelper.PutBuf(buf)
	var (
		bin []byte
		err error
	)
	buf.WriteByte('{')
	buf.WriteString("\"\\u0026mah\\u0026\":")
	bin, err = json.Marshal(d.U0026mahu0026)
	if err != nil {
		return nil, err
	}
	buf.Write(bin)
	buf.WriteByte(',')
	// This field is tagged with ",omitempty".
	if !reflect.ValueOf(d.FooBar__).IsZero() {
		buf.WriteString("\"__foo_bar\":")
		bin, err = json.Marshal(d.FooBar__)
		if err != nil {
			return nil, err
		}
		buf.Write(bin)
		buf.WriteByte(',')
	}
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
func (d *AdditionalPropEscapeU003chmu003eObject) UnmarshalJSON(data []byte) error {
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
		case "&mah&":
			err = dec.Decode(&d.U0026mahu0026)
		case "__foo_bar":
			err = dec.Decode(&d.FooBar__)
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

type AdditionalPropEscapeU003chmu003eObjectRaw struct {
	U0026mahu0026         elastic.Elastic[string]    `json:"&mah&,omitempty"`
	FooBar__              elastic.Elastic[*struct{}] `json:"__foo_bar,omitempty"`
	AdditionalProperties_ map[string]any
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AdditionalPropEscapeU003chmu003eObjectRaw) ToPlain() AdditionalPropEscapeU003chmu003eObject {
	return AdditionalPropEscapeU003chmu003eObject{
		U0026mahu0026:         d.U0026mahu0026.Value(),
		AdditionalProperties_: d.AdditionalProperties_,
	}
}

// MarshalJSON implements json.Marshaler
// so that both known fields and additional properties are marshaled into a same JSON object.
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d AdditionalPropEscapeU003chmu003eObjectRaw) MarshalJSON() ([]byte, error) {
	buf := gentypehelper.GetBuf()
	defer gentypehelper.PutBuf(buf)
	var (
		bin []byte
		err error
	)
	buf.WriteByte('{')
	// This field is tagged with ",omitempty".
	if !reflect.ValueOf(d.U0026mahu0026).IsZero() {
		buf.WriteString("\"\\u0026mah\\u0026\":")
		bin, err = json.Marshal(d.U0026mahu0026)
		if err != nil {
			return nil, err
		}
		buf.Write(bin)
		buf.WriteByte(',')
	}
	// This field is tagged with ",omitempty".
	if !reflect.ValueOf(d.FooBar__).IsZero() {
		buf.WriteString("\"__foo_bar\":")
		bin, err = json.Marshal(d.FooBar__)
		if err != nil {
			return nil, err
		}
		buf.Write(bin)
		buf.WriteByte(',')
	}
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
func (d *AdditionalPropEscapeU003chmu003eObjectRaw) UnmarshalJSON(data []byte) error {
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
		case "&mah&":
			err = dec.Decode(&d.U0026mahu0026)
		case "__foo_bar":
			err = dec.Decode(&d.FooBar__)
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

type AdditionalPropEscapeU2728Object struct {
	Yay                   string `json:"yay"`
	AdditionalProperties_ map[string]any
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AdditionalPropEscapeU2728Object) ToRaw() AdditionalPropEscapeU2728ObjectRaw {
	return AdditionalPropEscapeU2728ObjectRaw{
		Yay:                   gentypehelper.MapSingleValueToElastic[string](d.Yay),
		AdditionalProperties_: d.AdditionalProperties_,
	}
}

// MarshalJSON implements json.Marshaler
// so that both known fields and additional properties are marshaled into a same JSON object.
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d AdditionalPropEscapeU2728Object) MarshalJSON() ([]byte, error) {
	buf := gentypehelper.GetBuf()
	defer gentypehelper.PutBuf(buf)
	var (
		bin []byte
		err error
	)
	buf.WriteByte('{')
	buf.WriteString("\"yay\":")
	bin, err = json.Marshal(d.Yay)
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
func (d *AdditionalPropEscapeU2728Object) UnmarshalJSON(data []byte) error {
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
		case "yay":
			err = dec.Decode(&d.Yay)
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

type AdditionalPropEscapeU2728ObjectRaw struct {
	Yay                   elastic.Elastic[string] `json:"yay,omitempty"`
	AdditionalProperties_ map[string]any
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AdditionalPropEscapeU2728ObjectRaw) ToPlain() AdditionalPropEscapeU2728Object {
	return AdditionalPropEscapeU2728Object{
		Yay:                   d.Yay.Value(),
		AdditionalProperties_: d.AdditionalProperties_,
	}
}

// MarshalJSON implements json.Marshaler
// so that both known fields and additional properties are marshaled into a same JSON object.
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d AdditionalPropEscapeU2728ObjectRaw) MarshalJSON() ([]byte, error) {
	buf := gentypehelper.GetBuf()
	defer gentypehelper.PutBuf(buf)
	var (
		bin []byte
		err error
	)
	buf.WriteByte('{')
	// This field is tagged with ",omitempty".
	if !reflect.ValueOf(d.Yay).IsZero() {
		buf.WriteString("\"yay\":")
		bin, err = json.Marshal(d.Yay)
		if err != nil {
			return nil, err
		}
		buf.Write(bin)
		buf.WriteByte(',')
	}
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
func (d *AdditionalPropEscapeU2728ObjectRaw) UnmarshalJSON(data []byte) error {
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
		case "yay":
			err = dec.Decode(&d.Yay)
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
