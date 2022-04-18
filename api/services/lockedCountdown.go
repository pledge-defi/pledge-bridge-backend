package services

import (
	"pledge-bridge-backend/api/common"
	"pledge-bridge-backend/api/models/response"
	"pledge-bridge-backend/config"
	"pledge-bridge-backend/utils"
	"time"
)

type LockedCountdown struct{}

func NewLockedCountdown() *LockedCountdown {
	return &LockedCountdown{}
}

func (c *LockedCountdown) NextTimeSeconds(res *response.LockedCountdown) int {

	seconds := c.GetNextReleaseTimeStamp() - time.Now().Unix()

	res.Time = "UTC"
	res.Timestamp = utils.Int64ToString(seconds)

	return 0
}

func (c *LockedCountdown) GetNextReleaseTimeStamp() int64 {

	if config.Config.Env.IsDev {
		if common.ReleaseTime-time.Now().Unix() < 0 {
			sec := config.Config.Env.ReleaseDurDev - ((time.Now().Unix() - common.ReleaseTime) % config.Config.Env.ReleaseDurDev)
			common.SetReleaseTime(sec)
		}
		return common.ReleaseTime
	} else {

		now := time.Now()
		offset := int(time.Sunday - now.Weekday())
		if offset < 0 {
			offset = 7 + offset
		}
		weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)

		return weekStartDate.Unix()
	}
}
