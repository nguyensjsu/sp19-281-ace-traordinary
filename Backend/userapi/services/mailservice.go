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
	"github.com/sp19-281-ace-traordinary/Backend/userapi/models"
	"github.com/sp19-281-ace-traordinary/Backend/userapi/utils"
)

const (
	CharSet = "UTF-8"
)

//SendEmail to send
func SendEmail(mail models.Email, fileName string, indata models.TemplateData) {
	// Create a new session in the us-west-2 region.
	// Replace us-west-2 with the AWS Region you're using for Amazon SES.

	HtmlBody, err := getHTMLBody(fileName, indata)
	//	log.Println(HtmlBod)
	//	mail.HTMLBody = HtmlBod
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
				aws.String(mail.To),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(utils.CHARSET),
					Data:    aws.String(HtmlBody),
				},
				Text: &ses.Content{
					Charset: aws.String(utils.CHARSET),
					Data:    aws.String(mail.TextBody),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(utils.CHARSET),
				Data:    aws.String(mail.Subject),
			},
		},
		Source: aws.String(mail.From),
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
	fmt.Println("Email Sent to address: " + mail.To)
	fmt.Println(result)
}

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
