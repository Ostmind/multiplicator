package main

import (
	"flag"
	"github.com/Ostmind/multiplicator/cmd/server"
	"log"
)

func main() {
	rtpPtr := flag.Float64("rtp", 0, "rtp parameter (must be > 0 and ≤ 1.0)")
	flag.Parse()

	if *rtpPtr <= 0 || *rtpPtr > 1.0 {
		log.Fatal("rtp должен быть > 0 и ≤ 1.0")
	}
	server.StartServer(*rtpPtr)
}
