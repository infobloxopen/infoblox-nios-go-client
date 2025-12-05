# AdmingroupAdminShowCommands

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShowAdminGroupAcl** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowAnalyticsParameter** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowArp** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowBfd** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowBgp** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowCapacity** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowClusterdInfo** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowConfig** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowCpu** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDate** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDebug** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDebugAnalytics** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDeleteTasksInterval** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDisk** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowHardwareType** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowHardwareStatus** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowHwid** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowIbtrap** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowHwIdent** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowLog** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowLogfiles** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowMemory** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowNtp** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowReportingUserCapabilities** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowRpzRecursiveOnly** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowScheduled** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowSnmp** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowStatus** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowTechSupport** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowTemperature** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowThresholdtrap** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowUpgradeHistory** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowUptime** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowVersion** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowAnalyticsDatabaseDumps** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowCores** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowCoresummary** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowCspThreatDb** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowHsmGroup** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowHsmInfo** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowPmap** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowProcess** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowPstack** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowSafenetSupportInfo** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowWredStats** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowWredStatus** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowNtpStratum** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowPcDomain** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowReportFrequency** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowQueryLoggingWarnings** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**EnableAll** | Pointer to **bool** | If True then enable all fields | [optional] 
**DisableAll** | Pointer to **bool** | If True then disable all fields | [optional] 

## Methods

### NewAdmingroupAdminShowCommands

`func NewAdmingroupAdminShowCommands() *AdmingroupAdminShowCommands`

NewAdmingroupAdminShowCommands instantiates a new AdmingroupAdminShowCommands object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdmingroupAdminShowCommandsWithDefaults

`func NewAdmingroupAdminShowCommandsWithDefaults() *AdmingroupAdminShowCommands`

NewAdmingroupAdminShowCommandsWithDefaults instantiates a new AdmingroupAdminShowCommands object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowAdminGroupAcl

`func (o *AdmingroupAdminShowCommands) GetShowAdminGroupAcl() bool`

GetShowAdminGroupAcl returns the ShowAdminGroupAcl field if non-nil, zero value otherwise.

### GetShowAdminGroupAclOk

`func (o *AdmingroupAdminShowCommands) GetShowAdminGroupAclOk() (*bool, bool)`

GetShowAdminGroupAclOk returns a tuple with the ShowAdminGroupAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAdminGroupAcl

`func (o *AdmingroupAdminShowCommands) SetShowAdminGroupAcl(v bool)`

SetShowAdminGroupAcl sets ShowAdminGroupAcl field to given value.

### HasShowAdminGroupAcl

`func (o *AdmingroupAdminShowCommands) HasShowAdminGroupAcl() bool`

HasShowAdminGroupAcl returns a boolean if a field has been set.

### GetShowAnalyticsParameter

`func (o *AdmingroupAdminShowCommands) GetShowAnalyticsParameter() bool`

GetShowAnalyticsParameter returns the ShowAnalyticsParameter field if non-nil, zero value otherwise.

### GetShowAnalyticsParameterOk

`func (o *AdmingroupAdminShowCommands) GetShowAnalyticsParameterOk() (*bool, bool)`

GetShowAnalyticsParameterOk returns a tuple with the ShowAnalyticsParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAnalyticsParameter

`func (o *AdmingroupAdminShowCommands) SetShowAnalyticsParameter(v bool)`

SetShowAnalyticsParameter sets ShowAnalyticsParameter field to given value.

### HasShowAnalyticsParameter

`func (o *AdmingroupAdminShowCommands) HasShowAnalyticsParameter() bool`

HasShowAnalyticsParameter returns a boolean if a field has been set.

### GetShowArp

`func (o *AdmingroupAdminShowCommands) GetShowArp() bool`

GetShowArp returns the ShowArp field if non-nil, zero value otherwise.

### GetShowArpOk

`func (o *AdmingroupAdminShowCommands) GetShowArpOk() (*bool, bool)`

GetShowArpOk returns a tuple with the ShowArp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowArp

`func (o *AdmingroupAdminShowCommands) SetShowArp(v bool)`

SetShowArp sets ShowArp field to given value.

### HasShowArp

`func (o *AdmingroupAdminShowCommands) HasShowArp() bool`

HasShowArp returns a boolean if a field has been set.

### GetShowBfd

`func (o *AdmingroupAdminShowCommands) GetShowBfd() bool`

GetShowBfd returns the ShowBfd field if non-nil, zero value otherwise.

### GetShowBfdOk

`func (o *AdmingroupAdminShowCommands) GetShowBfdOk() (*bool, bool)`

GetShowBfdOk returns a tuple with the ShowBfd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowBfd

`func (o *AdmingroupAdminShowCommands) SetShowBfd(v bool)`

SetShowBfd sets ShowBfd field to given value.

### HasShowBfd

`func (o *AdmingroupAdminShowCommands) HasShowBfd() bool`

HasShowBfd returns a boolean if a field has been set.

### GetShowBgp

`func (o *AdmingroupAdminShowCommands) GetShowBgp() bool`

GetShowBgp returns the ShowBgp field if non-nil, zero value otherwise.

### GetShowBgpOk

