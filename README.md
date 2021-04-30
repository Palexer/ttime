# ttime (terminal time)

ttime is a command line program to create timers, stopwatches, alarms
and to get the current time.

## Installation

Prebuilt binaries are available for Windows (amd64), macOS (amd64, arm64) and Linux (amd64, arm64).
You can download them [here](https://github.com/Palexer/ttime/releases).
Alternatively, you can compile from source by cloning this repository and running:

```
make
```
to build the binary and 

```
sudo make install
```

to build and install the binary.

## Usage

### Current time

```
ttime
```

You can also use the -u flag to update the time every second:

```
ttime -u
```

### Stopwatch

```
ttime stopwatch
```

Then press enter to stop the stopwatch.

### Timer

The following will create a timer for 15 seconds

```
ttime timer 15
```

To add minutes, hours and/or days, put the values infront of the smaller unit like this:

```
ttime timer 3-15
```
will create a timer for 3 minutes and 15 seconds.

```
ttimer timer 1-2-0-30
```
will create a timer for one day, two hours and 30 seconds

### Alarm

In order to create an alarm, you have to pass the date and the exact time like this:

```
ttime alarm 2.1.2006-15:04:05
```

The command above will create an alarm on the 2nd of January 2006 at 3:04:05 PM
(the time has to be passed in the 24-hour-format).


## Libraries

- github.com/pterm/pterm for pretty terminal output
- github.com/gen2brain/beeep for notifications

## LICENSE

BSD 3-Clause License
