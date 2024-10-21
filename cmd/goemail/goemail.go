package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/49pctber/goemail"
)

func main() {
	err := goemail.Configure(
		os.Getenv("EMAIL_NAME"),
		os.Getenv("EMAIL_ADDRESS"),
		os.Getenv("EMAIL_PASSWORD"),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var to, subject, body string
	var html bool

	flag.StringVar(&to, "to", "", "recipient's email address")
	flag.StringVar(&subject, "subject", "", "subject of message")
	flag.StringVar(&body, "body", "", "the body of the message")
	flag.BoolVar(&html, "html", false, "indicate body contains html")

	flag.Parse()

	valid := true

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
