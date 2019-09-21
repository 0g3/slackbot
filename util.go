package slackbot

// for *bool type
func NewTrue() *bool {
	ret := true
	return &ret
}

// for *bool type
func NewFalse() *bool {
	ret := false
	return &ret
}
