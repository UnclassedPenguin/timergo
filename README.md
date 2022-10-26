# TimerGo

#### A simple timer for the command line written in go.

Go has become my favorite language of late. Previously I had written this simple command line timer in python, but wanted to write it in go instead, so I did!

Just a note, I use linux. These have never been intended to run on windows. If you'd like to fix it though, feel free to send a pull request. 

And another note. The notification is supposed to beep. On my system, I haven't been able to get it to work. So I have no idea if it will work on other systems or not, but you may get an annoying beep. I have no idea how loud/annoying it is...So just be aware if you try running it. 

## To install

```shell
$ go install github.com/unclassedpenguin/timergo@latest
```


## Usage
To just use a simple timer that will run infinitely

```shell
$ timergo
```

---

To do a simple time with a limit, that will pop up a desktop notification at the end of the limit:

```shell
$ timergo 30
```

This would run a 30 second timer. The time increment as of now is only accepted as seconds. 

---

There are a few other options as well. You can use -b to count in binary. If you want to set a time limit as well
as use binary, you must use the -t option.

```shell
$ timergo -b -t 30
```

This would run a 30 second timer, and the output would be in binary. 

---

The last command is -c. This is the "command" option, and allows you to run a command after the timer is reached. This option requires -t.

```shell
$ timergo -t 30 -c "echo 'hello'"
```

This would run the command echo hello after the 30 seconds is up.

## Todo

- Add ability to take time as minutes/hours at least. things like "5m" or "1h"
