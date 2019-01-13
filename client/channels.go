package main

import (
	"github.com/bhargavbhegde7/GoChat/common"
	"github.com/fatih/color"
)

var messageChannel = make(chan common.Response)

func channelSelector(){
	for{
		go messageHandler(<-messageChannel)
	}
}

func messageHandler(messageResponse common.Response){
	color.Yellow(messageResponse.Username+" : "+messageResponse.Message)
}
