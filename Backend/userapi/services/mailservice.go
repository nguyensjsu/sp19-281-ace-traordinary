package services

//go get -u github.com/aws/aws-sdk-go
import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	//go get -u github.com/aws/aws-sdk-go
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/sp19-281-ace-traordinary/Backend/userapi/utils"
)

const (
	Sender    = "1ra4vi3@gmail.com"
	Recipient = "1ra4vi3@gmail.com"
	// The subject line for the email.
	Subject = "Confirmation email"

	// The HTML body for the email.
	HtmlBod = "<h1>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
		"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
		"<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"

	//The email body for recipients with non-HTML email clients.
	TextBody = "This email was sent with Amazon SES using the AWS SDK for Go."

	// The character encoding for the email.
	CharSet = "UTF-8"
)

//Sender string, Recipient string, Subject string, HtmlBody string
//SendConfirmationemail to send
func SendConfirmationemail() {
	// Create a new session in the us-west-2 region.
	// Replace us-west-2 with the AWS Region you're using for Amazon SES.
	var a map[string]string
	HtmlBody, err := getHTMLBody("templates/confirmation_email.gohtml", a)
	if err != nil {
		log.Println(err)
	}
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(utils.REGIONSES),
		Credentials: credentials.NewStaticCredentials(utils.ACCESSKEYSES, utils.SECRETKEYSES, ""),
	})

	// Create an SES session.
	svc := ses.New(sess)

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(Recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(HtmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(TextBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(Subject),
			},
		},
		Source: aws.String(Sender),
		// Uncomment to use a configuration set
		//ConfigurationSetName: aws.String(ConfigurationSet),
	}

	// Attempt to send the email.
	result, err := svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}

		return
	}
	fmt.Println("Email Sent to address: " + Recipient)
	fmt.Println(result)
}

func getHTMLBody(fileName string, data map[string]string) (string, error) {
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

//TestSend Non is available
func TestSend() {
	SetConfiguration("AKIAQKC4VVTZIUJMY3GQ", "1nMiMQTnaXWgl1VWMk6/R3pILWj82ZS4zWw1c2D8", "us-west-2")

	emailData := Email{
		To:      "1ra4vi3@gmail.com",
		From:    "1ra4vi3@gmail.com",
		Text:    "Hi this is the text message body",
		Subject: "Sending email from aws ses api",
		ReplyTo: "1ra4vi3@gmail.com",
	}

	resp := SendEmail(emailData)

	fmt.Println(resp)

}
