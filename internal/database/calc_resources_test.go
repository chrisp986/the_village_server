package database

const testBuildingString string = "h1=1,h2=2,h3=3,h4=4,h5=5,h6=6,w1=10,w2=20,w3=30,w4=40,w5=50,w6=60,q1=01,q2=02,q3=03,q4=04,q5=05,q6=06,m1=07,m2=08,m3=09,m4=10,m5=550,m6=660,f1=1110,f2=2210,f3=330,f4=440,f5=550,f6=606,"

const oldtestBuildingString string = "(0)[0],(1)[0],(2)[0],(3)[0],(4)[0],(5)[0],(6)[0],(7)[0],(8)[0],(9)[0],(10)[0],(11)[0],(12)[0],(13)[0],(14)[0],(15)[0],(16)[0],(17)[0],(18)[0],(19)[0],(20)[0],(21)[0],(22)[0],(23)[0],(24)[0],"

// func BenchmarkSplitString(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		splitString(testBuildingString)
// 	}
// }

// func TestSplitString(t *testing.T) {

// 	got := splitString(testBuildingString)
// 	expected := []models.BuildingCount{
// 		{WorkerID: "h1", Count: 1},
// 		{WorkerID: "h2", Count: 2},
// 		{WorkerID: "h3", Count: 3},
// 		{WorkerID: "h4", Count: 4},
// 		{WorkerID: "h5", Count: 5},
// 		{WorkerID: "h6", Count: 6},
// 		{WorkerID: "w1", Count: 10},
// 		{WorkerID: "w2", Count: 20},
// 		{WorkerID: "w3", Count: 30},
// 		{WorkerID: "w4", Count: 40},
// 		{WorkerID: "w5", Count: 50},
// 		{WorkerID: "w6", Count: 60},
// 		{WorkerID: "q1", Count: 1},
// 		{WorkerID: "q2", Count: 2},
// 		{WorkerID: "q3", Count: 3},
// 		{WorkerID: "q4", Count: 4},
// 		{WorkerID: "q5", Count: 5},
// 		{WorkerID: "q6", Count: 6},
// 		{WorkerID: "m1", Count: 7},
// 		{WorkerID: "m2", Count: 8},
// 		{WorkerID: "m3", Count: 9},
// 		{WorkerID: "m4", Count: 10},
// 		{WorkerID: "m5", Count: 550},
// 		{WorkerID: "m6", Count: 660},
// 		{WorkerID: "f1", Count: 1110},
// 		{WorkerID: "f2", Count: 2210},
// 		{WorkerID: "f3", Count: 330},
// 		{WorkerID: "f4", Count: 440},
// 		{WorkerID: "f5", Count: 550},
// 		{WorkerID: "f6", Count: 606},
// 	}

// 	if !reflect.DeepEqual(got, expected) {
// 		t.Errorf("Expected %v, got %v", expected, got)
// 	}

// }
