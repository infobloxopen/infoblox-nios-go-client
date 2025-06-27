# GridThreatinsight

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
**Name** | Pointer to **string** | The Grid name. | [optional] [readonly] 
**ScheduledAllowlistDownload** | Pointer to [**GridThreatinsightScheduledAllowlistDownload**](GridThreatinsightScheduledAllowlistDownload.md) |  | [optional] 
**ScheduledDownload** | Pointer to [**GridThreatinsightScheduledDownload**](GridThreatinsightScheduledDownload.md) |  | [optional] 

## Methods

### NewGridThreatinsight

`func NewGridThreatinsight() *GridThreatinsight`

NewGridThreatinsight instantiates a new GridThreatinsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridThreatinsightWithDefaults

`func NewGridThreatinsightWithDefaults() *GridThreatinsight`

NewGridThreatinsightWithDefaults instantiates a new GridThreatinsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GridThreatinsight) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GridThreatinsight) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GridThreatinsight) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GridThreatinsight) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllowlistUpdatePolicy

`func (o *GridThreatinsight) GetAllowlistUpdatePolicy() string`

GetAllowlistUpdatePolicy returns the AllowlistUpdatePolicy field if non-nil, zero value otherwise.

### GetAllowlistUpdatePolicyOk

`func (o *GridThreatinsight) GetAllowlistUpdatePolicyOk() (*string, bool)`

GetAllowlistUpdatePolicyOk returns a tuple with the AllowlistUpdatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlistUpdatePolicy

`func (o *GridThreatinsight) SetAllowlistUpdatePolicy(v string)`

SetAllowlistUpdatePolicy sets AllowlistUpdatePolicy field to given value.

### HasAllowlistUpdatePolicy

`func (o *GridThreatinsight) HasAllowlistUpdatePolicy() bool`

HasAllowlistUpdatePolicy returns a boolean if a field has been set.

### GetConfigureDomainCollapsing

`func (o *GridThreatinsight) GetConfigureDomainCollapsing() bool`

GetConfigureDomainCollapsing returns the ConfigureDomainCollapsing field if non-nil, zero value otherwise.

### GetConfigureDomainCollapsingOk

`func (o *GridThreatinsight) GetConfigureDomainCollapsingOk() (*bool, bool)`

GetConfigureDomainCollapsingOk returns a tuple with the ConfigureDomainCollapsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureDomainCollapsing

`func (o *GridThreatinsight) SetConfigureDomainCollapsing(v bool)`

SetConfigureDomainCollapsing sets ConfigureDomainCollapsing field to given value.

### HasConfigureDomainCollapsing

`func (o *GridThreatinsight) HasConfigureDomainCollapsing() bool`

HasConfigureDomainCollapsing returns a boolean if a field has been set.

### GetCurrentAllowlist

`func (o *GridThreatinsight) GetCurrentAllowlist() string`

GetCurrentAllowlist returns the CurrentAllowlist field if non-nil, zero value otherwise.

### GetCurrentAllowlistOk

`func (o *GridThreatinsight) GetCurrentAllowlistOk() (*string, bool)`

GetCurrentAllowlistOk returns a tuple with the CurrentAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAllowlist

`func (o *GridThreatinsight) SetCurrentAllowlist(v string)`

SetCurrentAllowlist sets CurrentAllowlist field to given value.

### HasCurrentAllowlist

`func (o *GridThreatinsight) HasCurrentAllowlist() bool`

HasCurrentAllowlist returns a boolean if a field has been set.

### GetCurrentModuleset

`func (o *GridThreatinsight) GetCurrentModuleset() string`

GetCurrentModuleset returns the CurrentModuleset field if non-nil, zero value otherwise.

### GetCurrentModulesetOk

`func (o *GridThreatinsight) GetCurrentModulesetOk() (*string, bool)`

GetCurrentModulesetOk returns a tuple with the CurrentModuleset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentModuleset

