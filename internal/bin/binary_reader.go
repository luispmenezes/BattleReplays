package bin

import (
	"encoding/binary"
	"math"
)

func ReadByte(srcBuff []byte, len, start int) byte {
	index := start >> 3
	num1 := start - index*8
	if num1 == 0 && len == 8 {
		return srcBuff[index]
	}
	var num2 byte = srcBuff[index] >> num1
	num3 := len - (8 - num1)
	if num3 < 1 {
		return num2 & (uint8(255)>>8 - uint8(len))
	}
	num4 := srcBuff[index+1] & (uint8(255)>>8 - uint8(num3))
	return num2 | (num4 << uint8(len-num3))
}

func ReadBytes(srcBuf []byte, len int, post int, destBuf *[]byte, destOffset int) {
	srcOffset := post >> 3
	num1 := post - srcOffset*8
	if num1 == 0 {
		blockCopy(srcBuf, srcOffset, destBuf, destOffset, len)
	} else {
		num2 := 8 - num1
		num3 := 255 >> num2
		for index := 0; index < len; index++ {
			num4 := srcBuf[srcOffset] >> num1
			srcOffset++
			num5 := srcBuf[srcOffset] & byte(num3)
			(*destBuf)[destOffset] = byte(num4 | num5<<num2)
			destOffset++
		}
	}
}

func ReadUInt32(buffer []byte, len, start int) uint32 {
	if len <= 8 {
		return uint32(ReadByte(buffer, len, start))
	}
	var num1 = uint32(ReadByte(buffer, 8, start))
	len -= 8
	start += 8
	if len <= 8 {
		return num1 | uint32(ReadByte(buffer, len, start))<<8
	}
	var num2 = num1 | uint32(ReadByte(buffer, 8, start))<<8
	len -= 8
	start += 8
	if len <= 8 {
		var num3 = uint32(ReadByte(buffer, len, start)) << 16
		return num2 | num3
	}
	var num4 = num2 | uint32(ReadByte(buffer, 8, start))<<16
	len -= 8
	start += 8
	return num4 | uint32(ReadByte(buffer, len, start))<<24
}

func ReadUInt16(buffer []byte, len, start int) uint16 {
	if len <= 8 {
		return uint16(ReadByte(buffer, len, start))
	}
	num := uint16(ReadByte(buffer, 8, start))
	len -= 8
	start += 8
	if len <= 8 {
		num |= uint16(uint(ReadByte(buffer, len, start)) << 8)
	}
	return num
}

func ToSingle(source []byte, startIdx int) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(source[startIdx : startIdx+4]))
}

func blockCopy(src []byte, srcOffset int, dst *[]byte, dstOffset int, count int) {
	for i := 0; i < count; i++ {
		(*dst)[dstOffset+i] = src[srcOffset+i]
	}
}

func Pot(in uint) int {
	num := 1
	for {
		in >>= 1
		if in == uint(0) {
			break
		}
		num++
	}
	return num
}
