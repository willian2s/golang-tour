package packages

import "io"

type Rot13Reader struct {
	R io.Reader
}

func (rr Rot13Reader) Read(p []byte) (int, error) {
	n, err := rr.R.Read(p)
	for i := 0; i < n; i++ {
		if (p[i] >= 'A' && p[i] <= 'M') || (p[i] >= 'a' && p[i] <= 'm') {
			p[i] += 13
		} else if (p[i] >= 'N' && p[i] <= 'Z') || (p[i] >= 'n' && p[i] <= 'z') {
			p[i] -= 13
		}
	}
	return n, err
}
