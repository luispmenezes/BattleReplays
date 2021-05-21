package bitreader

import "github.com/luispmenezes/battle-replays/internal/bin"

type NetBuffer struct {
	buffer     []byte
	pos        int
	bitsToRead int
}

func NewNetBuffer(buffer []byte, pos, bitsToRead int) *NetBuffer {
	return &NetBuffer{
		buffer:     buffer,
		pos:        pos,
		bitsToRead: bitsToRead,
	}
}

func (n *NetBuffer) ReadUInt32() uint32 {
	num := int(bin.ReadUInt32(n.buffer, 32, n.pos))
	n.pos += 32
	return uint32(num)
}

func (n *NetBuffer) ReadInt32() int32 {
	num := int32(bin.ReadUInt32(n.buffer, 32, n.pos))
	n.pos += 32
	return num
}

func (n *NetBuffer) ReadByte() byte {
	num := bin.ReadByte(n.buffer, 8, n.pos)
	n.pos += 8
	return num
}

func (n *NetBuffer) ReadBytes(numberOfBytes int) []byte {
	dest := make([]byte, numberOfBytes)
	bin.ReadBytes(n.buffer, numberOfBytes, n.pos, &dest, 0)
	n.pos += 8 * numberOfBytes
	return dest
}

func (n *NetBuffer) ReadBytesWithLenOffset(into *[]byte, offset, numberOfBytes int) {
	bin.ReadBytes(n.buffer, numberOfBytes, n.pos, into, offset)
	n.pos += 8 * numberOfBytes
}

func (n *NetBuffer) ReadFloat() float32 {
	if (n.pos & 7) == 0 {
		res := bin.ToSingle(n.buffer, n.pos>>3)
		n.pos += 32
		return res
	}
	dest := make([]byte, 1024)
	n.ReadBytesWithLenOffset(&dest, 0, 4)
	return bin.ToSingle(dest, 0)
}

func (n *NetBuffer) ReadBoolean() bool {
	num := bin.ReadByte(n.buffer, 1, n.pos)
	n.pos++
	return num > 0
}

func (n *NetBuffer) ReadUInt16() uint16 {
	num := bin.ReadUInt16(n.buffer, 16, n.pos)
	n.pos += 16
	return num
}

func (n *NetBuffer) ReadInt16() int16 {
	num := int16(bin.ReadUInt16(n.buffer, 16, n.pos))
	n.pos += 16
	return num
}

func (n *NetBuffer) ReadInt32VariableBits(numberOfBits int) int32 {
	num1 := bin.ReadUInt32(n.buffer, numberOfBits, n.pos)
	n.pos += numberOfBits
	if numberOfBits == 32 {
		return int32(num1)
	}
	num2 := 1<<numberOfBits - 1
	if int64(num1)&int64(num2) == 0 {
		return int32(num1)
	}
	num3 := 4294967295>>33 - numberOfBits
	return -(int32(num1)&int32(num3) + 1)
}

func (n *NetBuffer) ReadRangedInteger(min, max int) int {
	num := n.ReadInt32VariableBits(bin.Pot(uint(max - min)))
	return int(int64(min) + int64(num))
}

func (n *NetBuffer) ReadVariableUInt32() uint {
	num1 := 0
	num2 := 0

	for n.bitsToRead-n.pos >= 8 {
		num3 := n.ReadByte()
		num1 |= (int(num3) & 127) << num2
		num2 += 7
		if (int(num3) & 128) == 0 {
			return uint(num1)
		}
	}
	return uint(num1)
}

func (n *NetBuffer) ReadString() string {
	num := n.ReadVariableUInt32()
	if num <= 0 {
		return ""
	}
	if uint64(n.bitsToRead-n.pos) < uint64(num*8) {
		n.pos = n.bitsToRead
		return ""
	}
	if (n.pos & 7) == 0 {
		startIdx := n.pos >> 3
		stopIdx := startIdx + int(num)
		str := string(n.buffer[startIdx:stopIdx])
		n.pos += 8 * int(num)
		return str
	}
	bytes := n.ReadBytes(int(num))
	return string(bytes)
}

func (n *NetBuffer) ReadUInt64() uint64 {
	num1 := int64(bin.ReadUInt32(n.buffer, 32, n.pos))
	n.pos += 32
	num2 := int64(bin.ReadUInt32(n.buffer, 32, n.pos)) << 32
	num3 := num1 + num2
	n.pos += 32
	return uint64(num3)
}
