package service

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"path/filepath"
	"time"
)

var baseDirectory string

func generateHash(filename string) string {
	s := fmt.Sprintf("%s%s%d", time.Now().String(), filename, rand.Intn(100000))
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash)
}

func getPathFromHash(hash string) string {
	blockSize := 5
	sliceLen := len(hash) / blockSize
	path := make([]string, sliceLen+1)
	path[0] = baseDirectory
	for i := 1; i < sliceLen; i++ {
		from, to := i*blockSize, (i+1)*blockSize
		path[i] = hash[from:to]
	}
	return filepath.Join(path...)
}
