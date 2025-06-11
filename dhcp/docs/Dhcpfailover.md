# Dhcpfailover

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AssociationType** | Pointer to **string** | The value indicating whether the failover assoctaion is Microsoft or Grid based. This is a read-only attribute. | [optional] [readonly] 
**Comment** | Pointer to **string** | A descriptive comment about a DHCP failover object. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
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

## Methods

### NewDhcpfailover

`func NewDhcpfailover() *Dhcpfailover`

NewDhcpfailover instantiates a new Dhcpfailover object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpfailoverWithDefaults

`func NewDhcpfailoverWithDefaults() *Dhcpfailover`

NewDhcpfailoverWithDefaults instantiates a new Dhcpfailover object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Dhcpfailover) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Dhcpfailover) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Dhcpfailover) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Dhcpfailover) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAssociationType

`func (o *Dhcpfailover) GetAssociationType() string`

GetAssociationType returns the AssociationType field if non-nil, zero value otherwise.

### GetAssociationTypeOk

`func (o *Dhcpfailover) GetAssociationTypeOk() (*string, bool)`

GetAssociationTypeOk returns a tuple with the AssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationType

`func (o *Dhcpfailover) SetAssociationType(v string)`

SetAssociationType sets AssociationType field to given value.

### HasAssociationType

`func (o *Dhcpfailover) HasAssociationType() bool`

HasAssociationType returns a boolean if a field has been set.

### GetComment

`func (o *Dhcpfailover) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Dhcpfailover) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Dhcpfailover) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Dhcpfailover) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtattrs

`func (o *Dhcpfailover) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *Dhcpfailover) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *Dhcpfailover) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *Dhcpfailover) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetFailoverPort

`func (o *Dhcpfailover) GetFailoverPort() int64`

GetFailoverPort returns the FailoverPort field if non-nil, zero value otherwise.

### GetFailoverPortOk

`func (o *Dhcpfailover) GetFailoverPortOk() (*int64, bool)`

GetFailoverPortOk returns a tuple with the FailoverPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverPort

`func (o *Dhcpfailover) SetFailoverPort(v int64)`

SetFailoverPort sets FailoverPort field to given value.

### HasFailoverPort

`func (o *Dhcpfailover) HasFailoverPort() bool`

HasFailoverPort returns a boolean if a field has been set.

### GetLoadBalanceSplit

`func (o *Dhcpfailover) GetLoadBalanceSplit() int64`

GetLoadBalanceSplit returns the LoadBalanceSplit field if non-nil, zero value otherwise.

### GetLoadBalanceSplitOk

`func (o *Dhcpfailover) GetLoadBalanceSplitOk() (*int64, bool)`

GetLoadBalanceSplitOk returns a tuple with the LoadBalanceSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalanceSplit

`func (o *Dhcpfailover) SetLoadBalanceSplit(v int64)`

SetLoadBalanceSplit sets LoadBalanceSplit field to given value.

### HasLoadBalanceSplit

`func (o *Dhcpfailover) HasLoadBalanceSplit() bool`

HasLoadBalanceSplit returns a boolean if a field has been set.

### GetMaxClientLeadTime

`func (o *Dhcpfailover) GetMaxClientLeadTime() int64`

GetMaxClientLeadTime returns the MaxClientLeadTime field if non-nil, zero value otherwise.

### GetMaxClientLeadTimeOk

`func (o *Dhcpfailover) GetMaxClientLeadTimeOk() (*int64, bool)`

GetMaxClientLeadTimeOk returns a tuple with the MaxClientLeadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxClientLeadTime

`func (o *Dhcpfailover) SetMaxClientLeadTime(v int64)`

SetMaxClientLeadTime sets MaxClientLeadTime field to given value.

### HasMaxClientLeadTime

`func (o *Dhcpfailover) HasMaxClientLeadTime() bool`

HasMaxClientLeadTime returns a boolean if a field has been set.

### GetMaxLoadBalanceDelay

`func (o *Dhcpfailover) GetMaxLoadBalanceDelay() int64`

GetMaxLoadBalanceDelay returns the MaxLoadBalanceDelay field if non-nil, zero value otherwise.

### GetMaxLoadBalanceDelayOk

`func (o *Dhcpfailover) GetMaxLoadBalanceDelayOk() (*int64, bool)`

