package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/49pctber/goemail"
)

func main() {
	var to, subject, body string
	var html bool

	flag.StringVar(&to, "to", "", "recipient's email address")
	flag.StringVar(&subject, "subject", "", "subject of message")
	flag.StringVar(&body, "body", "", "the body of the message")
	flag.BoolVar(&html, "html", false, "indicate body contains html")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(), "\nBe sure to set the EMAIL_NAME, EMAIL_ADDRESS, and EMAIL_PASSWORD environment variables.\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  EMAIL_NAME is the friendly name to display to the recipient (e.g. John Doe)\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  EMAIL_ADDRESS is email address from which you will send the email\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  EMAIL_PASSWORD is used to authenticate EMAIL_ADDRESS\n")
	}

	flag.Parse()

	valid := true

	err := goemail.Configure(
		os.Getenv("EMAIL_NAME"),
		os.Getenv("EMAIL_ADDRESS"),
		os.Getenv("EMAIL_PASSWORD"),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(to) == 0 {
		fmt.Println("specify recipient using -to flag")
		valid = false
	}
	if len(subject) == 0 {
		fmt.Println("specify subject using -subject flag")
		valid = false
	}
	if len(body) == 0 {
		fmt.Println("specify body using -body flag")
		valid = false
	}

	if !valid {
		os.Exit(1)
	}

	// send email
	if html {
		err = goemail.SendHTMLEmail(to, subject, body)
	} else {
		err = goemail.SendEmail(to, subject, body)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
