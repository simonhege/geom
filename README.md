# geom

Package geom is a geometry library in Go.

It provides definition of basic geometry structures (Point, LineString, Polygon), including Z, M and ZM variants. 
MultiGeometry and geometries collections are also provided.

Sub-packages allow [encoding to GeoJSON](https://github.com/xeonx/geom/tree/master/encoding/geojson) and [decoding from Well Known Binary](https://github.com/xeonx/geom/tree/master/encoding/wkb).
  
## Install

    go get github.com/xeonx/geom

## Docs

[![GoDoc](https://godoc.org/github.com/xeonx/geom?status.svg)](https://godoc.org/github.com/xeonx/geom)
	
## Tests

`go test` is used for testing.

## Roadmap
  * WKB encoding
  * GeoJSON decoding
  * WKT encoding/decoding
  * interoperability with popular geospatial libraries
     * GEOS via [github.com/paulsmith/gogeos](http://paulsmith.github.io/gogeos/)
	 * GDAL via [github.com/lukeroth/gdal](https://github.com/lukeroth/gdal)
  * adding more tests and benchmarks

## License

This code is licensed under the MIT license. See [LICENSE](https://github.com/xeonx/geom/blob/master/LICENSE).