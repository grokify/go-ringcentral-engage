/*
 * Engage Digital API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package engagedigital

import (
	"time"
)

type Attachment struct {
	ContentType   string    `json:"content_type,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	Embed         bool      `json:"embed,omitempty"`
	Filename      string    `json:"filename,omitempty"`
	ForeignId     string    `json:"foreign_id,omitempty"`
	Id            string    `json:"id,omitempty"`
	Public        bool      `json:"public?,omitempty"`
	Size          int32     `json:"size,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	Url           string    `json:"url,omitempty"`
	VideoMetadata []string  `json:"video_metadata,omitempty"`
}
