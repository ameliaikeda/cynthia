package handler

import (
	"testing"
)

func TestRandom(t *testing.T) {
	crit, miss, fail := 0, 0, 0
	for range make([]struct{}, 1000000) {
		res := random()

		switch res {
		case 1:
			miss++
		case 20:
			crit++
		case 0, 21:
			fail++
		}
	}

	if crit < 40000 {
		t.Logf("crit values: %d", crit)
		t.FailNow()
	}

	if fail > 0 {
		t.Log("0 should not be returned")
		t.FailNow()
	}
	if miss < 40000 {
		t.Logf("miss values: %d", miss)
		t.FailNow()
	}
}