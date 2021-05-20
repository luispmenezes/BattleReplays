package header

import (
	"BattleReplays/internal"
	"BattleReplays/pkg/utils"
	"encoding/hex"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"log"
)

type Header struct {
	headerSize         int32
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

func DeserializeHeader(reader *bitreader.NetBuffer) (Header, error) {
	header := Header{}

	headerUUIDSize := reader.ReadByte()
	log.Printf("HeaderUUIDSize: %d", headerUUIDSize)

	headerUUIDBytes := reader.ReadBytes(int(headerUUIDSize))
	headerUUID, err := uuid.FromBytes(headerUUIDBytes)
	if err != nil {
		return header, errors.Wrap(err, "failed to parse Header UUID from bytes")
	}
	header.headerUUID = headerUUID.String()

	if headerUUID.String() != expectedHeaderVersionUUID {
		log.Printf("header UUID does not match got: %s expected: %s", headerUUID.String(), expectedHeaderVersionUUID)
		//return header, errors.New(fmt.Sprintf("Unexpected Header UUID got %s", headerUUID.String()))
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
		log.Printf("Gameplay Version: %d AssetsRevisions: %d", header.gameplayVersion, header.AssetsRevision)
		log.Printf("Length: %f StartBaseTime: %f StartScaledTime: %f", header.Length, header.startBaseTime, header.startScaledTime)
		log.Printf("NumOfEvents: %d NumOfInputEvents: %d CompletelyRecorded: %v", header.numOfEvents, header.numOfInputEvents, header.CompletelyRecorded)
		log.Printf("ThumbnailSize: %d", thumbnailSize)

		if thumbnailSize > 0 {
			header.thumbnail = reader.ReadBytes(int(thumbnailSize))
		}

		numSnapshots := reader.ReadInt16()
		log.Printf("NumSnapshots: %d", numSnapshots)

		for i := 0; i < int(numSnapshots); i++ {
			header.Snapshots = append(header.Snapshots, deserializeSnapshot(reader))
		}

		numTimedEvents := reader.ReadInt16()
		log.Printf("numTimedEvents: %d", numTimedEvents)

		for i := 0; i < int(numTimedEvents); i++ {
			header.timedEventsData = append(header.timedEventsData, deserializeTimedEventData(reader))
		}

		matchIdSize := reader.ReadByte()
		log.Printf("matchIdSize: %d", matchIdSize)
		if matchIdSize > 0 {
			header.matchIdAsArray = reader.ReadBytes(int(matchIdSize))
		}
		header.MatchId = hex.EncodeToString(header.matchIdAsArray)
		log.Println("MatchId: " + header.MatchId)

		header.matchTypeObj = MatchType{typeId: reader.ReadByte()}
		header.MatchType = header.matchTypeObj.AsString()
		log.Printf("MatchType id: %d ( %s )", header.matchTypeObj.typeId, header.MatchType)

		baseTypesLen := reader.ReadInt16()
		log.Printf("baseTypesLen: %d", baseTypesLen)
		for i := 0; i < int(baseTypesLen); i++ {
			header.BaseTypesToLoad = append(header.BaseTypesToLoad, deserializeGameObjectType(reader))
		}

		header.mapAsset = deserializeAssetGUID(reader)
		header.MapName = utils.GetMapFromId(header.mapAsset.String())
		log.Printf("Map Asset GUID: %v (%s)", header.mapAsset, header.MapName)

		header.LocalUserId = deserializeUserId(reader)
		log.Printf("LocalUserId: %v", header.LocalUserId)

		header.TeamSize = int(reader.ReadByte())
		log.Printf("TeamSize: %d", header.TeamSize)

		header.RoundsToWin = int(reader.ReadByte())
		log.Printf("RoundsToWin: %d", header.RoundsToWin)

		header.LockedRounds = int(reader.ReadByte())
		log.Printf("LockedRounds: %d", header.LockedRounds)
		header.Team1Score = int(reader.ReadByte())
		header.Team2Score = int(reader.ReadByte())
		log.Printf("Score: %d - %d", header.Team1Score, header.Team2Score)

		numUsers := int(reader.ReadByte())
		log.Printf("Num of Users: %d", numUsers)

		for i := 0; i < header.TeamSize*2; i++ {
			header.Users = append(header.Users, deserializeUserData(reader))
		}

		for _, userData := range header.Users {
			log.Println(userData)
		}
	}
	return header, nil
}
