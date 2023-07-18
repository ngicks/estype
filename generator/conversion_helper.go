package generator

const (
	escapeValueId     = "escapeValue"
	escapeSliceId     = "escapeSlice"
	mapToPlainId      = "mapToPlain"
	mapToRawId        = "mapToRaw"
	mapToRawPointerId = "mapToRawPointer"
)

func generateEscapeValue(ctx *GeneratorContext) {
	if ctx.globalState.escapeValueGenerated {
		return
	}
	ctx.globalState.escapeValueGenerated = true
	ctx.file.
		Line().
		Line().
		Id(
			"func escapeValue[T any](v T) *T {\n" +
				"	return &v\n" +
				"}",
		).
		Line().
		Line()
}

func generateEscapeSlice(ctx *GeneratorContext) {
	if ctx.globalState.escapeSliceGenerated {
		return
	}
	ctx.globalState.escapeSliceGenerated = true

	ctx.file.
		Line().
		Line().
		Id(
			"func escapeSlice[T any](sl []T) *[]T {\n" +
				"	if sl == nil {\n" +
				"		return nil\n" +
				"	}\n" +
				"	return &sl\n" +
				"}",
		).
		Line().
		Line()
}

func generateMapToPlain(ctx *GeneratorContext) {
	if ctx.globalState.mapToPlainGenerated {
		return
	}
	ctx.globalState.mapToPlainGenerated = true

	ctx.
		file.
		Line().
		Line().
		Id(
			"func mapToPlain[T any, U interface{ ToPlain() T }](l []U) []T {\n" +
				"	if l == nil {\n" +
				"		return nil\n" +
				"	}\n" +
				"	out := make([]T, len(l))\n" +
				"	for i, v := range l {\n" +
				"		out[i] = v.ToPlain()\n" +
				"	}\n" +
				"	return out\n" +
				"}\n",
		).
		Line().
		Line()
}
func generateMapToRaw(ctx *GeneratorContext) {
	if ctx.globalState.mapToRawGenerated {
		return
	}
	ctx.globalState.mapToRawGenerated = true

	ctx.
		file.
		Line().
		Line().
		Id(
			"func mapToRaw[T any, U interface{ ToRaw() T }](l []U) []T {\n" +
				"	if l == nil {\n" +
				"		return nil\n" +
				"	}\n" +
				"	out := make([]T, len(l))\n" +
				"	for i, v := range l {\n" +
				"		out[i] = v.ToRaw()\n" +
				"	}\n" +
				"	return out\n" +
				"}\n",
		).
		Line().
		Line()
}

func generateMapToRawPointer(ctx *GeneratorContext) {
	if ctx.globalState.mapToRawPointerGenerated {
		return
	}
	ctx.globalState.mapToRawPointerGenerated = true

	ctx.file.
		Line().
		Line().
		Id(
			"func mapToRawPointer[T any, U interface{ ToRaw() T }](l *[]U) []T {\n" +
				"	if l == nil {\n" +
				"		return nil\n" +
				"	}\n" +
				"	out := make([]T, len(*l))\n" +
				"	for i, v := range *l {\n" +
				"		out[i] = v.ToRaw()\n" +
				"	}\n" +
				"	return out\n" +
				"}\n",
		).
		Line().
		Line()
}
