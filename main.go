package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bron10/parking-lot/packages"
)

const outputFile string = "output.txt"

func askUserForFilename() string {
	fmt.Println("Enter file name")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimRight(str, "\n")
	return str
}

func main() {
	fileName := askUserForFilename()
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	outputFile, outputError := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file creation\n")
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		output, err := packages.Serve(args)
		if err != nil {
			fmt.Printf("Error occured while executing command")
			return
		}
		outputWriter.WriteString(output + "\n")
	}
	outputWriter.Flush()
}
