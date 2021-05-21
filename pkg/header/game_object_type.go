package header

import bitreader "github.com/luispmenezes/battle-replays/internal"

type GameObjectType struct {
	Id int32
}

func deserializeGameObjectType(reader *bitreader.NetBuffer) GameObjectType {
	return GameObjectType{
		Id: reader.ReadInt32(),
	}
}
