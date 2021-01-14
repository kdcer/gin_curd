package file

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"strings"
)

const (
	daoDir = "handler/dao/"
	tag    = "curd"
)

func init() {

	os.MkdirAll(daoDir, 0777)

}

//BuildDao 生成Dao
func BuildDao(data []interface{}, db string) {
	entityPkg:=""
	for _, da := range data {
		doFile := new(File)
		doFile.Pkg = "dao"
		of := reflect.TypeOf(da)
		modelPkg := of.Elem().PkgPath()
		doFile.Import = append(doFile.Import, modelPkg)

		name := of.Elem().Name()
		lowerName := strings.ToLower(string(name[0])) + name[1:]
		file, err := createFile(daoDir + lowerName)
		if err != nil {
			panic(err)
		}
		 doFile.AddClass(name+"Dao", "d")

		doFile.Import = append(doFile.Import, db)
		search,field, isSearch := entityField(of)
		if isSearch {//
			if entityPkg=="" {
				entityPkg,_=path.Split(modelPkg)
				entityPkg=entityPkg+"entity"
			}
			doFile.Import = append(doFile.Import, entityPkg)
			InitEntity(lowerName,Class{
				Type:        "e",
				Name:        name+"ListReq",
				ParentClass: nil,
				Fields:      field,
			})
			doFile.NewFunc(name+"Dao","List", "req entity."+name,fmt.Sprintf("list []models.%s,err error", name), listDaoStr(search))
		}else{
			doFile.NewFunc(name+"Dao","List", "",fmt.Sprintf("list []models.%s,err error", name), listDaoStr(search))
		}
		doFile.NewFunc(name+"Dao","Create", "data models."+name, "err error", createDaoStr(of))
		doFile.NewFunc(name+"Dao","Update", "id int,data models."+name,  "err error", upDaoStr(of))
		doFile.NewFunc(name+"Dao","Delete", "id int", "err error",  delDaoStr(of))
		doFile.NewFunc(name+"Dao","Info", "id int", fmt.Sprintf("data models.%s,err error", name), infoDaoStr(of))
		Build(file, doFile)
	}
}
func listDaoStr(search strings.Builder) string {
	var bodyStr strings.Builder
	bodyStr.WriteString("\tDb:=db.Db\n")

	bodyStr.WriteString(search.String())
	bodyStr.WriteString("\tDb.Find(&list) \n")
	bodyStr.WriteString("\treturn  \n")

	return bodyStr.String()
}

func buildSearchStr(field string, like bool) string {
	if like {
		return fmt.Sprintf("\t\tDb=Db.Where(\"%s like ?\",%%%s%%)\n", snakeString(field), "req."+field)
	} else {
		return fmt.Sprintf("\t\tDb=Db.Where(\"%s = ?\",%s)\n", snakeString(field), "req."+field)
	}
}

func createDaoStr(p reflect.Type) string {
	bodyStr := `
	err=db.Db.Create(&data).Error 
	return `
	return bodyStr
}

func upDaoStr(p reflect.Type) string {
	bodyStr := `
	err=db.Db.Where("id=?",id).Update(&data).Error
	return `
	return bodyStr
}

func delDaoStr(p reflect.Type) string {
	bodyStr := `
	err=db.Db.Where("id=?",id).Delete(%s{}).Error
	return `
	return fmt.Sprintf(bodyStr, p.Elem())
}

func infoDaoStr(p reflect.Type) string {
	bodyStr := `
	err=db.Db.Where("id=?",id).First(&data).Error
	return `
	return bodyStr
}
