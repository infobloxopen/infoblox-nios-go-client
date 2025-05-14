# GetRecordDnskeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Algorithm** | Pointer to **string** | The public key encryption algorithm of a DNSKEY Record object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The comment for the record. | [optional] [readonly] 
**CreationTime** | Pointer to **int64** | The creation time of the record. | [optional] [readonly] 
**Creator** | Pointer to **string** | The record creator. | [optional] [readonly] 
**DnsName** | Pointer to **string** | Name of a DNSKEY record in punycode format. | [optional] [readonly] 
**Flags** | Pointer to **string** | The flags field is a 16-bit unsigned integer. Currently, only two bits of this value are used: the least significant bit and bit 7. The other bits are reserved for future use and must be zero. If bit 7 is set to 1, the key is a DNS zone key. Otherwise, the key is not a zone key and cannot be used to verify zone data. The least significant bit indicates \&quot;secure entry point property\&quot;. If it is not zero, the key is a key signing key (KSK type). Otherwise, the key type is ZSK. | [optional] [readonly] 
**KeyTag** | Pointer to **int64** | The key tag identifying the public key of a DNSKEY Record object. | [optional] [readonly] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the DNSKEY record in FQDN format. It has to be the same as the zone, where the record resides. | [optional] [readonly] 
**PublicKey** | Pointer to **string** | The public key. The format of the returned value depends on the key algorithm. | [optional] [readonly] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for the record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] [readonly] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] [readonly] 
**View** | Pointer to **string** | The name of the DNS View in which the record resides. Example: \&quot;external\&quot;. | [optional] [readonly] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 
**Result** | Pointer to [**RecordDnskey**](RecordDnskey.md) |  | [optional] 

## Methods

### NewGetRecordDnskeyResponse

`func NewGetRecordDnskeyResponse() *GetRecordDnskeyResponse`

NewGetRecordDnskeyResponse instantiates a new GetRecordDnskeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRecordDnskeyResponseWithDefaults

`func NewGetRecordDnskeyResponseWithDefaults() *GetRecordDnskeyResponse`

NewGetRecordDnskeyResponseWithDefaults instantiates a new GetRecordDnskeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetRecordDnskeyResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetRecordDnskeyResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetRecordDnskeyResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetRecordDnskeyResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAlgorithm

`func (o *GetRecordDnskeyResponse) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *GetRecordDnskeyResponse) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *GetRecordDnskeyResponse) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *GetRecordDnskeyResponse) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetComment

`func (o *GetRecordDnskeyResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetRecordDnskeyResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetRecordDnskeyResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetRecordDnskeyResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreationTime

`func (o *GetRecordDnskeyResponse) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *GetRecordDnskeyResponse) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *GetRecordDnskeyResponse) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *GetRecordDnskeyResponse) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreator

`func (o *GetRecordDnskeyResponse) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *GetRecordDnskeyResponse) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *GetRecordDnskeyResponse) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *GetRecordDnskeyResponse) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDnsName

`func (o *GetRecordDnskeyResponse) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *GetRecordDnskeyResponse) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *GetRecordDnskeyResponse) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *GetRecordDnskeyResponse) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetFlags

`func (o *GetRecordDnskeyResponse) GetFlags() string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *GetRecordDnskeyResponse) GetFlagsOk() (*string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *GetRecordDnskeyResponse) SetFlags(v string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *GetRecordDnskeyResponse) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetKeyTag

`func (o *GetRecordDnskeyResponse) GetKeyTag() int64`

GetKeyTag returns the KeyTag field if non-nil, zero value otherwise.

### GetKeyTagOk

`func (o *GetRecordDnskeyResponse) GetKeyTagOk() (*int64, bool)`

GetKeyTagOk returns a tuple with the KeyTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTag

`func (o *GetRecordDnskeyResponse) SetKeyTag(v int64)`

SetKeyTag sets KeyTag field to given value.

### HasKeyTag

`func (o *GetRecordDnskeyResponse) HasKeyTag() bool`

HasKeyTag returns a boolean if a field has been set.

### GetLastQueried

`func (o *GetRecordDnskeyResponse) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *GetRecordDnskeyResponse) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *GetRecordDnskeyResponse) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *GetRecordDnskeyResponse) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetName

`func (o *GetRecordDnskeyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRecordDnskeyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRecordDnskeyResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetRecordDnskeyResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicKey

`func (o *GetRecordDnskeyResponse) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *GetRecordDnskeyResponse) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *GetRecordDnskeyResponse) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *GetRecordDnskeyResponse) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetTtl

`func (o *GetRecordDnskeyResponse) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetRecordDnskeyResponse) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetRecordDnskeyResponse) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetRecordDnskeyResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *GetRecordDnskeyResponse) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *GetRecordDnskeyResponse) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *GetRecordDnskeyResponse) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *GetRecordDnskeyResponse) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *GetRecordDnskeyResponse) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetRecordDnskeyResponse) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetRecordDnskeyResponse) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *GetRecordDnskeyResponse) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *GetRecordDnskeyResponse) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetRecordDnskeyResponse) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetRecordDnskeyResponse) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetRecordDnskeyResponse) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetResult

`func (o *GetRecordDnskeyResponse) GetResult() RecordDnskey`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetRecordDnskeyResponse) GetResultOk() (*RecordDnskey, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetRecordDnskeyResponse) SetResult(v RecordDnskey)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetRecordDnskeyResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


