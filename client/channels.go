package main

import (
	"github.com/bhargavbhegde7/GoChat/common"
	"github.com/fatih/color"
)

var messageChannel = make(chan common.Response)

func channelSelector() {
	for {
		go messageHandler(<-messageChannel)
	}
}

func messageHandler(messageResponse common.Response) {
	message := common.AsymmetricPrivateKeyDecryption(privKey, messageResponse.Message)
	color.Yellow(messageResponse.Username + " : " + message)
}
