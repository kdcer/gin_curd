package dao
type UserProductDao struct{
} 
func(d UserProductDao)Create(data model.UserProduct)(err error){

	err=db.Db.Create(&data) 
	return 
}
func(d UserProductDao)Update(id int,data model.UserProduct)(err error){

	err=db.Db.Where("id=?",id).Update(&data) 
	return 
}
func(d UserProductDao)Delete(id int)(err error){

	err=db.Db.Where("id=?",id).Delete(&data) 
	return 
}
func(d UserProductDao)Info(id int)(data model.UserProduct,err error){

	err=db.Db.Where("id=?",id).first(&data) 
	return 
}
