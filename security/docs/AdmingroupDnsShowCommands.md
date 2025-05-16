# AdmingroupDnsShowCommands

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShowLogGuestLookups** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDnsGssTsig** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDns** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDnstapStats** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDnstapStatus** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDnsOverTlsConfig** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDnsOverTlsStatus** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDnsOverTlsStats** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDohConfig** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDohStatus** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDohStats** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowExtraDnsNameValidations** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowMsStickyIp** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDnsRrl** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowEnableMatchRecursiveOnly** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowMaxRecursionDepth** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowMaxRecursionQueries** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowMonitor** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowQueryCapture** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDtcEa** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDtcGeoip** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowRestartAnycastWithDnsRestart** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDnsAccel** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowDnsAccelDebug** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowAllowQueryDomain** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**ShowAllowQueryDomainViews** | Pointer to **bool** | If True then CLI user has permission to run the command | [optional] 
**EnableAll** | Pointer to **bool** | If True then enable all fields | [optional] 
**DisableAll** | Pointer to **bool** | If True then disable all fields | [optional] 

## Methods

### NewAdmingroupDnsShowCommands

`func NewAdmingroupDnsShowCommands() *AdmingroupDnsShowCommands`

NewAdmingroupDnsShowCommands instantiates a new AdmingroupDnsShowCommands object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdmingroupDnsShowCommandsWithDefaults

`func NewAdmingroupDnsShowCommandsWithDefaults() *AdmingroupDnsShowCommands`

NewAdmingroupDnsShowCommandsWithDefaults instantiates a new AdmingroupDnsShowCommands object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowLogGuestLookups

`func (o *AdmingroupDnsShowCommands) GetShowLogGuestLookups() bool`

GetShowLogGuestLookups returns the ShowLogGuestLookups field if non-nil, zero value otherwise.

### GetShowLogGuestLookupsOk

`func (o *AdmingroupDnsShowCommands) GetShowLogGuestLookupsOk() (*bool, bool)`

GetShowLogGuestLookupsOk returns a tuple with the ShowLogGuestLookups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLogGuestLookups

`func (o *AdmingroupDnsShowCommands) SetShowLogGuestLookups(v bool)`

SetShowLogGuestLookups sets ShowLogGuestLookups field to given value.

### HasShowLogGuestLookups

`func (o *AdmingroupDnsShowCommands) HasShowLogGuestLookups() bool`

HasShowLogGuestLookups returns a boolean if a field has been set.

### GetShowDnsGssTsig

`func (o *AdmingroupDnsShowCommands) GetShowDnsGssTsig() bool`

GetShowDnsGssTsig returns the ShowDnsGssTsig field if non-nil, zero value otherwise.

### GetShowDnsGssTsigOk

`func (o *AdmingroupDnsShowCommands) GetShowDnsGssTsigOk() (*bool, bool)`

GetShowDnsGssTsigOk returns a tuple with the ShowDnsGssTsig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDnsGssTsig

`func (o *AdmingroupDnsShowCommands) SetShowDnsGssTsig(v bool)`

SetShowDnsGssTsig sets ShowDnsGssTsig field to given value.

### HasShowDnsGssTsig

`func (o *AdmingroupDnsShowCommands) HasShowDnsGssTsig() bool`

HasShowDnsGssTsig returns a boolean if a field has been set.

### GetShowDns

`func (o *AdmingroupDnsShowCommands) GetShowDns() bool`

GetShowDns returns the ShowDns field if non-nil, zero value otherwise.

### GetShowDnsOk

`func (o *AdmingroupDnsShowCommands) GetShowDnsOk() (*bool, bool)`

GetShowDnsOk returns a tuple with the ShowDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDns

`func (o *AdmingroupDnsShowCommands) SetShowDns(v bool)`

SetShowDns sets ShowDns field to given value.

### HasShowDns

`func (o *AdmingroupDnsShowCommands) HasShowDns() bool`

HasShowDns returns a boolean if a field has been set.

### GetShowDnstapStats

`func (o *AdmingroupDnsShowCommands) GetShowDnstapStats() bool`

GetShowDnstapStats returns the ShowDnstapStats field if non-nil, zero value otherwise.

### GetShowDnstapStatsOk

