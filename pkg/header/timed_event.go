package header

import (
	"BattleReplays/internal"
)

type TimedEventData struct {
	a int32
	b float32
	c int32
}

func deserializeTimedEventData(reader *bitreader.NetBuffer) TimedEventData {
	return TimedEventData{
		a: reader.ReadInt32(),
		b: reader.ReadFloat(),
		c: reader.ReadInt32(),
	}
}
