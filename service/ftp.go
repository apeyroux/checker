package service

type FtpService struct {
	Service
	Host      string `json:'host'`
	Port      int    `json:'port'`
	Login     string `json:'login,omitempty'`
	Password  string `json:'password,omitempty'`
	Anonymous bool   `json:'anonymous,omitempty'`
}

func (*FtpService) State() state {
	return ALIVE
}
