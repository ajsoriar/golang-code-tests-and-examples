# Find Common Colours in an Image

We use a package to do this. The package, called prominentcolor, uses the Kmeans++ algorithm to work this out. By default, it will return us the three most popular colours after both cropping and resizing the image.

## Read this

- https://github.com/EdlinOrg/prominentcolor

## To install the package prominentcolor

> go get -v github.com/EdlinOrg/prominentcolor

## To run the example without building

> go run .\getcolors.go