`func (o *AdmingroupAdminShowCommands) GetShowBgpOk() (*bool, bool)`

GetShowBgpOk returns a tuple with the ShowBgp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowBgp

`func (o *AdmingroupAdminShowCommands) SetShowBgp(v bool)`

SetShowBgp sets ShowBgp field to given value.

### HasShowBgp

`func (o *AdmingroupAdminShowCommands) HasShowBgp() bool`

HasShowBgp returns a boolean if a field has been set.

### GetShowCapacity

`func (o *AdmingroupAdminShowCommands) GetShowCapacity() bool`

GetShowCapacity returns the ShowCapacity field if non-nil, zero value otherwise.

### GetShowCapacityOk

`func (o *AdmingroupAdminShowCommands) GetShowCapacityOk() (*bool, bool)`

GetShowCapacityOk returns a tuple with the ShowCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCapacity

`func (o *AdmingroupAdminShowCommands) SetShowCapacity(v bool)`

SetShowCapacity sets ShowCapacity field to given value.

### HasShowCapacity

`func (o *AdmingroupAdminShowCommands) HasShowCapacity() bool`

HasShowCapacity returns a boolean if a field has been set.

### GetShowClusterdInfo

`func (o *AdmingroupAdminShowCommands) GetShowClusterdInfo() bool`

GetShowClusterdInfo returns the ShowClusterdInfo field if non-nil, zero value otherwise.

### GetShowClusterdInfoOk

`func (o *AdmingroupAdminShowCommands) GetShowClusterdInfoOk() (*bool, bool)`

GetShowClusterdInfoOk returns a tuple with the ShowClusterdInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowClusterdInfo

`func (o *AdmingroupAdminShowCommands) SetShowClusterdInfo(v bool)`

SetShowClusterdInfo sets ShowClusterdInfo field to given value.

### HasShowClusterdInfo

`func (o *AdmingroupAdminShowCommands) HasShowClusterdInfo() bool`

HasShowClusterdInfo returns a boolean if a field has been set.

### GetShowConfig

`func (o *AdmingroupAdminShowCommands) GetShowConfig() bool`

GetShowConfig returns the ShowConfig field if non-nil, zero value otherwise.

### GetShowConfigOk

`func (o *AdmingroupAdminShowCommands) GetShowConfigOk() (*bool, bool)`

GetShowConfigOk returns a tuple with the ShowConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowConfig

`func (o *AdmingroupAdminShowCommands) SetShowConfig(v bool)`

SetShowConfig sets ShowConfig field to given value.

### HasShowConfig

`func (o *AdmingroupAdminShowCommands) HasShowConfig() bool`

HasShowConfig returns a boolean if a field has been set.

### GetShowCpu

`func (o *AdmingroupAdminShowCommands) GetShowCpu() bool`

GetShowCpu returns the ShowCpu field if non-nil, zero value otherwise.

### GetShowCpuOk

`func (o *AdmingroupAdminShowCommands) GetShowCpuOk() (*bool, bool)`

GetShowCpuOk returns a tuple with the ShowCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCpu

`func (o *AdmingroupAdminShowCommands) SetShowCpu(v bool)`

SetShowCpu sets ShowCpu field to given value.

### HasShowCpu

`func (o *AdmingroupAdminShowCommands) HasShowCpu() bool`

HasShowCpu returns a boolean if a field has been set.

### GetShowDate

`func (o *AdmingroupAdminShowCommands) GetShowDate() bool`

GetShowDate returns the ShowDate field if non-nil, zero value otherwise.

### GetShowDateOk

`func (o *AdmingroupAdminShowCommands) GetShowDateOk() (*bool, bool)`

GetShowDateOk returns a tuple with the ShowDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDate

`func (o *AdmingroupAdminShowCommands) SetShowDate(v bool)`

SetShowDate sets ShowDate field to given value.

### HasShowDate

`func (o *AdmingroupAdminShowCommands) HasShowDate() bool`

HasShowDate returns a boolean if a field has been set.

### GetShowDebug

`func (o *AdmingroupAdminShowCommands) GetShowDebug() bool`

GetShowDebug returns the ShowDebug field if non-nil, zero value otherwise.

### GetShowDebugOk

`func (o *AdmingroupAdminShowCommands) GetShowDebugOk() (*bool, bool)`

GetShowDebugOk returns a tuple with the ShowDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDebug

`func (o *AdmingroupAdminShowCommands) SetShowDebug(v bool)`

SetShowDebug sets ShowDebug field to given value.

### HasShowDebug

`func (o *AdmingroupAdminShowCommands) HasShowDebug() bool`

HasShowDebug returns a boolean if a field has been set.

### GetShowDebugAnalytics

`func (o *AdmingroupAdminShowCommands) GetShowDebugAnalytics() bool`

GetShowDebugAnalytics returns the ShowDebugAnalytics field if non-nil, zero value otherwise.

### GetShowDebugAnalyticsOk

`func (o *AdmingroupAdminShowCommands) GetShowDebugAnalyticsOk() (*bool, bool)`

GetShowDebugAnalyticsOk returns a tuple with the ShowDebugAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDebugAnalytics

