/*
 * RingCentral Engage Voice API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package engagevoice

type UploadLeadsResponse struct {
	DeletedCount     int32  `json:"deletedCount"`
	DncReturnedCount int32  `json:"dncReturnedCount"`
	DncUploadCount   int32  `json:"dncUploadCount"`
	HasDeletedLeads  bool   `json:"hasDeletedLeads"`
	InternalDncCount int32  `json:"internalDncCount"`
	LeadsAccepted    int32  `json:"leadsAccepted"`
	LeadsConverted   int32  `json:"leadsConverted"`
	LeadsInserted    int32  `json:"leadsInserted"`
	LeadsSupplied    int32  `json:"leadsSupplied"`
	ListState        string `json:"listState"`
	// Values can be `Your uploaded lead list has successfully completed processing` or `Your uploaded lead list file processing has failed`
	Message string `json:"message"`
	// Values can be `OK` or `Failed`
	ProcessingResult string `json:"processingResult"`
	// Values can be `DEFAULT_NOT_A_FAILURE` or `GENERAL_FAILURE`
	ProcessingStatus string `json:"processingStatus"`
	QuotaCount       int32  `json:"quotaCount"`
	TimeZoneOption   string `json:"timeZoneOption"`
	UploadFileName   string `json:"uploadFileName"`
	WhitelistCount   int32  `json:"whitelistCount"`
}
