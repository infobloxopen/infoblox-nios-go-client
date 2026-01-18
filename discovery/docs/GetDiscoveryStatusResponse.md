# GetDiscoveryStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Address** | Pointer to **string** | The IPv4 Address or IPv6 Address of the device. | [optional] [readonly] 
**CliCollectionEnabled** | Pointer to **bool** | Indicates if CLI collection is enabled. | [optional] [readonly] 
**CliCredentialInfo** | Pointer to [**DiscoveryStatusCliCredentialInfo**](DiscoveryStatusCliCredentialInfo.md) |  | [optional] 
**ExistenceInfo** | Pointer to [**DiscoveryStatusExistenceInfo**](DiscoveryStatusExistenceInfo.md) |  | [optional] 
**FingerprintEnabled** | Pointer to **bool** | Indicates if DHCP finterprinting is enabled. | [optional] [readonly] 
**FingerprintInfo** | Pointer to [**DiscoveryStatusFingerprintInfo**](DiscoveryStatusFingerprintInfo.md) |  | [optional] 
**FirstSeen** | Pointer to **int64** | The timestamp when the device was first discovered. | [optional] [readonly] 
**LastAction** | Pointer to **string** | The timestamp of the last detected interface property change. | [optional] [readonly] 
**LastSeen** | Pointer to **int64** | The timestamp when the device was last discovered. | [optional] [readonly] 
**LastTimestamp** | Pointer to **int64** | The timestamp of the last executed action for the device. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the device. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The name of the network view in which this device resides. | [optional] [readonly] 
**ReachableInfo** | Pointer to [**DiscoveryStatusReachableInfo**](DiscoveryStatusReachableInfo.md) |  | [optional] 
**SdnCollectionEnabled** | Pointer to **bool** | Indicate whether SDN collection enabled for the device. | [optional] [readonly] 
**SdnCollectionInfo** | Pointer to [**DiscoveryStatusSdnCollectionInfo**](DiscoveryStatusSdnCollectionInfo.md) |  | [optional] 
**SnmpCollectionEnabled** | Pointer to **bool** | Indicates if SNMP collection is enabled. | [optional] [readonly] 
**SnmpCollectionInfo** | Pointer to [**DiscoveryStatusSnmpCollectionInfo**](DiscoveryStatusSnmpCollectionInfo.md) |  | [optional] 
**SnmpCredentialInfo** | Pointer to [**DiscoveryStatusSnmpCredentialInfo**](DiscoveryStatusSnmpCredentialInfo.md) |  | [optional] 
**Status** | Pointer to **string** | The overall status of the device. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of device. | [optional] [readonly] 
**Result** | Pointer to [**DiscoveryStatus**](DiscoveryStatus.md) |  | [optional] 

## Methods

### NewGetDiscoveryStatusResponse

`func NewGetDiscoveryStatusResponse() *GetDiscoveryStatusResponse`

NewGetDiscoveryStatusResponse instantiates a new GetDiscoveryStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDiscoveryStatusResponseWithDefaults

`func NewGetDiscoveryStatusResponseWithDefaults() *GetDiscoveryStatusResponse`

NewGetDiscoveryStatusResponseWithDefaults instantiates a new GetDiscoveryStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDiscoveryStatusResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDiscoveryStatusResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDiscoveryStatusResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDiscoveryStatusResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *GetDiscoveryStatusResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetDiscoveryStatusResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetDiscoveryStatusResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetDiscoveryStatusResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCliCollectionEnabled

`func (o *GetDiscoveryStatusResponse) GetCliCollectionEnabled() bool`

GetCliCollectionEnabled returns the CliCollectionEnabled field if non-nil, zero value otherwise.

### GetCliCollectionEnabledOk

`func (o *GetDiscoveryStatusResponse) GetCliCollectionEnabledOk() (*bool, bool)`

GetCliCollectionEnabledOk returns a tuple with the CliCollectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliCollectionEnabled

`func (o *GetDiscoveryStatusResponse) SetCliCollectionEnabled(v bool)`

