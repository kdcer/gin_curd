package dao
type UserTaskDao struct{
} 
func(d UserTaskDao)Create(data model.UserTask)(err error){

	err=db.Db.Create(&data) 
	return 
}
func(d UserTaskDao)Update(id int,data model.UserTask)(err error){

	err=db.Db.Where("id=?",id).Update(&data) 
	return 
}
func(d UserTaskDao)Delete(id int)(err error){

	err=db.Db.Where("id=?",id).Delete(&data) 
	return 
}
func(d UserTaskDao)Info(id int)(data model.UserTask,err error){

	err=db.Db.Where("id=?",id).first(&data) 
	return 
}
