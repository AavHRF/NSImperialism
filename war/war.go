package war

import "github.com/brickman1444/NSImperialism/nationstates_api"

type War struct {
	Attacker *nationstates_api.Nation
	Defender *nationstates_api.Nation
	Score    int // 100 is attacker wins, -100 is defender wins
	Name     string
}

func (war *War) ScoreChangePerYear() int {
	return (100 - war.Attacker.GetDefenseForces()) - (100 - war.Defender.GetDefenseForces())
}
