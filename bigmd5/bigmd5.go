package bigmd5

import (
	"bufio"
	"crypto/md5"
	"io"
	"os"
)

const bufferSize = 65536

// MD5File returns the MD5 checksum of the file
func MD5File(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	hash := md5.New()
	for buf, reader := make([]byte, bufferSize), bufio.NewReader(file); ; {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		hash.Write(buf[:n])
	}

	return hash.Sum(nil), nil
}
