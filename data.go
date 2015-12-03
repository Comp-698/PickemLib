
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
    Scoreboard Scoreboard
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

type Scoreboard struct {
    ID string
    Week Week
    Results map[string]int
}

type GamesChoices struct {
    Player Player
    Week []Game
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
    var week1 = "Week1"
    var first = "first"
    p = new(Pickem)
    p.Name = s
    // initialize the high level PickemGames, with a single PickemGame
    p.PickemGames = make(map[string]PickemGame)
    p.PickemGames["first"] = PickemGame { 
        Weeks : make(map[string]Week),
        Players : make(map[string]Player),
    }

    // a PickemGame gets weeks and players, let's make some
    // a week
    var wk1 = Week{ ID : week1,
        Ordinal : 1,
        Games : make(map[string]Game),
    }
    p.PickemGames[first].Weeks[wk1.ID] = wk1 
    
    // players
    var playerNames = []string{ "Jon", "Tim", "JonnyT", "Huong", "Vitali" }
    for _, player := range playerNames {
        p.PickemGames["first"].Players[player] = Player {
            Name : player,
            Picks : make(map[string]GamesPicked),
        }
    }

    wk1.Games["Pittsburg, PA"] = Game { TeamHome : "Bengals", TeamAway : "Steelers", Location : "Pittsburg, PA" }
    wk1.Games["Phoenix, AZ"] = Game { TeamHome : "Cardinals", TeamAway : "Vikings", Location : "Phoenix, AZ" }
    wk1.Games["Charlotte, SC"] = Game { TeamHome : "Panthers", TeamAway : "Falcons", Location : "Charlotte, SC" }
    wk1.Games["Washington"] = Game { TeamHome : "Bears", TeamAway : "Native American Heroes", Location : "Washington" }

    return p
}