`func (o *AdmingroupAdminShowCommands) SetShowDebugAnalytics(v bool)`

SetShowDebugAnalytics sets ShowDebugAnalytics field to given value.

### HasShowDebugAnalytics

`func (o *AdmingroupAdminShowCommands) HasShowDebugAnalytics() bool`

HasShowDebugAnalytics returns a boolean if a field has been set.

### GetShowDeleteTasksInterval

`func (o *AdmingroupAdminShowCommands) GetShowDeleteTasksInterval() bool`

GetShowDeleteTasksInterval returns the ShowDeleteTasksInterval field if non-nil, zero value otherwise.

### GetShowDeleteTasksIntervalOk

`func (o *AdmingroupAdminShowCommands) GetShowDeleteTasksIntervalOk() (*bool, bool)`

GetShowDeleteTasksIntervalOk returns a tuple with the ShowDeleteTasksInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDeleteTasksInterval

`func (o *AdmingroupAdminShowCommands) SetShowDeleteTasksInterval(v bool)`

SetShowDeleteTasksInterval sets ShowDeleteTasksInterval field to given value.

### HasShowDeleteTasksInterval

`func (o *AdmingroupAdminShowCommands) HasShowDeleteTasksInterval() bool`

HasShowDeleteTasksInterval returns a boolean if a field has been set.

### GetShowDisk

`func (o *AdmingroupAdminShowCommands) GetShowDisk() bool`

GetShowDisk returns the ShowDisk field if non-nil, zero value otherwise.

### GetShowDiskOk

`func (o *AdmingroupAdminShowCommands) GetShowDiskOk() (*bool, bool)`

GetShowDiskOk returns a tuple with the ShowDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDisk

`func (o *AdmingroupAdminShowCommands) SetShowDisk(v bool)`

SetShowDisk sets ShowDisk field to given value.

### HasShowDisk

`func (o *AdmingroupAdminShowCommands) HasShowDisk() bool`

HasShowDisk returns a boolean if a field has been set.

### GetShowHardwareType

`func (o *AdmingroupAdminShowCommands) GetShowHardwareType() bool`

GetShowHardwareType returns the ShowHardwareType field if non-nil, zero value otherwise.

### GetShowHardwareTypeOk

`func (o *AdmingroupAdminShowCommands) GetShowHardwareTypeOk() (*bool, bool)`

GetShowHardwareTypeOk returns a tuple with the ShowHardwareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHardwareType

`func (o *AdmingroupAdminShowCommands) SetShowHardwareType(v bool)`

SetShowHardwareType sets ShowHardwareType field to given value.

### HasShowHardwareType

`func (o *AdmingroupAdminShowCommands) HasShowHardwareType() bool`

HasShowHardwareType returns a boolean if a field has been set.

### GetShowHardwareStatus

`func (o *AdmingroupAdminShowCommands) GetShowHardwareStatus() bool`

GetShowHardwareStatus returns the ShowHardwareStatus field if non-nil, zero value otherwise.

### GetShowHardwareStatusOk

`func (o *AdmingroupAdminShowCommands) GetShowHardwareStatusOk() (*bool, bool)`

GetShowHardwareStatusOk returns a tuple with the ShowHardwareStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHardwareStatus

`func (o *AdmingroupAdminShowCommands) SetShowHardwareStatus(v bool)`

SetShowHardwareStatus sets ShowHardwareStatus field to given value.

### HasShowHardwareStatus

`func (o *AdmingroupAdminShowCommands) HasShowHardwareStatus() bool`

HasShowHardwareStatus returns a boolean if a field has been set.

### GetShowHwid

`func (o *AdmingroupAdminShowCommands) GetShowHwid() bool`

GetShowHwid returns the ShowHwid field if non-nil, zero value otherwise.

### GetShowHwidOk

`func (o *AdmingroupAdminShowCommands) GetShowHwidOk() (*bool, bool)`

GetShowHwidOk returns a tuple with the ShowHwid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHwid

`func (o *AdmingroupAdminShowCommands) SetShowHwid(v bool)`

SetShowHwid sets ShowHwid field to given value.

### HasShowHwid

`func (o *AdmingroupAdminShowCommands) HasShowHwid() bool`

HasShowHwid returns a boolean if a field has been set.

### GetShowIbtrap

`func (o *AdmingroupAdminShowCommands) GetShowIbtrap() bool`

GetShowIbtrap returns the ShowIbtrap field if non-nil, zero value otherwise.

### GetShowIbtrapOk

`func (o *AdmingroupAdminShowCommands) GetShowIbtrapOk() (*bool, bool)`

GetShowIbtrapOk returns a tuple with the ShowIbtrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowIbtrap

`func (o *AdmingroupAdminShowCommands) SetShowIbtrap(v bool)`

SetShowIbtrap sets ShowIbtrap field to given value.

### HasShowIbtrap

`func (o *AdmingroupAdminShowCommands) HasShowIbtrap() bool`

HasShowIbtrap returns a boolean if a field has been set.

### GetShowHwIdent

`func (o *AdmingroupAdminShowCommands) GetShowHwIdent() bool`

GetShowHwIdent returns the ShowHwIdent field if non-nil, zero value otherwise.

