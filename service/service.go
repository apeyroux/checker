package service

type state int

const (
	ALIVE state = iota
	DOWN
)

func StateToString(s state) string {
	switch s {
	case 0:
		return "a live"
	case 1:
		return "down"
	default:
		return "no state"
	}
}

type Service struct {
	Name   string      `json:"name"`
	Config interface{} `json:"config"` // Nature; ApacheService, FTPService ...
}

type Serve interface {
	State() state
	Name() string
	String() string
}
