package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"os"
	"strings"
	"text/template"
	"unicode"
)

//go:embed blocks.json
var jsonBytes []byte

type jsonBlock struct {
	Properties map[string][]string `json:"properties"`

	States []struct {
		Id         int  `json:"id"`
		Default    bool `json:"default"`
		Properties map[string]string
	}
}

func main() {
	var jsonBlocks map[string]jsonBlock

	if err := json.Unmarshal(jsonBytes, &jsonBlocks); err != nil {
		panic(err)
	}

	biggst := 0
	p := map[string][]string{}
	for _, block := range jsonBlocks {
		var n int
		for _, v := range block.Properties {
			n = len(v)
		}

		if n > biggst {
			biggst = n
			p = block.Properties
		}
	}
	fmt.Println("biggest properties", p)
	codeTemplate := `
package block
//THIS FILE IS AUTO GENERATED PLEASE DO NOT TOUCH

{{ range $k, $v := . }}
{{- $GoName := stringToFieldName $k false }}
type {{ $GoName }} struct {

{{- if $v.Properties -}}

{{- range $prop, $value := $v.Properties }}
{{ stringToFieldName $prop true }} {{ findType $value }}
{{- end -}}

{{- end -}}
}

{{- range $i, $state := $v.States }}

{{- if and $state.Default $state.Properties }}
{{- $receiver := slice $GoName 0 1 }}
func ({{ $receiver }} *{{ $GoName }}) New(prop map[string]interface{}){

{{- range $prop, $value := $state.Properties -}}
{{- $GoType := findType2 $value -}}
{{- $skip := matchesGoDefault $GoType $value -}}
{{- if not $skip }}
{{ $receiver }}.{{ stringToFieldName $prop true }}=
{{- if eq $GoType "string" -}}
"{{ $value }}"
{{- else -}}
{{ $value }}

{{- end -}}
{{- end -}}
{{- end -}}
}
{{ end -}}
{{- end -}}

{{- end }}
`

	tmpl := template.New("blocks").Funcs(template.FuncMap{
		"stringToFieldName": stringToFieldName,
		"findType":          findType,
		"findType2":         findType2,
		"matchesGoDefault":  matchesGoDefault,
	})

	tmpl, err := tmpl.Parse(codeTemplate)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	wr := io.MultiWriter(&buf)
	if err = tmpl.Execute(wr, jsonBlocks); err != nil {
		panic(err)
	}

	fmtData, _ := format.Source(buf.Bytes())

	if err = os.WriteFile("blocks.go", fmtData, 0666); err != nil {
		panic(err)
	}
}

func matchesGoDefault(typ, value string) bool {
	if typ == "bool" && value == "false" {
		return true
	}

	if typ == "int" && value == "0" {
		return true
	}

	return false
}

func stringToFieldName(s string, private bool) string {
	if s == "type" {
		if private {
			return "_type"
		}
		return "Type"
	}

	words := strings.FieldsFunc(s, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})

	for i, word := range words {
		if word == "minecraft" {
			words[i] = ""
			continue
		}
		words[i] = strings.Title(word)
	}

	r := []byte(strings.Join(words, ""))
	if private {
		r[0] = strings.ToLower(string(r[:1]))[0]
	}

	return string(r)
}

func findType2(s string) string {
	return findType([]string{s})
}

func findType(s []string) string {
	for _, v := range s {

		if isInt(v) {
			return "int"
		}

		if v == "true" || v == "false" {
			return "bool"
		}
	}

	return "string"
}

func isInt(s string) bool {
	for _, s2 := range s {

		if !unicode.IsDigit(s2) {
			return false
		}
	}

	return true
}
