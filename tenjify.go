package tenjify

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"

	"github.com/nfnt/resize"
)

const (
	TenjiCols = 2
	TenjiRows = 4
)

func numberToTenji(num int, fillBlank bool) string {
	flags := 0
	flags += (num & 0b00001000) << 3
	flags += (num & 0b01110000) >> 1
	flags += num & 0b10000111

	if fillBlank && flags == 0 {
		return "⡀"
	}
	return string(rune(flags + 0x2800))
}

func Tenjify(fileName string, width int, threshold int, reverse bool, fillBlank bool) string {
	reader, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()

	canvasWidth := TenjiCols * width
	canvasHeight := int(math.Round(float64(bounds.Size().Y) * (float64(canvasWidth) / float64(bounds.Size().X))))

	m = resize.Resize(uint(canvasWidth), uint(canvasHeight), m, resize.Lanczos3)

	var result string

	for sy := 0; sy < canvasHeight; sy += TenjiRows {
		for sx := 0; sx < canvasWidth; sx += TenjiCols {
			n := 0

			for dy := 0; dy < TenjiRows && sy+dy < canvasHeight; dy++ {
				for dx := 0; dx < TenjiCols; dx++ {
					r, g, b, _ := m.At(sx+dx, sy+dy).RGBA()
					isBelow := (int(r>>8)+int(g>>8)+int(b>>8))/3 < threshold

					if (isBelow && !reverse) || (!isBelow && reverse) {
						diff := 1
						diff <<= 4 * dx
						diff <<= dy
						n += diff
					}
				}
			}
			result += numberToTenji(n, fillBlank)
		}

		result += "\n"
	}

	return result
}
