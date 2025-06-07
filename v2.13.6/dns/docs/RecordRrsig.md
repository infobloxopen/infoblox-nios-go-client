# RecordRrsig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Algorithm** | Pointer to **string** | The cryptographic algorithm that was used to create the signature. It uses the same algorithm types as the DNSKEY record indicated in the key tag field. | [optional] [readonly] 
**CloudInfo** | Pointer to [**RecordRrsigCloudInfo**](RecordRrsigCloudInfo.md) |  | [optional] 
**CreationTime** | Pointer to **int64** | The creation time of the record. | [optional] [readonly] 
**Creator** | Pointer to **string** | The record creator. | [optional] [readonly] 
**DnsName** | Pointer to **string** | Name for an RRSIG record in punycode format. | [optional] [readonly] 
**DnsSignerName** | Pointer to **string** | The domain name, in punycode format, of the zone that contains the signed RRset. | [optional] [readonly] 
**ExpirationTime** | Pointer to **int64** | The expiry time of an RRSIG record in Epoch seconds format. | [optional] [readonly] 
**InceptionTime** | Pointer to **int64** | The inception time of an RRSIG record in Epoch seconds format. | [optional] [readonly] 
**KeyTag** | Pointer to **int64** | The key tag value of the DNSKEY RR that validates the signature. | [optional] [readonly] 
**Labels** | Pointer to **int64** | The number of labels in the name of the RRset signed with the RRSIG object. | [optional] [readonly] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the RRSIG record in FQDN format. | [optional] [readonly] 
**OriginalTtl** | Pointer to **int64** | The TTL value of the RRset covered by the RRSIG record. | [optional] [readonly] 
**Signature** | Pointer to **string** | The Base64 encoded cryptographic signature that covers the RRSIG RDATA of the RRSIG Record object. | [optional] [readonly] 
**SignerName** | Pointer to **string** | The domain name of the zone in FQDN format that contains the signed RRset. | [optional] [readonly] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for the record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] [readonly] 
**TypeCovered** | Pointer to **string** | The RR type covered by the RRSIG record. | [optional] [readonly] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] [readonly] 
**View** | Pointer to **string** | The name of the DNS View in which the record resides. Example: \&quot;external\&quot;. | [optional] [readonly] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 

## Methods

### NewRecordRrsig

`func NewRecordRrsig() *RecordRrsig`

NewRecordRrsig instantiates a new RecordRrsig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordRrsigWithDefaults

`func NewRecordRrsigWithDefaults() *RecordRrsig`

NewRecordRrsigWithDefaults instantiates a new RecordRrsig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordRrsig) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordRrsig) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordRrsig) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordRrsig) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAlgorithm

`func (o *RecordRrsig) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *RecordRrsig) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *RecordRrsig) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *RecordRrsig) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetCloudInfo

`func (o *RecordRrsig) GetCloudInfo() RecordRrsigCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *RecordRrsig) GetCloudInfoOk() (*RecordRrsigCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *RecordRrsig) SetCloudInfo(v RecordRrsigCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *RecordRrsig) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetCreationTime

`func (o *RecordRrsig) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RecordRrsig) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RecordRrsig) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *RecordRrsig) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreator

`func (o *RecordRrsig) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RecordRrsig) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RecordRrsig) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RecordRrsig) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDnsName

