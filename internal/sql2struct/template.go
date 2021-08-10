package sql2struct

const structTPL = `type {{.TableName}} struct{
	{{range .Columns}}  {{$length := len .Comment}} {{if gt $length 0}} // {{.Comment}} {{else}} // {{.Name}}  {{end}} 
	{{$Typelen := len .Type}} {{if gt $Typelen 0}} {{.Name}} {{.Type}}  {{.Tag}} {{else}} {{.Name}}{{end}} 
	{{end}}
}
func (model {{.TableName}})TableName() string{
	return "{{.TableName}}"
}`


type StructTemplate struct{
	structTpl string
}

type StructColumn struct{
	Name string
	Type string
	Tag string
	Comment string
}
type StructTemplateDB struct{
	TableName string
	Columns []*StructColumn
}