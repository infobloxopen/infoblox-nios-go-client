# RecordTlsa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**CertificateData** | Pointer to **string** | Hex dump of either raw data for matching type 0, or the hash of the raw data for matching types 1 and 2. | [optional] 
**CertificateUsage** | Pointer to **int64** | Specifies the provided association that will be used to match the certificate presented in the TLS handshake. Based on RFC-6698. | [optional] 
**CloudInfo** | Pointer to [**RecordTlsaCloudInfo**](RecordTlsaCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the record; maximum 256 characters. | [optional] 
**Creator** | Pointer to **string** | The record creator. Note that changing creator from or to &#39;SYSTEM&#39; value is not allowed. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**DnsName** | Pointer to **string** | The name of the TLSA record in punycode format. | [optional] [readonly] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**MatchedType** | Pointer to **int64** | Specifies how the certificate association is presented. Based on RFC-6698. | [optional] 
**Name** | Pointer to **string** | The TLSA record name in FQDN format. This value can be in unicode format. | [optional] 
**Selector** | Pointer to **int64** | Specifies which part of the TLS certificate presented by the server will be matched against the association data. Based on RFC-6698. | [optional] 
**Ttl** | Pointer to **int64** | The Time to Live (TTL) value for the record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**View** | Pointer to **string** | The name of the DNS view in which the record resides. Example: \&quot;external\&quot;. | [optional] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 

## Methods

### NewRecordTlsa

`func NewRecordTlsa() *RecordTlsa`

NewRecordTlsa instantiates a new RecordTlsa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordTlsaWithDefaults

`func NewRecordTlsaWithDefaults() *RecordTlsa`

NewRecordTlsaWithDefaults instantiates a new RecordTlsa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordTlsa) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordTlsa) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordTlsa) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordTlsa) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCertificateData

`func (o *RecordTlsa) GetCertificateData() string`

GetCertificateData returns the CertificateData field if non-nil, zero value otherwise.

### GetCertificateDataOk

`func (o *RecordTlsa) GetCertificateDataOk() (*string, bool)`

GetCertificateDataOk returns a tuple with the CertificateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateData

`func (o *RecordTlsa) SetCertificateData(v string)`

SetCertificateData sets CertificateData field to given value.

### HasCertificateData

`func (o *RecordTlsa) HasCertificateData() bool`

HasCertificateData returns a boolean if a field has been set.

### GetCertificateUsage

`func (o *RecordTlsa) GetCertificateUsage() int64`

GetCertificateUsage returns the CertificateUsage field if non-nil, zero value otherwise.

### GetCertificateUsageOk

`func (o *RecordTlsa) GetCertificateUsageOk() (*int64, bool)`

GetCertificateUsageOk returns a tuple with the CertificateUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateUsage

`func (o *RecordTlsa) SetCertificateUsage(v int64)`

SetCertificateUsage sets CertificateUsage field to given value.

### HasCertificateUsage

`func (o *RecordTlsa) HasCertificateUsage() bool`

HasCertificateUsage returns a boolean if a field has been set.

### GetCloudInfo

`func (o *RecordTlsa) GetCloudInfo() RecordTlsaCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *RecordTlsa) GetCloudInfoOk() (*RecordTlsaCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *RecordTlsa) SetCloudInfo(v RecordTlsaCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *RecordTlsa) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *RecordTlsa) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordTlsa) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordTlsa) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordTlsa) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreator

`func (o *RecordTlsa) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RecordTlsa) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RecordTlsa) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RecordTlsa) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDisable

`func (o *RecordTlsa) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *RecordTlsa) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *RecordTlsa) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *RecordTlsa) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDnsName

`func (o *RecordTlsa) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *RecordTlsa) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *RecordTlsa) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *RecordTlsa) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *RecordTlsa) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *RecordTlsa) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *RecordTlsa) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *RecordTlsa) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *RecordTlsa) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *RecordTlsa) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *RecordTlsa) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *RecordTlsa) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *RecordTlsa) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *RecordTlsa) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *RecordTlsa) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *RecordTlsa) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetLastQueried

`func (o *RecordTlsa) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *RecordTlsa) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *RecordTlsa) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *RecordTlsa) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetMatchedType

`func (o *RecordTlsa) GetMatchedType() int64`

GetMatchedType returns the MatchedType field if non-nil, zero value otherwise.

### GetMatchedTypeOk

`func (o *RecordTlsa) GetMatchedTypeOk() (*int64, bool)`

GetMatchedTypeOk returns a tuple with the MatchedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedType

`func (o *RecordTlsa) SetMatchedType(v int64)`

SetMatchedType sets MatchedType field to given value.

### HasMatchedType

`func (o *RecordTlsa) HasMatchedType() bool`

HasMatchedType returns a boolean if a field has been set.

### GetName

`func (o *RecordTlsa) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordTlsa) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordTlsa) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordTlsa) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSelector

`func (o *RecordTlsa) GetSelector() int64`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *RecordTlsa) GetSelectorOk() (*int64, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *RecordTlsa) SetSelector(v int64)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *RecordTlsa) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetTtl

`func (o *RecordTlsa) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordTlsa) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordTlsa) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordTlsa) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordTlsa) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordTlsa) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordTlsa) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordTlsa) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordTlsa) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordTlsa) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordTlsa) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordTlsa) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordTlsa) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordTlsa) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordTlsa) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordTlsa) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


