package estime

import "strings"

func HasGoTimeToken(layout string) bool {
	for i := 0; i < len(layout); i++ {
		if tokens, ok := goTimeTokenSearchTable[layout[i]]; ok {
			for _, t := range tokens {
				if strings.HasPrefix(layout[i:], t) {
					return true
				}
			}
		}
	}
	return false
}

var goTimeTokenSearchTable = map[byte][]string{
	'0': {
		"002", "01",
		"02", "03",
		"04", "05",
		"06",
	},
	'1': {"15", "1"},
	'2': {"2006", "2"},
	'3': {"3"},
	'4': {"4"},
	'5': {"5"},
	'J': {"January", "Jan"},
	'M': {"Monday", "Mon", "MST"},
	'_': {"__2", "_2"},
	'P': {"PM"},
	'p': {"pm"},
	'Z': {"Z07:00:00", "Z07:00", "Z070000", "Z0700", "Z07"},
	'-': {"-07:00:00", "-07:00", "-070000", "-0700", "-07"},
	'.': {".0", ".9"},
}