### GetShowHwIdentOk

`func (o *AdmingroupAdminShowCommands) GetShowHwIdentOk() (*bool, bool)`

GetShowHwIdentOk returns a tuple with the ShowHwIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHwIdent

`func (o *AdmingroupAdminShowCommands) SetShowHwIdent(v bool)`

SetShowHwIdent sets ShowHwIdent field to given value.

### HasShowHwIdent

`func (o *AdmingroupAdminShowCommands) HasShowHwIdent() bool`

HasShowHwIdent returns a boolean if a field has been set.

### GetShowLog

`func (o *AdmingroupAdminShowCommands) GetShowLog() bool`

GetShowLog returns the ShowLog field if non-nil, zero value otherwise.

### GetShowLogOk

`func (o *AdmingroupAdminShowCommands) GetShowLogOk() (*bool, bool)`

GetShowLogOk returns a tuple with the ShowLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLog

`func (o *AdmingroupAdminShowCommands) SetShowLog(v bool)`

SetShowLog sets ShowLog field to given value.

### HasShowLog

`func (o *AdmingroupAdminShowCommands) HasShowLog() bool`

HasShowLog returns a boolean if a field has been set.

### GetShowLogfiles

`func (o *AdmingroupAdminShowCommands) GetShowLogfiles() bool`

GetShowLogfiles returns the ShowLogfiles field if non-nil, zero value otherwise.

### GetShowLogfilesOk

`func (o *AdmingroupAdminShowCommands) GetShowLogfilesOk() (*bool, bool)`

GetShowLogfilesOk returns a tuple with the ShowLogfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLogfiles

`func (o *AdmingroupAdminShowCommands) SetShowLogfiles(v bool)`

SetShowLogfiles sets ShowLogfiles field to given value.

### HasShowLogfiles

`func (o *AdmingroupAdminShowCommands) HasShowLogfiles() bool`

HasShowLogfiles returns a boolean if a field has been set.

### GetShowMemory

`func (o *AdmingroupAdminShowCommands) GetShowMemory() bool`

GetShowMemory returns the ShowMemory field if non-nil, zero value otherwise.

### GetShowMemoryOk

`func (o *AdmingroupAdminShowCommands) GetShowMemoryOk() (*bool, bool)`

GetShowMemoryOk returns a tuple with the ShowMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowMemory

`func (o *AdmingroupAdminShowCommands) SetShowMemory(v bool)`

SetShowMemory sets ShowMemory field to given value.

### HasShowMemory

`func (o *AdmingroupAdminShowCommands) HasShowMemory() bool`

HasShowMemory returns a boolean if a field has been set.

### GetShowNtp

`func (o *AdmingroupAdminShowCommands) GetShowNtp() bool`

GetShowNtp returns the ShowNtp field if non-nil, zero value otherwise.

### GetShowNtpOk

`func (o *AdmingroupAdminShowCommands) GetShowNtpOk() (*bool, bool)`

GetShowNtpOk returns a tuple with the ShowNtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNtp

`func (o *AdmingroupAdminShowCommands) SetShowNtp(v bool)`

SetShowNtp sets ShowNtp field to given value.

### HasShowNtp

`func (o *AdmingroupAdminShowCommands) HasShowNtp() bool`

HasShowNtp returns a boolean if a field has been set.

### GetShowReportingUserCapabilities

`func (o *AdmingroupAdminShowCommands) GetShowReportingUserCapabilities() bool`

GetShowReportingUserCapabilities returns the ShowReportingUserCapabilities field if non-nil, zero value otherwise.

### GetShowReportingUserCapabilitiesOk

`func (o *AdmingroupAdminShowCommands) GetShowReportingUserCapabilitiesOk() (*bool, bool)`

GetShowReportingUserCapabilitiesOk returns a tuple with the ShowReportingUserCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowReportingUserCapabilities

`func (o *AdmingroupAdminShowCommands) SetShowReportingUserCapabilities(v bool)`

SetShowReportingUserCapabilities sets ShowReportingUserCapabilities field to given value.

### HasShowReportingUserCapabilities

`func (o *AdmingroupAdminShowCommands) HasShowReportingUserCapabilities() bool`

HasShowReportingUserCapabilities returns a boolean if a field has been set.

### GetShowRpzRecursiveOnly

`func (o *AdmingroupAdminShowCommands) GetShowRpzRecursiveOnly() bool`

GetShowRpzRecursiveOnly returns the ShowRpzRecursiveOnly field if non-nil, zero value otherwise.

### GetShowRpzRecursiveOnlyOk

`func (o *AdmingroupAdminShowCommands) GetShowRpzRecursiveOnlyOk() (*bool, bool)`

GetShowRpzRecursiveOnlyOk returns a tuple with the ShowRpzRecursiveOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowRpzRecursiveOnly

`func (o *AdmingroupAdminShowCommands) SetShowRpzRecursiveOnly(v bool)`

SetShowRpzRecursiveOnly sets ShowRpzRecursiveOnly field to given value.

### HasShowRpzRecursiveOnly

`func (o *AdmingroupAdminShowCommands) HasShowRpzRecursiveOnly() bool`

