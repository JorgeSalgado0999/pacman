package contexts

import (
	"sync"

	"github.com/JorgeSalgado0999/MultithreadedPacman/src/constants"
	"github.com/JorgeSalgado0999/MultithreadedPacman/src/interfaces"
	"github.com/JorgeSalgado0999/MultithreadedPacman/src/modules"
	"github.com/JorgeSalgado0999/MultithreadedPacman/src/structures"
	"golang.org/x/image/font"
)

// GameContext represents the game context
type GameContext struct {
	MainPlayer  interfaces.MovableGameObject
	MazeMutex   sync.Mutex
	Maze        *structures.Maze
	GhostHome   interfaces.Location
	GhostBases  map[constants.GhostType]interfaces.Location
	SoundPlayer *modules.SoundPlayer
	Msg         *structures.MessageBroker
}

// AnchorContext represents the game context shared among screens
type AnchorContext struct {
	ChangeState  chan constants.GameState
	AssetManager *modules.AssetManager
	SoundPlayer  *modules.SoundPlayer
	GameScore    uint
	FontFace     font.Face
}
