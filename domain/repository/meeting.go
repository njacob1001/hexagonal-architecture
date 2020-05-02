package repository

import (
	"time"

	"github.com/njacob1001/sinapsis-back/domain/entity"
)

// MeetingRepository interface
type MeetingRepository interface {
	CreateMeeting(*entity.Meeting) (*entity.Meeting, error)
	StartMeeting(meetingUUID string) error
	EndMeeting(meetingUUID string) error
	CancelMeeting(meetingUUID string) error
	ChangeTime(meetingUUID string, start time.Time, end time.Time) (*entity.Meeting, error)
}
