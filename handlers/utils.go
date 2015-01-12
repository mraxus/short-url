package handlers

import (
)

// default host value
var host = "http://localhost:8080/"

// Sets the host name to be used for returning correct host name to the shortening service
func SetHost(value string) {
	host = value
}

// Gets the host name to be used for returning correct host name to the shortening service
func GetHost() string {
	return host
}