package main

import (
	"bufio"
	"fmt"
	"os"
    "flag"
	"time"

    "github.com/logrusorgru/aurora"
)

var slowMS *int
var useColor *bool
var lineTime *bool

func init() {
    useColor = flag.Bool("c", false, "Use Color")
    slowMS = flag.Int("s", 2, "If using color, Time in MS to color slow requests in Red")
    lineTime = flag.Bool("t", false, "Show time between lines")
	flag.Parse()
}

func main() {
    var ts time.Time
    var lastStamp time.Time
    var timeOut string
    var au aurora.Aurora
    au = aurora.NewAurora(*useColor)

/*    if *useColor {
        errTime := *slowMS 
        warnTime:= (*slowMS/2)
    }
*/
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
        if *lineTime {
            ts=time.Now()
            timeOut=fmt.Sprintf("(%7s): ",ts.Sub(lastStamp).Truncate(time.Millisecond))
            fmt.Printf("%s %s %s\n",au.BrightBlack(ts.Format(time.StampMilli)),au.Green(timeOut),input) 
            lastStamp=ts
        } else {
            ts=time.Now()
            fmt.Printf("%s %s\n",au.BrightBlack(ts.Format(time.StampMilli)),input) 
        }
    }
}
