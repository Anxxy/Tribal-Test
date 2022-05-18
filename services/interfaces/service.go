package interfaces

import (
	"tribal/structures"
)

type Service interface {
	GetNMusicJokes(n int) ([]*structures.Joke, error)
}
