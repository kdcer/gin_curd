package dao
type FloorDao struct{
} 
func(d FloorDao)Create(data model.Floor)(err error){

	err=db.Db.Create(&data) 
	return 
}
func(d FloorDao)Update(id int,data model.Floor)(err error){

	err=db.Db.Where("id=?",id).Update(&data) 
	return 
}
func(d FloorDao)Delete(id int)(err error){

	err=db.Db.Where("id=?",id).Delete(&data) 
	return 
}
func(d FloorDao)Info(id int)(data model.Floor,err error){

	err=db.Db.Where("id=?",id).first(&data) 
	return 
}
