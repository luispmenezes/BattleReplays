package header

import bitreader "BattleReplays/internal"

type AssetGUID struct {
	a int32
	b int32
	c int32
	d int32
}

func deserializeAssetGUID(reader *bitreader.NetBuffer) AssetGUID {
	return AssetGUID{
		a: reader.ReadInt32(),
		b: reader.ReadInt32(),
		c: reader.ReadInt32(),
		d: reader.ReadInt32(),
	}
}
