package sql2struct

import (
	"fmt"
	"os"
	"text/template"
)

const structTpl = `type {{.TableName}} struct{
	{{range .Columns}}  {{$length := len .Comment}} {{if gt $length 0}} // {{.Comment}} {{else}} // {{.Name}}  {{end}} 
	{{$Typelen := len .Type}} {{if gt $Typelen 0}} {{.Name}} {{.Type}}  {{.Tag}} {{else}} {{.Name}}{{end}} 
	{{end}}
}
func (model {{.TableName}})TableName() string{
	return "{{.TableName}}"
}`

type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}
type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, v := range tbColumns {
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", v.ColumnName)
		tplColumns = append(tplColumns, &StructColumn{
			Name:    v.ColumnName,
			Type:    DBTypeToStructType[v.ColumnType],
			Tag:     tag,
			Comment: v.ColumnComment,
		})
	}
	return tplColumns
}

func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Parse(t.structTpl))
	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	//把结果输出到stdout中；
	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}
	return nil
}
