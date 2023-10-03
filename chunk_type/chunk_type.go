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
	b                [4]byte
	valid            bool
	critical         bool
	public           bool
	reservedBitValid bool
	safeToCopy       bool
}

func (c *ChunkType) IsValid() bool {
	return c.valid
}

func (c *ChunkType) IsCritical() bool {
	return c.critical
}

func (c *ChunkType) IsPublic() bool {
	return c.public
}

func (c *ChunkType) IsReservedBitValid() bool {
	return c.reservedBitValid
}

func (c *ChunkType) IsSafeToCopy() bool {
	return c.safeToCopy
}

func (c *ChunkType) Bytes() [4]byte {
	return c.b
}

func (c *ChunkType) ToString() string {
	str := ""

	for _, char := range c.b {
		str += string(rune(char))
	}

	return str
}

func ChunkTypeFromBytes(bytes [4]byte) (ChunkType, error) {
	new_chunk_type := ChunkType{}

	new_chunk_type.valid = true

	for index, b := range bytes {

		char_case, err := low_or_uppercase(b)

		if err != nil {
			new_chunk_type.valid = false

			new_chunk_type.b = [4]byte{0, 0, 0, 0}
			new_chunk_type.critical = false
			new_chunk_type.reservedBitValid = false
			new_chunk_type.public = false
			new_chunk_type.safeToCopy = false

			return new_chunk_type, errors.New("Invalid byte")
		}

		new_chunk_type.b[index] = b
		switch index {
		case 0:
			if char_case == Upper {
				new_chunk_type.critical = true
			} else {
				new_chunk_type.critical = false
			}
		case 1:
			if char_case == Upper {
				new_chunk_type.public = true
			} else {
				new_chunk_type.public = false
			}
		case 2:
			if char_case == Upper {
				new_chunk_type.reservedBitValid = true
			} else {
				new_chunk_type.reservedBitValid = false
			}
		case 3:
			if char_case == Upper {
				new_chunk_type.safeToCopy = false
			} else {
				new_chunk_type.safeToCopy = true
			}
		}

	}

	return new_chunk_type, nil
}

func ChunkTypeFromString(str string) (ChunkType, error) {
	byteArr := [4]byte{}

	for index, char := range str {
		byteArr[index] = byte(char)
	}

	return ChunkTypeFromBytes(byteArr)
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
