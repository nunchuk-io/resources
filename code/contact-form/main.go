package main

//https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/ses-example-send-email.html

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/mail"
	"net/url"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/kataras/hcaptcha"
)

type ContactForm struct {
	Name             string    `json:"name"`
	Mail             string    `json:"mail"`
	Subject          string    `json:"subject"`
	Message          string    `json:"message"`
	HCaptchaResponse string    `json:"h-captcha-response,omitempty"`
	Date             time.Time `json:"date"`
}

const (
	DevTestEmail        = "giahuy@nunchuk.io"
	NunchukSupportEmail = "support@nunchuk.io"
	CharSet             = "UTF-8"
)

func HandleContactFormSubmission(ctx context.Context, submission ContactForm) (string, error) {
	log.Println("Request", submission)

	if err := ValidateEmail(submission.Mail); err != nil {
		return "", err
	}

	submission.Date = time.Now()

	hCaptchaClient := hcaptcha.New(os.Getenv("HCAPTCHA_SECRET_KEY"))
	if captchaResp := hCaptchaClient.VerifyToken(submission.HCaptchaResponse); !captchaResp.Success {
		return "", fmt.Errorf("invalid captcha: %v", captchaResp.ErrorCodes)
	}

	// Remove hcaptcha when send email
	submission.HCaptchaResponse = ""
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
					Data:    aws.String(BuildReplyEmail(&submission)),
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

func BuildReplyEmail(submission *ContactForm) string {
	const (
		HtmlBody = `<body>
<p><b>Name:</b> {{ .name }}</p>
<p><b>Email:</b> {{ .mail }}</p>
<p><b>Date:</b> {{ .date }}</p>
<p><b>Subject:</b> {{ .subject }}</p>
<p><b>Message:</b></p>
<p>{{ .message }}</p>

<br/>
<a href="{{ .reply }}">Click here to reply</a>
</body>`
	)

	subject := url.PathEscape(fmt.Sprintf("[Nunchuk Support] - %s", submission.Subject))
	body := url.PathEscape(fmt.Sprintf("Hi %s,\n", submission.Name))
	reply := fmt.Sprintf("mailto:%s?&subject=%s&body=%s", strings.Join([]string{submission.Mail, NunchukSupportEmail}, ","), subject, body)

	tmpl := template.Must(template.New("reply_email").Parse(HtmlBody))

	var buf bytes.Buffer
	err := tmpl.Execute(&buf, map[string]string{
		"name":    submission.Name,
		"mail":    submission.Mail,
		"date":    submission.Date.Format("2006-01-02 15:04:05"),
		"subject": submission.Subject,
		"message": submission.Message,
		"reply":   reply,
	})

	if err != nil {
		log.Panic(err)
	}

	return buf.String()
}

func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}
