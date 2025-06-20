# GetDhcpfailoverResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AssociationType** | Pointer to **string** | The value indicating whether the failover assoctaion is Microsoft or Grid based. This is a read-only attribute. | [optional] [readonly] 
**Comment** | Pointer to **string** | A descriptive comment about a DHCP failover object. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**FailoverPort** | Pointer to **int64** | Determines the TCP port on which the server should listen for connections from its failover peer. Valid values are between 1 and 63999. | [optional] 
**LoadBalanceSplit** | Pointer to **int64** | A load balancing split value of a DHCP failover object. Specify the value of the maximum load balancing delay in a 8-bit integer format (range from 0 to 256). | [optional] 
**MaxClientLeadTime** | Pointer to **int64** | The maximum client lead time value of a DHCP failover object. Specify the value of the maximum client lead time in a 32-bit integer format (range from 0 to 4294967295) that represents the duration in seconds. Valid values are between 1 and 4294967295. | [optional] 
**MaxLoadBalanceDelay** | Pointer to **int64** | The maximum load balancing delay value of a DHCP failover object. Specify the value of the maximum load balancing delay in a 32-bit integer format (range from 0 to 4294967295) that represents the duration in seconds. Valid values are between 1 and 4294967295. | [optional] 
**MaxResponseDelay** | Pointer to **int64** | The maximum response delay value of a DHCP failover object. Specify the value of the maximum response delay in a 32-bit integer format (range from 0 to 4294967295) that represents the duration in seconds. Valid values are between 1 and 4294967295. | [optional] 
**MaxUnackedUpdates** | Pointer to **int64** | The maximum number of unacked updates value of a DHCP failover object. Specify the value of the maximum number of unacked updates in a 32-bit integer format (range from 0 to 4294967295) that represents the number of messages. Valid values are between 1 and 4294967295. | [optional] 
**MsAssociationMode** | Pointer to **string** | The value that indicates whether the failover association is read-write or read-only. This is a read-only attribute. | [optional] [readonly] 
**MsEnableAuthentication** | Pointer to **bool** | Determines if the authentication for the failover association is enabled or not. | [optional] 
**MsEnableSwitchoverInterval** | Pointer to **bool** | Determines if the switchover interval is enabled or not. | [optional] 
**MsFailoverMode** | Pointer to **string** | The mode for the failover association. | [optional] 
**MsFailoverPartner** | Pointer to **string** | Failover partner defined in the association with the Microsoft Server. | [optional] 
**MsHotstandbyPartnerRole** | Pointer to **string** | The partner role in the case of HotStandby. | [optional] 
**MsIsConflict** | Pointer to **bool** | Determines if the matching Microsfot failover association (if any) is in synchronization (False) or not (True). If there is no matching failover association the returned values is False. This is a read-only attribute. | [optional] [readonly] 
**MsPreviousState** | Pointer to **string** | The previous failover association state. This is a read-only attribute. | [optional] [readonly] 
**MsServer** | Pointer to **string** | The primary Microsoft Server. | [optional] 
**MsSharedSecret** | Pointer to **string** | The failover association authentication. This is a write-only attribute. | [optional] 
**MsState** | Pointer to **string** | The failover association state. This is a read-only attribute. | [optional] [readonly] 
**MsSwitchoverInterval** | Pointer to **int64** | The time (in seconds) that DHCPv4 server will wait before transitioning the server from the COMMUNICATION-INT state to PARTNER-DOWN state. | [optional] 
**Name** | Pointer to **string** | The name of a DHCP failover object. | [optional] 
**Primary** | Pointer to **string** | The primary server of a DHCP failover object. | [optional] 
**PrimaryServerType** | Pointer to **string** | The type of the primary server of DHCP Failover association object. | [optional] 
**PrimaryState** | Pointer to **string** | The primary server status of a DHCP failover object. | [optional] [readonly] 
**RecycleLeases** | Pointer to **bool** | Determines if the leases are kept in recycle bin until one week after expiration or not. | [optional] 
**Secondary** | Pointer to **string** | The secondary server of a DHCP failover object. | [optional] 
**SecondaryServerType** | Pointer to **string** | The type of the secondary server of DHCP Failover association object. | [optional] 
**SecondaryState** | Pointer to **string** | The secondary server status of a DHCP failover object. | [optional] [readonly] 
**UseFailoverPort** | Pointer to **bool** | Use flag for: failover_port | [optional] 
**UseMsSwitchoverInterval** | Pointer to **bool** | Use flag for: ms_switchover_interval | [optional] 
**UseRecycleLeases** | Pointer to **bool** | Use flag for: recycle_leases | [optional] 
**Result** | Pointer to [**Dhcpfailover**](Dhcpfailover.md) |  | [optional] 

