# HsmThaleslunagroupThalesluna

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The HSM Thales Luna device IPv4 Address or FQDN. | [optional] 
**PartitionSerialNumber** | Pointer to **string** | The HSM Thales Luna device partition serial number (PSN). | [optional] 
**Disable** | Pointer to **bool** | Determines whether the HSM Thales Luna device is disabled. | [optional] 
**PartitionId** | Pointer to **string** | Partition ID that is displayed after the appliance has successfully connected to the HSM Thales Luna device. | [optional] [readonly] 
**IsFipsCompliant** | Pointer to **bool** | Determines whether the HSM Thales Luna device is FIPS compliant. | [optional] [readonly] 
**ServerCert** | Pointer to **string** | The token returned by the uploadinit function call in object fileop for a Thales Luna HSM device certificate. | [optional] 
**PartitionCapacity** | Pointer to **int64** | The HSM Thales Luna device partition capacity percentage used. | [optional] [readonly] 
**Status** | Pointer to **string** | The HSM Thales Luna device status. | [optional] [readonly] 

## Methods

### NewHsmThaleslunagroupThalesluna

`func NewHsmThaleslunagroupThalesluna() *HsmThaleslunagroupThalesluna`

NewHsmThaleslunagroupThalesluna instantiates a new HsmThaleslunagroupThalesluna object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHsmThaleslunagroupThaleslunaWithDefaults

`func NewHsmThaleslunagroupThaleslunaWithDefaults() *HsmThaleslunagroupThalesluna`

NewHsmThaleslunagroupThaleslunaWithDefaults instantiates a new HsmThaleslunagroupThalesluna object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HsmThaleslunagroupThalesluna) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HsmThaleslunagroupThalesluna) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HsmThaleslunagroupThalesluna) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HsmThaleslunagroupThalesluna) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartitionSerialNumber

`func (o *HsmThaleslunagroupThalesluna) GetPartitionSerialNumber() string`

GetPartitionSerialNumber returns the PartitionSerialNumber field if non-nil, zero value otherwise.

### GetPartitionSerialNumberOk

`func (o *HsmThaleslunagroupThalesluna) GetPartitionSerialNumberOk() (*string, bool)`

GetPartitionSerialNumberOk returns a tuple with the PartitionSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionSerialNumber

`func (o *HsmThaleslunagroupThalesluna) SetPartitionSerialNumber(v string)`

SetPartitionSerialNumber sets PartitionSerialNumber field to given value.

### HasPartitionSerialNumber

`func (o *HsmThaleslunagroupThalesluna) HasPartitionSerialNumber() bool`

HasPartitionSerialNumber returns a boolean if a field has been set.

### GetDisable

`func (o *HsmThaleslunagroupThalesluna) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *HsmThaleslunagroupThalesluna) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *HsmThaleslunagroupThalesluna) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *HsmThaleslunagroupThalesluna) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetPartitionId

`func (o *HsmThaleslunagroupThalesluna) GetPartitionId() string`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *HsmThaleslunagroupThalesluna) GetPartitionIdOk() (*string, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *HsmThaleslunagroupThalesluna) SetPartitionId(v string)`

SetPartitionId sets PartitionId field to given value.

### HasPartitionId

`func (o *HsmThaleslunagroupThalesluna) HasPartitionId() bool`

HasPartitionId returns a boolean if a field has been set.

### GetIsFipsCompliant

`func (o *HsmThaleslunagroupThalesluna) GetIsFipsCompliant() bool`

GetIsFipsCompliant returns the IsFipsCompliant field if non-nil, zero value otherwise.

### GetIsFipsCompliantOk

`func (o *HsmThaleslunagroupThalesluna) GetIsFipsCompliantOk() (*bool, bool)`

GetIsFipsCompliantOk returns a tuple with the IsFipsCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFipsCompliant

`func (o *HsmThaleslunagroupThalesluna) SetIsFipsCompliant(v bool)`

SetIsFipsCompliant sets IsFipsCompliant field to given value.

### HasIsFipsCompliant

`func (o *HsmThaleslunagroupThalesluna) HasIsFipsCompliant() bool`

HasIsFipsCompliant returns a boolean if a field has been set.

### GetServerCert

`func (o *HsmThaleslunagroupThalesluna) GetServerCert() string`

GetServerCert returns the ServerCert field if non-nil, zero value otherwise.

### GetServerCertOk

`func (o *HsmThaleslunagroupThalesluna) GetServerCertOk() (*string, bool)`

GetServerCertOk returns a tuple with the ServerCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCert

`func (o *HsmThaleslunagroupThalesluna) SetServerCert(v string)`

SetServerCert sets ServerCert field to given value.

### HasServerCert

`func (o *HsmThaleslunagroupThalesluna) HasServerCert() bool`

HasServerCert returns a boolean if a field has been set.

### GetPartitionCapacity

`func (o *HsmThaleslunagroupThalesluna) GetPartitionCapacity() int64`

GetPartitionCapacity returns the PartitionCapacity field if non-nil, zero value otherwise.

### GetPartitionCapacityOk

`func (o *HsmThaleslunagroupThalesluna) GetPartitionCapacityOk() (*int64, bool)`

GetPartitionCapacityOk returns a tuple with the PartitionCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionCapacity

`func (o *HsmThaleslunagroupThalesluna) SetPartitionCapacity(v int64)`

SetPartitionCapacity sets PartitionCapacity field to given value.

### HasPartitionCapacity

`func (o *HsmThaleslunagroupThalesluna) HasPartitionCapacity() bool`

HasPartitionCapacity returns a boolean if a field has been set.

### GetStatus

`func (o *HsmThaleslunagroupThalesluna) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HsmThaleslunagroupThalesluna) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HsmThaleslunagroupThalesluna) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HsmThaleslunagroupThalesluna) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


