package main

import "github.com/gen2brain/beeep"

func notify(message string, sound bool) {
	err := beeep.Notify("ttime", message, "")
	if err != nil {
		printErrExit("failed to send notification: ", err)
	}

	if sound {
		beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
		if err != nil {
			printErrExit("failed to play notification sound: ", err)
		}
	}
}
