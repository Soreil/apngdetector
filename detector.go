package apngdetector

import "bytes"

const (
	//Image header
	ihdr = "IHDR"

	//First frame
	idat = "IDAT"

	//APNG header
	actl = "acTL"
)

//check if an acTL follows between IHDR and first IDAT
func Detect(img []byte) bool {
	index := bytes.Index(img, []byte(ihdr))
	if index == -1 {
		return false
	}
	//Cut head off
	img = img[index:]

	index = bytes.Index(img, []byte(idat))
	if index == -1 {
		return false
	}
	//Cut tail off
	img = img[:index]

	index = bytes.Index(img, []byte(actl))
	//whether found the APNG header in the right location
	return index != -1
}
