# RecordUnknown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**CloudInfo** | Pointer to [**RecordUnknownCloudInfo**](RecordUnknownCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the record; maximum 256 characters. | [optional] 
**Creator** | Pointer to **string** | The record creator. Note that changing creator from or to &#39;SYSTEM&#39; value is not allowed. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**DisplayRdata** | Pointer to **string** | Standard textual representation of the RDATA. | [optional] [readonly] 
**DnsName** | Pointer to **string** | The name of the unknown record in punycode format. | [optional] [readonly] 
**EnableHostNamePolicy** | Pointer to **bool** | Determines if host name policy is applicable for the record. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**Name** | Pointer to **string** | The Unknown record name in FQDN format. This value can be in unicode format. | [optional] 
**Policy** | Pointer to **string** | The host name policy for the record. | [optional] [readonly] 
**RecordType** | Pointer to **string** | Specifies type of unknown resource record. | [optional] 
**SubfieldValues** | Pointer to [**[]RecordUnknownSubfieldValues**](RecordUnknownSubfieldValues.md) | The list of rdata subfield values of unknown resource record. | [optional] 
**Ttl** | Pointer to **int64** | The Time to Live (TTL) value for the record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**View** | Pointer to **string** | The name of the DNS view in which the record resides. Example: \&quot;external\&quot;. | [optional] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 

## Methods

### NewRecordUnknown

`func NewRecordUnknown() *RecordUnknown`

NewRecordUnknown instantiates a new RecordUnknown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordUnknownWithDefaults

`func NewRecordUnknownWithDefaults() *RecordUnknown`

NewRecordUnknownWithDefaults instantiates a new RecordUnknown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordUnknown) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordUnknown) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordUnknown) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordUnknown) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCloudInfo

`func (o *RecordUnknown) GetCloudInfo() RecordUnknownCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *RecordUnknown) GetCloudInfoOk() (*RecordUnknownCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *RecordUnknown) SetCloudInfo(v RecordUnknownCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *RecordUnknown) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *RecordUnknown) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordUnknown) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordUnknown) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordUnknown) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreator

`func (o *RecordUnknown) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RecordUnknown) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RecordUnknown) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RecordUnknown) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDisable

`func (o *RecordUnknown) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *RecordUnknown) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *RecordUnknown) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *RecordUnknown) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDisplayRdata

`func (o *RecordUnknown) GetDisplayRdata() string`

GetDisplayRdata returns the DisplayRdata field if non-nil, zero value otherwise.

### GetDisplayRdataOk

`func (o *RecordUnknown) GetDisplayRdataOk() (*string, bool)`

GetDisplayRdataOk returns a tuple with the DisplayRdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayRdata

`func (o *RecordUnknown) SetDisplayRdata(v string)`

SetDisplayRdata sets DisplayRdata field to given value.

### HasDisplayRdata

`func (o *RecordUnknown) HasDisplayRdata() bool`

HasDisplayRdata returns a boolean if a field has been set.

### GetDnsName

`func (o *RecordUnknown) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *RecordUnknown) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *RecordUnknown) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *RecordUnknown) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetEnableHostNamePolicy

`func (o *RecordUnknown) GetEnableHostNamePolicy() bool`

GetEnableHostNamePolicy returns the EnableHostNamePolicy field if non-nil, zero value otherwise.

### GetEnableHostNamePolicyOk

`func (o *RecordUnknown) GetEnableHostNamePolicyOk() (*bool, bool)`

GetEnableHostNamePolicyOk returns a tuple with the EnableHostNamePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHostNamePolicy

`func (o *RecordUnknown) SetEnableHostNamePolicy(v bool)`

SetEnableHostNamePolicy sets EnableHostNamePolicy field to given value.

### HasEnableHostNamePolicy

`func (o *RecordUnknown) HasEnableHostNamePolicy() bool`

HasEnableHostNamePolicy returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *RecordUnknown) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *RecordUnknown) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *RecordUnknown) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *RecordUnknown) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *RecordUnknown) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *RecordUnknown) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *RecordUnknown) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *RecordUnknown) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *RecordUnknown) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *RecordUnknown) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *RecordUnknown) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *RecordUnknown) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetLastQueried

`func (o *RecordUnknown) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *RecordUnknown) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *RecordUnknown) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *RecordUnknown) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetName

`func (o *RecordUnknown) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordUnknown) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordUnknown) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordUnknown) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicy

`func (o *RecordUnknown) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *RecordUnknown) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *RecordUnknown) SetPolicy(v string)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *RecordUnknown) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetRecordType

`func (o *RecordUnknown) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *RecordUnknown) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *RecordUnknown) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *RecordUnknown) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetSubfieldValues

`func (o *RecordUnknown) GetSubfieldValues() []RecordUnknownSubfieldValues`

GetSubfieldValues returns the SubfieldValues field if non-nil, zero value otherwise.

### GetSubfieldValuesOk

`func (o *RecordUnknown) GetSubfieldValuesOk() (*[]RecordUnknownSubfieldValues, bool)`

GetSubfieldValuesOk returns a tuple with the SubfieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubfieldValues

`func (o *RecordUnknown) SetSubfieldValues(v []RecordUnknownSubfieldValues)`

SetSubfieldValues sets SubfieldValues field to given value.

### HasSubfieldValues

`func (o *RecordUnknown) HasSubfieldValues() bool`

HasSubfieldValues returns a boolean if a field has been set.

### GetTtl

`func (o *RecordUnknown) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordUnknown) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordUnknown) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordUnknown) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordUnknown) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordUnknown) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordUnknown) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordUnknown) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordUnknown) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordUnknown) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordUnknown) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordUnknown) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordUnknown) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordUnknown) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordUnknown) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordUnknown) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


