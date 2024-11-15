package iomanager

type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(job interface{}) error
}