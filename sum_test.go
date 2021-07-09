package main

import "testing"

func TestSum(t *testing.T) {

	result := sum(2, 3)

	if result != 5 {
		t.Error("The result must be 5")
	}
}

func TestMain(t *testing.T) {
	main()
	t.Log(7)
}
