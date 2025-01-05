package mail

import (
	"bytes"
	"embed"
	"text/template"
	"time"

	"gopkg.in/mail.v2"
)

//go:embed "templates"
var templateFS embed.FS

type Mailer interface {
	Send(recipient, templateFile string, data interface{}) error
}

type mailer struct {
	dialer *mail.Dialer
	sender string
}

func New(host string, port int, username, password, sender string) Mailer {
	dialer := mail.NewDialer(host, port, username, password)
	dialer.Timeout = time.Second * 5
	return mailer{
		dialer: dialer,
		sender: sender,
	}
}

// Send sends a template email with data to a recipient with configured mailer settings.
func (m mailer) Send(recipient, templateFile string, data interface{}) error {
	// Parse the required templates from the embedded file system
	tmpl, err := template.New("email").ParseFS(templateFS, "templates/"+templateFile)
	if err != nil {
		return err
	}

	// Execute the named template "subject", pass in the dynamic data and store it
	// in a bytes.Buffer variable
	subject := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return err
	}

	plainBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(plainBody, "plainBody", data)
	if err != nil {
		return err
	}

	htmlBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
	if err != nil {
		return err
	}

	// Initialize new message and set the headers and content.
	msg := mail.NewMessage()
	msg.SetHeader("To", recipient)
	msg.SetHeader("From", m.sender)
	msg.SetHeader("Subject", subject.String())
	msg.SetBody("text/plain", plainBody.String())
	// AddAlternative() should always be called after SetBody()
	msg.AddAlternative("text/html", htmlBody.String())

	err = m.dialer.DialAndSend(msg)
	// will return "dial tcp: i/o timeout" if there's a timeout
	if err != nil {
		return err
	}

	return nil
}
