package core

type Response interface {
	Bytes() []byte
	Code() int
}
