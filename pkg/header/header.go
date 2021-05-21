package header

import (
	"encoding/hex"
	"github.com/google/uuid"
	bitreader "github.com/luispmenezes/battle-replays/internal"
	"github.com/luispmenezes/battle-replays/pkg/utils"
	"github.com/pkg/errors"
	"log"
)

type Header struct {
	HeaderSize         int
	headerUUID         string
	gameplayVersion    int32
	AssetsRevision     int32
	Length             float32
	startBaseTime      float32
	startScaledTime    float32
	numOfEvents        int32
	numOfInputEvents   int32
	CompletelyRecorded bool
	thumbnail          []byte
	Snapshots          []Snapshot
	timedEventsData    []TimedEventData
	matchIdAsArray     []byte
	MatchId            string
	matchTypeObj       MatchType
	MatchType          string
	BaseTypesToLoad    []GameObjectType
	mapAsset           AssetGUID
	MapName            string
	LocalUserId        UserId
	TeamSize           int
	RoundsToWin        int
	LockedRounds       int
	Team1Score         int
	Team2Score         int
	Users              []UserData
}

const expectedHeaderVersionUUID = "d7c1be6b-8c69-5446-9d30-de9c7853975d"

func DeserializeHeader(reader *bitreader.NetBuffer, byteCount int) (Header, error) {
	header := Header{}

	header.HeaderSize = byteCount
	headerUUIDSize := reader.ReadByte()
	headerUUIDBytes := reader.ReadBytes(int(headerUUIDSize))
	headerUUID, err := uuid.FromBytes(headerUUIDBytes)
	if err != nil {
		return header, errors.Wrap(err, "failed to parse Header UUID from bytes")
	}
	header.headerUUID = headerUUID.String()

	if headerUUID.String() != expectedHeaderVersionUUID {
		log.Printf("header UUID does not match got: %s expected: %s", headerUUID.String(), expectedHeaderVersionUUID)
	}

	checkpoint := reader.ReadInt32()

	if checkpoint == 1 {
		header.gameplayVersion = reader.ReadInt32()
		header.AssetsRevision = reader.ReadInt32()
		header.Length = reader.ReadFloat()
		header.startBaseTime = reader.ReadFloat()
		header.startScaledTime = reader.ReadFloat()
		header.numOfEvents = reader.ReadInt32()
		header.numOfInputEvents = reader.ReadInt32()
		header.CompletelyRecorded = reader.ReadBoolean()
		thumbnailSize := reader.ReadUInt16()

		if thumbnailSize > 0 {
			header.thumbnail = reader.ReadBytes(int(thumbnailSize))
		}

		numSnapshots := reader.ReadInt16()

		for i := 0; i < int(numSnapshots); i++ {
			header.Snapshots = append(header.Snapshots, deserializeSnapshot(reader))
		}

		numTimedEvents := reader.ReadInt16()

		for i := 0; i < int(numTimedEvents); i++ {
			header.timedEventsData = append(header.timedEventsData, deserializeTimedEventData(reader))
		}

		matchIdSize := reader.ReadByte()
		if matchIdSize > 0 {
			header.matchIdAsArray = reader.ReadBytes(int(matchIdSize))
		}
		header.MatchId = hex.EncodeToString(header.matchIdAsArray)
		header.matchTypeObj = MatchType{typeId: reader.ReadByte()}
		header.MatchType = header.matchTypeObj.AsString()

		baseTypesLen := reader.ReadInt16()
		for i := 0; i < int(baseTypesLen); i++ {
			header.BaseTypesToLoad = append(header.BaseTypesToLoad, deserializeGameObjectType(reader))
		}

		header.mapAsset = deserializeAssetGUID(reader)
		header.MapName = utils.GetMapFromId(header.mapAsset.String())
		header.LocalUserId = deserializeUserId(reader)
		header.TeamSize = int(reader.ReadByte())
		header.RoundsToWin = int(reader.ReadByte())
		header.LockedRounds = int(reader.ReadByte())
		header.Team1Score = int(reader.ReadByte())
		header.Team2Score = int(reader.ReadByte())

		reader.ReadByte()

		for i := 0; i < header.TeamSize*2; i++ {
			header.Users = append(header.Users, deserializeUserData(reader))
		}
	}
	return header, nil
}
