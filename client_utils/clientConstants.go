package client_utils

import "net"

const HELP = "~~"
const LOGIN = "~~1"
const SIGNUP = "~~2"
const USERS = "~~3"
const SELECT = "~~4"

/*
var targetpubkey []byte
var username string
var serverPubKey []byte
var serverKey []byte
var PubKey []byte
var PrivKey []byte
*/

type Client struct {
	Conn         net.Conn
	Targetpubkey []byte
	Username     string
	ServerPubKey []byte
	ServerKey    []byte
	PubKey       []byte
	PrivKey      []byte
}
