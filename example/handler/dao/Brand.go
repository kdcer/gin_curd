package dao
type BrandDao struct{
} 
func(d BrandDao)Create(data model.Brand)(err error){

	err=db.Db.Create(&data) 
	return 
}
func(d BrandDao)Update(id int,data model.Brand)(err error){

	err=db.Db.Where("id=?",id).Update(&data) 
	return 
}
func(d BrandDao)Delete(id int)(err error){

	err=db.Db.Where("id=?",id).Delete(&data) 
	return 
}
func(d BrandDao)Info(id int)(data model.Brand,err error){

	err=db.Db.Where("id=?",id).first(&data) 
	return 
}