`func (o *RecordRrsig) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *RecordRrsig) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *RecordRrsig) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *RecordRrsig) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetDnsSignerName

`func (o *RecordRrsig) GetDnsSignerName() string`

GetDnsSignerName returns the DnsSignerName field if non-nil, zero value otherwise.

### GetDnsSignerNameOk

`func (o *RecordRrsig) GetDnsSignerNameOk() (*string, bool)`

GetDnsSignerNameOk returns a tuple with the DnsSignerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSignerName

`func (o *RecordRrsig) SetDnsSignerName(v string)`

SetDnsSignerName sets DnsSignerName field to given value.

### HasDnsSignerName

`func (o *RecordRrsig) HasDnsSignerName() bool`

HasDnsSignerName returns a boolean if a field has been set.

### GetExpirationTime

`func (o *RecordRrsig) GetExpirationTime() int64`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *RecordRrsig) GetExpirationTimeOk() (*int64, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *RecordRrsig) SetExpirationTime(v int64)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *RecordRrsig) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetInceptionTime

`func (o *RecordRrsig) GetInceptionTime() int64`

GetInceptionTime returns the InceptionTime field if non-nil, zero value otherwise.

### GetInceptionTimeOk

`func (o *RecordRrsig) GetInceptionTimeOk() (*int64, bool)`

GetInceptionTimeOk returns a tuple with the InceptionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInceptionTime

`func (o *RecordRrsig) SetInceptionTime(v int64)`

SetInceptionTime sets InceptionTime field to given value.

### HasInceptionTime

`func (o *RecordRrsig) HasInceptionTime() bool`

HasInceptionTime returns a boolean if a field has been set.

### GetKeyTag

`func (o *RecordRrsig) GetKeyTag() int64`

GetKeyTag returns the KeyTag field if non-nil, zero value otherwise.

### GetKeyTagOk

`func (o *RecordRrsig) GetKeyTagOk() (*int64, bool)`

GetKeyTagOk returns a tuple with the KeyTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTag

`func (o *RecordRrsig) SetKeyTag(v int64)`

SetKeyTag sets KeyTag field to given value.

### HasKeyTag

`func (o *RecordRrsig) HasKeyTag() bool`

HasKeyTag returns a boolean if a field has been set.

### GetLabels

`func (o *RecordRrsig) GetLabels() int64`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RecordRrsig) GetLabelsOk() (*int64, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RecordRrsig) SetLabels(v int64)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RecordRrsig) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLastQueried

`func (o *RecordRrsig) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *RecordRrsig) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *RecordRrsig) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *RecordRrsig) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetName

`func (o *RecordRrsig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordRrsig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordRrsig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordRrsig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOriginalTtl

`func (o *RecordRrsig) GetOriginalTtl() int64`

GetOriginalTtl returns the OriginalTtl field if non-nil, zero value otherwise.

### GetOriginalTtlOk

`func (o *RecordRrsig) GetOriginalTtlOk() (*int64, bool)`

GetOriginalTtlOk returns a tuple with the OriginalTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTtl

`func (o *RecordRrsig) SetOriginalTtl(v int64)`

SetOriginalTtl sets OriginalTtl field to given value.

### HasOriginalTtl

`func (o *RecordRrsig) HasOriginalTtl() bool`

HasOriginalTtl returns a boolean if a field has been set.

### GetSignature

`func (o *RecordRrsig) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *RecordRrsig) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *RecordRrsig) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *RecordRrsig) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSignerName

`func (o *RecordRrsig) GetSignerName() string`

GetSignerName returns the SignerName field if non-nil, zero value otherwise.

### GetSignerNameOk

`func (o *RecordRrsig) GetSignerNameOk() (*string, bool)`

GetSignerNameOk returns a tuple with the SignerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerName

`func (o *RecordRrsig) SetSignerName(v string)`

SetSignerName sets SignerName field to given value.

### HasSignerName

`func (o *RecordRrsig) HasSignerName() bool`

HasSignerName returns a boolean if a field has been set.

### GetTtl

`func (o *RecordRrsig) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordRrsig) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordRrsig) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordRrsig) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetTypeCovered

`func (o *RecordRrsig) GetTypeCovered() string`

GetTypeCovered returns the TypeCovered field if non-nil, zero value otherwise.

### GetTypeCoveredOk

`func (o *RecordRrsig) GetTypeCoveredOk() (*string, bool)`

GetTypeCoveredOk returns a tuple with the TypeCovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCovered

`func (o *RecordRrsig) SetTypeCovered(v string)`

SetTypeCovered sets TypeCovered field to given value.

### HasTypeCovered

`func (o *RecordRrsig) HasTypeCovered() bool`

HasTypeCovered returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordRrsig) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordRrsig) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordRrsig) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordRrsig) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordRrsig) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordRrsig) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordRrsig) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordRrsig) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordRrsig) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordRrsig) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordRrsig) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordRrsig) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


