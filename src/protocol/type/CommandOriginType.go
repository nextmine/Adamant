package _type

type CommandOriginType int

const (
	Player CommandOriginType = iota
	Block
	MinecartBlock
	DevConsole
	Test
	AutomationPlayer
	ClientAutomation
	DedicatedServer
	Entity
	Virtual
	GameArgument
	EntityServer
	Precompiled
	GameDirectorEntityServer
	Script
)
