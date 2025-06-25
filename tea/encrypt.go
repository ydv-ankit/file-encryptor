package tea

import (
	"encoding/binary"
)

func EncryptData(data []byte, keyBytes []byte) []byte {
	if len(keyBytes) != 16 {
		panic("key must be 16 bytes (128-bit)")
	}

	// Convert key bytes to 4 uint32 values
	key := [4]uint32{
		binary.BigEndian.Uint32(keyBytes[0:4]),
		binary.BigEndian.Uint32(keyBytes[4:8]),
		binary.BigEndian.Uint32(keyBytes[8:12]),
		binary.BigEndian.Uint32(keyBytes[12:16]),
	}

	// Store original length at the beginning (8 bytes)
	originalLength := uint64(len(data))
	lengthBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(lengthBytes, originalLength)

	// Combine length and data
	dataWithLength := append(lengthBytes, data...)

	// Pad data to be multiple of 8 bytes
	if len(dataWithLength)%8 != 0 {
		padding := make([]byte, 8-len(dataWithLength)%8)
		dataWithLength = append(dataWithLength, padding...)
	}

	encrypted := make([]byte, len(dataWithLength))
	copy(encrypted, dataWithLength)

	const delta = 0x9E3779B9

	for i := 0; i < len(encrypted); i += 8 {
		v0 := binary.BigEndian.Uint32(encrypted[i : i+4])
		v1 := binary.BigEndian.Uint32(encrypted[i+4 : i+8])
		var sum uint32 = 0

		for range 32 {
			sum += delta
			v0 += ((v1 << 4) + key[0]) ^ (v1 + sum) ^ ((v1 >> 5) + key[1])
			v1 += ((v0 << 4) + key[2]) ^ (v0 + sum) ^ ((v0 >> 5) + key[3])
		}

		binary.BigEndian.PutUint32(encrypted[i:i+4], v0)
		binary.BigEndian.PutUint32(encrypted[i+4:i+8], v1)
	}

	return encrypted
}
