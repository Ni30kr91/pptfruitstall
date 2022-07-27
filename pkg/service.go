package pkg

func FruitInsertService(reqbody Fruits) bool {

	res := FruitInsert(reqbody)

	return res
}

func GetFruitsService() []Fruits {

	data := GetFruits()
	return data
}

func UpdateFruitsService(reqbody Fruits) bool {

	res := FruitUpdate(reqbody)

	return res
}

func DeleteFruitsService(id int) bool {

	res := FruitDelete(id)

	return res
}
