package protocol

type Response interface {
	Bytes() []byte
	Code() int
}

func (s *Sellex) AwaitFor(p interface{}) <-chan Response {
	r := CreateRequestWith(p)
	return r.Output()
}
