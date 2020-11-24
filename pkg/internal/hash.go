package internal

import (
	"fmt"
	"hash/fnv"
)

func Hash(s string) string {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	if err != nil {
		return s
	}
	return fmt.Sprintf("%d", h.Sum32())
}
