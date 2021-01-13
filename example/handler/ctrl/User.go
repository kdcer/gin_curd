package ctrl
type UserCtrl struct{
	BaseCtrl
} 
func(c UserCtrl)Create(ctx *gin.Context){
 
		var req models.User
        if err := ctx.ShouldBindJSON(&req);err==nil{
		   c.resErr40XJson(ctx,400,err,nil)
			return
		}
		err:= dao.UserDao{}.Create(req).Error
		if err!=nil{
			c.resErr500Json(ctx, err, nil)
			return
        }
		c.resSuccessJson(ctx, "ok", nil)
	
}
func(c UserCtrl)Update(ctx *gin.Context){
 
		param := c.Param("id")
        id, _ := strconv.Atoi(param)
		var req models.User
        if err := ctx.ShouldBindJSON(&req);err==nil{
		   c.resErr40XJson(ctx,400,err,nil)
			return
		}
		err:= dao.UserDao{}.Update(id,req).Error
		if err!=nil{
			c.resErr500Json(ctx, err, nil)
			return
        }
		c.resSuccessJson(ctx, "ok", nil)
	
}
func(c UserCtrl)Delete(ctx *gin.Context){
  
		param := c.Param("id")
        id, _ := strconv.Atoi(param)
	
		err:= dao.UserDao{}.Delete(id).Error
			if err!=nil{
			c.resErr500Json(ctx, err, nil)
			return
        }
		c.resSuccessJson(ctx, "ok", nil)
	
}
func(c UserCtrl)Info(ctx *gin.Context){
 
		param := c.Param("id")
        id, _ := strconv.Atoi(param)
	
		info,err:= dao.UserDao{}.Info(id).Error
			if err!=nil{
			c.resErr500Json(ctx, err, nil)
			return
        }
		c.resSuccessJson(ctx, "ok", map[string]interface{}{"info":info})
	
}
