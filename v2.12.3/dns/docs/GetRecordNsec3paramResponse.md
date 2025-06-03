# GetRecordNsec3paramResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
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
**Result** | Pointer to [**RecordNsec3param**](RecordNsec3param.md) |  | [optional] 

## Methods

### NewGetRecordNsec3paramResponse

`func NewGetRecordNsec3paramResponse() *GetRecordNsec3paramResponse`

NewGetRecordNsec3paramResponse instantiates a new GetRecordNsec3paramResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRecordNsec3paramResponseWithDefaults

`func NewGetRecordNsec3paramResponseWithDefaults() *GetRecordNsec3paramResponse`

NewGetRecordNsec3paramResponseWithDefaults instantiates a new GetRecordNsec3paramResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetRecordNsec3paramResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetRecordNsec3paramResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetRecordNsec3paramResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetRecordNsec3paramResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAlgorithm

`func (o *GetRecordNsec3paramResponse) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *GetRecordNsec3paramResponse) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *GetRecordNsec3paramResponse) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *GetRecordNsec3paramResponse) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetCloudInfo

`func (o *GetRecordNsec3paramResponse) GetCloudInfo() RecordNsec3paramCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *GetRecordNsec3paramResponse) GetCloudInfoOk() (*RecordNsec3paramCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *GetRecordNsec3paramResponse) SetCloudInfo(v RecordNsec3paramCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *GetRecordNsec3paramResponse) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetCreationTime

`func (o *GetRecordNsec3paramResponse) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *GetRecordNsec3paramResponse) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *GetRecordNsec3paramResponse) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *GetRecordNsec3paramResponse) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreator

`func (o *GetRecordNsec3paramResponse) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *GetRecordNsec3paramResponse) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *GetRecordNsec3paramResponse) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *GetRecordNsec3paramResponse) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDnsName

`func (o *GetRecordNsec3paramResponse) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *GetRecordNsec3paramResponse) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *GetRecordNsec3paramResponse) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *GetRecordNsec3paramResponse) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetFlags

`func (o *GetRecordNsec3paramResponse) GetFlags() int64`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *GetRecordNsec3paramResponse) GetFlagsOk() (*int64, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *GetRecordNsec3paramResponse) SetFlags(v int64)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *GetRecordNsec3paramResponse) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetIterations

`func (o *GetRecordNsec3paramResponse) GetIterations() int64`

GetIterations returns the Iterations field if non-nil, zero value otherwise.

### GetIterationsOk

`func (o *GetRecordNsec3paramResponse) GetIterationsOk() (*int64, bool)`

GetIterationsOk returns a tuple with the Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIterations

`func (o *GetRecordNsec3paramResponse) SetIterations(v int64)`

SetIterations sets Iterations field to given value.

### HasIterations

`func (o *GetRecordNsec3paramResponse) HasIterations() bool`

HasIterations returns a boolean if a field has been set.

### GetLastQueried

`func (o *GetRecordNsec3paramResponse) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *GetRecordNsec3paramResponse) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *GetRecordNsec3paramResponse) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *GetRecordNsec3paramResponse) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetName

`func (o *GetRecordNsec3paramResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRecordNsec3paramResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRecordNsec3paramResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetRecordNsec3paramResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSalt

`func (o *GetRecordNsec3paramResponse) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *GetRecordNsec3paramResponse) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *GetRecordNsec3paramResponse) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *GetRecordNsec3paramResponse) HasSalt() bool`

HasSalt returns a boolean if a field has been set.

### GetTtl

`func (o *GetRecordNsec3paramResponse) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetRecordNsec3paramResponse) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetRecordNsec3paramResponse) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetRecordNsec3paramResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *GetRecordNsec3paramResponse) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *GetRecordNsec3paramResponse) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *GetRecordNsec3paramResponse) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *GetRecordNsec3paramResponse) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *GetRecordNsec3paramResponse) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetRecordNsec3paramResponse) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetRecordNsec3paramResponse) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *GetRecordNsec3paramResponse) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *GetRecordNsec3paramResponse) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetRecordNsec3paramResponse) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetRecordNsec3paramResponse) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetRecordNsec3paramResponse) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetResult

`func (o *GetRecordNsec3paramResponse) GetResult() RecordNsec3param`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetRecordNsec3paramResponse) GetResultOk() (*RecordNsec3param, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetRecordNsec3paramResponse) SetResult(v RecordNsec3param)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetRecordNsec3paramResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


