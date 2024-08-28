package goemail

import (
	"fmt"
	"net/smtp"
	"strings"
)

var configured bool = false

var display_name string // display name for from_address
var from_address string // the address from which the email will originate
var password string     // password for authentication

var smtpHost string
var smtpPort string

func Configure(name, address, pswd string) error {

	configured = false

	display_name = name
	from_address = address
	password = pswd

	// validate sender's email address
	if len(from_address) == 0 {
		return fmt.Errorf("from address not set")
	}

	// ensure email address's domain is supported
	at := strings.LastIndex(from_address, "@")

	if at == -1 {
		return fmt.Errorf("%s is an invalid email address", from_address)
	}

	domain := from_address[at+1:]
	switch domain {
	case "gmail.com":
		smtpHost = "smtp.gmail.com"
		smtpPort = "587"
	default:
		return fmt.Errorf("%s is an unsupported domain", domain)
	}

	// validate display name
	if len(display_name) == 0 {
		return fmt.Errorf("name not set")
	}

	// validate password
	if len(password) == 0 {
		return fmt.Errorf("password not set")
	}

	configured = true

	return nil
}

/*
Sends an email from your email address.

If you'd like to send an SMS, you first need to get the number's gateway address.
This is done by looking up the appropriate gateway for the number's carrier.
e.g. <phone number>@vtext.com
e.g. <phone number>@tmomail.net
*/
func SendEmail(to, subject, body string) error {

	if !configured {
		return fmt.Errorf("not configured")
	}

	header := fmt.Sprintf("From: %s <%s>\nTo: <%s>\nSubject: %s\n", display_name, from_address, to, subject)
	message := []byte(header + "\n" + body)

	// send it, bro
	auth := smtp.PlainAuth("", from_address, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from_address, []string{to}, message)
	if err != nil {
		return err
	}
	return nil
}
