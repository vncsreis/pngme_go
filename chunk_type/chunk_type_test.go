package chunk_type

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChunkTypeFromBytes(t *testing.T) {
	expected := [4]byte{82, 117, 83, 116}

	actual, _ := ChunkTypeFromBytes([4]byte{82, 117, 83, 116})

	assert.Equal(t, expected, actual.Bytes(), "Should be equal")

}

func TestChunkTypeFromString(t *testing.T) {
	expected := [4]byte{82, 117, 83, 116}

	actual, _ := ChunkTypeFromString("RuSt")

	assert.Equal(t, expected, actual.Bytes(), "Should be equal")
}

func TestChunkTypeIsCritical(t *testing.T) {
	chunk, _ := ChunkTypeFromString("RuSt")
	assert.True(t, chunk.IsCritical(), "isCritical() should be true")
}

func TestChunkTypeIsNotCritical(t *testing.T) {
	chunk, _ := ChunkTypeFromString("ruSt")
	assert.False(t, chunk.IsCritical(), "isCritical() should be false")
}

func TestChunkTypeIsPublic(t *testing.T) {
	chunk, _ := ChunkTypeFromString("RUSt")
	assert.True(t, chunk.IsPublic(), "isPublic() should be true")
}

func TestChunkTypeIsNotPublic(t *testing.T) {
	chunk, _ := ChunkTypeFromString("RuSt")
	assert.False(t, chunk.IsPublic(), "isPublic() should be false")
}

func TestChunkTypeIsReservedBitValid(t *testing.T) {
	chunk, _ := ChunkTypeFromString("RUSt")
	assert.True(t,
		chunk.IsReservedBitValid(),
		"isReservedBitValid() should be true",
	)
}

func TestChunkTypeIsNotReservedBitValid(t *testing.T) {
	chunk, _ := ChunkTypeFromString("Rust")
	assert.False(t,
		chunk.IsReservedBitValid(),
		"isReservedBitValid() should be false",
	)
}

func TestChunkTypeIsSafeToCopy(t *testing.T) {
	chunk, _ := ChunkTypeFromString("RUSt")
	assert.True(t, chunk.IsSafeToCopy(), "isSafeToCopy() should be true")
}

func TestChunkTypeIsUnsafeToCopy(t *testing.T) {
	chunk, _ := ChunkTypeFromString("RuST")
	assert.False(t, chunk.IsSafeToCopy(), "isSafeToCopy() should be false")
}

func TestValidChunkTypeIsValid(t *testing.T) {
	chunk, _ := ChunkTypeFromString("RUSt")
	assert.True(t, chunk.IsValid(), "isValid() should be true")
}

func TestInvalidChunkTypeIsNotValid(t *testing.T) {
	chunk, err := ChunkTypeFromString("Ru1t")
	assert.False(t, chunk.IsValid(), "isValid() should be false")
	assert.NotNil(t, err, "err should be not nil")
}

func TestChunkTypeToString(t *testing.T) {
	chunk, _ := ChunkTypeFromString("RuSt")
	assert.Equal(t, "RuSt", chunk.toString(), "String should be \"RuSt\"")
}

func TestChunkEqual(t *testing.T) {
	chunk1, _ := ChunkTypeFromString("RuSt")
	chunk2, _ := ChunkTypeFromString("RuSt")
	chunk3, _ := ChunkTypeFromString("RUSt")

	assert.Equal(t, chunk1, chunk2, "chunk1 and chunk2 should be equal")
	assert.NotEqual(t, chunk1, chunk3, "chunk1 and chunk2 should not be equal")
}
