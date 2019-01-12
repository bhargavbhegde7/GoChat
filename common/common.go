package common

type Response struct {
	ResTag  string `json:"restag"`
	Message string `json:"message"`
	Username string `json:"username"`
}

type Request struct {
	ReqTag   string `json:"reqtag"`
	Username string `json:"username"`
	Pubkey   string `json:"pubkey"`
	Message  string `json:"message"`
}

func NewRequest(reqTag string, username string, pubkey string, message string) *Request {
	return &Request{ReqTag: reqTag, Username: username, Pubkey:pubkey, Message:message}
}

func NewResponse(resTag string, message string, username string) *Response {
	return &Response{ResTag: resTag, Message:message, Username: username}
}

const GET_CLIENTS   		= "get_clients"
const LOGIN 	    		= "login"
const SIGNUP 	    		= "signup"
const SELECT_TARGET 		= "selectTarget"
const TARGET_FAIL   		= "targetFail"
const TARGET_SET    		= "targetset"
const SIGNUP_FAILURE    	= "signupfailure"
const SIGNUP_SUCCESSFUL 	= "signupsuccess"
const ERROR    				= "error"
const LOGIN_SUCCESS    		= "login_success"
const CLIENTS_LIST    		= "clients_list"
const NONE					= "NONE"
const CLIENT_MESSAGE 		= "message"
const CONNECTION_SUCCESSFUL = "connection_successful"