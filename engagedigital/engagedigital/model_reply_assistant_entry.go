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

type ReplyAssistantEntry struct {
	CategoryIds  []string  `json:"category_ids,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	EntryGroupId string    `json:"entry_group_id,omitempty"`
	ForeignId    string    `json:"foreign_id,omitempty"`
	Id           string    `json:"id"`
	Label        string    `json:"label,omitempty"`
	Shortcuts    string    `json:"shortcuts,omitempty"`
	SourceIds    []string  `json:"source_ids,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	VersionIds   []string  `json:"version_ids,omitempty"`
}