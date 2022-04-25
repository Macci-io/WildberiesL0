package ProducerStack

import (
	"MyProjectForWB/src"
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
)

func getAllFilesByte(path string) (result [][]byte) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	for _, v := range dir {
		if !v.IsDir() { //&& strings.HasSuffix(v.Name(), ".json")
			file, err := ioutil.ReadFile(path + "/" + v.Name())
			if err != nil {
				fmt.Println(err)
				return nil
			}
			result = append(result, file)
		}
	}
	return result
}

func Producer(conf *src.Config) {
	sc, er := stan.Connect(conf.ClusterID, "Ruslen")
	if er != nil {
		fmt.Println(er)
		return
	}

	bam := getAllFilesByte("src/Postgres/ProducerStack")

	for _, v := range bam {
		er = sc.Publish("foo", v)
	}

	er = sc.Close()
}
