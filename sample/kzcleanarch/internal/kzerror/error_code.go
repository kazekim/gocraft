package kzerror

const (

)

type ErrorCode int

func (enum ErrorCode) Int() int {
	return int(enum)
}
