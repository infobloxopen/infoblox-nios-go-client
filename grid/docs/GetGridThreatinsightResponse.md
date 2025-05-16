# GetGridThreatinsightResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AllowlistUpdatePolicy** | Pointer to **string** | allowlist update policy (manual or automatic) | [optional] 
**ConfigureDomainCollapsing** | Pointer to **bool** | Disable domain collapsing at grid level | [optional] 
**CurrentAllowlist** | Pointer to **string** | The Grid allowlist. | [optional] [readonly] 
**CurrentModuleset** | Pointer to **string** | The current threat insight module set. | [optional] [readonly] 
**DnsTunnelBlockListRpzZones** | Pointer to **[]string** | The list of response policy zones for DNS tunnelling requests. | [optional] 
**DomainCollapsingLevel** | Pointer to **int64** | Level of domain collapsing | [optional] 
**DownloadThreatInsightAllowlistUpdate** | Pointer to **map[string]interface{}** |  | [optional] 
**DownloadThreatInsightModulesetUpdate** | Pointer to **map[string]interface{}** |  | [optional] 
**EnableAllowlistAutoDownload** | Pointer to **bool** | Indicates whether auto download service is enabled | [optional] 
**EnableAllowlistScheduledDownload** | Pointer to **bool** | Indicates whether the custom scheduled settings for auto download is enabled. If false then default frequency is once per 24 hours | [optional] 
**EnableAutoDownload** | Pointer to **bool** | Determines whether the automatic threat insight module set download is enabled. | [optional] 
**EnableScheduledDownload** | Pointer to **bool** | Determines whether the scheduled download of the threat insight module set is enabled. | [optional] 
**LastAllowlistUpdateTime** | Pointer to **int64** | The last update time for the threat insight allowlist. | [optional] [readonly] 
**LastAllowlistUpdateVersion** | Pointer to **string** | The version number of the last updated threat insight allowlist. | [optional] [readonly] 
**LastCheckedForAllowlistUpdate** | Pointer to **int64** | Timestamp of last checked allowlist | [optional] [readonly] 
**LastCheckedForPackageUpdate** | Pointer to **int64** | The last update time for the threat analytics moduleset package. | [optional] [readonly] 
**LastCheckedForUpdate** | Pointer to **int64** | The last time when the threat insight module set was checked for the update. | [optional] [readonly] 
**LastModuleUpdateTime** | Pointer to **int64** | The last update time for the threat insight module set. | [optional] [readonly] 
**LastModuleUpdateVersion** | Pointer to **string** | The version number of the last updated threat insight module set. | [optional] [readonly] 
**LastUpdatedPackageVersion** | Pointer to **string** | The version number of the last updated Moduleset package. | [optional] [readonly] 
**ModuleUpdatePolicy** | Pointer to **string** | The update policy for the threat insight module set. | [optional] 
**MoveBlocklistRpzToAllowList** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** | The Grid name. | [optional] [readonly] 
**ScheduledAllowlistDownload** | Pointer to [**GridThreatinsightScheduledAllowlistDownload**](GridThreatinsightScheduledAllowlistDownload.md) |  | [optional] 
**ScheduledDownload** | Pointer to [**GridThreatinsightScheduledDownload**](GridThreatinsightScheduledDownload.md) |  | [optional] 
**SetLastUploadedThreatInsightModuleset** | Pointer to **map[string]interface{}** |  | [optional] 
**TestThreatInsightServerConnectivity** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdateThreatInsightModuleset** | Pointer to **map[string]interface{}** |  | [optional] 
**Result** | Pointer to [**GridThreatinsight**](GridThreatinsight.md) |  | [optional] 

## Methods

### NewGetGridThreatinsightResponse

`func NewGetGridThreatinsightResponse() *GetGridThreatinsightResponse`

NewGetGridThreatinsightResponse instantiates a new GetGridThreatinsightResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGridThreatinsightResponseWithDefaults

`func NewGetGridThreatinsightResponseWithDefaults() *GetGridThreatinsightResponse`

NewGetGridThreatinsightResponseWithDefaults instantiates a new GetGridThreatinsightResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGridThreatinsightResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGridThreatinsightResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGridThreatinsightResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGridThreatinsightResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllowlistUpdatePolicy

`func (o *GetGridThreatinsightResponse) GetAllowlistUpdatePolicy() string`

