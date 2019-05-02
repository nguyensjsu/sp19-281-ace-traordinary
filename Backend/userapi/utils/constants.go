package utils

//MONGODB MongoDb Credentials
var MONGODB = map[string]string{
	"SERVER":                 "52.11.201.189",
	"DATABASE":               "cmpe281",
	"USERCOLLECTION":         "User",
	"REGISTRATIONCOLLECTION": "Registration",
}

//CHARSET for mail
const CHARSET = "UTF-8"

//ACCESSKEYSES access key for SES
var ACCESSKEYSES = "AKIAQKC4VVTZIUJMY3GQ"

//SECRETKEYSES secret Key for SES
var SECRETKEYSES = "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"

//FROM email ID
var FROM = "support@picassa.awsapps.com"

//REGIONSES SES Region
var REGIONSES = "us-west-2"

//REGISTRATIONEMAIL mail sent before confirming user registration
var REGISTRATIONEMAIL = "Registration Email From Picassa"

//REGISTRATIONCONFIRMATIONTEMPLATE  After Successfull Payment
var REGISTRATIONCONFIRMATIONTEMPLATE = "src/github.com/sp19-281-ace-traordinary/Backend/userapi/templates/registration_email.gohtml"

//CONFIRMATIONEMAIL After Confirming user registration
var CONFIRMATIONEMAIL = "Welcome to Picassa"

//CONFIRMATIONTEMPLATE After Confirming user registration
var CONFIRMATIONTEMPLATE = "src/github.com/sp19-281-ace-traordinary/Backend/userapi/templates/confirm_registation.gohtml"

//PAYMENTCONFIRMATION After SuccesfulPayment
var PAYMENTCONFIRMATION = "Payment Confirmation Picassa"

//PAYMENTCONFIRMATIONTEMPLATE  After Successfull Payment
var PAYMENTCONFIRMATIONTEMPLATE = "src/github.com/sp19-281-ace-traordinary/Backend/userapi/templates/paymentconfirmation_email.gohtml"

// FORGOTPASSWORD subject for forgot password
var FORGOTPASSWORD = "Please Find your Temporary Password"

//FORGOTPASSWORDTEMPLATE when user forgets password
var FORGOTPASSWORDTEMPLATE = "src/github.com/sp19-281-ace-traordinary/Backend/userapi/templates/forgot_password.gohtml"
