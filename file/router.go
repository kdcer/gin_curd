package file

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"strings"
)

const dir = "router"
const routerFile = "router"

var (

	router *File
)

func init() {

	router = new(File)
	_ = os.Mkdir(dir, 0777)


}
//BuildRouter 生成路由
func BuildRouter(data []interface{}) {
	file, err := createFile(dir+"/"+routerFile)
	if err != nil {
		panic(err)
	}
	var ctrlPkg string
	router.Pkg = "router \n"
	router.Import = append(router.Import, "github.com/gin-gonic/gin")

	router.NewFunc("InitRouter", "", "engine *gin.Engine", "\tengine = gin.Default() \n")
	for _, da := range data {
		of := reflect.TypeOf(da)
		modelPkg:=of.Elem().PkgPath()

		if ctrlPkg=="" {
			ctrlPkg,_=path.Split(modelPkg)
			router.Import = append(router.Import, ctrlPkg+"ctrl")
		}

		name := of.Elem().Name()
		lowerName := strings.ToLower(string(name[0])) + name[1:]
		body := fmt.Sprintf("\t%s:=engine.Group(\"%s\") \n \t { \n", lowerName, lowerName)
		//body += fmt.Sprintf("\t\t%s.GET(\"\",%s.%sCtrl.%s) \n", lowerName, "ctrl", name, "List")
		ctrlStr:=fmt.Sprintf("\t\t%sCtrl:=new(ctrl.%sCtrl)\n",lowerName,name)
		body += fmt.Sprintf(ctrlStr)
		body += fmt.Sprintf("\t\t%s.GET(\"/:id\",%sCtrl.%s) \n", lowerName, lowerName, "Info")
		body += fmt.Sprintf("\t\t%s.POST(\"\",%sCtrl.%s) \n", lowerName, lowerName, "Create")
		body += fmt.Sprintf("\t\t%s.PUT(\"/:id\",%sCtrl.%s) \n", lowerName, lowerName, "Update")
		body += fmt.Sprintf("\t\t%s.DELETE(\"/:id\",%sCtrl.%s) \n", lowerName, lowerName, "Delete")
		body += "\t } \n"
		if findFunc := router.findFunc("InitRouter");findFunc!=nil{

			findFunc.AddBody(body)
		}

	}
	router.findFunc("InitRouter").AddBody("\treturn engine \n")
	Build(file, router)

}
