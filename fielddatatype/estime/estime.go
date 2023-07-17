package estime

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ngicks/estype/fielddatatype/estime/optionalstring"
)

// NumberParser is a parser which decodes JSON number to time.Time.
// It only can be "epoch_millis" or "epoch_second".
type NumberParser string

func (p NumberParser) Parse(val int64) (time.Time, error) {
	switch p {
	case Millis:
		return time.UnixMilli(val), nil
	case Second:
		return time.Unix(val, 0), nil
	default:
		return time.Time{}, fmt.Errorf("estime: parser is not allowed to process int64. input is %d", val)
	}
}

func (p NumberParser) Format(t time.Time) int64 {
	switch p {
	case Second:
		return t.Unix()
	default:
		return t.UnixMilli()
	}
}

const (
	None   NumberParser = ""
	Millis NumberParser = "epoch_millis"
	Second NumberParser = "epoch_second"
)

func NumParser(numberFmt string) NumberParser {
	switch numberFmt {
	case string(Millis):
		return Millis
	case string(Second):
		return Second
	}
	return None
}

type EsTimeParser struct {
	numParser   NumberParser
	multiLayout MultiLayout
}

func FromJavaDateTimeLike(layouts []string, numParser NumberParser) (EsTimeParser, error) {
	goLayouts := make([]string, 0)
	for _, format := range layouts {
		opts, err := optionalstring.EnumerateOptionalString(format)
		if err != nil {
			return EsTimeParser{}, err
		}
		for _, layout := range opts {
			goLayout, err := ConvertTimeToken(layout)
			if err != nil {
				return EsTimeParser{}, err
			}
			goLayouts = append(goLayouts, goLayout)
		}
	}

	return FromGoTimeLayout(goLayouts, numParser)
}

func FromGoTimeLayout(layouts []string, numParser NumberParser) (EsTimeParser, error) {
	multiLayout, err := NewMultiLayout(layouts)
	if err != nil {
		return EsTimeParser{}, err
	}

	switch numParser {
	case Millis, Second:
	default:
		numParser = ""
	}

	return EsTimeParser{
		numParser:   numParser,
		multiLayout: multiLayout,
	}, nil
}

func FromGoTimeLayoutUnsafe(layouts []string, numParser NumberParser) EsTimeParser {
	return EsTimeParser{
		numParser:   numParser,
		multiLayout: NewMultiLayoutUnsafe(layouts),
	}
}

func (t EsTimeParser) ParseJson(data []byte) (time.Time, error) {
	if data[0] == '"' && data[len(data)-1] == '"' {
		return t.ParseString(string(data[1 : len(data)-1]))
	}
	if data[0] == '-' || ('0' <= data[0] && data[0] <= '9') {
		num, err := strconv.ParseInt(string(data), 10, 64)
		if err != nil {
			return time.Time{}, err
		}
		return t.ParseNumber(num)
	}
	return time.Time{}, fmt.Errorf(
		"estime: ParseJson received an unknown type."+
			" The input must be JSON string or number without decimal point. input is %s",
		data,
	)
}

func (t EsTimeParser) ParseString(s string) (time.Time, error) {
	if len(t.multiLayout.layouts) == 0 {
		return time.Time{}, fmt.Errorf("estime: parser is not allowed to process string input. input = %s", s)
	}
	return t.multiLayout.Parse(s)
}

func (t EsTimeParser) ParseNumber(n int64) (time.Time, error) {
	return t.numParser.Parse(n)
}

func (p EsTimeParser) FormatString(t time.Time, idx uint) string {
	if len(p.multiLayout.layouts) <= int(idx) {
		idx = uint(len(p.multiLayout.layouts) - 1)
	}
	return t.Format(p.multiLayout.layouts[idx])
}

func (p EsTimeParser) FormatNumber(t time.Time) int64 {
	return p.numParser.Format(t)
}

func (t EsTimeParser) Layout() []string {
	return t.multiLayout.Clone()
}
