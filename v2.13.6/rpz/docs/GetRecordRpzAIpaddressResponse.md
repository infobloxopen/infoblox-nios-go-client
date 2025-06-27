# GetRecordRpzAIpaddressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | The comment for the record; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Ipv4addr** | Pointer to **string** | The IPv4 Address of the substitute rule. | [optional] 
**Name** | Pointer to **string** | The name for a record in FQDN format. This value cannot be in unicode format. | [optional] 
**RpZone** | Pointer to **string** | The name of a response policy zone in which the record resides. | [optional] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**View** | Pointer to **string** | The name of the DNS View in which the record resides. Example: \&quot;external\&quot;. | [optional] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 
**Result** | Pointer to [**RecordRpzAIpaddress**](RecordRpzAIpaddress.md) |  | [optional] 

## Methods

### NewGetRecordRpzAIpaddressResponse

`func NewGetRecordRpzAIpaddressResponse() *GetRecordRpzAIpaddressResponse`

NewGetRecordRpzAIpaddressResponse instantiates a new GetRecordRpzAIpaddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRecordRpzAIpaddressResponseWithDefaults

`func NewGetRecordRpzAIpaddressResponseWithDefaults() *GetRecordRpzAIpaddressResponse`

NewGetRecordRpzAIpaddressResponseWithDefaults instantiates a new GetRecordRpzAIpaddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetRecordRpzAIpaddressResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetRecordRpzAIpaddressResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetRecordRpzAIpaddressResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetRecordRpzAIpaddressResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetRecordRpzAIpaddressResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetRecordRpzAIpaddressResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetRecordRpzAIpaddressResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetRecordRpzAIpaddressResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *GetRecordRpzAIpaddressResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetRecordRpzAIpaddressResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetRecordRpzAIpaddressResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetRecordRpzAIpaddressResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetRecordRpzAIpaddressResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetRecordRpzAIpaddressResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetRecordRpzAIpaddressResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetRecordRpzAIpaddressResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetIpv4addr

`func (o *GetRecordRpzAIpaddressResponse) GetIpv4addr() string`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *GetRecordRpzAIpaddressResponse) GetIpv4addrOk() (*string, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *GetRecordRpzAIpaddressResponse) SetIpv4addr(v string)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *GetRecordRpzAIpaddressResponse) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### GetName

`func (o *GetRecordRpzAIpaddressResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRecordRpzAIpaddressResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRecordRpzAIpaddressResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetRecordRpzAIpaddressResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRpZone

`func (o *GetRecordRpzAIpaddressResponse) GetRpZone() string`

GetRpZone returns the RpZone field if non-nil, zero value otherwise.

### GetRpZoneOk

`func (o *GetRecordRpzAIpaddressResponse) GetRpZoneOk() (*string, bool)`

GetRpZoneOk returns a tuple with the RpZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpZone

`func (o *GetRecordRpzAIpaddressResponse) SetRpZone(v string)`

SetRpZone sets RpZone field to given value.

### HasRpZone

`func (o *GetRecordRpzAIpaddressResponse) HasRpZone() bool`

HasRpZone returns a boolean if a field has been set.

### GetTtl

`func (o *GetRecordRpzAIpaddressResponse) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetRecordRpzAIpaddressResponse) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetRecordRpzAIpaddressResponse) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetRecordRpzAIpaddressResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *GetRecordRpzAIpaddressResponse) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *GetRecordRpzAIpaddressResponse) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *GetRecordRpzAIpaddressResponse) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *GetRecordRpzAIpaddressResponse) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *GetRecordRpzAIpaddressResponse) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetRecordRpzAIpaddressResponse) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetRecordRpzAIpaddressResponse) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *GetRecordRpzAIpaddressResponse) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *GetRecordRpzAIpaddressResponse) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetRecordRpzAIpaddressResponse) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetRecordRpzAIpaddressResponse) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetRecordRpzAIpaddressResponse) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetResult

`func (o *GetRecordRpzAIpaddressResponse) GetResult() RecordRpzAIpaddress`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetRecordRpzAIpaddressResponse) GetResultOk() (*RecordRpzAIpaddress, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetRecordRpzAIpaddressResponse) SetResult(v RecordRpzAIpaddress)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetRecordRpzAIpaddressResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


