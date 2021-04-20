package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pterm/pterm"
)

func printErrExit(a ...interface{}) {
	fmt.Fprintf(os.Stderr, "ttime: ")
	fmt.Fprint(os.Stderr, a...)
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

func main() {
	nosound := flag.Bool("nosound", false, "play no sound when sending a notification")
	nonotify := flag.Bool("nonotify", false, "don't send a notification when a timer is finished")
	flag.Parse()

	switch flag.Arg(0) {
	case "":
		pterm.FgYellow.Printf("%s\n", time.Now().Format(time.RFC1123))
	case "stopwatch":
		starttime := time.Now()
		spinner, err := pterm.DefaultSpinner.WithRemoveWhenDone(true).Start("Running Stopwatch")
		if err != nil {
			printErrExit(err)
		}
		fmt.Scanln()
		stoptime := time.Now()

		pterm.FgYellow.Printf("Measured time: %s\n", stoptime.Sub(starttime).Round(time.Millisecond))
		spinner.Stop()
	case "alarm":
		if len(flag.Args()) < 2 {
			printErrExit("not enough arguments: no time provided")
		}
		// genaues Datum, 24-Stunden bzw. lokales Format
		alarmtime, err := time.ParseInLocation("2-1-2006:3-4-5", flag.Arg(1), time.Local)
		if err != nil {
			printErrExit("failed to parse time: ", err)
		}
		println(alarmtime.String())

	case "timer":
		if len(flag.Args()) < 2 {
			printErrExit("not enough arguments: no duration provided")
		}
		durationStr := strings.Split(flag.Arg(1), "-")
		length := len(durationStr)
		var seconds, minutes, hours, days int

		seconds, err := strconv.Atoi(durationStr[length-1])
		if err != nil {
			printErrExit(err)
		}

		if length > 1 {
			minutes, err = strconv.Atoi(durationStr[len(durationStr)-2])
			if err != nil {
				printErrExit(err)
			}

			if length > 2 {
				hours, err = strconv.Atoi(durationStr[len(durationStr)-3])
				if err != nil {
					printErrExit(err)
				}

				if length > 3 {
					days, err = strconv.Atoi(durationStr[len(durationStr)-4])
					if err != nil {
						printErrExit(err)
					}
				}
			}
		}

		timeInSeconds := days*24*60*60 + hours*60*60 + minutes*60 + seconds
		pterm.FgYellow.Printf("Timer for %d days, %d hours, %d minutes and %d seconds\n\n", days, hours, minutes, seconds)

		duration, err := time.ParseDuration(strconv.Itoa(timeInSeconds) + "s")
		if err != nil {
			printErrExit(err)
		}

		timer := time.NewTimer(duration)
		go func() {
			// progressbar
			bar, err := pterm.DefaultProgressbar.WithTotal(timeInSeconds).WithTitle("Timer: ").Start()
			if err != nil {
				printErrExit(err)
			}

			for i := 0; i < bar.Total; i++ {
				time.Sleep(time.Second)
				bar.Increment()
			}
			pterm.FgYellow.Println("\nFinished")
		}()
		<-timer.C

		if !*nonotify {
			notify("Timer finished!", !*nosound)
		}
	}
}
