
package PickemLib

import (
)

type Pickem struct {
    Name string
    PickemGames map[string]PickemGame
}

type PickemGame struct {
    Players map[string]Player
    Weeks map[string]int
}

type Player struct {
    Name string
    Picks map[string]GamesPicked
}

type Week struct {
    ID string

}

type GamesPicked struct {
    Game Game 
    Points int
}

type Game struct {
    teamHome string
    teamAway string
    location string
}

var p *Pickem

func Data(s string) *Pickem {
    p = new(Pickem)
    p.Name = s
    return p
}

func saturate(p *Pickem) *Pickem {
    p.PickemGames = make(map[string]PickemGame)
    p.PickemGames["first"] = PickemGame { 
        Weeks : make(map[string]int),
        Players : make(map[string]Player),
    }
    p.PickemGames["first"].Weeks["one"] = 1
    var wk1 = Week{ ID : "Week1" }
    var playerNames = []string{ "Jon", "Tim", "JonnyT", "Huong", "Vitali" }
    var g1 = Game { teamHome : "Patriots", teamAway : "Bills", location : "Foxboro, MA" }
    for _, player := range playerNames {
        p.PickemGames["first"].Players[player] = Player {
            Name : player,
            Picks : make(map[string]GamesPicked),
        }
        p.PickemGames["first"].Players[player].Picks[wk1.ID] = GamesPicked{ Game : g1, Points : 5 }
    }

return p
}
