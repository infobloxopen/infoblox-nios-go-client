# GetTaxiiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**EnableService** | Pointer to **bool** | Indicates whether the Taxii service is running on the given member or not. | [optional] 
**Ipv4addr** | Pointer to **string** | The IPv4 Address of the Grid member. | [optional] [readonly] 
**Ipv6addr** | Pointer to **string** | The IPv6 Address of the Grid member. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the Taxii Member. | [optional] [readonly] 
**TaxiiRpzConfig** | Pointer to [**[]TaxiiTaxiiRpzConfig**](TaxiiTaxiiRpzConfig.md) | Taxii service RPZ configuration list. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**Taxii**](Taxii.md) |  | [optional] 

## Methods

### NewGetTaxiiResponse

`func NewGetTaxiiResponse() *GetTaxiiResponse`

NewGetTaxiiResponse instantiates a new GetTaxiiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTaxiiResponseWithDefaults

`func NewGetTaxiiResponseWithDefaults() *GetTaxiiResponse`

NewGetTaxiiResponseWithDefaults instantiates a new GetTaxiiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetTaxiiResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetTaxiiResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetTaxiiResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetTaxiiResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetEnableService

`func (o *GetTaxiiResponse) GetEnableService() bool`

GetEnableService returns the EnableService field if non-nil, zero value otherwise.

### GetEnableServiceOk

`func (o *GetTaxiiResponse) GetEnableServiceOk() (*bool, bool)`

GetEnableServiceOk returns a tuple with the EnableService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableService

`func (o *GetTaxiiResponse) SetEnableService(v bool)`

SetEnableService sets EnableService field to given value.

### HasEnableService

`func (o *GetTaxiiResponse) HasEnableService() bool`

HasEnableService returns a boolean if a field has been set.

### GetIpv4addr

`func (o *GetTaxiiResponse) GetIpv4addr() string`

GetIpv4addr returns the Ipv4addr field if non-nil, zero value otherwise.

### GetIpv4addrOk

`func (o *GetTaxiiResponse) GetIpv4addrOk() (*string, bool)`

GetIpv4addrOk returns a tuple with the Ipv4addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addr

`func (o *GetTaxiiResponse) SetIpv4addr(v string)`

SetIpv4addr sets Ipv4addr field to given value.

### HasIpv4addr

`func (o *GetTaxiiResponse) HasIpv4addr() bool`

HasIpv4addr returns a boolean if a field has been set.

### GetIpv6addr

`func (o *GetTaxiiResponse) GetIpv6addr() string`

GetIpv6addr returns the Ipv6addr field if non-nil, zero value otherwise.

### GetIpv6addrOk

`func (o *GetTaxiiResponse) GetIpv6addrOk() (*string, bool)`

GetIpv6addrOk returns a tuple with the Ipv6addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6addr

`func (o *GetTaxiiResponse) SetIpv6addr(v string)`

SetIpv6addr sets Ipv6addr field to given value.

### HasIpv6addr

`func (o *GetTaxiiResponse) HasIpv6addr() bool`

HasIpv6addr returns a boolean if a field has been set.

### GetName

`func (o *GetTaxiiResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTaxiiResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTaxiiResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetTaxiiResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTaxiiRpzConfig

`func (o *GetTaxiiResponse) GetTaxiiRpzConfig() []TaxiiTaxiiRpzConfig`

GetTaxiiRpzConfig returns the TaxiiRpzConfig field if non-nil, zero value otherwise.

### GetTaxiiRpzConfigOk

`func (o *GetTaxiiResponse) GetTaxiiRpzConfigOk() (*[]TaxiiTaxiiRpzConfig, bool)`

GetTaxiiRpzConfigOk returns a tuple with the TaxiiRpzConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxiiRpzConfig

`func (o *GetTaxiiResponse) SetTaxiiRpzConfig(v []TaxiiTaxiiRpzConfig)`

SetTaxiiRpzConfig sets TaxiiRpzConfig field to given value.

### HasTaxiiRpzConfig

`func (o *GetTaxiiResponse) HasTaxiiRpzConfig() bool`

HasTaxiiRpzConfig returns a boolean if a field has been set.

### GetUuid

`func (o *GetTaxiiResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetTaxiiResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetTaxiiResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetTaxiiResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetTaxiiResponse) GetResult() Taxii`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetTaxiiResponse) GetResultOk() (*Taxii, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetTaxiiResponse) SetResult(v Taxii)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetTaxiiResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


