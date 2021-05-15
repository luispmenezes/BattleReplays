package header

import (
	"BattleReplays/internal"
)

type Snapshot struct {
	Position         int32
	NumOfEvents      int32
	NumOfInputEvents int32
	StartBaseTime    float32
	StartScaledTime  float32
}

func deserializeSnapshot(reader *bitreader.NetBuffer) Snapshot {
	return Snapshot{
		Position:         reader.ReadInt32(),
		NumOfEvents:      reader.ReadInt32(),
		NumOfInputEvents: reader.ReadInt32(),
		StartBaseTime:    reader.ReadFloat(),
		StartScaledTime:  reader.ReadFloat(),
	}
}
