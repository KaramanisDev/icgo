// Port of lmembers.shaw.ca/el.supremo/MagickWand/bunny.htm to Go
package main

import (
	"gopkg.in/gographics/imagick.v2/imagick"
	"net/http"
	"log"
	"io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	imgURL := r.URL.Query().Get("image")
	filter := r.URL.Query().Get("filter")

	response, _ := http.Get(imgURL)
	blob, _ := ioutil.ReadAll(response.Body)


	resultBlob := applyFilter(blob, filter)

	w.Header().Set("Content-Type", "image/jpg")
	w.Write(resultBlob)
}

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func applyFilter(imageBlob []byte, filterType string) []byte{
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	mw.ReadImageBlob(imageBlob);

	switch filterType {
		case "sketch":
			mw.SketchImage(100,5,20)
		case "sepia":
			mw.SepiaToneImage(20)
		case "blur":
			mw.BlurImage(0,10)
		case "negative":
			mw.NegateImage(false)
	}

	return mw.GetImageBlob()
}