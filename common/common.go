package common

import (
	"crypto/aes"
	"crypto/cipher"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
)

//usernames in both request and response indicate who the message is from

type Response struct {
	ResTag   string `json:"restag"`
	Message  []byte `json:"message"`
	Username string `json:"username"`
}

type Request struct {
	ReqTag   string `json:"reqtag"`
	Username string `json:"username"`
	Pubkey   []byte `json:"pubkey"`
	Message  []byte `json:"message"`
}

func readTextFromFile(fileName string) []byte {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}

	return buffer
}

func InitRSA() ([]byte, []byte) {
	pub_fileName := "pub_key"
	priv_fileName := "priv_key"

	pubKey := readTextFromFile(pub_fileName)
	privKey := readTextFromFile(priv_fileName)

	fmt.Println(pubKey)
	fmt.Println(privKey)

	return pubKey, privKey
}

func GenerateRandomKey() []byte {
	return []byte(RandomString(32))
}

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func NewRequest(reqTag string, username string, pubkey []byte, message []byte) *Request {
	return &Request{ReqTag: reqTag, Username: username, Pubkey: pubkey, Message: message}
}

func NewResponse(resTag string, message []byte, username string) *Response {
	return &Response{ResTag: resTag, Message: message, Username: username}
}

func AsymmetricPublicKeyEncryption(key []byte, message []byte) []byte {

	pub_key, err := ConvertStringPubKeyToRsaKey(key)

	if err != nil {
		fmt.Println(err)
	}

	return GetEncrypted(message, pub_key)
}

func GetEncrypted(message []byte, key *rsa.PublicKey) []byte {
	label := []byte("")
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(
		hash,
		crand.Reader,
		key,
		message,
		label)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return ciphertext
}

func ConvertStringPubKeyToRsaKey(pubPEM []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pubPEM)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break // fall through
	}
	return nil, errors.New("Key type is not RSA")
}

func SymmetricEncryption(key string, message string) []byte {

	keyByteArray := []byte(key)

	c, err := aes.NewCipher(keyByteArray)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())

	_, err = io.ReadFull(crand.Reader, nonce)
	if err != nil {
		fmt.Println(err)
	}

	return gcm.Seal(nonce, nonce, []byte(message), nil)
}

func AsymmetricPrivateKeyDecryption(privKey_str []byte, cipherText []byte) string {
	priv_key, err := ConvertStringPrivKeyToRsaKey(privKey_str)

	if err != nil {
		fmt.Println(err)
	}

	return string(GetDecrypted(cipherText, priv_key))

}

func GetDecrypted(ciphertext []byte, key *rsa.PrivateKey) []byte {
	hash := sha256.New()
	label := []byte("")
	plainText, err := rsa.DecryptOAEP(
		hash,
		crand.Reader,
		key,
		ciphertext,
		label)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return plainText
}

func ConvertStringPrivKeyToRsaKey(privPEM []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privPEM)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priv, nil
}

func SymmetricDecryption(key []byte, cipherText []byte) string {
	keyByteArray := []byte(key)
	cipherTextByteArray := []byte(cipherText)

	c, err := aes.NewCipher(keyByteArray)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(cipherTextByteArray) < nonceSize {
		fmt.Println(err)
	}

	nonce, cipherTextByteArray := cipherTextByteArray[:nonceSize], cipherTextByteArray[nonceSize:]

	plainText, err := gcm.Open(nil, nonce, cipherTextByteArray, nil)
	if err != nil {
		fmt.Println(err)
	}

	return string(plainText)
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
