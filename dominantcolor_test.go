package dominantcolor_test

import (
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"math"
	"os"
	"testing"

	"github.com/brunomvsouza/dominantcolor"
)

// https://www.mozilla.org/en-US/styleguide/identity/firefox/color/
var firefoxOrange = color.RGBA{R: 230, G: 96}

func hex(c color.RGBA) string {
	return "#" + fmt.Sprintf("%.2X%.2X%.2X", c.R, c.G, c.B)
}

func distance(a, b color.RGBA) float64 {
	dr := uint32(a.R) - uint32(b.R)
	dg := uint32(a.G) - uint32(b.G)
	db := uint32(a.B) - uint32(b.B)
	return math.Sqrt(float64(dr*dr + dg*dg + db*db))
}

func TestDominantColorFromImage(t *testing.T) {
	f, err := os.Open("firefox.png")
	if err != nil {
		t.Error(err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	dc := dominantcolor.NewDefault()
	color := dc.FromImage(img)
	distance := distance(color, firefoxOrange)
	t.Log("Found dominant color:", hex(color))
	t.Log("Firefox orange:      ", hex(firefoxOrange))
	t.Logf("Distance:             %.2f", distance)
	if distance > 50 {
		t.Error("found color is not close")
	}
}

func BenchmarkFind(b *testing.B) {
	f, err := os.Open("firefox.png")
	if err != nil {
		b.Fatal(err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		b.Fatal(err)
	}
	f.Close()
	dc := dominantcolor.NewDefault()
	for i := 0; i < b.N; i++ {
		dc.FromImage(img)
	}
}
