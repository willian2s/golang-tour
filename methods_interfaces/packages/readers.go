package packages

import (
	"fmt"
	"strings"
)

func Reader(e string) {
	reader := strings.NewReader(e)

	bytes := make([]byte, 8)

	for {
		n, err := reader.Read(bytes)

		fmt.Printf("n = %v, err =  %v, byte = %v\n", n, err, bytes)
		fmt.Printf("bytes[:n] = %q\n", bytes[:n])
		if err != nil {
			break
		}
	}
}
