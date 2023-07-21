package generator

import (
	"github.com/dave/jennifer/jen"
)

type TypeIdRenderOption interface {
	IsOptional() bool
	IsSingle() bool
}

type typeIdRenderOption struct {
	opt, single bool
}

func RenderOption(isOptional, isSingle bool) typeIdRenderOption {
	return typeIdRenderOption{
		opt:    isOptional,
		single: isSingle,
	}
}

func (o typeIdRenderOption) IsOptional() bool {
	return o.opt
}

func (o typeIdRenderOption) IsSingle() bool {
	return o.single
}

type TypeId struct {
	Qualifier    string
	TypeParam    []TypeId
	Id           string
	NonWritable  bool // A field data type for which the Elasticsearch does not allow store value.
	AlwaysSingle bool // A field data type for which the Elasticsearch only accepts T or T[] with single element.
}

func (t TypeId) Render(option TypeIdRenderOption) *jen.Statement {
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
				g.Add(p.Render(RenderOption(false, true)))
			}
		})
	}

	return stmt
}

func (t TypeId) IsSingle(option TypeIdRenderOption) bool {
	return t.AlwaysSingle || option.IsSingle()
}

func (t TypeId) IsOptional(option TypeIdRenderOption) bool {
	return option.IsOptional()
}
