package png

import (
	"encoding/binary"
	"errors"
	"pngme/chunks"
)

var STANDARD_HEADER [8]byte = [8]byte{137, 80, 78, 71, 13, 10, 26, 10}

type Png struct {
	Header [8]byte
	Chunks []chunks.Chunk
}

func FromChunks(chunks []chunks.Chunk) Png {
	png := Png{
		Header: STANDARD_HEADER,
		Chunks: chunks,
	}

	return png
}

func FromBytes(bytes []byte) (*Png, error) {
	header := [8]byte{}

	copy(header[:], bytes[:8])
	if header != STANDARD_HEADER {
		return nil, errors.New("Invalid PNG header")
	}

	offset := 8
	totalLen := len(bytes)

	chunksArr := []chunks.Chunk{}

	for offset < totalLen {
		newChunkLen := int(binary.BigEndian.Uint32(bytes[offset : offset+4]))

		newChunk, err := chunks.FromBytes(bytes[offset : offset+12+newChunkLen])
		if err != nil {
			return nil, err
		}

		chunksArr = append(chunksArr, *newChunk)

		offset += 12 + newChunkLen

	}

	png := FromChunks(chunksArr)

	return &png, nil

}