GetMaxLoadBalanceDelayOk returns a tuple with the MaxLoadBalanceDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLoadBalanceDelay

`func (o *Dhcpfailover) SetMaxLoadBalanceDelay(v int64)`

SetMaxLoadBalanceDelay sets MaxLoadBalanceDelay field to given value.

### HasMaxLoadBalanceDelay

`func (o *Dhcpfailover) HasMaxLoadBalanceDelay() bool`

HasMaxLoadBalanceDelay returns a boolean if a field has been set.

### GetMaxResponseDelay

`func (o *Dhcpfailover) GetMaxResponseDelay() int64`

GetMaxResponseDelay returns the MaxResponseDelay field if non-nil, zero value otherwise.

### GetMaxResponseDelayOk

`func (o *Dhcpfailover) GetMaxResponseDelayOk() (*int64, bool)`

GetMaxResponseDelayOk returns a tuple with the MaxResponseDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResponseDelay

`func (o *Dhcpfailover) SetMaxResponseDelay(v int64)`

SetMaxResponseDelay sets MaxResponseDelay field to given value.

### HasMaxResponseDelay

`func (o *Dhcpfailover) HasMaxResponseDelay() bool`

HasMaxResponseDelay returns a boolean if a field has been set.

### GetMaxUnackedUpdates

`func (o *Dhcpfailover) GetMaxUnackedUpdates() int64`

GetMaxUnackedUpdates returns the MaxUnackedUpdates field if non-nil, zero value otherwise.

### GetMaxUnackedUpdatesOk

`func (o *Dhcpfailover) GetMaxUnackedUpdatesOk() (*int64, bool)`

GetMaxUnackedUpdatesOk returns a tuple with the MaxUnackedUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUnackedUpdates

`func (o *Dhcpfailover) SetMaxUnackedUpdates(v int64)`

SetMaxUnackedUpdates sets MaxUnackedUpdates field to given value.

### HasMaxUnackedUpdates

`func (o *Dhcpfailover) HasMaxUnackedUpdates() bool`

HasMaxUnackedUpdates returns a boolean if a field has been set.

### GetMsAssociationMode

`func (o *Dhcpfailover) GetMsAssociationMode() string`

GetMsAssociationMode returns the MsAssociationMode field if non-nil, zero value otherwise.

### GetMsAssociationModeOk

`func (o *Dhcpfailover) GetMsAssociationModeOk() (*string, bool)`

GetMsAssociationModeOk returns a tuple with the MsAssociationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAssociationMode

`func (o *Dhcpfailover) SetMsAssociationMode(v string)`

SetMsAssociationMode sets MsAssociationMode field to given value.

### HasMsAssociationMode

`func (o *Dhcpfailover) HasMsAssociationMode() bool`

HasMsAssociationMode returns a boolean if a field has been set.

### GetMsEnableAuthentication

`func (o *Dhcpfailover) GetMsEnableAuthentication() bool`

GetMsEnableAuthentication returns the MsEnableAuthentication field if non-nil, zero value otherwise.

### GetMsEnableAuthenticationOk

`func (o *Dhcpfailover) GetMsEnableAuthenticationOk() (*bool, bool)`

GetMsEnableAuthenticationOk returns a tuple with the MsEnableAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsEnableAuthentication

`func (o *Dhcpfailover) SetMsEnableAuthentication(v bool)`

SetMsEnableAuthentication sets MsEnableAuthentication field to given value.

### HasMsEnableAuthentication

`func (o *Dhcpfailover) HasMsEnableAuthentication() bool`

HasMsEnableAuthentication returns a boolean if a field has been set.

### GetMsEnableSwitchoverInterval

`func (o *Dhcpfailover) GetMsEnableSwitchoverInterval() bool`

GetMsEnableSwitchoverInterval returns the MsEnableSwitchoverInterval field if non-nil, zero value otherwise.

### GetMsEnableSwitchoverIntervalOk

`func (o *Dhcpfailover) GetMsEnableSwitchoverIntervalOk() (*bool, bool)`

GetMsEnableSwitchoverIntervalOk returns a tuple with the MsEnableSwitchoverInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsEnableSwitchoverInterval

`func (o *Dhcpfailover) SetMsEnableSwitchoverInterval(v bool)`

SetMsEnableSwitchoverInterval sets MsEnableSwitchoverInterval field to given value.

