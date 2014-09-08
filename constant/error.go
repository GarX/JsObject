package constant

import "errors"

var EmptyIDError error = errors.New("object with empty identifier")
var NotMapError error = errors.New("object is not a map")
