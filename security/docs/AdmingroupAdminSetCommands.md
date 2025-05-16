# AdmingroupAdminSetCommands

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SetAdminGroupAcl** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**EtBfd** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetBfd** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetBgp** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetCleanMscache** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetDebug** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetDebugAnalytics** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetDeleteTasksInterval** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetDisableGuiOneClickSupport** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetHardwareType** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetIbtrap** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetHwIdent** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetLines** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetMsMaxConnection** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetNosafemode** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetOcsp** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetPurgeRestartObjects** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetReportingUserCapabilities** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetRpzRecursiveOnly** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetSafemode** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetScheduled** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetSnmptrap** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetSysname** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetTerm** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetThresholdtrap** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetExpertmode** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetMaintenancemode** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetTransferReportingData** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetTransferSupportbundle** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetAnalyticsDatabaseDump** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetAnalyticsParameter** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetCollectOldLogs** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetCoreFilesQuota** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetHsmGroup** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetWred** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetEnableDohKeyLogging** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetEnableDotKeyLogging** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetHotfix** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetMgm** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetNtpStratum** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetPcDomain** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**SetReportFrequency** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**EnableAll** | Pointer to **bool** | If True then enable all fields | [optional] 
**DisableAll** | Pointer to **bool** | If True then disable all fields | [optional] 

## Methods

### NewAdmingroupAdminSetCommands

`func NewAdmingroupAdminSetCommands() *AdmingroupAdminSetCommands`

NewAdmingroupAdminSetCommands instantiates a new AdmingroupAdminSetCommands object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdmingroupAdminSetCommandsWithDefaults

`func NewAdmingroupAdminSetCommandsWithDefaults() *AdmingroupAdminSetCommands`

NewAdmingroupAdminSetCommandsWithDefaults instantiates a new AdmingroupAdminSetCommands object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSetAdminGroupAcl

`func (o *AdmingroupAdminSetCommands) GetSetAdminGroupAcl() bool`

GetSetAdminGroupAcl returns the SetAdminGroupAcl field if non-nil, zero value otherwise.

### GetSetAdminGroupAclOk

`func (o *AdmingroupAdminSetCommands) GetSetAdminGroupAclOk() (*bool, bool)`

GetSetAdminGroupAclOk returns a tuple with the SetAdminGroupAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetAdminGroupAcl

`func (o *AdmingroupAdminSetCommands) SetSetAdminGroupAcl(v bool)`

SetSetAdminGroupAcl sets SetAdminGroupAcl field to given value.

### HasSetAdminGroupAcl

`func (o *AdmingroupAdminSetCommands) HasSetAdminGroupAcl() bool`

HasSetAdminGroupAcl returns a boolean if a field has been set.

### GetEtBfd

`func (o *AdmingroupAdminSetCommands) GetEtBfd() bool`

GetEtBfd returns the EtBfd field if non-nil, zero value otherwise.

### GetEtBfdOk

`func (o *AdmingroupAdminSetCommands) GetEtBfdOk() (*bool, bool)`

GetEtBfdOk returns a tuple with the EtBfd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtBfd

`func (o *AdmingroupAdminSetCommands) SetEtBfd(v bool)`

SetEtBfd sets EtBfd field to given value.

### HasEtBfd

`func (o *AdmingroupAdminSetCommands) HasEtBfd() bool`

HasEtBfd returns a boolean if a field has been set.

### GetSetBfd

`func (o *AdmingroupAdminSetCommands) GetSetBfd() bool`

GetSetBfd returns the SetBfd field if non-nil, zero value otherwise.

### GetSetBfdOk

`func (o *AdmingroupAdminSetCommands) GetSetBfdOk() (*bool, bool)`

GetSetBfdOk returns a tuple with the SetBfd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetBfd

`func (o *AdmingroupAdminSetCommands) SetSetBfd(v bool)`

SetSetBfd sets SetBfd field to given value.

### HasSetBfd

`func (o *AdmingroupAdminSetCommands) HasSetBfd() bool`

HasSetBfd returns a boolean if a field has been set.

### GetSetBgp

`func (o *AdmingroupAdminSetCommands) GetSetBgp() bool`

GetSetBgp returns the SetBgp field if non-nil, zero value otherwise.

### GetSetBgpOk

`func (o *AdmingroupAdminSetCommands) GetSetBgpOk() (*bool, bool)`

GetSetBgpOk returns a tuple with the SetBgp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetBgp

`func (o *AdmingroupAdminSetCommands) SetSetBgp(v bool)`

SetSetBgp sets SetBgp field to given value.

### HasSetBgp

