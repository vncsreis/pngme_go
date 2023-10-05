package png_test

import (
	"os"
	"pngme/chunk_type"
	"pngme/chunks"
	"pngme/png"
	"testing"

	"github.com/stretchr/testify/assert"
)

const PNG_PATH string = "../tux_png.png"

func createTestingChunks() []chunks.Chunk {

	chunk1, _ := chunkFromStrings("FrSt", "I am the first chunk")
	chunk2, _ := chunkFromStrings("miDl", "I am another chunk")
	chunk3, _ := chunkFromStrings("LASt", "I am the last chunk")

	testingChunks := []chunks.Chunk{}

	testingChunks = append(testingChunks, *chunk1)
	testingChunks = append(testingChunks, *chunk2)
	testingChunks = append(testingChunks, *chunk3)

	return testingChunks
}

func chunkFromStrings(chunkTypeString string, dataString string) (*chunks.Chunk, error) {
	chunkType, err := chunk_type.FromString(chunkTypeString)
	if err != nil {
		return nil, err
	}

	data := []byte(dataString)

	chunk := chunks.New(chunkType, data)

	return &chunk, nil

}

func createTestingPng() png.Png {
	chunks := createTestingChunks()
	newPng := png.FromChunks(chunks)

	return newPng
}

func chunksToByteArray(chunkArr []chunks.Chunk) []byte {
	chunkBytes := []byte{}

	chunkBytes = append(chunkBytes, png.STANDARD_HEADER[:]...)
	for i := 0; i < len(chunkArr); i++ {

		chunkBytes = append(chunkBytes, chunkArr[i].AsBytes()...)
	}

	return chunkBytes

}

func TestFromChunks(t *testing.T) {
	chunks := createTestingChunks()
	newPng := png.FromChunks(chunks)

	assert.Equal(t, 3, len(newPng.Chunks), "Png chunk array len should be 3")
}

func TestValidFromBytes(t *testing.T) {
	testingChunks := createTestingChunks()
	chunkBytes := chunksToByteArray(testingChunks)

	newPng, err := png.FromBytes(chunkBytes)
	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, newPng, "newPng Should be not nil")

}

func TestInvalidHeader(t *testing.T) {
	testingChunks := createTestingChunks()
	chunkBytes := chunksToByteArray(testingChunks)

	badBytes := []byte{13, 80, 78, 71, 13, 10, 26, 10}

	badBytes = append(badBytes, chunkBytes...)

	newPng, err := png.FromBytes(badBytes)
	assert.NotNil(t, err, "err should be not nil")
	assert.Nil(t, newPng, "newPng should be nil")
}

func TestInvalidChunk(t *testing.T) {
	testingChunks := createTestingChunks()
	chunkBytes := chunksToByteArray(testingChunks)

	bytes := []byte{}

	bytes = append(bytes, chunkBytes...)

	badChunk := []byte{
		0, 0, 0, 5, // length
		32, 117, 83, 116, // Chunk Type (bad)
		65, 64, 65, 66, 67, // Data
		1, 2, 3, 4, // CRC (bad)
	}

	bytes = append(bytes, badChunk...)

	newPng, err := png.FromBytes(bytes)

	assert.NotNil(t, err, "err should be not nil")
	assert.Nil(t, newPng, "newPng should be nil")
}

func TestListChunks(t *testing.T) {
	testPng := createTestingPng()
	pngChunks := testPng.Chunks

	assert.Equal(t, 3, len(pngChunks), "Chunk list len should be 3")
}

func TestChunkByType(t *testing.T) {
	testPng := createTestingPng()
	testChunk := testPng.GetChunkByType("FrSt")

	assert.NotNil(t, testChunk, "testChunk should not be nil")
	assert.Equal(t,
		"I am the first chunk",
		testChunk.DataAsString(),
		"testChunk data string should be \"I am the first chunk\"",
	)
}

func TestAppendChunk(t *testing.T) {
	testPng := createTestingPng()
	newChunk, _ := chunks.FromStrings("TeSt", "Message")
	testPng.AppendChunk(*newChunk)
	foundChunk := testPng.GetChunkByType("TeSt")

	assert.NotNil(t, foundChunk, "foundChunk should not be nil")
	assert.Equal(t,
		4,
		len(testPng.Chunks),
		"testPng Chunk len should be 4",
	)
	assert.Equal(t,
		"TeSt",
		foundChunk.Type.AsString(),
		"foundChunk type should be \"TeSt\"",
	)
	assert.Equal(t,
		"Message",
		foundChunk.DataAsString(),
		"foundChunk data should be \"Message\"",
	)

}

func TestRemoveChunk(t *testing.T) {
	testPng := createTestingPng()
	testChunk, _ := chunks.FromStrings("TeSt", "Message")
	testPng.AppendChunk(*testChunk)

	err := testPng.RemoveChunk("TeSt")
	assert.Nil(t, err, "err should be nil")

	foundChunk := testPng.GetChunkByType("TeSt")
	assert.Nil(t, foundChunk, "foundChunk should be nil")

}

func TestFromPath(t *testing.T) {
	testPng, err := png.FromPath(PNG_PATH)

	assert.Nil(t, err, "err should be nil")
	assert.NotNil(t, testPng, "testPng should not be nil")
}

func TestAsBytes(t *testing.T) {
	testPng, _ := png.FromPath(PNG_PATH)
	actual, _ := os.ReadFile(PNG_PATH)

	expected := testPng.AsBytes()

	assert.Equal(t, expected, actual, "expected and actual should be equal")
}
