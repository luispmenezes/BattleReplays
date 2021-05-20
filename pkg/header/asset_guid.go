package header

import (
	bitreader "BattleReplays/internal"
	"encoding/binary"
	"encoding/hex"
)

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

func (ag AssetGUID) String() string {
	a := make([]byte, 4)
	binary.LittleEndian.PutUint32(a, uint32(ag.a))
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(ag.b))
	c := make([]byte, 4)
	binary.LittleEndian.PutUint32(c, uint32(ag.c))
	d := make([]byte, 4)
	binary.LittleEndian.PutUint32(d, uint32(ag.d))
	var bytes []byte
	bytes = append(bytes, a...)
	bytes = append(bytes, b...)
	bytes = append(bytes, c...)
	bytes = append(bytes, d...)
	return hex.EncodeToString(bytes)
}