`func (o *AdmingroupAdminSetCommands) HasSetBgp() bool`

HasSetBgp returns a boolean if a field has been set.

### GetSetCleanMscache

`func (o *AdmingroupAdminSetCommands) GetSetCleanMscache() bool`

GetSetCleanMscache returns the SetCleanMscache field if non-nil, zero value otherwise.

### GetSetCleanMscacheOk

`func (o *AdmingroupAdminSetCommands) GetSetCleanMscacheOk() (*bool, bool)`

GetSetCleanMscacheOk returns a tuple with the SetCleanMscache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetCleanMscache

`func (o *AdmingroupAdminSetCommands) SetSetCleanMscache(v bool)`

SetSetCleanMscache sets SetCleanMscache field to given value.

### HasSetCleanMscache

`func (o *AdmingroupAdminSetCommands) HasSetCleanMscache() bool`

HasSetCleanMscache returns a boolean if a field has been set.

### GetSetDebug

`func (o *AdmingroupAdminSetCommands) GetSetDebug() bool`

GetSetDebug returns the SetDebug field if non-nil, zero value otherwise.

### GetSetDebugOk

`func (o *AdmingroupAdminSetCommands) GetSetDebugOk() (*bool, bool)`

GetSetDebugOk returns a tuple with the SetDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDebug

`func (o *AdmingroupAdminSetCommands) SetSetDebug(v bool)`

SetSetDebug sets SetDebug field to given value.

### HasSetDebug

`func (o *AdmingroupAdminSetCommands) HasSetDebug() bool`

HasSetDebug returns a boolean if a field has been set.

### GetSetDebugAnalytics

`func (o *AdmingroupAdminSetCommands) GetSetDebugAnalytics() bool`

GetSetDebugAnalytics returns the SetDebugAnalytics field if non-nil, zero value otherwise.

### GetSetDebugAnalyticsOk

`func (o *AdmingroupAdminSetCommands) GetSetDebugAnalyticsOk() (*bool, bool)`

GetSetDebugAnalyticsOk returns a tuple with the SetDebugAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDebugAnalytics

`func (o *AdmingroupAdminSetCommands) SetSetDebugAnalytics(v bool)`

SetSetDebugAnalytics sets SetDebugAnalytics field to given value.

### HasSetDebugAnalytics

`func (o *AdmingroupAdminSetCommands) HasSetDebugAnalytics() bool`

HasSetDebugAnalytics returns a boolean if a field has been set.

### GetSetDeleteTasksInterval

`func (o *AdmingroupAdminSetCommands) GetSetDeleteTasksInterval() bool`

GetSetDeleteTasksInterval returns the SetDeleteTasksInterval field if non-nil, zero value otherwise.

### GetSetDeleteTasksIntervalOk

`func (o *AdmingroupAdminSetCommands) GetSetDeleteTasksIntervalOk() (*bool, bool)`

GetSetDeleteTasksIntervalOk returns a tuple with the SetDeleteTasksInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDeleteTasksInterval

`func (o *AdmingroupAdminSetCommands) SetSetDeleteTasksInterval(v bool)`

SetSetDeleteTasksInterval sets SetDeleteTasksInterval field to given value.

### HasSetDeleteTasksInterval

`func (o *AdmingroupAdminSetCommands) HasSetDeleteTasksInterval() bool`

HasSetDeleteTasksInterval returns a boolean if a field has been set.

### GetSetDisableGuiOneClickSupport

`func (o *AdmingroupAdminSetCommands) GetSetDisableGuiOneClickSupport() bool`

GetSetDisableGuiOneClickSupport returns the SetDisableGuiOneClickSupport field if non-nil, zero value otherwise.

### GetSetDisableGuiOneClickSupportOk

`func (o *AdmingroupAdminSetCommands) GetSetDisableGuiOneClickSupportOk() (*bool, bool)`

GetSetDisableGuiOneClickSupportOk returns a tuple with the SetDisableGuiOneClickSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetDisableGuiOneClickSupport

`func (o *AdmingroupAdminSetCommands) SetSetDisableGuiOneClickSupport(v bool)`

SetSetDisableGuiOneClickSupport sets SetDisableGuiOneClickSupport field to given value.

### HasSetDisableGuiOneClickSupport

`func (o *AdmingroupAdminSetCommands) HasSetDisableGuiOneClickSupport() bool`

HasSetDisableGuiOneClickSupport returns a boolean if a field has been set.

### GetSetHardwareType

`func (o *AdmingroupAdminSetCommands) GetSetHardwareType() bool`

GetSetHardwareType returns the SetHardwareType field if non-nil, zero value otherwise.

