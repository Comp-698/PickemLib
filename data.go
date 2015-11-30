
package PickemLib

import (
)

type Pickem struct {
    Name string
    PickemGames map[string]PickemGame
}

type PickemGame struct {
    Players map[string]Player
    Weeks map[string]Week
}

type Player struct {
    Name string
    Picks map[string]GamesPicked
}

type Week struct {
    ID string
    Ordinal int
    Games map[string]Game
}

type GamesChoices struct {
    Player Player
    Week Week
}

type GamesPicked struct {
    Game Game 
    Points int
}

type Game struct {
    TeamHome string
    TeamAway string
    Location string
}

var p *Pickem

func Data(s string) *Pickem {
    p = new(Pickem)
    p.Name = s
    return p
}

func Saturate(s string) *Pickem {
    p = new(Pickem)
    p.Name = s
    p.PickemGames = make(map[string]PickemGame)
    p.PickemGames["first"] = PickemGame { 
        Weeks : make(map[string]Week),
        Players : make(map[string]Player),
    }
    var wk1 = Week{ ID : "Week1",
        Ordinal : 1,
        Games : make(map[string]Game),
    }
    p.PickemGames["first"].Weeks[wk1.ID] = wk1 
    var playerNames = []string{ "Jon", "Tim", "JonnyT", "Huong", "Vitali" }
    var g1 = Game { TeamHome : "Patriots", TeamAway : "Bills", Location : "Foxboro, MA" }
    wk1.Games[g1.Location] = g1;
    for _, player := range playerNames {
        p.PickemGames["first"].Players[player] = Player {
            Name : player,
            Picks : make(map[string]GamesPicked),
        }
        //p.PickemGames["first"].Players[player].Picks[wk1.ID] = GamesPicked{ Game : g1, Points : 5 }
    }

    wk1.Games["Phoenix, AZ"] = Game { TeamHome : "Cardinals", TeamAway : "Vikings", Location : "Phoenix, AZ" }
    wk1.Games["Charlotte, SC"] = Game { TeamHome : "Panthers", TeamAway : "Falcons", Location : "Charlotte, SC" }

    return p
}
