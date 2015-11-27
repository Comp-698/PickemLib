package PickemLib

import (
    "testing"
)

func TestData(t *testing.T) {
    var pickemName = "Comp-698"
    var d = Data(pickemName)
    if d.Name != pickemName  {
        t.Errorf("Error!  Data(%s) == %s, not %s", pickemName, d.Name, pickemName)
    }
}

func TestSaturate(t *testing.T) {
    var pickemName = "Week1"
    // var pickem = Data(pickemName)
    var pickem = Saturate(pickemName)
    var one = 1
    if pickem.PickemGames["first"].Weeks["one"] != one {
        t.Errorf("Error!  pickem.Weeks[\"one\"] == %d, not %d", pickem.PickemGames["first"].Weeks["one"], one)
    }
    var wk1 = Week{ ID : "Week1"}
    if pickem.PickemGames["first"].Players["Jon"].Picks[wk1.ID].Points != 5 {
        t.Errorf("%s", pickem.PickemGames["first"].Players["Jon"].Picks[wk1.ID].Game.teamHome)
        t.Errorf("%s", pickem.PickemGames["first"].Players["Jon"].Picks[wk1.ID].Game.teamAway)
        t.Errorf("Error!  pickem.PickemGames[\"first\"].Players[\"Jon\"].Picks[wk1.ID].Points == %d, not %d", pickem.PickemGames["first"].Players["Jon"].Picks[wk1.ID].Points, 5)
    }
}
