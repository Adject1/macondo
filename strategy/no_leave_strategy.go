package strategy

import (
	"github.com/adject1/macondo/alphabet"
	"github.com/adject1/macondo/board"
	"github.com/adject1/macondo/move"
)

// NoLeaveStrategy does not take leave into account at all.
type NoLeaveStrategy struct{}

func NewNoLeaveStrategy() *NoLeaveStrategy {

	return &NoLeaveStrategy{}
}

func (nls *NoLeaveStrategy) Equity(play *move.Move, board *board.GameBoard,
	bag *alphabet.Bag, oppRack *alphabet.Rack) float64 {
	score := play.Score()
	otherAdjustments := 0.0

	if board.IsEmpty() {
		otherAdjustments += placementAdjustment(play)
	}

	if bag.TilesRemaining() == 0 {
		otherAdjustments += endgameAdjustment(play, oppRack, bag.LetterDistribution())
	}
	return float64(score) + otherAdjustments
}

func (nls *NoLeaveStrategy) LeaveValue(leave alphabet.MachineWord) float64 {
	return 0.0
}
