package Postgres

import (
	"MyProjectForWB/src"
	"database/sql"
	"fmt"
	"log"
)

func DbConnect(config src.Config) *sql.DB {

	dbSql, err := sql.Open("postgres", "postgresql://"+config.UserBase+":sensation05@localhost/one?sslmode=disable")
	if err != nil {
		fmt.Println("Can't connect to database one")
		log.Fatal(err)
	}

	// open json and fill struct with data from jsonfile

	//file, er := ioutil.ReadFile("MyProjectForWB/src/Postgres/ProducerStack/model.json")
	//if er != nil {
	//	er = fmt.Errorf("error opening %w", er)
	//}

	// create jsonStruct struct instance
	//jsData := new(JsonStruct.JsonStruct)

	// fill jsonStruct with data from json file
	//err = json.Unmarshal(file, &jsData)
	//if err != nil {
	//	fmt.Println("Can't unmarshal json file to jsonStruct")
	//	log.Fatal(err)
	//}

	//	_, err = dbSql.Exec("insert into models (model) values ($1);", file)
	//	if err != nil {
	//		fmt.Println("Can't unmarshall json file to jsonStruct struct...")
	//		log.Fatal(err)

	//queryCalls, er := dbSql.Query("INSERT INTO models (model) VALUES ($1)", file)
	//if er != nil {
	//	fmt.Println(er)
	//	return
	//}
	//defer queryCalls.Close()
	return dbSql
}