HasShowRpzRecursiveOnly returns a boolean if a field has been set.

### GetShowScheduled

`func (o *AdmingroupAdminShowCommands) GetShowScheduled() bool`

GetShowScheduled returns the ShowScheduled field if non-nil, zero value otherwise.

### GetShowScheduledOk

`func (o *AdmingroupAdminShowCommands) GetShowScheduledOk() (*bool, bool)`

GetShowScheduledOk returns a tuple with the ShowScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowScheduled

`func (o *AdmingroupAdminShowCommands) SetShowScheduled(v bool)`

SetShowScheduled sets ShowScheduled field to given value.

### HasShowScheduled

`func (o *AdmingroupAdminShowCommands) HasShowScheduled() bool`

HasShowScheduled returns a boolean if a field has been set.

### GetShowSnmp

`func (o *AdmingroupAdminShowCommands) GetShowSnmp() bool`

GetShowSnmp returns the ShowSnmp field if non-nil, zero value otherwise.

### GetShowSnmpOk

`func (o *AdmingroupAdminShowCommands) GetShowSnmpOk() (*bool, bool)`

GetShowSnmpOk returns a tuple with the ShowSnmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSnmp

`func (o *AdmingroupAdminShowCommands) SetShowSnmp(v bool)`

SetShowSnmp sets ShowSnmp field to given value.

### HasShowSnmp

`func (o *AdmingroupAdminShowCommands) HasShowSnmp() bool`

HasShowSnmp returns a boolean if a field has been set.

### GetShowStatus

`func (o *AdmingroupAdminShowCommands) GetShowStatus() bool`

GetShowStatus returns the ShowStatus field if non-nil, zero value otherwise.

### GetShowStatusOk

`func (o *AdmingroupAdminShowCommands) GetShowStatusOk() (*bool, bool)`

GetShowStatusOk returns a tuple with the ShowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowStatus

`func (o *AdmingroupAdminShowCommands) SetShowStatus(v bool)`

SetShowStatus sets ShowStatus field to given value.

### HasShowStatus

`func (o *AdmingroupAdminShowCommands) HasShowStatus() bool`

HasShowStatus returns a boolean if a field has been set.

### GetShowTechSupport

`func (o *AdmingroupAdminShowCommands) GetShowTechSupport() bool`

GetShowTechSupport returns the ShowTechSupport field if non-nil, zero value otherwise.

### GetShowTechSupportOk

`func (o *AdmingroupAdminShowCommands) GetShowTechSupportOk() (*bool, bool)`

GetShowTechSupportOk returns a tuple with the ShowTechSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTechSupport

`func (o *AdmingroupAdminShowCommands) SetShowTechSupport(v bool)`

SetShowTechSupport sets ShowTechSupport field to given value.

### HasShowTechSupport

`func (o *AdmingroupAdminShowCommands) HasShowTechSupport() bool`

HasShowTechSupport returns a boolean if a field has been set.

### GetShowTemperature

`func (o *AdmingroupAdminShowCommands) GetShowTemperature() bool`

GetShowTemperature returns the ShowTemperature field if non-nil, zero value otherwise.

### GetShowTemperatureOk

`func (o *AdmingroupAdminShowCommands) GetShowTemperatureOk() (*bool, bool)`

GetShowTemperatureOk returns a tuple with the ShowTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowTemperature

`func (o *AdmingroupAdminShowCommands) SetShowTemperature(v bool)`

SetShowTemperature sets ShowTemperature field to given value.

### HasShowTemperature

`func (o *AdmingroupAdminShowCommands) HasShowTemperature() bool`

HasShowTemperature returns a boolean if a field has been set.

### GetShowThresholdtrap

`func (o *AdmingroupAdminShowCommands) GetShowThresholdtrap() bool`

GetShowThresholdtrap returns the ShowThresholdtrap field if non-nil, zero value otherwise.

### GetShowThresholdtrapOk

`func (o *AdmingroupAdminShowCommands) GetShowThresholdtrapOk() (*bool, bool)`

GetShowThresholdtrapOk returns a tuple with the ShowThresholdtrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowThresholdtrap

`func (o *AdmingroupAdminShowCommands) SetShowThresholdtrap(v bool)`

SetShowThresholdtrap sets ShowThresholdtrap field to given value.

### HasShowThresholdtrap

`func (o *AdmingroupAdminShowCommands) HasShowThresholdtrap() bool`

HasShowThresholdtrap returns a boolean if a field has been set.

### GetShowUpgradeHistory

`func (o *AdmingroupAdminShowCommands) GetShowUpgradeHistory() bool`

GetShowUpgradeHistory returns the ShowUpgradeHistory field if non-nil, zero value otherwise.

### GetShowUpgradeHistoryOk

`func (o *AdmingroupAdminShowCommands) GetShowUpgradeHistoryOk() (*bool, bool)`

GetShowUpgradeHistoryOk returns a tuple with the ShowUpgradeHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowUpgradeHistory

`func (o *AdmingroupAdminShowCommands) SetShowUpgradeHistory(v bool)`

SetShowUpgradeHistory sets ShowUpgradeHistory field to given value.

