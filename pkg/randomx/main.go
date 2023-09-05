package randomx

import "math/rand"

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandString(length int) string {
	out := ""
	for i:=0; i<length;i++ {
		out += string(alphabet[int(rand.Intn(len(alphabet)))])
	}
	return out
}

func RandInt(min, max int) int {
	if min > max {
		return 0
	}
	return min + rand.Intn(max - min + 1)
}