GetAllowlistUpdatePolicy returns the AllowlistUpdatePolicy field if non-nil, zero value otherwise.

### GetAllowlistUpdatePolicyOk

`func (o *GetGridThreatinsightResponse) GetAllowlistUpdatePolicyOk() (*string, bool)`

GetAllowlistUpdatePolicyOk returns a tuple with the AllowlistUpdatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlistUpdatePolicy

`func (o *GetGridThreatinsightResponse) SetAllowlistUpdatePolicy(v string)`

SetAllowlistUpdatePolicy sets AllowlistUpdatePolicy field to given value.

### HasAllowlistUpdatePolicy

`func (o *GetGridThreatinsightResponse) HasAllowlistUpdatePolicy() bool`

HasAllowlistUpdatePolicy returns a boolean if a field has been set.

### GetConfigureDomainCollapsing

`func (o *GetGridThreatinsightResponse) GetConfigureDomainCollapsing() bool`

GetConfigureDomainCollapsing returns the ConfigureDomainCollapsing field if non-nil, zero value otherwise.

### GetConfigureDomainCollapsingOk

`func (o *GetGridThreatinsightResponse) GetConfigureDomainCollapsingOk() (*bool, bool)`

GetConfigureDomainCollapsingOk returns a tuple with the ConfigureDomainCollapsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureDomainCollapsing

`func (o *GetGridThreatinsightResponse) SetConfigureDomainCollapsing(v bool)`

SetConfigureDomainCollapsing sets ConfigureDomainCollapsing field to given value.

### HasConfigureDomainCollapsing

`func (o *GetGridThreatinsightResponse) HasConfigureDomainCollapsing() bool`

HasConfigureDomainCollapsing returns a boolean if a field has been set.

### GetCurrentAllowlist

`func (o *GetGridThreatinsightResponse) GetCurrentAllowlist() string`

GetCurrentAllowlist returns the CurrentAllowlist field if non-nil, zero value otherwise.

### GetCurrentAllowlistOk

`func (o *GetGridThreatinsightResponse) GetCurrentAllowlistOk() (*string, bool)`

GetCurrentAllowlistOk returns a tuple with the CurrentAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAllowlist

`func (o *GetGridThreatinsightResponse) SetCurrentAllowlist(v string)`

SetCurrentAllowlist sets CurrentAllowlist field to given value.

### HasCurrentAllowlist

`func (o *GetGridThreatinsightResponse) HasCurrentAllowlist() bool`

HasCurrentAllowlist returns a boolean if a field has been set.

### GetCurrentModuleset

`func (o *GetGridThreatinsightResponse) GetCurrentModuleset() string`

GetCurrentModuleset returns the CurrentModuleset field if non-nil, zero value otherwise.

### GetCurrentModulesetOk

`func (o *GetGridThreatinsightResponse) GetCurrentModulesetOk() (*string, bool)`

GetCurrentModulesetOk returns a tuple with the CurrentModuleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentModuleset

`func (o *GetGridThreatinsightResponse) SetCurrentModuleset(v string)`

SetCurrentModuleset sets CurrentModuleset field to given value.

### HasCurrentModuleset

`func (o *GetGridThreatinsightResponse) HasCurrentModuleset() bool`

HasCurrentModuleset returns a boolean if a field has been set.

### GetDnsTunnelBlockListRpzZones

`func (o *GetGridThreatinsightResponse) GetDnsTunnelBlockListRpzZones() []string`

GetDnsTunnelBlockListRpzZones returns the DnsTunnelBlockListRpzZones field if non-nil, zero value otherwise.

### GetDnsTunnelBlockListRpzZonesOk

`func (o *GetGridThreatinsightResponse) GetDnsTunnelBlockListRpzZonesOk() (*[]string, bool)`

GetDnsTunnelBlockListRpzZonesOk returns a tuple with the DnsTunnelBlockListRpzZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTunnelBlockListRpzZones

`func (o *GetGridThreatinsightResponse) SetDnsTunnelBlockListRpzZones(v []string)`

SetDnsTunnelBlockListRpzZones sets DnsTunnelBlockListRpzZones field to given value.

### HasDnsTunnelBlockListRpzZones

`func (o *GetGridThreatinsightResponse) HasDnsTunnelBlockListRpzZones() bool`

HasDnsTunnelBlockListRpzZones returns a boolean if a field has been set.

### GetDomainCollapsingLevel

`func (o *GetGridThreatinsightResponse) GetDomainCollapsingLevel() int64`

