package main

import (
	_ "embed"
	"os"

	"github.com/49pctber/goemail"
)

//go:embed message.html
var body string

func main() {
	var err error

	goemail.Configure(
		os.Getenv("EMAIL_NAME"),
		os.Getenv("EMAIL_ADDRESS"),
		os.Getenv("EMAIL_PASSWORD"),
	)

	err = goemail.SendHTMLEmail(os.Getenv("EMAIL_ADDRESS"), "HTML Test", body)
	if err != nil {
		panic(err)
	}
}
