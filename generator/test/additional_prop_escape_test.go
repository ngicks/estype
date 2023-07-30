package test

var sampleAddtionalPropEscape = AddtionalPropEscape{
	U003chmu003e: AddtionalPropEscapeU003chmu003eObject{
		U0026mahu0026:         "<not now>",
		AdditionalProperties_: map[string]any{"foo": "bar"},
	},
	U2728: AddtionalPropEscapeU2728Object{
		Yay:                   "yay",
		AdditionalProperties_: map[string]any{"baz": "quux"},
	},
	AdditionalProperties_: map[string]any{"qux": "corge"},
}