`func (o *GridThreatinsight) SetCurrentModuleset(v string)`

SetCurrentModuleset sets CurrentModuleset field to given value.

### HasCurrentModuleset

`func (o *GridThreatinsight) HasCurrentModuleset() bool`

HasCurrentModuleset returns a boolean if a field has been set.

### GetDnsTunnelBlockListRpzZones

`func (o *GridThreatinsight) GetDnsTunnelBlockListRpzZones() []string`

GetDnsTunnelBlockListRpzZones returns the DnsTunnelBlockListRpzZones field if non-nil, zero value otherwise.

### GetDnsTunnelBlockListRpzZonesOk

`func (o *GridThreatinsight) GetDnsTunnelBlockListRpzZonesOk() (*[]string, bool)`

GetDnsTunnelBlockListRpzZonesOk returns a tuple with the DnsTunnelBlockListRpzZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTunnelBlockListRpzZones

`func (o *GridThreatinsight) SetDnsTunnelBlockListRpzZones(v []string)`

SetDnsTunnelBlockListRpzZones sets DnsTunnelBlockListRpzZones field to given value.

### HasDnsTunnelBlockListRpzZones

`func (o *GridThreatinsight) HasDnsTunnelBlockListRpzZones() bool`

HasDnsTunnelBlockListRpzZones returns a boolean if a field has been set.

### GetDomainCollapsingLevel

`func (o *GridThreatinsight) GetDomainCollapsingLevel() int64`

GetDomainCollapsingLevel returns the DomainCollapsingLevel field if non-nil, zero value otherwise.

### GetDomainCollapsingLevelOk

`func (o *GridThreatinsight) GetDomainCollapsingLevelOk() (*int64, bool)`

GetDomainCollapsingLevelOk returns a tuple with the DomainCollapsingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainCollapsingLevel

`func (o *GridThreatinsight) SetDomainCollapsingLevel(v int64)`

SetDomainCollapsingLevel sets DomainCollapsingLevel field to given value.

### HasDomainCollapsingLevel

`func (o *GridThreatinsight) HasDomainCollapsingLevel() bool`

HasDomainCollapsingLevel returns a boolean if a field has been set.

### GetEnableAllowlistAutoDownload

`func (o *GridThreatinsight) GetEnableAllowlistAutoDownload() bool`

GetEnableAllowlistAutoDownload returns the EnableAllowlistAutoDownload field if non-nil, zero value otherwise.

### GetEnableAllowlistAutoDownloadOk

`func (o *GridThreatinsight) GetEnableAllowlistAutoDownloadOk() (*bool, bool)`

GetEnableAllowlistAutoDownloadOk returns a tuple with the EnableAllowlistAutoDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllowlistAutoDownload

`func (o *GridThreatinsight) SetEnableAllowlistAutoDownload(v bool)`

SetEnableAllowlistAutoDownload sets EnableAllowlistAutoDownload field to given value.

### HasEnableAllowlistAutoDownload

`func (o *GridThreatinsight) HasEnableAllowlistAutoDownload() bool`

HasEnableAllowlistAutoDownload returns a boolean if a field has been set.

### GetEnableAllowlistScheduledDownload

`func (o *GridThreatinsight) GetEnableAllowlistScheduledDownload() bool`

GetEnableAllowlistScheduledDownload returns the EnableAllowlistScheduledDownload field if non-nil, zero value otherwise.

### GetEnableAllowlistScheduledDownloadOk

`func (o *GridThreatinsight) GetEnableAllowlistScheduledDownloadOk() (*bool, bool)`

GetEnableAllowlistScheduledDownloadOk returns a tuple with the EnableAllowlistScheduledDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAllowlistScheduledDownload

`func (o *GridThreatinsight) SetEnableAllowlistScheduledDownload(v bool)`

SetEnableAllowlistScheduledDownload sets EnableAllowlistScheduledDownload field to given value.

