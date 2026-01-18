# DtcServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AutoCreateHostRecord** | Pointer to **bool** | Enabling this option will auto-create a single read-only A/AAAA/CNAME record corresponding to the configured hostname and update it if the hostname changes. | [optional] 
**Comment** | Pointer to **string** | Comment for the DTC Server; maximum 256 characters. | [optional] 
**Disable** | Pointer to **bool** | Determines whether the DTC Server is disabled or not. When this is set to False, the fixed address is enabled. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Health** | Pointer to [**DtcServerHealth**](DtcServerHealth.md) |  | [optional] 
**Host** | Pointer to **string** | The address or FQDN of the server. | [optional] 
**Monitors** | Pointer to [**[]DtcServerMonitors**](DtcServerMonitors.md) | List of IP/FQDN and monitor pairs to be used for additional monitoring. | [optional] 
**Name** | Pointer to **string** | The DTC Server display name. | [optional] 
**SniHostname** | Pointer to **string** | The hostname for Server Name Indication (SNI) in FQDN format. | [optional] 
**UseSniHostname** | Pointer to **bool** | Use flag for: sni_hostname | [optional] 

## Methods

### NewDtcServer

`func NewDtcServer() *DtcServer`

NewDtcServer instantiates a new DtcServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcServerWithDefaults

`func NewDtcServerWithDefaults() *DtcServer`

NewDtcServerWithDefaults instantiates a new DtcServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DtcServer) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DtcServer) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DtcServer) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DtcServer) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAutoCreateHostRecord

`func (o *DtcServer) GetAutoCreateHostRecord() bool`

GetAutoCreateHostRecord returns the AutoCreateHostRecord field if non-nil, zero value otherwise.

### GetAutoCreateHostRecordOk

`func (o *DtcServer) GetAutoCreateHostRecordOk() (*bool, bool)`

GetAutoCreateHostRecordOk returns a tuple with the AutoCreateHostRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateHostRecord

`func (o *DtcServer) SetAutoCreateHostRecord(v bool)`

SetAutoCreateHostRecord sets AutoCreateHostRecord field to given value.

### HasAutoCreateHostRecord

`func (o *DtcServer) HasAutoCreateHostRecord() bool`

HasAutoCreateHostRecord returns a boolean if a field has been set.

### GetComment

`func (o *DtcServer) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DtcServer) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DtcServer) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DtcServer) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisable

`func (o *DtcServer) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *DtcServer) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *DtcServer) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *DtcServer) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *DtcServer) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *DtcServer) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *DtcServer) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *DtcServer) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *DtcServer) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *DtcServer) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *DtcServer) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *DtcServer) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *DtcServer) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *DtcServer) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *DtcServer) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *DtcServer) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetHealth

`func (o *DtcServer) GetHealth() DtcServerHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *DtcServer) GetHealthOk() (*DtcServerHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *DtcServer) SetHealth(v DtcServerHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *DtcServer) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHost

`func (o *DtcServer) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DtcServer) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DtcServer) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *DtcServer) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetMonitors

`func (o *DtcServer) GetMonitors() []DtcServerMonitors`

GetMonitors returns the Monitors field if non-nil, zero value otherwise.

### GetMonitorsOk

`func (o *DtcServer) GetMonitorsOk() (*[]DtcServerMonitors, bool)`

GetMonitorsOk returns a tuple with the Monitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitors

`func (o *DtcServer) SetMonitors(v []DtcServerMonitors)`

SetMonitors sets Monitors field to given value.

### HasMonitors

`func (o *DtcServer) HasMonitors() bool`

HasMonitors returns a boolean if a field has been set.

### GetName

`func (o *DtcServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtcServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtcServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtcServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSniHostname

`func (o *DtcServer) GetSniHostname() string`

GetSniHostname returns the SniHostname field if non-nil, zero value otherwise.

### GetSniHostnameOk

`func (o *DtcServer) GetSniHostnameOk() (*string, bool)`

GetSniHostnameOk returns a tuple with the SniHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSniHostname

`func (o *DtcServer) SetSniHostname(v string)`

SetSniHostname sets SniHostname field to given value.

### HasSniHostname

`func (o *DtcServer) HasSniHostname() bool`

HasSniHostname returns a boolean if a field has been set.

### GetUseSniHostname

`func (o *DtcServer) GetUseSniHostname() bool`

GetUseSniHostname returns the UseSniHostname field if non-nil, zero value otherwise.

### GetUseSniHostnameOk

`func (o *DtcServer) GetUseSniHostnameOk() (*bool, bool)`

GetUseSniHostnameOk returns a tuple with the UseSniHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSniHostname

`func (o *DtcServer) SetUseSniHostname(v bool)`

SetUseSniHostname sets UseSniHostname field to given value.

### HasUseSniHostname

`func (o *DtcServer) HasUseSniHostname() bool`

HasUseSniHostname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


