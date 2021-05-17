package header

import bitreader "BattleReplays/internal"

type UserId struct {
	Index      uint
	Generation byte
}

func deserializeUserId(reader *bitreader.NetBuffer) UserId {
	return UserId{
		Index:      uint(reader.ReadRangedInteger(0, 42)),
		Generation: reader.ReadByte(),
	}
}
