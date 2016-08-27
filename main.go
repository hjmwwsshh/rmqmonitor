package main

import (
	"flag"
	"rmqmon/g"
	"time"
	"rmqmon/falcon"
	"fmt"
	"os"
)

func Collect() {
	go collect(g.Config().Interval)
}

func collect(sec int64) {
	t := time.NewTicker(time.Second * time.Duration(sec)).C
	for {
		<-t
		falcon.Collector()
	}
}

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	ver := flag.Bool("v", false, "show agent version")

	flag.Parse()

	if *ver {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)

	Collect()

	select {}

}
