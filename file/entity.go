package file

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

const (
	entityDir = "handler/entity/"
)

func init() {
	router = new(File)
	_ = os.Mkdir(entityDir, 0777)
}

func InitEntity(fileName string, class Class) {
	file, err := createFile(entityDir + fileName)
	if err != nil {
		panic(err)
	}
	file.WriteString("package entity \n")
	file.WriteString(class.String())

}

// 解析tag
func entityField(p reflect.Type) (bodyStr strings.Builder, fields []Field, search bool) {
	elem := p.Elem()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		ctag := field.Tag.Get(tag)
		bindTag := field.Tag.Get("binding")
		setting := parseTagSetting(ctag)
		var tag strings.Builder

		if bindTag!="" {
			tag.WriteString( fmt.Sprintf("binding:\"%s\"", bindTag))
		}

		if val, ok := setting["search"]; ok { // 搜索
			search = true
			if tag.Len()!=0 {
				tag.WriteString( ";")
			}
			tag.WriteString( fmt.Sprintf("from:\"%s\"", snakeString(field.Name)))


			bodyStr.WriteString(fmt.Sprintf("\tif req.%s !=\"\"{\n", field.Name))
			searchStr := ""
			if val == "like" {
				searchStr = buildSearchStr(field.Name, true)
			} else {
				searchStr = buildSearchStr(field.Name, false)
			}
			bodyStr.WriteString(searchStr)
			bodyStr.WriteString("\t}\n")
		}

		if tag.Len()!=0 {
			fields = append(fields, NewField(field.Name, "string",tag.String()))
		}

	}
	return
}
