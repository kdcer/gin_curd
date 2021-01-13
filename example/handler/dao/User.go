package dao
type UserDao struct{
} 
func(d UserDao)Create(data model.User)(err error){

	err=db.Db.Create(&data) 
	return 
}
func(d UserDao)Update(id int,data model.User)(err error){

	err=db.Db.Where("id=?",id).Update(&data) 
	return 
}
func(d UserDao)Delete(id int)(err error){

	err=db.Db.Where("id=?",id).Delete(&data) 
	return 
}
func(d UserDao)Info(id int)(data model.User,err error){

	err=db.Db.Where("id=?",id).first(&data) 
	return 
}