### GetSetHardwareTypeOk

`func (o *AdmingroupAdminSetCommands) GetSetHardwareTypeOk() (*bool, bool)`

GetSetHardwareTypeOk returns a tuple with the SetHardwareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetHardwareType

`func (o *AdmingroupAdminSetCommands) SetSetHardwareType(v bool)`

SetSetHardwareType sets SetHardwareType field to given value.

### HasSetHardwareType

`func (o *AdmingroupAdminSetCommands) HasSetHardwareType() bool`

HasSetHardwareType returns a boolean if a field has been set.

### GetSetIbtrap

`func (o *AdmingroupAdminSetCommands) GetSetIbtrap() bool`

GetSetIbtrap returns the SetIbtrap field if non-nil, zero value otherwise.

### GetSetIbtrapOk

`func (o *AdmingroupAdminSetCommands) GetSetIbtrapOk() (*bool, bool)`

GetSetIbtrapOk returns a tuple with the SetIbtrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetIbtrap

`func (o *AdmingroupAdminSetCommands) SetSetIbtrap(v bool)`

SetSetIbtrap sets SetIbtrap field to given value.

### HasSetIbtrap

`func (o *AdmingroupAdminSetCommands) HasSetIbtrap() bool`

HasSetIbtrap returns a boolean if a field has been set.

### GetSetHwIdent

`func (o *AdmingroupAdminSetCommands) GetSetHwIdent() bool`

GetSetHwIdent returns the SetHwIdent field if non-nil, zero value otherwise.

### GetSetHwIdentOk

`func (o *AdmingroupAdminSetCommands) GetSetHwIdentOk() (*bool, bool)`

GetSetHwIdentOk returns a tuple with the SetHwIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetHwIdent

`func (o *AdmingroupAdminSetCommands) SetSetHwIdent(v bool)`

SetSetHwIdent sets SetHwIdent field to given value.

### HasSetHwIdent

`func (o *AdmingroupAdminSetCommands) HasSetHwIdent() bool`

HasSetHwIdent returns a boolean if a field has been set.

### GetSetLines

`func (o *AdmingroupAdminSetCommands) GetSetLines() bool`

GetSetLines returns the SetLines field if non-nil, zero value otherwise.

### GetSetLinesOk

`func (o *AdmingroupAdminSetCommands) GetSetLinesOk() (*bool, bool)`

GetSetLinesOk returns a tuple with the SetLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetLines

`func (o *AdmingroupAdminSetCommands) SetSetLines(v bool)`

SetSetLines sets SetLines field to given value.

### HasSetLines

`func (o *AdmingroupAdminSetCommands) HasSetLines() bool`

HasSetLines returns a boolean if a field has been set.

### GetSetMsMaxConnection

`func (o *AdmingroupAdminSetCommands) GetSetMsMaxConnection() bool`

GetSetMsMaxConnection returns the SetMsMaxConnection field if non-nil, zero value otherwise.

### GetSetMsMaxConnectionOk

`func (o *AdmingroupAdminSetCommands) GetSetMsMaxConnectionOk() (*bool, bool)`

GetSetMsMaxConnectionOk returns a tuple with the SetMsMaxConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetMsMaxConnection

`func (o *AdmingroupAdminSetCommands) SetSetMsMaxConnection(v bool)`

SetSetMsMaxConnection sets SetMsMaxConnection field to given value.

### HasSetMsMaxConnection

`func (o *AdmingroupAdminSetCommands) HasSetMsMaxConnection() bool`

HasSetMsMaxConnection returns a boolean if a field has been set.

### GetSetNosafemode

`func (o *AdmingroupAdminSetCommands) GetSetNosafemode() bool`

GetSetNosafemode returns the SetNosafemode field if non-nil, zero value otherwise.

### GetSetNosafemodeOk

`func (o *AdmingroupAdminSetCommands) GetSetNosafemodeOk() (*bool, bool)`

GetSetNosafemodeOk returns a tuple with the SetNosafemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetNosafemode

`func (o *AdmingroupAdminSetCommands) SetSetNosafemode(v bool)`

SetSetNosafemode sets SetNosafemode field to given value.

### HasSetNosafemode

`func (o *AdmingroupAdminSetCommands) HasSetNosafemode() bool`

HasSetNosafemode returns a boolean if a field has been set.

### GetSetOcsp

`func (o *AdmingroupAdminSetCommands) GetSetOcsp() bool`

GetSetOcsp returns the SetOcsp field if non-nil, zero value otherwise.

### GetSetOcspOk

`func (o *AdmingroupAdminSetCommands) GetSetOcspOk() (*bool, bool)`

