package Backend

import (
	"MyProjectForWB/src"
	"MyProjectForWB/src/JsonStruct"
	"MyProjectForWB/src/Postgres"
	_ "MyProjectForWB/src/Postgres"
	pr "MyProjectForWB/src/Postgres/ProducerStack"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"time"
)

func StanServer(conf *src.Config) {
	var me JsonStruct.JsonStruct
	bam := make([]JsonStruct.JsonStruct, 0, 10)

	dbSql := Postgres.DbConnect(*conf)

	sc, er := stan.Connect(conf.ClusterID, "Ruslan")
	if er != nil {
		fmt.Println(er)
		return
	}
	countInvalidModels := 1
	var sub, err = sc.Subscribe("foo", func(m *stan.Msg) {
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

	pr.Producer(conf)

	time.Sleep(time.Second * 1)

	err = sub.Unsubscribe()
	if err != nil {
		return
	}

	err = sc.Close()
}