## Methods

### NewGetDhcpfailoverResponse

`func NewGetDhcpfailoverResponse() *GetDhcpfailoverResponse`

NewGetDhcpfailoverResponse instantiates a new GetDhcpfailoverResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDhcpfailoverResponseWithDefaults

`func NewGetDhcpfailoverResponseWithDefaults() *GetDhcpfailoverResponse`

NewGetDhcpfailoverResponseWithDefaults instantiates a new GetDhcpfailoverResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDhcpfailoverResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDhcpfailoverResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDhcpfailoverResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDhcpfailoverResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAssociationType

`func (o *GetDhcpfailoverResponse) GetAssociationType() string`

GetAssociationType returns the AssociationType field if non-nil, zero value otherwise.

### GetAssociationTypeOk

`func (o *GetDhcpfailoverResponse) GetAssociationTypeOk() (*string, bool)`

GetAssociationTypeOk returns a tuple with the AssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationType

`func (o *GetDhcpfailoverResponse) SetAssociationType(v string)`

SetAssociationType sets AssociationType field to given value.

### HasAssociationType

`func (o *GetDhcpfailoverResponse) HasAssociationType() bool`

HasAssociationType returns a boolean if a field has been set.

### GetComment

`func (o *GetDhcpfailoverResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetDhcpfailoverResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetDhcpfailoverResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetDhcpfailoverResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetDhcpfailoverResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetDhcpfailoverResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetDhcpfailoverResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetDhcpfailoverResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetFailoverPort

`func (o *GetDhcpfailoverResponse) GetFailoverPort() int64`

GetFailoverPort returns the FailoverPort field if non-nil, zero value otherwise.

### GetFailoverPortOk

`func (o *GetDhcpfailoverResponse) GetFailoverPortOk() (*int64, bool)`

GetFailoverPortOk returns a tuple with the FailoverPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverPort

`func (o *GetDhcpfailoverResponse) SetFailoverPort(v int64)`

SetFailoverPort sets FailoverPort field to given value.

### HasFailoverPort

`func (o *GetDhcpfailoverResponse) HasFailoverPort() bool`

HasFailoverPort returns a boolean if a field has been set.

### GetLoadBalanceSplit

`func (o *GetDhcpfailoverResponse) GetLoadBalanceSplit() int64`

GetLoadBalanceSplit returns the LoadBalanceSplit field if non-nil, zero value otherwise.

### GetLoadBalanceSplitOk

`func (o *GetDhcpfailoverResponse) GetLoadBalanceSplitOk() (*int64, bool)`

GetLoadBalanceSplitOk returns a tuple with the LoadBalanceSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceSplit

`func (o *GetDhcpfailoverResponse) SetLoadBalanceSplit(v int64)`

SetLoadBalanceSplit sets LoadBalanceSplit field to given value.

### HasLoadBalanceSplit

`func (o *GetDhcpfailoverResponse) HasLoadBalanceSplit() bool`

HasLoadBalanceSplit returns a boolean if a field has been set.

### GetMaxClientLeadTime

`func (o *GetDhcpfailoverResponse) GetMaxClientLeadTime() int64`

GetMaxClientLeadTime returns the MaxClientLeadTime field if non-nil, zero value otherwise.

### GetMaxClientLeadTimeOk

`func (o *GetDhcpfailoverResponse) GetMaxClientLeadTimeOk() (*int64, bool)`

GetMaxClientLeadTimeOk returns a tuple with the MaxClientLeadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClientLeadTime

`func (o *GetDhcpfailoverResponse) SetMaxClientLeadTime(v int64)`

SetMaxClientLeadTime sets MaxClientLeadTime field to given value.

### HasMaxClientLeadTime

`func (o *GetDhcpfailoverResponse) HasMaxClientLeadTime() bool`

HasMaxClientLeadTime returns a boolean if a field has been set.

### GetMaxLoadBalanceDelay

`func (o *GetDhcpfailoverResponse) GetMaxLoadBalanceDelay() int64`

GetMaxLoadBalanceDelay returns the MaxLoadBalanceDelay field if non-nil, zero value otherwise.

### GetMaxLoadBalanceDelayOk

`func (o *GetDhcpfailoverResponse) GetMaxLoadBalanceDelayOk() (*int64, bool)`

GetMaxLoadBalanceDelayOk returns a tuple with the MaxLoadBalanceDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLoadBalanceDelay

`func (o *GetDhcpfailoverResponse) SetMaxLoadBalanceDelay(v int64)`

SetMaxLoadBalanceDelay sets MaxLoadBalanceDelay field to given value.

### HasMaxLoadBalanceDelay

`func (o *GetDhcpfailoverResponse) HasMaxLoadBalanceDelay() bool`

HasMaxLoadBalanceDelay returns a boolean if a field has been set.

### GetMaxResponseDelay

`func (o *GetDhcpfailoverResponse) GetMaxResponseDelay() int64`

GetMaxResponseDelay returns the MaxResponseDelay field if non-nil, zero value otherwise.

### GetMaxResponseDelayOk

`func (o *GetDhcpfailoverResponse) GetMaxResponseDelayOk() (*int64, bool)`

GetMaxResponseDelayOk returns a tuple with the MaxResponseDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResponseDelay

`func (o *GetDhcpfailoverResponse) SetMaxResponseDelay(v int64)`

SetMaxResponseDelay sets MaxResponseDelay field to given value.

### HasMaxResponseDelay

`func (o *GetDhcpfailoverResponse) HasMaxResponseDelay() bool`

HasMaxResponseDelay returns a boolean if a field has been set.

### GetMaxUnackedUpdates

`func (o *GetDhcpfailoverResponse) GetMaxUnackedUpdates() int64`

GetMaxUnackedUpdates returns the MaxUnackedUpdates field if non-nil, zero value otherwise.

### GetMaxUnackedUpdatesOk

`func (o *GetDhcpfailoverResponse) GetMaxUnackedUpdatesOk() (*int64, bool)`

GetMaxUnackedUpdatesOk returns a tuple with the MaxUnackedUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUnackedUpdates

`func (o *GetDhcpfailoverResponse) SetMaxUnackedUpdates(v int64)`

SetMaxUnackedUpdates sets MaxUnackedUpdates field to given value.

### HasMaxUnackedUpdates

`func (o *GetDhcpfailoverResponse) HasMaxUnackedUpdates() bool`

HasMaxUnackedUpdates returns a boolean if a field has been set.

### GetMsAssociationMode

`func (o *GetDhcpfailoverResponse) GetMsAssociationMode() string`

GetMsAssociationMode returns the MsAssociationMode field if non-nil, zero value otherwise.

### GetMsAssociationModeOk

`func (o *GetDhcpfailoverResponse) GetMsAssociationModeOk() (*string, bool)`

GetMsAssociationModeOk returns a tuple with the MsAssociationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAssociationMode

`func (o *GetDhcpfailoverResponse) SetMsAssociationMode(v string)`

SetMsAssociationMode sets MsAssociationMode field to given value.

### HasMsAssociationMode

`func (o *GetDhcpfailoverResponse) HasMsAssociationMode() bool`

HasMsAssociationMode returns a boolean if a field has been set.

### GetMsEnableAuthentication

`func (o *GetDhcpfailoverResponse) GetMsEnableAuthentication() bool`

GetMsEnableAuthentication returns the MsEnableAuthentication field if non-nil, zero value otherwise.

### GetMsEnableAuthenticationOk

`func (o *GetDhcpfailoverResponse) GetMsEnableAuthenticationOk() (*bool, bool)`

GetMsEnableAuthenticationOk returns a tuple with the MsEnableAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsEnableAuthentication

`func (o *GetDhcpfailoverResponse) SetMsEnableAuthentication(v bool)`

SetMsEnableAuthentication sets MsEnableAuthentication field to given value.

### HasMsEnableAuthentication

`func (o *GetDhcpfailoverResponse) HasMsEnableAuthentication() bool`

HasMsEnableAuthentication returns a boolean if a field has been set.

### GetMsEnableSwitchoverInterval

`func (o *GetDhcpfailoverResponse) GetMsEnableSwitchoverInterval() bool`

GetMsEnableSwitchoverInterval returns the MsEnableSwitchoverInterval field if non-nil, zero value otherwise.

### GetMsEnableSwitchoverIntervalOk

`func (o *GetDhcpfailoverResponse) GetMsEnableSwitchoverIntervalOk() (*bool, bool)`

GetMsEnableSwitchoverIntervalOk returns a tuple with the MsEnableSwitchoverInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsEnableSwitchoverInterval

`func (o *GetDhcpfailoverResponse) SetMsEnableSwitchoverInterval(v bool)`

SetMsEnableSwitchoverInterval sets MsEnableSwitchoverInterval field to given value.

### HasMsEnableSwitchoverInterval

`func (o *GetDhcpfailoverResponse) HasMsEnableSwitchoverInterval() bool`

HasMsEnableSwitchoverInterval returns a boolean if a field has been set.

### GetMsFailoverMode

`func (o *GetDhcpfailoverResponse) GetMsFailoverMode() string`

GetMsFailoverMode returns the MsFailoverMode field if non-nil, zero value otherwise.

### GetMsFailoverModeOk

`func (o *GetDhcpfailoverResponse) GetMsFailoverModeOk() (*string, bool)`

GetMsFailoverModeOk returns a tuple with the MsFailoverMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsFailoverMode

`func (o *GetDhcpfailoverResponse) SetMsFailoverMode(v string)`

SetMsFailoverMode sets MsFailoverMode field to given value.

### HasMsFailoverMode

`func (o *GetDhcpfailoverResponse) HasMsFailoverMode() bool`

HasMsFailoverMode returns a boolean if a field has been set.

### GetMsFailoverPartner

`func (o *GetDhcpfailoverResponse) GetMsFailoverPartner() string`

GetMsFailoverPartner returns the MsFailoverPartner field if non-nil, zero value otherwise.

### GetMsFailoverPartnerOk

`func (o *GetDhcpfailoverResponse) GetMsFailoverPartnerOk() (*string, bool)`

GetMsFailoverPartnerOk returns a tuple with the MsFailoverPartner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsFailoverPartner

`func (o *GetDhcpfailoverResponse) SetMsFailoverPartner(v string)`

SetMsFailoverPartner sets MsFailoverPartner field to given value.

### HasMsFailoverPartner

`func (o *GetDhcpfailoverResponse) HasMsFailoverPartner() bool`

HasMsFailoverPartner returns a boolean if a field has been set.

### GetMsHotstandbyPartnerRole

`func (o *GetDhcpfailoverResponse) GetMsHotstandbyPartnerRole() string`

GetMsHotstandbyPartnerRole returns the MsHotstandbyPartnerRole field if non-nil, zero value otherwise.

### GetMsHotstandbyPartnerRoleOk

`func (o *GetDhcpfailoverResponse) GetMsHotstandbyPartnerRoleOk() (*string, bool)`

GetMsHotstandbyPartnerRoleOk returns a tuple with the MsHotstandbyPartnerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsHotstandbyPartnerRole

`func (o *GetDhcpfailoverResponse) SetMsHotstandbyPartnerRole(v string)`

SetMsHotstandbyPartnerRole sets MsHotstandbyPartnerRole field to given value.

### HasMsHotstandbyPartnerRole

`func (o *GetDhcpfailoverResponse) HasMsHotstandbyPartnerRole() bool`

HasMsHotstandbyPartnerRole returns a boolean if a field has been set.

### GetMsIsConflict

`func (o *GetDhcpfailoverResponse) GetMsIsConflict() bool`

GetMsIsConflict returns the MsIsConflict field if non-nil, zero value otherwise.

### GetMsIsConflictOk

`func (o *GetDhcpfailoverResponse) GetMsIsConflictOk() (*bool, bool)`

GetMsIsConflictOk returns a tuple with the MsIsConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsIsConflict

`func (o *GetDhcpfailoverResponse) SetMsIsConflict(v bool)`

SetMsIsConflict sets MsIsConflict field to given value.

### HasMsIsConflict

`func (o *GetDhcpfailoverResponse) HasMsIsConflict() bool`

HasMsIsConflict returns a boolean if a field has been set.

### GetMsPreviousState

`func (o *GetDhcpfailoverResponse) GetMsPreviousState() string`

GetMsPreviousState returns the MsPreviousState field if non-nil, zero value otherwise.

### GetMsPreviousStateOk

`func (o *GetDhcpfailoverResponse) GetMsPreviousStateOk() (*string, bool)`

GetMsPreviousStateOk returns a tuple with the MsPreviousState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsPreviousState

`func (o *GetDhcpfailoverResponse) SetMsPreviousState(v string)`

SetMsPreviousState sets MsPreviousState field to given value.

### HasMsPreviousState

`func (o *GetDhcpfailoverResponse) HasMsPreviousState() bool`

HasMsPreviousState returns a boolean if a field has been set.

### GetMsServer

`func (o *GetDhcpfailoverResponse) GetMsServer() string`

GetMsServer returns the MsServer field if non-nil, zero value otherwise.

### GetMsServerOk

`func (o *GetDhcpfailoverResponse) GetMsServerOk() (*string, bool)`

GetMsServerOk returns a tuple with the MsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsServer

`func (o *GetDhcpfailoverResponse) SetMsServer(v string)`

SetMsServer sets MsServer field to given value.

### HasMsServer

`func (o *GetDhcpfailoverResponse) HasMsServer() bool`

HasMsServer returns a boolean if a field has been set.

### GetMsSharedSecret

`func (o *GetDhcpfailoverResponse) GetMsSharedSecret() string`

GetMsSharedSecret returns the MsSharedSecret field if non-nil, zero value otherwise.

### GetMsSharedSecretOk

`func (o *GetDhcpfailoverResponse) GetMsSharedSecretOk() (*string, bool)`

GetMsSharedSecretOk returns a tuple with the MsSharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsSharedSecret

`func (o *GetDhcpfailoverResponse) SetMsSharedSecret(v string)`

SetMsSharedSecret sets MsSharedSecret field to given value.

### HasMsSharedSecret

`func (o *GetDhcpfailoverResponse) HasMsSharedSecret() bool`

HasMsSharedSecret returns a boolean if a field has been set.

### GetMsState

`func (o *GetDhcpfailoverResponse) GetMsState() string`

GetMsState returns the MsState field if non-nil, zero value otherwise.

### GetMsStateOk

`func (o *GetDhcpfailoverResponse) GetMsStateOk() (*string, bool)`

GetMsStateOk returns a tuple with the MsState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsState

`func (o *GetDhcpfailoverResponse) SetMsState(v string)`

SetMsState sets MsState field to given value.

### HasMsState

`func (o *GetDhcpfailoverResponse) HasMsState() bool`

HasMsState returns a boolean if a field has been set.

### GetMsSwitchoverInterval

`func (o *GetDhcpfailoverResponse) GetMsSwitchoverInterval() int64`

GetMsSwitchoverInterval returns the MsSwitchoverInterval field if non-nil, zero value otherwise.

### GetMsSwitchoverIntervalOk

`func (o *GetDhcpfailoverResponse) GetMsSwitchoverIntervalOk() (*int64, bool)`

GetMsSwitchoverIntervalOk returns a tuple with the MsSwitchoverInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsSwitchoverInterval

`func (o *GetDhcpfailoverResponse) SetMsSwitchoverInterval(v int64)`

SetMsSwitchoverInterval sets MsSwitchoverInterval field to given value.

### HasMsSwitchoverInterval

`func (o *GetDhcpfailoverResponse) HasMsSwitchoverInterval() bool`

HasMsSwitchoverInterval returns a boolean if a field has been set.

### GetName

`func (o *GetDhcpfailoverResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDhcpfailoverResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDhcpfailoverResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDhcpfailoverResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimary

`func (o *GetDhcpfailoverResponse) GetPrimary() string`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *GetDhcpfailoverResponse) GetPrimaryOk() (*string, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *GetDhcpfailoverResponse) SetPrimary(v string)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *GetDhcpfailoverResponse) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetPrimaryServerType

