package types

type Player struct {
	Name string `json:"name"`
	Role string `json:"role"`
}
type Teams struct {
	Striker  string `json:"striker"`
	Defender string `json:"defender"`
}
