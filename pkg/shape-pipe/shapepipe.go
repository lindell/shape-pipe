package shapepipe

import (
	"io"
)

type Shape []byte

const emptyByte = byte(' ')
const newlineByte = byte('\n')
const tabByte = byte('	')

type ShapeReader struct {
	Shape           Shape
	Reader          io.Reader
	currentPosition int
}

func (sw *ShapeReader) Read(p []byte) (int, error) {
	read := make([]byte, 1)

	for i := range p {
		switch sw.Shape[sw.currentPosition%len(sw.Shape)] {
		case emptyByte:
			p[i] = emptyByte
		case newlineByte:
			p[i] = newlineByte
		case tabByte:
			p[i] = emptyByte
		default:
			for {
				n, err := sw.Reader.Read(read)
				// If the underlying reader does not give us any data, stop the read and return any potential error (might be nil)
				// We should also return if the underlying reader errors
				if n == 0 || err != nil {
					return i, err
				}

				// We don't want to print newlines from the underlying reader
				if read[0] == newlineByte {
					continue
				}

				p[i] = read[0]
				break
			}
		}
		sw.currentPosition++
	}

	return len(p), nil
}
