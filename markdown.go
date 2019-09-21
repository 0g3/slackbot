package slackbot

import "fmt"

func AdaptBold(txt string) string {
	return fmt.Sprintf("*%s*", txt)
}

func AdaptCrossedOut(txt string) string {
	return fmt.Sprintf("~%s~", txt)
}

func AdaptLinkText(txt, link string) string {
	return fmt.Sprintf("<%s|%s>", link, txt)
}
