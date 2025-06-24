# RecordAlias

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AwsRte53RecordInfo** | Pointer to [**RecordAliasAwsRte53RecordInfo**](RecordAliasAwsRte53RecordInfo.md) |  | [optional] 
**CloudInfo** | Pointer to [**RecordAliasCloudInfo**](RecordAliasCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the record; maximum 256 characters. | [optional] 
**Creator** | Pointer to **string** | The record creator. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**DnsName** | Pointer to **string** | The name for an Alias record in punycode format. | [optional] [readonly] 
**DnsTargetName** | Pointer to **string** | Target name in punycode format. | [optional] [readonly] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**Name** | Pointer to **string** | The name for an Alias record in FQDN format. This value can be in unicode format. Regular expression search is not supported for unicode values. | [optional] 
**TargetName** | Pointer to **string** | Target name in FQDN format. This value can be in unicode format. | [optional] 
**TargetType** | Pointer to **string** | Target type. | [optional] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**View** | Pointer to **string** | The name of the DNS View in which the record resides. Example: \&quot;external\&quot;. | [optional] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 

## Methods

### NewRecordAlias

`func NewRecordAlias() *RecordAlias`

NewRecordAlias instantiates a new RecordAlias object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordAliasWithDefaults

`func NewRecordAliasWithDefaults() *RecordAlias`

NewRecordAliasWithDefaults instantiates a new RecordAlias object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordAlias) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordAlias) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordAlias) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordAlias) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAwsRte53RecordInfo

`func (o *RecordAlias) GetAwsRte53RecordInfo() RecordAliasAwsRte53RecordInfo`

GetAwsRte53RecordInfo returns the AwsRte53RecordInfo field if non-nil, zero value otherwise.

### GetAwsRte53RecordInfoOk

`func (o *RecordAlias) GetAwsRte53RecordInfoOk() (*RecordAliasAwsRte53RecordInfo, bool)`

GetAwsRte53RecordInfoOk returns a tuple with the AwsRte53RecordInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRte53RecordInfo

`func (o *RecordAlias) SetAwsRte53RecordInfo(v RecordAliasAwsRte53RecordInfo)`

SetAwsRte53RecordInfo sets AwsRte53RecordInfo field to given value.

### HasAwsRte53RecordInfo

`func (o *RecordAlias) HasAwsRte53RecordInfo() bool`

HasAwsRte53RecordInfo returns a boolean if a field has been set.

### GetCloudInfo

`func (o *RecordAlias) GetCloudInfo() RecordAliasCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *RecordAlias) GetCloudInfoOk() (*RecordAliasCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *RecordAlias) SetCloudInfo(v RecordAliasCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *RecordAlias) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *RecordAlias) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordAlias) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordAlias) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordAlias) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreator

`func (o *RecordAlias) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RecordAlias) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RecordAlias) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RecordAlias) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDisable

`func (o *RecordAlias) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *RecordAlias) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *RecordAlias) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *RecordAlias) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDnsName

`func (o *RecordAlias) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *RecordAlias) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *RecordAlias) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *RecordAlias) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetDnsTargetName

`func (o *RecordAlias) GetDnsTargetName() string`

GetDnsTargetName returns the DnsTargetName field if non-nil, zero value otherwise.

### GetDnsTargetNameOk

`func (o *RecordAlias) GetDnsTargetNameOk() (*string, bool)`

GetDnsTargetNameOk returns a tuple with the DnsTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTargetName

`func (o *RecordAlias) SetDnsTargetName(v string)`

SetDnsTargetName sets DnsTargetName field to given value.

### HasDnsTargetName

`func (o *RecordAlias) HasDnsTargetName() bool`

HasDnsTargetName returns a boolean if a field has been set.

### GetExtAttrs

`func (o *RecordAlias) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *RecordAlias) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *RecordAlias) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *RecordAlias) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetLastQueried

`func (o *RecordAlias) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *RecordAlias) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *RecordAlias) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *RecordAlias) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetName

`func (o *RecordAlias) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordAlias) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordAlias) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordAlias) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTargetName

`func (o *RecordAlias) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *RecordAlias) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *RecordAlias) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *RecordAlias) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetTargetType

`func (o *RecordAlias) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *RecordAlias) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *RecordAlias) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *RecordAlias) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetTtl

`func (o *RecordAlias) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordAlias) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordAlias) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordAlias) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordAlias) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordAlias) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordAlias) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordAlias) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordAlias) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordAlias) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordAlias) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordAlias) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordAlias) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordAlias) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordAlias) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordAlias) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


