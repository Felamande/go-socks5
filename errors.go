package socks5

type ListenError struct {
	Err error
}

func (l ListenError) Error() string {
	return l.Err.Error()
}