### HasShowUpgradeHistory

`func (o *AdmingroupAdminShowCommands) HasShowUpgradeHistory() bool`

HasShowUpgradeHistory returns a boolean if a field has been set.

### GetShowUptime

`func (o *AdmingroupAdminShowCommands) GetShowUptime() bool`

GetShowUptime returns the ShowUptime field if non-nil, zero value otherwise.

### GetShowUptimeOk

`func (o *AdmingroupAdminShowCommands) GetShowUptimeOk() (*bool, bool)`

GetShowUptimeOk returns a tuple with the ShowUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowUptime

`func (o *AdmingroupAdminShowCommands) SetShowUptime(v bool)`

SetShowUptime sets ShowUptime field to given value.

### HasShowUptime

`func (o *AdmingroupAdminShowCommands) HasShowUptime() bool`

HasShowUptime returns a boolean if a field has been set.

### GetShowVersion

`func (o *AdmingroupAdminShowCommands) GetShowVersion() bool`

GetShowVersion returns the ShowVersion field if non-nil, zero value otherwise.

### GetShowVersionOk

`func (o *AdmingroupAdminShowCommands) GetShowVersionOk() (*bool, bool)`

GetShowVersionOk returns a tuple with the ShowVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowVersion

`func (o *AdmingroupAdminShowCommands) SetShowVersion(v bool)`

SetShowVersion sets ShowVersion field to given value.

### HasShowVersion

`func (o *AdmingroupAdminShowCommands) HasShowVersion() bool`

HasShowVersion returns a boolean if a field has been set.

### GetShowAnalyticsDatabaseDumps

`func (o *AdmingroupAdminShowCommands) GetShowAnalyticsDatabaseDumps() bool`

GetShowAnalyticsDatabaseDumps returns the ShowAnalyticsDatabaseDumps field if non-nil, zero value otherwise.

### GetShowAnalyticsDatabaseDumpsOk

`func (o *AdmingroupAdminShowCommands) GetShowAnalyticsDatabaseDumpsOk() (*bool, bool)`

GetShowAnalyticsDatabaseDumpsOk returns a tuple with the ShowAnalyticsDatabaseDumps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAnalyticsDatabaseDumps

`func (o *AdmingroupAdminShowCommands) SetShowAnalyticsDatabaseDumps(v bool)`

SetShowAnalyticsDatabaseDumps sets ShowAnalyticsDatabaseDumps field to given value.

### HasShowAnalyticsDatabaseDumps

`func (o *AdmingroupAdminShowCommands) HasShowAnalyticsDatabaseDumps() bool`

HasShowAnalyticsDatabaseDumps returns a boolean if a field has been set.

### GetShowCores

`func (o *AdmingroupAdminShowCommands) GetShowCores() bool`

GetShowCores returns the ShowCores field if non-nil, zero value otherwise.

### GetShowCoresOk

`func (o *AdmingroupAdminShowCommands) GetShowCoresOk() (*bool, bool)`

GetShowCoresOk returns a tuple with the ShowCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCores

`func (o *AdmingroupAdminShowCommands) SetShowCores(v bool)`

SetShowCores sets ShowCores field to given value.

### HasShowCores

`func (o *AdmingroupAdminShowCommands) HasShowCores() bool`

HasShowCores returns a boolean if a field has been set.

### GetShowCoresummary

`func (o *AdmingroupAdminShowCommands) GetShowCoresummary() bool`

GetShowCoresummary returns the ShowCoresummary field if non-nil, zero value otherwise.

### GetShowCoresummaryOk

`func (o *AdmingroupAdminShowCommands) GetShowCoresummaryOk() (*bool, bool)`

GetShowCoresummaryOk returns a tuple with the ShowCoresummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCoresummary

`func (o *AdmingroupAdminShowCommands) SetShowCoresummary(v bool)`

SetShowCoresummary sets ShowCoresummary field to given value.

### HasShowCoresummary

`func (o *AdmingroupAdminShowCommands) HasShowCoresummary() bool`

HasShowCoresummary returns a boolean if a field has been set.

### GetShowCspThreatDb

`func (o *AdmingroupAdminShowCommands) GetShowCspThreatDb() bool`

GetShowCspThreatDb returns the ShowCspThreatDb field if non-nil, zero value otherwise.

### GetShowCspThreatDbOk

`func (o *AdmingroupAdminShowCommands) GetShowCspThreatDbOk() (*bool, bool)`

GetShowCspThreatDbOk returns a tuple with the ShowCspThreatDb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCspThreatDb

`func (o *AdmingroupAdminShowCommands) SetShowCspThreatDb(v bool)`

SetShowCspThreatDb sets ShowCspThreatDb field to given value.

### HasShowCspThreatDb

`func (o *AdmingroupAdminShowCommands) HasShowCspThreatDb() bool`

HasShowCspThreatDb returns a boolean if a field has been set.

### GetShowHsmGroup

`func (o *AdmingroupAdminShowCommands) GetShowHsmGroup() bool`

GetShowHsmGroup returns the ShowHsmGroup field if non-nil, zero value otherwise.

### GetShowHsmGroupOk

