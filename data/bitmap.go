package main

import (
	"fmt"
)

// Bitmap 采用的是以空间换时间的思想，数组中最大元素值为7，所以在内存中开辟8位的存储空间，存储空间大小的确定方法是（元素最大值进位到8的倍数/8），之所以除以8，是因为开辟空间的时候以byte为单位，1byte=8bit
type BitMap struct {
	data    []byte
	bitsize uint64
	maxpos  uint64
}

func main() {

}

// SetBit 将 offset 位置的 bit 置为 value (0/1)
func (this *BitMap) SetBit(offset uint64, value uint8) bool {
	index, pos := offset/8, offset%8

	if this.bitsize < offset {
		return false
	}

	if value == 0 {
		// &^ 清位
		this.data[index] &^= 0x01 << pos
	} else {
		this.data[index] |= 0x01 << pos

		// 记录曾经设置为 1 的最大位置
		if this.maxpos < offset {
			this.maxpos = offset
		}
	}

	return true
}

func demo() {
	// 1. bitmap
	var bitmap [8]byte
	fmt.Printf("%T \n", bitmap)
	fmt.Printf("%v \n", bitmap)
	data := []int{4, 6, 3, 1, 7}
	for _, v := range data {
		bitmap[v] = 1
	}
	fmt.Printf("%v \n", bitmap)
}
