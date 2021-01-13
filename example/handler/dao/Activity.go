package dao
type ActivityDao struct{
} 
func(d ActivityDao)Create(data model.Activity)(err error){

	err=db.Db.Create(&data) 
	return 
}
func(d ActivityDao)Update(id int,data model.Activity)(err error){

	err=db.Db.Where("id=?",id).Update(&data) 
	return 
}
func(d ActivityDao)Delete(id int)(err error){

	err=db.Db.Where("id=?",id).Delete(&data) 
	return 
}
func(d ActivityDao)Info(id int)(data model.Activity,err error){

	err=db.Db.Where("id=?",id).first(&data) 
	return 
}