`func (o *AdmingroupDnsShowCommands) GetShowDnstapStatsOk() (*bool, bool)`

GetShowDnstapStatsOk returns a tuple with the ShowDnstapStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDnstapStats

`func (o *AdmingroupDnsShowCommands) SetShowDnstapStats(v bool)`

SetShowDnstapStats sets ShowDnstapStats field to given value.

### HasShowDnstapStats

`func (o *AdmingroupDnsShowCommands) HasShowDnstapStats() bool`

HasShowDnstapStats returns a boolean if a field has been set.

### GetShowDnstapStatus

`func (o *AdmingroupDnsShowCommands) GetShowDnstapStatus() bool`

GetShowDnstapStatus returns the ShowDnstapStatus field if non-nil, zero value otherwise.

### GetShowDnstapStatusOk

`func (o *AdmingroupDnsShowCommands) GetShowDnstapStatusOk() (*bool, bool)`

GetShowDnstapStatusOk returns a tuple with the ShowDnstapStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDnstapStatus

`func (o *AdmingroupDnsShowCommands) SetShowDnstapStatus(v bool)`

SetShowDnstapStatus sets ShowDnstapStatus field to given value.

### HasShowDnstapStatus

`func (o *AdmingroupDnsShowCommands) HasShowDnstapStatus() bool`

HasShowDnstapStatus returns a boolean if a field has been set.

### GetShowDnsOverTlsConfig

`func (o *AdmingroupDnsShowCommands) GetShowDnsOverTlsConfig() bool`

GetShowDnsOverTlsConfig returns the ShowDnsOverTlsConfig field if non-nil, zero value otherwise.

### GetShowDnsOverTlsConfigOk

`func (o *AdmingroupDnsShowCommands) GetShowDnsOverTlsConfigOk() (*bool, bool)`

GetShowDnsOverTlsConfigOk returns a tuple with the ShowDnsOverTlsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDnsOverTlsConfig

`func (o *AdmingroupDnsShowCommands) SetShowDnsOverTlsConfig(v bool)`

SetShowDnsOverTlsConfig sets ShowDnsOverTlsConfig field to given value.

### HasShowDnsOverTlsConfig

`func (o *AdmingroupDnsShowCommands) HasShowDnsOverTlsConfig() bool`

HasShowDnsOverTlsConfig returns a boolean if a field has been set.

### GetShowDnsOverTlsStatus

`func (o *AdmingroupDnsShowCommands) GetShowDnsOverTlsStatus() bool`

GetShowDnsOverTlsStatus returns the ShowDnsOverTlsStatus field if non-nil, zero value otherwise.

### GetShowDnsOverTlsStatusOk

`func (o *AdmingroupDnsShowCommands) GetShowDnsOverTlsStatusOk() (*bool, bool)`

GetShowDnsOverTlsStatusOk returns a tuple with the ShowDnsOverTlsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDnsOverTlsStatus

`func (o *AdmingroupDnsShowCommands) SetShowDnsOverTlsStatus(v bool)`

SetShowDnsOverTlsStatus sets ShowDnsOverTlsStatus field to given value.

### HasShowDnsOverTlsStatus

`func (o *AdmingroupDnsShowCommands) HasShowDnsOverTlsStatus() bool`

HasShowDnsOverTlsStatus returns a boolean if a field has been set.

### GetShowDnsOverTlsStats

`func (o *AdmingroupDnsShowCommands) GetShowDnsOverTlsStats() bool`

GetShowDnsOverTlsStats returns the ShowDnsOverTlsStats field if non-nil, zero value otherwise.

### GetShowDnsOverTlsStatsOk

`func (o *AdmingroupDnsShowCommands) GetShowDnsOverTlsStatsOk() (*bool, bool)`

GetShowDnsOverTlsStatsOk returns a tuple with the ShowDnsOverTlsStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDnsOverTlsStats

`func (o *AdmingroupDnsShowCommands) SetShowDnsOverTlsStats(v bool)`

SetShowDnsOverTlsStats sets ShowDnsOverTlsStats field to given value.

### HasShowDnsOverTlsStats

`func (o *AdmingroupDnsShowCommands) HasShowDnsOverTlsStats() bool`

HasShowDnsOverTlsStats returns a boolean if a field has been set.

