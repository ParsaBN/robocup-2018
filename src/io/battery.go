package main

// import "strings"

type Battery struct {
  path string
  device Device
}

func (battery Battery) new() Battery {
  battery.path = BATTERY_PATH
  battery.device = Device{path: battery.path}
  return battery
}

func (battery Battery) voltage() string {
  voltage := battery.device.get("voltage_now")
  // voltage = voltage[0:len(voltage) - 6]
  return string(voltage[0]) + "." + string(voltage[1])
}

// func (battery Battery) voltageNumber() int {
//   voltage := battery.device.get("voltage_now")
//   return int(voltage[0:1])
// }