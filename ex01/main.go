package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	timeNtp, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error while getting time: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Current time: %s\n", timeNtp.Format(time.RFC1123))
}
