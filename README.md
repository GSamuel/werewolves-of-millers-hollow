# Werewolves of Millers Hollow

Welcome to the Werewolves of Millers Hollow project.
With this library I make it possible for you to simulate every Werwolves of Millers Hollow game.

Contributions are welcome

## Some keywords that may or may not end up in library
* Game
* Room
* Player
* Character / Role
* Special powers
* Voting
* Village
* Elimination
* Werewolves
* Night
* Day
* Execution
* Narrator
* Game balance
* Custom Rules
* Events
* Deal Cards. (Give everyone their secret identity)
* Death, Alive, Protection


## Notes
* A typical game of Werewolves of Miller's Hollow has around 8 - 20 players
* All players have exactly 1 secret role
* All players are given their identity at the start of the game.
* All players only know their own role
* During the game players are eliminated.
* During the day players vote for a player to be eliminated.
* When a player is eliminated their role is revealed (optional)
* Players who are eliminated lose the ability to interact with the game.
* Certain roles have unique abilities that impact the course of the game.
* Werewolves win the game when all civilians are eliminated.
* Civilians (Humnans) win the game when all Werewolves are eliminated.
* In the night the werewolves wake up and together choose a victim.
* The victim of the werewolves is announced dead at the beginning of the day.
* At the beginning of the day all players who did not survive the night are announced.
* Once a game has started no player can join the game anymore. (Like almost all games)
* The narrater chooses a deck of card. (certain composition of roles)
* The deck of roles influences the balance of the game. Balance has an effect on the amount of fun players have in the game.

## Requirements
* Players can take part in a game.
* A deck of role cards contains exactly the amount of cards as the amount of players in the game.
* The cards are distributed between the players. Each player gets exactly 1 role card.
* The game is split into two phases. Night and Day.
* Some roles have unique abilities that can be used when certain events happen.
* In the night the werewolves always choose a victim.
* Roles that have a special ability in the night wake up and may use their ability.
* When the night is over all players who dies during the night are announced.
* It is possible that some roles have an unique ability when they die. (event)
* There may be a vote for Mayor.
* During the day all players vote for a player to be eliminated.
* There is room for discussing, but there may be a time limit.
* The player voted to be elimated is removed from the game.
* Everytime deaths are anounced the role of the dead player may become public. (depending on the game mode/settings)
* After the daily vote the night next night starts.
* When all players left belong to the same team the game is over and the surviving team wins the game.
* Who may and who may not vote?
* Ballot box for gathering all votes and calculating the result


## Events
* a player dies.
* a player is revealed.
* a player is chosen as (werewolf) victim.
* a player is revived/protected
* a player fell in love (cupid)


## Possible Goals
* Simulate the game so the narrater makes less mistakes
* Remove the need for a narrater.
* Make it possible to play this game online against people you do or don't know.


## Architecture
Roles have states. i.e. Witch has to know if she can still use her potions or not.
Player should not have dependencies on specific roles. (Strategy pattern for roles?)
Roles React to different events. Every role knows to which events to react.
Roles should know when an event happens on their player or on an other player. (identity)
Because roles react to different events maybe use template pattern? Are the events set in stone? else template pattern might not be a good idea.
Roles can sometimes choose players. Do they need some way to access all players in the game?
It is probably the players responsibility to choose other players. (In the real game players get asked by the narrator)