package PickemLib_test

import (
    "testing"
     "github.com/Comp-698/PickemLib"
)

func TestData(t *testing.T) {
    var pickemName = "Comp-698"
    var d = PickemLib.Data(pickemName)
    if d.Name != pickemName  {
        t.Errorf("Error!  Data(%s) == %s, not %s", pickemName, d.Name, pickemName)
    }
}

func TestSaturate(t *testing.T) {
    var pickemName = "first"
    var pickem = PickemLib.Saturate(pickemName)
    
    var week1 = "Week1"
    var first = "first"

    // check to see if the "first" PickemGame, has had "week1" setup the way we expect
    if pickem.PickemGames[first].Weeks[week1].ID != week1 {
        t.Errorf("Error!  pickem.Weeks[\"week1\"] == %d, not %d", pickem.PickemGames[first].Weeks[week1].ID, week1)
    }

    // make sure the number of players is 5
    if len( pickem.PickemGames[first].Players ) != 5 {
        t.Errorf("Error! len( pickem.PickemGames[\"first\"].Players ) == %d, not %d",  len( pickem.PickemGames[first].Players ), 5)
    }

    // make sure player Jon is in there, and the players Picks are not initialized
    if len( pickem.PickemGames[first].Players["Jon"].Picks ) != 0 {
        t.Errorf("Error! len( pickem.PickemGames[\"first\"].Players[\"Jon\"].Picks ) ) == %d, not %d",  len( pickem.PickemGames[first].Players ), 0)
    }

    // let's also make sure the week in question has 3 games
    if len( pickem.PickemGames[first].Weeks[week1].Games ) != 3 {
        t.Errorf("Error! len( pickem.PickemGames[\"first\"].Weeks[\"week1\"].Games ) ) == %d, not %d",  len( pickem.PickemGames[first].Weeks[week1].Games ), 3)
    }
}