### HasEnableAllowlistScheduledDownload

`func (o *GridThreatinsight) HasEnableAllowlistScheduledDownload() bool`

HasEnableAllowlistScheduledDownload returns a boolean if a field has been set.

### GetEnableAutoDownload

`func (o *GridThreatinsight) GetEnableAutoDownload() bool`

GetEnableAutoDownload returns the EnableAutoDownload field if non-nil, zero value otherwise.

### GetEnableAutoDownloadOk

`func (o *GridThreatinsight) GetEnableAutoDownloadOk() (*bool, bool)`

GetEnableAutoDownloadOk returns a tuple with the EnableAutoDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoDownload

`func (o *GridThreatinsight) SetEnableAutoDownload(v bool)`

SetEnableAutoDownload sets EnableAutoDownload field to given value.

### HasEnableAutoDownload

`func (o *GridThreatinsight) HasEnableAutoDownload() bool`

HasEnableAutoDownload returns a boolean if a field has been set.

### GetEnableScheduledDownload

`func (o *GridThreatinsight) GetEnableScheduledDownload() bool`

GetEnableScheduledDownload returns the EnableScheduledDownload field if non-nil, zero value otherwise.

### GetEnableScheduledDownloadOk

`func (o *GridThreatinsight) GetEnableScheduledDownloadOk() (*bool, bool)`

GetEnableScheduledDownloadOk returns a tuple with the EnableScheduledDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableScheduledDownload

`func (o *GridThreatinsight) SetEnableScheduledDownload(v bool)`

SetEnableScheduledDownload sets EnableScheduledDownload field to given value.

### HasEnableScheduledDownload

`func (o *GridThreatinsight) HasEnableScheduledDownload() bool`

HasEnableScheduledDownload returns a boolean if a field has been set.

### GetLastAllowlistUpdateTime

`func (o *GridThreatinsight) GetLastAllowlistUpdateTime() int64`

GetLastAllowlistUpdateTime returns the LastAllowlistUpdateTime field if non-nil, zero value otherwise.

### GetLastAllowlistUpdateTimeOk

`func (o *GridThreatinsight) GetLastAllowlistUpdateTimeOk() (*int64, bool)`

GetLastAllowlistUpdateTimeOk returns a tuple with the LastAllowlistUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAllowlistUpdateTime

`func (o *GridThreatinsight) SetLastAllowlistUpdateTime(v int64)`

SetLastAllowlistUpdateTime sets LastAllowlistUpdateTime field to given value.

### HasLastAllowlistUpdateTime

`func (o *GridThreatinsight) HasLastAllowlistUpdateTime() bool`

HasLastAllowlistUpdateTime returns a boolean if a field has been set.

### GetLastAllowlistUpdateVersion

`func (o *GridThreatinsight) GetLastAllowlistUpdateVersion() string`

GetLastAllowlistUpdateVersion returns the LastAllowlistUpdateVersion field if non-nil, zero value otherwise.

### GetLastAllowlistUpdateVersionOk

`func (o *GridThreatinsight) GetLastAllowlistUpdateVersionOk() (*string, bool)`

GetLastAllowlistUpdateVersionOk returns a tuple with the LastAllowlistUpdateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAllowlistUpdateVersion

`func (o *GridThreatinsight) SetLastAllowlistUpdateVersion(v string)`

SetLastAllowlistUpdateVersion sets LastAllowlistUpdateVersion field to given value.

### HasLastAllowlistUpdateVersion

`func (o *GridThreatinsight) HasLastAllowlistUpdateVersion() bool`

HasLastAllowlistUpdateVersion returns a boolean if a field has been set.

### GetLastCheckedForAllowlistUpdate

`func (o *GridThreatinsight) GetLastCheckedForAllowlistUpdate() int64`

GetLastCheckedForAllowlistUpdate returns the LastCheckedForAllowlistUpdate field if non-nil, zero value otherwise.