GetSetOcspOk returns a tuple with the SetOcsp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetOcsp

`func (o *AdmingroupAdminSetCommands) SetSetOcsp(v bool)`

SetSetOcsp sets SetOcsp field to given value.

### HasSetOcsp

`func (o *AdmingroupAdminSetCommands) HasSetOcsp() bool`

HasSetOcsp returns a boolean if a field has been set.

### GetSetPurgeRestartObjects

`func (o *AdmingroupAdminSetCommands) GetSetPurgeRestartObjects() bool`

GetSetPurgeRestartObjects returns the SetPurgeRestartObjects field if non-nil, zero value otherwise.

### GetSetPurgeRestartObjectsOk

`func (o *AdmingroupAdminSetCommands) GetSetPurgeRestartObjectsOk() (*bool, bool)`

GetSetPurgeRestartObjectsOk returns a tuple with the SetPurgeRestartObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetPurgeRestartObjects

`func (o *AdmingroupAdminSetCommands) SetSetPurgeRestartObjects(v bool)`

SetSetPurgeRestartObjects sets SetPurgeRestartObjects field to given value.

### HasSetPurgeRestartObjects

`func (o *AdmingroupAdminSetCommands) HasSetPurgeRestartObjects() bool`

HasSetPurgeRestartObjects returns a boolean if a field has been set.

### GetSetReportingUserCapabilities

`func (o *AdmingroupAdminSetCommands) GetSetReportingUserCapabilities() bool`

GetSetReportingUserCapabilities returns the SetReportingUserCapabilities field if non-nil, zero value otherwise.

### GetSetReportingUserCapabilitiesOk

`func (o *AdmingroupAdminSetCommands) GetSetReportingUserCapabilitiesOk() (*bool, bool)`

GetSetReportingUserCapabilitiesOk returns a tuple with the SetReportingUserCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetReportingUserCapabilities

`func (o *AdmingroupAdminSetCommands) SetSetReportingUserCapabilities(v bool)`

SetSetReportingUserCapabilities sets SetReportingUserCapabilities field to given value.

### HasSetReportingUserCapabilities

`func (o *AdmingroupAdminSetCommands) HasSetReportingUserCapabilities() bool`

HasSetReportingUserCapabilities returns a boolean if a field has been set.

### GetSetRpzRecursiveOnly

`func (o *AdmingroupAdminSetCommands) GetSetRpzRecursiveOnly() bool`

GetSetRpzRecursiveOnly returns the SetRpzRecursiveOnly field if non-nil, zero value otherwise.

### GetSetRpzRecursiveOnlyOk

`func (o *AdmingroupAdminSetCommands) GetSetRpzRecursiveOnlyOk() (*bool, bool)`

GetSetRpzRecursiveOnlyOk returns a tuple with the SetRpzRecursiveOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetRpzRecursiveOnly

`func (o *AdmingroupAdminSetCommands) SetSetRpzRecursiveOnly(v bool)`

SetSetRpzRecursiveOnly sets SetRpzRecursiveOnly field to given value.

### HasSetRpzRecursiveOnly

`func (o *AdmingroupAdminSetCommands) HasSetRpzRecursiveOnly() bool`

HasSetRpzRecursiveOnly returns a boolean if a field has been set.

### GetSetSafemode

`func (o *AdmingroupAdminSetCommands) GetSetSafemode() bool`

GetSetSafemode returns the SetSafemode field if non-nil, zero value otherwise.

### GetSetSafemodeOk

`func (o *AdmingroupAdminSetCommands) GetSetSafemodeOk() (*bool, bool)`

GetSetSafemodeOk returns a tuple with the SetSafemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetSafemode

`func (o *AdmingroupAdminSetCommands) SetSetSafemode(v bool)`

SetSetSafemode sets SetSafemode field to given value.

### HasSetSafemode

`func (o *AdmingroupAdminSetCommands) HasSetSafemode() bool`

HasSetSafemode returns a boolean if a field has been set.

### GetSetScheduled

`func (o *AdmingroupAdminSetCommands) GetSetScheduled() bool`

GetSetScheduled returns the SetScheduled field if non-nil, zero value otherwise.

### GetSetScheduledOk

`func (o *AdmingroupAdminSetCommands) GetSetScheduledOk() (*bool, bool)`

GetSetScheduledOk returns a tuple with the SetScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetScheduled

`func (o *AdmingroupAdminSetCommands) SetSetScheduled(v bool)`

SetSetScheduled sets SetScheduled field to given value.

### HasSetScheduled

`func (o *AdmingroupAdminSetCommands) HasSetScheduled() bool`

