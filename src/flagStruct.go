package src

import "flag"

type Config struct {
	UserBase  string
	ClusterID string
}

func ParFla() *Config {
	conf := &Config{}

	flag.StringVar(&conf.UserBase, "u", "one", "Psql name")
	flag.StringVar(&conf.ClusterID, "cid", "test-cluster", "Named ClusterID")

	flag.Parse()
	return conf
}
