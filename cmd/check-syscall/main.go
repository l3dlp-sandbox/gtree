package main

import (
	"time"

	"github.com/ddddddO/work/syscaller"
	"github.com/ddddddO/work/syscaller/file"
)

const (
	duration = 1 * time.Second
)

func main() {
	interval := time.NewTicker(duration)
	end := time.NewTicker(duration * 5)

END:
	for {
		select {
		case <-interval.C:
			sc := file.Gen()
			syscaller.Run(sc)
		case <-end.C:
			break END
		}
	}
}