### GetShowDohConfig

`func (o *AdmingroupDnsShowCommands) GetShowDohConfig() bool`

GetShowDohConfig returns the ShowDohConfig field if non-nil, zero value otherwise.

### GetShowDohConfigOk

`func (o *AdmingroupDnsShowCommands) GetShowDohConfigOk() (*bool, bool)`

GetShowDohConfigOk returns a tuple with the ShowDohConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDohConfig

`func (o *AdmingroupDnsShowCommands) SetShowDohConfig(v bool)`

SetShowDohConfig sets ShowDohConfig field to given value.

### HasShowDohConfig

`func (o *AdmingroupDnsShowCommands) HasShowDohConfig() bool`

HasShowDohConfig returns a boolean if a field has been set.

### GetShowDohStatus

`func (o *AdmingroupDnsShowCommands) GetShowDohStatus() bool`

GetShowDohStatus returns the ShowDohStatus field if non-nil, zero value otherwise.

### GetShowDohStatusOk

`func (o *AdmingroupDnsShowCommands) GetShowDohStatusOk() (*bool, bool)`

GetShowDohStatusOk returns a tuple with the ShowDohStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDohStatus

`func (o *AdmingroupDnsShowCommands) SetShowDohStatus(v bool)`

SetShowDohStatus sets ShowDohStatus field to given value.

### HasShowDohStatus

`func (o *AdmingroupDnsShowCommands) HasShowDohStatus() bool`

HasShowDohStatus returns a boolean if a field has been set.

### GetShowDohStats

`func (o *AdmingroupDnsShowCommands) GetShowDohStats() bool`

GetShowDohStats returns the ShowDohStats field if non-nil, zero value otherwise.

### GetShowDohStatsOk

`func (o *AdmingroupDnsShowCommands) GetShowDohStatsOk() (*bool, bool)`

GetShowDohStatsOk returns a tuple with the ShowDohStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDohStats

`func (o *AdmingroupDnsShowCommands) SetShowDohStats(v bool)`

SetShowDohStats sets ShowDohStats field to given value.

### HasShowDohStats

`func (o *AdmingroupDnsShowCommands) HasShowDohStats() bool`

HasShowDohStats returns a boolean if a field has been set.

### GetShowExtraDnsNameValidations

`func (o *AdmingroupDnsShowCommands) GetShowExtraDnsNameValidations() bool`

GetShowExtraDnsNameValidations returns the ShowExtraDnsNameValidations field if non-nil, zero value otherwise.

### GetShowExtraDnsNameValidationsOk

`func (o *AdmingroupDnsShowCommands) GetShowExtraDnsNameValidationsOk() (*bool, bool)`

GetShowExtraDnsNameValidationsOk returns a tuple with the ShowExtraDnsNameValidations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowExtraDnsNameValidations

`func (o *AdmingroupDnsShowCommands) SetShowExtraDnsNameValidations(v bool)`

SetShowExtraDnsNameValidations sets ShowExtraDnsNameValidations field to given value.

### HasShowExtraDnsNameValidations

`func (o *AdmingroupDnsShowCommands) HasShowExtraDnsNameValidations() bool`

HasShowExtraDnsNameValidations returns a boolean if a field has been set.

### GetShowMsStickyIp

`func (o *AdmingroupDnsShowCommands) GetShowMsStickyIp() bool`

GetShowMsStickyIp returns the ShowMsStickyIp field if non-nil, zero value otherwise.

### GetShowMsStickyIpOk

`func (o *AdmingroupDnsShowCommands) GetShowMsStickyIpOk() (*bool, bool)`

GetShowMsStickyIpOk returns a tuple with the ShowMsStickyIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowMsStickyIp

`func (o *AdmingroupDnsShowCommands) SetShowMsStickyIp(v bool)`

SetShowMsStickyIp sets ShowMsStickyIp field to given value.

### HasShowMsStickyIp

`func (o *AdmingroupDnsShowCommands) HasShowMsStickyIp() bool`

HasShowMsStickyIp returns a boolean if a field has been set.

### GetShowDnsRrl

`func (o *AdmingroupDnsShowCommands) GetShowDnsRrl() bool`

GetShowDnsRrl returns the ShowDnsRrl field if non-nil, zero value otherwise.

