package dev01

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

/*
	Use go get github.com/eashabalin/accuratetime to get this module
*/

func GetTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting time from ntp pool: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(time)
}
