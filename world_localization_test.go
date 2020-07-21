package mundo_pequeno

import (
	"reflect"
	"testing"
)

func TestLocalization(t *testing.T) {

	t.Run("only me in the world", func(t *testing.T) {
		me := Position{1, 1}
		world := World{}
		world.Add(me)
		//want := []Position{} why?
		var want []Position

		got := world.Nearest(me)

		assertPositions(t, got, want)
	})

	t.Run("me and another person in the world", func(t *testing.T) {
		me := Position{1, 1}
		friend := Position{2, 2}
		world := World{}
		world.Add(me)
		world.Add(friend)
		want := []Position{friend}

		got := world.Nearest(me)

		assertPositions(t, got, want)
	})

	t.Run("friends at the same position", func(t *testing.T) {
		me := Position{1, 1}
		friendA := Position{1, 1}
		world := World{}
		world.Add(me)
		err := world.Add(friendA)

		assertSamePositionError(t, err, ErrFriendAtSamePosition)
	})

	t.Run("me and two other people in the world", func(t *testing.T) {
		me := Position{1, 1}
		friendA := Position{2, 2}
		friendB := Position{3, 3}
		world := World{}
		world.Add(me)
		world.Add(friendA)
		world.Add(friendB)
		want := []Position{friendA, friendB}

		got := world.Nearest(me)

		assertPositions(t, got, want)
	})

	t.Run("me and three other friends", func(t *testing.T) {
		me := Position{1, 1}
		friendA := Position{2, 2}
		friendB := Position{3, 3}
		friendC := Position{4, 4}
		world := World{}
		world.Add(me)
		world.Add(friendA)
		world.Add(friendB)
		world.Add(friendC)
		want := []Position{friendA, friendB, friendC}

		got := world.Nearest(me)

		assertPositions(t, got, want)
	})


	t.Run("me and a crowd", func(t *testing.T) {
		me := Position{1, 1}
		friendA := Position{2, 2}
		friendB := Position{3, 3}
		friendD := Position{5, 5}
		friendC := Position{4, 4}
		world := World{}
		world.Add(me)
		world.Add(friendA)
		world.Add(friendB)
		world.Add(friendC)
		world.Add(friendD)
		want := []Position{friendA, friendB, friendC}

		got := world.Nearest(me)

		assertPositions(t, got, want)
	})
}

func assertPositions(t *testing.T, got []Position, want []Position) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertSamePositionError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
