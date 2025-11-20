# Allrpzrecords

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
**View** | Pointer to **string** | The DNS view name of the record associated with the allrpzrecords object. | [optional] [readonly] 
**Zone** | Pointer to **string** | The Response Policy Zone name of the record associated with the allrpzrecords object. | [optional] [readonly] 

## Methods

### NewAllrpzrecords

`func NewAllrpzrecords() *Allrpzrecords`

NewAllrpzrecords instantiates a new Allrpzrecords object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllrpzrecordsWithDefaults

`func NewAllrpzrecordsWithDefaults() *Allrpzrecords`

NewAllrpzrecordsWithDefaults instantiates a new Allrpzrecords object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Allrpzrecords) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Allrpzrecords) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Allrpzrecords) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Allrpzrecords) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAlertType

`func (o *Allrpzrecords) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *Allrpzrecords) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *Allrpzrecords) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *Allrpzrecords) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetComment

`func (o *Allrpzrecords) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Allrpzrecords) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Allrpzrecords) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Allrpzrecords) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *Allrpzrecords) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *Allrpzrecords) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *Allrpzrecords) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *Allrpzrecords) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExpirationTime

`func (o *Allrpzrecords) GetExpirationTime() int64`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *Allrpzrecords) GetExpirationTimeOk() (*int64, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *Allrpzrecords) SetExpirationTime(v int64)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *Allrpzrecords) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Allrpzrecords) GetLastUpdated() int64`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Allrpzrecords) GetLastUpdatedOk() (*int64, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Allrpzrecords) SetLastUpdated(v int64)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Allrpzrecords) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetName

`func (o *Allrpzrecords) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Allrpzrecords) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Allrpzrecords) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Allrpzrecords) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecord

`func (o *Allrpzrecords) GetRecord() string`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *Allrpzrecords) GetRecordOk() (*string, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *Allrpzrecords) SetRecord(v string)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *Allrpzrecords) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetRpzRule

`func (o *Allrpzrecords) GetRpzRule() string`

GetRpzRule returns the RpzRule field if non-nil, zero value otherwise.

### GetRpzRuleOk

`func (o *Allrpzrecords) GetRpzRuleOk() (*string, bool)`

GetRpzRuleOk returns a tuple with the RpzRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzRule

`func (o *Allrpzrecords) SetRpzRule(v string)`

SetRpzRule sets RpzRule field to given value.

### HasRpzRule

`func (o *Allrpzrecords) HasRpzRule() bool`

HasRpzRule returns a boolean if a field has been set.

### GetTtl

`func (o *Allrpzrecords) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *Allrpzrecords) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *Allrpzrecords) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *Allrpzrecords) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetType

`func (o *Allrpzrecords) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Allrpzrecords) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Allrpzrecords) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Allrpzrecords) HasType() bool`

HasType returns a boolean if a field has been set.

### GetView

`func (o *Allrpzrecords) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *Allrpzrecords) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *Allrpzrecords) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *Allrpzrecords) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *Allrpzrecords) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Allrpzrecords) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Allrpzrecords) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Allrpzrecords) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


