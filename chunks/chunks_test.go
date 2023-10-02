package chunks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunkTypeFromBytes(t *testing.T) {
	expected := [4]int{82, 117, 83, 116}

	actual, _ := ChunkTypeFromBytes([4]int{82, 117, 83, 116})

	assert.Equal(t, expected, actual.bytes(), "Should be equal")

}

func TestChunkTypeFromString(t *testing.T) {
	expected := [4]int{82, 117, 83, 116}

	actual, _ := ChunkTypeFromString("RuSt")

	assert.Equal(t, expected, actual.bytes(), "Should be equal")
}

func TestChunkTypeIsCritical(t *testing.T) {
	chunk, _ := ChunkTypeFromString("RuSt")
	assert.True(t, chunk.isCritical(), "isCritical() should be true")
}

func TestChunkTypeIsNotCritical(t *testing.T) {
	chunk, _ := ChunkTypeFromString("ruSt")
	assert.False(t, chunk.isCritical(), "isCritical() should be false")
}

func TestChunkTypeIsPublic(t *testing.T) {
	chunk, _ := ChunkTypeFromString("RUSt")
	assert.True(t, chunk.isPublic(), "isPublic() should be true")
}

func TestChunkTypeIsNotPublic(t *testing.T) {
	chunk, _ := ChunkTypeFromString("RuSt")
	assert.False(t, chunk.isPublic(), "isPublic() should be false")
}

func TestChunkTypeIsReservedBitValid(t *testing.T) {
	chunk, _ := ChunkTypeFromString("RUSt")
	assert.True(t,
		chunk.isReservedBitValid(),
		"isReservedBitValid() should be true",
	)
}

func TestChunkTypeIsNotReservedBitValid(t *testing.T) {
	chunk, _ := ChunkTypeFromString("Rust")
	assert.False(t,
		chunk.isReservedBitValid(),
		"isReservedBitValid() should be false",
	)
}

func TestChunkTypeIsSafeToCopy(t *testing.T) {
	chunk, _ := ChunkTypeFromString("RUSt")
	assert.True(t, chunk.isSafeToCopy(), "isSafeToCopy() should be true")
}

func TestChunkTypeIsUnsafeToCopy(t *testing.T) {
	chunk, _ := ChunkTypeFromString("RuST")
	assert.False(t, chunk.isSafeToCopy(), "isSafeToCopy() should be false")
}

func TestValidChunkTypeIsValid(t *testing.T) {
	chunk, _ := ChunkTypeFromString("RUSt")
	assert.True(t, chunk.isValid(), "isValid() should be true")
}

func TestInvalidChunkTypeIsNotValid(t *testing.T) {
	chunk, _ := ChunkTypeFromString("Ru1t")
	assert.False(t, chunk.isValid(), "isValid() should be false")
}
