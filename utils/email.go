package utils

import (
	"bytes"
	"fmt"
	"github.com/k3a/html2text"
	"github.com/yusufbagussh/pet_hotel_backend/entity"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

type Account struct {
	Email    string
	Password string
}

const (
	CONFIG_SMTP_HOST     = "smtp.mailtrap.io"
	CONFIG_SMTP_PORT     = 2525
	CONFIG_SENDER_NAME   = "Anabul Hotel Indonesia <anabulhotelindonesia@gmail.com>"
	CONFIG_AUTH_EMAIL    = "b57d8063295c6c"
	CONFIG_AUTH_PASSWORD = "210f76284bf441"
)

func ParseTemplateAccept(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	body := new(bytes.Buffer)
	if err = t.Execute(body, data); err != nil {
		fmt.Println(err)
		return "", err
	}
	return body.String(), nil
}

func CreateEmailAccept(to string, subject string, data Account, templateFile string) error {
	result, _ := ParseTemplateAccept(templateFile, data)
	m := gomail.NewMessage()
	m.SetHeader("From", CONFIG_SENDER_NAME)
	m.SetHeader("To", to)
	// m.SetAddressHeader("Cc", "<RECIPIENT CC>", "<RECIPIENT CC NAME>")
	m.SetHeader("Subject", subject)
	//m.SetBody("text/html", "Email <b>"+to+"</b> and Password <i>"+data.Password+"</i>!")
	m.SetBody("text/html", result)
	//m.Attach(templateFile) // attach whatever you want
	//senderPort := 2525
	d := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD)
	err := d.DialAndSend(m)
	return err
}

func SendEmailAccept(to string, data Account) error {
	var err error
	template := "./templates/emailAccept.html"
	subject := "no reply"
	err = CreateEmailAccept(to, subject, data, template)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		fmt.Println("send email '" + subject + "' success")
		return nil
	}
}

func ParseTemplateReject(templateFileName string) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	body := new(bytes.Buffer)
	if err = t.Execute(body, nil); err != nil {
		fmt.Println(err)
		return "", err
	}
	return body.String(), nil
}

func CreateEmailReject(to string, subject string, templateFile string) error {
	result, _ := ParseTemplateReject(templateFile)
	m := gomail.NewMessage()
	m.SetHeader("From", CONFIG_SENDER_NAME)
	m.SetHeader("To", to)
	// m.SetAddressHeader("Cc", "<RECIPIENT CC>", "<RECIPIENT CC NAME>")
	m.SetHeader("Subject", subject)
	//m.SetBody("text/html", "DITOLAK")
	m.SetBody("text/html", result)
	//m.Attach(templateFile) // attach whatever you want
	//senderPort := 2525
	d := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD)
	err := d.DialAndSend(m)
	return err
}

func SendEmailReject(to string) error {
	fmt.Println(to)
	var err error
	template := "./templates/emailReject.html"
	subject := "no reply"
	err = CreateEmailReject(to, subject, template)
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		fmt.Println("send email '" + subject + "' success")
		return nil
	}
}

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func SendEmailStatus(user *entity.Request, data *EmailData, emailTemp string) error {
	const CONFIG_SMTP_HOST = "smtp.mailtrap.io"
	const CONFIG_SMTP_PORT = 2525
	const CONFIG_SENDER_NAME = "PT. Makmur Subur Jaya <anabulhotelindonesia@gmail.com>"
	const CONFIG_AUTH_EMAIL = "b57d8063295c6c"
	const CONFIG_AUTH_PASSWORD = "210f76284bf441"

	var body bytes.Buffer

	template, err := ParseTemplateDir("templates")
	if err != nil {
		return err
	}

	template.ExecuteTemplate(&body, emailTemp, &data)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", user.HotelEmail)
	mailer.SetAddressHeader("Cc", "tralalala@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", data.Subject)
	mailer.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	errSent := dialer.DialAndSend(mailer)
	if errSent != nil {
		return errSent
	}

	log.Println("Mail sent!")
	return nil
}

func SendEmail(user *entity.User, data *EmailData, emailTemp string) error {
	const CONFIG_SMTP_HOST = "smtp.mailtrap.io"
	const CONFIG_SMTP_PORT = 2525
	const CONFIG_SENDER_NAME = "PT. Makmur Subur Jaya <anabulhotelindonesia@gmail.com>"
	const CONFIG_AUTH_EMAIL = "b57d8063295c6c"
	const CONFIG_AUTH_PASSWORD = "210f76284bf441"

	var body bytes.Buffer

	template, err := ParseTemplateDir("templates")
	if err != nil {
		return err
	}

	template.ExecuteTemplate(&body, emailTemp, &data)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", user.Email)
	mailer.SetAddressHeader("Cc", "tralalala@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", data.Subject)
	mailer.AddAlternative("text/html", body.String())

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	errSent := dialer.DialAndSend(mailer)
	if errSent != nil {
		return errSent
	}

	log.Println("Mail sent!")
	return nil
}
