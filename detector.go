package apngdetector

import (
	"bytes"
	"io"
)

const (
	//Image header
	ihdr = "IHDR"

	//First frame
	idat = "IDAT"

	//APNG header
	actl = "acTL"
)

// Detect checks, if an acTL follows between IHDR and first IDAT
func Detect(r io.Reader) (isAPNG bool, err error) {
	img := make([]byte, 512)
	_, err = r.Read(img)
	if err != nil {
		return
	}
	index := bytes.Index(img, []byte(ihdr))
	if index == -1 {
		return
	}
	//Cut head off
	img = img[index:]

	index = bytes.Index(img, []byte(idat))
	if index == -1 {
		return
	}
	//Cut tail off
	img = img[:index]

	index = bytes.Index(img, []byte(actl))
	//whether found the APNG header in the right location
	isAPNG = index != -1
	return
}
