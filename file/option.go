package file

import (
	"fmt"
	"sort"
	"strings"
)

type (
	//File go文件的组成
	File struct {
		Pkg    string
		Class  ClassList
		Import []string
		Func   []Func
	}
	Func struct {
		Class     string
		Name      string
		Parameter string
		Ret       string
		Body      *[]string
	}
	ClassList []Class
	Class     struct {
		Type        string // c ctrl d:dao,e entity
		Name        string
		ParentClass []string
		Fields      []Field
	}
	Field struct {
		Name string
		Type string
		Tag  string
	}
)

//SetClass 设置struct
func (f *File) AddClass(name, Type string)*Class {
	f.Class = append(f.Class, Class{
		Type: Type,
		Name: name,
	})
	return &f.Class[len(f.Class)-1]
}

//NewFunc
func (f *File) NewFunc(class, name, parameter, ret, body string) {

	f.Func = append(f.Func, Func{
		Class:     class,
		Name:      name,
		Parameter: parameter,
		Ret:       ret,
		Body:      &([]string{body}),
	})

}

func (f *File) findFunc(name string) *Func {
	for _, fc := range f.Func {
		if fc.Name == name {
			return &fc
		}
	}
	return nil
}

//Func2String
func (f *File) Func2String() string {
	var funcStr strings.Builder
	for _, fu := range f.Func {

		var funcBody strings.Builder
		for _, s := range *fu.Body {
			funcBody.WriteString(s)
		}

		funcStr.WriteString("func ")
		if fu.Class != "" {
			for _, class := range f.Class {
				if fu.Class == class.Name {
					funcStr.WriteString(fmt.Sprintf("(%s %s)",class.Type, class.Name))
				}
			}

		}
		funcStr.WriteString(fmt.Sprintf("%s", fu.Name))
		if fu.Parameter != "" {
			funcStr.WriteString(fmt.Sprintf("(%s)", fu.Parameter))
		} else {
			funcStr.WriteString("()")
		}
		if fu.Ret != "" {
			funcStr.WriteString(fmt.Sprintf("(%s)", fu.Ret))
		}
		funcStr.WriteString(fmt.Sprintf("{\n%s\n}\n", funcBody.String()))

	}
	return funcStr.String()
}

//Pkg2String
func (f *File) Pkg2String() string {
	return "package " + f.Pkg + "\n"

}

//Import2String
func (f *File) Import2String() string {
	var importBody strings.Builder
	importBody.WriteString("import(\n")
	sort.Strings(f.Import)
	for _, im := range f.Import {

		importBody.WriteString(fmt.Sprintf("\t\"%s\"\n", im))

	}
	importBody.WriteString(")\n")
	return importBody.String()
}

func (c *Class) AddParent(name ...string) {
	c.ParentClass = append(c.ParentClass, name...)
}
func (c *Class) AddField(fields ...Field) {
	c.Fields = append(c.Fields, fields...)
}

func NewField(name, Type, Tag string) Field {
	return Field{
		Name: name,
		Type: Type,
		Tag:  Tag,
	}
}
func (c *Class) String() string {
	var classBody strings.Builder
	for _, class := range c.ParentClass {
		classBody.WriteString("\t")
		classBody.WriteString(class)
		classBody.WriteString("\n")
	}
	for _, field := range c.Fields {
		classBody.WriteString("\t")
		classBody.WriteString(fmt.Sprintf("\t%s %s `%s`", field.Name, field.Type, field.Tag))
		classBody.WriteString("\n")
	}
	return fmt.Sprintf("type %s struct{\n%s} \n", c.Name, classBody.String())
}

//AddBody func 添加内容
func (f *Func) AddBody(body ...string) {

	*(f.Body) = append(*(f.Body), body...)
}

func (list ClassList) String() string {
	var classBuf strings.Builder
	for _, class := range list {
		classBuf.WriteString(class.String())
	}
	return classBuf.String()
}
