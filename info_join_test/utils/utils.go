package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

//ToBytes encode i to slice of bytes
func ToBytes(i interface{}) []byte {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	HandleError(encoder.Encode(i))
	return buffer.Bytes()
}

//HandleError handles error
func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
