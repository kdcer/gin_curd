package file

import (
	"fmt"
	"os"
	"path"
	"reflect"
)

const ctrlDir = "handler/ctrl/"

func init() {
	os.MkdirAll(ctrlDir, 0777)
}

func BuildCtrl(data []interface{}) {
	var daoPkg string
	for _, da := range data {
		doFile := new(File)

		doFile.Pkg = "ctrl"

		of := reflect.TypeOf(da)
		modelPkg:=of.Elem().PkgPath()
		doFile.Import = append(doFile.Import, modelPkg,"github.com/gin-gonic/gin","strconv")
		if daoPkg=="" {
			daoPkg,_=path.Split(modelPkg)
			daoPkg=daoPkg+"dao"
		}
		doFile.Import = append(doFile.Import, daoPkg)
		name := of.Elem().Name()
		file, err := createFile(ctrlDir+name)
		if err != nil {
			panic(err)
		}
		doFile.SetClass(name + "Ctrl","c").AddParent("BaseCtrl")
		doFile.NewFunc("Create", "ctx *gin.Context", "", createCtrlStr(of))
		doFile.NewFunc("Update", "ctx *gin.Context", "", upCtrlStr(of))
		doFile.NewFunc("Delete", "ctx *gin.Context", "", delCtrlStr(of))
		doFile.NewFunc("Info", "ctx *gin.Context","", infoCtrlStr(of))
		Build(file, doFile)
	}
}




func createCtrlStr(p reflect.Type) string {
	code := `	var req %s
	if err := ctx.ShouldBindJSON(&req);err==nil{
		c.resErr40XJson(ctx,400,err,nil)
		return
	}
	err:= %sDao{}.Create(req)
	if err!=nil{
		c.resErr500Json(ctx, err, nil)
		return
	}
	c.resSuccessJson(ctx, "ok", nil)
	`
	bodyStr := fmt.Sprintf(code,p.Elem().String(), "dao."+p.Elem().Name())
	return bodyStr
}

func upCtrlStr(p reflect.Type) string {
	code := `	param := ctx.Param("id")
	id, _ := strconv.Atoi(param)
	var req %s
	if err := ctx.ShouldBindJSON(&req);err==nil{
	   c.resErr40XJson(ctx,400,err,nil)
	   return
	}
	err:= %sDao{}.Update(id,req)
	if err!=nil{
		c.resErr500Json(ctx, err, nil)
		return
	}
	c.resSuccessJson(ctx, "ok", nil)
	`
	bodyStr := fmt.Sprintf(code, p.Elem().String(),"dao."+p.Elem().Name())
	return bodyStr
}

func delCtrlStr(p reflect.Type) string {
	code := `	param := ctx.Param("id")
	id, _ := strconv.Atoi(param)
	err:= %sDao{}.Delete(id)
	if err!=nil{
		c.resErr500Json(ctx, err, nil)
		return
	}
	c.resSuccessJson(ctx, "ok", nil)
	`
	bodyStr := fmt.Sprintf(code, "dao."+p.Elem().Name())
	return bodyStr
}

func infoCtrlStr(p reflect.Type) string {
	code := `	param := ctx.Param("id")
	id, _ := strconv.Atoi(param)
	info,err:= %sDao{}.Info(id)
	if err!=nil{
		c.resErr500Json(ctx, err, nil)
		return
	}
	c.resSuccessJson(ctx, "ok", map[string]interface{}{"info":info})
	`
	bodyStr := fmt.Sprintf(code,"dao."+p.Elem().Name())
	return bodyStr
}
