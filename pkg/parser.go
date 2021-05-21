package battlereplays

import (
	"bufio"
	"encoding/binary"
	bitreader "github.com/luispmenezes/battle-replays/internal"
	"github.com/luispmenezes/battle-replays/pkg/header"
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

	buf := make([]byte, count)
	_, _ = replayStream.Read(buf)

	p := parser{
		netBuffer: bitreader.NewNetBuffer(buf, 0, int(bitsToRead)),
	}

	h, err := header.DeserializeHeader(p.netBuffer, int(count))

	if err != nil {
		return nil, err
	}

	p.Header = h

	return &p, nil
}