`func (o *GetDhcpfailoverResponse) GetPrimaryServerType() string`

GetPrimaryServerType returns the PrimaryServerType field if non-nil, zero value otherwise.

### GetPrimaryServerTypeOk

`func (o *GetDhcpfailoverResponse) GetPrimaryServerTypeOk() (*string, bool)`

GetPrimaryServerTypeOk returns a tuple with the PrimaryServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryServerType

`func (o *GetDhcpfailoverResponse) SetPrimaryServerType(v string)`

SetPrimaryServerType sets PrimaryServerType field to given value.

### HasPrimaryServerType

`func (o *GetDhcpfailoverResponse) HasPrimaryServerType() bool`

HasPrimaryServerType returns a boolean if a field has been set.

### GetPrimaryState

`func (o *GetDhcpfailoverResponse) GetPrimaryState() string`

GetPrimaryState returns the PrimaryState field if non-nil, zero value otherwise.

### GetPrimaryStateOk

`func (o *GetDhcpfailoverResponse) GetPrimaryStateOk() (*string, bool)`

GetPrimaryStateOk returns a tuple with the PrimaryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryState

`func (o *GetDhcpfailoverResponse) SetPrimaryState(v string)`

SetPrimaryState sets PrimaryState field to given value.

### HasPrimaryState

