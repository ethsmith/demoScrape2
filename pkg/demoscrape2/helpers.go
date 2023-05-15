package demoscrape2

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"

	dem "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
)

type Dictionary map[string]interface{}

func isDuringExpectedRound(game *Game, p dem.Parser) bool {
	isPreWinCon := int(game.PotentialRound.RoundNum) == p.GameState().TotalRoundsPlayed()+1
	isAfterWinCon := int(game.PotentialRound.RoundNum) == p.GameState().TotalRoundsPlayed() && game.Flags.PostWinCon
	return isPreWinCon || isAfterWinCon
}

func printPlayers(game *Game, team *common.TeamState) {
	for _, teamMember := range team.Members() {
		if teamMember.IsAlive() && game.PotentialRound.PlayerStats[teamMember.SteamID64].Health > 0 {
			log.Debug(teamMember, " is alive on team", team)
		} else {
			log.Debug(teamMember, " is dead on team", team)
		}
	}
}

func validateTeamName(game *Game, teamName string, teamNum common.Team) string {
	if teamName != "" {
		name := ""
		if strings.HasPrefix(teamName, "[") {
			if len(teamName) == 31 {
				//name here will be truncated
				name = strings.Split(teamName, "] ")[1]
				for _, team := range game.Teams {
					if strings.Contains(team.Name, name) {
						return team.Name
					}
				}
				fmt.Print("OH NOEY")
				return name
			} else {
				name = strings.Split(teamName, "] ")[1]
				return name
			}
		} else {
			return teamName
		}
	} else {
		//this demo no have team names, so we are big fucked
		//we are hardcoding during what rounds each team will have what side
		round := game.PotentialRound.RoundNum
		swap := false
		if round >= 16 && round <= 33 {
			swap = true
		} else if round >= 34 {
			//we are now in OT hell :)
			if (round-34)/6%2 != 0 {
				swap = true
			}
		}
		if !swap {
			if teamNum == 2 {
				return "StartedT"
			} else if teamNum == 3 {
				return "StartedCT"
			}
		} else {
			if teamNum == 2 {
				return "StartedCT"
			} else if teamNum == 3 {
				return "StartedT"
			}
			return "SPECs"
		}
		return "SPECs"
	}
}

func calculateTeamEquipmentValue(game *Game, team *common.TeamState) int {
	equipment := 0
	for _, teamMember := range team.Members() {
		if teamMember.IsAlive() && game.PotentialRound.PlayerStats[teamMember.SteamID64].Health > 0 {
			equipment += teamMember.EquipmentValueCurrent()
		}
	}
	return equipment
}

// works for grenades, needs to be modified for other types
func calculateTeamEquipmentNum(game *Game, team *common.TeamState, equipmentENUM int) int {
	equipment := 0
	for _, teamMember := range team.Members() {
		if teamMember.IsAlive() && game.PotentialRound.PlayerStats[teamMember.SteamID64].Health > 0 {
			//log.Debug(teamMember.Inventory)
			//log.Debug(teamMember.Weapons())
			//log.Debug(teamMember.AmmoLeft)
			//gren := teamMember.Inventory[equipmentENUM]
			equipment += teamMember.AmmoLeft[equipmentENUM]
		}
	}
	return equipment
}

func closestCTDisttoBomb(game *Game, team *common.TeamState, bomb *common.Bomb) int {
	var distance = 999999
	for _, teamMember := range team.Members() {
		if teamMember.IsAlive() && game.PotentialRound.PlayerStats[teamMember.SteamID64].Health > 0 {
			if bomb.Position().Distance(teamMember.Position()) < float64(distance) {
				distance = int(bomb.Position().Distance(teamMember.Position()))
			}
		}
	}
	return distance
}

func numOfKits(game *Game, team *common.TeamState) int {
	kits := 0
	for _, teamMember := range team.Members() {
		if teamMember.IsAlive() && game.PotentialRound.PlayerStats[teamMember.SteamID64].Health > 0 {
			if teamMember.HasDefuseKit() {
				kits += 1
			}
		}
	}
	return kits
}

func playersWithArmor(game *Game, team *common.TeamState) int {
	armor := 0
	for _, teamMember := range team.Members() {
		if teamMember.IsAlive() && game.PotentialRound.PlayerStats[teamMember.SteamID64].Health > 0 {
			if teamMember.Armor() > 0 {
				armor += 1
			}
		}
	}
	return armor
}

func getPlayerAPIDict(side string, player *playerStats) Dictionary {

	return Dictionary{
		"playerSteamId": strconv.FormatUint(player.SteamID, 10),
		"side":          side,
		"teamName":      player.TeamClanName,
		"adp":           player.DeathPlacement,
		"adr":           player.Adr,
		"assists":       player.Assists,
		"atd":           player.Atd,
		"awpK":          player.AwpKills,
		"damageDealt":   player.Damage,
		"damageTaken":   player.DamageTaken,
		"deaths":        player.Deaths,
		"eac":           player.Eac,
		"ef":            player.Ef,
		"eft":           player.EnemyFlashTime,
		"fAss":          player.FAss,
		"fDeaths":       player.Ol,
		"fireDamage":    player.InfernoDmg,
		"fires":         player.FiresThrown,
		"fiveK":         player.FiveK,
		"fourK":         player.FourK,
		"threeK":        player.ThreeK,
		"twoK":          player.TwoK,
		"fKills":        player.Ok,
		"flashes":       player.FlashThrown,
		"hs":            player.Hs,
		"impact":        player.ImpactRating,
		"iwr":           player.Iiwr,
		"jumps":         0,
		"kast":          player.Kast,
		"kills":         player.Kills,
		"kpa":           player.KillPointAvg,
		"lurks":         player.LurkRounds,
		"mip":           player.Mip,
		"nadeDamage":    player.NadeDmg,
		"nades":         player.NadesThrown,
		"oneVFive":      player.Cl_5,
		"oneVFour":      player.Cl_4,
		"oneVThree":     player.Cl_3,
		"oneVTwo":       player.Cl_2,
		"oneVOne":       player.Cl_1,
		"ra":            player.RA,
		"rating":        player.Rating,
		"rf":            player.RF,
		"rounds":        player.Rounds,
		"rwk":           player.Rwk,
		"rws":           player.Rws,
		"saves":         player.Saves,
		"smokes":        player.SmokeThrown,
		"suppR":         player.SuppRounds,
		"suppX":         player.SuppDamage,
		"traded":        player.Traded,
		"trades":        player.Trades,
		"ud":            player.UtilDmg,
		"util":          player.UtilThrown,
		"wlp":           player.Wlp,
	}
}
