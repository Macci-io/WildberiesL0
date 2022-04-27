package main

import (
	"flag"
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
)

type Config struct {
	UserBase  string
	ClusterID string
	PassBase  string
	AddrBase  string
	NameDB    string
}

func getAllFilesByte(path string) (result [][]byte, err error) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, v := range dir {
		if !v.IsDir() { //&& strings.HasSuffix(v.Name(), ".json")
			file, err := ioutil.ReadFile(path + "/" + v.Name())
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			result = append(result, file)
		}
	}
	return result, err
}

func main() {
	conf := &Config{}

	flag.StringVar(&conf.UserBase, "u", "one", "Psql name")
	flag.StringVar(&conf.ClusterID, "cid", "test-cluster", "Named ClusterID")
	flag.StringVar(&conf.PassBase, "p", "sensation05", "Password of Database")
	flag.StringVar(&conf.AddrBase, "h", "localhost", "Address of Database")
	flag.StringVar(&conf.NameDB, "d", "one", "Name of Database")

	flag.Parse()

	bam, er := getAllFilesByte("/home/one/GolandProjects/MyProjectForWB/src/ProducerStack")

	sc, er := stan.Connect(conf.ClusterID, "Rodrig")
	if er != nil {
		fmt.Println(er)
		return
	}
	fmt.Println(len(bam))
	for _, v := range bam {
		er = sc.Publish("foo", v)
	}

	er = sc.Close()
}
