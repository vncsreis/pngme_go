package chunks

import (
	"encoding/binary"
	"errors"
	"hash/crc32"
	"pngme/chunk_type"
)

type Chunk struct {
	Len  int
	Type *chunk_type.ChunkType
	Data []byte
	Crc  int
}

func (c *Chunk) AsBytes() []byte {
	byteArr := make([]byte, 12+c.Len)

	lenBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lenBytes, uint32(c.Len))
	copy(byteArr[0:4], lenBytes)
	copy(byteArr[4:8], c.Type.Bytes[:])
	copy(byteArr[8:8+c.Len], c.Data)

	crcBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(crcBytes, uint32(c.Crc))
	copy(byteArr[8+c.Len:], crcBytes)

	return byteArr
}

func (c *Chunk) DataAsString() string {
	return string(c.Data)
}

func New(chunkType chunk_type.ChunkType, data []byte) Chunk {
	newChunk := Chunk{}

	newChunk.Len = len(data)

	newChunk.Type = &chunkType
	newChunk.Data = data
	newChunk.Crc = Crc(chunkType, data)

	return newChunk
}

func Crc(chunkType chunk_type.ChunkType, data []byte) int {
	dataLen := len(data)
	crcTarget := make([]byte, 4+dataLen)

	copy(crcTarget[0:4], chunkType.Bytes[:])
	copy(crcTarget[4:], data)

	crcAsBytes := make([]byte, 4)

	binary.BigEndian.PutUint32(crcAsBytes, crc32.ChecksumIEEE(crcTarget))

	return int(binary.BigEndian.Uint32(crcAsBytes))
}

func FromBytes(data []byte) (*Chunk, error) {
	dataLenB := make([]byte, 4)
	copy(dataLenB, data[0:4])
	dataLen := int(binary.BigEndian.Uint32(dataLenB))

	chunkTypeB := make([]byte, 4)
	copy(chunkTypeB, data[4:8])
	chunkType, err := chunk_type.FromBytes([4]byte(chunkTypeB))
	if err != nil {
		return nil, err
	}

	dataB := make([]byte, dataLen)
	copy(dataB, data[8:dataLen+8])

	chunkCrcTest := Crc(chunkType, dataB)

	crcB := make([]byte, 4)
	copy(crcB, data[dataLen+8:])
	crc := int(binary.BigEndian.Uint32(crcB))

	if chunkCrcTest != crc {
		return nil, errors.New("Invalid checksum")
	}

	newChunk := Chunk{}

	newChunk.Len = dataLen
	newChunk.Type = &chunkType
	newChunk.Data = dataB
	newChunk.Crc = crc

	return &newChunk, nil
}

func FromStrings(typeString string, dataString string) (*Chunk, error) {
	chunkType, err := chunk_type.FromString(typeString)
	if err != nil {
		return nil, err
	}

	newChunk := New(chunkType, []byte(dataString))

	return &newChunk, nil
}
