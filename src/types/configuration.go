package types

type Config struct {
	AppName string `json:"appname"`
	HostName string `json:"hostname"`
	HostPort int `json:"hostport"`
}
