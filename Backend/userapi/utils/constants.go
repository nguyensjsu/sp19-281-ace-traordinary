package utils

//MONGODB MongoDb Credentials
var MONGODB = map[string]string{
	"SERVER":                 "52.11.201.189",
	"DATABASE":               "cmpe281",
	"USERCOLLECTION":         "User",
	"REGISTRATIONCOLLECTION": "Registration",
}

//ACCESSKEYSES access key for SES
var ACCESSKEYSES = "AKIAQKC4VVTZIUJMY3GQ"

//SECRETKEYSES secret Key for SES
var SECRETKEYSES = "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"

//REGIONSES SES Region
var REGIONSES = "us-west-2"

//Email Subjects and Templates
//REGISTRATIONEMAIL mail sent before confirming user registration
var REGISTRATIONEMAIL = "Registration Email From Picassa"

//CONFIRMATIONEMAIL After Confirming user registration
var CONFIRMATIONEMAIL = "Welcome to Picassa"
