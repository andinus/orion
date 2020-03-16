package main

import (
	"os"
	"time"

	"framagit.org/andinus/orion/hibp"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func main() {
	var pass string

	prompt := &survey.Password{
		Message: "Password:",
		Help:    "Enter password to be checked against HIBP's Database",
	}
	err := survey.AskOne(prompt, &pass, survey.WithValidator(survey.Required))
	if err == terminal.InterruptErr {
		color.Yellow("Interrupt Received")
		os.Exit(0)
	} else if err != nil {
		panic(err)
	}

	s := spinner.New(spinner.CharSets[12], 32*time.Millisecond)
	s.Start()
	s.Color("cyan")

	// get password hash
	hsh := hibp.GetHsh(pass)

	// get list of pwned passwords
	list, err := hibp.GetPwned(hsh)
	if err != nil {
		color.Yellow(err.Error())
		os.Exit(1)
	}

	// check if pass is pwned
	pwn, fq := hibp.ChkPwn(list, hsh)
	s.Stop()

	if pwn {
		color.New(color.FgRed).Add(color.Bold).Println("\nPwned!")
		color.Yellow("This password has been seen %s times before.", fq)
		return
	}

	color.Green("\nPassword wasn't found in Have I Been Pwned's Database")
}
