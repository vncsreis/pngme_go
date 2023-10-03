package chunks

import (
	"encoding/binary"
	"fmt"
	"pngme/chunk_type"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateTestChunk() Chunk {
	dataLen := 42
	chunkTypeAsBytes := []byte("RuSt")
	messageAsBytes := []byte("This is where your secret message will be!")
	crc := 2882656334

	lenAsBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lenAsBytes, uint32(dataLen))

	crcAsBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(crcAsBytes, uint32(crc))

	amnt := 12 + dataLen

	chunkDataAsBytes := make([]byte, amnt)

	copy(chunkDataAsBytes[0:4], lenAsBytes)
	copy(chunkDataAsBytes[4:8], chunkTypeAsBytes)
	copy(chunkDataAsBytes[8:8+dataLen], messageAsBytes)
	copy(chunkDataAsBytes[8+dataLen:], crcAsBytes)

	return ChunkFromBytes(chunkDataAsBytes)
}

func TestNewChunk(t *testing.T) {
	chunkType, _ := chunk_type.ChunkTypeFromString("RuSt")
	data := []byte("This is where your secret message will be!")
	chunk := ChunkNew(chunkType, data)

	assert.Equal(t, 42, chunk.length(), "Chunk length should be 42")
	assert.Equal(t, 2882656334, chunk.crc(), "Chunk CRC should be 2882656334")
}

func TestChunkLenght(t *testing.T) {
	chunk := CreateTestChunk()
	assert.Equal(t, 42, chunk.length(), "Chunk data length should be 42")
}

func TestChunkType(t *testing.T) {
	chunk := CreateTestChunk()
	assert.Equal(t, "RuSt", chunk.chunkType().ToString(), "Chunk type should be \"RuSt\"")
}

func TestChunkString(t *testing.T) {
	chunk := CreateTestChunk()
	chunkString := chunk.dataAsString()
	excpectedChunkString := "This is where your secret message will be!"

	assert.Equal(t,
		excpectedChunkString,
		chunkString,
		fmt.Sprintf("Chunk string should be \"%s\"", excpectedChunkString),
	)
}

func TestChunkCrc(t *testing.T) {
	chunk := CreateTestChunk()
	assert.Equal(t, 2882656334, chunk.crc(), "Chunk Crc should be 2882656334")
}

func TestValidChunkFromBytes(t *testing.T) {
	dataLen := 42
	chunkTypeAsBytes := []byte("RuSt")
	messageAsBytes := []byte("This is where your secret message will be!")
	crc := 2882656334

	lenAsBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lenAsBytes, uint32(dataLen))

	crcAsBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(crcAsBytes, uint32(crc))

	amnt := 12 + dataLen

	chunkDataAsBytes := make([]byte, amnt)

	copy(chunkDataAsBytes[0:4], lenAsBytes)
	copy(chunkDataAsBytes[4:8], chunkTypeAsBytes)
	copy(chunkDataAsBytes[8:8+dataLen], messageAsBytes)
	copy(chunkDataAsBytes[8+dataLen:], crcAsBytes)

	chunk := ChunkFromBytes(chunkDataAsBytes)

	chunkString := chunk.dataAsString()
	expectedChunkString := "This is where your secret message will be!"

	assert.Equal(t, 42, chunk.length(), "Chunk length should be 42")
	assert.Equal(t, "RuSt", chunk.chunkType().ToString(), "Chunk type should be \"RuSt\"")
	assert.Equal(t, expectedChunkString, chunkString, fmt.Sprintf("Chunk message should be \"%s\"", expectedChunkString))
	assert.Equal(t, crc, chunk.crc(), fmt.Sprintf("Chunk Crc should be %d", crc))
}

// TODO: Test invalid chunks
