package main

import (
	"fmt"
	"log"
	"os"

	"bufio"
	"strings"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the text to generate QR code: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	input = strings.TrimSpace(input)

	fmt.Print("Enter the output file name (e.g., my_massage): ")
	outputFile, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading output file name: %v", err)
	}
	outputFile = strings.TrimSpace(outputFile)
	outputFile = outputFile + ".png"
	err = generateQRCode(input, outputFile)
	if err != nil {
		log.Fatalf("Error generating QR code: %v", err)
	}

	fmt.Println("QR code generated successfully!")
}

func generateQRCode(text, outputFile string) error {
	err := qrcode.WriteFile(text, qrcode.Medium, 256, outputFile)
	if err != nil {
		return err
	}

	file, err := os.Open("qr.png")
	if err != nil {
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	fmt.Printf("QR code image size: %d bytes\n", fileInfo.Size())

	return nil
}
