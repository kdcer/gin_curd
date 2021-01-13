package dao
type ProductDao struct{
} 
func(d ProductDao)Create(data model.Product)(err error){

	err=db.Db.Create(&data) 
	return 
}
func(d ProductDao)Update(id int,data model.Product)(err error){

	err=db.Db.Where("id=?",id).Update(&data) 
	return 
}
func(d ProductDao)Delete(id int)(err error){

	err=db.Db.Where("id=?",id).Delete(&data) 
	return 
}
func(d ProductDao)Info(id int)(data model.Product,err error){

	err=db.Db.Where("id=?",id).first(&data) 
	return 
}