### HasMsEnableSwitchoverInterval

`func (o *Dhcpfailover) HasMsEnableSwitchoverInterval() bool`

HasMsEnableSwitchoverInterval returns a boolean if a field has been set.

### GetMsFailoverMode

`func (o *Dhcpfailover) GetMsFailoverMode() string`

GetMsFailoverMode returns the MsFailoverMode field if non-nil, zero value otherwise.

### GetMsFailoverModeOk

`func (o *Dhcpfailover) GetMsFailoverModeOk() (*string, bool)`

GetMsFailoverModeOk returns a tuple with the MsFailoverMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsFailoverMode

`func (o *Dhcpfailover) SetMsFailoverMode(v string)`

SetMsFailoverMode sets MsFailoverMode field to given value.

### HasMsFailoverMode

`func (o *Dhcpfailover) HasMsFailoverMode() bool`

HasMsFailoverMode returns a boolean if a field has been set.

### GetMsFailoverPartner

`func (o *Dhcpfailover) GetMsFailoverPartner() string`

GetMsFailoverPartner returns the MsFailoverPartner field if non-nil, zero value otherwise.

### GetMsFailoverPartnerOk

`func (o *Dhcpfailover) GetMsFailoverPartnerOk() (*string, bool)`

GetMsFailoverPartnerOk returns a tuple with the MsFailoverPartner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsFailoverPartner

`func (o *Dhcpfailover) SetMsFailoverPartner(v string)`

SetMsFailoverPartner sets MsFailoverPartner field to given value.

### HasMsFailoverPartner

`func (o *Dhcpfailover) HasMsFailoverPartner() bool`

HasMsFailoverPartner returns a boolean if a field has been set.

### GetMsHotstandbyPartnerRole

`func (o *Dhcpfailover) GetMsHotstandbyPartnerRole() string`

GetMsHotstandbyPartnerRole returns the MsHotstandbyPartnerRole field if non-nil, zero value otherwise.

### GetMsHotstandbyPartnerRoleOk

`func (o *Dhcpfailover) GetMsHotstandbyPartnerRoleOk() (*string, bool)`

GetMsHotstandbyPartnerRoleOk returns a tuple with the MsHotstandbyPartnerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsHotstandbyPartnerRole

`func (o *Dhcpfailover) SetMsHotstandbyPartnerRole(v string)`

SetMsHotstandbyPartnerRole sets MsHotstandbyPartnerRole field to given value.

### HasMsHotstandbyPartnerRole

`func (o *Dhcpfailover) HasMsHotstandbyPartnerRole() bool`

HasMsHotstandbyPartnerRole returns a boolean if a field has been set.

### GetMsIsConflict

`func (o *Dhcpfailover) GetMsIsConflict() bool`

GetMsIsConflict returns the MsIsConflict field if non-nil, zero value otherwise.

### GetMsIsConflictOk

`func (o *Dhcpfailover) GetMsIsConflictOk() (*bool, bool)`

GetMsIsConflictOk returns a tuple with the MsIsConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsIsConflict

`func (o *Dhcpfailover) SetMsIsConflict(v bool)`

SetMsIsConflict sets MsIsConflict field to given value.

### HasMsIsConflict

`func (o *Dhcpfailover) HasMsIsConflict() bool`

HasMsIsConflict returns a boolean if a field has been set.

### GetMsPreviousState

`func (o *Dhcpfailover) GetMsPreviousState() string`

GetMsPreviousState returns the MsPreviousState field if non-nil, zero value otherwise.

### GetMsPreviousStateOk

`func (o *Dhcpfailover) GetMsPreviousStateOk() (*string, bool)`

GetMsPreviousStateOk returns a tuple with the MsPreviousState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsPreviousState

`func (o *Dhcpfailover) SetMsPreviousState(v string)`

SetMsPreviousState sets MsPreviousState field to given value.

### HasMsPreviousState

`func (o *Dhcpfailover) HasMsPreviousState() bool`

HasMsPreviousState returns a boolean if a field has been set.

### GetMsServer

`func (o *Dhcpfailover) GetMsServer() string`

GetMsServer returns the MsServer field if non-nil, zero value otherwise.

### GetMsServerOk

`func (o *Dhcpfailover) GetMsServerOk() (*string, bool)`

GetMsServerOk returns a tuple with the MsServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsServer

