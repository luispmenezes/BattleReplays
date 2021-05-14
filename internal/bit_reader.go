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
	position    int
}

func NewBitReader(reader *bufio.Reader) *BitReader {
	return &BitReader{
		reader:      reader,
		bitioReader: bitio.NewReader(reader),
		position:    0,
	}
}

func (b *BitReader) ReadUInt8() uint8 {
	b.position += 8
	return b.bitioReader.TryReadByte()
}

func (b *BitReader) ReadInt8() int8 {
	return int8(b.ReadUInt8())
}

func (b *BitReader) ReadInt16() int16 {
	return int16(b.ReadUInt16())
}

func (b *BitReader) ReadUInt16() uint16 {
	b.position += 16
	r := make([]byte, 2)
	b.bitioReader.TryRead(r)
	return binary.LittleEndian.Uint16(r)
}

func (b *BitReader) ReadInt32() int32 {
	return int32(b.ReadUInt32())
}

func (b *BitReader) ReadUInt32() uint32 {
	b.position += 32
	r := make([]byte, 4)
	b.bitioReader.TryRead(r)
	return binary.LittleEndian.Uint32(r)
}

func (b *BitReader) ReadBytes(n int) []byte {
	b.position += n * 8
	r := make([]byte, n)
	b.bitioReader.TryRead(r)
	return r
}

func (b *BitReader) ReadFloat() float32 {
	/*if b.position & 7 != 0 {
		log.Println(b.position)
		log.Println(b.position >> 3)
		log.Println(b.position / 8)
		log.Println(hex.EncodeToString(b.Peek(4)))
		skipped := b.bitioReader.Align()
		if skipped >  0{
			log.Println(skipped)
			b.position += int(skipped)
		}
	}*/
	b.position += 32
	rBytes := make([]byte, 4)
	b.bitioReader.TryRead(rBytes)
	return math.Float32frombits(binary.LittleEndian.Uint32(rBytes))
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

func (b *BitReader) Align() {
	b.bitioReader.Align()
}