GetDomainCollapsingLevel returns the DomainCollapsingLevel field if non-nil, zero value otherwise.

### GetDomainCollapsingLevelOk

`func (o *GetGridThreatinsightResponse) GetDomainCollapsingLevelOk() (*int64, bool)`

GetDomainCollapsingLevelOk returns a tuple with the DomainCollapsingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainCollapsingLevel

`func (o *GetGridThreatinsightResponse) SetDomainCollapsingLevel(v int64)`

SetDomainCollapsingLevel sets DomainCollapsingLevel field to given value.

### HasDomainCollapsingLevel

`func (o *GetGridThreatinsightResponse) HasDomainCollapsingLevel() bool`

HasDomainCollapsingLevel returns a boolean if a field has been set.

### GetDownloadThreatInsightAllowlistUpdate

`func (o *GetGridThreatinsightResponse) GetDownloadThreatInsightAllowlistUpdate() map[string]interface{}`

GetDownloadThreatInsightAllowlistUpdate returns the DownloadThreatInsightAllowlistUpdate field if non-nil, zero value otherwise.

### GetDownloadThreatInsightAllowlistUpdateOk

`func (o *GetGridThreatinsightResponse) GetDownloadThreatInsightAllowlistUpdateOk() (*map[string]interface{}, bool)`

GetDownloadThreatInsightAllowlistUpdateOk returns a tuple with the DownloadThreatInsightAllowlistUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadThreatInsightAllowlistUpdate

`func (o *GetGridThreatinsightResponse) SetDownloadThreatInsightAllowlistUpdate(v map[string]interface{})`

SetDownloadThreatInsightAllowlistUpdate sets DownloadThreatInsightAllowlistUpdate field to given value.

### HasDownloadThreatInsightAllowlistUpdate

`func (o *GetGridThreatinsightResponse) HasDownloadThreatInsightAllowlistUpdate() bool`

HasDownloadThreatInsightAllowlistUpdate returns a boolean if a field has been set.

### GetDownloadThreatInsightModulesetUpdate

`func (o *GetGridThreatinsightResponse) GetDownloadThreatInsightModulesetUpdate() map[string]interface{}`

GetDownloadThreatInsightModulesetUpdate returns the DownloadThreatInsightModulesetUpdate field if non-nil, zero value otherwise.

### GetDownloadThreatInsightModulesetUpdateOk

`func (o *GetGridThreatinsightResponse) GetDownloadThreatInsightModulesetUpdateOk() (*map[string]interface{}, bool)`

GetDownloadThreatInsightModulesetUpdateOk returns a tuple with the DownloadThreatInsightModulesetUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadThreatInsightModulesetUpdate

`func (o *GetGridThreatinsightResponse) SetDownloadThreatInsightModulesetUpdate(v map[string]interface{})`

SetDownloadThreatInsightModulesetUpdate sets DownloadThreatInsightModulesetUpdate field to given value.

### HasDownloadThreatInsightModulesetUpdate

`func (o *GetGridThreatinsightResponse) HasDownloadThreatInsightModulesetUpdate() bool`

HasDownloadThreatInsightModulesetUpdate returns a boolean if a field has been set.

### GetEnableAllowlistAutoDownload

`func (o *GetGridThreatinsightResponse) GetEnableAllowlistAutoDownload() bool`

GetEnableAllowlistAutoDownload returns the EnableAllowlistAutoDownload field if non-nil, zero value otherwise.

### GetEnableAllowlistAutoDownloadOk

`func (o *GetGridThreatinsightResponse) GetEnableAllowlistAutoDownloadOk() (*bool, bool)`

GetEnableAllowlistAutoDownloadOk returns a tuple with the EnableAllowlistAutoDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllowlistAutoDownload

`func (o *GetGridThreatinsightResponse) SetEnableAllowlistAutoDownload(v bool)`

SetEnableAllowlistAutoDownload sets EnableAllowlistAutoDownload field to given value.

### HasEnableAllowlistAutoDownload

`func (o *GetGridThreatinsightResponse) HasEnableAllowlistAutoDownload() bool`

HasEnableAllowlistAutoDownload returns a boolean if a field has been set.

### GetEnableAllowlistScheduledDownload

`func (o *GetGridThreatinsightResponse) GetEnableAllowlistScheduledDownload() bool`

GetEnableAllowlistScheduledDownload returns the EnableAllowlistScheduledDownload field if non-nil, zero value otherwise.

