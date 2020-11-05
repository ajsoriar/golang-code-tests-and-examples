# Ggenerating source png images

This program scales all the images in source folder generating png files in results folder.

## Download dependencies

> go get -u golang.org/x/image/

## To run the example without building

> go run .\scaleall.2.go

## Advanced functionality

The last example `scaleall.3.go` is spectring for marams (target width) in order to calculate the scale.

Example:

> go run .\scaleall.3.go 1000

- Will scale all pictures in sources folder to width = 1000
- Height of each picture will be automatically adjusted.
