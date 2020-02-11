package main

import (
	"fmt"
	"io/ioutil"

	g "github.com/moznion/gowrtr/generator"
)

const pkgName = "yamlnillable"

func main() {
	typeName2StructName := map[string]string{
		"bool":    "Bool",
		"string":  "String",
		"int":     "Int",
		"int8":    "Int8",
		"int16":   "Int16",
		"int32":   "Int32",
		"int64":   "Int64",
		"uint":    "Uint",
		"uint8":   "Uint8",
		"uint16":  "Uint16",
		"uint32":  "Uint32",
		"uint64":  "Uint64",
		"byte":    "Byte",
		"rune":    "Rune",
		"float32": "Float32",
		"float64": "Float64",
	}
	valAttributeName := "Val"
	assignedAttributeName := "IsAssigned"

	for typeName, structName := range typeName2StructName {
		structureComment := g.NewCommentf(" %s is a data type for nillable %s value for YAML marshaling/unmarshaling.", structName, typeName)
		structure := g.NewStruct(structName).
			AddField(valAttributeName, typeName).
			AddField(assignedAttributeName, "bool")

		factoryMethodComment := g.NewCommentf(" %sOf makes a non-nil value with given %s.", structName, typeName)
		factoryMethodSignature := g.NewFuncSignature(fmt.Sprintf("%sOf", structName)).
			AddParameters(g.NewFuncParameter("val", typeName)).
			AddReturnTypes("*" + structName)
		factoryMethodReturnStmt := g.NewReturnStatement(fmt.Sprintf("&%s{%s: val, %s: true}", structName, valAttributeName, assignedAttributeName))
		factoryMethod := g.NewFunc(nil, factoryMethodSignature, factoryMethodReturnStmt)

		marshalingMethodComment := g.NewRoot(
			g.NewCommentf(" MarshalYAML marshals %s as YAML.", structName),
			g.NewComment(" This method used on marshaling YAML internally."),
		)
		marshalingMethodSignature := g.NewFuncSignature("MarshalYAML").
			ReturnTypes("interface{}", "error")
		marshalingMethodBody := g.NewRoot(
			g.NewIf("v."+assignedAttributeName, g.NewReturnStatement("v."+valAttributeName, "nil")),
			g.NewReturnStatement("nil", "nil"),
		)
		marshalingMethod := g.NewFunc(g.NewFuncReceiver("v", "*"+structName), marshalingMethodSignature, marshalingMethodBody)

		unmarshalingMethodComment := g.NewRoot(
			g.NewCommentf(" UnarshalYAML unmarshals %s as YAML.", structName),
			g.NewComment(" This method used on unmarshaling YAML internally."),
		)
		unmarshalingMethodSignature := g.NewFuncSignature("UnmarshalYAML").
			AddParameters(g.NewFuncParameter("unmarshal", "func(interface{}) error")).
			AddReturnTypes("error")
		unmarshalingMethodBody := g.NewRoot(
			g.NewRawStatementf("var val %s", typeName),
			g.NewIf("err := unmarshal(&val); err != nil", g.NewReturnStatement("err")),
			g.NewRawStatementf("v.%s = val", valAttributeName),
			g.NewRawStatementf("v.%s = true", assignedAttributeName),
			g.NewReturnStatement("nil"),
		)
		unmarshalingMethod := g.NewFunc(g.NewFuncReceiver("v", "*"+structName), unmarshalingMethodSignature, unmarshalingMethodBody)

		generated, err := g.NewRoot(
			g.NewComment(" Code generated by go-yaml-nillable generator; DO NOT EDIT."),
			g.NewNewline(),
			g.NewPackage(pkgName),
			g.NewNewline(),
			structureComment,
			structure,
			g.NewNewline(),
			factoryMethodComment,
			factoryMethod,
			g.NewNewline(),
			marshalingMethodComment,
			marshalingMethod,
			g.NewNewline(),
			unmarshalingMethodComment,
			unmarshalingMethod,
		).Gofmt("-s").Generate(0)

		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(typeName+".go", []byte(generated), 0644)
		if err != nil {
			panic(err)
		}
	}
}
