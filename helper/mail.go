package helper

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func SendMail(receiver, subject string, confirmationNr int) error {

	// Sender data.
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")

	// Receiver email address.
	toList := []string{receiver}

	// smtp server configuration.
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")

	// This is the message to send in the mail
	// msg := fmt.Sprintf("Hello, your confirmation number is :%v", confirmationNr)
	msg := fmt.Sprintf("To: %v \r\n"+
		"Subject: %v\r\n"+
		"\r\n"+
		"Dear User, \n\n your confirmation number is :%v. \r\n\n Best regards, \n fridgiGO Team", receiver, subject, confirmationNr)
	body := []byte(msg)

	// body := []byte(msg)

	auth := smtp.PlainAuth("", from, password, host)

	// SendMail uses TLS connection to send the mail
	err := smtp.SendMail(host+":"+port, auth, from, toList, body)

	// handling the errors
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println("Successfully sent mail to all user in toList")

	return nil
}
