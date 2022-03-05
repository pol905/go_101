package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("code")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		word := "next"
		definition := "value"

		err := dict.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		dict := Dictionary{"next": "value"}
		word := "next"
		definition := "value"

		err := dict.Add(word, definition)

		assertError(t, err, ErrAlreadyExists)
		assertDefinition(t, dict, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "Ukraine \u2665"

		dict := Dictionary{word: definition}
		newDef := "Russia \u2695"

		err := dict.Update(word, newDef)

		assertError(t, err, nil)
		assertDefinition(t, dict, word, newDef)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"

		dict := Dictionary{}
		newDef := "Russia \u2695"

		err := dict.Update(word, newDef)

		assertError(t, err, ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dict := Dictionary{word: "test words"}

	dict.Delete(word)

	_, err := dict.Search(word)

	if err == nil {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertDefinition(t *testing.T, dict Dictionary, word, definition string) {
	t.Helper()

	got, err := dict.Search(word)

	if err != nil {
		t.Fatal("should find the word", err)
	}

	if definition != got {
		t.Errorf("Got: %q, want: %q", definition, got)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if want != nil && got == nil {
		t.Fatal("expected an err")
	}

	if got != want {
		t.Errorf("got %q, want %q given %q", got, want, "test")
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q given %q", got, want, "test")
	}
}
