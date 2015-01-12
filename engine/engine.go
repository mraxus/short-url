package engine

import (
	"log"
	"math/rand"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

var currentHashGenerationLength int = 3
var lookup map[string]string

func init() {
	log.Println("engine init()")
	
	lookup = make(map[string]string)
}

func generateHash(length int) string {
	
	hashRune := make([]rune, length)
    
    for i := range hashRune {
        hashRune[i] = runes[rand.Intn(len(runes))]
    }
    
    return string(hashRune)
}

func Resolve(hash string) (url string, exists bool) {
	url, exists = lookup[hash]
	return url, exists
}

func Shorten(url string) (hash string) {
	
	maxTries := 4
	
	for i := 1; i <= maxTries ; i++ {
		
		hash = generateHash(currentHashGenerationLength)
		_, exists := lookup[hash]
		
		// If hash does not exist
		if !exists { break }

		if i == maxTries {
        	// if the hash has bounced "maxTries" times,
        	// its time to expand the hash length and generate a new hash 
        	currentHashGenerationLength += 1
        	hash = generateHash(currentHashGenerationLength)
        }
    }
    
    // Add the hash with its associated url
    lookup[hash] = url
	
	return hash
}