### GetEnableAllowlistScheduledDownloadOk

`func (o *GetGridThreatinsightResponse) GetEnableAllowlistScheduledDownloadOk() (*bool, bool)`

GetEnableAllowlistScheduledDownloadOk returns a tuple with the EnableAllowlistScheduledDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllowlistScheduledDownload

`func (o *GetGridThreatinsightResponse) SetEnableAllowlistScheduledDownload(v bool)`

SetEnableAllowlistScheduledDownload sets EnableAllowlistScheduledDownload field to given value.

### HasEnableAllowlistScheduledDownload

`func (o *GetGridThreatinsightResponse) HasEnableAllowlistScheduledDownload() bool`

HasEnableAllowlistScheduledDownload returns a boolean if a field has been set.

### GetEnableAutoDownload

`func (o *GetGridThreatinsightResponse) GetEnableAutoDownload() bool`

GetEnableAutoDownload returns the EnableAutoDownload field if non-nil, zero value otherwise.

### GetEnableAutoDownloadOk

`func (o *GetGridThreatinsightResponse) GetEnableAutoDownloadOk() (*bool, bool)`

GetEnableAutoDownloadOk returns a tuple with the EnableAutoDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoDownload

`func (o *GetGridThreatinsightResponse) SetEnableAutoDownload(v bool)`

SetEnableAutoDownload sets EnableAutoDownload field to given value.

### HasEnableAutoDownload

`func (o *GetGridThreatinsightResponse) HasEnableAutoDownload() bool`

HasEnableAutoDownload returns a boolean if a field has been set.

### GetEnableScheduledDownload

`func (o *GetGridThreatinsightResponse) GetEnableScheduledDownload() bool`

GetEnableScheduledDownload returns the EnableScheduledDownload field if non-nil, zero value otherwise.

### GetEnableScheduledDownloadOk

`func (o *GetGridThreatinsightResponse) GetEnableScheduledDownloadOk() (*bool, bool)`

GetEnableScheduledDownloadOk returns a tuple with the EnableScheduledDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableScheduledDownload

`func (o *GetGridThreatinsightResponse) SetEnableScheduledDownload(v bool)`

SetEnableScheduledDownload sets EnableScheduledDownload field to given value.

### HasEnableScheduledDownload

`func (o *GetGridThreatinsightResponse) HasEnableScheduledDownload() bool`

HasEnableScheduledDownload returns a boolean if a field has been set.

### GetLastAllowlistUpdateTime

`func (o *GetGridThreatinsightResponse) GetLastAllowlistUpdateTime() int64`

GetLastAllowlistUpdateTime returns the LastAllowlistUpdateTime field if non-nil, zero value otherwise.

### GetLastAllowlistUpdateTimeOk

`func (o *GetGridThreatinsightResponse) GetLastAllowlistUpdateTimeOk() (*int64, bool)`

GetLastAllowlistUpdateTimeOk returns a tuple with the LastAllowlistUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAllowlistUpdateTime

`func (o *GetGridThreatinsightResponse) SetLastAllowlistUpdateTime(v int64)`

SetLastAllowlistUpdateTime sets LastAllowlistUpdateTime field to given value.

### HasLastAllowlistUpdateTime

`func (o *GetGridThreatinsightResponse) HasLastAllowlistUpdateTime() bool`

HasLastAllowlistUpdateTime returns a boolean if a field has been set.

### GetLastAllowlistUpdateVersion

`func (o *GetGridThreatinsightResponse) GetLastAllowlistUpdateVersion() string`

GetLastAllowlistUpdateVersion returns the LastAllowlistUpdateVersion field if non-nil, zero value otherwise.

### GetLastAllowlistUpdateVersionOk

`func (o *GetGridThreatinsightResponse) GetLastAllowlistUpdateVersionOk() (*string, bool)`

GetLastAllowlistUpdateVersionOk returns a tuple with the LastAllowlistUpdateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAllowlistUpdateVersion

`func (o *GetGridThreatinsightResponse) SetLastAllowlistUpdateVersion(v string)`

SetLastAllowlistUpdateVersion sets LastAllowlistUpdateVersion field to given value.

### HasLastAllowlistUpdateVersion

`func (o *GetGridThreatinsightResponse) HasLastAllowlistUpdateVersion() bool`

HasLastAllowlistUpdateVersion returns a boolean if a field has been set.

