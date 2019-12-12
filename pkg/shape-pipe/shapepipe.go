package shapepipe

import (
	"io"
)

type Shape []byte

const emptyByte = byte(' ')
const newlineByte = byte('\n')
const tabByte = byte('\t')

type ShapeReader struct {
	Shape           Shape
	Reader          io.Reader
	currentPosition int
	lastByte        byte
}

func (sw *ShapeReader) Read(p []byte) (int, error) {
	read := make([]byte, 1)

	for i := range p {
		switch sw.Shape[sw.currentPosition%len(sw.Shape)] {
		case emptyByte:
			p[i] = emptyByte
		case newlineByte:
			p[i] = newlineByte
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
				// Tabs will mess with the printing, replace it with a space
				if read[0] == tabByte {
					sw.lastByte = emptyByte
					p[i] = emptyByte
					break
				}
				// Don't print multiple spaces right after another
				if sw.lastByte == emptyByte && read[0] == emptyByte {
					continue
				}

				sw.lastByte = read[0]
				p[i] = read[0]
				break
			}
		}
		sw.currentPosition++
	}

	return len(p), nil
}
