package chunks

import (
	"encoding/binary"
	"hash/crc32"
	"pngme/chunk_type"
)

type Chunk struct {
	cLen  uint32
	cType chunk_type.ChunkType
	cData []byte
	cCrc  uint32
}

func (c *Chunk) dataAsString() string {
	return string(c.cData)
}

func (c *Chunk) asBytes() []byte {
	return c.cData
}

func (c *Chunk) length() int {
	return int(c.cLen)
}

func (c *Chunk) chunkType() *chunk_type.ChunkType {
	return &c.cType
}

func (c *Chunk) data() *[]byte {
	return &c.cData
}

func (c *Chunk) crc() int {
	return int(c.cCrc)
}
func ChunkNew(chunkType chunk_type.ChunkType, data []byte) Chunk {
	newChunk := Chunk{}

	dataLen := uint32(len(data))
	newChunk.cLen = dataLen

	newChunk.cType = chunkType
	newChunk.cData = data
	newChunk.cCrc = chunkCrc(chunkType, data)

	return newChunk
}

func chunkCrc(chunkType chunk_type.ChunkType, data []byte) uint32 {
	dataLen := uint32(len(data))
	crcTarget := make([]byte, 4+dataLen)
	chunkTypeAsBytes := chunkType.Bytes()

	copy(crcTarget[0:4], chunkTypeAsBytes[:])
	copy(crcTarget[4:], data)

	crcAsBytes := make([]byte, 4)

	binary.BigEndian.PutUint32(crcAsBytes, crc32.ChecksumIEEE(crcTarget))

	return binary.BigEndian.Uint32(crcAsBytes)
}

func ChunkFromBytes(data []byte) Chunk {
	dataLenB := make([]byte, 4)
	copy(dataLenB, data[0:4])
	dataLen := binary.BigEndian.Uint32(dataLenB)

	chunkTypeB := make([]byte, 4)
	copy(chunkTypeB, data[4:8])
	chunkType, err := chunk_type.ChunkTypeFromBytes([4]byte(chunkTypeB))
	if err != nil {
		panic("error")
	}

	dataB := make([]byte, dataLen)
	copy(dataB, data[8:dataLen+8])

	crcB := make([]byte, 4)
	copy(crcB, data[dataLen+8:])
	crc := binary.BigEndian.Uint32(crcB)

	newChunk := Chunk{}

	newChunk.cLen = dataLen
	newChunk.cType = chunkType
	newChunk.cData = dataB
	newChunk.cCrc = crc

	return newChunk
}

// func ChunkNew(chunkType chunk_type.ChunkType, data []byte) Chunk {
// 	dataLen := uint32(len(data))
// 	dataLenAsBytes := make([]byte, 4)
// 	binary.BigEndian.PutUint32(dataLenAsBytes, dataLen)

// 	crcTarget := make([]byte, 4+dataLen)
// 	chunkTypeAsBytes := chunkType.Bytes()

// 	copy(crcTarget[0:4], chunkTypeAsBytes[:])
// 	copy(crcTarget[4:], data)

// 	crcAsBytes := make([]byte, 4)

// 	binary.BigEndian.PutUint32(crcAsBytes, crc32.ChecksumIEEE(crcTarget))

// 	chunkData := make([]byte, 12+dataLen)

// 	copy(chunkData[0:4], dataLenAsBytes)
// 	copy(chunkData[4:8], chunkTypeAsBytes[:])
// 	copy(chunkData[8:8+dataLen], data)
// 	copy(chunkData[8+dataLen:], crcAsBytes)

// }
