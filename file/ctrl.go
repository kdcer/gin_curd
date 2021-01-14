package file

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"strings"
)

const (
	ctrlDir = "handler/ctrl/"
)

func init() {
	os.MkdirAll(ctrlDir, 0777)
}

//BuildCtrl 生成ctrl
func BuildCtrl(data []interface{}) {
	var daoPkg string
	for _, da := range data {
		doFile := new(File)

		doFile.Pkg = "ctrl"

		of := reflect.TypeOf(da)
		modelPkg := of.Elem().PkgPath()
		doFile.Import = append(doFile.Import, modelPkg, "github.com/gin-gonic/gin", "strconv")
		if daoPkg == "" {
			daoPkg, _ = path.Split(modelPkg)
			daoPkg = daoPkg + "dao"
		}

		doFile.Import = append(doFile.Import, daoPkg)
		name := of.Elem().Name()
		lowerName := strings.ToLower(string(name[0])) + name[1:]
		file, err := createFile(ctrlDir + lowerName)
		if err != nil {
			panic(err)
		}
		doFile.AddClass(name+"Ctrl", "c").AddParent("BaseCtrl")
		str, search := listCtrlStr(of)
		if search {
			entityPkg, _ := path.Split(modelPkg)
			entityPkg = entityPkg + "entity"
			doFile.Import = append(doFile.Import, entityPkg)
		}
		doFile.NewFunc(name+"Ctrl", "List", "ctx *gin.Context", "", str)
		doFile.NewFunc(name+"Ctrl", "Create", "ctx *gin.Context", "", createCtrlStr(of))
		doFile.NewFunc(name+"Ctrl", "Update", "ctx *gin.Context", "", upCtrlStr(of))
		doFile.NewFunc(name+"Ctrl", "Delete", "ctx *gin.Context", "", delCtrlStr(of))
		doFile.NewFunc(name+"Ctrl", "Info", "ctx *gin.Context", "", infoCtrlStr(of))
		Build(file, doFile)
	}
}

func listCtrlStr(p reflect.Type) (bodyStr string, search bool) {
	_, _, search = entityField(p)
	var code string
	if search {

		code = `	var req %s
	if err := ctx.ShouldBind(&req);err==nil{
		c.resErr40XJson(ctx,400,err,nil)
		return
	}
	list,err:= %sDao{}.List(req)
	if err!=nil{
		c.resErr500Json(ctx, err, nil)
		return
	}
	c.resSuccessJson(ctx, "ok", list)
	`
		bodyStr = fmt.Sprintf(code, "entity."+p.Elem().Name()+"ListReq", "dao."+p.Elem().Name())
		return
	} else {
		code = `	
	list,err:= %sDao{}.List()
	if err!=nil{
		c.resErr500Json(ctx, err, nil)
		return
	}
	c.resSuccessJson(ctx, "ok", list)
	`
	}

	bodyStr = fmt.Sprintf(code, "dao."+p.Elem().Name())
	return

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
	bodyStr := fmt.Sprintf(code, p.Elem().String(), "dao."+p.Elem().Name())
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
	bodyStr := fmt.Sprintf(code, p.Elem().String(), "dao."+p.Elem().Name())
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
	bodyStr := fmt.Sprintf(code, "dao."+p.Elem().Name())
	return bodyStr
}
