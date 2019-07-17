# Source

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) |  | [optional] 
**Active** | **bool** | Activate/deactivate the source | [optional] 
**CommunityId** | **string** |  | [optional] 
**ChannelId** | **string** |  | [optional] 
**Name** | **string** | Source name | [optional] 
**Status** | **string** |  | [optional] 
**Color** | **int32** | Color of the icon: Default: 0 Blue: 1 Green: 2 Turquoise: 3 Purple: 4 Yellow: 5 Orange: 6 Red: 7 Asphalt: 8 Grey: 9 | [optional] 
**SlaResponse** | **int32** | Response time (seconds) | [optional] 
**SlaExpiredStrategy** | **string** | SLA expired strategy (\&quot;max\&quot;, \&quot;half\&quot; or \&quot;base\&quot;) | [optional] 
**InterventionMessagesBoost** | **int32** | Priority boost of messages with intervention | [optional] 
**TransferredTasksBoost** | **int32** | Priority boost of transferred tasks | [optional] 
**HiddenFromStats** | **bool** | Hide from statistics | [optional] 
**DefaultCategoryIds** | **[]string** | Default categories | [optional] 
**UserThreadDefaultCategoryIds** | **[]string** | Default categories (agent messages) | [optional] 
**DefaultContentLanguage** | **string** | Default content language | [optional] 
**AutoDetectContentLanguage** | **bool** | Auto-detect content language (Boolean) | [optional] 
**TimeSheetIds** | **[]string** |  | [optional] 
**ContentArchiving** | **bool** | Automatic archiving of old contents (Boolean) | [optional] 
**ContentArchivingPeriod** | **int64** | Archive contents older than (seconds) | [optional] 
**ContentLanguages** | **[]string** |  | [optional] 
**Type** | **string** |  | [optional] 
**ErrorMessage** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


