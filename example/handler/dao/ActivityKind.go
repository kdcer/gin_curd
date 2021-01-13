package dao
type ActivityKindDao struct{
} 
func(d ActivityKindDao)Create(data model.ActivityKind)(err error){

	err=db.Db.Create(&data) 
	return 
}
func(d ActivityKindDao)Update(id int,data model.ActivityKind)(err error){

	err=db.Db.Where("id=?",id).Update(&data) 
	return 
}
func(d ActivityKindDao)Delete(id int)(err error){

	err=db.Db.Where("id=?",id).Delete(&data) 
	return 
}
func(d ActivityKindDao)Info(id int)(data model.ActivityKind,err error){

	err=db.Db.Where("id=?",id).first(&data) 
	return 
}
