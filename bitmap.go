package bitmap

import (
	"errors"
)

type bitmap struct {
	data []uint64
	size uint64
}

func NewBitmap(size uint64) *bitmap {
	// hardcode 64bit
	sizeArr := size / 64
	arr := make([]uint64, sizeArr)
	if size%64 != 0 {
		arr = append(arr, 0)
	}
	return &bitmap{size: size, data: arr}
}

func (b *bitmap) Set(val bool, position uint8) error {
	if position < 0 || position >= 64 {
		return errors.New("position only available in range [0, 64)")
	}
	if val {
		b.data |= uint64(1 << position)
	} else {
		b.data &= uint64(^(1 << position))
	}
	return nil
}

func (b *bitmap) Get(position uint8) (bool, error) {
	if position < 0 || position >= 64 {
		return false, errors.New("position only available in range [0, 64)")
	}

	return b.data&(1<<position) > 0, nil
}

func (b *bitmap) ClearBit(position uint8) (bool, error) {

}

func (b *bitmap) Clear() {
	b.data = 0
}
