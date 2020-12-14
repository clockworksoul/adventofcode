package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeXMasks(t *testing.T) {
	// fmt.Println(strconv.FormatInt(42, 2))
	bitmask := newBitmask("000000000000000000000000000000X1001X")

	masks := bitmask.xmasks()

	assert.Len(t, masks, 4)
	assert.Contains(t, masks, int64(0b000000))
	assert.Contains(t, masks, int64(0b000001))
	assert.Contains(t, masks, int64(0b100000))
	assert.Contains(t, masks, int64(0b100001))
}

func TestMaskAddress(t *testing.T) {
	// fmt.Println(strconv.FormatInt(42, 2))
	bitmask := newBitmask("000000000000000000000000000000X1001X")

	addrs := bitmask.maskAddress(42)

	t.Log(addrs)

	assert.Len(t, addrs, 4)
	assert.Contains(t, addrs, int64(26))
	assert.Contains(t, addrs, int64(27))
	assert.Contains(t, addrs, int64(58))
	assert.Contains(t, addrs, int64(59))
}
