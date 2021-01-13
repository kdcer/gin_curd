package dao
type BannerDao struct{
} 
func(d BannerDao)Create(data model.Banner)(err error){

	err=db.Db.Create(&data) 
	return 
}
func(d BannerDao)Update(id int,data model.Banner)(err error){

	err=db.Db.Where("id=?",id).Update(&data) 
	return 
}
func(d BannerDao)Delete(id int)(err error){

	err=db.Db.Where("id=?",id).Delete(&data) 
	return 
}
func(d BannerDao)Info(id int)(data model.Banner,err error){

	err=db.Db.Where("id=?",id).first(&data) 
	return 
}