### GetLastCheckedForAllowlistUpdate

`func (o *GetGridThreatinsightResponse) GetLastCheckedForAllowlistUpdate() int64`

GetLastCheckedForAllowlistUpdate returns the LastCheckedForAllowlistUpdate field if non-nil, zero value otherwise.

### GetLastCheckedForAllowlistUpdateOk

`func (o *GetGridThreatinsightResponse) GetLastCheckedForAllowlistUpdateOk() (*int64, bool)`

GetLastCheckedForAllowlistUpdateOk returns a tuple with the LastCheckedForAllowlistUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedForAllowlistUpdate

`func (o *GetGridThreatinsightResponse) SetLastCheckedForAllowlistUpdate(v int64)`

SetLastCheckedForAllowlistUpdate sets LastCheckedForAllowlistUpdate field to given value.

### HasLastCheckedForAllowlistUpdate

`func (o *GetGridThreatinsightResponse) HasLastCheckedForAllowlistUpdate() bool`

HasLastCheckedForAllowlistUpdate returns a boolean if a field has been set.

### GetLastCheckedForPackageUpdate

`func (o *GetGridThreatinsightResponse) GetLastCheckedForPackageUpdate() int64`

GetLastCheckedForPackageUpdate returns the LastCheckedForPackageUpdate field if non-nil, zero value otherwise.

### GetLastCheckedForPackageUpdateOk

`func (o *GetGridThreatinsightResponse) GetLastCheckedForPackageUpdateOk() (*int64, bool)`

GetLastCheckedForPackageUpdateOk returns a tuple with the LastCheckedForPackageUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedForPackageUpdate

`func (o *GetGridThreatinsightResponse) SetLastCheckedForPackageUpdate(v int64)`

SetLastCheckedForPackageUpdate sets LastCheckedForPackageUpdate field to given value.

### HasLastCheckedForPackageUpdate

`func (o *GetGridThreatinsightResponse) HasLastCheckedForPackageUpdate() bool`

HasLastCheckedForPackageUpdate returns a boolean if a field has been set.

### GetLastCheckedForUpdate

`func (o *GetGridThreatinsightResponse) GetLastCheckedForUpdate() int64`

GetLastCheckedForUpdate returns the LastCheckedForUpdate field if non-nil, zero value otherwise.

### GetLastCheckedForUpdateOk

`func (o *GetGridThreatinsightResponse) GetLastCheckedForUpdateOk() (*int64, bool)`

GetLastCheckedForUpdateOk returns a tuple with the LastCheckedForUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedForUpdate

`func (o *GetGridThreatinsightResponse) SetLastCheckedForUpdate(v int64)`

SetLastCheckedForUpdate sets LastCheckedForUpdate field to given value.

### HasLastCheckedForUpdate

`func (o *GetGridThreatinsightResponse) HasLastCheckedForUpdate() bool`

HasLastCheckedForUpdate returns a boolean if a field has been set.

### GetLastModuleUpdateTime

`func (o *GetGridThreatinsightResponse) GetLastModuleUpdateTime() int64`

GetLastModuleUpdateTime returns the LastModuleUpdateTime field if non-nil, zero value otherwise.

### GetLastModuleUpdateTimeOk

`func (o *GetGridThreatinsightResponse) GetLastModuleUpdateTimeOk() (*int64, bool)`

GetLastModuleUpdateTimeOk returns a tuple with the LastModuleUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModuleUpdateTime

`func (o *GetGridThreatinsightResponse) SetLastModuleUpdateTime(v int64)`

SetLastModuleUpdateTime sets LastModuleUpdateTime field to given value.

### HasLastModuleUpdateTime

`func (o *GetGridThreatinsightResponse) HasLastModuleUpdateTime() bool`

HasLastModuleUpdateTime returns a boolean if a field has been set.

### GetLastModuleUpdateVersion

`func (o *GetGridThreatinsightResponse) GetLastModuleUpdateVersion() string`

GetLastModuleUpdateVersion returns the LastModuleUpdateVersion field if non-nil, zero value otherwise.

### GetLastModuleUpdateVersionOk

`func (o *GetGridThreatinsightResponse) GetLastModuleUpdateVersionOk() (*string, bool)`

GetLastModuleUpdateVersionOk returns a tuple with the LastModuleUpdateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModuleUpdateVersion

`func (o *GetGridThreatinsightResponse) SetLastModuleUpdateVersion(v string)`

