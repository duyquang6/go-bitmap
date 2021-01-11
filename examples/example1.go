package main

import "bitmap"

const (
	canCall = iota
	canConnectInternet
	canMessage
)

func main() {

	myphoneCapability := bitmap.NewBitmap(100)

	myphoneCapability.Set(canConnectInternet, true)
	myphoneCapability.Set(canMessage, true)

	if ok, _ := myphoneCapability.Get(canCall); ok {
		println("My phone can call")
	}

	if ok, _ := myphoneCapability.Get(canConnectInternet); ok {
		println("My phone can connect internet")
	}

	if ok, _ := myphoneCapability.Get(canMessage); ok {
		println("My phone can message")
	}
}
