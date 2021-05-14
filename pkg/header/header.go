package header

import (
	"BattleReplays/internal"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"log"
)

type Header struct {
	headerSize         int32
	headerUUID         string
	GameplayVersion    int32
	AssetsRevision     int32
	Length             float32
	StartBaseTime      float32
	StartScaledTime    float32
	NumOfEvents        int32
	NumOfInputEvents   int32
	CompletelyRecorded bool
	Thumbnail          []byte
	Snapshots          []Snapshot
	TimedEventsData    []TimedEventData
	MatchIdAsArray     []byte
	MatchType          MatchType
	BaseTypesToLoad    []GameObjectTypeId
}

const expectedHeaderVersionUUID = "d7c1be6b-8c69-5446-9d30-de9c7853975d"

func DeserializeHeader(reader *bitreader.BitReader) (Header, error) {
	header := Header{}

	header.headerSize = (reader.ReadInt32() + 7) >> 3
	log.Printf("HeaderSize: %d", header.headerSize)

	headerUUIDSize := reader.ReadInt8()
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
		header.GameplayVersion = reader.ReadInt32()
		header.AssetsRevision = reader.ReadInt32()
		header.Length = reader.ReadFloat()
		header.StartBaseTime = reader.ReadFloat()
		header.StartScaledTime = reader.ReadFloat()
		header.NumOfEvents = reader.ReadInt32()
		header.NumOfInputEvents = reader.ReadInt32()
		header.CompletelyRecorded = reader.ReadBoolean()
		thumbnailSize := reader.ReadUInt16()
		log.Printf("Gameplay Version: %d AssetsRevisions: %d", header.GameplayVersion, header.AssetsRevision)
		log.Printf("Length: %f StartBaseTime: %f StartScaledTime: %f", header.Length, header.StartBaseTime, header.StartScaledTime)
		log.Printf("NumOfEvents: %d NumOfInputEvents: %d CompletelyRecorded: %v", header.NumOfEvents, header.NumOfInputEvents, header.CompletelyRecorded)
		log.Printf("ThumbnailSize: %d", thumbnailSize)

		if thumbnailSize > 0 {
			header.Thumbnail = reader.ReadBytes(int(thumbnailSize))
		}

		numSnapshots := reader.ReadInt16()
		log.Printf("NumSnapshots: %d", numSnapshots)

		for i := 0; i < int(numSnapshots); i++ {
			header.Snapshots = append(header.Snapshots, deserializeSnapshot(reader))
		}

		numTimedEvents := reader.ReadInt16()
		log.Printf("numTimedEvents: %d", numTimedEvents)

		for i := 0; i < int(numTimedEvents); i++ {
			header.TimedEventsData = append(header.TimedEventsData, deserializeTimedEventData(reader))
		}

		matchIdSize := reader.ReadUInt8()
		log.Printf("matchIdSize: %d", matchIdSize)
		if matchIdSize > 0 {
			header.MatchIdAsArray = reader.ReadBytes(int(matchIdSize))
		}
		header.MatchType = MatchType{typeId: reader.ReadUInt8()}
		log.Printf("MatchType is %d: %s ", header.MatchType.typeId, header.MatchType.AsString())
	}
	return header, nil
}
