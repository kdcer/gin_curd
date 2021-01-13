package dao
type TaskDao struct{
} 
func(d TaskDao)Create(data model.Task)(err error){

	err=db.Db.Create(&data) 
	return 
}
func(d TaskDao)Update(id int,data model.Task)(err error){

	err=db.Db.Where("id=?",id).Update(&data) 
	return 
}
func(d TaskDao)Delete(id int)(err error){

	err=db.Db.Where("id=?",id).Delete(&data) 
	return 
}
func(d TaskDao)Info(id int)(data model.Task,err error){

	err=db.Db.Where("id=?",id).first(&data) 
	return 
}
