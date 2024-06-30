package generator

import (
	"github.com/dave/jennifer/jen"
)

type typeIdRenderOption interface {
	IsOptional() bool
	IsSingle() bool
}

type simpleRenderOption struct {
	opt, single bool
}

func newSimpleRenderOption(isOptional, isSingle bool) simpleRenderOption {
	return simpleRenderOption{
		opt:    isOptional,
		single: isSingle,
	}
}

func (o simpleRenderOption) IsOptional() bool {
	return o.opt
}

func (o simpleRenderOption) IsSingle() bool {
	return o.single
}

type typeId struct {
	Qualifier     string
	TypeParam     []typeId
	Id            string
	NonWritable   bool // A field data type for which the Elasticsearch does not allow store value.
	AlwaysSingle  bool // A field data type for which the Elasticsearch only accepts T or T[] with single element.
	DisallowArray bool // Some types can not even be an array of a single element. see ./field.go for details.
	DisallowNull  bool // Some types can not be `null` which means it is not possible to overwrite those fields with a null value.
}

func (t typeId) Render(option typeIdRenderOption) *jen.Statement {
	stmt := new(jen.Statement)

	if t.NonWritable {
		return stmt.Op("*").Id("struct{}")
	}
	if option.IsOptional() {
		stmt = stmt.Op("*")
	}
	if !t.AlwaysSingle && !option.IsSingle() {
		stmt = stmt.Index()
	}

	if t.Qualifier != "" {
		stmt = stmt.Qual(t.Qualifier, t.Id)
	} else {
		stmt = stmt.Id(t.Id)
	}

	if len(t.TypeParam) > 0 {
		stmt = stmt.IndexFunc(func(g *jen.Group) {
			for _, p := range t.TypeParam {
				g.Add(p.Render(newSimpleRenderOption(false, true)))
			}
		})
	}

	return stmt
}

func (t typeId) IsSingle(option typeIdRenderOption) bool {
	return t.AlwaysSingle || option.IsSingle()
}

func (t typeId) IsOptional(option typeIdRenderOption) bool {
	return option.IsOptional()
}

func (t typeId) MustOmit(option typeIdRenderOption) bool {
	return isUnd(t) || t.NonWritable || (t.DisallowNull && option.IsOptional())
}

// IgnoreInConversion reports whether t must not be converted between plain and raw types.
func (t typeId) IgnoreInConversion() bool {
	return t.NonWritable
}
