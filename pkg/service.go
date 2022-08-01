package pkg

import (
	"encoding/json"
	"fmt"
	"pptfruitstall/server/db"
	"time"
)

func FruitInsertService(reqbody Fruits) bool {

	res := FruitInsert(reqbody)
	if res == true {
		data := GetFruits()
		jsonData, _ := json.Marshal(data)
		r := db.Client.Set("all", jsonData, time.Minute*30)

		fmt.Println("test", r)
	}
	return res
}

func GetFruitsService() (f []Fruits) {

	val, err := db.Client.Get("all").Bytes()
	json.Unmarshal(val, &f)
	if err != nil {
		return f
	}
	fmt.Println("data from redis: ", f, err)

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
