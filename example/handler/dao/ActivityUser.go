package dao
type ActivityUserDao struct{
} 
func(d ActivityUserDao)Create(data model.ActivityUser)(err error){

	err=db.Db.Create(&data) 
	return 
}
func(d ActivityUserDao)Update(id int,data model.ActivityUser)(err error){

	err=db.Db.Where("id=?",id).Update(&data) 
	return 
}
func(d ActivityUserDao)Delete(id int)(err error){

	err=db.Db.Where("id=?",id).Delete(&data) 
	return 
}
func(d ActivityUserDao)Info(id int)(data model.ActivityUser,err error){

	err=db.Db.Where("id=?",id).first(&data) 
	return 
}
