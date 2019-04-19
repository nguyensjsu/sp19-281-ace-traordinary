package services

import (
	"bytes"
	"html/template"

	"github.com/sp19-281-ace-traordinary/Backend/userapi/models"
	"github.com/sp19-281-ace-traordinary/Backend/userapi/utils"
)

type Data interface {
}

//SendRegistrationEmail service Builder
func SendRegistrationEmail(user models.Registration) {
	var mail models.Email
	mail.From = utils.FROM
	mail.To = user.Userid
	mail.Subject = utils.REGISTRATIONEMAIL
	mail.HTMLBody, _ = getHTMLBody(utils.PAYMENTCONFIRMATIONTEMPLATE, user)
	go SendEmail(mail)
}

//SendConfirmationEmail service Builder
func SendConfirmationEmail(toAddress string, data map[string]string) {
	var mail models.Email
	mail.From = utils.FROM
	mail.To = toAddress
	mail.Subject = utils.REGISTRATIONEMAIL
	mail.HTMLBody, _ = getHTMLBody(utils.PAYMENTCONFIRMATIONTEMPLATE, data)
	go SendEmail(mail)
}

//SendTemporeryPasswordEmail service Builder
func SendTemporeryPasswordEmail(toAddress string, data map[string]string) {
	var mail models.Email
	mail.From = utils.FROM
	mail.To = toAddress
	mail.Subject = utils.REGISTRATIONEMAIL
	mail.HTMLBody, _ = getHTMLBody(utils.PAYMENTCONFIRMATIONTEMPLATE, data)
	go SendEmail(mail)
}

//SendPaymentConfirmationEmail Service
func SendPaymentConfirmationEmail(toAddress string, data map[string]string) {
	var mail models.Email
	mail.From = utils.FROM
	mail.To = toAddress
	mail.Subject = utils.REGISTRATIONEMAIL
	mail.HTMLBody, _ = getHTMLBody(utils.PAYMENTCONFIRMATIONTEMPLATE, data)
	go SendEmail(mail)
}

func getHTMLBody(fileName string, data Data) (string, error) {

	t, err := template.ParseFiles(fileName)
	if err != nil {
		return "", err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return "", err
	}
	html := buffer.String()
	return html, nil
}
