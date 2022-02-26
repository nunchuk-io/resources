package main

//https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/ses-example-send-email.html

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type ContactForm struct {
	Name    string `json:"name"`
	Mail    string `json:"mail"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

const (
	DevTestEmail        = "giahuy@nunchuk.io"
	NunchukSupportEmail = "support@nunchuk.io"
	CharSet             = "UTF-8"
	HtmlBody            = `<body>
<p>%s</p>
<pre>%s</pre>
</body>`
)

func ValidateEmail(email string) error {
	emailSp := strings.Split(email, "@")
	if len(emailSp) != 2 {
		return errors.New("Wrong email format")
	}
	return nil
}

func HandleContactFormSubmission(ctx context.Context, submission ContactForm) (string, error) {
	log.Println("Request", submission)

	if err := ValidateEmail(submission.Mail); err != nil {
		return "", err
	}

	bytes, err := json.Marshal(submission)
	if err != nil {
		return "", err
	}

	// Create a new session in the ap-southeast-1 region.
	// Replace ap-southeast-1 with the AWS Region you're using for Amazon SES.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1")},
	)

	// Create an SES session.
	svc := ses.New(sess)

	// Assemble the email.
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{
				aws.String(DevTestEmail),
			},
			ToAddresses: []*string{
				aws.String(NunchukSupportEmail),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(fmt.Sprintf(HtmlBody, submission.Message, string(bytes))),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(string(bytes)),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(submission.Subject),
			},
		},
		Source: aws.String(NunchukSupportEmail),
	}

	// Attempt to send the email.
	result, err := svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				log.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				log.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				log.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				log.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Println(err.Error())
		}
	}

	log.Println("Send result:", result)

	return string(bytes), nil
}

func main() {
	lambda.Start(HandleContactFormSubmission)
}
