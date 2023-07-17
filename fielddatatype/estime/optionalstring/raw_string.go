package optionalstring

type valueType int

const (
	Normal valueType = iota
	SingleQuoteEscaped
	SlashEscaped
)

type textNode struct {
	typ   valueType
	value string
}

func (v textNode) Typ() valueType {
	return v.typ
}

func (v textNode) Len() int {
	return len(v.value)
}

func (v textNode) Value() string {
	return v.value
}

func (v textNode) Unescaped() string {
	switch v.typ {
	case Normal:
		return v.value
	case SingleQuoteEscaped:
		if val := v.Value(); val == "''" {
			return "'"
		} else {
			return val[1 : v.Len()-1]
		}
	case SlashEscaped:
		return v.Value()[1:]
	}
	panic("unknown")
}

type RawString []textNode

func newRawString() RawString {
	return make(RawString, 0)
}

func (rs RawString) append(str ...RawString) RawString {
	c := rs.clone()
	for _, v := range str {
		c = append(c, v...)
	}
	return c
}

func (rs RawString) clone() RawString {
	cloned := make(RawString, len(rs))
	copy(cloned, rs)
	return cloned
}

// String returns rs's raw internal value.
func (rs RawString) String() string {
	var out string
	for _, v := range rs {
		out += v.value
	}
	return out
}

// Unescaped returns rs's internal value unescaped.
func (rs RawString) Unescaped() string {
	var out string
	for _, v := range rs {
		out += v.Unescaped()
	}
	return out
}