HasSetScheduled returns a boolean if a field has been set.

### GetSetSnmptrap

`func (o *AdmingroupAdminSetCommands) GetSetSnmptrap() bool`

GetSetSnmptrap returns the SetSnmptrap field if non-nil, zero value otherwise.

### GetSetSnmptrapOk

`func (o *AdmingroupAdminSetCommands) GetSetSnmptrapOk() (*bool, bool)`

GetSetSnmptrapOk returns a tuple with the SetSnmptrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetSnmptrap

`func (o *AdmingroupAdminSetCommands) SetSetSnmptrap(v bool)`

SetSetSnmptrap sets SetSnmptrap field to given value.

### HasSetSnmptrap

`func (o *AdmingroupAdminSetCommands) HasSetSnmptrap() bool`

HasSetSnmptrap returns a boolean if a field has been set.

### GetSetSysname

`func (o *AdmingroupAdminSetCommands) GetSetSysname() bool`

GetSetSysname returns the SetSysname field if non-nil, zero value otherwise.

### GetSetSysnameOk

`func (o *AdmingroupAdminSetCommands) GetSetSysnameOk() (*bool, bool)`

GetSetSysnameOk returns a tuple with the SetSysname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetSysname

`func (o *AdmingroupAdminSetCommands) SetSetSysname(v bool)`

SetSetSysname sets SetSysname field to given value.

### HasSetSysname

`func (o *AdmingroupAdminSetCommands) HasSetSysname() bool`

HasSetSysname returns a boolean if a field has been set.

### GetSetTerm

`func (o *AdmingroupAdminSetCommands) GetSetTerm() bool`

GetSetTerm returns the SetTerm field if non-nil, zero value otherwise.

### GetSetTermOk

`func (o *AdmingroupAdminSetCommands) GetSetTermOk() (*bool, bool)`

GetSetTermOk returns a tuple with the SetTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetTerm

`func (o *AdmingroupAdminSetCommands) SetSetTerm(v bool)`

SetSetTerm sets SetTerm field to given value.

### HasSetTerm

`func (o *AdmingroupAdminSetCommands) HasSetTerm() bool`

HasSetTerm returns a boolean if a field has been set.

### GetSetThresholdtrap

`func (o *AdmingroupAdminSetCommands) GetSetThresholdtrap() bool`

GetSetThresholdtrap returns the SetThresholdtrap field if non-nil, zero value otherwise.

### GetSetThresholdtrapOk

`func (o *AdmingroupAdminSetCommands) GetSetThresholdtrapOk() (*bool, bool)`

GetSetThresholdtrapOk returns a tuple with the SetThresholdtrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetThresholdtrap

`func (o *AdmingroupAdminSetCommands) SetSetThresholdtrap(v bool)`

SetSetThresholdtrap sets SetThresholdtrap field to given value.

### HasSetThresholdtrap

`func (o *AdmingroupAdminSetCommands) HasSetThresholdtrap() bool`

HasSetThresholdtrap returns a boolean if a field has been set.

### GetSetExpertmode

`func (o *AdmingroupAdminSetCommands) GetSetExpertmode() bool`

GetSetExpertmode returns the SetExpertmode field if non-nil, zero value otherwise.

### GetSetExpertmodeOk

`func (o *AdmingroupAdminSetCommands) GetSetExpertmodeOk() (*bool, bool)`

GetSetExpertmodeOk returns a tuple with the SetExpertmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetExpertmode

`func (o *AdmingroupAdminSetCommands) SetSetExpertmode(v bool)`

SetSetExpertmode sets SetExpertmode field to given value.

### HasSetExpertmode

`func (o *AdmingroupAdminSetCommands) HasSetExpertmode() bool`

HasSetExpertmode returns a boolean if a field has been set.

### GetSetMaintenancemode

`func (o *AdmingroupAdminSetCommands) GetSetMaintenancemode() bool`

GetSetMaintenancemode returns the SetMaintenancemode field if non-nil, zero value otherwise.

### GetSetMaintenancemodeOk

`func (o *AdmingroupAdminSetCommands) GetSetMaintenancemodeOk() (*bool, bool)`

GetSetMaintenancemodeOk returns a tuple with the SetMaintenancemode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetMaintenancemode

`func (o *AdmingroupAdminSetCommands) SetSetMaintenancemode(v bool)`

SetSetMaintenancemode sets SetMaintenancemode field to given value.

### HasSetMaintenancemode

`func (o *AdmingroupAdminSetCommands) HasSetMaintenancemode() bool`

HasSetMaintenancemode returns a boolean if a field has been set.

