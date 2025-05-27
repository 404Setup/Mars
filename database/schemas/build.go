package schemas

import (
	"time"

	"Mars/shared/schemas"
	"Mars/shared/utils/json"
)

type Build struct {
	ID        uint      `gorm:"primary_key;autoIncrement"`
	Project   string    `gorm:"type:varchar(48);size:48;index"`
	Version   string    `gorm:"type:varchar(48);size:48;index"`
	Number    int       `gorm:"index"`
	Time      time.Time `gorm:"index"`
	Changes   []byte    `gorm:"type:blob"`
	Downloads []byte    `gorm:"type:blob"`
	Channel   Channel   `gorm:"type:varchar(14);size:14;default:default"`
	Promoted  bool      `gorm:"default:false"`
}

func (b *Build) MarshalChanges(changes []schemas.Change) error {
	d, err := json.JSON.Marshal(changes)
	if err != nil {
		return err
	}
	b.Changes = d
	return nil
}

func (b *Build) UnmarshalChanges() []schemas.Change {
	var changes []schemas.Change
	_ = json.JSON.Unmarshal(b.Changes, &changes)
	return changes
}

func (b *Build) ChannelOrDefault() Channel {
	if b.Channel == "default" {
		return ChannelDefault
	}
	return ChannelExperimental
}

func (b *Build) PromotedOrDefault() bool {
	return b.Promoted
}

// MarshalDownloads marshals the Downloads map to JSON
func (b *Build) MarshalDownloads(downloads map[string]schemas.ApplicationVersionsSchema) error {
	d, err := json.JSON.Marshal(downloads)
	if err != nil {
		return err
	}
	b.Downloads = d
	return nil
}

// UnmarshalDownloads unmarshals the JSON to a Downloads map
func (b *Build) UnmarshalDownloads() map[string]schemas.ApplicationVersionsSchema {
	var downloads map[string]schemas.ApplicationVersionsSchema
	_ = json.JSON.Unmarshal(b.Downloads, &downloads)
	return downloads
}
