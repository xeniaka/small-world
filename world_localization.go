package mundo_pequeno

import (
	"math"
	"sort"
)

const numberOfNearestFriends = 3

const (
	ErrFriendAtSamePosition = WorldErr("you could not have friends at the same position")
)

type WorldErr string

func (e WorldErr) Error() string {
	return string(e)
}

type Position struct {
	X int
	Y int
}

type World struct {
	positions []Position
}

func (w *World) Add(p Position) error {
	var contains = false
	for _, pos := range w.positions {
		if p == pos {
			contains = true
		}
	}

	if contains {
		return ErrFriendAtSamePosition
	}

	w.positions = append(w.positions, p)
	return nil
}

func (w *World) Nearest(p Position) []Position {
	var friends = make(map[float64]Position)

	// Calculate distance to friends
	for _, position := range w.positions {
		if position != p {
			d := distance(p, position)
			friends[d] = position
		}
	}

	// Redefine number of friends
	var numberOfFriends = len(friends)
	if numberOfFriends > numberOfNearestFriends {
		numberOfFriends = numberOfNearestFriends
	}

	// Sort friends by distance
	distances := sortDistances(friends)

	// Get only nearest friends
	var nearestFriends []Position
	for i := 0; i < numberOfFriends; i++ {
		nearestFriends = append(nearestFriends, friends[distances[i]])
	}

	return nearestFriends
}

func distance(p1 Position, p2 Position) float64 {
	return math.Sqrt(math.Pow(float64(p2.X-p1.X), 2) + math.Pow(float64(p2.Y-p1.Y), 2))
}

func sortDistances(friends map[float64]Position) []float64 {
	var keys []float64
	for k := range friends {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	return keys
}
