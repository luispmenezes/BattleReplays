package header

import bitreader "github.com/luispmenezes/BattleReplays/internal"

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
