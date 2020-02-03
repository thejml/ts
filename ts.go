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
var stripe *bool

func init() {
    slowMS = flag.Int("m", 2, "If using color, Time in MS to color slow requests in Red")
    useColor = flag.Bool("c", false, "Use Color")
    lineTime = flag.Bool("t", false, "Show time between lines")
    stripe   = flag.Bool("s", false, "Use a striped output")
    flag.Parse()
}

func main() {
    var ts time.Time
    var lastStamp time.Time
    var timeOut string
    var lineCount int
    var bg uint8
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
    lineCount = 0
    lastStamp=time.Now()
	// Read in from the pipe until we can't.
	for reader.Scan() {
		input := reader.Text()
        if *lineTime {
            ts=time.Now()
            timeOut=fmt.Sprintf("(%7s): ",ts.Sub(lastStamp).Truncate(time.Millisecond))
            if *stripe {
                if lineCount%2==0 {
                    bg=3
                } else {
                    bg=0
                }
                fmt.Printf("%s %s %s\n",au.BrightBlack(ts.Format(time.StampMilli)).BgGray(bg),au.Green(timeOut).BgGray(bg),au.BgGray(bg,input)) 
            } else {
                fmt.Printf("%s %s %s\n",au.BrightBlack(ts.Format(time.StampMilli)),au.Green(timeOut),input) 
            }
            lastStamp=ts
        } else {
            ts=time.Now()
            fmt.Printf("%s %s\n",au.BrightBlack(ts.Format(time.StampMilli)),input) 
        }
        lineCount++
    }
}
