package models

const (
	holdTypeDesign = 0
	holdTypeStart  = 1
	holdTypeTop    = 2
	holdTypeHold   = 3
	holdTypeFeet   = 4

	holdSizeSmall  = 0
	holdSizeMedium = 1
	holdSizeBig    = 2
)

type HoldType int
type HoldSize int

type Hold struct {
	X    float64  `json:"x"`
	Y    float64  `json:"y"`
	Type HoldType `json:"type"`
	Size HoldSize `json:"size"`
}
