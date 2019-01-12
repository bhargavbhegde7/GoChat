package main

import (
	"github.com/bhargavbhegde7/GoChat/common"
	"github.com/fatih/color"
)

var errorChannel = make(chan common.Response)
var messageChannel = make(chan common.Response)

var controlInputChannel = make(chan common.Request)
var messageInputChannel = make(chan common.Request)

func channelSelector(){
	for{
		select{
		case errorResponse := <-errorChannel:
			go errorHandler(errorResponse)
			break
		case messageResponse := <-messageChannel:
			go messageHandler(messageResponse)
			break
		case controlRequest := <-controlInputChannel:
			go controlInputHandler(controlRequest)
			break
		case messageRequest := <-messageInputChannel:
			go messageInputHandler(messageRequest)
			break
		}
	}
}

func controlInputHandler(controlRequest common.Request){

}

func messageInputHandler(messageRequest common.Request){

}

func errorHandler(errorResponse common.Response){
	color.Red(errorResponse.Message)
}

func messageHandler(messageResponse common.Response){
	color.Yellow(messageResponse.Message)
}
