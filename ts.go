package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
    var ts time.Time
    var lastStamp time.Time

    info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: <stdin> | ts")
		return
	}

	reader := bufio.NewScanner(os.Stdin)
    lastStamp=time.Now()
	// Read in from the pipe until we can't.
	for reader.Scan() {
		input := reader.Text()
        ts=time.Now()
        fmt.Printf("%s (%7s): %s\n",ts.Format(time.StampMilli),ts.Sub(lastStamp).Truncate(time.Millisecond),input)
        lastStamp=ts
    }
}
