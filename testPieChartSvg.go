package main

// This chart package (i.e. directory) needs to be accessible for this file to use it
// the src directory that it belongs to, must have its parent in the GOPATH env var
// The other packages must on the GOPATH as well but they are shipped with go

import (
	"chart"
	"fmt"
	"io"
	"os"
)

// Write to a file
func stringToFile(fileName string, stringToWrite string) {

	fmt.Println("Writing: " + fileName)
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}
	n, err := io.WriteString(f, stringToWrite)
	if err != nil {
		fmt.Println(n, err)
	}
	f.Close()
}

func main() {
	testContent := chart.PieDraw([]float64{10, 30, 80}, "Sample Title")
	//fmt.Println(testContent)

	// Write to a file
	fileName := "SamplePie.svg"

	stringToFile(fileName, testContent)

}
