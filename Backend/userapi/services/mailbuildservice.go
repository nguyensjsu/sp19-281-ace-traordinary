package services

import (
	"log"

	"github.com/sp19-281-ace-traordinary/Backend/userapi/models"
	"github.com/sp19-281-ace-traordinary/Backend/userapi/utils"
)

//SendRegistrationEmail service Builder
func SendRegistrationEmail(user models.Registration, host string) {
	log.Println("In Send Registration Email function")
	var data models.TemplateData
	data.Firstname = user.Firstname
	data.URL = "http://" + host + "/user?userid=" + user.Userid + "?verificationcode=" + user.Verificationcode
	var mail models.Email
	mail.From = utils.FROM
	mail.To = user.Userid
	mail.Subject = utils.REGISTRATIONEMAIL
	mail.HTMLBody, _ = getHTMLBody(utils.REGISTRATIONCONFIRMATIONTEMPLATE, data)
	SendEmail(mail, utils.REGISTRATIONCONFIRMATIONTEMPLATE, data)
}

//SendConfirmationEmail service Builder
func SendConfirmationEmail(user models.Registration) {
	var data models.TemplateData
	data.Firstname = user.Firstname
	var mail models.Email
	mail.From = utils.FROM
	mail.To = user.Userid
	mail.Subject = utils.REGISTRATIONEMAIL
	//	mail.HTMLBody, _ = getHTMLBody(utils.PAYMENTCONFIRMATIONTEMPLATE, data)
	SendEmail(mail, utils.PAYMENTCONFIRMATIONTEMPLATE, data)
}

//SendTemporeryPasswordEmail service Builder
func SendTemporeryPasswordEmail(user models.User) {
	var data models.TemplateData
	data.Firstname = user.Firstname
	data.Password = user.Password
	var mail models.Email
	mail.From = utils.FROM
	mail.To = user.Userid
	mail.Subject = utils.FORGOTPASSWORD
	Htmldata, err := getHTMLBody(utils.FORGOTPASSWORDTEMPLATE, data)
	if err != nil {

	}
	mail.HTMLBody = Htmldata
	go SendEmail(mail, utils.FORGOTPASSWORDTEMPLATE, data)
}

//DeleteUserEmail service Builder
func DeleteUserEmail(userid string) {
	//	var data models.TemplateData
	//data.Firstname = user.Firstname
	var user models.User
	user.Userid = userid
	var mail models.Email
	mail.From = utils.FROM
	mail.To = userid
	mail.Subject = utils.FORGOTPASSWORD
	//mail.HTMLBody, _ = getHTMLBody(utils.FORGOTPASSWORDTEMPLATE, data)
	//go SendEmail(mail, utils.FORGOTPASSWORDTEMPLATE, data)
}

//SendPaymentConfirmationEmail Service
func SendPaymentConfirmationEmail(picture models.Picture) {
	var data models.TemplateData
	data.Firstname = picture.UserId
	data.URL = picture.OrigUrl
	data.Price = picture.Price
	data.Verificationcode = utils.GenerateVerificationTocken()
	var mail models.Email
	mail.From = utils.FROM
	mail.To = picture.UserId
	mail.Subject = utils.PAYMENTCONFIRMATION
	mail.HTMLBody, _ = getHTMLBody(utils.PAYMENTCONFIRMATION, data)
	SendEmail(mail, utils.PAYMENTCONFIRMATIONTEMPLATE, data)
}

/**
func getHTMLBody(fileName string, indata models.TemplateData) (string, error) {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		log.Println(err)
		return "", err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, indata); err != nil {
		log.Println(err)
		return "", err
	}
	html := buffer.String()
	return html, nil
}
**/
