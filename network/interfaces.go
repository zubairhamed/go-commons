package network
import "errors"

type Request interface {

}

type Response interface {

}

var ERR_NO_MATCHING_ROUTE = errors.New("No matching route found")
var ERR_UNSUPPORTED_CONTENT_FORMAT = errors.New("Unsupported Content-Format")
var ERR_NO_MATCHING_METHOD = errors.New("No matching method")