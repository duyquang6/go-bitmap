package bitmap

import (
	"errors"
)

// hardcode 8bit
const BYTE_SIZE = 8

type Bitmap struct {
	data []byte
	size uint64
}

func NewBitmap(size uint64) *Bitmap {
	sizeArr := size / BYTE_SIZE
	if size%64 != 0 {
		sizeArr++
	}
	arr := make([]byte, sizeArr)
	return &Bitmap{size: size, data: arr}
}

func (b *Bitmap) Set(position uint64, val bool) error {
	if position >= b.size {
		return errors.New("invalid index, out of range")
	}
	arrOffset := position / BYTE_SIZE

	if val {
		b.data[arrOffset] |= byte(1 << (position % BYTE_SIZE))
	} else {
		b.data[arrOffset] &= byte(^(1 << (position % BYTE_SIZE)))
	}
	return nil
}

func (b *Bitmap) Get(position uint64) (bool, error) {
	if position >= b.size {
		return false, errors.New("invalid index, out of range")
	}
	arrOffset := position / BYTE_SIZE
	return b.data[arrOffset]&(1<<(position%BYTE_SIZE)) > 0, nil
}

func (b *Bitmap) ClearBit(position uint64) error {
	return b.Set(position, false)
}

func (b *Bitmap) Clear() {
	for i := range b.data {
		b.data[i] = 0
	}
}