### GetSetTransferReportingData

`func (o *AdmingroupAdminSetCommands) GetSetTransferReportingData() bool`

GetSetTransferReportingData returns the SetTransferReportingData field if non-nil, zero value otherwise.

### GetSetTransferReportingDataOk

`func (o *AdmingroupAdminSetCommands) GetSetTransferReportingDataOk() (*bool, bool)`

GetSetTransferReportingDataOk returns a tuple with the SetTransferReportingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetTransferReportingData

`func (o *AdmingroupAdminSetCommands) SetSetTransferReportingData(v bool)`

SetSetTransferReportingData sets SetTransferReportingData field to given value.

### HasSetTransferReportingData

`func (o *AdmingroupAdminSetCommands) HasSetTransferReportingData() bool`

HasSetTransferReportingData returns a boolean if a field has been set.

### GetSetTransferSupportbundle

`func (o *AdmingroupAdminSetCommands) GetSetTransferSupportbundle() bool`

GetSetTransferSupportbundle returns the SetTransferSupportbundle field if non-nil, zero value otherwise.

### GetSetTransferSupportbundleOk

`func (o *AdmingroupAdminSetCommands) GetSetTransferSupportbundleOk() (*bool, bool)`

GetSetTransferSupportbundleOk returns a tuple with the SetTransferSupportbundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetTransferSupportbundle

`func (o *AdmingroupAdminSetCommands) SetSetTransferSupportbundle(v bool)`

SetSetTransferSupportbundle sets SetTransferSupportbundle field to given value.

### HasSetTransferSupportbundle

`func (o *AdmingroupAdminSetCommands) HasSetTransferSupportbundle() bool`

HasSetTransferSupportbundle returns a boolean if a field has been set.

### GetSetAnalyticsDatabaseDump

`func (o *AdmingroupAdminSetCommands) GetSetAnalyticsDatabaseDump() bool`

GetSetAnalyticsDatabaseDump returns the SetAnalyticsDatabaseDump field if non-nil, zero value otherwise.

### GetSetAnalyticsDatabaseDumpOk

`func (o *AdmingroupAdminSetCommands) GetSetAnalyticsDatabaseDumpOk() (*bool, bool)`

GetSetAnalyticsDatabaseDumpOk returns a tuple with the SetAnalyticsDatabaseDump field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetAnalyticsDatabaseDump

`func (o *AdmingroupAdminSetCommands) SetSetAnalyticsDatabaseDump(v bool)`

SetSetAnalyticsDatabaseDump sets SetAnalyticsDatabaseDump field to given value.

### HasSetAnalyticsDatabaseDump

`func (o *AdmingroupAdminSetCommands) HasSetAnalyticsDatabaseDump() bool`

HasSetAnalyticsDatabaseDump returns a boolean if a field has been set.

### GetSetAnalyticsParameter

`func (o *AdmingroupAdminSetCommands) GetSetAnalyticsParameter() bool`

GetSetAnalyticsParameter returns the SetAnalyticsParameter field if non-nil, zero value otherwise.

### GetSetAnalyticsParameterOk

`func (o *AdmingroupAdminSetCommands) GetSetAnalyticsParameterOk() (*bool, bool)`

GetSetAnalyticsParameterOk returns a tuple with the SetAnalyticsParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetAnalyticsParameter

`func (o *AdmingroupAdminSetCommands) SetSetAnalyticsParameter(v bool)`

SetSetAnalyticsParameter sets SetAnalyticsParameter field to given value.

### HasSetAnalyticsParameter

`func (o *AdmingroupAdminSetCommands) HasSetAnalyticsParameter() bool`

HasSetAnalyticsParameter returns a boolean if a field has been set.

### GetSetCollectOldLogs

`func (o *AdmingroupAdminSetCommands) GetSetCollectOldLogs() bool`

GetSetCollectOldLogs returns the SetCollectOldLogs field if non-nil, zero value otherwise.

### GetSetCollectOldLogsOk

`func (o *AdmingroupAdminSetCommands) GetSetCollectOldLogsOk() (*bool, bool)`

GetSetCollectOldLogsOk returns a tuple with the SetCollectOldLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetCollectOldLogs

`func (o *AdmingroupAdminSetCommands) SetSetCollectOldLogs(v bool)`

SetSetCollectOldLogs sets SetCollectOldLogs field to given value.

### HasSetCollectOldLogs

`func (o *AdmingroupAdminSetCommands) HasSetCollectOldLogs() bool`

HasSetCollectOldLogs returns a boolean if a field has been set.

### GetSetCoreFilesQuota

