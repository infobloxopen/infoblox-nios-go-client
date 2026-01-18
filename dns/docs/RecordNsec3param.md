# RecordNsec3param

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Algorithm** | Pointer to **string** | The hash algorithm that was used. | [optional] [readonly] 
**CloudInfo** | Pointer to [**RecordNsec3paramCloudInfo**](RecordNsec3paramCloudInfo.md) |  | [optional] 
**CreationTime** | Pointer to **int64** | The creation time of the record. | [optional] [readonly] 
**Creator** | Pointer to **string** | Creator of the record. | [optional] [readonly] 
**DnsName** | Pointer to **string** | Name for an NSEC3PARAM record in punycode format. | [optional] [readonly] 
**Flags** | Pointer to **int64** | The set of 8 one-bit flags, of which only one flag, the Opt-Out flag, is defined by RFC 5155. The Opt-Out flag indicates whether the NSEC3 record covers unsigned delegations. | [optional] [readonly] 
**Iterations** | Pointer to **int64** | The number of times the hash function was performed. | [optional] [readonly] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the NSEC3PARAM record in FQDN format. It has to be the same as the zone, where the record resides. | [optional] [readonly] 
**Salt** | Pointer to **string** | A series of case-insensitive hexadecimal digits. It is appended to the original owner name as protection against pre-calculated dictionary attacks. A new salt value is generated when the ZSK rolls over, for which the user can control the period. For a random salt value, the selected length is between one and 15 octets. | [optional] [readonly] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for the record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] [readonly] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] [readonly] 
**View** | Pointer to **string** | The name of the DNS View in which the record resides. Example: \&quot;external\&quot;. | [optional] [readonly] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 

## Methods

### NewRecordNsec3param

`func NewRecordNsec3param() *RecordNsec3param`

NewRecordNsec3param instantiates a new RecordNsec3param object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordNsec3paramWithDefaults

`func NewRecordNsec3paramWithDefaults() *RecordNsec3param`

NewRecordNsec3paramWithDefaults instantiates a new RecordNsec3param object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordNsec3param) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordNsec3param) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordNsec3param) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordNsec3param) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAlgorithm

`func (o *RecordNsec3param) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *RecordNsec3param) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *RecordNsec3param) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *RecordNsec3param) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetCloudInfo

`func (o *RecordNsec3param) GetCloudInfo() RecordNsec3paramCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *RecordNsec3param) GetCloudInfoOk() (*RecordNsec3paramCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *RecordNsec3param) SetCloudInfo(v RecordNsec3paramCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *RecordNsec3param) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetCreationTime

`func (o *RecordNsec3param) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RecordNsec3param) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RecordNsec3param) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *RecordNsec3param) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreator

`func (o *RecordNsec3param) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RecordNsec3param) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RecordNsec3param) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RecordNsec3param) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDnsName

`func (o *RecordNsec3param) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *RecordNsec3param) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *RecordNsec3param) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *RecordNsec3param) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetFlags

`func (o *RecordNsec3param) GetFlags() int64`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *RecordNsec3param) GetFlagsOk() (*int64, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *RecordNsec3param) SetFlags(v int64)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *RecordNsec3param) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetIterations

`func (o *RecordNsec3param) GetIterations() int64`

GetIterations returns the Iterations field if non-nil, zero value otherwise.

### GetIterationsOk

`func (o *RecordNsec3param) GetIterationsOk() (*int64, bool)`

GetIterationsOk returns a tuple with the Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterations

`func (o *RecordNsec3param) SetIterations(v int64)`

SetIterations sets Iterations field to given value.

### HasIterations

`func (o *RecordNsec3param) HasIterations() bool`

HasIterations returns a boolean if a field has been set.

### GetLastQueried

`func (o *RecordNsec3param) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *RecordNsec3param) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *RecordNsec3param) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *RecordNsec3param) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetName

`func (o *RecordNsec3param) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordNsec3param) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordNsec3param) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordNsec3param) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSalt

`func (o *RecordNsec3param) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *RecordNsec3param) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *RecordNsec3param) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *RecordNsec3param) HasSalt() bool`

HasSalt returns a boolean if a field has been set.

### GetTtl

`func (o *RecordNsec3param) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordNsec3param) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordNsec3param) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordNsec3param) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordNsec3param) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordNsec3param) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordNsec3param) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordNsec3param) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordNsec3param) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordNsec3param) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordNsec3param) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordNsec3param) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordNsec3param) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordNsec3param) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordNsec3param) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordNsec3param) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