`func (o *GetDhcpfailoverResponse) HasPrimaryState() bool`

HasPrimaryState returns a boolean if a field has been set.

### GetRecycleLeases

`func (o *GetDhcpfailoverResponse) GetRecycleLeases() bool`

GetRecycleLeases returns the RecycleLeases field if non-nil, zero value otherwise.

### GetRecycleLeasesOk

`func (o *GetDhcpfailoverResponse) GetRecycleLeasesOk() (*bool, bool)`

GetRecycleLeasesOk returns a tuple with the RecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycleLeases

`func (o *GetDhcpfailoverResponse) SetRecycleLeases(v bool)`

SetRecycleLeases sets RecycleLeases field to given value.

### HasRecycleLeases

`func (o *GetDhcpfailoverResponse) HasRecycleLeases() bool`

HasRecycleLeases returns a boolean if a field has been set.

### GetSecondary

`func (o *GetDhcpfailoverResponse) GetSecondary() string`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *GetDhcpfailoverResponse) GetSecondaryOk() (*string, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *GetDhcpfailoverResponse) SetSecondary(v string)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *GetDhcpfailoverResponse) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.

### GetSecondaryServerType

`func (o *GetDhcpfailoverResponse) GetSecondaryServerType() string`

GetSecondaryServerType returns the SecondaryServerType field if non-nil, zero value otherwise.