`func (o *AdmingroupAdminSetCommands) GetSetCoreFilesQuota() bool`

GetSetCoreFilesQuota returns the SetCoreFilesQuota field if non-nil, zero value otherwise.

### GetSetCoreFilesQuotaOk

`func (o *AdmingroupAdminSetCommands) GetSetCoreFilesQuotaOk() (*bool, bool)`

GetSetCoreFilesQuotaOk returns a tuple with the SetCoreFilesQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetCoreFilesQuota

`func (o *AdmingroupAdminSetCommands) SetSetCoreFilesQuota(v bool)`

SetSetCoreFilesQuota sets SetCoreFilesQuota field to given value.

### HasSetCoreFilesQuota

`func (o *AdmingroupAdminSetCommands) HasSetCoreFilesQuota() bool`

HasSetCoreFilesQuota returns a boolean if a field has been set.

### GetSetHsmGroup

`func (o *AdmingroupAdminSetCommands) GetSetHsmGroup() bool`

GetSetHsmGroup returns the SetHsmGroup field if non-nil, zero value otherwise.

### GetSetHsmGroupOk

`func (o *AdmingroupAdminSetCommands) GetSetHsmGroupOk() (*bool, bool)`

GetSetHsmGroupOk returns a tuple with the SetHsmGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetHsmGroup

`func (o *AdmingroupAdminSetCommands) SetSetHsmGroup(v bool)`

SetSetHsmGroup sets SetHsmGroup field to given value.

### HasSetHsmGroup

`func (o *AdmingroupAdminSetCommands) HasSetHsmGroup() bool`

HasSetHsmGroup returns a boolean if a field has been set.

### GetSetWred

`func (o *AdmingroupAdminSetCommands) GetSetWred() bool`

GetSetWred returns the SetWred field if non-nil, zero value otherwise.

### GetSetWredOk

`func (o *AdmingroupAdminSetCommands) GetSetWredOk() (*bool, bool)`

GetSetWredOk returns a tuple with the SetWred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetWred

`func (o *AdmingroupAdminSetCommands) SetSetWred(v bool)`

SetSetWred sets SetWred field to given value.

### HasSetWred

`func (o *AdmingroupAdminSetCommands) HasSetWred() bool`

HasSetWred returns a boolean if a field has been set.

### GetSetEnableDohKeyLogging

`func (o *AdmingroupAdminSetCommands) GetSetEnableDohKeyLogging() bool`

GetSetEnableDohKeyLogging returns the SetEnableDohKeyLogging field if non-nil, zero value otherwise.

### GetSetEnableDohKeyLoggingOk

`func (o *AdmingroupAdminSetCommands) GetSetEnableDohKeyLoggingOk() (*bool, bool)`

GetSetEnableDohKeyLoggingOk returns a tuple with the SetEnableDohKeyLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetEnableDohKeyLogging

`func (o *AdmingroupAdminSetCommands) SetSetEnableDohKeyLogging(v bool)`

SetSetEnableDohKeyLogging sets SetEnableDohKeyLogging field to given value.

### HasSetEnableDohKeyLogging

`func (o *AdmingroupAdminSetCommands) HasSetEnableDohKeyLogging() bool`

HasSetEnableDohKeyLogging returns a boolean if a field has been set.

### GetSetEnableDotKeyLogging

`func (o *AdmingroupAdminSetCommands) GetSetEnableDotKeyLogging() bool`

GetSetEnableDotKeyLogging returns the SetEnableDotKeyLogging field if non-nil, zero value otherwise.

### GetSetEnableDotKeyLoggingOk

`func (o *AdmingroupAdminSetCommands) GetSetEnableDotKeyLoggingOk() (*bool, bool)`

GetSetEnableDotKeyLoggingOk returns a tuple with the SetEnableDotKeyLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetEnableDotKeyLogging

`func (o *AdmingroupAdminSetCommands) SetSetEnableDotKeyLogging(v bool)`

SetSetEnableDotKeyLogging sets SetEnableDotKeyLogging field to given value.

### HasSetEnableDotKeyLogging

`func (o *AdmingroupAdminSetCommands) HasSetEnableDotKeyLogging() bool`

HasSetEnableDotKeyLogging returns a boolean if a field has been set.

### GetSetHotfix

`func (o *AdmingroupAdminSetCommands) GetSetHotfix() bool`

GetSetHotfix returns the SetHotfix field if non-nil, zero value otherwise.

### GetSetHotfixOk

`func (o *AdmingroupAdminSetCommands) GetSetHotfixOk() (*bool, bool)`

GetSetHotfixOk returns a tuple with the SetHotfix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetHotfix

