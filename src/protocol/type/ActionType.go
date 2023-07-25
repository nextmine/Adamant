package _type

type ActionType int

const (
	StartBreak ActionType = iota
	AbortBreak
	StopBreak
	GetUpdatedBlock
	DropItem
	StartSleeping
	StopSleeping
	Respawn
	Jump
	StartSprint
	StopSprint
	StartSneak
	StopSneak
	CreativePlayerDestroyBlock
	DimensionChangeAck
	StartGlide
	StopGlide
	BuildDenied
	CrackBreak
	ChangeSkin
	SetEnchantmentSeed
	Swimming
	StopSwimming
	StartSpinAttack
	StopSpinAttack
	InteractBlock
	PredictBreak
	ContinueBreak
)
