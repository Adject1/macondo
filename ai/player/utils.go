package player

import (
	"github.com/adject1/macondo/game"
	"github.com/adject1/macondo/move"
	"github.com/adject1/macondo/movegen"
)

// GenBestStaticTurn is a useful utility function for sims and autoplaying.
func GenBestStaticTurn(game *game.Game, movegen movegen.MoveGenerator,
	aiplayer AIPlayer, playerIdx int) *move.Move {

	opp := (playerIdx + 1) % game.NumPlayers()

	// Add an exchange only if there are 7 or more tiles in the bag.
	movegen.GenAll(game.RackFor(playerIdx), game.Bag().TilesRemaining() >= 7)
	aiplayer.AssignEquity(movegen.Plays(), game.Board(), game.Bag(),
		game.RackFor(opp))
	return aiplayer.BestPlay(movegen.Plays())
}
