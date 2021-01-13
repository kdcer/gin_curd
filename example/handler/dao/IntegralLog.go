package dao
type IntegralLogDao struct{
} 
func(d IntegralLogDao)Create(data model.IntegralLog)(err error){

	err=db.Db.Create(&data) 
	return 
}
func(d IntegralLogDao)Update(id int,data model.IntegralLog)(err error){

	err=db.Db.Where("id=?",id).Update(&data) 
	return 
}
func(d IntegralLogDao)Delete(id int)(err error){

	err=db.Db.Where("id=?",id).Delete(&data) 
	return 
}
func(d IntegralLogDao)Info(id int)(data model.IntegralLog,err error){

	err=db.Db.Where("id=?",id).first(&data) 
	return 
}
