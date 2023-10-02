package chunks

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
	b                [4]int
	valid            bool
	critical         bool
	public           bool
	reservedBitValid bool
	safeToCopy       bool
}

func (c *ChunkType) isValid() bool {
	return c.valid
}

func (c *ChunkType) isCritical() bool {
	return c.critical
}

func (c *ChunkType) isPublic() bool {
	return c.public
}

func (c *ChunkType) isReservedBitValid() bool {
	return c.reservedBitValid
}

func (c *ChunkType) isSafeToCopy() bool {
	return c.safeToCopy
}

func (c *ChunkType) bytes() [4]int {
	return c.b
}

func ChunkTypeFromBytes(bytes [4]int) (ChunkType, error) {
	new_chunk_type := ChunkType{}

	new_chunk_type.valid = true

	for index, byte := range bytes {

		char_case, err := low_or_uppercase(byte)

		if err != nil {
			new_chunk_type.valid = false

			new_chunk_type.b = [4]int{0, 0, 0, 0}
			new_chunk_type.critical = false
			new_chunk_type.reservedBitValid = false
			new_chunk_type.public = false
			new_chunk_type.safeToCopy = false

			return new_chunk_type, errors.New("Invalid byte")
		}

		new_chunk_type.b[index] = byte
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
	int_arr := [4]int{}

	for index, char := range str {
		int_arr[index] = int(char)
	}

	return ChunkTypeFromBytes(int_arr)
}

func low_or_uppercase(num int) (Case, error) {
	if num >= LOWERCASE_LOW && num <= LOWERCASE_HIGH {
		return Lower, nil
	} else if num >= UPPERCASE_LOW && num <= UPPERCASE_HIGH {
		return Upper, nil
	} else {
		return 2, errors.New("Invalid case number")
	}
}