`func (o *AdmingroupAdminSetCommands) SetSetHotfix(v bool)`

SetSetHotfix sets SetHotfix field to given value.

### HasSetHotfix

`func (o *AdmingroupAdminSetCommands) HasSetHotfix() bool`

HasSetHotfix returns a boolean if a field has been set.

### GetSetMgm

`func (o *AdmingroupAdminSetCommands) GetSetMgm() bool`

GetSetMgm returns the SetMgm field if non-nil, zero value otherwise.

### GetSetMgmOk

`func (o *AdmingroupAdminSetCommands) GetSetMgmOk() (*bool, bool)`

GetSetMgmOk returns a tuple with the SetMgm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetMgm

`func (o *AdmingroupAdminSetCommands) SetSetMgm(v bool)`

SetSetMgm sets SetMgm field to given value.

### HasSetMgm

`func (o *AdmingroupAdminSetCommands) HasSetMgm() bool`

HasSetMgm returns a boolean if a field has been set.

### GetSetNtpStratum

`func (o *AdmingroupAdminSetCommands) GetSetNtpStratum() bool`

GetSetNtpStratum returns the SetNtpStratum field if non-nil, zero value otherwise.

### GetSetNtpStratumOk

`func (o *AdmingroupAdminSetCommands) GetSetNtpStratumOk() (*bool, bool)`

GetSetNtpStratumOk returns a tuple with the SetNtpStratum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetNtpStratum

`func (o *AdmingroupAdminSetCommands) SetSetNtpStratum(v bool)`

SetSetNtpStratum sets SetNtpStratum field to given value.

### HasSetNtpStratum

`func (o *AdmingroupAdminSetCommands) HasSetNtpStratum() bool`

HasSetNtpStratum returns a boolean if a field has been set.

### GetSetPcDomain

`func (o *AdmingroupAdminSetCommands) GetSetPcDomain() bool`

GetSetPcDomain returns the SetPcDomain field if non-nil, zero value otherwise.

### GetSetPcDomainOk

`func (o *AdmingroupAdminSetCommands) GetSetPcDomainOk() (*bool, bool)`

GetSetPcDomainOk returns a tuple with the SetPcDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetPcDomain

`func (o *AdmingroupAdminSetCommands) SetSetPcDomain(v bool)`

SetSetPcDomain sets SetPcDomain field to given value.

### HasSetPcDomain

`func (o *AdmingroupAdminSetCommands) HasSetPcDomain() bool`

HasSetPcDomain returns a boolean if a field has been set.

### GetSetReportFrequency

`func (o *AdmingroupAdminSetCommands) GetSetReportFrequency() bool`

GetSetReportFrequency returns the SetReportFrequency field if non-nil, zero value otherwise.

### GetSetReportFrequencyOk

`func (o *AdmingroupAdminSetCommands) GetSetReportFrequencyOk() (*bool, bool)`

GetSetReportFrequencyOk returns a tuple with the SetReportFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetReportFrequency

`func (o *AdmingroupAdminSetCommands) SetSetReportFrequency(v bool)`

SetSetReportFrequency sets SetReportFrequency field to given value.

### HasSetReportFrequency

`func (o *AdmingroupAdminSetCommands) HasSetReportFrequency() bool`

HasSetReportFrequency returns a boolean if a field has been set.

### GetEnableAll

`func (o *AdmingroupAdminSetCommands) GetEnableAll() bool`

GetEnableAll returns the EnableAll field if non-nil, zero value otherwise.

### GetEnableAllOk

`func (o *AdmingroupAdminSetCommands) GetEnableAllOk() (*bool, bool)`

GetEnableAllOk returns a tuple with the EnableAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAll

`func (o *AdmingroupAdminSetCommands) SetEnableAll(v bool)`

SetEnableAll sets EnableAll field to given value.

### HasEnableAll

`func (o *AdmingroupAdminSetCommands) HasEnableAll() bool`

HasEnableAll returns a boolean if a field has been set.

### GetDisableAll

`func (o *AdmingroupAdminSetCommands) GetDisableAll() bool`

GetDisableAll returns the DisableAll field if non-nil, zero value otherwise.

### GetDisableAllOk

`func (o *AdmingroupAdminSetCommands) GetDisableAllOk() (*bool, bool)`

GetDisableAllOk returns a tuple with the DisableAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAll

`func (o *AdmingroupAdminSetCommands) SetDisableAll(v bool)`

SetDisableAll sets DisableAll field to given value.

### HasDisableAll

`func (o *AdmingroupAdminSetCommands) HasDisableAll() bool`

HasDisableAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


