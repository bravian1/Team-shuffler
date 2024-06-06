package types

type Player struct {
	Name string `json:"name"`
	Role string `json:"role"`
}
type Teams struct {
	Name  string `json:"name"`
	Striker  string `json:"striker"`
	Defender string `json:"defender"`
}
