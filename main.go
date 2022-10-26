//------------------------------------------------------------------------------
//------------------------------------------------------------------------------
//
// Tyler(UnclassedPenguin) TimerGo 2022
//
//      Author: Tyler(UnclassedPenguin)
//         URL: https://unclassed.ca
//      GitHub: https://github.com/UnclassedPenguin/timergo.git
// Description: I just wanted a simple timer. Now in Go!
//
//------------------------------------------------------------------------------
//------------------------------------------------------------------------------

package main

import (
  "fmt"
  "os"
  "os/exec"
  "strings"
  "flag"
  "time"
  "strconv"
  "github.com/gen2brain/beeep"
)

// Gets the current time, and returns it as a string "hh:mm:ss" 
func getTime() string {
  currentTime := time.Now()

  currentHour := currentTime.Hour()
  var currentHourString string
  if currentHour < 10 {
    currentHourString = fmt.Sprintf("%02d", currentHour)
  } else {
    currentHourString = strconv.Itoa(currentHour)
  }

  currentMinute := currentTime.Minute()
  var currentMinuteString string
  if currentMinute < 10 {
    currentMinuteString = fmt.Sprintf("%02d", currentMinute)
  } else {
    currentMinuteString = strconv.Itoa(currentMinute)
  }

  currentSecond := currentTime.Second()
  var currentSecondString string
  if currentSecond < 10 {
    currentSecondString = fmt.Sprintf("%02d", currentSecond)
  } else {
    currentSecondString = strconv.Itoa(currentSecond)
  }

  nowTime := currentHourString + ":" + currentMinuteString + ":" + currentSecondString
  return nowTime
}

// Takes a single digit, and returns a binary value with all the leading zeros.
// ie, takes 7, returns 00000111
// This seems to be an unused function? I should probably remove it...
func getBinary(num int) string {
  binary := fmt.Sprintf("%08b", num)
  return binary
}

func executeCmd(command string) {
  splitCommand := strings.Split(command, " ")

  cmd, stdout, stderr := exec.Command(splitCommand[0], splitCommand[1:]...), new(strings.Builder), new(strings.Builder)
  cmd.Stdout = stdout
  cmd.Stderr = stderr
  err := cmd.Run()
  if err != nil {
    fmt.Println("Error: ", err)
  }
  if stdout != nil {
    fmt.Println(stdout.String())
  }

  if stderr != nil {
    fmt.Println(stderr.String())
  }
}

func main() {
  //programName := os.Args[0]
  //fmt.Println(programName)

  var verbose bool
  var binary bool
  var timerTime int
  var timerReached bool
  var command string

  // beep frequency float
  beepFreq := 440.0
  // beep duration in milliseconds
  beepLength := 1000

  flag.BoolVar(&verbose, "v", false, "Print slightly more information")
  flag.BoolVar(&binary, "b", false, "Print the timer in binary")
  flag.IntVar(&timerTime, "t", 0, "Set a time to send a notification, in seconds")
  flag.StringVar(&command, "c", "", "Command to run when timer hit, requires -t")

  flag.Parse()

  if len(os.Args) > 1 && timerTime == 0 {
    if os.Args[1] != "" {
      timerTime, _ = strconv.Atoi(os.Args[1])
    }
  }


  //fmt.Println("DIAGNOSTICS")
  //fmt.Println("----------------------------")
  //fmt.Println("Verbose: ", verbose)
  //fmt.Println(" Binary: ", binary)
  //fmt.Println("  Timer: ", timerTime)
  //fmt.Println("Command: ", command)
  //fmt.Println("   Time: ", getTime())
  //fmt.Println("----------------------------")

  // Program Start

  if command != "" && timerTime == 0 {
    fmt.Println("The -c/--command option requires a -t/--time amount! Try again using -t")
    fmt.Println("Or -h for help")
    os.Exit(1)
  }

  if verbose {
    fmt.Println("Start Time: ", getTime())
  }

  count := 1

  if binary {
    fmt.Println(fmt.Sprintf("%08b", count))
  } else {
    fmt.Println(count)
  }

  time.Sleep(1 * time.Second)
  fmt.Print("\033[1A\033[K")

  // Main loop where it counts in 1 second intervals
  for true {
    count += 1
    if binary {
      fmt.Println(fmt.Sprintf("%08b", count))
    } else {
      fmt.Println(count)
    }
    time.Sleep(1 * time.Second)
    fmt.Print("\033[1A\033[K")

    if count == timerTime {
      if command != "" {
        if binary {
          binaryTimerTime := fmt.Sprintf("%08b", timerTime)
          message := "Time limit of " + binaryTimerTime + "s reached." + "\n" + "Executing command: " + command
          beeep.Beep(beepFreq, beepLength)
          beeep.Notify("TimerGo", message, "")
          fmt.Println("Executing command: ", command)
          //This doesn't work, figure it out? I want it to print out the value to terminal at the end...
          //timerReached = true
          executeCmd(command)
        } else {
          message := "Time limit of " + strconv.Itoa(timerTime) + "s reached." + "\n" + "Executing command: " + command
          beeep.Beep(beepFreq, beepLength)
          beeep.Notify("TimerGo", message, "")
          fmt.Println("Executing command: ", command)
          //This doesn't work, figure it out? I want it to print out the value to terminal at the end...
          //timerReached = true
          executeCmd(command)
        }
      } else {
        if binary {
          binaryTimerTime := fmt.Sprintf("%08b", timerTime)
          message := "Time limit of " + binaryTimerTime + "s reached."
          beeep.Beep(beepFreq, beepLength)
          beeep.Notify("TimerGo", message, "")
        } else {
          message := "Time limit of " + strconv.Itoa(timerTime) + "s reached."
          beeep.Beep(beepFreq, beepLength)
          beeep.Notify("TimerGo", message, "")
        }
      }
      if verbose {
        fmt.Println("Finish Time: ", getTime())
      }
      if timerReached{
        fmt.Println("timerReached")
      }
      os.Exit(0)
    }
  }
}
