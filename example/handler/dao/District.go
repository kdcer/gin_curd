package dao
type DistrictDao struct{
} 
func(d DistrictDao)Create(data model.District)(err error){

	err=db.Db.Create(&data) 
	return 
}
func(d DistrictDao)Update(id int,data model.District)(err error){

	err=db.Db.Where("id=?",id).Update(&data) 
	return 
}
func(d DistrictDao)Delete(id int)(err error){

	err=db.Db.Where("id=?",id).Delete(&data) 
	return 
}
func(d DistrictDao)Info(id int)(data model.District,err error){

	err=db.Db.Where("id=?",id).first(&data) 
	return 
}