`func (o *Dhcpfailover) SetMsServer(v string)`

SetMsServer sets MsServer field to given value.

### HasMsServer

`func (o *Dhcpfailover) HasMsServer() bool`

HasMsServer returns a boolean if a field has been set.

### GetMsSharedSecret

`func (o *Dhcpfailover) GetMsSharedSecret() string`

GetMsSharedSecret returns the MsSharedSecret field if non-nil, zero value otherwise.

### GetMsSharedSecretOk

`func (o *Dhcpfailover) GetMsSharedSecretOk() (*string, bool)`

GetMsSharedSecretOk returns a tuple with the MsSharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsSharedSecret

`func (o *Dhcpfailover) SetMsSharedSecret(v string)`

SetMsSharedSecret sets MsSharedSecret field to given value.

### HasMsSharedSecret

`func (o *Dhcpfailover) HasMsSharedSecret() bool`

HasMsSharedSecret returns a boolean if a field has been set.

### GetMsState

`func (o *Dhcpfailover) GetMsState() string`

GetMsState returns the MsState field if non-nil, zero value otherwise.

### GetMsStateOk

`func (o *Dhcpfailover) GetMsStateOk() (*string, bool)`

GetMsStateOk returns a tuple with the MsState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsState

`func (o *Dhcpfailover) SetMsState(v string)`

SetMsState sets MsState field to given value.

### HasMsState

`func (o *Dhcpfailover) HasMsState() bool`

HasMsState returns a boolean if a field has been set.

### GetMsSwitchoverInterval

`func (o *Dhcpfailover) GetMsSwitchoverInterval() int64`

GetMsSwitchoverInterval returns the MsSwitchoverInterval field if non-nil, zero value otherwise.

### GetMsSwitchoverIntervalOk

`func (o *Dhcpfailover) GetMsSwitchoverIntervalOk() (*int64, bool)`

GetMsSwitchoverIntervalOk returns a tuple with the MsSwitchoverInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsSwitchoverInterval

`func (o *Dhcpfailover) SetMsSwitchoverInterval(v int64)`

SetMsSwitchoverInterval sets MsSwitchoverInterval field to given value.

### HasMsSwitchoverInterval

`func (o *Dhcpfailover) HasMsSwitchoverInterval() bool`

HasMsSwitchoverInterval returns a boolean if a field has been set.

### GetName

