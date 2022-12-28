package movie

import "fmt"

func ReviewMovie(name string, rating float64) {
	fmt.Printf("I reviewd %s and it's rating is %.2f\n", name, rating)
}

func FindMovieName(imdbId string) string {
	switch imdbId {
	case "tt1":
		return "AvenJeab"
	case "tt2":
		return "AvaTeab"
	}
	return "not good"
}