### GetSecondaryServerTypeOk

`func (o *GetDhcpfailoverResponse) GetSecondaryServerTypeOk() (*string, bool)`

GetSecondaryServerTypeOk returns a tuple with the SecondaryServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryServerType

`func (o *GetDhcpfailoverResponse) SetSecondaryServerType(v string)`

SetSecondaryServerType sets SecondaryServerType field to given value.

### HasSecondaryServerType

`func (o *GetDhcpfailoverResponse) HasSecondaryServerType() bool`

HasSecondaryServerType returns a boolean if a field has been set.

### GetSecondaryState

`func (o *GetDhcpfailoverResponse) GetSecondaryState() string`

GetSecondaryState returns the SecondaryState field if non-nil, zero value otherwise.

### GetSecondaryStateOk

`func (o *GetDhcpfailoverResponse) GetSecondaryStateOk() (*string, bool)`

GetSecondaryStateOk returns a tuple with the SecondaryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryState

`func (o *GetDhcpfailoverResponse) SetSecondaryState(v string)`

SetSecondaryState sets SecondaryState field to given value.

### HasSecondaryState

`func (o *GetDhcpfailoverResponse) HasSecondaryState() bool`

HasSecondaryState returns a boolean if a field has been set.

### GetUseFailoverPort

