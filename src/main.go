package main

import "time"
import "os/signal"
import "os"
// import "fmt"
// import "strconv"

var log Logger = Logger{flag: "test", level: 7}.new("start")
var bot Bot

func main() {
  log.notice("program started")

  log.info("setting up io")
  bot = Bot{
    battery: Battery{}.new(),
    colorSensor: ColorSensor{port: IN_2}.new(),
    speaker: Speaker{}.new(),
    button: Button{
      onKeypress: func (key int, state int) {
        if key == KEY_ESCAPE {
          end("escape")
        }
      },
    }.new(),
  }

  log.inc(":status")
  if bot.battery.voltage() > 72 {
    log.info("voltage is at " + log.value(bot.battery.voltageString() + "v"))
  } else {
    log.warn("voltage is at " + log.value(bot.battery.voltageString() + "v"))
  }
  log.dec()

  go bot.speaker.song([]int{300, 400, 500, 600}, 100, 1)

  log.info("looping")
  log.rep("loop")

  setupInterrupt()
  loop()
}

func loop() {
  time.Sleep(time.Second / 20)
  log.trace("looping")
  loop()
}

func setupInterrupt() {
  stop := make(chan os.Signal, 1)
  signal.Notify(stop, os.Interrupt)
  go func() {
    <-stop
    end("ctrl-c")
  }()
}

func end(catch string) {
  log.set([]string{"end"})

  go bot.speaker.song([]int{600, 500, 400, 300}, 100, 1)
  log.notice("caught " + catch)
  log.level = 0
  time.Sleep(time.Millisecond * 500)
  os.Exit(0)
}