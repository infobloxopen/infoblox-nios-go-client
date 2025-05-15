# ZoneAuthAwsRte53ZoneInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedVpcs** | Pointer to **[]string** | List of AWS VPC strings that are associated with this zone. | [optional] [readonly] 
**CallerReference** | Pointer to **string** | User specified caller reference when zone was created. | [optional] [readonly] 
**DelegationSetId** | Pointer to **string** | ID of delegation set associated with this zone. | [optional] [readonly] 
**HostedZoneId** | Pointer to **string** | AWS route 53 assigned ID for this zone. | [optional] [readonly] 
**NameServers** | Pointer to **[]string** | List of AWS name servers that are authoritative for this domain name. | [optional] [readonly] 
**RecordSetCount** | Pointer to **int64** | Number of resource record sets in the hosted zone. | [optional] [readonly] 
**Type** | Pointer to **string** | Indicates whether private or public zone. | [optional] [readonly] 

## Methods

### NewZoneAuthAwsRte53ZoneInfo

`func NewZoneAuthAwsRte53ZoneInfo() *ZoneAuthAwsRte53ZoneInfo`

NewZoneAuthAwsRte53ZoneInfo instantiates a new ZoneAuthAwsRte53ZoneInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneAuthAwsRte53ZoneInfoWithDefaults

`func NewZoneAuthAwsRte53ZoneInfoWithDefaults() *ZoneAuthAwsRte53ZoneInfo`

NewZoneAuthAwsRte53ZoneInfoWithDefaults instantiates a new ZoneAuthAwsRte53ZoneInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedVpcs

`func (o *ZoneAuthAwsRte53ZoneInfo) GetAssociatedVpcs() []string`

GetAssociatedVpcs returns the AssociatedVpcs field if non-nil, zero value otherwise.

### GetAssociatedVpcsOk

`func (o *ZoneAuthAwsRte53ZoneInfo) GetAssociatedVpcsOk() (*[]string, bool)`

GetAssociatedVpcsOk returns a tuple with the AssociatedVpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedVpcs

`func (o *ZoneAuthAwsRte53ZoneInfo) SetAssociatedVpcs(v []string)`

SetAssociatedVpcs sets AssociatedVpcs field to given value.

### HasAssociatedVpcs

`func (o *ZoneAuthAwsRte53ZoneInfo) HasAssociatedVpcs() bool`

HasAssociatedVpcs returns a boolean if a field has been set.

### GetCallerReference

`func (o *ZoneAuthAwsRte53ZoneInfo) GetCallerReference() string`

GetCallerReference returns the CallerReference field if non-nil, zero value otherwise.

### GetCallerReferenceOk

`func (o *ZoneAuthAwsRte53ZoneInfo) GetCallerReferenceOk() (*string, bool)`

GetCallerReferenceOk returns a tuple with the CallerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerReference

`func (o *ZoneAuthAwsRte53ZoneInfo) SetCallerReference(v string)`

SetCallerReference sets CallerReference field to given value.

### HasCallerReference

`func (o *ZoneAuthAwsRte53ZoneInfo) HasCallerReference() bool`

HasCallerReference returns a boolean if a field has been set.

### GetDelegationSetId

`func (o *ZoneAuthAwsRte53ZoneInfo) GetDelegationSetId() string`

GetDelegationSetId returns the DelegationSetId field if non-nil, zero value otherwise.

### GetDelegationSetIdOk

`func (o *ZoneAuthAwsRte53ZoneInfo) GetDelegationSetIdOk() (*string, bool)`

GetDelegationSetIdOk returns a tuple with the DelegationSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegationSetId

`func (o *ZoneAuthAwsRte53ZoneInfo) SetDelegationSetId(v string)`

SetDelegationSetId sets DelegationSetId field to given value.

### HasDelegationSetId

`func (o *ZoneAuthAwsRte53ZoneInfo) HasDelegationSetId() bool`

HasDelegationSetId returns a boolean if a field has been set.

### GetHostedZoneId

`func (o *ZoneAuthAwsRte53ZoneInfo) GetHostedZoneId() string`

GetHostedZoneId returns the HostedZoneId field if non-nil, zero value otherwise.

### GetHostedZoneIdOk

`func (o *ZoneAuthAwsRte53ZoneInfo) GetHostedZoneIdOk() (*string, bool)`

GetHostedZoneIdOk returns a tuple with the HostedZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZoneId

`func (o *ZoneAuthAwsRte53ZoneInfo) SetHostedZoneId(v string)`

SetHostedZoneId sets HostedZoneId field to given value.

### HasHostedZoneId

`func (o *ZoneAuthAwsRte53ZoneInfo) HasHostedZoneId() bool`

HasHostedZoneId returns a boolean if a field has been set.

### GetNameServers

`func (o *ZoneAuthAwsRte53ZoneInfo) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *ZoneAuthAwsRte53ZoneInfo) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *ZoneAuthAwsRte53ZoneInfo) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.

### HasNameServers

`func (o *ZoneAuthAwsRte53ZoneInfo) HasNameServers() bool`

HasNameServers returns a boolean if a field has been set.

### GetRecordSetCount

`func (o *ZoneAuthAwsRte53ZoneInfo) GetRecordSetCount() int64`

GetRecordSetCount returns the RecordSetCount field if non-nil, zero value otherwise.

### GetRecordSetCountOk

`func (o *ZoneAuthAwsRte53ZoneInfo) GetRecordSetCountOk() (*int64, bool)`

GetRecordSetCountOk returns a tuple with the RecordSetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordSetCount

`func (o *ZoneAuthAwsRte53ZoneInfo) SetRecordSetCount(v int64)`

SetRecordSetCount sets RecordSetCount field to given value.

### HasRecordSetCount

`func (o *ZoneAuthAwsRte53ZoneInfo) HasRecordSetCount() bool`

HasRecordSetCount returns a boolean if a field has been set.

### GetType

`func (o *ZoneAuthAwsRte53ZoneInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ZoneAuthAwsRte53ZoneInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ZoneAuthAwsRte53ZoneInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ZoneAuthAwsRte53ZoneInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


