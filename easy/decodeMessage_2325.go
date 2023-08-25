package easy

import "bytes"

func DecodeMessage(key string, message string) string {
	convert := map[byte]byte{}
	alpha := 'a'
	for i := 0; i < len(key); i++ {
		b := byte(key[i])
		if _, exsit := convert[b]; b != ' ' && !exsit {
			convert[b] = byte(alpha)
			alpha++
		}
	}
	var buffer bytes.Buffer
	for i := 0; i < len(message); i++ {
		b := byte(message[i])
		if _, exsit := convert[b]; b != ' ' && exsit {
			buffer.WriteByte(convert[b])
		} else {
			buffer.WriteByte(b)
		}

	}
	return buffer.String()
}
