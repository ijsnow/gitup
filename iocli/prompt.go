package iocli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"gitup.io/isaac/gitup/services/validate"

	"github.com/chzyer/readline"
)

// PromptInput is the data type returned with input
type PromptInput struct {
	Response string
}

// IsYes is a helper to see if the user entered yes
func (p PromptInput) IsYes() bool {
	return p.Response == "y" || p.Response == "Y"
}

// IsNo is a helper to see if the user entered no
func (p PromptInput) IsNo() bool {
	return p.Response == "n" || p.Response == "N"
}

// PromptRune is used to prompt the user to enter a single character
func PromptRune(format string, a ...interface{}) PromptInput {

	reader := bufio.NewReader(os.Stdin)
	printPrompt(format, a...)

	r, _, _ := reader.ReadRune()

	return PromptInput{
		Response: string(r),
	}
}

// PromptString is used to prompt the user to enter a string
func PromptString(format string, a ...interface{}) PromptInput {
	pformat := fmt.Sprintf(format, a...)
	cfg := &readline.Config{
		Prompt:          fmt.Sprintf("%s %s: ", promptColor("?>"), pformat),
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	}

	rl, err := readline.NewEx(cfg)
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	r, err := rl.Readline()
	if err == readline.ErrInterrupt {
		os.Exit(0)
	}

	return PromptInput{
		Response: strings.TrimRight(r, "\n"),
	}
}

func getListener(rl *readline.Instance, valid func(string) bool, format string) readline.Listener {
	return readline.FuncListener(func(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool) {
		var successChar string
		if valid(string(line)) {
			successChar = successColor(pass)
		} else if len(line) > 0 {
			successChar = errorColor(fail)
		} else {
			successChar = promptColor(ambiguous)
		}

		rl.SetPrompt(fmt.Sprintf("%s %s (%s): ", promptColor("?>"), format, successChar))

		rl.Refresh()

		return nil, 0, false
	})
}

func promptWithValidation(valid func(string) bool, printHelp func(), format string, a ...interface{}) PromptInput {
	pformat := fmt.Sprintf(format, a...)
	cfg := &readline.Config{
		Prompt:          fmt.Sprintf("%s %s: ", promptColor("?>"), pformat),
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	}

	rl, err := readline.NewEx(cfg)
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	cfg.Listener = getListener(rl, valid, format)
	rl.SetConfig(cfg)

	var line string
	isFirst := true

	for !valid(line) {
		if !isFirst {
			printHelp()
		} else {
			isFirst = false
		}

		line, err = rl.Readline()
		if err == readline.ErrInterrupt {
			os.Exit(0)
		} else if err == io.EOF {
			break
		}

		var successChar string
		if valid(line) {
			successChar = successColor(pass)
		} else {
			successChar = errorColor(fail)
		}

		rl.SetPrompt(fmt.Sprintf("%s %s (%s): ", promptColor("?>"), pformat, successChar))
	}

	return PromptInput{
		Response: line,
	}
}

func printUnameHelp() {
	Error("Oops! The username you entered was invalid.")
	Info("username may only contain alphanumeric characters or hyphens")
	Info("username cannot have multiple consecutive hyphens")
	Info("username cannot begin or end with a hyphen")
	Info("all letters must be lowercase")
	Info("maximum is 39 characters")
}

// PromptUname reads input and validates a user name
func PromptUname(format string, a ...interface{}) PromptInput {
	return promptWithValidation(validate.Uname, printUnameHelp, format, a...)
}

func printRepoNameHelp() {
	Error("Oops! The repo name you entered was invalid.")
	Info("repo name cannot end with `.git`")
	Info("repo name may only contain alphanumeric characters or hyphens")
	Info("repo name cannot have multiple consecutive hyphens")
	Info("repo name cannot begin or end with a hyphen")
	Info("all letters must be lowercase")
	Info("maximum is 39 characters")
}

// PromptRepoName reads input and validates a reponame
func PromptRepoName(format string, a ...interface{}) PromptInput {
	return promptWithValidation(validate.RepoName, printRepoNameHelp, format, a...)
}

func printHostHelp() {
	Error("Oops! The host you entered is invalid.")
	Info("The host name must either be a domain name(ex: git.mydomain.com)")
}

// PromptHost prompts for an email
func PromptHost(format string, a ...interface{}) PromptInput {
	return promptWithValidation(validate.Host, printHostHelp, format, a...)
}

func printEmailHelp() {
	Error("Oops! The email you entered was invalid.")
	Info("Enter a valid email address in the following format:")
	Info("  you@example.com")
}

// PromptEmail prompts for an email
func PromptEmail(format string, a ...interface{}) PromptInput {
	return promptWithValidation(validate.Email, printEmailHelp, format, a...)
}

func printPasswordHelp() {
	Error("Invalid password!")
	Info("Password must be at least 6 characters may not contain any spaces")
}

// PromptPassword prompts the user to enter a password
func PromptPassword(format string, a ...interface{}) PromptInput {
	pformat := fmt.Sprintf(format, a...)
	rl, err := readline.NewEx(&readline.Config{
		Prompt:          fmt.Sprintf("%s %s: ", promptColor("?>"), pformat),
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	cfg := rl.GenPasswordConfig()
	cfg.SetListener(func(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool) {
		var successChar string
		if validate.Password(string(line)) {
			successChar = successColor(pass)
		} else if len(line) > 0 {
			successChar = errorColor(fail)
		} else {
			successChar = promptColor(ambiguous)
		}

		rl.SetPrompt(fmt.Sprintf("%s %s (%s): ", promptColor("?>"), format, successChar))

		rl.Refresh()

		return nil, 0, false
	})

	var line []byte
	var lineStr string

	isFirst := true

	for !validate.Password(lineStr) {
		if !isFirst {
			printPasswordHelp()
		} else {
			isFirst = false
		}

		line, err = rl.ReadPasswordWithConfig(cfg)
		lineStr = string(line)
		if err == readline.ErrInterrupt {
			os.Exit(0)
		} else if err == io.EOF {
			break
		}

		var successChar string
		if validate.Password(lineStr) {
			successChar = successColor(pass)
		} else {
			successChar = errorColor(fail)
		}

		rl.SetPrompt(fmt.Sprintf("%s %s (%s): ", promptColor("?>"), pformat, successChar))
	}

	return PromptInput{
		Response: lineStr,
	}
}
