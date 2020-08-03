package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"sync"
)

var auth smtp.Auth

func main() {
	auth = smtp.PlainAuth("", "dicomiotservice@gmail.com", "uuijiwcohupaskza", "smtp.gmail.com")
	templateData := struct {
		Name string
		URL  string
	}{
		Name: "DICOMITOSERVICE",
		URL:  "http://geektrust.in",
	}

	r := NewRequest([]string{"namhoangle1996@gmail.com"}, "Hello Nam!", "Hello, World!")
	if err := r.ParseTemplate("mail.html", templateData); err == nil {
		r.wg.Add(1)
		go func() {
			r.mutex.Lock()
			defer r.wg.Done()

			ok, _ := r.SendEmail()
			if ok  {
				fmt.Print("send success")
			}
			defer r.mutex.Unlock()

		}()
		fmt.Print("okokoko")

	}

}

//Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
	wg sync.WaitGroup
	mutex	sync.Mutex
}

func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) SendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, auth, "dicomiotservice@gmail.com", r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}