`func (o *GetDhcpfailoverResponse) GetUseFailoverPort() bool`

GetUseFailoverPort returns the UseFailoverPort field if non-nil, zero value otherwise.

### GetUseFailoverPortOk

`func (o *GetDhcpfailoverResponse) GetUseFailoverPortOk() (*bool, bool)`

GetUseFailoverPortOk returns a tuple with the UseFailoverPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFailoverPort

`func (o *GetDhcpfailoverResponse) SetUseFailoverPort(v bool)`

SetUseFailoverPort sets UseFailoverPort field to given value.

### HasUseFailoverPort

`func (o *GetDhcpfailoverResponse) HasUseFailoverPort() bool`

HasUseFailoverPort returns a boolean if a field has been set.

### GetUseMsSwitchoverInterval

`func (o *GetDhcpfailoverResponse) GetUseMsSwitchoverInterval() bool`

GetUseMsSwitchoverInterval returns the UseMsSwitchoverInterval field if non-nil, zero value otherwise.

### GetUseMsSwitchoverIntervalOk

`func (o *GetDhcpfailoverResponse) GetUseMsSwitchoverIntervalOk() (*bool, bool)`

GetUseMsSwitchoverIntervalOk returns a tuple with the UseMsSwitchoverInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMsSwitchoverInterval

