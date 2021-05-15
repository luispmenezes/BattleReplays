package BattleReplays

import (
	"BattleReplays/internal"
	"BattleReplays/pkg/header"
	"bufio"
	"encoding/binary"
	"log"
	"os"
)

type parser struct {
	netBuffer *bitreader.NetBuffer
	Header    header.Header
}

func NewParser(f *os.File) (*parser, error) {

	replayStream := bufio.NewReader(f)

	r := make([]byte, 4)
	_, _ = replayStream.Read(r)
	bitsToRead := int32(binary.LittleEndian.Uint32(r))
	count := bitsToRead + 7>>3
	log.Println(count)

	buf := make([]byte, count)
	_, _ = replayStream.Read(buf)

	p := parser{
		netBuffer: bitreader.NewNetBuffer(buf, 0),
	}

	h, err := header.DeserializeHeader(p.netBuffer)

	if err != nil {
		return nil, err
	}

	p.Header = h

	return &p, nil
}
