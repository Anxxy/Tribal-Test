package services

import (
	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"
	"tribal/structures"
)

// jmonterroso@tribalworldwide.gt
// jgiron@tribalworldwide.gt
func (s service) GetNMusicJokes(n int) ([]*structures.Joke, error) {
	const url = "https://api.chucknorris.io/jokes/random"

	jokesRegister := make(map[string]bool)

	var jokes []*structures.Joke

	for len(jokes) < 15 {
		resp, err := s.RestClient.R().
			SetResult(&structures.Joke{}).
			SetQueryParams(map[string]string{
				"category": "music",
			}).
			Get(url)
		if err != nil {
			return nil, err
		}

		joke := resp.Result().(*structures.Joke)

		log.Debugln(joke.Id)
		if ok := jokesRegister[joke.Id]; !ok {
			jokes = append(jokes, joke)
			jokesRegister[joke.Id] = true
		}

		log.Debugln(spew.Sdump(jokesRegister))
	}

	return jokes, nil
}
