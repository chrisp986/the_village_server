package database

import (
	"testing"
)

func TestBuildingString(t *testing.T) {

	buildingsID := []string{"h1", "h2", "h3"}
	got := InitBuildingsString(buildingsID)
	expected := "h1=0,h2=0,h3=0,"

	// (h4)[0],(h5)[0],(h6)[0],(w1)[0],(w2)[0],(w3)[0],(w4)[0],(w5)[0],(w6)[0],(q1)[0],(q2)[0],(q3)[0],(q4)[0],(q5)[0],(q6)[0],(m1)[0],(m2)[0],(m3)[0],(m4)[0],(m5)[0],(m6)[0],(f1)[0],(f2)[0],(f3)[0],(f4)[0],(f5)[0],(f6)[0],

	if got != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, got)
	}
}

// func TestUpdateBuildingString(t *testing.T) {

// 	bString := "(h1)[0],(h2)[0],(h3)[0],"
// 	got := updateBuildingString(bString, "h2", 1)
// 	expected := "(h1)[0],(h2)[1],(h3)[0],"

// }
