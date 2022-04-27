package src

import "flag"

type Config struct {
	UserBase  string
	ClusterID string
	PassBase  string
	AddrBase  string
	NameDB    string
}

func ParFla() *Config {
	conf := &Config{}

	flag.StringVar(&conf.UserBase, "u", "one", "Psql name")
	flag.StringVar(&conf.ClusterID, "cid", "test-cluster", "Named ClusterID")
	flag.StringVar(&conf.PassBase, "p", "sensation05", "Password of Database")
	flag.StringVar(&conf.AddrBase, "h", "localhost", "Address of Database")
	flag.StringVar(&conf.NameDB, "d", "one", "Name of Database")

	flag.Parse()
	return conf
}
