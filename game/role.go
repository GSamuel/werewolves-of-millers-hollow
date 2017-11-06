package game

import ()

type Role interface {
	Name() string         //Role name or player name?
	ID() int              //Should Role really have identity? Player has identity
	SetID(int)            //Same
	Alive() bool          //Player is alive or dead, not the role
	SetAlive(bool)        // same
	Die(*Game, bool)      //Player dies, the role does not die.
	Alliance() Alliance   //Should be bound to player. Alliance can possibly change
	SetAlliance(Alliance) //same
	EventListener         //belongs here probably.
}
