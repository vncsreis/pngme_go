package chunks_test

import (
	"encoding/binary"
	"fmt"
	"pngme/chunk_type"
	"pngme/chunks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateTestChunk() (*chunks.Chunk, error) {
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

	chunk, err := chunks.FromBytes(chunkDataAsBytes)
	if err != nil {
		return nil, err
	}

	return chunk, nil
}

func TestNewChunk(t *testing.T) {
	chunkType, _ := chunk_type.FromString("RuSt")
	data := []byte("This is where your secret message will be!")
	chunk := chunks.New(chunkType, data)

	assert.Equal(t, 42, chunk.Len, "Chunk length should be 42")
	assert.Equal(t, 2882656334, chunk.Crc, "Chunk CRC should be 2882656334")
}

func TestChunkLenght(t *testing.T) {
	chunk, _ := CreateTestChunk()
	assert.Equal(t, 42, chunk.Len, "Chunk data length should be 42")
}

func TestChunkType(t *testing.T) {
	chunk, _ := CreateTestChunk()
	assert.Equal(t, "RuSt", chunk.Type.ToString(), "Chunk type should be \"RuSt\"")
}

func TestChunkString(t *testing.T) {
	chunk, _ := CreateTestChunk()
	chunkString := chunk.DataAsString()
	excpectedChunkString := "This is where your secret message will be!"

	assert.Equal(t,
		excpectedChunkString,
		chunkString,
		fmt.Sprintf("Chunk string should be \"%s\"", excpectedChunkString),
	)
}

func TestChunkCrc(t *testing.T) {
	chunk, _ := CreateTestChunk()
	assert.Equal(t, 2882656334, chunk.Crc, "Chunk Crc should be 2882656334")
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

	chunk, _ := chunks.FromBytes(chunkDataAsBytes)

	chunkString := chunk.DataAsString()
	expectedChunkString := "This is where your secret message will be!"

	assert.Equal(t, 42, chunk.Len, "Chunk length should be 42")
	assert.Equal(t, "RuSt", chunk.Type.ToString(), "Chunk type should be \"RuSt\"")
	assert.Equal(t, expectedChunkString, chunkString, fmt.Sprintf("Chunk message should be \"%s\"", expectedChunkString))
	assert.Equal(t, crc, chunk.Crc, fmt.Sprintf("Chunk Crc should be %d", crc))
}

func TestInvalidChunkFromBytes(t *testing.T) {
	dataLen := 42
	chunkTypeAsBytes := []byte("RuSt")
	messageAsBytes := []byte("This is where your secret message will be!")
	crc := 2882656333

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

	chunk, err := chunks.FromBytes(chunkDataAsBytes)

	assert.NotNil(t, err, "error should not be nil")
	assert.Nil(t, chunk, "chunk should be nil")
}

func TestChunkToBytes(t *testing.T) {
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

	chunkType, _ := chunk_type.New("RuSt")
	newChunk := chunks.New(chunkType, []byte("This is where your secret message will be!"))

	newChunkAsBytes := newChunk.AsBytes()

	assert.Equal(t, chunkDataAsBytes, newChunkAsBytes, "chunkDataAsBytes and newChunkAsBytes should be equal")
}
