package service

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"
)

/*
Resizes an image
*/
func ResizeImage(inputBytes []byte, width, height int) ([]byte, error) {
	// Decode the image
	img, _, err := image.Decode(bytes.NewReader(inputBytes))
	if err != nil {
		return nil, err
	}

	// Creates new image of desired size
	newImg := image.NewRGBA(image.Rect(0, 0, width, height))

	// Warps given image to new image
	draw.Draw(newImg, newImg.Bounds(), img, img.Bounds().Min, draw.Src)

	var outputBuf bytes.Buffer
	err = png.Encode(&outputBuf, newImg)
	if err != nil {
		return nil, err
	}

	return outputBuf.Bytes(), nil
}
