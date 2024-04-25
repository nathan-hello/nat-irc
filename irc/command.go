package irc

type Command struct {
	Prefix       string
	Args         struct{}
	FullCmd      string
	Validator    func(string) bool
	ErrorHandler func(string) any
	Callback     func(string) any
}
