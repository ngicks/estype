package estime

import (
	"fmt"
	"strings"
)

type FormatError struct {
	idx      int
	expected string
	actual   string
	msg      string
}

func (e *FormatError) Error() string {
	return fmt.Sprintf("index [%d]: %s but %s. %s", e.idx, e.expected, e.actual, e.msg)
}

// ConvertTimeToken converts Java Time tokens included in input into Go std time counterparts.
// input is still allowed to have Go std time token. Conversion for those tokens are no-op.
//
// This functions does not process the optional part (enclosed by a `[` and a `]`) of time tokens.
// The input must be treated by optionalstring package if you wish to use the optional format.
//
// Allowed input tokens are what they are for https://docs.oracle.com/javase/8/docs/api/java/time/format/DateTimeFormatter.html.
// But this package omits support for tokens described below as the std time package does not support it:
//
//   - G(era, AD; Anno Domini; A)
//   - Q/q(quarter-of-year, 3; 03; Q3; 3rd quarter)
//   - Y(week-based-year, 1996; 96)
//   - w(week-of-week-based-year, 27)
//   - W(week-of-month, 4)
//   - e/c(localized day-of-week, 2; 02; Tue; Tuesday; T)
//   - F(week-of-month, 3)
//   - K(hour-of-am-pm (0-11), 0)
//   - k(clock-hour-of-am-pm (1-24), 0)
//   - A(milli-of-day, 1234)
//   - n(nano-of-second, 987654321)
//   - N(nano-of-day, 1234000000)
//   - V(time-zone ID, America/Los_Angeles; Z; -08:30)
//   - O(localized zone-offset, GMT+8; GMT+08:00; UTC-08:00;)
//
// `S` is omitted since Go supports fraction-of-time only when there is a preceding dot character.
// You can still use `S` but only with a leading dot.
//
// Its detailed conversion rule is defined as a following conversion table:
//
//		var tokenTable = map[timeFormatToken]goTimeFmtToken{
//			`uuuu`: `2006`,
//			`uu`:   `06`,
//			`yyyy`: `2006`,
//			`yy`:   `06`,
//			`DDD`:  `002`,
//			`M`:    `1`,
//			`MM`:   `01`,
//			`L`:    `1`,
//			`LL`:   `01`,
//			`d`:    `2`,
//			`dd`:   `02`,
//			`EEEE`: `Monday`,
//			`EEE`:  `Mon`,
//			`a`:    `PM`,
//			`hh`:   `03`,
//			`h`:    `3`,
//			`HH`:   `15`,
//			`mm`:   `04`,
//			`m`:    `4`,
//			`ss`:   `05`,
//			`s`:    `5`,
//			`z`:    `MST`,
//			`XXX`:  `Z070000`,
//			`XX`:   `Z0700`,
//			`X`:    `Z07`,
//			`xxx`:  `-070000`,
//			`xx`:   `-0700`,
//			`x`:    `-07`,
//			`Z`:    `Z0700`,
//			// further more
//			`MMMM`: `January`,
//			`MMM`:  `Jan`,
//			// preservation for Go tokens.
//			`MST`:       `MST`,
//			`Z070000`:   `Z070000`,
//			`Z07`:       `Z07`,
//			`Z07:00:00`: `Z07:00:00`,
//			`-0700`:     `-0700`,
//			`-070000`:   `-070000`,
//			`-07`:       `-07`,
//			`-07:00`:    `-07:00`,
//			`-07:00:00`: `-07:00:00`,
//	 }
func ConvertTimeToken(input string) (string, error) {
	var prefix, token string
	var isToken bool
	var err error

	var output string

	for len(input) > 0 {
		prefix, token, input, isToken, err = nextChunk(input)
		if err != nil {
			return "", err
		}
		output += prefix
		if isToken {
			output += timeFormatToken(token).toGoFmt()
		} else {
			output += token
		}
	}

	return output, nil
}

