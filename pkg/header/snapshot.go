package header

import (
	bitreader "BattleReplays/internal"
)

type Snapshot struct {
	a int32
	b int32
	c int32
	d float32
	e float32
}

func deserializeSnapshot(reader *bitreader.BitReader) Snapshot {
	return Snapshot{
		a: reader.ReadInt32(),
		b: reader.ReadInt32(),
		c: reader.ReadInt32(),
		d: reader.ReadFloat(),
		e: reader.ReadFloat(),
	}
}
