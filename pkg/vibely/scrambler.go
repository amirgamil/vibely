package vibely

import (
	"fmt"
	"math/rand"
	"strings"
)

type Word struct {
	Hidden bool   `json:"hidden"`
	Word   string `json:"word"`
}

type Payload struct {
	Song []Word `json:"data"`
}

func scramble(song string) Payload {
	song = strings.TrimSpace(song)
	words := []Word{}
	for _, sentence := range strings.Split(song, "\n") {
		if sentence == "" {
			words = append(words, Word{Hidden: false, Word: "\n"})
			continue
		}
		formatSentence := strings.Fields(sentence)
		i := 0
		for i < len(formatSentence) {
			if rand.Intn(4) == 2 && i+3 < len(formatSentence) {
				words = append(words, Word{Hidden: true, Word: fmt.Sprintf("%s %s %s", formatSentence[i], formatSentence[i+1], formatSentence[i+2])})
				i += 3
			} else {
				words = append(words, Word{Hidden: false, Word: formatSentence[i]})
				i += 1
			}

		}
		words = append(words, Word{Hidden: false, Word: "\n"})
	}

	return Payload{Song: words}
}