`func (o *Dhcpfailover) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dhcpfailover) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dhcpfailover) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Dhcpfailover) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimary

`func (o *Dhcpfailover) GetPrimary() string`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *Dhcpfailover) GetPrimaryOk() (*string, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *Dhcpfailover) SetPrimary(v string)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *Dhcpfailover) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetPrimaryServerType

`func (o *Dhcpfailover) GetPrimaryServerType() string`

GetPrimaryServerType returns the PrimaryServerType field if non-nil, zero value otherwise.

### GetPrimaryServerTypeOk

`func (o *Dhcpfailover) GetPrimaryServerTypeOk() (*string, bool)`

GetPrimaryServerTypeOk returns a tuple with the PrimaryServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryServerType

`func (o *Dhcpfailover) SetPrimaryServerType(v string)`

SetPrimaryServerType sets PrimaryServerType field to given value.

### HasPrimaryServerType

`func (o *Dhcpfailover) HasPrimaryServerType() bool`

HasPrimaryServerType returns a boolean if a field has been set.

### GetPrimaryState

`func (o *Dhcpfailover) GetPrimaryState() string`

GetPrimaryState returns the PrimaryState field if non-nil, zero value otherwise.

### GetPrimaryStateOk

`func (o *Dhcpfailover) GetPrimaryStateOk() (*string, bool)`

GetPrimaryStateOk returns a tuple with the PrimaryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryState

`func (o *Dhcpfailover) SetPrimaryState(v string)`

SetPrimaryState sets PrimaryState field to given value.

### HasPrimaryState

`func (o *Dhcpfailover) HasPrimaryState() bool`

HasPrimaryState returns a boolean if a field has been set.

### GetRecycleLeases

`func (o *Dhcpfailover) GetRecycleLeases() bool`

GetRecycleLeases returns the RecycleLeases field if non-nil, zero value otherwise.

### GetRecycleLeasesOk

`func (o *Dhcpfailover) GetRecycleLeasesOk() (*bool, bool)`

GetRecycleLeasesOk returns a tuple with the RecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycleLeases

`func (o *Dhcpfailover) SetRecycleLeases(v bool)`

SetRecycleLeases sets RecycleLeases field to given value.

### HasRecycleLeases

`func (o *Dhcpfailover) HasRecycleLeases() bool`

HasRecycleLeases returns a boolean if a field has been set.

### GetSecondary

`func (o *Dhcpfailover) GetSecondary() string`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *Dhcpfailover) GetSecondaryOk() (*string, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *Dhcpfailover) SetSecondary(v string)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *Dhcpfailover) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.

### GetSecondaryServerType

`func (o *Dhcpfailover) GetSecondaryServerType() string`

GetSecondaryServerType returns the SecondaryServerType field if non-nil, zero value otherwise.

### GetSecondaryServerTypeOk

`func (o *Dhcpfailover) GetSecondaryServerTypeOk() (*string, bool)`

GetSecondaryServerTypeOk returns a tuple with the SecondaryServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryServerType

`func (o *Dhcpfailover) SetSecondaryServerType(v string)`

SetSecondaryServerType sets SecondaryServerType field to given value.

### HasSecondaryServerType

`func (o *Dhcpfailover) HasSecondaryServerType() bool`

HasSecondaryServerType returns a boolean if a field has been set.

### GetSecondaryState

`func (o *Dhcpfailover) GetSecondaryState() string`

GetSecondaryState returns the SecondaryState field if non-nil, zero value otherwise.

### GetSecondaryStateOk

`func (o *Dhcpfailover) GetSecondaryStateOk() (*string, bool)`

GetSecondaryStateOk returns a tuple with the SecondaryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryState

`func (o *Dhcpfailover) SetSecondaryState(v string)`

SetSecondaryState sets SecondaryState field to given value.

### HasSecondaryState

`func (o *Dhcpfailover) HasSecondaryState() bool`

HasSecondaryState returns a boolean if a field has been set.

### GetUseFailoverPort

`func (o *Dhcpfailover) GetUseFailoverPort() bool`

GetUseFailoverPort returns the UseFailoverPort field if non-nil, zero value otherwise.

### GetUseFailoverPortOk

`func (o *Dhcpfailover) GetUseFailoverPortOk() (*bool, bool)`

GetUseFailoverPortOk returns a tuple with the UseFailoverPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFailoverPort

`func (o *Dhcpfailover) SetUseFailoverPort(v bool)`

SetUseFailoverPort sets UseFailoverPort field to given value.

### HasUseFailoverPort

`func (o *Dhcpfailover) HasUseFailoverPort() bool`

HasUseFailoverPort returns a boolean if a field has been set.

### GetUseMsSwitchoverInterval

`func (o *Dhcpfailover) GetUseMsSwitchoverInterval() bool`

GetUseMsSwitchoverInterval returns the UseMsSwitchoverInterval field if non-nil, zero value otherwise.

### GetUseMsSwitchoverIntervalOk

`func (o *Dhcpfailover) GetUseMsSwitchoverIntervalOk() (*bool, bool)`

GetUseMsSwitchoverIntervalOk returns a tuple with the UseMsSwitchoverInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseMsSwitchoverInterval

`func (o *Dhcpfailover) SetUseMsSwitchoverInterval(v bool)`

SetUseMsSwitchoverInterval sets UseMsSwitchoverInterval field to given value.

### HasUseMsSwitchoverInterval

`func (o *Dhcpfailover) HasUseMsSwitchoverInterval() bool`

HasUseMsSwitchoverInterval returns a boolean if a field has been set.

### GetUseRecycleLeases

`func (o *Dhcpfailover) GetUseRecycleLeases() bool`

GetUseRecycleLeases returns the UseRecycleLeases field if non-nil, zero value otherwise.

### GetUseRecycleLeasesOk

`func (o *Dhcpfailover) GetUseRecycleLeasesOk() (*bool, bool)`

GetUseRecycleLeasesOk returns a tuple with the UseRecycleLeases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecycleLeases

`func (o *Dhcpfailover) SetUseRecycleLeases(v bool)`

SetUseRecycleLeases sets UseRecycleLeases field to given value.

### HasUseRecycleLeases

`func (o *Dhcpfailover) HasUseRecycleLeases() bool`

HasUseRecycleLeases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


