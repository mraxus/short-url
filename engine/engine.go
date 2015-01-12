package engine

import (
	"math/rand"
)

// A slice of runes that are the building blocks for the unique keys generated
var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

// The length of the key to be generated
var currentKeyLength int = 1

// The lookup map stores all keys with its associated redirect urls
// Should this service become a real server and used in production,
// this data structure needs to be converted to a persistant one
var lookup map[string]string

// Initializes the engine
func init() {
	// setup a new key lookup object
	lookup = make(map[string]string)
}

// Private function to generate keys
func generateKey(length int) string {
	
	keyRune := make([]rune, length)
    
    for i := range keyRune {
        keyRune[i] = runes[rand.Intn(len(runes))]
    }
    
    return string(keyRune)
}

// Resolves a key and returns the corresponding URL and
// a second value to indicate if the key exists
func Resolve(key string) (url string, exists bool) {
	url, exists = lookup[key]
	return url, exists
}

// Generates a new key to be associated with the given url
// The function work in a best effort where it tries
// to randomly generate a key that is not already used.
// Should the key already exist, it retries a number of times
// and after a limit it increases the generator key length.
// By doing so, the key has a new ununsed range of possible permutations
// 
// This best-effort way of generating keys should also make it less likely
// to guess keys
func Shorten(url string) (key string) {
	
	maxTries := 2
	
	for i := 1; i <= maxTries ; i++ {
		
		key = generateKey(currentKeyLength)
		_, exists := lookup[key]
		
		// If key does not exist
		if !exists { break }

		if i == maxTries {
        	// if the key has bounced "maxTries" times,
        	// its time to expand the key length and generate a new key 
        	currentKeyLength += 1
        	key = generateKey(currentKeyLength)
        }
    }
    
    // Add the key with its associated url
    lookup[key] = url
	
	return key
}
