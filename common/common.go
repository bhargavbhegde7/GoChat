package common

import (
	"crypto/rand"
	"math/rand"
)

//usernames in both request and response indicate who the message is from

type Response struct {
	ResTag   string `json:"restag"`
	Message  string `json:"message"`
	Username string `json:"username"`
}

type Request struct {
	ReqTag   string `json:"reqtag"`
	Username string `json:"username"`
	Pubkey   string `json:"pubkey"`
	Message  string `json:"message"`
}

func GenerateRandomKey() string {
	return RandomString(32)
}

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func NewRequest(reqTag string, username string, pubkey string, message string) *Request {
	return &Request{ReqTag: reqTag, Username: username, Pubkey: pubkey, Message: message}
}

func NewResponse(resTag string, message string, username string) *Response {
	return &Response{ResTag: resTag, Message: message, Username: username}
}

func AsymmetricEncryption(key string, message string) string {
	return message + " encrypted with " + key
}

func SymmetricEncryption(key string, message string) string {
	return message + " encrypted with " + key
}

func AsymmetricDecryption(key string, message string) string {
	return "decrypted : " + message
}

func SymmetricDecryption(key string, message string) string {
	return "decrypted : " + message
}

const GET_CLIENTS = "get_clients"
const LOGIN = "login"
const SIGNUP = "signup"
const SELECT_TARGET = "selectTarget"
const TARGET_FAIL = "targetFail"
const TARGET_SET = "targetset"
const SIGNUP_FAILURE = "signupfailure"
const SIGNUP_SUCCESSFUL = "signupsuccess"
const ERROR = "error"
const LOGIN_SUCCESS = "login_success"
const CLIENTS_LIST = "clients_list"
const NONE = "NONE"
const CLIENT_MESSAGE = "message"
const CONNECTION_SUCCESSFUL = "connection_successful"
const SERVER_KEY_EXCHANGE = "server_key_exchange"
const SERVER_KEY_ACK = "server_key_ack"
const TARGET_NOT_SET = "target_not_set"
