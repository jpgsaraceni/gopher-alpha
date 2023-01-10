package testes_test

import (
	"testing"

	"github.com/jpgsaraceni/gopher-alpha/exemplos/testes"
)

func TestAnimal(t *testing.T) {
	testDog := testes.Animal{
		Species: "dog",
	}

	testCat := testes.Animal{
		Name:    "Miau",
		Species: "cat",
	}

	assertIsDog(t, true, testDog)
	assertIsDog(t, false, testCat)
	assertHasName(t, false, testDog)
	assertHasName(t, true, testCat)
}

func assertIsDog(t *testing.T, wantDog bool, animal testes.Animal) {
	t.Helper()

	gotDog := animal.IsDog()
	if gotDog != wantDog {
		t.Logf("expected IsDog to be %t but got %t for animal %v", wantDog, gotDog, animal)
	}
}

func assertHasName(t *testing.T, wantName bool, animal testes.Animal) {
	t.Helper()

	gotName := animal.HasName()
	name := animal.Name
	if len(name) > 0 != gotName {
		t.Logf("got HasName %t for animal.Name == %s", gotName, name)
	}
}
