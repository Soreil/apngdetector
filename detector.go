package apngdetector

import "bytes"

//Image header
var ihdr = [4]byte{'I', 'H', 'D', 'R'}

//First frame
var idat = [4]byte{'I', 'D', 'A', 'T'}

//APNG header
var actl = [4]byte{'a', 'c', 'T', 'L'}

//check if an acTL follows between IHDR and first IDAT
func Detect(img []byte) bool {
	index := bytes.Index(img, ihdr[:])
	if index == -1 {
		return false
	}
	//Cut head off
	img = img[index:]

	index = bytes.Index(img, idat[:])
	if index == -1 {
		return false
	}
	//Cut tail off
	img = img[:index]

	index = bytes.Index(img, actl[:])
	//whether found the APNG header in the right location
	return index != -1
}
