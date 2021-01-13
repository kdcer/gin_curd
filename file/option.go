package file

import (
	"fmt"
	"sort"
	"strings"
)

const (
	Tag = "curl"
)

type (
	File struct {
		Pkg    string
		Class  *Class
		Import []string
		Func   []Func
	}
	Func struct {
		Name      string
		Parameter string
		Ret       string
		Body      *[]string
	}
	Class struct {
		Type        string
		Name        string
		ParentClass []string
	}
)

func (f *File) SetClass(name, Type string) *Class {
	f.Class = &Class{
		Type: Type,
		Name: name,
	}

	return f.Class

}

func (f *File) NewFunc(name, parameter, ret, body string) {

	f.Func = append(f.Func, Func{
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

func (f *File) Func2String() string {
	var funcStr strings.Builder
	for _, fu := range f.Func {

		var funcBody strings.Builder
		for _, s := range *fu.Body {
			funcBody.WriteString(s)
		}

		funcStr.WriteString("func ")
		if f.Class!=nil{
			funcStr.WriteString(fmt.Sprintf("(%s %s)",f.Class.Type,f.Class.Name))
		}
		funcStr.WriteString(fmt.Sprintf("%s",fu.Name))
		if fu.Parameter!="" {
			funcStr.WriteString(fmt.Sprintf("(%s)",fu.Parameter))
		}else{
			funcStr.WriteString("()")
		}
		if fu.Ret!="" {
			funcStr.WriteString(fmt.Sprintf("(%s)",fu.Ret))
		}
		funcStr.WriteString(fmt.Sprintf("{\n%s\n}\n",funcBody.String()))

	}
	return funcStr.String()
}
func (f *File) Pkg2String() string {
	return "package " + f.Pkg + "\n"

}
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
func (c *Class) String() string {
	var classBody strings.Builder
	for _, class := range c.ParentClass {
		classBody.WriteString("\t")
		classBody.WriteString(class)
		classBody.WriteString("\n")
	}
	return fmt.Sprintf("type %s struct{\n%s} \n", c.Name, classBody.String())
}
func (f *Func) AddBody(body ...string) {

	*(f.Body) = append(*(f.Body), body...)
}
