package tea

import "encoding/binary"

func DecryptData(data []byte, keyBytes []byte) []byte {
	if len(keyBytes) != 16 {
		panic("key must be 16 bytes (128-bit)")
	}

	key := [4]uint32{
		binary.BigEndian.Uint32(keyBytes[0:4]),
		binary.BigEndian.Uint32(keyBytes[4:8]),
		binary.BigEndian.Uint32(keyBytes[8:12]),
		binary.BigEndian.Uint32(keyBytes[12:16]),
	}

	if len(data)%8 != 0 {
		panic("invalid encrypted file")
	}

	decrypted := make([]byte, len(data))
	copy(decrypted, data)

	for i := 0; i < len(decrypted); i += 8 {
		v0 := binary.BigEndian.Uint32(decrypted[i : i+4])
		v1 := binary.BigEndian.Uint32(decrypted[i+4 : i+8])
		const delta = uint32(0x9E3779B9)
		var sum uint32 = uint32((uint64(delta) * 32) % (1 << 32))

		for range 32 {
			v1 -= ((v0 << 4) + key[2]) ^ (v0 + sum) ^ ((v0 >> 5) + key[3])
			v0 -= ((v1 << 4) + key[0]) ^ (v1 + sum) ^ ((v1 >> 5) + key[1])
			sum -= delta
		}

		binary.BigEndian.PutUint32(decrypted[i:i+4], v0)
		binary.BigEndian.PutUint32(decrypted[i+4:i+8], v1)
	}

	// Read the original file length from the first 8 bytes
	if len(decrypted) < 8 {
		panic("invalid encrypted file: too short")
	}
	originalLength := binary.BigEndian.Uint64(decrypted[0:8])

	// Extract the original data (skip the 8-byte length header)
	if uint64(len(decrypted)) < 8+originalLength {
		panic("invalid encrypted file: corrupted data")
	}

	return decrypted[8 : 8+originalLength]
}
