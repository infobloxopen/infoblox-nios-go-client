# AdmingroupAdminToplevelCommands

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ps** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**Iostat** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**Netstat** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**Vmstat** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**Tcpdump** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**Rndc** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**Sar** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**Resilver** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**RestartProduct** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**Scrape** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SamlRestart** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**Synctime** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**EnableAll** | Pointer to **bool** | If True then enable all fields | [optional] 
**DisableAll** | Pointer to **bool** | If True then disable all fields | [optional] 

## Methods

### NewAdmingroupAdminToplevelCommands

`func NewAdmingroupAdminToplevelCommands() *AdmingroupAdminToplevelCommands`

NewAdmingroupAdminToplevelCommands instantiates a new AdmingroupAdminToplevelCommands object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdmingroupAdminToplevelCommandsWithDefaults

`func NewAdmingroupAdminToplevelCommandsWithDefaults() *AdmingroupAdminToplevelCommands`

NewAdmingroupAdminToplevelCommandsWithDefaults instantiates a new AdmingroupAdminToplevelCommands object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPs

`func (o *AdmingroupAdminToplevelCommands) GetPs() bool`

GetPs returns the Ps field if non-nil, zero value otherwise.

### GetPsOk

`func (o *AdmingroupAdminToplevelCommands) GetPsOk() (*bool, bool)`

GetPsOk returns a tuple with the Ps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPs

`func (o *AdmingroupAdminToplevelCommands) SetPs(v bool)`

SetPs sets Ps field to given value.

### HasPs

`func (o *AdmingroupAdminToplevelCommands) HasPs() bool`

HasPs returns a boolean if a field has been set.

### GetIostat

`func (o *AdmingroupAdminToplevelCommands) GetIostat() bool`

GetIostat returns the Iostat field if non-nil, zero value otherwise.

### GetIostatOk

`func (o *AdmingroupAdminToplevelCommands) GetIostatOk() (*bool, bool)`

GetIostatOk returns a tuple with the Iostat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIostat

`func (o *AdmingroupAdminToplevelCommands) SetIostat(v bool)`

SetIostat sets Iostat field to given value.

### HasIostat

`func (o *AdmingroupAdminToplevelCommands) HasIostat() bool`

HasIostat returns a boolean if a field has been set.

### GetNetstat

`func (o *AdmingroupAdminToplevelCommands) GetNetstat() bool`

GetNetstat returns the Netstat field if non-nil, zero value otherwise.

### GetNetstatOk

`func (o *AdmingroupAdminToplevelCommands) GetNetstatOk() (*bool, bool)`

GetNetstatOk returns a tuple with the Netstat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetstat

`func (o *AdmingroupAdminToplevelCommands) SetNetstat(v bool)`

SetNetstat sets Netstat field to given value.

### HasNetstat

`func (o *AdmingroupAdminToplevelCommands) HasNetstat() bool`

HasNetstat returns a boolean if a field has been set.

### GetVmstat

`func (o *AdmingroupAdminToplevelCommands) GetVmstat() bool`

GetVmstat returns the Vmstat field if non-nil, zero value otherwise.

### GetVmstatOk

`func (o *AdmingroupAdminToplevelCommands) GetVmstatOk() (*bool, bool)`

GetVmstatOk returns a tuple with the Vmstat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmstat

`func (o *AdmingroupAdminToplevelCommands) SetVmstat(v bool)`

SetVmstat sets Vmstat field to given value.

### HasVmstat

`func (o *AdmingroupAdminToplevelCommands) HasVmstat() bool`

HasVmstat returns a boolean if a field has been set.

### GetTcpdump

`func (o *AdmingroupAdminToplevelCommands) GetTcpdump() bool`

GetTcpdump returns the Tcpdump field if non-nil, zero value otherwise.

### GetTcpdumpOk

`func (o *AdmingroupAdminToplevelCommands) GetTcpdumpOk() (*bool, bool)`

GetTcpdumpOk returns a tuple with the Tcpdump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpdump

`func (o *AdmingroupAdminToplevelCommands) SetTcpdump(v bool)`

SetTcpdump sets Tcpdump field to given value.

### HasTcpdump

`func (o *AdmingroupAdminToplevelCommands) HasTcpdump() bool`

HasTcpdump returns a boolean if a field has been set.

### GetRndc

`func (o *AdmingroupAdminToplevelCommands) GetRndc() bool`

GetRndc returns the Rndc field if non-nil, zero value otherwise.

### GetRndcOk

`func (o *AdmingroupAdminToplevelCommands) GetRndcOk() (*bool, bool)`

GetRndcOk returns a tuple with the Rndc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRndc

`func (o *AdmingroupAdminToplevelCommands) SetRndc(v bool)`

SetRndc sets Rndc field to given value.

### HasRndc

`func (o *AdmingroupAdminToplevelCommands) HasRndc() bool`

HasRndc returns a boolean if a field has been set.

### GetSar

`func (o *AdmingroupAdminToplevelCommands) GetSar() bool`

GetSar returns the Sar field if non-nil, zero value otherwise.

### GetSarOk

