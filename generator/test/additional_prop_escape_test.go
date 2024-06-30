package test

var sampleAdditionalPropEscape = AdditionalPropEscape{
	U003chmu003e: AdditionalPropEscapeU003chmu003eObject{
		U0026mahu0026:         "<not now>",
		AdditionalProperties_: map[string]any{"foo": "bar"},
	},
	U2728: AdditionalPropEscapeU2728Object{
		Yay:                   "yay",
		AdditionalProperties_: map[string]any{"baz": "quux"},
	},
	AdditionalProperties_: map[string]any{"qux": "corge"},
}
