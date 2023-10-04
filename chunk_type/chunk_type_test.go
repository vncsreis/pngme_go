package chunk_type

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChunkTypeFromBytes(t *testing.T) {
	expected := [4]byte{82, 117, 83, 116}

	actual, _ := FromBytes([4]byte{82, 117, 83, 116})

	assert.Equal(t, expected, actual.Bytes, "Should be equal")

}

func TestChunkTypeFromString(t *testing.T) {
	expected := [4]byte{82, 117, 83, 116}

	actual, _ := FromString("RuSt")

	assert.Equal(t, expected, actual.Bytes, "Should be equal")
}

func TestChunkTypeIsCritical(t *testing.T) {
	chunk, _ := FromString("RuSt")
	assert.True(t, chunk.Critical, "isCritical() should be true")
}

func TestChunkTypeIsNotCritical(t *testing.T) {
	chunk, _ := FromString("ruSt")
	assert.False(t, chunk.Critical, "isCritical() should be false")
}

func TestChunkTypeIsPublic(t *testing.T) {
	chunk, _ := FromString("RUSt")
	assert.True(t, chunk.Public, "isPublic() should be true")
}

func TestChunkTypeIsNotPublic(t *testing.T) {
	chunk, _ := FromString("RuSt")
	assert.False(t, chunk.Public, "isPublic() should be false")
}

func TestChunkTypeIsReservedBitValid(t *testing.T) {
	chunk, _ := FromString("RUSt")
	assert.True(t,
		chunk.ReservedBitValid,
		"ReservedBitValid should be true",
	)
}

func TestChunkTypeIsNotReservedBitValid(t *testing.T) {
	chunk, _ := FromString("Rust")
	assert.False(t,
		chunk.ReservedBitValid,
		"isReservedBitValid() should be false",
	)
}

func TestChunkTypeIsSafeToCopy(t *testing.T) {
	chunk, _ := FromString("RUSt")
	assert.True(t, chunk.SafeToCopy, "isSafeToCopy() should be true")
}

func TestChunkTypeIsUnsafeToCopy(t *testing.T) {
	chunk, _ := FromString("RuST")
	assert.False(t, chunk.SafeToCopy, "isSafeToCopy() should be false")
}

func TestValidChunkTypeIsValid(t *testing.T) {
	chunk, _ := FromString("RUSt")
	assert.True(t, chunk.Valid, "isValid() should be true")
}

func TestInvalidChunkTypeIsNotValid(t *testing.T) {
	chunk, err := FromString("Ru1t")
	assert.False(t, chunk.Valid, "isValid() should be false")
	assert.NotNil(t, err, "err should be not nil")
}

func TestChunkTypeToString(t *testing.T) {
	chunk, _ := FromString("RuSt")
	assert.Equal(t, "RuSt", chunk.ToString(), "String should be \"RuSt\"")
}

func TestChunkEqual(t *testing.T) {
	chunk1, _ := FromString("RuSt")
	chunk2, _ := FromString("RuSt")
	chunk3, _ := FromString("RUSt")

	assert.Equal(t, chunk1, chunk2, "chunk1 and chunk2 should be equal")
	assert.NotEqual(t, chunk1, chunk3, "chunk1 and chunk2 should not be equal")
}
