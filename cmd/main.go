package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
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
	flag.StringVar(&srcImage, "src", "", "Source image, it can be of any dimensions and JPEG or PNG encoded.")
	flag.StringVar(&srcMask, "mask", "", "Mask file of any dimensions. Source and mask will be center aligned and only those pixels that match a white pixel from the mask will be used during computation.")
	flag.BoolVar(&outputSegmentation, "s", false, "If true, it will output a black and white segmentation result as a PNG encoded image.")
	flag.StringVar(&outputFile, "out", "", "Output file name (if segmentation -s is enabled)")
	flag.Parse()

	if srcImage == "" {
		flag.Usage()
		fmt.Println("Missing source image.")
		os.Exit(1)
	}
	if srcMask == "" {
		flag.Usage()
		fmt.Println("Missing mask image.")
		os.Exit(1)
	}
	if outputSegmentation && outputFile == "" {
		flag.Usage()
		fmt.Println("Missing output file, please set when enabling segmentation output.")
		os.Exit(1)
	}
}

func main() {
	src, mask := loadImages()

	cci := alg.NewCCI(src, mask)
	index := cci.Run()
	fmt.Printf("Cloud Cover Index: %f\n", index)

	if outputSegmentation {
		img, err := cci.SaveSegmentation()
		if err != nil {
			fmt.Printf("Failed saving segmentation output\n\t%v", err)
			os.Exit(1)
		}
		outFile, err := os.Create(outputFile)
		if err != nil {
			fmt.Printf("Failed creating output file\n\t%v", err)
			os.Exit(1)
		}
		if err := png.Encode(outFile, img); err != nil {
			fmt.Printf("Failed encoding result image\n%v", err)
			os.Exit(1)
		}
		fmt.Printf("Segmentation image saved successfully at %s\n", outputFile)
	}
}

// loadImages returns a tuple of the parsed sky picture and circle mask respectively.
// If an error occurs, program will exit with non-zero exit code.
func loadImages() (image.Image, image.Image) {
	imgReader, err := os.Open(srcImage)
	if err != nil {
		fmt.Printf("Failed opening source file\n\t%v", err)
		os.Exit(1)
	}
	defer imgReader.Close()
	src, _, err := image.Decode(imgReader)
	if err != nil {
		fmt.Printf("Failed reading source image\n\t%v", err)
		os.Exit(1)
	}

	maskReader, err := os.Open(srcMask)
	if err != nil {
		fmt.Printf("Failed opening mask file\n\t%v", err)
		os.Exit(1)
	}
	defer maskReader.Close()
	mask, _, err := image.Decode(maskReader)
	if err != nil {
		fmt.Printf("Failed reading mask image\n\t%v", err)
		os.Exit(1)
	}
	return src, mask
}
