package dao
type BrandKindDao struct{
} 
func(d BrandKindDao)Create(data model.BrandKind)(err error){

	err=db.Db.Create(&data) 
	return 
}
func(d BrandKindDao)Update(id int,data model.BrandKind)(err error){

	err=db.Db.Where("id=?",id).Update(&data) 
	return 
}
func(d BrandKindDao)Delete(id int)(err error){

	err=db.Db.Where("id=?",id).Delete(&data) 
	return 
}
func(d BrandKindDao)Info(id int)(data model.BrandKind,err error){

	err=db.Db.Where("id=?",id).first(&data) 
	return 
}
