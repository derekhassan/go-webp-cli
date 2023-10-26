package main

import (
	"flag"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

// Note to self: We need to manually install webp
func main() {
	inputFilePathFlag := flag.String("i", "", "Input file path")
	outputFilePathFlag := flag.String("o", "", "Output file path")
	qualityFlag := flag.Float64("q", 75, "Quality when converting image")

	flag.Parse()

	if *inputFilePathFlag == "" {
		log.Fatalln("Please supply an input image path!")
	}

	if *outputFilePathFlag == "" {
		log.Fatalln("Please supply an output image path!")
	}

	file, err := os.Open(*inputFilePathFlag)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}

	output, err := os.Create(*outputFilePathFlag)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, float32(*qualityFlag))
	if err != nil {
		log.Fatalln(err)
	}

	if err := webp.Encode(output, img, options); err != nil {
		log.Fatalln(err)
	}
	log.Printf("Created %s successfully", *outputFilePathFlag)
}
