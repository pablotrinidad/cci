# Cloud Cover Index

CCI is a school project that computes the cloud cover index of any image of the sky using
a basic RGB-based cloud segmentation algorithm.

## How to use

#### Build at home

1. Download 1.15 Golang compiler. Instructions [here](https://golang.org/doc/install).
2. Build binary with `go build -o cci cmd/main.go`
3. (optionally) Run and build in the same command `go run cmd/main.go`

#### Instructions

Use built binary with the following flags
- **`-src`** path to the source image
- **`-mask`** path to the mask image
- **`-s`** if set, an output file with the segmentation results will be created
- **`-out`** path to the output file (will be created for you)

**Examples**

- Just output cloud cover index:

```sh
./cci -src myImage.jpg -mask myMask.png
```

- Output cloud cover image and segmentation file:

```sh
./cci -src myImage.jpg -mask myMask.png -s -out myOutputFile.png
```

## Package structure

* `cmd` is a command line interface that read input files and parses them into Golang's `image` pgk abstractions.
* `alg` is a package that deals with `image` objects and run segmentation algorithms.


## Results

| Source Image   |                Segmentation result           | CCI           |
| :-------------------------: | :-----------------------------: | :-----------: |
| ![360-1](samples/360-1.jpg) | ![360-1 seg](samples/360-1.png) | 0.181054      |

*command*: **`./cci -src samples/360-1.jpg -mask samples/mask-1350-sq.png -s -out samples/360-1.png`**


| Source Image   |                Segmentation result           | CCI           |
| :-------------------------: | :-----------------------------: | :-----------: |
| ![360-2](samples/360-2.jpg) | ![360-2 seg](samples/360-2.png) | 0.292510      |

*command*: **`./cci -src samples/360-2.jpg -mask samples/mask-1350-sq.png -s -out samples/360-2.png`**

| Source Image   |                Segmentation result           | CCI           |
| :-------------------------: | :-----------------------------: | :-----------: |
| ![360-3](samples/360-3.jpg) | ![360-3 seg](samples/360-3.png) | 0.44803       |

*command*: **`./cci -src samples/360-3.jpg -mask samples/mask-1350-sq.png -s -out samples/360-3.png`**

| Source Image   |                Segmentation result           | CCI           |
| :-------------------------: | :-----------------------------: | :-----------: |
| ![360-4](samples/360-4.jpg) | ![360-4 seg](samples/360-4.png) | 0.505722      |


*command*: **`./cci -src samples/360-4.jpg -mask samples/mask-1350-sq.png -s -out samples/360-4.png`**

| Source Image   |                Segmentation result           | CCI           |
| :-------------------------: | :-----------------------------: | :-----------: |
| ![360-5](samples/360-5.jpg) | ![360-5 seg](samples/360-5.png) | 0.517248      |

*command*: **`./cci -src samples/360-5.jpg -mask samples/mask-1350-sq.png -s -out samples/360-5.png`**

| Source Image   |                Segmentation result           | CCI           |
| :-------------------------: | :-----------------------------: | :-----------: |
| ![360-6](samples/360-6.jpg) | ![360-6 seg](samples/360-6.png) | 0.324765      |

*command*: **`./cci -src samples/360-6.jpg -mask samples/mask-1350-sq.png -s -out samples/360-6.png`**

| Source Image   |                Segmentation result           | CCI           |
| :-------------------------: | :-----------------------------: | :-----------: |
| ![rect-1](samples/rect-1.jpg) | ![rect-1 seg](samples/rect-1-circle-mask.png) |  0.121358      |

*command*: **`./cci -src samples/rect-1.jpg -mask samples/mask-1350-sq.png -s -out samples/rect-1-circle-mask.png`**

| Source Image   |                Segmentation result           | CCI           |
| :-------------------------: | :-----------------------------: | :-----------: |
| ![rect-1](samples/rect-1.jpg) | ![rect-1 seg](samples/rect-1-rect-full-mask.png) |  0.167027     |

*command*: **`./cci -src samples/rect-1.jpg -mask samples/3264-1831-rect-mask.png -s -out samples/rect-1-rect-full-mask.png`**

| Source Image   |                Segmentation result           | CCI           |
| :-------------------------: | :-----------------------------: | :-----------: |
| ![rect-1](samples/rect-1.jpg) | ![rect-1 seg](samples/rect-1-rect-happy-mask.png) |  0.007702     |

*command*: **`./cci -src samples/rect-1.jpg -mask samples/3264-1831-happy-mask.png -s -out samples/rect-1-rect-happy-mask.png`**

| Source Image   |                Segmentation result           | CCI           |
| :-------------------------: | :-----------------------------: | :-----------: |
| ![rect-1](samples/small-lots-of-clouds.jpg) | ![rect-1 seg](samples/small-lots-of-clouds.png) |  0.308125     |

*command*: **`./cci -src samples/small-lots-of-clouds.jpg -mask samples/3264-1831-rect-mask.png -s -out samples/small-lots-of-clouds.png`**


## Contributors
 
 - Pablo Trinidad ([@pablotrinidad](https://github.com/pablotrinidad))
 - Diego Jardón ([@jardondiego](https://github.com/jardondiego))
 



