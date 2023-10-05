package png

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"pngme/chunks"
)

var STANDARD_HEADER [8]byte = [8]byte{137, 80, 78, 71, 13, 10, 26, 10}

type Png struct {
	Header [8]byte
	Chunks []chunks.Chunk
}

func (p *Png) GetChunkByType(typeName string) *chunks.Chunk {
	for _, chunk := range p.Chunks {
		if chunk.Type.AsString() == typeName {
			return &chunk
		}
	}

	return nil
}

func (p *Png) AppendChunk(chunk chunks.Chunk) {
	p.Chunks = append(p.Chunks, chunk)
}

func (p *Png) RemoveChunk(chunkType string) error {
	for i, chunk := range p.Chunks {
		if chunk.Type.AsString() == chunkType {
			p.Chunks = append(p.Chunks[:i], p.Chunks[i+1:]...)
			return nil

		}
	}

	return errors.New(fmt.Sprintf("Chunk with type %s not found", chunkType))
}

func (p *Png) AsBytes() []byte {
	totalBytes := []byte{}

	totalBytes = append(totalBytes, p.Header[:]...)

	for _, chunk := range p.Chunks {
		totalBytes = append(totalBytes, chunk.AsBytes()...)
	}

	return totalBytes
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

func FromPath(path string) (*Png, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return FromBytes(fileBytes)
}
