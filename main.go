package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jawher/mow.cli"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	app := cli.App("bcrypt", "Simple bcrypt hasher")
	password := app.String(cli.StringArg{
		Name:      "PASSWORD",
		Desc:      "the password to hash",
		HideValue: true,
	})
	app.Action = func() {
		foo, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(foo))
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
