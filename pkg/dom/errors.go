package dom

type stringError string

func (err stringError) Error() string {
	return string(err)
}

func newError(s string) error {
	return stringError(s)
}
