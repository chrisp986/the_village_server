package database

import (
	"reflect"
	"testing"

	"github.com/chrisp986/the_village_server/internal/models"
)

const testBuildingString string = "h1=1,h2=2,h3=3,h4=4,h5=5,h6=6,w1=10,w2=20,w3=30,w4=40,w5=50,w6=60,q1=01,q2=02,q3=03,q4=04,q5=05,q6=06,m1=07,m2=08,m3=09,m4=10,m5=550,m6=660,f1=1110,f2=2210,f3=330,f4=440,f5=550,f6=606,"

const oldtestBuildingString string = "(0)[0],(1)[0],(2)[0],(3)[0],(4)[0],(5)[0],(6)[0],(7)[0],(8)[0],(9)[0],(10)[0],(11)[0],(12)[0],(13)[0],(14)[0],(15)[0],(16)[0],(17)[0],(18)[0],(19)[0],(20)[0],(21)[0],(22)[0],(23)[0],(24)[0],"

func BenchmarkSplitString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		splitString(testBuildingString)
	}
}

func TestSplitString(t *testing.T) {

	got := splitString(testBuildingString)
	expected := []models.BuildingCount{
		{BuildingID: "h1", Count: 1},
		{BuildingID: "h2", Count: 2},
		{BuildingID: "h3", Count: 3},
		{BuildingID: "h4", Count: 4},
		{BuildingID: "h5", Count: 5},
		{BuildingID: "h6", Count: 6},
		{BuildingID: "w1", Count: 10},
		{BuildingID: "w2", Count: 20},
		{BuildingID: "w3", Count: 30},
		{BuildingID: "w4", Count: 40},
		{BuildingID: "w5", Count: 50},
		{BuildingID: "w6", Count: 60},
		{BuildingID: "q1", Count: 1},
		{BuildingID: "q2", Count: 2},
		{BuildingID: "q3", Count: 3},
		{BuildingID: "q4", Count: 4},
		{BuildingID: "q5", Count: 5},
		{BuildingID: "q6", Count: 6},
		{BuildingID: "m1", Count: 7},
		{BuildingID: "m2", Count: 8},
		{BuildingID: "m3", Count: 9},
		{BuildingID: "m4", Count: 10},
		{BuildingID: "m5", Count: 550},
		{BuildingID: "m6", Count: 660},
		{BuildingID: "f1", Count: 1110},
		{BuildingID: "f2", Count: 2210},
		{BuildingID: "f3", Count: 330},
		{BuildingID: "f4", Count: 440},
		{BuildingID: "f5", Count: 550},
		{BuildingID: "f6", Count: 606},
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v, got %v", expected, got)
	}

}

func BenchmarkSplitBuildingsString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SplitBuildingsString(oldtestBuildingString)
	}
}
