package main

import (
	"flag"
	"image"
	"log"
	"os"

	_ "image/jpeg"
	_ "image/png"

	"github.com/pablotrinidad/cci/alg"
)

var srcImage string
var srcMask string
var outputSegmentation bool
var outputFile string

func init() {
	log.SetPrefix("")

	flag.StringVar(&srcImage, "src", "", "source image")
	flag.StringVar(&srcMask, "mask_src", "mask-1350-sq.png", "circle mask used to filter out ")
	flag.BoolVar(&outputSegmentation, "s", false, "output black and white segmentation result")
	flag.StringVar(&outputFile, "out", "out.png", "output file if segmentation is enabled")
	flag.Parse()

	if srcImage == "" {
		flag.Usage()
		log.Fatalln("Missing source image.")
	}
	if srcMask == "" {
		flag.Usage()
		log.Fatalln("Missing circle mask.")
	}
	if outputSegmentation && outputFile == "" {
		flag.Usage()
		log.Fatalln("Missing output file, please set when enabling segmentation output.")
	}
}

func main() {
	src, mask := loadImages()
	cci := alg.NewCCI(src, mask)
	cci.Run()
	if outputSegmentation {
		if err := cci.SaveSegmentation(outputFile); err != nil {
			log.Fatal(err)
		}
	}
}

// loadImages returns a tuple of the parsed sky picture and circle mask respectively.
// If an error occurs, program will exit with non-zero exit code.
func loadImages() (image.Image, image.Image) {
	imgReader, err := os.Open(srcImage)
	if err != nil {
		log.Fatalf("Failed opening source file\n%v", err)
	}
	defer imgReader.Close()
	src, _, err := image.Decode(imgReader)
	if err != nil {
		log.Fatalf("Failed reading source image\n%v", err)
	}

	maskReader, err := os.Open(srcMask)
	if err != nil {
		log.Fatalf("Failed opening mask file\n%v", err)
	}
	defer maskReader.Close()
	mask, _, err := image.Decode(maskReader)
	if err != nil {
		log.Fatalf("Failed reading mask image\n%v", err)
	}
	return src, mask
}
