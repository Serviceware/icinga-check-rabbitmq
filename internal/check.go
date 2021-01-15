package internal

const (
	OK       = 0
	WARNING  = 1
	CRITICAL = 2
	UNKNOWN  = 3
)

type Check interface {
	DoCheck() int
}
