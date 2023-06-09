package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io"
	"os"
)

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "input format = ", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 99})
}
func main() {
	fi, _ := os.Open("/Users/wuzhiyong/Desktop/a.png")
	if err := toJPEG(fi, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}
