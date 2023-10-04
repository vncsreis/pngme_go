package chunk_type

import (
	"errors"
)

const LOWERCASE_LOW = 97
const LOWERCASE_HIGH = 122
const UPPERCASE_LOW = 65
const UPPERCASE_HIGH = 90

type Case uint8

const (
	Upper Case = 0
	Lower Case = 1
)

type ChunkType struct {
	Bytes            [4]byte
	Valid            bool
	Critical         bool
	Public           bool
	ReservedBitValid bool
	SafeToCopy       bool
}

func (c *ChunkType) ToString() string {
	str := ""

	for _, char := range c.Bytes {
		str += string(rune(char))
	}

	return str
}

func FromBytes(bytes [4]byte) (ChunkType, error) {
	new_chunk_type := ChunkType{}

	new_chunk_type.Valid = true

	for index, b := range bytes {

		char_case, err := low_or_uppercase(b)

		if err != nil {
			new_chunk_type.Valid = false

			new_chunk_type.Bytes = [4]byte{0, 0, 0, 0}
			new_chunk_type.Critical = false
			new_chunk_type.ReservedBitValid = false
			new_chunk_type.Public = false
			new_chunk_type.SafeToCopy = false

			return new_chunk_type, errors.New("Invalid byte")
		}

		new_chunk_type.Bytes[index] = b
		switch index {
		case 0:
			if char_case == Upper {
				new_chunk_type.Critical = true
			} else {
				new_chunk_type.Critical = false
			}
		case 1:
			if char_case == Upper {
				new_chunk_type.Public = true
			} else {
				new_chunk_type.Public = false
			}
		case 2:
			if char_case == Upper {
				new_chunk_type.ReservedBitValid = true
			} else {
				new_chunk_type.ReservedBitValid = false
			}
		case 3:
			if char_case == Upper {
				new_chunk_type.SafeToCopy = false
			} else {
				new_chunk_type.SafeToCopy = true
			}
		}

	}

	return new_chunk_type, nil
}

func FromString(str string) (ChunkType, error) {
	byteArr := [4]byte{}

	for index, char := range str {
		byteArr[index] = byte(char)
	}

	return FromBytes(byteArr)
}

func New(str string) (ChunkType, error) {
	return FromString(str)
}

func low_or_uppercase(num byte) (Case, error) {
	if num >= LOWERCASE_LOW && num <= LOWERCASE_HIGH {
		return Lower, nil
	} else if num >= UPPERCASE_LOW && num <= UPPERCASE_HIGH {
		return Upper, nil
	} else {
		return 2, errors.New("Invalid case number")
	}

}
