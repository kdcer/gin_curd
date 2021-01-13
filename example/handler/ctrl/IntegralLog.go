package ctrl
type IntegralLogCtrl struct{
	BaseCtrl
} 
func(c IntegralLogCtrl)Create(ctx *gin.Context){
 
		var req models.IntegralLog
        if err := ctx.ShouldBindJSON(&req);err==nil{
		   c.resErr40XJson(ctx,400,err,nil)
			return
		}
		err:= dao.IntegralLogDao{}.Create(req).Error
		if err!=nil{
			c.resErr500Json(ctx, err, nil)
			return
        }
		c.resSuccessJson(ctx, "ok", nil)
	
}
func(c IntegralLogCtrl)Update(ctx *gin.Context){
 
		param := c.Param("id")
        id, _ := strconv.Atoi(param)
		var req models.IntegralLog
        if err := ctx.ShouldBindJSON(&req);err==nil{
		   c.resErr40XJson(ctx,400,err,nil)
			return
		}
		err:= dao.IntegralLogDao{}.Update(id,req).Error
		if err!=nil{
			c.resErr500Json(ctx, err, nil)
			return
        }
		c.resSuccessJson(ctx, "ok", nil)
	
}
func(c IntegralLogCtrl)Delete(ctx *gin.Context){
  
		param := c.Param("id")
        id, _ := strconv.Atoi(param)
	
		err:= dao.IntegralLogDao{}.Delete(id).Error
			if err!=nil{
			c.resErr500Json(ctx, err, nil)
			return
        }
		c.resSuccessJson(ctx, "ok", nil)
	
}
func(c IntegralLogCtrl)Info(ctx *gin.Context){
 
		param := c.Param("id")
        id, _ := strconv.Atoi(param)
	
		info,err:= dao.IntegralLogDao{}.Info(id).Error
			if err!=nil{
			c.resErr500Json(ctx, err, nil)
			return
        }
		c.resSuccessJson(ctx, "ok", map[string]interface{}{"info":info})
	
}
