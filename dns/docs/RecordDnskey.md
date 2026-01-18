# RecordDnskey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Algorithm** | Pointer to **string** | The public key encryption algorithm of a DNSKEY Record object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The comment for the record. | [optional] [readonly] 
**CreationTime** | Pointer to **int64** | The creation time of the record. | [optional] [readonly] 
**Creator** | Pointer to **string** | The record creator. | [optional] [readonly] 
**DnsName** | Pointer to **string** | Name of a DNSKEY record in punycode format. | [optional] [readonly] 
**Flags** | Pointer to **int64** | The flags field is a 16-bit unsigned integer. Currently, only two bits of this value are used: the least significant bit and bit 7. The other bits are reserved for future use and must be zero. If bit 7 is set to 1, the key is a DNS zone key. Otherwise, the key is not a zone key and cannot be used to verify zone data. The least significant bit indicates \&quot;secure entry point property\&quot;. If it is not zero, the key is a key signing key (KSK type). Otherwise, the key type is ZSK. | [optional] [readonly] 
**KeyTag** | Pointer to **int64** | The key tag identifying the public key of a DNSKEY Record object. | [optional] [readonly] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the DNSKEY record in FQDN format. It has to be the same as the zone, where the record resides. | [optional] [readonly] 
**PublicKey** | Pointer to **string** | The public key. The format of the returned value depends on the key algorithm. | [optional] [readonly] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for the record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] [readonly] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] [readonly] 
**View** | Pointer to **string** | The name of the DNS View in which the record resides. Example: \&quot;external\&quot;. | [optional] [readonly] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 

## Methods

### NewRecordDnskey

`func NewRecordDnskey() *RecordDnskey`

NewRecordDnskey instantiates a new RecordDnskey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordDnskeyWithDefaults

`func NewRecordDnskeyWithDefaults() *RecordDnskey`

NewRecordDnskeyWithDefaults instantiates a new RecordDnskey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordDnskey) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordDnskey) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordDnskey) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordDnskey) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAlgorithm

`func (o *RecordDnskey) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *RecordDnskey) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *RecordDnskey) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *RecordDnskey) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetComment

`func (o *RecordDnskey) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordDnskey) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordDnskey) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordDnskey) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreationTime

`func (o *RecordDnskey) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RecordDnskey) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RecordDnskey) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *RecordDnskey) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreator

`func (o *RecordDnskey) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RecordDnskey) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RecordDnskey) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RecordDnskey) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDnsName

`func (o *RecordDnskey) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *RecordDnskey) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *RecordDnskey) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *RecordDnskey) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetFlags

`func (o *RecordDnskey) GetFlags() int64`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *RecordDnskey) GetFlagsOk() (*int64, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *RecordDnskey) SetFlags(v int64)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *RecordDnskey) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetKeyTag

`func (o *RecordDnskey) GetKeyTag() int64`

GetKeyTag returns the KeyTag field if non-nil, zero value otherwise.

### GetKeyTagOk

`func (o *RecordDnskey) GetKeyTagOk() (*int64, bool)`

GetKeyTagOk returns a tuple with the KeyTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTag

`func (o *RecordDnskey) SetKeyTag(v int64)`

SetKeyTag sets KeyTag field to given value.

### HasKeyTag

`func (o *RecordDnskey) HasKeyTag() bool`

HasKeyTag returns a boolean if a field has been set.

### GetLastQueried

`func (o *RecordDnskey) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *RecordDnskey) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *RecordDnskey) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *RecordDnskey) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetName

`func (o *RecordDnskey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordDnskey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordDnskey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordDnskey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicKey

`func (o *RecordDnskey) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *RecordDnskey) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *RecordDnskey) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *RecordDnskey) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetTtl

`func (o *RecordDnskey) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordDnskey) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordDnskey) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordDnskey) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordDnskey) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordDnskey) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordDnskey) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordDnskey) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordDnskey) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordDnskey) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordDnskey) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordDnskey) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordDnskey) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordDnskey) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordDnskey) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordDnskey) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


