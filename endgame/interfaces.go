package endgame

import (
	"github.com/adject1/macondo/game"
	"github.com/adject1/macondo/move"
	"github.com/adject1/macondo/movegen"
)

// Solver is an interface for an endgame solver. The scores don't matter, the
// solver maximizes spread.
type Solver interface {
	Init(movegen movegen.MoveGenerator, game *game.Game)
	Solve(plies int) (float32, *move.Move)
}