`func (o *GetDhcpfailoverResponse) SetUseMsSwitchoverInterval(v bool)`

SetUseMsSwitchoverInterval sets UseMsSwitchoverInterval field to given value.

### HasUseMsSwitchoverInterval

`func (o *GetDhcpfailoverResponse) HasUseMsSwitchoverInterval() bool`

HasUseMsSwitchoverInterval returns a boolean if a field has been set.

### GetUseRecycleLeases

`func (o *GetDhcpfailoverResponse) GetUseRecycleLeases() bool`

GetUseRecycleLeases returns the UseRecycleLeases field if non-nil, zero value otherwise.

### GetUseRecycleLeasesOk

`func (o *GetDhcpfailoverResponse) GetUseRecycleLeasesOk() (*bool, bool)`

GetUseRecycleLeasesOk returns a tuple with the UseRecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecycleLeases

`func (o *GetDhcpfailoverResponse) SetUseRecycleLeases(v bool)`

SetUseRecycleLeases sets UseRecycleLeases field to given value.

### HasUseRecycleLeases

`func (o *GetDhcpfailoverResponse) HasUseRecycleLeases() bool`

HasUseRecycleLeases returns a boolean if a field has been set.

### GetResult

`func (o *GetDhcpfailoverResponse) GetResult() Dhcpfailover`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDhcpfailoverResponse) GetResultOk() (*Dhcpfailover, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDhcpfailoverResponse) SetResult(v Dhcpfailover)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDhcpfailoverResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


