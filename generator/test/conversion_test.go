package test

var sampleConversion = Conversion{
	MultipleOptional: escape([]ConversionMultipleOptionalNested{
		{
			Age: 123,
			Name: ConversionMultipleOptionalNameObject{
				First: "john",
				Last:  "doe",
			},
		},
		{
			Age: 224,
			Name: ConversionMultipleOptionalNameObject{
				First: "jane",
				Last:  "doe",
			},
		},
	}),
	MultipleRequired: []ConversionMultipleRequiredNested{
		{
			Age: 11,
			Name: ConversionMultipleRequiredNameObject{
				First: "foo",
				Last:  "bar",
			},
		},
		{
			Age: 55,
			Name: ConversionMultipleRequiredNameObject{
				First: "baz",
				Last:  "qux",
			},
		},
	},
	SingleOptional: &ConversionSingleOptionalObject{
		Age: 81,
		Name: ConversionSingleOptionalNameObject{
			First: "nyan",
			Last:  "cat",
		},
	},
	SingleRequired: ConversionSingleRequiredObject{
		Age: 77,
		Name: ConversionSingleRequiredNameObject{
			First: "shibe",
			Last:  "dog",
		},
	},
}
