package file

import (
	"fmt"
	"os"
	"reflect"
)

const daoDir = "handler/dao/"

func init() {
	os.MkdirAll(daoDir, 0777)

}

//BuildDao 生成Dao
func BuildDao(data []interface{},db string) {

	for _, da := range data {
		doFile := new(File)
		doFile.Pkg = "dao"
		of := reflect.TypeOf(da)
		modelPkg:=of.Elem().PkgPath()
		doFile.Import = append(doFile.Import, modelPkg)

		name := of.Elem().Name()
		file, err := createFile(daoDir+name)
		if err != nil {
			panic(err)
		}
		doFile.SetClass(name+"Dao", "d")

		doFile.Import = append(doFile.Import,db )
		doFile.NewFunc("Create", "data models."+name, "err error", createDaoStr(of))
		doFile.NewFunc("Update", "id int,data models."+name, "err error", upDaoStr(of))
		doFile.NewFunc("Delete", "id int", "err error", delDaoStr(of))
		doFile.NewFunc("Info", "id int", fmt.Sprintf("data models.%s,err error", name), infoDaoStr(of))
		Build(file, doFile)
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
	return fmt.Sprintf(bodyStr,p.Elem())
}

func infoDaoStr(p reflect.Type) string {
	bodyStr := `
	err=db.Db.Where("id=?",id).First(&data).Error
	return `
	return bodyStr
}
