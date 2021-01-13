package ctrl
type TaskCtrl struct{
	BaseCtrl
} 
func(c TaskCtrl)Create(ctx *gin.Context){
 
		var req models.Task
        if err := ctx.ShouldBindJSON(&req);err==nil{
		   c.resErr40XJson(ctx,400,err,nil)
			return
		}
		err:= dao.TaskDao{}.Create(req).Error
		if err!=nil{
			c.resErr500Json(ctx, err, nil)
			return
        }
		c.resSuccessJson(ctx, "ok", nil)
	
}
func(c TaskCtrl)Update(ctx *gin.Context){
 
		param := c.Param("id")
        id, _ := strconv.Atoi(param)
		var req models.Task
        if err := ctx.ShouldBindJSON(&req);err==nil{
		   c.resErr40XJson(ctx,400,err,nil)
			return
		}
		err:= dao.TaskDao{}.Update(id,req).Error
		if err!=nil{
			c.resErr500Json(ctx, err, nil)
			return
        }
		c.resSuccessJson(ctx, "ok", nil)
	
}
func(c TaskCtrl)Delete(ctx *gin.Context){
  
		param := c.Param("id")
        id, _ := strconv.Atoi(param)
	
		err:= dao.TaskDao{}.Delete(id).Error
			if err!=nil{
			c.resErr500Json(ctx, err, nil)
			return
        }
		c.resSuccessJson(ctx, "ok", nil)
	
}
func(c TaskCtrl)Info(ctx *gin.Context){
 
		param := c.Param("id")
        id, _ := strconv.Atoi(param)
	
		info,err:= dao.TaskDao{}.Info(id).Error
			if err!=nil{
			c.resErr500Json(ctx, err, nil)
			return
        }
		c.resSuccessJson(ctx, "ok", map[string]interface{}{"info":info})
	
}
