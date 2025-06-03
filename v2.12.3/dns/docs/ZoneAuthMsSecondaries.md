# ZoneAuthMsSecondaries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The address of the server. | [optional] 
**IsMaster** | Pointer to **bool** | This flag indicates if this server is a synchronization master. | [optional] 
**NsIp** | Pointer to **string** | This address is used when generating the NS record in the zone, which can be different in case of multihomed hosts. | [optional] 
**NsName** | Pointer to **string** | This name is used when generating the NS record in the zone, which can be different in case of multihomed hosts. | [optional] 
**Stealth** | Pointer to **bool** | Set this flag to hide the NS record for the primary name server from DNS queries. | [optional] 
**SharedWithMsParentDelegation** | Pointer to **bool** | This flag represents whether the name server is shared with the parent Microsoft primary zone&#39;s delegation server. | [optional] [readonly] 

## Methods

### NewZoneAuthMsSecondaries

`func NewZoneAuthMsSecondaries() *ZoneAuthMsSecondaries`

NewZoneAuthMsSecondaries instantiates a new ZoneAuthMsSecondaries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneAuthMsSecondariesWithDefaults

`func NewZoneAuthMsSecondariesWithDefaults() *ZoneAuthMsSecondaries`

NewZoneAuthMsSecondariesWithDefaults instantiates a new ZoneAuthMsSecondaries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ZoneAuthMsSecondaries) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ZoneAuthMsSecondaries) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ZoneAuthMsSecondaries) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ZoneAuthMsSecondaries) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIsMaster

`func (o *ZoneAuthMsSecondaries) GetIsMaster() bool`

GetIsMaster returns the IsMaster field if non-nil, zero value otherwise.

### GetIsMasterOk

`func (o *ZoneAuthMsSecondaries) GetIsMasterOk() (*bool, bool)`

GetIsMasterOk returns a tuple with the IsMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMaster

`func (o *ZoneAuthMsSecondaries) SetIsMaster(v bool)`

SetIsMaster sets IsMaster field to given value.

### HasIsMaster

`func (o *ZoneAuthMsSecondaries) HasIsMaster() bool`

HasIsMaster returns a boolean if a field has been set.

### GetNsIp

`func (o *ZoneAuthMsSecondaries) GetNsIp() string`

GetNsIp returns the NsIp field if non-nil, zero value otherwise.

### GetNsIpOk

`func (o *ZoneAuthMsSecondaries) GetNsIpOk() (*string, bool)`

GetNsIpOk returns a tuple with the NsIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsIp

`func (o *ZoneAuthMsSecondaries) SetNsIp(v string)`

SetNsIp sets NsIp field to given value.

### HasNsIp

`func (o *ZoneAuthMsSecondaries) HasNsIp() bool`

HasNsIp returns a boolean if a field has been set.

### GetNsName

`func (o *ZoneAuthMsSecondaries) GetNsName() string`

GetNsName returns the NsName field if non-nil, zero value otherwise.

### GetNsNameOk

`func (o *ZoneAuthMsSecondaries) GetNsNameOk() (*string, bool)`

GetNsNameOk returns a tuple with the NsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsName

`func (o *ZoneAuthMsSecondaries) SetNsName(v string)`

SetNsName sets NsName field to given value.

### HasNsName

`func (o *ZoneAuthMsSecondaries) HasNsName() bool`

HasNsName returns a boolean if a field has been set.

### GetStealth

`func (o *ZoneAuthMsSecondaries) GetStealth() bool`

GetStealth returns the Stealth field if non-nil, zero value otherwise.

### GetStealthOk

`func (o *ZoneAuthMsSecondaries) GetStealthOk() (*bool, bool)`

GetStealthOk returns a tuple with the Stealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStealth

`func (o *ZoneAuthMsSecondaries) SetStealth(v bool)`

SetStealth sets Stealth field to given value.

### HasStealth

`func (o *ZoneAuthMsSecondaries) HasStealth() bool`

HasStealth returns a boolean if a field has been set.

### GetSharedWithMsParentDelegation

`func (o *ZoneAuthMsSecondaries) GetSharedWithMsParentDelegation() bool`

GetSharedWithMsParentDelegation returns the SharedWithMsParentDelegation field if non-nil, zero value otherwise.

### GetSharedWithMsParentDelegationOk

`func (o *ZoneAuthMsSecondaries) GetSharedWithMsParentDelegationOk() (*bool, bool)`

GetSharedWithMsParentDelegationOk returns a tuple with the SharedWithMsParentDelegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedWithMsParentDelegation

`func (o *ZoneAuthMsSecondaries) SetSharedWithMsParentDelegation(v bool)`

SetSharedWithMsParentDelegation sets SharedWithMsParentDelegation field to given value.

### HasSharedWithMsParentDelegation

`func (o *ZoneAuthMsSecondaries) HasSharedWithMsParentDelegation() bool`

HasSharedWithMsParentDelegation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