`func (o *AdmingroupAdminToplevelCommands) GetSarOk() (*bool, bool)`

GetSarOk returns a tuple with the Sar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSar

`func (o *AdmingroupAdminToplevelCommands) SetSar(v bool)`

SetSar sets Sar field to given value.

### HasSar

`func (o *AdmingroupAdminToplevelCommands) HasSar() bool`

HasSar returns a boolean if a field has been set.

### GetResilver

`func (o *AdmingroupAdminToplevelCommands) GetResilver() bool`

GetResilver returns the Resilver field if non-nil, zero value otherwise.

### GetResilverOk

`func (o *AdmingroupAdminToplevelCommands) GetResilverOk() (*bool, bool)`

GetResilverOk returns a tuple with the Resilver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResilver

`func (o *AdmingroupAdminToplevelCommands) SetResilver(v bool)`

SetResilver sets Resilver field to given value.

### HasResilver

`func (o *AdmingroupAdminToplevelCommands) HasResilver() bool`

HasResilver returns a boolean if a field has been set.

### GetRestartProduct

`func (o *AdmingroupAdminToplevelCommands) GetRestartProduct() bool`

GetRestartProduct returns the RestartProduct field if non-nil, zero value otherwise.

### GetRestartProductOk

`func (o *AdmingroupAdminToplevelCommands) GetRestartProductOk() (*bool, bool)`

GetRestartProductOk returns a tuple with the RestartProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartProduct

`func (o *AdmingroupAdminToplevelCommands) SetRestartProduct(v bool)`

SetRestartProduct sets RestartProduct field to given value.

### HasRestartProduct

`func (o *AdmingroupAdminToplevelCommands) HasRestartProduct() bool`

HasRestartProduct returns a boolean if a field has been set.

### GetScrape

`func (o *AdmingroupAdminToplevelCommands) GetScrape() bool`

GetScrape returns the Scrape field if non-nil, zero value otherwise.

### GetScrapeOk

`func (o *AdmingroupAdminToplevelCommands) GetScrapeOk() (*bool, bool)`

GetScrapeOk returns a tuple with the Scrape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrape

`func (o *AdmingroupAdminToplevelCommands) SetScrape(v bool)`

SetScrape sets Scrape field to given value.

### HasScrape

`func (o *AdmingroupAdminToplevelCommands) HasScrape() bool`

HasScrape returns a boolean if a field has been set.

### GetSamlRestart

`func (o *AdmingroupAdminToplevelCommands) GetSamlRestart() bool`

GetSamlRestart returns the SamlRestart field if non-nil, zero value otherwise.

### GetSamlRestartOk

`func (o *AdmingroupAdminToplevelCommands) GetSamlRestartOk() (*bool, bool)`

GetSamlRestartOk returns a tuple with the SamlRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlRestart

`func (o *AdmingroupAdminToplevelCommands) SetSamlRestart(v bool)`

SetSamlRestart sets SamlRestart field to given value.

### HasSamlRestart

`func (o *AdmingroupAdminToplevelCommands) HasSamlRestart() bool`

HasSamlRestart returns a boolean if a field has been set.

### GetSynctime

`func (o *AdmingroupAdminToplevelCommands) GetSynctime() bool`

GetSynctime returns the Synctime field if non-nil, zero value otherwise.

### GetSynctimeOk

`func (o *AdmingroupAdminToplevelCommands) GetSynctimeOk() (*bool, bool)`

GetSynctimeOk returns a tuple with the Synctime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynctime

`func (o *AdmingroupAdminToplevelCommands) SetSynctime(v bool)`

SetSynctime sets Synctime field to given value.

### HasSynctime

`func (o *AdmingroupAdminToplevelCommands) HasSynctime() bool`

HasSynctime returns a boolean if a field has been set.

### GetEnableAll

`func (o *AdmingroupAdminToplevelCommands) GetEnableAll() bool`

GetEnableAll returns the EnableAll field if non-nil, zero value otherwise.

### GetEnableAllOk

`func (o *AdmingroupAdminToplevelCommands) GetEnableAllOk() (*bool, bool)`

GetEnableAllOk returns a tuple with the EnableAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAll

`func (o *AdmingroupAdminToplevelCommands) SetEnableAll(v bool)`

SetEnableAll sets EnableAll field to given value.

### HasEnableAll

`func (o *AdmingroupAdminToplevelCommands) HasEnableAll() bool`

HasEnableAll returns a boolean if a field has been set.

### GetDisableAll

`func (o *AdmingroupAdminToplevelCommands) GetDisableAll() bool`

GetDisableAll returns the DisableAll field if non-nil, zero value otherwise.

### GetDisableAllOk

`func (o *AdmingroupAdminToplevelCommands) GetDisableAllOk() (*bool, bool)`

GetDisableAllOk returns a tuple with the DisableAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAll

`func (o *AdmingroupAdminToplevelCommands) SetDisableAll(v bool)`

SetDisableAll sets DisableAll field to given value.

### HasDisableAll

`func (o *AdmingroupAdminToplevelCommands) HasDisableAll() bool`

HasDisableAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


