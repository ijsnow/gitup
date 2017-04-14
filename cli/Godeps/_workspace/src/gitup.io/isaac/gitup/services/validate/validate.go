package validate

import (
	"regexp"

	"github.com/asaskevich/govalidator"
)

// Uname validates a username
func Uname(uname string) bool {
	return regexp.
		MustCompile(`^[a-z\d](?:[a-z\d]|-(?:[a-z\d])){0,38}$`).
		MatchString(uname)
}

// Password validates a password
func Password(pass string) bool {
	if len(pass) < 6 {
		return false
	}

	for _, s := range pass {
		if s == ' ' {
			return false
		}
	}

	return true
}

// Email validates an email
func Email(email string) bool {
	return govalidator.IsEmail(email)
}

// Host validates a host name. We actually validate by pinging the server
func Host(url string) bool {
	return url != ""
}

// RepoName validates a repo name. It is the same as uname validation
func RepoName(name string) bool {
	isEndsWithGit := regexp.
		MustCompile(".git$").
		MatchString(name)

	if isEndsWithGit {
		return false
	}

	return Uname(name)
}
