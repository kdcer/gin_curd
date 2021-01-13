package dao
type UserSignUpDao struct{
} 
func(d UserSignUpDao)Create(data model.UserSignUp)(err error){

	err=db.Db.Create(&data) 
	return 
}
func(d UserSignUpDao)Update(id int,data model.UserSignUp)(err error){

	err=db.Db.Where("id=?",id).Update(&data) 
	return 
}
func(d UserSignUpDao)Delete(id int)(err error){

	err=db.Db.Where("id=?",id).Delete(&data) 
	return 
}
func(d UserSignUpDao)Info(id int)(data model.UserSignUp,err error){

	err=db.Db.Where("id=?",id).first(&data) 
	return 
}
