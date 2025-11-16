package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// LiveSessionStatus represents lifecycle states for live transcription sessions
type LiveSessionStatus string

const (
	LiveStatusActive     LiveSessionStatus = "active"
	LiveStatusFinalizing LiveSessionStatus = "finalizing"
	LiveStatusCompleted  LiveSessionStatus = "completed"
	LiveStatusCancelled  LiveSessionStatus = "cancelled"
)

// LiveTranscriptionSession persists metadata for progressive transcription jobs
type LiveTranscriptionSession struct {
	ID                    string            `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Title                 *string           `json:"title,omitempty" gorm:"type:text"`
	Status                LiveSessionStatus `json:"status" gorm:"type:varchar(20);not null;default:'active'"`
	Parameters            WhisperXParams    `json:"parameters" gorm:"embedded"`
	ChunkCount            int               `json:"chunk_count" gorm:"not null;default:0"`
	LastSequence          int               `json:"last_sequence" gorm:"not null;default:0"`
	AccumulatedTranscript *string           `json:"accumulated_transcript,omitempty" gorm:"type:text"`
	OutputAudioPath       *string           `json:"output_audio_path,omitempty" gorm:"type:text"`
	FinalJobID            *string           `json:"final_job_id,omitempty" gorm:"type:varchar(36)"`
	CreatedAt             time.Time         `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt             time.Time         `json:"updated_at" gorm:"autoUpdateTime"`
	CompletedAt           *time.Time        `json:"completed_at,omitempty"`

	Chunks []LiveTranscriptionChunk `json:"chunks,omitempty" gorm:"foreignKey:SessionID"`
}

// BeforeCreate assigns an ID if missing
func (s *LiveTranscriptionSession) BeforeCreate(tx *gorm.DB) error {
	if s.ID == "" {
		s.ID = uuid.New().String()
	}
	return nil
}

// LiveTranscriptionChunk stores received audio chunk metadata
type LiveTranscriptionChunk struct {
	ID             uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	SessionID      string    `json:"session_id" gorm:"type:varchar(36);index;not null"`
	Sequence       int       `json:"sequence" gorm:"not null"`
	StartOffset    float64   `json:"start_offset" gorm:"type:real"`
	EndOffset      float64   `json:"end_offset" gorm:"type:real"`
	AudioPath      string    `json:"audio_path" gorm:"type:text;not null"`
	TranscriptJSON *string   `json:"transcript_json,omitempty" gorm:"type:text"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func (LiveTranscriptionChunk) TableName() string {
	return "live_transcription_chunks"
}
