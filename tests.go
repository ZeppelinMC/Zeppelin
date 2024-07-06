package main

import (
	"bytes"
	"fmt"
	"reflect"

	"aether/nbt"
)

// ANSI escape codes for colors
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
)

type TestStruct struct {
	Name      string   `nbt:"name"`
	Age       int32    `nbt:"age"`
	Balance   float64  `nbt:"balance"`
	Nicknames []string `nbt:"nicknames"`
}

type NestedStruct struct {
	ID   int32  `nbt:"id"`
	Name string `nbt:"name"`
}

type ComplexStruct struct {
	Title    string           `nbt:"title"`
	Values   []float64        `nbt:"values"`
	Details  NestedStruct     `nbt:"details"`
	Metadata map[string]int32 `nbt:"metadata"`
}

type EdgeCaseStruct struct {
	EmptyString     string            `nbt:"emptyString"`
	ZeroInt         int32             `nbt:"zeroInt"`
	NegativeInt     int32             `nbt:"negativeInt"`
	LargeFloat      float64           `nbt:"largeFloat"`
	EmptyList       []string          `nbt:"emptyList"`
	EmptyMap        map[string]int32  `nbt:"emptyMap"`
	NestedEmptyList []NestedStruct    `nbt:"nestedEmptyList"`
	ComplexMap      map[string]string `nbt:"complexMap"`
}

var loadingFrames = []rune{'⠋', '⠙', '⠹', '⠸', '⠼', '⠴', '⠦', '⠧', '⠇', '⠏'}

func printHeader(title string) {
	fmt.Printf("%s\n%s%s\n", Cyan, title, Reset)
}

func printSuccess(message string) {
	fmt.Printf("%s✔ %s%s\n", Green, message, Reset)
}

func printError(message string) {
	fmt.Printf("%s✘ %s%s\n", Red, message, Reset)
}

func runTest(testFunc func() bool, index, total int) bool {
	for i := 0; i < len(loadingFrames); i++ {
		fmt.Printf("\r%s %d/%d", string(loadingFrames[i]), index, total)
	}
	fmt.Printf("\r%s %d/%d passed (%d failed)\n", Green, index, total, 0)
	return testFunc()
}

func testSimpleStruct() bool {
	// Create a buffer to write the encoded data to
	buf := new(bytes.Buffer)

	// Create an encoder and encode the data
	encoder := nbt.NewEncoder(buf)
	data := TestStruct{
		Name:      "Alice",
		Age:       30,
		Balance:   1000.5,
		Nicknames: []string{"Al", "Alicia"},
	}
	if err := encoder.Encode("root", data); err != nil {
		printError(fmt.Sprintf("Encoding failed: %v", err))
		return false
	}

	// Create a decoder and decode the data
	decoder := nbt.NewDecoder(buf)
	var decodedData TestStruct
	if _, err := decoder.Decode(&decodedData); err != nil {
		printError(fmt.Sprintf("Decoding failed: %v", err))
		return false
	}

	// Compare the original and decoded data
	if !reflect.DeepEqual(data, decodedData) {
		printError(fmt.Sprintf("Decoded data does not match original: got %+v, want %+v", decodedData, data))
		return false
	}

	return true
}

func testNestedStruct() bool {
	// Create a buffer to write the encoded data to
	buf := new(bytes.Buffer)

	// Create an encoder and encode the data
	encoder := nbt.NewEncoder(buf)
	data := NestedStruct{
		ID:   1,
		Name: "Nested",
	}
	if err := encoder.Encode("root", data); err != nil {
		printError(fmt.Sprintf("Encoding failed: %v", err))
		return false
	}

	// Create a decoder and decode the data
	decoder := nbt.NewDecoder(buf)
	var decodedData NestedStruct
	if _, err := decoder.Decode(&decodedData); err != nil {
		printError(fmt.Sprintf("Decoding failed: %v", err))
		return false
	}

	// Compare the original and decoded data
	if !reflect.DeepEqual(data, decodedData) {
		printError(fmt.Sprintf("Decoded data does not match original: got %+v, want %+v", decodedData, data))
		return false
	}

	return true
}

func testComplexStruct() bool {
	// Create a buffer to write the encoded data to
	buf := new(bytes.Buffer)

	// Create an encoder and encode the data
	encoder := nbt.NewEncoder(buf)
	data := ComplexStruct{
		Title:   "Complex",
		Values:  []float64{1.1, 2.2, 3.3},
		Details: NestedStruct{ID: 2, Name: "NestedDetail"},
		Metadata: map[string]int32{
			"key1": 1,
			"key2": 2,
		},
	}
	if err := encoder.Encode("root", data); err != nil {
		printError(fmt.Sprintf("Encoding failed: %v", err))
		return false
	}

	// Create a decoder and decode the data
	decoder := nbt.NewDecoder(buf)
	var decodedData ComplexStruct
	if _, err := decoder.Decode(&decodedData); err != nil {
		printError(fmt.Sprintf("Decoding failed: %v", err))
		return false
	}

	// Compare the original and decoded data
	if !reflect.DeepEqual(data, decodedData) {
		printError(fmt.Sprintf("Decoded data does not match original: got %+v, want %+v", decodedData, data))
		return false
	}

	return true
}

func testEdgeCases() bool {
	// Create a buffer to write the encoded data to
	buf := new(bytes.Buffer)

	// Create an encoder and encode the data
	encoder := nbt.NewEncoder(buf)
	data := EdgeCaseStruct{
		EmptyString:     "",
		ZeroInt:         0,
		NegativeInt:     -12345,
		LargeFloat:      1e+30,
		EmptyList:       []string{},
		EmptyMap:        map[string]int32{},
		NestedEmptyList: []NestedStruct{},
		ComplexMap: map[string]string{
			"first":  "value1",
			"second": "value2",
		},
	}
	if err := encoder.Encode("root", data); err != nil {
		printError(fmt.Sprintf("Encoding failed: %v", err))
		return false
	}

	// Create a decoder and decode the data
	decoder := nbt.NewDecoder(buf)
	var decodedData EdgeCaseStruct
	if _, err := decoder.Decode(&decodedData); err != nil {
		printError(fmt.Sprintf("Decoding failed: %v", err))
		return false
	}

	// Compare the original and decoded data
	if !reflect.DeepEqual(data, decodedData) {
		printError(fmt.Sprintf("Decoded data does not match original: got %+v, want %+v", decodedData, data))
		return false
	}

	return true
}

func runTests() {
	printHeader("Running NBT Tests")

	tests := []func() bool{
		testSimpleStruct,
		testNestedStruct,
		testComplexStruct,
		testEdgeCases,
	}

	totalTests := len(tests)
	passedTests := 0
	failedTests := 0

	for i, test := range tests {
		if runTest(test, i+1, totalTests) {
			passedTests++
		} else {
			failedTests++
		}
	}

	fmt.Printf("\nTest Results: %d/%d passed, %d failed\n", passedTests, totalTests, failedTests)
	if failedTests > 0 {
		printError("Some tests failed. Check the above output for details.")
	} else {
		printSuccess("All tests passed successfully!")
	}
}
