package main

import (
	"MyProjectForWB/src"
	"MyProjectForWB/src/Backend"
	"MyProjectForWB/src/Frontend"
)

func main() {
	conf := src.ParFla()
	bam := Backend.StanServer(conf)

	Frontend.WebServ(bam)
}