`func (o *AdmingroupAdminShowCommands) GetShowHsmGroupOk() (*bool, bool)`

GetShowHsmGroupOk returns a tuple with the ShowHsmGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHsmGroup

`func (o *AdmingroupAdminShowCommands) SetShowHsmGroup(v bool)`

SetShowHsmGroup sets ShowHsmGroup field to given value.

### HasShowHsmGroup

`func (o *AdmingroupAdminShowCommands) HasShowHsmGroup() bool`

HasShowHsmGroup returns a boolean if a field has been set.

### GetShowHsmInfo

`func (o *AdmingroupAdminShowCommands) GetShowHsmInfo() bool`

GetShowHsmInfo returns the ShowHsmInfo field if non-nil, zero value otherwise.

### GetShowHsmInfoOk

`func (o *AdmingroupAdminShowCommands) GetShowHsmInfoOk() (*bool, bool)`

GetShowHsmInfoOk returns a tuple with the ShowHsmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHsmInfo

`func (o *AdmingroupAdminShowCommands) SetShowHsmInfo(v bool)`

SetShowHsmInfo sets ShowHsmInfo field to given value.

### HasShowHsmInfo

`func (o *AdmingroupAdminShowCommands) HasShowHsmInfo() bool`

HasShowHsmInfo returns a boolean if a field has been set.

### GetShowPmap

`func (o *AdmingroupAdminShowCommands) GetShowPmap() bool`

GetShowPmap returns the ShowPmap field if non-nil, zero value otherwise.

### GetShowPmapOk

`func (o *AdmingroupAdminShowCommands) GetShowPmapOk() (*bool, bool)`

GetShowPmapOk returns a tuple with the ShowPmap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPmap

`func (o *AdmingroupAdminShowCommands) SetShowPmap(v bool)`

SetShowPmap sets ShowPmap field to given value.

### HasShowPmap

`func (o *AdmingroupAdminShowCommands) HasShowPmap() bool`

HasShowPmap returns a boolean if a field has been set.

### GetShowProcess

`func (o *AdmingroupAdminShowCommands) GetShowProcess() bool`

GetShowProcess returns the ShowProcess field if non-nil, zero value otherwise.

### GetShowProcessOk

`func (o *AdmingroupAdminShowCommands) GetShowProcessOk() (*bool, bool)`

GetShowProcessOk returns a tuple with the ShowProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowProcess

`func (o *AdmingroupAdminShowCommands) SetShowProcess(v bool)`

SetShowProcess sets ShowProcess field to given value.

### HasShowProcess

`func (o *AdmingroupAdminShowCommands) HasShowProcess() bool`

HasShowProcess returns a boolean if a field has been set.

### GetShowPstack

`func (o *AdmingroupAdminShowCommands) GetShowPstack() bool`

GetShowPstack returns the ShowPstack field if non-nil, zero value otherwise.

### GetShowPstackOk

`func (o *AdmingroupAdminShowCommands) GetShowPstackOk() (*bool, bool)`

GetShowPstackOk returns a tuple with the ShowPstack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPstack

`func (o *AdmingroupAdminShowCommands) SetShowPstack(v bool)`

SetShowPstack sets ShowPstack field to given value.

### HasShowPstack

`func (o *AdmingroupAdminShowCommands) HasShowPstack() bool`

HasShowPstack returns a boolean if a field has been set.

### GetShowSafenetSupportInfo

`func (o *AdmingroupAdminShowCommands) GetShowSafenetSupportInfo() bool`

GetShowSafenetSupportInfo returns the ShowSafenetSupportInfo field if non-nil, zero value otherwise.

### GetShowSafenetSupportInfoOk

`func (o *AdmingroupAdminShowCommands) GetShowSafenetSupportInfoOk() (*bool, bool)`

GetShowSafenetSupportInfoOk returns a tuple with the ShowSafenetSupportInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSafenetSupportInfo

`func (o *AdmingroupAdminShowCommands) SetShowSafenetSupportInfo(v bool)`

SetShowSafenetSupportInfo sets ShowSafenetSupportInfo field to given value.

### HasShowSafenetSupportInfo

`func (o *AdmingroupAdminShowCommands) HasShowSafenetSupportInfo() bool`

HasShowSafenetSupportInfo returns a boolean if a field has been set.

### GetShowWredStats

`func (o *AdmingroupAdminShowCommands) GetShowWredStats() bool`

GetShowWredStats returns the ShowWredStats field if non-nil, zero value otherwise.

### GetShowWredStatsOk

`func (o *AdmingroupAdminShowCommands) GetShowWredStatsOk() (*bool, bool)`

GetShowWredStatsOk returns a tuple with the ShowWredStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowWredStats

`func (o *AdmingroupAdminShowCommands) SetShowWredStats(v bool)`

SetShowWredStats sets ShowWredStats field to given value.

### HasShowWredStats

`func (o *AdmingroupAdminShowCommands) HasShowWredStats() bool`

HasShowWredStats returns a boolean if a field has been set.

### GetShowWredStatus

`func (o *AdmingroupAdminShowCommands) GetShowWredStatus() bool`

GetShowWredStatus returns the ShowWredStatus field if non-nil, zero value otherwise.