### GetLastCheckedForAllowlistUpdateOk

`func (o *GridThreatinsight) GetLastCheckedForAllowlistUpdateOk() (*int64, bool)`

GetLastCheckedForAllowlistUpdateOk returns a tuple with the LastCheckedForAllowlistUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedForAllowlistUpdate

`func (o *GridThreatinsight) SetLastCheckedForAllowlistUpdate(v int64)`

SetLastCheckedForAllowlistUpdate sets LastCheckedForAllowlistUpdate field to given value.

### HasLastCheckedForAllowlistUpdate

`func (o *GridThreatinsight) HasLastCheckedForAllowlistUpdate() bool`

HasLastCheckedForAllowlistUpdate returns a boolean if a field has been set.

### GetLastCheckedForPackageUpdate

`func (o *GridThreatinsight) GetLastCheckedForPackageUpdate() int64`

GetLastCheckedForPackageUpdate returns the LastCheckedForPackageUpdate field if non-nil, zero value otherwise.

### GetLastCheckedForPackageUpdateOk

`func (o *GridThreatinsight) GetLastCheckedForPackageUpdateOk() (*int64, bool)`

GetLastCheckedForPackageUpdateOk returns a tuple with the LastCheckedForPackageUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedForPackageUpdate

`func (o *GridThreatinsight) SetLastCheckedForPackageUpdate(v int64)`

SetLastCheckedForPackageUpdate sets LastCheckedForPackageUpdate field to given value.

### HasLastCheckedForPackageUpdate

`func (o *GridThreatinsight) HasLastCheckedForPackageUpdate() bool`

HasLastCheckedForPackageUpdate returns a boolean if a field has been set.

### GetLastCheckedForUpdate

`func (o *GridThreatinsight) GetLastCheckedForUpdate() int64`

GetLastCheckedForUpdate returns the LastCheckedForUpdate field if non-nil, zero value otherwise.

### GetLastCheckedForUpdateOk

`func (o *GridThreatinsight) GetLastCheckedForUpdateOk() (*int64, bool)`

GetLastCheckedForUpdateOk returns a tuple with the LastCheckedForUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedForUpdate

`func (o *GridThreatinsight) SetLastCheckedForUpdate(v int64)`

SetLastCheckedForUpdate sets LastCheckedForUpdate field to given value.

### HasLastCheckedForUpdate

`func (o *GridThreatinsight) HasLastCheckedForUpdate() bool`

HasLastCheckedForUpdate returns a boolean if a field has been set.

### GetLastModuleUpdateTime

`func (o *GridThreatinsight) GetLastModuleUpdateTime() int64`

GetLastModuleUpdateTime returns the LastModuleUpdateTime field if non-nil, zero value otherwise.

### GetLastModuleUpdateTimeOk

`func (o *GridThreatinsight) GetLastModuleUpdateTimeOk() (*int64, bool)`

GetLastModuleUpdateTimeOk returns a tuple with the LastModuleUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModuleUpdateTime

`func (o *GridThreatinsight) SetLastModuleUpdateTime(v int64)`

SetLastModuleUpdateTime sets LastModuleUpdateTime field to given value.

### HasLastModuleUpdateTime

`func (o *GridThreatinsight) HasLastModuleUpdateTime() bool`

HasLastModuleUpdateTime returns a boolean if a field has been set.

### GetLastModuleUpdateVersion

`func (o *GridThreatinsight) GetLastModuleUpdateVersion() string`

GetLastModuleUpdateVersion returns the LastModuleUpdateVersion field if non-nil, zero value otherwise.

### GetLastModuleUpdateVersionOk

`func (o *GridThreatinsight) GetLastModuleUpdateVersionOk() (*string, bool)`

GetLastModuleUpdateVersionOk returns a tuple with the LastModuleUpdateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModuleUpdateVersion

`func (o *GridThreatinsight) SetLastModuleUpdateVersion(v string)`