SetLastModuleUpdateVersion sets LastModuleUpdateVersion field to given value.

### HasLastModuleUpdateVersion

`func (o *GetGridThreatinsightResponse) HasLastModuleUpdateVersion() bool`

HasLastModuleUpdateVersion returns a boolean if a field has been set.

### GetLastUpdatedPackageVersion

`func (o *GetGridThreatinsightResponse) GetLastUpdatedPackageVersion() string`

GetLastUpdatedPackageVersion returns the LastUpdatedPackageVersion field if non-nil, zero value otherwise.

### GetLastUpdatedPackageVersionOk

`func (o *GetGridThreatinsightResponse) GetLastUpdatedPackageVersionOk() (*string, bool)`

GetLastUpdatedPackageVersionOk returns a tuple with the LastUpdatedPackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedPackageVersion

`func (o *GetGridThreatinsightResponse) SetLastUpdatedPackageVersion(v string)`

SetLastUpdatedPackageVersion sets LastUpdatedPackageVersion field to given value.

### HasLastUpdatedPackageVersion

`func (o *GetGridThreatinsightResponse) HasLastUpdatedPackageVersion() bool`

HasLastUpdatedPackageVersion returns a boolean if a field has been set.

### GetModuleUpdatePolicy

`func (o *GetGridThreatinsightResponse) GetModuleUpdatePolicy() string`

GetModuleUpdatePolicy returns the ModuleUpdatePolicy field if non-nil, zero value otherwise.

### GetModuleUpdatePolicyOk

`func (o *GetGridThreatinsightResponse) GetModuleUpdatePolicyOk() (*string, bool)`

GetModuleUpdatePolicyOk returns a tuple with the ModuleUpdatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleUpdatePolicy

`func (o *GetGridThreatinsightResponse) SetModuleUpdatePolicy(v string)`

SetModuleUpdatePolicy sets ModuleUpdatePolicy field to given value.

### HasModuleUpdatePolicy

`func (o *GetGridThreatinsightResponse) HasModuleUpdatePolicy() bool`

HasModuleUpdatePolicy returns a boolean if a field has been set.

### GetMoveBlocklistRpzToAllowList

`func (o *GetGridThreatinsightResponse) GetMoveBlocklistRpzToAllowList() map[string]interface{}`

GetMoveBlocklistRpzToAllowList returns the MoveBlocklistRpzToAllowList field if non-nil, zero value otherwise.

### GetMoveBlocklistRpzToAllowListOk

`func (o *GetGridThreatinsightResponse) GetMoveBlocklistRpzToAllowListOk() (*map[string]interface{}, bool)`

GetMoveBlocklistRpzToAllowListOk returns a tuple with the MoveBlocklistRpzToAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveBlocklistRpzToAllowList

`func (o *GetGridThreatinsightResponse) SetMoveBlocklistRpzToAllowList(v map[string]interface{})`

SetMoveBlocklistRpzToAllowList sets MoveBlocklistRpzToAllowList field to given value.

### HasMoveBlocklistRpzToAllowList

`func (o *GetGridThreatinsightResponse) HasMoveBlocklistRpzToAllowList() bool`

HasMoveBlocklistRpzToAllowList returns a boolean if a field has been set.

### GetName

