# GetRecordTlsaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
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
**Result** | Pointer to [**RecordTlsa**](RecordTlsa.md) |  | [optional] 

## Methods

### NewGetRecordTlsaResponse

`func NewGetRecordTlsaResponse() *GetRecordTlsaResponse`

NewGetRecordTlsaResponse instantiates a new GetRecordTlsaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRecordTlsaResponseWithDefaults

`func NewGetRecordTlsaResponseWithDefaults() *GetRecordTlsaResponse`

NewGetRecordTlsaResponseWithDefaults instantiates a new GetRecordTlsaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetRecordTlsaResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetRecordTlsaResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetRecordTlsaResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetRecordTlsaResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetCertificateData

`func (o *GetRecordTlsaResponse) GetCertificateData() string`

GetCertificateData returns the CertificateData field if non-nil, zero value otherwise.

### GetCertificateDataOk

`func (o *GetRecordTlsaResponse) GetCertificateDataOk() (*string, bool)`

GetCertificateDataOk returns a tuple with the CertificateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateData

`func (o *GetRecordTlsaResponse) SetCertificateData(v string)`

SetCertificateData sets CertificateData field to given value.

### HasCertificateData

`func (o *GetRecordTlsaResponse) HasCertificateData() bool`

HasCertificateData returns a boolean if a field has been set.

### GetCertificateUsage

`func (o *GetRecordTlsaResponse) GetCertificateUsage() int64`

GetCertificateUsage returns the CertificateUsage field if non-nil, zero value otherwise.

### GetCertificateUsageOk

`func (o *GetRecordTlsaResponse) GetCertificateUsageOk() (*int64, bool)`

GetCertificateUsageOk returns a tuple with the CertificateUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateUsage

`func (o *GetRecordTlsaResponse) SetCertificateUsage(v int64)`

SetCertificateUsage sets CertificateUsage field to given value.

### HasCertificateUsage

`func (o *GetRecordTlsaResponse) HasCertificateUsage() bool`

HasCertificateUsage returns a boolean if a field has been set.

### GetCloudInfo

`func (o *GetRecordTlsaResponse) GetCloudInfo() RecordTlsaCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *GetRecordTlsaResponse) GetCloudInfoOk() (*RecordTlsaCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *GetRecordTlsaResponse) SetCloudInfo(v RecordTlsaCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *GetRecordTlsaResponse) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *GetRecordTlsaResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetRecordTlsaResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetRecordTlsaResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetRecordTlsaResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreator

`func (o *GetRecordTlsaResponse) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *GetRecordTlsaResponse) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *GetRecordTlsaResponse) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *GetRecordTlsaResponse) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDisable

`func (o *GetRecordTlsaResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetRecordTlsaResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetRecordTlsaResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetRecordTlsaResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDnsName

`func (o *GetRecordTlsaResponse) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *GetRecordTlsaResponse) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *GetRecordTlsaResponse) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *GetRecordTlsaResponse) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetRecordTlsaResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetRecordTlsaResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetRecordTlsaResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetRecordTlsaResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetRecordTlsaResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetRecordTlsaResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetRecordTlsaResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetRecordTlsaResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetRecordTlsaResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetRecordTlsaResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetRecordTlsaResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetRecordTlsaResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetLastQueried

`func (o *GetRecordTlsaResponse) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *GetRecordTlsaResponse) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *GetRecordTlsaResponse) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *GetRecordTlsaResponse) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetMatchedType

`func (o *GetRecordTlsaResponse) GetMatchedType() int64`

GetMatchedType returns the MatchedType field if non-nil, zero value otherwise.

### GetMatchedTypeOk

`func (o *GetRecordTlsaResponse) GetMatchedTypeOk() (*int64, bool)`

GetMatchedTypeOk returns a tuple with the MatchedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedType

`func (o *GetRecordTlsaResponse) SetMatchedType(v int64)`

SetMatchedType sets MatchedType field to given value.

### HasMatchedType

`func (o *GetRecordTlsaResponse) HasMatchedType() bool`

HasMatchedType returns a boolean if a field has been set.

### GetName

`func (o *GetRecordTlsaResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRecordTlsaResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRecordTlsaResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetRecordTlsaResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSelector

`func (o *GetRecordTlsaResponse) GetSelector() int64`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *GetRecordTlsaResponse) GetSelectorOk() (*int64, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *GetRecordTlsaResponse) SetSelector(v int64)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *GetRecordTlsaResponse) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetTtl

`func (o *GetRecordTlsaResponse) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetRecordTlsaResponse) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetRecordTlsaResponse) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetRecordTlsaResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *GetRecordTlsaResponse) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *GetRecordTlsaResponse) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *GetRecordTlsaResponse) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *GetRecordTlsaResponse) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *GetRecordTlsaResponse) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetRecordTlsaResponse) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetRecordTlsaResponse) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *GetRecordTlsaResponse) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *GetRecordTlsaResponse) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetRecordTlsaResponse) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetRecordTlsaResponse) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetRecordTlsaResponse) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetResult

`func (o *GetRecordTlsaResponse) GetResult() RecordTlsa`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetRecordTlsaResponse) GetResultOk() (*RecordTlsa, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetRecordTlsaResponse) SetResult(v RecordTlsa)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetRecordTlsaResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


