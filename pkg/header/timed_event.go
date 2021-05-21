package header

import (
	bitreader "github.com/luispmenezes/BattleReplays/internal"
)

type TimedEventData struct {
	EventId   int32
	EventTime float32
	Round     int32
}

func deserializeTimedEventData(reader *bitreader.NetBuffer) TimedEventData {
	return TimedEventData{
		EventId:   reader.ReadInt32(),
		EventTime: reader.ReadFloat(),
		Round:     reader.ReadInt32(),
	}
}
