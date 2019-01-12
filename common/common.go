package common

type Response struct {
	ResTag  string `json:"restag"`
	Message string `json:"message"`
}

//const PREFIX 			= "~&#"
//const SUFFIX			= "#&~"
const GET_CLIENTS   	= "get_clients"
const LOGIN 	    	= "login"
const SIGNUP 	    	= "signup"
const SELECT_TARGET 	= "selectTarget"
const TARGET_FAIL   	= "targetFail"
const TARGET_SET    	= "targetset"
const SIGNUP_FAILURE    = "signupfailure"
const SIGNUP_SUCCESSFUL = "signupsuccess"
const ERROR    			= "error"
const INFO    			= "info"
const LOGIN_SUCCESS    	= "login_success"
const CLIENTS_LIST    	= "clients_list"
const NONE				= "NONE"
const CONTROL_MSG    	= "control"
const MESSAGE    	    = "control"