package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/pterm/pterm"
)

const version = "0.8"

func printErrExit(a ...interface{}) {
	fmt.Fprintf(os.Stderr, "ttime: ")
	fmt.Fprint(os.Stderr, a...)
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(1)
}

func notify(message string, sound bool) {
	if sound {
		err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
		if err != nil {
			printErrExit("failed to play notification sound: ", err)
		}
	}
	err := beeep.Notify("ttime", message, "")
	if err != nil {
		printErrExit("failed to send notification: ", err)
	}
}

func main() {
	nosound := flag.Bool("nosound", false, "play no sound when sending a notification")
	nonotify := flag.Bool("nonotify", false, "don't send a notification when a timer is finished")
	update := flag.Bool("u", false, "update time every second")
	flag.Parse()

	bwstyle := pterm.NewStyle(pterm.FgWhite, pterm.BgBlack)

	switch strings.ToLower(flag.Arg(0)) {
	case "":
		if *update {
			for range time.Tick(time.Second) {
				fmt.Printf("\r%s", time.Now().Format(time.RFC1123))
			}
		} else {
			fmt.Printf("%s\n", time.Now().Format(time.RFC1123))
		}

	case "stopwatch":
		starttime := time.Now()
		fmt.Println("Press Enter to stop")

		go func() {
			for range time.Tick(time.Millisecond) {
				fmt.Printf("\rRunning stopwatch: %s", time.Now().Sub(starttime).Truncate(time.Millisecond))
			}
		}()
		fmt.Scanln()

	case "alarm":
		if len(flag.Args()) < 2 {
			printErrExit("not enough arguments: no time provided")
		}

		alarmtime, err := time.ParseInLocation("02.01.2006-15:04:05", flag.Arg(1), time.Local)
		if err != nil {
			printErrExit("failed to parse time: ", err)
		}

		now := time.Now()

		duration := alarmtime.Sub(now)
		if err != nil {
			printErrExit(err)
		}

		if alarmtime.Before(now) {
			printErrExit("please provide a time in the future")
		}

		fmt.Printf("Setting an alarm to %s (%s)\n", alarmtime.Round(time.Second).String(), duration.Round(time.Second).String())

		timer := time.NewTimer(duration)
		go func() {
			// progressbar
			bar, err := pterm.DefaultProgressbar.WithTotal(int(duration.Seconds())).WithTitle("Alarm: ").WithBarStyle(bwstyle).Start()
			if err != nil {
				printErrExit(err)
			}

			for i := 0; i < bar.Total; i++ {
				bar.Increment()
				time.Sleep(time.Second)
			}
			fmt.Println("\nFinished")
		}()
		<-timer.C

		if !*nonotify {
			notify("Alarm finished!", !*nosound)
		}

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
		fmt.Printf("Timer for %d days, %d hours, %d minutes and %d seconds\n\n", days, hours, minutes, seconds)

		duration, err := time.ParseDuration(strconv.Itoa(timeInSeconds) + "s")
		if err != nil {
			printErrExit(err)
		}

		timer := time.NewTimer(duration)
		go func() {
			// progressbar
			bar, err := pterm.DefaultProgressbar.WithTotal(timeInSeconds).WithTitle("Timer: ").WithBarStyle(bwstyle).Start()
			if err != nil {
				printErrExit(err)
			}

			for i := 0; i < bar.Total; i++ {
				bar.Increment()
				time.Sleep(time.Second)
			}
			fmt.Println("\nFinished")
		}()
		<-timer.C

		if !*nonotify {
			notify("Timer finished!", !*nosound)
		}
	case "help":
		fmt.Printf("ttime help:\n")
		fmt.Println("\tstopwatch: create a stopwatch")
		fmt.Println("\ttimer N: create a timer for a duration of N")
		fmt.Println("\talarm N: create an alarm at time N")
		fmt.Println("\tversion: print the currently used version of ttime")
		fmt.Println("\nsee ttime -h for available flags")
	case "version":
		fmt.Printf("ttime version %s\n", version)
	case "license":
		fmt.Println("(C) Palexer, 2021\nttime is licensed under the BSD-3-Clause License\nSee the LICENSE file for more details.\nRepository: https://github.com/Palexer/ttime")
	case "copyright":
		fmt.Println("(C) Palexer, 2021\nttime is licensed under the BSD-3-Clause License\nSee the LICENSE file for more details.\nRepository: https://github.com/Palexer/ttime")
	default:
		printErrExit("command not found\navailable commands: timer, stopwatch, alarm")
	}
}
