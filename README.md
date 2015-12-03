# PickemLib

Use the following approach to develop the Lib

_We need to find a way to install the lib locally, without requiring a round trip
to through the github_

just like the PickemApp

fork PickemLib, and then clone it and add your remote upstream which in this case
is going to be the Comp-698/PickemLib

and then modify data.go, and use go build in your PickemLib clone to confirm it 
is valid go

since we don't know how to install it without pushing it through github, we can't 
run go test on data_test.go as a reflection of the our changes to data.go

this is awkward

but at risk of breaking the app, do so until we figure this out and then update 
data_test.go to validate our expectations of the app