### GetShowWredStatusOk

`func (o *AdmingroupAdminShowCommands) GetShowWredStatusOk() (*bool, bool)`

GetShowWredStatusOk returns a tuple with the ShowWredStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowWredStatus

`func (o *AdmingroupAdminShowCommands) SetShowWredStatus(v bool)`

SetShowWredStatus sets ShowWredStatus field to given value.

### HasShowWredStatus

`func (o *AdmingroupAdminShowCommands) HasShowWredStatus() bool`

HasShowWredStatus returns a boolean if a field has been set.

### GetShowNtpStratum

`func (o *AdmingroupAdminShowCommands) GetShowNtpStratum() bool`

GetShowNtpStratum returns the ShowNtpStratum field if non-nil, zero value otherwise.

### GetShowNtpStratumOk

`func (o *AdmingroupAdminShowCommands) GetShowNtpStratumOk() (*bool, bool)`

GetShowNtpStratumOk returns a tuple with the ShowNtpStratum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNtpStratum

`func (o *AdmingroupAdminShowCommands) SetShowNtpStratum(v bool)`

SetShowNtpStratum sets ShowNtpStratum field to given value.

### HasShowNtpStratum

`func (o *AdmingroupAdminShowCommands) HasShowNtpStratum() bool`

HasShowNtpStratum returns a boolean if a field has been set.

### GetShowPcDomain

`func (o *AdmingroupAdminShowCommands) GetShowPcDomain() bool`

GetShowPcDomain returns the ShowPcDomain field if non-nil, zero value otherwise.

### GetShowPcDomainOk

`func (o *AdmingroupAdminShowCommands) GetShowPcDomainOk() (*bool, bool)`

GetShowPcDomainOk returns a tuple with the ShowPcDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPcDomain

`func (o *AdmingroupAdminShowCommands) SetShowPcDomain(v bool)`

SetShowPcDomain sets ShowPcDomain field to given value.

### HasShowPcDomain

`func (o *AdmingroupAdminShowCommands) HasShowPcDomain() bool`

HasShowPcDomain returns a boolean if a field has been set.

### GetShowReportFrequency

`func (o *AdmingroupAdminShowCommands) GetShowReportFrequency() bool`

GetShowReportFrequency returns the ShowReportFrequency field if non-nil, zero value otherwise.

### GetShowReportFrequencyOk

`func (o *AdmingroupAdminShowCommands) GetShowReportFrequencyOk() (*bool, bool)`

GetShowReportFrequencyOk returns a tuple with the ShowReportFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowReportFrequency

`func (o *AdmingroupAdminShowCommands) SetShowReportFrequency(v bool)`

SetShowReportFrequency sets ShowReportFrequency field to given value.

### HasShowReportFrequency

`func (o *AdmingroupAdminShowCommands) HasShowReportFrequency() bool`

HasShowReportFrequency returns a boolean if a field has been set.

### GetShowQueryLoggingWarnings

`func (o *AdmingroupAdminShowCommands) GetShowQueryLoggingWarnings() bool`

GetShowQueryLoggingWarnings returns the ShowQueryLoggingWarnings field if non-nil, zero value otherwise.

### GetShowQueryLoggingWarningsOk

`func (o *AdmingroupAdminShowCommands) GetShowQueryLoggingWarningsOk() (*bool, bool)`

GetShowQueryLoggingWarningsOk returns a tuple with the ShowQueryLoggingWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowQueryLoggingWarnings

`func (o *AdmingroupAdminShowCommands) SetShowQueryLoggingWarnings(v bool)`

SetShowQueryLoggingWarnings sets ShowQueryLoggingWarnings field to given value.

### HasShowQueryLoggingWarnings

`func (o *AdmingroupAdminShowCommands) HasShowQueryLoggingWarnings() bool`

HasShowQueryLoggingWarnings returns a boolean if a field has been set.

### GetEnableAll

`func (o *AdmingroupAdminShowCommands) GetEnableAll() bool`

GetEnableAll returns the EnableAll field if non-nil, zero value otherwise.

### GetEnableAllOk

`func (o *AdmingroupAdminShowCommands) GetEnableAllOk() (*bool, bool)`

GetEnableAllOk returns a tuple with the EnableAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAll

`func (o *AdmingroupAdminShowCommands) SetEnableAll(v bool)`

SetEnableAll sets EnableAll field to given value.

### HasEnableAll

`func (o *AdmingroupAdminShowCommands) HasEnableAll() bool`

HasEnableAll returns a boolean if a field has been set.

### GetDisableAll

`func (o *AdmingroupAdminShowCommands) GetDisableAll() bool`

GetDisableAll returns the DisableAll field if non-nil, zero value otherwise.

### GetDisableAllOk

`func (o *AdmingroupAdminShowCommands) GetDisableAllOk() (*bool, bool)`

GetDisableAllOk returns a tuple with the DisableAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAll

`func (o *AdmingroupAdminShowCommands) SetDisableAll(v bool)`

SetDisableAll sets DisableAll field to given value.

### HasDisableAll

`func (o *AdmingroupAdminShowCommands) HasDisableAll() bool`

HasDisableAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


