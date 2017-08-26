package main

import "strings"

type Device struct {
  path string
}

func (device Device) get(attribute string) string {
  return read(device.path + attribute)
}

func (device Device) set(attribute string, data string) {
  write(device.path + attribute, data)
}

func (device Device) bytes(file string) []byte {
  return readBytes(device.path + file)
}

type IndexedDevice struct {
  port string
  path string
}

func (device *IndexedDevice) findDeviceFromPort() {
  files := list(device.path)

  for _, file := range files {
    if strings.Contains(read(device.path + "/" + file + "/address"), device.port) {
      device.path = device.path + "/" + file + "/"
      break
    }
  }
}

func (device IndexedDevice) get(attribute string) string {
  return read(device.path + attribute)
}

func (device IndexedDevice) set(attribute string, data string) {
  write(device.path + attribute, data)
}