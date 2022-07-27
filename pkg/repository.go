package pkg

import (
	"pptfruitstall/server/db"
)

func FruitInsert(reqbody Fruits) bool {

	sql := `INSERT INTO fruits(fruits_name)VALUES($1)`

	_, err := db.DB.Exec(sql, reqbody.Fruitsname)

	if err != nil {
		return false
	}
	return true
}

func FruitUpdate(reqbody Fruits) bool {

	SQL := `UPDATE fruits SET  fruits_name=$1 WHERE id=$2`

	_, err2 := db.DB.Exec(SQL, reqbody.Fruitsname, reqbody.Id)

	if err2 != nil {
		//log.Fatal("ERror in update: ", err2)
		return false
	}

	return true
}

func GetFruits() []Fruits {
	data := []Fruits{}
	sql := `SELECT id,fruits_name FROM "fruits"`

	rows, _ := db.DB.Query(sql)

	row := Fruits{}

	for rows.Next() {

		rows.Scan(&row.Id, &row.Fruitsname)

		data = append(data, row)
	}

	return data

}

func FruitDelete(id int) bool {
	SQL := `DELETE FROM "fruits"  WHERE id=$1`

	_, err2 := db.DB.Exec(SQL, id)

	if err2 != nil {
		//log.Fatal("ERror in update: ", err2)
		return false
	}

	return true
}
