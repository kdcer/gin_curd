package dao
type AdDao struct{
} 
func(d AdDao)Create(data model.Ad)(err error){

	err=db.Db.Create(&data) 
	return 
}
func(d AdDao)Update(id int,data model.Ad)(err error){

	err=db.Db.Where("id=?",id).Update(&data) 
	return 
}
func(d AdDao)Delete(id int)(err error){

	err=db.Db.Where("id=?",id).Delete(&data) 
	return 
}
func(d AdDao)Info(id int)(data model.Ad,err error){

	err=db.Db.Where("id=?",id).first(&data) 
	return 
}
