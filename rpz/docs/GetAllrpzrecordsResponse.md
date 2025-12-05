# GetAllrpzrecordsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AlertType** | Pointer to **string** | The alert type of the record associated with the allrpzrecords object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The descriptive comment of the record associated with the allrpzrecords object. | [optional] [readonly] 
**Disable** | Pointer to **bool** | The disable flag of the record associated with the allrpzrecords object (if present). | [optional] [readonly] 
**ExpirationTime** | Pointer to **int64** | The expiration time of the record associated with the allrpzrecords object. | [optional] [readonly] 
**LastUpdated** | Pointer to **int64** | The time when the record associated with the allrpzrecords object was last updated. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the record associated with the allrpzrecords object. Note that this value might be different than the value of the name field for the associated record. | [optional] [readonly] 
**Record** | Pointer to **string** | The record object associated with the allrpzrecords object. | [optional] [readonly] 
**RpzRule** | Pointer to **string** | The RPZ rule type of the record associated with the allrpzrecrods object. | [optional] [readonly] 
**Ttl** | Pointer to **int64** | The TTL value of the record associated with the allrpzrecords object (if present). | [optional] [readonly] 
**Type** | Pointer to **string** | The type of record associated with the allrpzrecords object. This is a descriptive string that identifies the record to which this allrpzrecords object refers. (Examples: &#39;record:rpz:a&#39;, &#39;record:rpz:mx&#39;, etc.) | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**View** | Pointer to **string** | The DNS view name of the record associated with the allrpzrecords object. | [optional] [readonly] 
**Zone** | Pointer to **string** | The Response Policy Zone name of the record associated with the allrpzrecords object. | [optional] [readonly] 
**Result** | Pointer to [**Allrpzrecords**](Allrpzrecords.md) |  | [optional] 

## Methods

### NewGetAllrpzrecordsResponse

`func NewGetAllrpzrecordsResponse() *GetAllrpzrecordsResponse`

NewGetAllrpzrecordsResponse instantiates a new GetAllrpzrecordsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllrpzrecordsResponseWithDefaults

`func NewGetAllrpzrecordsResponseWithDefaults() *GetAllrpzrecordsResponse`

NewGetAllrpzrecordsResponseWithDefaults instantiates a new GetAllrpzrecordsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetAllrpzrecordsResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetAllrpzrecordsResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetAllrpzrecordsResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetAllrpzrecordsResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAlertType

`func (o *GetAllrpzrecordsResponse) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *GetAllrpzrecordsResponse) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *GetAllrpzrecordsResponse) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *GetAllrpzrecordsResponse) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetComment

`func (o *GetAllrpzrecordsResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetAllrpzrecordsResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetAllrpzrecordsResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetAllrpzrecordsResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetAllrpzrecordsResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetAllrpzrecordsResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetAllrpzrecordsResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetAllrpzrecordsResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExpirationTime

`func (o *GetAllrpzrecordsResponse) GetExpirationTime() int64`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *GetAllrpzrecordsResponse) GetExpirationTimeOk() (*int64, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *GetAllrpzrecordsResponse) SetExpirationTime(v int64)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *GetAllrpzrecordsResponse) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetAllrpzrecordsResponse) GetLastUpdated() int64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetAllrpzrecordsResponse) GetLastUpdatedOk() (*int64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetAllrpzrecordsResponse) SetLastUpdated(v int64)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetAllrpzrecordsResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *GetAllrpzrecordsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAllrpzrecordsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAllrpzrecordsResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAllrpzrecordsResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecord

`func (o *GetAllrpzrecordsResponse) GetRecord() string`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *GetAllrpzrecordsResponse) GetRecordOk() (*string, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *GetAllrpzrecordsResponse) SetRecord(v string)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *GetAllrpzrecordsResponse) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetRpzRule

`func (o *GetAllrpzrecordsResponse) GetRpzRule() string`

GetRpzRule returns the RpzRule field if non-nil, zero value otherwise.

### GetRpzRuleOk

`func (o *GetAllrpzrecordsResponse) GetRpzRuleOk() (*string, bool)`

GetRpzRuleOk returns a tuple with the RpzRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzRule

`func (o *GetAllrpzrecordsResponse) SetRpzRule(v string)`

SetRpzRule sets RpzRule field to given value.

### HasRpzRule

`func (o *GetAllrpzrecordsResponse) HasRpzRule() bool`

HasRpzRule returns a boolean if a field has been set.

### GetTtl

`func (o *GetAllrpzrecordsResponse) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetAllrpzrecordsResponse) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetAllrpzrecordsResponse) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetAllrpzrecordsResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetType

`func (o *GetAllrpzrecordsResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetAllrpzrecordsResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetAllrpzrecordsResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetAllrpzrecordsResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *GetAllrpzrecordsResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetAllrpzrecordsResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetAllrpzrecordsResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetAllrpzrecordsResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetView

`func (o *GetAllrpzrecordsResponse) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetAllrpzrecordsResponse) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetAllrpzrecordsResponse) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *GetAllrpzrecordsResponse) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *GetAllrpzrecordsResponse) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetAllrpzrecordsResponse) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetAllrpzrecordsResponse) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetAllrpzrecordsResponse) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetResult

`func (o *GetAllrpzrecordsResponse) GetResult() Allrpzrecords`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetAllrpzrecordsResponse) GetResultOk() (*Allrpzrecords, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetAllrpzrecordsResponse) SetResult(v Allrpzrecords)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetAllrpzrecordsResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


