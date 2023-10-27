package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

func main() {
	inputFilePathFlag := flag.String("i", "", "Input file path")
	outputFilePathFlag := flag.String("o", "", "Output file path")
	qualityFlag := flag.Float64("q", 75, "Quality when converting image")

	flag.Parse()

	if *inputFilePathFlag == "" {
		fmt.Println("Please supply an input image path!")
		os.Exit(1)
	}

	if *outputFilePathFlag == "" {
		fmt.Println("Please supply an output image path!")
		os.Exit(1)
	}

	file, err := os.Open(*inputFilePathFlag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output, err := os.Create(*outputFilePathFlag)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer output.Close()

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, float32(*qualityFlag))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := webp.Encode(output, img, options); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Created %s successfully\n", *outputFilePathFlag)
}
