package BattleReplays

import (
	"BattleReplays/internal"
	"BattleReplays/pkg/header"
	"bufio"
	"os"
)

type parser struct {
	bitReader *bitreader.BitReader
	Header    header.Header
}

func NewParser(f *os.File) (*parser, error) {

	replayStream := bufio.NewReader(f)

	p := parser{
		bitReader: bitreader.NewBitReader(replayStream),
	}

	h, err := header.DeserializeHeader(p.bitReader)

	if err != nil {
		return nil, err
	}

	p.Header = h

	return &p, nil
}
