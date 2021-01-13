package ctrl

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type ActivityCtrl struct{
	BaseCtrl
} 
func(c ActivityCtrl)Create(ctx *gin.Context){
 
		var req models.Activity
        if err := ctx.ShouldBindJSON(&req);err==nil{
		   c.resErr40XJson(ctx,400,err,nil)
			return
		}
		err:= dao.ActivityDao{}.Create(req).Error
		if err!=nil{
			c.resErr500Json(ctx, err, nil)
			return
        }
		c.resSuccessJson(ctx, "ok", nil)
	
}
func(c ActivityCtrl)Update(ctx *gin.Context){
 
		param := c.Param("id")
        id, _ := strconv.Atoi(param)
		var req models.Activity
        if err := ctx.ShouldBindJSON(&req);err==nil{
		   c.resErr40XJson(ctx,400,err,nil)
			return
		}
		err:= dao.ActivityDao{}.Update(id,req).Error
		if err!=nil{
			c.resErr500Json(ctx, err, nil)
			return
        }
		c.resSuccessJson(ctx, "ok", nil)
	
}
func(c ActivityCtrl)Delete(ctx *gin.Context){
  
		param := c.Param("id")
        id, _ := strconv.Atoi(param)
	
		err:= dao.ActivityDao{}.Delete(id).Error
			if err!=nil{
			c.resErr500Json(ctx, err, nil)
			return
        }
		c.resSuccessJson(ctx, "ok", nil)
	
}
func(c ActivityCtrl)Info(ctx *gin.Context){
 
		param := c.Param("id")
        id, _ := strconv.Atoi(param)
	
		info,err:= dao.ActivityDao{}.Info(id).Error
			if err!=nil{
			c.resErr500Json(ctx, err, nil)
			return
        }
		c.resSuccessJson(ctx, "ok", map[string]interface{}{"info":info})
	
}
