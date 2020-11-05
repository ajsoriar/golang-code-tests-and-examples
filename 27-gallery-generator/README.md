# Pictures gallery JSON

This group of programs scale all the images in `source` folder generating the necessary files to run the poject `https://github.com/ajsoriar/media-gallery-web-site-2020`

## Generation sequence

In order to get the json of the gallery run the following sequence

> go run .\cleanresults.go
> go run .\scalepictures.go 150
> go run .\renameresults.go
> go run .\gallerygen.go

150 is the desired width of the thumbnails

At the end of this process the file `gallery.json` will be generated ready to be used with `https://github.com/ajsoriar/media-gallery-web-site-2020`