### GetShowDnsRrlOk

`func (o *AdmingroupDnsShowCommands) GetShowDnsRrlOk() (*bool, bool)`

GetShowDnsRrlOk returns a tuple with the ShowDnsRrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDnsRrl

`func (o *AdmingroupDnsShowCommands) SetShowDnsRrl(v bool)`

SetShowDnsRrl sets ShowDnsRrl field to given value.

### HasShowDnsRrl

`func (o *AdmingroupDnsShowCommands) HasShowDnsRrl() bool`

HasShowDnsRrl returns a boolean if a field has been set.

### GetShowEnableMatchRecursiveOnly

`func (o *AdmingroupDnsShowCommands) GetShowEnableMatchRecursiveOnly() bool`

GetShowEnableMatchRecursiveOnly returns the ShowEnableMatchRecursiveOnly field if non-nil, zero value otherwise.

### GetShowEnableMatchRecursiveOnlyOk

`func (o *AdmingroupDnsShowCommands) GetShowEnableMatchRecursiveOnlyOk() (*bool, bool)`

GetShowEnableMatchRecursiveOnlyOk returns a tuple with the ShowEnableMatchRecursiveOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowEnableMatchRecursiveOnly

`func (o *AdmingroupDnsShowCommands) SetShowEnableMatchRecursiveOnly(v bool)`

SetShowEnableMatchRecursiveOnly sets ShowEnableMatchRecursiveOnly field to given value.

### HasShowEnableMatchRecursiveOnly

`func (o *AdmingroupDnsShowCommands) HasShowEnableMatchRecursiveOnly() bool`

HasShowEnableMatchRecursiveOnly returns a boolean if a field has been set.

### GetShowMaxRecursionDepth

`func (o *AdmingroupDnsShowCommands) GetShowMaxRecursionDepth() bool`

GetShowMaxRecursionDepth returns the ShowMaxRecursionDepth field if non-nil, zero value otherwise.

### GetShowMaxRecursionDepthOk

`func (o *AdmingroupDnsShowCommands) GetShowMaxRecursionDepthOk() (*bool, bool)`

GetShowMaxRecursionDepthOk returns a tuple with the ShowMaxRecursionDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowMaxRecursionDepth

`func (o *AdmingroupDnsShowCommands) SetShowMaxRecursionDepth(v bool)`

SetShowMaxRecursionDepth sets ShowMaxRecursionDepth field to given value.

### HasShowMaxRecursionDepth

`func (o *AdmingroupDnsShowCommands) HasShowMaxRecursionDepth() bool`

HasShowMaxRecursionDepth returns a boolean if a field has been set.

### GetShowMaxRecursionQueries

`func (o *AdmingroupDnsShowCommands) GetShowMaxRecursionQueries() bool`

GetShowMaxRecursionQueries returns the ShowMaxRecursionQueries field if non-nil, zero value otherwise.

### GetShowMaxRecursionQueriesOk

`func (o *AdmingroupDnsShowCommands) GetShowMaxRecursionQueriesOk() (*bool, bool)`

GetShowMaxRecursionQueriesOk returns a tuple with the ShowMaxRecursionQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowMaxRecursionQueries

`func (o *AdmingroupDnsShowCommands) SetShowMaxRecursionQueries(v bool)`

SetShowMaxRecursionQueries sets ShowMaxRecursionQueries field to given value.

### HasShowMaxRecursionQueries

`func (o *AdmingroupDnsShowCommands) HasShowMaxRecursionQueries() bool`

HasShowMaxRecursionQueries returns a boolean if a field has been set.

### GetShowMonitor

`func (o *AdmingroupDnsShowCommands) GetShowMonitor() bool`

GetShowMonitor returns the ShowMonitor field if non-nil, zero value otherwise.

### GetShowMonitorOk

`func (o *AdmingroupDnsShowCommands) GetShowMonitorOk() (*bool, bool)`

GetShowMonitorOk returns a tuple with the ShowMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowMonitor

`func (o *AdmingroupDnsShowCommands) SetShowMonitor(v bool)`

SetShowMonitor sets ShowMonitor field to given value.

### HasShowMonitor

`func (o *AdmingroupDnsShowCommands) HasShowMonitor() bool`

