package test

var sampleDynamic = Dynamic{
	NestedInherit: DynamicNestedInheritNested{
		Age: -222,
	},
	NestedRuntime: DynamicNestedRuntimeNested{
		Age:                   259540,
		AdditionalProperties_: map[string]any{"baz": "qux"},
	},
	NestedStrict: DynamicNestedStrictNested{},
	ObjectFalse: DynamicObjectFalseObject{
		Age:                   555,
		AdditionalProperties_: map[string]any{"quux": "corge"},
	},
}
