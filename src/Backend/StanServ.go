package Backend

import (
	"MyProjectForWB/src"
	"MyProjectForWB/src/JsonStruct"
	"MyProjectForWB/src/Postgres"
	_ "MyProjectForWB/src/Postgres"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"log"
)

func StanServer(conf *src.Config) *[]JsonStruct.JsonStruct {
	var me JsonStruct.JsonStruct
	bam := make([]JsonStruct.JsonStruct, 0, 10)

	dbSql := Postgres.DbConnect(*conf)

	var jsonData []byte
	query, er := dbSql.Query("SELECT model FROM models;")
	if er != nil {
		log.Panic(er)
	}

	for query.Next() {
		err := query.Scan(&jsonData)
		if err != nil {
			log.Panic(err)
		}
		e := json.Unmarshal(jsonData, &me)
		if e != nil {
			fmt.Print(e)
		} else {
			bam = append(bam, me)
		}

	}

	sc, er := stan.Connect(conf.ClusterID, "Ruslan")
	if er != nil {
		fmt.Println(er)
	}
	countInvalidModels := 1
	_, err := sc.Subscribe("foo", func(m *stan.Msg) {
		err := json.Unmarshal(m.Data, &me)
		if err != nil {
			fmt.Printf("Invalid model %d\n", countInvalidModels)
			countInvalidModels++
			return
		}
		bam = append(bam, me)

		_, err = dbSql.Exec("insert into models (model) values ($1);", m.Data)
		if err != nil {
			fmt.Println(err)
		}
	})
	if err != nil {
		log.Fatal(err)
	}
	//
	//err = sub.Unsubscribe()
	//if err != nil {
	//	fmt.Println(err)
	//}

	//defer sc.Close()
	return &bam
}