`func (o *GetGridThreatinsightResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetGridThreatinsightResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetGridThreatinsightResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetGridThreatinsightResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheduledAllowlistDownload

`func (o *GetGridThreatinsightResponse) GetScheduledAllowlistDownload() GridThreatinsightScheduledAllowlistDownload`

GetScheduledAllowlistDownload returns the ScheduledAllowlistDownload field if non-nil, zero value otherwise.

### GetScheduledAllowlistDownloadOk

`func (o *GetGridThreatinsightResponse) GetScheduledAllowlistDownloadOk() (*GridThreatinsightScheduledAllowlistDownload, bool)`

GetScheduledAllowlistDownloadOk returns a tuple with the ScheduledAllowlistDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAllowlistDownload

`func (o *GetGridThreatinsightResponse) SetScheduledAllowlistDownload(v GridThreatinsightScheduledAllowlistDownload)`

SetScheduledAllowlistDownload sets ScheduledAllowlistDownload field to given value.

### HasScheduledAllowlistDownload

`func (o *GetGridThreatinsightResponse) HasScheduledAllowlistDownload() bool`

HasScheduledAllowlistDownload returns a boolean if a field has been set.

### GetScheduledDownload

`func (o *GetGridThreatinsightResponse) GetScheduledDownload() GridThreatinsightScheduledDownload`

GetScheduledDownload returns the ScheduledDownload field if non-nil, zero value otherwise.

### GetScheduledDownloadOk

`func (o *GetGridThreatinsightResponse) GetScheduledDownloadOk() (*GridThreatinsightScheduledDownload, bool)`

GetScheduledDownloadOk returns a tuple with the ScheduledDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDownload

`func (o *GetGridThreatinsightResponse) SetScheduledDownload(v GridThreatinsightScheduledDownload)`

SetScheduledDownload sets ScheduledDownload field to given value.

### HasScheduledDownload

`func (o *GetGridThreatinsightResponse) HasScheduledDownload() bool`

HasScheduledDownload returns a boolean if a field has been set.

### GetSetLastUploadedThreatInsightModuleset

`func (o *GetGridThreatinsightResponse) GetSetLastUploadedThreatInsightModuleset() map[string]interface{}`

GetSetLastUploadedThreatInsightModuleset returns the SetLastUploadedThreatInsightModuleset field if non-nil, zero value otherwise.

### GetSetLastUploadedThreatInsightModulesetOk

`func (o *GetGridThreatinsightResponse) GetSetLastUploadedThreatInsightModulesetOk() (*map[string]interface{}, bool)`

GetSetLastUploadedThreatInsightModulesetOk returns a tuple with the SetLastUploadedThreatInsightModuleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetLastUploadedThreatInsightModuleset

`func (o *GetGridThreatinsightResponse) SetSetLastUploadedThreatInsightModuleset(v map[string]interface{})`

SetSetLastUploadedThreatInsightModuleset sets SetLastUploadedThreatInsightModuleset field to given value.

### HasSetLastUploadedThreatInsightModuleset

`func (o *GetGridThreatinsightResponse) HasSetLastUploadedThreatInsightModuleset() bool`

HasSetLastUploadedThreatInsightModuleset returns a boolean if a field has been set.

### GetTestThreatInsightServerConnectivity

`func (o *GetGridThreatinsightResponse) GetTestThreatInsightServerConnectivity() map[string]interface{}`

GetTestThreatInsightServerConnectivity returns the TestThreatInsightServerConnectivity field if non-nil, zero value otherwise.

### GetTestThreatInsightServerConnectivityOk

`func (o *GetGridThreatinsightResponse) GetTestThreatInsightServerConnectivityOk() (*map[string]interface{}, bool)`

GetTestThreatInsightServerConnectivityOk returns a tuple with the TestThreatInsightServerConnectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestThreatInsightServerConnectivity

`func (o *GetGridThreatinsightResponse) SetTestThreatInsightServerConnectivity(v map[string]interface{})`

SetTestThreatInsightServerConnectivity sets TestThreatInsightServerConnectivity field to given value.

### HasTestThreatInsightServerConnectivity

`func (o *GetGridThreatinsightResponse) HasTestThreatInsightServerConnectivity() bool`

HasTestThreatInsightServerConnectivity returns a boolean if a field has been set.

### GetUpdateThreatInsightModuleset

`func (o *GetGridThreatinsightResponse) GetUpdateThreatInsightModuleset() map[string]interface{}`

GetUpdateThreatInsightModuleset returns the UpdateThreatInsightModuleset field if non-nil, zero value otherwise.

### GetUpdateThreatInsightModulesetOk

`func (o *GetGridThreatinsightResponse) GetUpdateThreatInsightModulesetOk() (*map[string]interface{}, bool)`

GetUpdateThreatInsightModulesetOk returns a tuple with the UpdateThreatInsightModuleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateThreatInsightModuleset

`func (o *GetGridThreatinsightResponse) SetUpdateThreatInsightModuleset(v map[string]interface{})`

SetUpdateThreatInsightModuleset sets UpdateThreatInsightModuleset field to given value.

### HasUpdateThreatInsightModuleset

`func (o *GetGridThreatinsightResponse) HasUpdateThreatInsightModuleset() bool`

HasUpdateThreatInsightModuleset returns a boolean if a field has been set.

### GetResult

`func (o *GetGridThreatinsightResponse) GetResult() GridThreatinsight`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGridThreatinsightResponse) GetResultOk() (*GridThreatinsight, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGridThreatinsightResponse) SetResult(v GridThreatinsight)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGridThreatinsightResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