// nextChunk reads input string from its head, up to a first time token or espaced string.
//
// prefix is non time token string which is read up before the first hit.
// found is next chunk string. If isTokein is true, chunk is a time token, an unescaped string otherwise.
// suffix is rest of input.
// err would be non nil if token has wrong length.
func nextChunk(input string) (prefix string, found string, suffix string, isToken bool, err error) {
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '\\':
			return input[:i], input[i+1 : i+2], input[i+2:], false, nil
		case '.':
			if strings.HasPrefix(input[i:], ".S") ||
				strings.HasPrefix(input[i:], ".9") ||
				strings.HasPrefix(input[i:], ".0") {
				repeated := getRepeatOf(input[i+1:], input[i+1:i+2])
				if repeated[0] == 'S' {
					repeated = strings.Replace(repeated, "S", "0", -1)
				}
				return input[:i], "." + repeated, input[i+len("."+repeated):], true, nil
			}
		case '\'':
			unescaped := getUntilClosingSingleQuote(input[i+1:])
			returned := unescaped
			if unescaped == "" {
				returned = "'"
			}
			return input[:i], returned, input[i+len(`'`+unescaped+`'`):], false, nil
		}

		possibleSequences, ok := tokenSearchTable[input[i]]
		if ok {
			for _, possible := range possibleSequences {
				if strings.HasPrefix(string(input[i:]), string(possible)) {
					return input[:i], string(possible), input[i+len(possible):], true, nil
				}
			}
			if input[i] == '-' {
				continue
			}
			return "", "", "", false, &FormatError{
				idx:      i,
				expected: fmt.Sprintf("must be prefixed with one of %+v", possibleSequences),
				actual:   input[i:],
				msg:      "maybe wrong len, like Y or YYY.",
			}
		}
	}
	return input, "", "", false, nil
}

func getRepeatOf(input string, target string) string {
	for i := 0; i < len(input); i++ {
		if input[i:i+len(target)] != target {
			return input[:i+len(target)-1]
		}
	}
	return input
}

// getUntilClosingSingleQuote returns `aaaaa` if input is `aaaaa'`.
func getUntilClosingSingleQuote(input string) string {
	for i := 0; i < len(input); i++ {
		if input[i] == '\'' {
			if i == 0 {
				return ""
			}
			if input[i-1] != '\\' || strings.HasSuffix(input[:i+1], `\\'`) {
				return input[:i]
			}
		}
	}
	return input
}

// tokenTable maps java-like time token to Go std time token.
var tokenTable = map[timeFormatToken]goTimeFmtToken{
	`uuuu`: `2006`,
	`uu`:   `06`,
	`yyyy`: `2006`,
	`yy`:   `06`,
	`DDD`:  `002`,
	`M`:    `1`,
	`MM`:   `01`,
	`L`:    `1`,
	`LL`:   `01`,
	`d`:    `2`,
	`dd`:   `02`,
	`EEEE`: `Monday`,
	`EEE`:  `Mon`,
	`a`:    `PM`,
	`hh`:   `03`,
	`h`:    `3`,
	`HH`:   `15`,
	`mm`:   `04`,
	`m`:    `4`,
	`ss`:   `05`,
	`s`:    `5`,
	`z`:    `MST`,
	`XXX`:  `Z070000`,
	`XX`:   `Z0700`,
	`X`:    `Z07`,
	`xxx`:  `-070000`,
	`xx`:   `-0700`,
	`x`:    `-07`,
	`Z`:    `Z0700`,
	// further more
	`MMMM`: `January`,
	`MMM`:  `Jan`,
	// preservation for Go tokens.
	`MST`:       `MST`,
	`Z070000`:   `Z070000`,
	`Z07`:       `Z07`,
	`Z07:00:00`: `Z07:00:00`,
	`-0700`:     `-0700`,
	`-070000`:   `-070000`,
	`-07`:       `-07`,
	`-07:00`:    `-07:00`,
	`-07:00:00`: `-07:00:00`,
}

// tokenSearchTable maps a leading character to possible sequence.
var tokenSearchTable = map[byte][]timeFormatToken{
	'u': {"uuuu", "uu"},
	'y': {"yyyy", "yy"},
	'D': {"DDD"},
	'M': {"MMMM", "MMM", "MST", "MM", "M"},
	'L': {"LL", "L"},
	'd': {"dd", "d"},
	'E': {"EEEE", "EEE"},
	'a': {"a"},
	'h': {"hh", "h"},
	'H': {"HH"},
	'm': {"mm", "m"},
	's': {"ss", "s"},
	'z': {"z"},
	'X': {"XXX", "XX", "X"},
	'x': {"xxx", "xx", "x"},
	'Z': {"Z07:00:00", "Z070000", "Z07", "Z"},
	// '-' with no succeeding 0 is a valid non-token.
	// '-':  {"-0700", "-070000", "-07", "-07:00", "-07:00:00"},
	// '.' with succeeding 0,9,S needs special handling.
	// single '.' is non-token.
}

type timeFormatToken string

type goTimeFmtToken string

func (tt timeFormatToken) toGoFmt() string {
	token, ok := tokenTable[tt]
	if ok {
		return string(token)
	}

	if strings.HasPrefix(string(tt), ".S") {
		return strings.ReplaceAll(string(tt), "S", "0")
	} else if strings.HasPrefix(string(tt), ".0") || strings.HasPrefix(string(tt), ".9") {
		return string(tt)
	}
	panic(fmt.Sprintf("unknown: %s", tt))
}
