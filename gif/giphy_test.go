package gif

import (
	"testing"
)

// // Single case function test
// func TestModifyInput(t *testing.T) {
// 	fmt.Println("TestModifyInput")
// 	expected := "Hello%20Dog!"
// 	result := ModifyInput("Hello Dog!")
// 	if expected != result {
// 		t.Error("Expected Hello Dog! to become Hello%20Dog!")
// 	}
// }

// Two table tests
// Table test #1
func TestTableModifyInput(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"Hello Dog!", "Hello%20Dog!"},
		{"Mama Miya!", "Mama%20Miya!"},
		{"Lets Run!", "Lets%20Run!"},
		{"I'm really thirsty.", "I'm%20really%20thirsty."},
		{"Here, have a sip of water.", "Here,%20have%20a%20sip%20of%20water."},
	}

	for _, test := range tests {
		if output := ModifyInput(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.input, test.expected, output)
		}
	}
}

//Table test #2
func TestTableBetween(t *testing.T) {
	var tests = []struct {
		given    string
		first    string
		second   string
		expected string
	}{
		{"My Todo list: Have fun, get a job, make money!", "list: ", "!", "Have fun, get a job, make money"},
		{"bob marley ok", "bob ", " ok", "marley"},
		{"https://thisisacoolurldontyouthink.com", "https://", ".com", "thisisacoolurldontyouthink"},
	}

	for _, test := range tests {
		if output := Between(test.given, test.first, test.second); output != test.expected {
			t.Error("Test Failed: {} given string, {} beginning point substring, {} end point, {} expected, recieved: {}", test.given, test.first, test.second, test.expected, output)
		}
	}
}

// Benchmark tests
/* I don't understand why the benchmark test loops multiple times:
91147             12215 ns/op
PASS
ok      github.com/kmurata798/goslackit/gif     2.260s
??? :O
*/

func benchmarkModifyInput(input string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		ModifyInput(input)
	}
}

func BenchmarkModifyInputOhMyGoodness(b *testing.B) { benchmarkModifyInput("Oh my goodness!", b) }

// All these 3 benchmark tests work, but they loop multiple times for some reason.
// func BenchmarkModifyInputGetToTheChoppo(b *testing.B) { benchmarkModifyInput("Get to the choppa!", b) }
// func BenchmarkModifyInputGrabTheBag(b *testing.B)     { benchmarkModifyInput("Grab the bag.", b) }