SetLastModuleUpdateVersion sets LastModuleUpdateVersion field to given value.

### HasLastModuleUpdateVersion

`func (o *GridThreatinsight) HasLastModuleUpdateVersion() bool`

HasLastModuleUpdateVersion returns a boolean if a field has been set.

### GetLastUpdatedPackageVersion

`func (o *GridThreatinsight) GetLastUpdatedPackageVersion() string`

GetLastUpdatedPackageVersion returns the LastUpdatedPackageVersion field if non-nil, zero value otherwise.

### GetLastUpdatedPackageVersionOk

`func (o *GridThreatinsight) GetLastUpdatedPackageVersionOk() (*string, bool)`

GetLastUpdatedPackageVersionOk returns a tuple with the LastUpdatedPackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedPackageVersion

`func (o *GridThreatinsight) SetLastUpdatedPackageVersion(v string)`

SetLastUpdatedPackageVersion sets LastUpdatedPackageVersion field to given value.

### HasLastUpdatedPackageVersion

`func (o *GridThreatinsight) HasLastUpdatedPackageVersion() bool`

HasLastUpdatedPackageVersion returns a boolean if a field has been set.

### GetModuleUpdatePolicy

`func (o *GridThreatinsight) GetModuleUpdatePolicy() string`

GetModuleUpdatePolicy returns the ModuleUpdatePolicy field if non-nil, zero value otherwise.

### GetModuleUpdatePolicyOk

`func (o *GridThreatinsight) GetModuleUpdatePolicyOk() (*string, bool)`

GetModuleUpdatePolicyOk returns a tuple with the ModuleUpdatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleUpdatePolicy

`func (o *GridThreatinsight) SetModuleUpdatePolicy(v string)`

SetModuleUpdatePolicy sets ModuleUpdatePolicy field to given value.

### HasModuleUpdatePolicy

`func (o *GridThreatinsight) HasModuleUpdatePolicy() bool`

HasModuleUpdatePolicy returns a boolean if a field has been set.

### GetName

`func (o *GridThreatinsight) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GridThreatinsight) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GridThreatinsight) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GridThreatinsight) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheduledAllowlistDownload

`func (o *GridThreatinsight) GetScheduledAllowlistDownload() GridThreatinsightScheduledAllowlistDownload`

GetScheduledAllowlistDownload returns the ScheduledAllowlistDownload field if non-nil, zero value otherwise.

### GetScheduledAllowlistDownloadOk

`func (o *GridThreatinsight) GetScheduledAllowlistDownloadOk() (*GridThreatinsightScheduledAllowlistDownload, bool)`

GetScheduledAllowlistDownloadOk returns a tuple with the ScheduledAllowlistDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAllowlistDownload

`func (o *GridThreatinsight) SetScheduledAllowlistDownload(v GridThreatinsightScheduledAllowlistDownload)`

SetScheduledAllowlistDownload sets ScheduledAllowlistDownload field to given value.

### HasScheduledAllowlistDownload

`func (o *GridThreatinsight) HasScheduledAllowlistDownload() bool`

HasScheduledAllowlistDownload returns a boolean if a field has been set.

### GetScheduledDownload

`func (o *GridThreatinsight) GetScheduledDownload() GridThreatinsightScheduledDownload`

GetScheduledDownload returns the ScheduledDownload field if non-nil, zero value otherwise.

### GetScheduledDownloadOk

`func (o *GridThreatinsight) GetScheduledDownloadOk() (*GridThreatinsightScheduledDownload, bool)`

GetScheduledDownloadOk returns a tuple with the ScheduledDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDownload

`func (o *GridThreatinsight) SetScheduledDownload(v GridThreatinsightScheduledDownload)`

SetScheduledDownload sets ScheduledDownload field to given value.

### HasScheduledDownload

`func (o *GridThreatinsight) HasScheduledDownload() bool`

HasScheduledDownload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


