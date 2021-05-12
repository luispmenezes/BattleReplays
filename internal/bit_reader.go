package bitreader

import (
	"bufio"
	"encoding/binary"
	"github.com/icza/bitio"
	"math"
)

type BitReader struct {
	reader      *bufio.Reader
	bitioReader *bitio.Reader
	order       binary.ByteOrder
	position    int
}

func NewBitReader(reader *bufio.Reader, bigEndian bool) *BitReader {
	var order binary.ByteOrder = binary.LittleEndian
	if bigEndian {
		order = binary.BigEndian
	}

	return &BitReader{
		reader:      reader,
		bitioReader: bitio.NewReader(reader),
		order:       order,
		position:    0,
	}
}

func (b *BitReader) ReadUInt8() uint8 {
	b.position += 8
	return uint8(b.bitioReader.TryReadBits(8))
}

func (b *BitReader) ReadInt8() int8 {
	b.position += 8
	return int8(b.bitioReader.TryReadBits(8))
}

func (b *BitReader) ReadInt16() int16 {
	b.position += 16
	r := make([]byte, 2)
	b.bitioReader.TryRead(r)
	return int16(b.order.Uint16(r))
}

func (b *BitReader) ReadUInt16() uint16 {
	b.position += 16
	r := make([]byte, 2)
	b.bitioReader.TryRead(r)
	return b.order.Uint16(r)
}

func (b *BitReader) ReadInt32() int32 {
	b.position += 32
	r := make([]byte, 4)
	b.bitioReader.TryRead(r)
	return int32(b.order.Uint32(r))
}

func (b *BitReader) ReadUInt32() uint32 {
	b.position += 32
	r := make([]byte, 4)
	b.bitioReader.TryRead(r)
	return b.order.Uint32(r)
}

func (b *BitReader) ReadBytes(n int) []byte {
	b.position += n * 8
	r := make([]byte, n)
	b.bitioReader.TryRead(r)
	return r
}

func (b *BitReader) ReadFloat() float32 {
	b.position += 32
	rBytes := make([]byte, 4)
	b.bitioReader.TryRead(rBytes)
	rInt := b.order.Uint32(rBytes)
	return math.Float32frombits(rInt)
}

func (b *BitReader) ReadBoolean() bool {
	b.position += 1
	return b.bitioReader.TryReadBool()
}

func (b *BitReader) GetPosition() int {
	return b.position
}

func (b *BitReader) Peek(n int) []byte {
	res, _ := b.reader.Peek(n)
	return res
}
