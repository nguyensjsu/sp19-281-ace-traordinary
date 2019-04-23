package services

import (
	"bytes"
	"html/template"

	"github.com/sp19-281-ace-traordinary/Backend/userapi/models"
	"github.com/sp19-281-ace-traordinary/Backend/userapi/utils"
)

//SendRegistrationEmail service Builder
func SendRegistrationEmail(user models.Registration, host string) {
	var data models.TemplateData
	data.Firstname = user.Firstname
	data.URL = "http://" + host + "/user?userid=" + user.Userid + "?verificationcode=" + user.Verificationcode
	var mail models.Email
	mail.From = utils.FROM
	mail.To = user.Userid
	mail.Subject = utils.REGISTRATIONEMAIL
	mail.HTMLBody, _ = getHTMLBody(utils.REGISTRATIONCONFIRMATIONTEMPLATE, data)
	go SendEmail(mail)
}

//SendConfirmationEmail service Builder
func SendConfirmationEmail(user models.Registration) {
	var data models.TemplateData
	data.Firstname = user.Firstname
	var mail models.Email
	mail.From = utils.FROM
	mail.To = user.Userid
	mail.Subject = utils.REGISTRATIONEMAIL
	mail.HTMLBody, _ = getHTMLBody(utils.PAYMENTCONFIRMATIONTEMPLATE, data)
	go SendEmail(mail)
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
	mail.HTMLBody, _ = getHTMLBody(utils.FORGOTPASSWORDTEMPLATE, data)
	go SendEmail(mail)
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
	go SendEmail(mail)
}

//SendPaymentConfirmationEmail Service
func SendPaymentConfirmationEmail(indata map[string]string) {
	var data models.TemplateData
	data.Firstname = indata["Firstname"]
	var mail models.Email
	mail.From = utils.FROM
	mail.To = indata["toAddress"]
	mail.Subject = utils.REGISTRATIONEMAIL
	mail.HTMLBody, _ = getHTMLBody(utils.PAYMENTCONFIRMATIONTEMPLATE, data)
	go SendEmail(mail)
}

func getHTMLBody(fileName string, indata models.TemplateData) (string, error) {
	t, err := template.ParseFiles(fileName)
	if err != nil {
		return "", err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, indata); err != nil {
		return "", err
	}
	html := buffer.String()
	return html, nil
}
