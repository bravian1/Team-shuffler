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

type Table struct {
	TeamName  string `json:"teamname"`
	Played    int    `json:"played"`
	Wins     int    `json:"wins"`
	Draws    int    `json:"draws"`
	Losses   int    `json:"losses"`
	GoalsFor   int    `json:"goalsfor"`
	GoalsAgainst   int    `json:"goalsagainst"`
	Points int    `json:"points"`
}

type Fixture struct {
	Home string `json:"home"`
	Away string `json:"away"`
	Date  string `json:"date"`
}

type GameWeek struct {
	Week     int    `json:"week"`
	Matches []Fixture `json:"matches"`
}