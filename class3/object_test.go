package class3_test

import (
	"github.com/victoryRo/tdd/class3"
	"reflect"
	"testing"
)

func TestDog(t *testing.T) {
	want := &class3.Dog{
		Name: "Firulais",
		Age:  1,
		Kind: class3.Kind{
			Name: "Criollo",
		},
	}
	got := &class3.Dog{
		Name: "Firulais",
		Age:  1,
		Kind: class3.Kind{
			Name: "Criollo",
		},
	}

	//t.Logf("want %p got %p", want, got)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("we expected %v and got %v", want, got)
	}
}
