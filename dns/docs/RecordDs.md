# RecordDs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Algorithm** | Pointer to **string** | The algorithm of the DNSKEY RR to which this DS RR refers. It uses the same algorithm values and types as the corresponding DNSKEY RR. | [optional] [readonly] 
**CloudInfo** | Pointer to [**RecordDsCloudInfo**](RecordDsCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | The comment for the record. | [optional] [readonly] 
**CreationTime** | Pointer to **int64** | The creation time of the record. | [optional] [readonly] 
**Creator** | Pointer to **string** | Creator of the record. | [optional] [readonly] 
**Digest** | Pointer to **string** | The digest of the DNSKEY resource record that is stored in a DS Record object. | [optional] [readonly] 
**DigestType** | Pointer to **string** | The algorithm used to construct the digest. | [optional] [readonly] 
**DnsName** | Pointer to **string** | The name for the DS record in punycode format. | [optional] [readonly] 
**KeyTag** | Pointer to **int64** | The key tag value that is used to determine which key to use to verify signatures. | [optional] [readonly] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the DNS DS record in FQDN format. | [optional] [readonly] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for the record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] [readonly] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] [readonly] 
**View** | Pointer to **string** | The name of the DNS View in which the record resides. Example: \&quot;external\&quot;. | [optional] [readonly] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 

## Methods

### NewRecordDs

`func NewRecordDs() *RecordDs`

NewRecordDs instantiates a new RecordDs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordDsWithDefaults

`func NewRecordDsWithDefaults() *RecordDs`

NewRecordDsWithDefaults instantiates a new RecordDs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordDs) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordDs) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordDs) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordDs) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAlgorithm

`func (o *RecordDs) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *RecordDs) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *RecordDs) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *RecordDs) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetCloudInfo

`func (o *RecordDs) GetCloudInfo() RecordDsCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *RecordDs) GetCloudInfoOk() (*RecordDsCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *RecordDs) SetCloudInfo(v RecordDsCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *RecordDs) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *RecordDs) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordDs) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordDs) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordDs) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreationTime

`func (o *RecordDs) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RecordDs) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RecordDs) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *RecordDs) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreator

`func (o *RecordDs) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RecordDs) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RecordDs) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RecordDs) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDigest

`func (o *RecordDs) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *RecordDs) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *RecordDs) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *RecordDs) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetDigestType

`func (o *RecordDs) GetDigestType() string`

GetDigestType returns the DigestType field if non-nil, zero value otherwise.

### GetDigestTypeOk

`func (o *RecordDs) GetDigestTypeOk() (*string, bool)`

GetDigestTypeOk returns a tuple with the DigestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestType

`func (o *RecordDs) SetDigestType(v string)`

SetDigestType sets DigestType field to given value.

### HasDigestType

`func (o *RecordDs) HasDigestType() bool`

HasDigestType returns a boolean if a field has been set.

### GetDnsName

`func (o *RecordDs) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *RecordDs) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *RecordDs) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *RecordDs) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetKeyTag

`func (o *RecordDs) GetKeyTag() int64`

GetKeyTag returns the KeyTag field if non-nil, zero value otherwise.

### GetKeyTagOk

`func (o *RecordDs) GetKeyTagOk() (*int64, bool)`

GetKeyTagOk returns a tuple with the KeyTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTag

`func (o *RecordDs) SetKeyTag(v int64)`

SetKeyTag sets KeyTag field to given value.

### HasKeyTag

`func (o *RecordDs) HasKeyTag() bool`

HasKeyTag returns a boolean if a field has been set.

### GetLastQueried

`func (o *RecordDs) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *RecordDs) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *RecordDs) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *RecordDs) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetName

`func (o *RecordDs) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordDs) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordDs) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordDs) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTtl

`func (o *RecordDs) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordDs) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordDs) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordDs) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordDs) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordDs) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordDs) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordDs) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordDs) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordDs) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordDs) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordDs) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordDs) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordDs) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordDs) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordDs) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


