
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

    wk1.Games["New England, NE"] = Game { TeamHome : "Patriots", TeamAway : "Steelers", Location : "New England, NE" }
    wk1.Games["San Diego, SD"] = Game { TeamHome : "Chargers", TeamAway : "Lions", Location : "San Diego, SD" }
    wk1.Games["Buffalo, NY"] = Game { TeamHome : "Bills", TeamAway : "Colts", Location : "Buffalo, NY" }
    wk1.Games["New York, NY"] = Game { TeamHome : "Jets", TeamAway : "Browns", Location : "New York, NY" }
    wk1.Games["Cincinatti, OH"] = Game { TeamHome : "Bengals", TeamAway : "Raiders", Location : "Cincinatti, OH" }

    wk1.Games["Arizona, AZ"] = Game { TeamHome : "Cardinals", TeamAway : "Saints", Location : "Arizona, AZ" }
    wk1.Games["Charlotte, NC"] = Game { TeamHome : "Panthers", TeamAway : "Jaguars", Location : "Charlotte, NC" }
    wk1.Games["Denver, CO"] = Game { TeamHome : "Broncos", TeamAway : "Ravens", Location : "Denver, CO" }
    wk1.Games["Dallas, TX"] = Game { TeamHome : "Cowboys", TeamAway : "Giants", Location : "Dallas, TX" }
    wk1.Games["St. Lious, MO"] = Game { TeamHome : "Rams", TeamAway : "Seahawks", Location : "St. Lious, MO" }

    wk1.Games["Kansas City, KA"] = Game { TeamHome : "Chiefs", TeamAway : "Texans", Location : "Kansas City, KA" }
    wk1.Games["Green Bay, WI"] = Game { TeamHome : "Packers", TeamAway : "Bears", Location : "Green Bay, WI" }
    wk1.Games["Miami, FL"] = Game { TeamHome : "Dolphins", TeamAway : "Redskins", Location : "Miami, FL" }
    wk1.Games["Tennessee, TN"] = Game { TeamHome : "Titans", TeamAway : "Buccaneers", Location : "Tennessee, TN" }
    wk1.Games["San Francisco, CA"] = Game { TeamHome : "49ers", TeamAway : "Vikings", Location : "San Francisco, CA" }

    wk1.Games["Atlanta, GA"] = Game { TeamHome : "Falcons", TeamAway : "Eagles", Location : "Atlanta, GA" }

    return p
}
