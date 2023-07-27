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

type AddtionalPropEscape struct {
	U003chmu003e          AddtionalPropEscapeU003chmu003eObject `json:"<hm>"`
	Ue29ca8               AddtionalPropEscapeUe29ca8Object      `json:"✨"`
	AdditionalProperties_ map[string]any
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AddtionalPropEscape) ToRaw() AddtionalPropEscapeRaw {
	return AddtionalPropEscapeRaw{
		U003chmu003e:          gentypehelper.MapPlainToRawElastic[AddtionalPropEscapeU003chmu003eObjectRaw](d.U003chmu003e),
		Ue29ca8:               gentypehelper.MapPlainToRawElastic[AddtionalPropEscapeUe29ca8ObjectRaw](d.Ue29ca8),
		AdditionalProperties_: d.AdditionalProperties_,
	}
}

// MarshalJSON implements json.Marshaler
// so that both known fields and additional properties are marshaled into a same JSON object.
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d AddtionalPropEscape) MarshalJSON() ([]byte, error) {
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
	bin, err = json.Marshal(d.Ue29ca8)
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
func (d *AddtionalPropEscape) UnmarshalJSON(data []byte) error {
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
			err = dec.Decode(&d.Ue29ca8)
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

type AddtionalPropEscapeRaw struct {
	U003chmu003e          elastic.Elastic[AddtionalPropEscapeU003chmu003eObjectRaw] `json:"<hm>"`
	Ue29ca8               elastic.Elastic[AddtionalPropEscapeUe29ca8ObjectRaw]      `json:"✨"`
	AdditionalProperties_ map[string]any
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AddtionalPropEscapeRaw) ToPlain() AddtionalPropEscape {
	return AddtionalPropEscape{
		U003chmu003e:          gentypehelper.MapElasticToPlainSingle[AddtionalPropEscapeU003chmu003eObject](d.U003chmu003e),
		Ue29ca8:               gentypehelper.MapElasticToPlainSingle[AddtionalPropEscapeUe29ca8Object](d.Ue29ca8),
		AdditionalProperties_: d.AdditionalProperties_,
	}
}

// MarshalJSON implements json.Marshaler
// so that both known fields and additional properties are marshaled into a same JSON object.
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d AddtionalPropEscapeRaw) MarshalJSON() ([]byte, error) {
	buf := gentypehelper.GetBuf()
	defer gentypehelper.PutBuf(buf)
	var (
		bin []byte
		err error
	)
	buf.WriteByte('{')
	buf.WriteString("\"\\u003chm\\u003e\":")
	bin, err = serde.Marshal(d.U003chmu003e)
	if err != nil {
		return nil, err
	}
	buf.Write(bin)
	buf.WriteByte(',')
	buf.WriteString("\"✨\":")
	bin, err = serde.Marshal(d.Ue29ca8)
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
func (d *AddtionalPropEscapeRaw) UnmarshalJSON(data []byte) error {
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
			err = dec.Decode(&d.Ue29ca8)
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

type AddtionalPropEscapeU003chmu003eObject struct {
	U0026mahu0026         string `json:"&mah&"`
	AdditionalProperties_ map[string]any
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AddtionalPropEscapeU003chmu003eObject) ToRaw() AddtionalPropEscapeU003chmu003eObjectRaw {
	return AddtionalPropEscapeU003chmu003eObjectRaw{
		U0026mahu0026:         gentypehelper.MapSingleValueToElastic[string](d.U0026mahu0026),
		AdditionalProperties_: d.AdditionalProperties_,
	}
}

// MarshalJSON implements json.Marshaler
// so that both known fields and additional properties are marshaled into a same JSON object.
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d AddtionalPropEscapeU003chmu003eObject) MarshalJSON() ([]byte, error) {
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
func (d *AddtionalPropEscapeU003chmu003eObject) UnmarshalJSON(data []byte) error {
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

type AddtionalPropEscapeU003chmu003eObjectRaw struct {
	U0026mahu0026         elastic.Elastic[string] `json:"&mah&"`
	AdditionalProperties_ map[string]any
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AddtionalPropEscapeU003chmu003eObjectRaw) ToPlain() AddtionalPropEscapeU003chmu003eObject {
	return AddtionalPropEscapeU003chmu003eObject{
		U0026mahu0026:         d.U0026mahu0026.ValueSingle(),
		AdditionalProperties_: d.AdditionalProperties_,
	}
}

// MarshalJSON implements json.Marshaler
// so that both known fields and additional properties are marshaled into a same JSON object.
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d AddtionalPropEscapeU003chmu003eObjectRaw) MarshalJSON() ([]byte, error) {
	buf := gentypehelper.GetBuf()
	defer gentypehelper.PutBuf(buf)
	var (
		bin []byte
		err error
	)
	buf.WriteByte('{')
	buf.WriteString("\"\\u0026mah\\u0026\":")
	bin, err = serde.Marshal(d.U0026mahu0026)
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
func (d *AddtionalPropEscapeU003chmu003eObjectRaw) UnmarshalJSON(data []byte) error {
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

type AddtionalPropEscapeUe29ca8Object struct {
	Yay                   string `json:"yay"`
	AdditionalProperties_ map[string]any
}

// ToRaw converts d into its plain equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AddtionalPropEscapeUe29ca8Object) ToRaw() AddtionalPropEscapeUe29ca8ObjectRaw {
	return AddtionalPropEscapeUe29ca8ObjectRaw{
		Yay:                   gentypehelper.MapSingleValueToElastic[string](d.Yay),
		AdditionalProperties_: d.AdditionalProperties_,
	}
}

// MarshalJSON implements json.Marshaler
// so that both known fields and additional properties are marshaled into a same JSON object.
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d AddtionalPropEscapeUe29ca8Object) MarshalJSON() ([]byte, error) {
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
func (d *AddtionalPropEscapeUe29ca8Object) UnmarshalJSON(data []byte) error {
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

type AddtionalPropEscapeUe29ca8ObjectRaw struct {
	Yay                   elastic.Elastic[string] `json:"yay"`
	AdditionalProperties_ map[string]any
}

// ToPlain converts d into its raw equivalent.
// It avoids copying data where it is possilbe. Mutation to fields is not advised.
func (d AddtionalPropEscapeUe29ca8ObjectRaw) ToPlain() AddtionalPropEscapeUe29ca8Object {
	return AddtionalPropEscapeUe29ca8Object{
		Yay:                   d.Yay.ValueSingle(),
		AdditionalProperties_: d.AdditionalProperties_,
	}
}

// MarshalJSON implements json.Marshaler
// so that both known fields and additional properties are marshaled into a same JSON object.
//
// The presence of this implementation indicates the dynamic field for this object are
// defined to be other than "strict" in its mapping.json.
func (d AddtionalPropEscapeUe29ca8ObjectRaw) MarshalJSON() ([]byte, error) {
	buf := gentypehelper.GetBuf()
	defer gentypehelper.PutBuf(buf)
	var (
		bin []byte
		err error
	)
	buf.WriteByte('{')
	buf.WriteString("\"yay\":")
	bin, err = serde.Marshal(d.Yay)
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
func (d *AddtionalPropEscapeUe29ca8ObjectRaw) UnmarshalJSON(data []byte) error {
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