SetCliCollectionEnabled sets CliCollectionEnabled field to given value.

### HasCliCollectionEnabled

`func (o *GetDiscoveryStatusResponse) HasCliCollectionEnabled() bool`

HasCliCollectionEnabled returns a boolean if a field has been set.

### GetCliCredentialInfo

`func (o *GetDiscoveryStatusResponse) GetCliCredentialInfo() DiscoveryStatusCliCredentialInfo`

GetCliCredentialInfo returns the CliCredentialInfo field if non-nil, zero value otherwise.

### GetCliCredentialInfoOk

`func (o *GetDiscoveryStatusResponse) GetCliCredentialInfoOk() (*DiscoveryStatusCliCredentialInfo, bool)`

GetCliCredentialInfoOk returns a tuple with the CliCredentialInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliCredentialInfo

`func (o *GetDiscoveryStatusResponse) SetCliCredentialInfo(v DiscoveryStatusCliCredentialInfo)`

SetCliCredentialInfo sets CliCredentialInfo field to given value.

### HasCliCredentialInfo

`func (o *GetDiscoveryStatusResponse) HasCliCredentialInfo() bool`

HasCliCredentialInfo returns a boolean if a field has been set.

### GetExistenceInfo

`func (o *GetDiscoveryStatusResponse) GetExistenceInfo() DiscoveryStatusExistenceInfo`

GetExistenceInfo returns the ExistenceInfo field if non-nil, zero value otherwise.

### GetExistenceInfoOk

`func (o *GetDiscoveryStatusResponse) GetExistenceInfoOk() (*DiscoveryStatusExistenceInfo, bool)`

GetExistenceInfoOk returns a tuple with the ExistenceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistenceInfo

`func (o *GetDiscoveryStatusResponse) SetExistenceInfo(v DiscoveryStatusExistenceInfo)`

SetExistenceInfo sets ExistenceInfo field to given value.

### HasExistenceInfo

`func (o *GetDiscoveryStatusResponse) HasExistenceInfo() bool`

HasExistenceInfo returns a boolean if a field has been set.

### GetFingerprintEnabled

`func (o *GetDiscoveryStatusResponse) GetFingerprintEnabled() bool`

GetFingerprintEnabled returns the FingerprintEnabled field if non-nil, zero value otherwise.

### GetFingerprintEnabledOk

`func (o *GetDiscoveryStatusResponse) GetFingerprintEnabledOk() (*bool, bool)`

GetFingerprintEnabledOk returns a tuple with the FingerprintEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintEnabled

`func (o *GetDiscoveryStatusResponse) SetFingerprintEnabled(v bool)`

SetFingerprintEnabled sets FingerprintEnabled field to given value.

### HasFingerprintEnabled

`func (o *GetDiscoveryStatusResponse) HasFingerprintEnabled() bool`

HasFingerprintEnabled returns a boolean if a field has been set.

### GetFingerprintInfo

`func (o *GetDiscoveryStatusResponse) GetFingerprintInfo() DiscoveryStatusFingerprintInfo`

GetFingerprintInfo returns the FingerprintInfo field if non-nil, zero value otherwise.

### GetFingerprintInfoOk

`func (o *GetDiscoveryStatusResponse) GetFingerprintInfoOk() (*DiscoveryStatusFingerprintInfo, bool)`

GetFingerprintInfoOk returns a tuple with the FingerprintInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintInfo

`func (o *GetDiscoveryStatusResponse) SetFingerprintInfo(v DiscoveryStatusFingerprintInfo)`

SetFingerprintInfo sets FingerprintInfo field to given value.

### HasFingerprintInfo

`func (o *GetDiscoveryStatusResponse) HasFingerprintInfo() bool`

HasFingerprintInfo returns a boolean if a field has been set.

### GetFirstSeen

`func (o *GetDiscoveryStatusResponse) GetFirstSeen() int64`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *GetDiscoveryStatusResponse) GetFirstSeenOk() (*int64, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *GetDiscoveryStatusResponse) SetFirstSeen(v int64)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *GetDiscoveryStatusResponse) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetLastAction

`func (o *GetDiscoveryStatusResponse) GetLastAction() string`

GetLastAction returns the LastAction field if non-nil, zero value otherwise.

### GetLastActionOk

`func (o *GetDiscoveryStatusResponse) GetLastActionOk() (*string, bool)`

GetLastActionOk returns a tuple with the LastAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAction

`func (o *GetDiscoveryStatusResponse) SetLastAction(v string)`

SetLastAction sets LastAction field to given value.

### HasLastAction

`func (o *GetDiscoveryStatusResponse) HasLastAction() bool`

HasLastAction returns a boolean if a field has been set.

### GetLastSeen

`func (o *GetDiscoveryStatusResponse) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *GetDiscoveryStatusResponse) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *GetDiscoveryStatusResponse) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *GetDiscoveryStatusResponse) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLastTimestamp

`func (o *GetDiscoveryStatusResponse) GetLastTimestamp() int64`

GetLastTimestamp returns the LastTimestamp field if non-nil, zero value otherwise.

### GetLastTimestampOk

`func (o *GetDiscoveryStatusResponse) GetLastTimestampOk() (*int64, bool)`

GetLastTimestampOk returns a tuple with the LastTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTimestamp

`func (o *GetDiscoveryStatusResponse) SetLastTimestamp(v int64)`

SetLastTimestamp sets LastTimestamp field to given value.

### HasLastTimestamp

`func (o *GetDiscoveryStatusResponse) HasLastTimestamp() bool`

HasLastTimestamp returns a boolean if a field has been set.

### GetName

`func (o *GetDiscoveryStatusResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDiscoveryStatusResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDiscoveryStatusResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDiscoveryStatusResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetDiscoveryStatusResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetDiscoveryStatusResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetDiscoveryStatusResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetDiscoveryStatusResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetReachableInfo

`func (o *GetDiscoveryStatusResponse) GetReachableInfo() DiscoveryStatusReachableInfo`

GetReachableInfo returns the ReachableInfo field if non-nil, zero value otherwise.

### GetReachableInfoOk

`func (o *GetDiscoveryStatusResponse) GetReachableInfoOk() (*DiscoveryStatusReachableInfo, bool)`

GetReachableInfoOk returns a tuple with the ReachableInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachableInfo

`func (o *GetDiscoveryStatusResponse) SetReachableInfo(v DiscoveryStatusReachableInfo)`

SetReachableInfo sets ReachableInfo field to given value.

### HasReachableInfo

`func (o *GetDiscoveryStatusResponse) HasReachableInfo() bool`

HasReachableInfo returns a boolean if a field has been set.

### GetSdnCollectionEnabled

`func (o *GetDiscoveryStatusResponse) GetSdnCollectionEnabled() bool`

GetSdnCollectionEnabled returns the SdnCollectionEnabled field if non-nil, zero value otherwise.

### GetSdnCollectionEnabledOk

`func (o *GetDiscoveryStatusResponse) GetSdnCollectionEnabledOk() (*bool, bool)`

GetSdnCollectionEnabledOk returns a tuple with the SdnCollectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdnCollectionEnabled

`func (o *GetDiscoveryStatusResponse) SetSdnCollectionEnabled(v bool)`

SetSdnCollectionEnabled sets SdnCollectionEnabled field to given value.

### HasSdnCollectionEnabled

`func (o *GetDiscoveryStatusResponse) HasSdnCollectionEnabled() bool`

HasSdnCollectionEnabled returns a boolean if a field has been set.

### GetSdnCollectionInfo

`func (o *GetDiscoveryStatusResponse) GetSdnCollectionInfo() DiscoveryStatusSdnCollectionInfo`

GetSdnCollectionInfo returns the SdnCollectionInfo field if non-nil, zero value otherwise.

### GetSdnCollectionInfoOk

`func (o *GetDiscoveryStatusResponse) GetSdnCollectionInfoOk() (*DiscoveryStatusSdnCollectionInfo, bool)`

GetSdnCollectionInfoOk returns a tuple with the SdnCollectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdnCollectionInfo

`func (o *GetDiscoveryStatusResponse) SetSdnCollectionInfo(v DiscoveryStatusSdnCollectionInfo)`

SetSdnCollectionInfo sets SdnCollectionInfo field to given value.

### HasSdnCollectionInfo

`func (o *GetDiscoveryStatusResponse) HasSdnCollectionInfo() bool`

HasSdnCollectionInfo returns a boolean if a field has been set.

### GetSnmpCollectionEnabled

`func (o *GetDiscoveryStatusResponse) GetSnmpCollectionEnabled() bool`

GetSnmpCollectionEnabled returns the SnmpCollectionEnabled field if non-nil, zero value otherwise.

### GetSnmpCollectionEnabledOk

`func (o *GetDiscoveryStatusResponse) GetSnmpCollectionEnabledOk() (*bool, bool)`

GetSnmpCollectionEnabledOk returns a tuple with the SnmpCollectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpCollectionEnabled

`func (o *GetDiscoveryStatusResponse) SetSnmpCollectionEnabled(v bool)`

SetSnmpCollectionEnabled sets SnmpCollectionEnabled field to given value.

### HasSnmpCollectionEnabled

`func (o *GetDiscoveryStatusResponse) HasSnmpCollectionEnabled() bool`

HasSnmpCollectionEnabled returns a boolean if a field has been set.

### GetSnmpCollectionInfo

`func (o *GetDiscoveryStatusResponse) GetSnmpCollectionInfo() DiscoveryStatusSnmpCollectionInfo`

GetSnmpCollectionInfo returns the SnmpCollectionInfo field if non-nil, zero value otherwise.

### GetSnmpCollectionInfoOk

`func (o *GetDiscoveryStatusResponse) GetSnmpCollectionInfoOk() (*DiscoveryStatusSnmpCollectionInfo, bool)`

GetSnmpCollectionInfoOk returns a tuple with the SnmpCollectionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpCollectionInfo

`func (o *GetDiscoveryStatusResponse) SetSnmpCollectionInfo(v DiscoveryStatusSnmpCollectionInfo)`

SetSnmpCollectionInfo sets SnmpCollectionInfo field to given value.

### HasSnmpCollectionInfo

`func (o *GetDiscoveryStatusResponse) HasSnmpCollectionInfo() bool`

HasSnmpCollectionInfo returns a boolean if a field has been set.

### GetSnmpCredentialInfo

`func (o *GetDiscoveryStatusResponse) GetSnmpCredentialInfo() DiscoveryStatusSnmpCredentialInfo`

GetSnmpCredentialInfo returns the SnmpCredentialInfo field if non-nil, zero value otherwise.

### GetSnmpCredentialInfoOk

`func (o *GetDiscoveryStatusResponse) GetSnmpCredentialInfoOk() (*DiscoveryStatusSnmpCredentialInfo, bool)`

GetSnmpCredentialInfoOk returns a tuple with the SnmpCredentialInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpCredentialInfo

`func (o *GetDiscoveryStatusResponse) SetSnmpCredentialInfo(v DiscoveryStatusSnmpCredentialInfo)`

SetSnmpCredentialInfo sets SnmpCredentialInfo field to given value.

### HasSnmpCredentialInfo

`func (o *GetDiscoveryStatusResponse) HasSnmpCredentialInfo() bool`

HasSnmpCredentialInfo returns a boolean if a field has been set.

### GetStatus

`func (o *GetDiscoveryStatusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDiscoveryStatusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDiscoveryStatusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetDiscoveryStatusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *GetDiscoveryStatusResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDiscoveryStatusResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDiscoveryStatusResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetDiscoveryStatusResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetResult

`func (o *GetDiscoveryStatusResponse) GetResult() DiscoveryStatus`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDiscoveryStatusResponse) GetResultOk() (*DiscoveryStatus, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDiscoveryStatusResponse) SetResult(v DiscoveryStatus)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDiscoveryStatusResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


