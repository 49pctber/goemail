package main

import (
	"os"

	"github.com/49pctber/goemail"
)

func main() {
	var err error

	goemail.Configure(
		os.Getenv("EMAIL_NAME"),
		os.Getenv("EMAIL_ADDRESS"),
		os.Getenv("EMAIL_PASSWORD"),
	)

	err = goemail.SendEmail(os.Getenv("EMAIL_ADDRESS"), "Library Test", "This is a test of the library API.")
	if err != nil {
		panic(err)
	}
}
