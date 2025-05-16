package main

import "testing"

func Example_tallylingo() {
	goMain([]string{"tallylingo"})
	// Output:
	// Welcome to tallylingo!
}

// func TestHello(t *testing.T) {
// 	got := hello()
// 	want := "Welcome to tallylingo!"
// 	if got != want {
// 		t.Errorf("hello() = %q, want %q", got, want)
// 	}
// }

func TestHelpMessage(t *testing.T) {
	goMain([]string{"tallylingo", "-h"})
}