HasShowMonitor returns a boolean if a field has been set.

### GetShowQueryCapture

`func (o *AdmingroupDnsShowCommands) GetShowQueryCapture() bool`

GetShowQueryCapture returns the ShowQueryCapture field if non-nil, zero value otherwise.

### GetShowQueryCaptureOk

`func (o *AdmingroupDnsShowCommands) GetShowQueryCaptureOk() (*bool, bool)`

GetShowQueryCaptureOk returns a tuple with the ShowQueryCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowQueryCapture

`func (o *AdmingroupDnsShowCommands) SetShowQueryCapture(v bool)`

SetShowQueryCapture sets ShowQueryCapture field to given value.

### HasShowQueryCapture

`func (o *AdmingroupDnsShowCommands) HasShowQueryCapture() bool`

HasShowQueryCapture returns a boolean if a field has been set.

### GetShowDtcEa

`func (o *AdmingroupDnsShowCommands) GetShowDtcEa() bool`

GetShowDtcEa returns the ShowDtcEa field if non-nil, zero value otherwise.

### GetShowDtcEaOk

`func (o *AdmingroupDnsShowCommands) GetShowDtcEaOk() (*bool, bool)`

GetShowDtcEaOk returns a tuple with the ShowDtcEa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDtcEa

`func (o *AdmingroupDnsShowCommands) SetShowDtcEa(v bool)`

SetShowDtcEa sets ShowDtcEa field to given value.

### HasShowDtcEa

`func (o *AdmingroupDnsShowCommands) HasShowDtcEa() bool`

HasShowDtcEa returns a boolean if a field has been set.

### GetShowDtcGeoip

`func (o *AdmingroupDnsShowCommands) GetShowDtcGeoip() bool`

GetShowDtcGeoip returns the ShowDtcGeoip field if non-nil, zero value otherwise.

### GetShowDtcGeoipOk

`func (o *AdmingroupDnsShowCommands) GetShowDtcGeoipOk() (*bool, bool)`

GetShowDtcGeoipOk returns a tuple with the ShowDtcGeoip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDtcGeoip

`func (o *AdmingroupDnsShowCommands) SetShowDtcGeoip(v bool)`

SetShowDtcGeoip sets ShowDtcGeoip field to given value.

### HasShowDtcGeoip

`func (o *AdmingroupDnsShowCommands) HasShowDtcGeoip() bool`

HasShowDtcGeoip returns a boolean if a field has been set.

### GetShowRestartAnycastWithDnsRestart

`func (o *AdmingroupDnsShowCommands) GetShowRestartAnycastWithDnsRestart() bool`

GetShowRestartAnycastWithDnsRestart returns the ShowRestartAnycastWithDnsRestart field if non-nil, zero value otherwise.

### GetShowRestartAnycastWithDnsRestartOk

`func (o *AdmingroupDnsShowCommands) GetShowRestartAnycastWithDnsRestartOk() (*bool, bool)`

GetShowRestartAnycastWithDnsRestartOk returns a tuple with the ShowRestartAnycastWithDnsRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowRestartAnycastWithDnsRestart

`func (o *AdmingroupDnsShowCommands) SetShowRestartAnycastWithDnsRestart(v bool)`

SetShowRestartAnycastWithDnsRestart sets ShowRestartAnycastWithDnsRestart field to given value.

### HasShowRestartAnycastWithDnsRestart

`func (o *AdmingroupDnsShowCommands) HasShowRestartAnycastWithDnsRestart() bool`

HasShowRestartAnycastWithDnsRestart returns a boolean if a field has been set.

### GetShowDnsAccel

`func (o *AdmingroupDnsShowCommands) GetShowDnsAccel() bool`

GetShowDnsAccel returns the ShowDnsAccel field if non-nil, zero value otherwise.

### GetShowDnsAccelOk

`func (o *AdmingroupDnsShowCommands) GetShowDnsAccelOk() (*bool, bool)`

GetShowDnsAccelOk returns a tuple with the ShowDnsAccel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDnsAccel

`func (o *AdmingroupDnsShowCommands) SetShowDnsAccel(v bool)`

SetShowDnsAccel sets ShowDnsAccel field to given value.

### HasShowDnsAccel

