package ctrl
type BrandKindCtrl struct{
	BaseCtrl
} 
func(c BrandKindCtrl)Create(ctx *gin.Context){
 
		var req models.BrandKind
        if err := ctx.ShouldBindJSON(&req);err==nil{
		   c.resErr40XJson(ctx,400,err,nil)
			return
		}
		err:= dao.BrandKindDao{}.Create(req).Error
		if err!=nil{
			c.resErr500Json(ctx, err, nil)
			return
        }
		c.resSuccessJson(ctx, "ok", nil)
	
}
func(c BrandKindCtrl)Update(ctx *gin.Context){
 
		param := c.Param("id")
        id, _ := strconv.Atoi(param)
		var req models.BrandKind
        if err := ctx.ShouldBindJSON(&req);err==nil{
		   c.resErr40XJson(ctx,400,err,nil)
			return
		}
		err:= dao.BrandKindDao{}.Update(id,req).Error
		if err!=nil{
			c.resErr500Json(ctx, err, nil)
			return
        }
		c.resSuccessJson(ctx, "ok", nil)
	
}
func(c BrandKindCtrl)Delete(ctx *gin.Context){
  
		param := c.Param("id")
        id, _ := strconv.Atoi(param)
	
		err:= dao.BrandKindDao{}.Delete(id).Error
			if err!=nil{
			c.resErr500Json(ctx, err, nil)
			return
        }
		c.resSuccessJson(ctx, "ok", nil)
	
}
func(c BrandKindCtrl)Info(ctx *gin.Context){
 
		param := c.Param("id")
        id, _ := strconv.Atoi(param)
	
		info,err:= dao.BrandKindDao{}.Info(id).Error
			if err!=nil{
			c.resErr500Json(ctx, err, nil)
			return
        }
		c.resSuccessJson(ctx, "ok", map[string]interface{}{"info":info})
	
}