`func (o *AdmingroupDnsShowCommands) HasShowDnsAccel() bool`

HasShowDnsAccel returns a boolean if a field has been set.

### GetShowDnsAccelDebug

`func (o *AdmingroupDnsShowCommands) GetShowDnsAccelDebug() bool`

GetShowDnsAccelDebug returns the ShowDnsAccelDebug field if non-nil, zero value otherwise.

### GetShowDnsAccelDebugOk

`func (o *AdmingroupDnsShowCommands) GetShowDnsAccelDebugOk() (*bool, bool)`

GetShowDnsAccelDebugOk returns a tuple with the ShowDnsAccelDebug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowDnsAccelDebug

`func (o *AdmingroupDnsShowCommands) SetShowDnsAccelDebug(v bool)`

SetShowDnsAccelDebug sets ShowDnsAccelDebug field to given value.

### HasShowDnsAccelDebug

`func (o *AdmingroupDnsShowCommands) HasShowDnsAccelDebug() bool`

HasShowDnsAccelDebug returns a boolean if a field has been set.

### GetShowAllowQueryDomain

`func (o *AdmingroupDnsShowCommands) GetShowAllowQueryDomain() bool`

GetShowAllowQueryDomain returns the ShowAllowQueryDomain field if non-nil, zero value otherwise.

### GetShowAllowQueryDomainOk

`func (o *AdmingroupDnsShowCommands) GetShowAllowQueryDomainOk() (*bool, bool)`

GetShowAllowQueryDomainOk returns a tuple with the ShowAllowQueryDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAllowQueryDomain

`func (o *AdmingroupDnsShowCommands) SetShowAllowQueryDomain(v bool)`

SetShowAllowQueryDomain sets ShowAllowQueryDomain field to given value.

### HasShowAllowQueryDomain

`func (o *AdmingroupDnsShowCommands) HasShowAllowQueryDomain() bool`

HasShowAllowQueryDomain returns a boolean if a field has been set.

### GetShowAllowQueryDomainViews

`func (o *AdmingroupDnsShowCommands) GetShowAllowQueryDomainViews() bool`

GetShowAllowQueryDomainViews returns the ShowAllowQueryDomainViews field if non-nil, zero value otherwise.

### GetShowAllowQueryDomainViewsOk

`func (o *AdmingroupDnsShowCommands) GetShowAllowQueryDomainViewsOk() (*bool, bool)`

GetShowAllowQueryDomainViewsOk returns a tuple with the ShowAllowQueryDomainViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAllowQueryDomainViews

`func (o *AdmingroupDnsShowCommands) SetShowAllowQueryDomainViews(v bool)`

SetShowAllowQueryDomainViews sets ShowAllowQueryDomainViews field to given value.

### HasShowAllowQueryDomainViews

`func (o *AdmingroupDnsShowCommands) HasShowAllowQueryDomainViews() bool`

HasShowAllowQueryDomainViews returns a boolean if a field has been set.

### GetEnableAll

`func (o *AdmingroupDnsShowCommands) GetEnableAll() bool`

GetEnableAll returns the EnableAll field if non-nil, zero value otherwise.

### GetEnableAllOk

`func (o *AdmingroupDnsShowCommands) GetEnableAllOk() (*bool, bool)`

GetEnableAllOk returns a tuple with the EnableAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAll

`func (o *AdmingroupDnsShowCommands) SetEnableAll(v bool)`

SetEnableAll sets EnableAll field to given value.

### HasEnableAll

`func (o *AdmingroupDnsShowCommands) HasEnableAll() bool`

HasEnableAll returns a boolean if a field has been set.

### GetDisableAll

`func (o *AdmingroupDnsShowCommands) GetDisableAll() bool`

GetDisableAll returns the DisableAll field if non-nil, zero value otherwise.

### GetDisableAllOk

`func (o *AdmingroupDnsShowCommands) GetDisableAllOk() (*bool, bool)`

GetDisableAllOk returns a tuple with the DisableAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAll

`func (o *AdmingroupDnsShowCommands) SetDisableAll(v bool)`

SetDisableAll sets DisableAll field to given value.

### HasDisableAll

`func (o *AdmingroupDnsShowCommands) HasDisableAll() bool`

HasDisableAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


