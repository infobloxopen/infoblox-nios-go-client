# Upgradestatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AllowDistribution** | Pointer to **bool** | Determines if distribution is allowed for the Grid. | [optional] [readonly] 
**AllowDistributionScheduling** | Pointer to **bool** | Determines if distribution scheduling is allowed. | [optional] [readonly] 
**AllowUpgrade** | Pointer to **bool** | Determines if upgrade is allowed for the Grid. | [optional] [readonly] 
**AllowUpgradeCancel** | Pointer to **bool** | Determines if the Grid is allowed to cancel an upgrade. | [optional] [readonly] 
**AllowUpgradePause** | Pointer to **bool** | Determines if the Grid is allowed to pause an upgrade. | [optional] [readonly] 
**AllowUpgradeResume** | Pointer to **bool** | Determines if the Grid is allowed to resume an upgrade. | [optional] [readonly] 
**AllowUpgradeScheduling** | Pointer to **bool** | Determine if the Grid is allowed to schedule an upgrade. | [optional] [readonly] 
**AllowUpgradeTest** | Pointer to **bool** | Determines if the Grid is allowed to test an upgrade. | [optional] [readonly] 
**AllowUpload** | Pointer to **bool** | Determine if the Grid is allowed to upload a build. | [optional] [readonly] 
**AlternateVersion** | Pointer to **string** | The alternative version. | [optional] [readonly] 
**Comment** | Pointer to **string** | Comment in readable format for an upgrade group a or virtual node. | [optional] [readonly] 
**CurrentVersion** | Pointer to **string** | The current version. | [optional] [readonly] 
**CurrentVersionSummary** | Pointer to **string** | Current version summary for the &#39;type&#39; requested. This field can be requested for the Grid, a certain group that has virtual nodes as subelements, or for the overall group status. | [optional] [readonly] 
**DistributionScheduleActive** | Pointer to **bool** | Determines if the distribution schedule is active for the Grid. | [optional] [readonly] 
**DistributionScheduleTime** | Pointer to **int64** | The Grid master distribution schedule time. | [optional] [readonly] 
**DistributionState** | Pointer to **string** | The current state of distribution process. | [optional] [readonly] 
**DistributionVersion** | Pointer to **string** | The version that is distributed. | [optional] [readonly] 
**DistributionVersionSummary** | Pointer to **string** | Distribution version summary for the &#39;type&#39; requested. This field can be requested for the Grid, a certain group that has virtual nodes as subelements, or for the overall group status. | [optional] [readonly] 
**ElementStatus** | Pointer to **string** | The status of a certain element with regards to the type requested. | [optional] [readonly] 
**GridState** | Pointer to **string** | The state of the Grid. | [optional] [readonly] 
**GroupState** | Pointer to **string** | The state of a group. | [optional] [readonly] 
**HaStatus** | Pointer to **string** | Status of the HA pair. | [optional] [readonly] 
**Hotfixes** | Pointer to [**[]UpgradestatusHotfixes**](UpgradestatusHotfixes.md) | The list of hotfixes. | [optional] [readonly] 
**Ipv4Address** | Pointer to **string** | The IPv4 Address of virtual node or physical one. | [optional] [readonly] 
**Ipv6Address** | Pointer to **string** | The IPv6 Address of virtual node or physical one. | [optional] [readonly] 
**Member** | Pointer to **string** | Member that participates in the upgrade process. | [optional] [readonly] 
**Message** | Pointer to **string** | The Grid message. | [optional] [readonly] 
**PnodeRole** | Pointer to **string** | Status of the physical node in the HA pair. | [optional] [readonly] 
**Reverted** | Pointer to **bool** | Determines if the upgrade process is reverted. | [optional] [readonly] 
**StatusTime** | Pointer to **int64** | The status time. | [optional] [readonly] 
**StatusValue** | Pointer to **string** | Status of a certain group, virtual node or physical node. | [optional] [readonly] 
**StatusValueUpdateTime** | Pointer to **int64** | Timestamp of when the status was updated. | [optional] [readonly] 
**Steps** | Pointer to [**[]UpgradestatusSteps**](UpgradestatusSteps.md) | The list of upgrade process steps. | [optional] [readonly] 
**StepsCompleted** | Pointer to **int32** | The number of steps done. | [optional] [readonly] 
**StepsTotal** | Pointer to **int32** | Total number steps in the upgrade process. | [optional] [readonly] 
**SubelementType** | Pointer to **string** | The type of subelements to be requested. If &#39;type&#39; is &#39;GROUP&#39;, or &#39;VNODE&#39;, then &#39;upgrade_group&#39; or &#39;member&#39; should have proper values for an operation to return data specific for the values passed. Otherwise, overall data is returned for every group or physical node. | [optional] [readonly] 
**SubelementsCompleted** | Pointer to **int32** | Number of subelements that have accomplished an upgrade. | [optional] [readonly] 
**SubelementsStatus** | Pointer to **[]string** | The upgrade process information of subelements. | [optional] 
**SubelementsTotal** | Pointer to **int32** | Number of subelements number in a certain group, virtual node, or the Grid. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of upper level elements to be requested. | [optional] [readonly] 
**UpgradeGroup** | Pointer to **string** | Upgrade group that participates in the upgrade process. | [optional] [readonly] 
**UpgradeScheduleActive** | Pointer to **bool** | Determines if the upgrade schedule is active. | [optional] [readonly] 
**UpgradeState** | Pointer to **string** | The upgrade state of the Grid. | [optional] [readonly] 
**UpgradeTestStatus** | Pointer to **string** | The upgrade test status of the Grid. | [optional] [readonly] 
**UploadVersion** | Pointer to **string** | The version that is uploaded. | [optional] [readonly] 
**UploadVersionSummary** | Pointer to **string** | Upload version summary for the &#39;type&#39; requested. This field can be requested for the Grid, a certain group that has virtual nodes as subelements, or overall group status. | [optional] [readonly] 

## Methods

### NewUpgradestatus

`func NewUpgradestatus() *Upgradestatus`

NewUpgradestatus instantiates a new Upgradestatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradestatusWithDefaults

`func NewUpgradestatusWithDefaults() *Upgradestatus`

NewUpgradestatusWithDefaults instantiates a new Upgradestatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Upgradestatus) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Upgradestatus) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Upgradestatus) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Upgradestatus) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllowDistribution

`func (o *Upgradestatus) GetAllowDistribution() bool`

GetAllowDistribution returns the AllowDistribution field if non-nil, zero value otherwise.

### GetAllowDistributionOk

`func (o *Upgradestatus) GetAllowDistributionOk() (*bool, bool)`

GetAllowDistributionOk returns a tuple with the AllowDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDistribution

`func (o *Upgradestatus) SetAllowDistribution(v bool)`

SetAllowDistribution sets AllowDistribution field to given value.

### HasAllowDistribution

`func (o *Upgradestatus) HasAllowDistribution() bool`

HasAllowDistribution returns a boolean if a field has been set.

### GetAllowDistributionScheduling

`func (o *Upgradestatus) GetAllowDistributionScheduling() bool`

GetAllowDistributionScheduling returns the AllowDistributionScheduling field if non-nil, zero value otherwise.

### GetAllowDistributionSchedulingOk

`func (o *Upgradestatus) GetAllowDistributionSchedulingOk() (*bool, bool)`

GetAllowDistributionSchedulingOk returns a tuple with the AllowDistributionScheduling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDistributionScheduling

`func (o *Upgradestatus) SetAllowDistributionScheduling(v bool)`

SetAllowDistributionScheduling sets AllowDistributionScheduling field to given value.

### HasAllowDistributionScheduling

`func (o *Upgradestatus) HasAllowDistributionScheduling() bool`

HasAllowDistributionScheduling returns a boolean if a field has been set.

### GetAllowUpgrade

`func (o *Upgradestatus) GetAllowUpgrade() bool`

GetAllowUpgrade returns the AllowUpgrade field if non-nil, zero value otherwise.

### GetAllowUpgradeOk

`func (o *Upgradestatus) GetAllowUpgradeOk() (*bool, bool)`

GetAllowUpgradeOk returns a tuple with the AllowUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUpgrade

`func (o *Upgradestatus) SetAllowUpgrade(v bool)`

SetAllowUpgrade sets AllowUpgrade field to given value.

### HasAllowUpgrade

`func (o *Upgradestatus) HasAllowUpgrade() bool`

HasAllowUpgrade returns a boolean if a field has been set.

### GetAllowUpgradeCancel

`func (o *Upgradestatus) GetAllowUpgradeCancel() bool`

GetAllowUpgradeCancel returns the AllowUpgradeCancel field if non-nil, zero value otherwise.

### GetAllowUpgradeCancelOk

`func (o *Upgradestatus) GetAllowUpgradeCancelOk() (*bool, bool)`

GetAllowUpgradeCancelOk returns a tuple with the AllowUpgradeCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUpgradeCancel

`func (o *Upgradestatus) SetAllowUpgradeCancel(v bool)`

SetAllowUpgradeCancel sets AllowUpgradeCancel field to given value.

### HasAllowUpgradeCancel

`func (o *Upgradestatus) HasAllowUpgradeCancel() bool`

HasAllowUpgradeCancel returns a boolean if a field has been set.

### GetAllowUpgradePause

`func (o *Upgradestatus) GetAllowUpgradePause() bool`

GetAllowUpgradePause returns the AllowUpgradePause field if non-nil, zero value otherwise.

### GetAllowUpgradePauseOk

`func (o *Upgradestatus) GetAllowUpgradePauseOk() (*bool, bool)`

GetAllowUpgradePauseOk returns a tuple with the AllowUpgradePause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUpgradePause

`func (o *Upgradestatus) SetAllowUpgradePause(v bool)`

SetAllowUpgradePause sets AllowUpgradePause field to given value.

### HasAllowUpgradePause

`func (o *Upgradestatus) HasAllowUpgradePause() bool`

HasAllowUpgradePause returns a boolean if a field has been set.

### GetAllowUpgradeResume

`func (o *Upgradestatus) GetAllowUpgradeResume() bool`

GetAllowUpgradeResume returns the AllowUpgradeResume field if non-nil, zero value otherwise.

### GetAllowUpgradeResumeOk

`func (o *Upgradestatus) GetAllowUpgradeResumeOk() (*bool, bool)`

GetAllowUpgradeResumeOk returns a tuple with the AllowUpgradeResume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUpgradeResume

`func (o *Upgradestatus) SetAllowUpgradeResume(v bool)`

SetAllowUpgradeResume sets AllowUpgradeResume field to given value.

### HasAllowUpgradeResume

`func (o *Upgradestatus) HasAllowUpgradeResume() bool`

HasAllowUpgradeResume returns a boolean if a field has been set.

### GetAllowUpgradeScheduling

`func (o *Upgradestatus) GetAllowUpgradeScheduling() bool`

GetAllowUpgradeScheduling returns the AllowUpgradeScheduling field if non-nil, zero value otherwise.

### GetAllowUpgradeSchedulingOk

`func (o *Upgradestatus) GetAllowUpgradeSchedulingOk() (*bool, bool)`

GetAllowUpgradeSchedulingOk returns a tuple with the AllowUpgradeScheduling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUpgradeScheduling

`func (o *Upgradestatus) SetAllowUpgradeScheduling(v bool)`

SetAllowUpgradeScheduling sets AllowUpgradeScheduling field to given value.

### HasAllowUpgradeScheduling

`func (o *Upgradestatus) HasAllowUpgradeScheduling() bool`

HasAllowUpgradeScheduling returns a boolean if a field has been set.

### GetAllowUpgradeTest

`func (o *Upgradestatus) GetAllowUpgradeTest() bool`

GetAllowUpgradeTest returns the AllowUpgradeTest field if non-nil, zero value otherwise.

### GetAllowUpgradeTestOk

`func (o *Upgradestatus) GetAllowUpgradeTestOk() (*bool, bool)`

GetAllowUpgradeTestOk returns a tuple with the AllowUpgradeTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUpgradeTest

`func (o *Upgradestatus) SetAllowUpgradeTest(v bool)`

SetAllowUpgradeTest sets AllowUpgradeTest field to given value.

### HasAllowUpgradeTest

`func (o *Upgradestatus) HasAllowUpgradeTest() bool`

HasAllowUpgradeTest returns a boolean if a field has been set.

### GetAllowUpload

`func (o *Upgradestatus) GetAllowUpload() bool`

GetAllowUpload returns the AllowUpload field if non-nil, zero value otherwise.

### GetAllowUploadOk

`func (o *Upgradestatus) GetAllowUploadOk() (*bool, bool)`

GetAllowUploadOk returns a tuple with the AllowUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUpload

`func (o *Upgradestatus) SetAllowUpload(v bool)`

SetAllowUpload sets AllowUpload field to given value.

### HasAllowUpload

`func (o *Upgradestatus) HasAllowUpload() bool`

HasAllowUpload returns a boolean if a field has been set.

### GetAlternateVersion

`func (o *Upgradestatus) GetAlternateVersion() string`

GetAlternateVersion returns the AlternateVersion field if non-nil, zero value otherwise.

### GetAlternateVersionOk

`func (o *Upgradestatus) GetAlternateVersionOk() (*string, bool)`

GetAlternateVersionOk returns a tuple with the AlternateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateVersion

`func (o *Upgradestatus) SetAlternateVersion(v string)`

SetAlternateVersion sets AlternateVersion field to given value.

### HasAlternateVersion

`func (o *Upgradestatus) HasAlternateVersion() bool`

HasAlternateVersion returns a boolean if a field has been set.

### GetComment

`func (o *Upgradestatus) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Upgradestatus) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Upgradestatus) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Upgradestatus) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *Upgradestatus) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *Upgradestatus) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *Upgradestatus) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *Upgradestatus) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetCurrentVersionSummary

`func (o *Upgradestatus) GetCurrentVersionSummary() string`

GetCurrentVersionSummary returns the CurrentVersionSummary field if non-nil, zero value otherwise.

### GetCurrentVersionSummaryOk

`func (o *Upgradestatus) GetCurrentVersionSummaryOk() (*string, bool)`

GetCurrentVersionSummaryOk returns a tuple with the CurrentVersionSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersionSummary

`func (o *Upgradestatus) SetCurrentVersionSummary(v string)`

SetCurrentVersionSummary sets CurrentVersionSummary field to given value.

### HasCurrentVersionSummary

`func (o *Upgradestatus) HasCurrentVersionSummary() bool`

HasCurrentVersionSummary returns a boolean if a field has been set.

### GetDistributionScheduleActive

`func (o *Upgradestatus) GetDistributionScheduleActive() bool`

GetDistributionScheduleActive returns the DistributionScheduleActive field if non-nil, zero value otherwise.

### GetDistributionScheduleActiveOk

`func (o *Upgradestatus) GetDistributionScheduleActiveOk() (*bool, bool)`

GetDistributionScheduleActiveOk returns a tuple with the DistributionScheduleActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionScheduleActive

`func (o *Upgradestatus) SetDistributionScheduleActive(v bool)`

SetDistributionScheduleActive sets DistributionScheduleActive field to given value.

### HasDistributionScheduleActive

`func (o *Upgradestatus) HasDistributionScheduleActive() bool`

HasDistributionScheduleActive returns a boolean if a field has been set.

### GetDistributionScheduleTime

`func (o *Upgradestatus) GetDistributionScheduleTime() int64`

GetDistributionScheduleTime returns the DistributionScheduleTime field if non-nil, zero value otherwise.

### GetDistributionScheduleTimeOk

`func (o *Upgradestatus) GetDistributionScheduleTimeOk() (*int64, bool)`

GetDistributionScheduleTimeOk returns a tuple with the DistributionScheduleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionScheduleTime

`func (o *Upgradestatus) SetDistributionScheduleTime(v int64)`

SetDistributionScheduleTime sets DistributionScheduleTime field to given value.

### HasDistributionScheduleTime

`func (o *Upgradestatus) HasDistributionScheduleTime() bool`

HasDistributionScheduleTime returns a boolean if a field has been set.

### GetDistributionState

`func (o *Upgradestatus) GetDistributionState() string`

GetDistributionState returns the DistributionState field if non-nil, zero value otherwise.

### GetDistributionStateOk

`func (o *Upgradestatus) GetDistributionStateOk() (*string, bool)`

GetDistributionStateOk returns a tuple with the DistributionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionState

`func (o *Upgradestatus) SetDistributionState(v string)`

SetDistributionState sets DistributionState field to given value.

### HasDistributionState

`func (o *Upgradestatus) HasDistributionState() bool`

HasDistributionState returns a boolean if a field has been set.

### GetDistributionVersion

`func (o *Upgradestatus) GetDistributionVersion() string`

GetDistributionVersion returns the DistributionVersion field if non-nil, zero value otherwise.

### GetDistributionVersionOk

`func (o *Upgradestatus) GetDistributionVersionOk() (*string, bool)`

GetDistributionVersionOk returns a tuple with the DistributionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionVersion

`func (o *Upgradestatus) SetDistributionVersion(v string)`

SetDistributionVersion sets DistributionVersion field to given value.

### HasDistributionVersion

`func (o *Upgradestatus) HasDistributionVersion() bool`

HasDistributionVersion returns a boolean if a field has been set.

### GetDistributionVersionSummary

`func (o *Upgradestatus) GetDistributionVersionSummary() string`

GetDistributionVersionSummary returns the DistributionVersionSummary field if non-nil, zero value otherwise.

### GetDistributionVersionSummaryOk

`func (o *Upgradestatus) GetDistributionVersionSummaryOk() (*string, bool)`

GetDistributionVersionSummaryOk returns a tuple with the DistributionVersionSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionVersionSummary

`func (o *Upgradestatus) SetDistributionVersionSummary(v string)`

SetDistributionVersionSummary sets DistributionVersionSummary field to given value.

### HasDistributionVersionSummary

`func (o *Upgradestatus) HasDistributionVersionSummary() bool`

HasDistributionVersionSummary returns a boolean if a field has been set.

### GetElementStatus

`func (o *Upgradestatus) GetElementStatus() string`

GetElementStatus returns the ElementStatus field if non-nil, zero value otherwise.

### GetElementStatusOk

`func (o *Upgradestatus) GetElementStatusOk() (*string, bool)`

GetElementStatusOk returns a tuple with the ElementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementStatus

`func (o *Upgradestatus) SetElementStatus(v string)`

SetElementStatus sets ElementStatus field to given value.

### HasElementStatus

`func (o *Upgradestatus) HasElementStatus() bool`

HasElementStatus returns a boolean if a field has been set.

### GetGridState

`func (o *Upgradestatus) GetGridState() string`

GetGridState returns the GridState field if non-nil, zero value otherwise.

### GetGridStateOk

`func (o *Upgradestatus) GetGridStateOk() (*string, bool)`

GetGridStateOk returns a tuple with the GridState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridState

`func (o *Upgradestatus) SetGridState(v string)`

SetGridState sets GridState field to given value.

### HasGridState

`func (o *Upgradestatus) HasGridState() bool`

HasGridState returns a boolean if a field has been set.

### GetGroupState

`func (o *Upgradestatus) GetGroupState() string`

GetGroupState returns the GroupState field if non-nil, zero value otherwise.

### GetGroupStateOk

`func (o *Upgradestatus) GetGroupStateOk() (*string, bool)`

GetGroupStateOk returns a tuple with the GroupState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupState

`func (o *Upgradestatus) SetGroupState(v string)`

SetGroupState sets GroupState field to given value.

### HasGroupState

`func (o *Upgradestatus) HasGroupState() bool`

HasGroupState returns a boolean if a field has been set.

### GetHaStatus

`func (o *Upgradestatus) GetHaStatus() string`

GetHaStatus returns the HaStatus field if non-nil, zero value otherwise.

### GetHaStatusOk

`func (o *Upgradestatus) GetHaStatusOk() (*string, bool)`

GetHaStatusOk returns a tuple with the HaStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaStatus

`func (o *Upgradestatus) SetHaStatus(v string)`

SetHaStatus sets HaStatus field to given value.

### HasHaStatus

`func (o *Upgradestatus) HasHaStatus() bool`

HasHaStatus returns a boolean if a field has been set.

### GetHotfixes

`func (o *Upgradestatus) GetHotfixes() []UpgradestatusHotfixes`

GetHotfixes returns the Hotfixes field if non-nil, zero value otherwise.

### GetHotfixesOk

`func (o *Upgradestatus) GetHotfixesOk() (*[]UpgradestatusHotfixes, bool)`

GetHotfixesOk returns a tuple with the Hotfixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotfixes

`func (o *Upgradestatus) SetHotfixes(v []UpgradestatusHotfixes)`

SetHotfixes sets Hotfixes field to given value.

### HasHotfixes

`func (o *Upgradestatus) HasHotfixes() bool`

HasHotfixes returns a boolean if a field has been set.

### GetIpv4Address

`func (o *Upgradestatus) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *Upgradestatus) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *Upgradestatus) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *Upgradestatus) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetIpv6Address

`func (o *Upgradestatus) GetIpv6Address() string`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *Upgradestatus) GetIpv6AddressOk() (*string, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *Upgradestatus) SetIpv6Address(v string)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *Upgradestatus) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### GetMember

`func (o *Upgradestatus) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *Upgradestatus) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *Upgradestatus) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *Upgradestatus) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetMessage

`func (o *Upgradestatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Upgradestatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Upgradestatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Upgradestatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPnodeRole

`func (o *Upgradestatus) GetPnodeRole() string`

GetPnodeRole returns the PnodeRole field if non-nil, zero value otherwise.

### GetPnodeRoleOk

`func (o *Upgradestatus) GetPnodeRoleOk() (*string, bool)`

GetPnodeRoleOk returns a tuple with the PnodeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPnodeRole

`func (o *Upgradestatus) SetPnodeRole(v string)`

SetPnodeRole sets PnodeRole field to given value.

### HasPnodeRole

`func (o *Upgradestatus) HasPnodeRole() bool`

HasPnodeRole returns a boolean if a field has been set.

### GetReverted

`func (o *Upgradestatus) GetReverted() bool`

GetReverted returns the Reverted field if non-nil, zero value otherwise.

### GetRevertedOk

`func (o *Upgradestatus) GetRevertedOk() (*bool, bool)`

GetRevertedOk returns a tuple with the Reverted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverted

`func (o *Upgradestatus) SetReverted(v bool)`

SetReverted sets Reverted field to given value.

### HasReverted

`func (o *Upgradestatus) HasReverted() bool`

HasReverted returns a boolean if a field has been set.

### GetStatusTime

`func (o *Upgradestatus) GetStatusTime() int64`

GetStatusTime returns the StatusTime field if non-nil, zero value otherwise.

### GetStatusTimeOk

`func (o *Upgradestatus) GetStatusTimeOk() (*int64, bool)`

GetStatusTimeOk returns a tuple with the StatusTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusTime

`func (o *Upgradestatus) SetStatusTime(v int64)`

SetStatusTime sets StatusTime field to given value.

### HasStatusTime

`func (o *Upgradestatus) HasStatusTime() bool`

HasStatusTime returns a boolean if a field has been set.

### GetStatusValue

`func (o *Upgradestatus) GetStatusValue() string`

GetStatusValue returns the StatusValue field if non-nil, zero value otherwise.

### GetStatusValueOk

`func (o *Upgradestatus) GetStatusValueOk() (*string, bool)`

GetStatusValueOk returns a tuple with the StatusValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusValue

`func (o *Upgradestatus) SetStatusValue(v string)`

SetStatusValue sets StatusValue field to given value.

### HasStatusValue

`func (o *Upgradestatus) HasStatusValue() bool`

HasStatusValue returns a boolean if a field has been set.

### GetStatusValueUpdateTime

`func (o *Upgradestatus) GetStatusValueUpdateTime() int64`

GetStatusValueUpdateTime returns the StatusValueUpdateTime field if non-nil, zero value otherwise.

### GetStatusValueUpdateTimeOk

`func (o *Upgradestatus) GetStatusValueUpdateTimeOk() (*int64, bool)`

GetStatusValueUpdateTimeOk returns a tuple with the StatusValueUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusValueUpdateTime

`func (o *Upgradestatus) SetStatusValueUpdateTime(v int64)`

SetStatusValueUpdateTime sets StatusValueUpdateTime field to given value.

### HasStatusValueUpdateTime

`func (o *Upgradestatus) HasStatusValueUpdateTime() bool`

HasStatusValueUpdateTime returns a boolean if a field has been set.

### GetSteps

`func (o *Upgradestatus) GetSteps() []UpgradestatusSteps`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *Upgradestatus) GetStepsOk() (*[]UpgradestatusSteps, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *Upgradestatus) SetSteps(v []UpgradestatusSteps)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *Upgradestatus) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetStepsCompleted

`func (o *Upgradestatus) GetStepsCompleted() int32`

GetStepsCompleted returns the StepsCompleted field if non-nil, zero value otherwise.

### GetStepsCompletedOk

`func (o *Upgradestatus) GetStepsCompletedOk() (*int32, bool)`

GetStepsCompletedOk returns a tuple with the StepsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsCompleted

`func (o *Upgradestatus) SetStepsCompleted(v int32)`

SetStepsCompleted sets StepsCompleted field to given value.

### HasStepsCompleted

`func (o *Upgradestatus) HasStepsCompleted() bool`

HasStepsCompleted returns a boolean if a field has been set.

### GetStepsTotal

`func (o *Upgradestatus) GetStepsTotal() int32`

GetStepsTotal returns the StepsTotal field if non-nil, zero value otherwise.

### GetStepsTotalOk

`func (o *Upgradestatus) GetStepsTotalOk() (*int32, bool)`

GetStepsTotalOk returns a tuple with the StepsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsTotal

`func (o *Upgradestatus) SetStepsTotal(v int32)`

SetStepsTotal sets StepsTotal field to given value.

### HasStepsTotal

`func (o *Upgradestatus) HasStepsTotal() bool`

HasStepsTotal returns a boolean if a field has been set.

### GetSubelementType

`func (o *Upgradestatus) GetSubelementType() string`

GetSubelementType returns the SubelementType field if non-nil, zero value otherwise.

### GetSubelementTypeOk

`func (o *Upgradestatus) GetSubelementTypeOk() (*string, bool)`

GetSubelementTypeOk returns a tuple with the SubelementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubelementType

`func (o *Upgradestatus) SetSubelementType(v string)`

SetSubelementType sets SubelementType field to given value.

### HasSubelementType

`func (o *Upgradestatus) HasSubelementType() bool`

HasSubelementType returns a boolean if a field has been set.

### GetSubelementsCompleted

`func (o *Upgradestatus) GetSubelementsCompleted() int32`

GetSubelementsCompleted returns the SubelementsCompleted field if non-nil, zero value otherwise.

### GetSubelementsCompletedOk

`func (o *Upgradestatus) GetSubelementsCompletedOk() (*int32, bool)`

GetSubelementsCompletedOk returns a tuple with the SubelementsCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubelementsCompleted

`func (o *Upgradestatus) SetSubelementsCompleted(v int32)`

SetSubelementsCompleted sets SubelementsCompleted field to given value.

### HasSubelementsCompleted

`func (o *Upgradestatus) HasSubelementsCompleted() bool`

HasSubelementsCompleted returns a boolean if a field has been set.

### GetSubelementsStatus

`func (o *Upgradestatus) GetSubelementsStatus() []string`

GetSubelementsStatus returns the SubelementsStatus field if non-nil, zero value otherwise.

### GetSubelementsStatusOk

`func (o *Upgradestatus) GetSubelementsStatusOk() (*[]string, bool)`

GetSubelementsStatusOk returns a tuple with the SubelementsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubelementsStatus

`func (o *Upgradestatus) SetSubelementsStatus(v []string)`

SetSubelementsStatus sets SubelementsStatus field to given value.

### HasSubelementsStatus

`func (o *Upgradestatus) HasSubelementsStatus() bool`

HasSubelementsStatus returns a boolean if a field has been set.

### GetSubelementsTotal

`func (o *Upgradestatus) GetSubelementsTotal() int32`

GetSubelementsTotal returns the SubelementsTotal field if non-nil, zero value otherwise.

### GetSubelementsTotalOk

`func (o *Upgradestatus) GetSubelementsTotalOk() (*int32, bool)`

GetSubelementsTotalOk returns a tuple with the SubelementsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubelementsTotal

`func (o *Upgradestatus) SetSubelementsTotal(v int32)`

SetSubelementsTotal sets SubelementsTotal field to given value.

### HasSubelementsTotal

`func (o *Upgradestatus) HasSubelementsTotal() bool`

HasSubelementsTotal returns a boolean if a field has been set.

### GetType

`func (o *Upgradestatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Upgradestatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Upgradestatus) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Upgradestatus) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpgradeGroup

`func (o *Upgradestatus) GetUpgradeGroup() string`

GetUpgradeGroup returns the UpgradeGroup field if non-nil, zero value otherwise.

### GetUpgradeGroupOk

`func (o *Upgradestatus) GetUpgradeGroupOk() (*string, bool)`

GetUpgradeGroupOk returns a tuple with the UpgradeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeGroup

`func (o *Upgradestatus) SetUpgradeGroup(v string)`

SetUpgradeGroup sets UpgradeGroup field to given value.

### HasUpgradeGroup

`func (o *Upgradestatus) HasUpgradeGroup() bool`

HasUpgradeGroup returns a boolean if a field has been set.

### GetUpgradeScheduleActive

`func (o *Upgradestatus) GetUpgradeScheduleActive() bool`

GetUpgradeScheduleActive returns the UpgradeScheduleActive field if non-nil, zero value otherwise.

### GetUpgradeScheduleActiveOk

`func (o *Upgradestatus) GetUpgradeScheduleActiveOk() (*bool, bool)`

GetUpgradeScheduleActiveOk returns a tuple with the UpgradeScheduleActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeScheduleActive

`func (o *Upgradestatus) SetUpgradeScheduleActive(v bool)`

SetUpgradeScheduleActive sets UpgradeScheduleActive field to given value.

### HasUpgradeScheduleActive

`func (o *Upgradestatus) HasUpgradeScheduleActive() bool`

HasUpgradeScheduleActive returns a boolean if a field has been set.

### GetUpgradeState

`func (o *Upgradestatus) GetUpgradeState() string`

GetUpgradeState returns the UpgradeState field if non-nil, zero value otherwise.

### GetUpgradeStateOk

`func (o *Upgradestatus) GetUpgradeStateOk() (*string, bool)`

GetUpgradeStateOk returns a tuple with the UpgradeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeState

`func (o *Upgradestatus) SetUpgradeState(v string)`

SetUpgradeState sets UpgradeState field to given value.

### HasUpgradeState

`func (o *Upgradestatus) HasUpgradeState() bool`

HasUpgradeState returns a boolean if a field has been set.

### GetUpgradeTestStatus

`func (o *Upgradestatus) GetUpgradeTestStatus() string`

GetUpgradeTestStatus returns the UpgradeTestStatus field if non-nil, zero value otherwise.

### GetUpgradeTestStatusOk

`func (o *Upgradestatus) GetUpgradeTestStatusOk() (*string, bool)`

GetUpgradeTestStatusOk returns a tuple with the UpgradeTestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeTestStatus

`func (o *Upgradestatus) SetUpgradeTestStatus(v string)`

SetUpgradeTestStatus sets UpgradeTestStatus field to given value.

### HasUpgradeTestStatus

`func (o *Upgradestatus) HasUpgradeTestStatus() bool`

HasUpgradeTestStatus returns a boolean if a field has been set.

### GetUploadVersion

`func (o *Upgradestatus) GetUploadVersion() string`

GetUploadVersion returns the UploadVersion field if non-nil, zero value otherwise.

### GetUploadVersionOk

`func (o *Upgradestatus) GetUploadVersionOk() (*string, bool)`

GetUploadVersionOk returns a tuple with the UploadVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadVersion

`func (o *Upgradestatus) SetUploadVersion(v string)`

SetUploadVersion sets UploadVersion field to given value.

### HasUploadVersion

`func (o *Upgradestatus) HasUploadVersion() bool`

HasUploadVersion returns a boolean if a field has been set.

### GetUploadVersionSummary

`func (o *Upgradestatus) GetUploadVersionSummary() string`

GetUploadVersionSummary returns the UploadVersionSummary field if non-nil, zero value otherwise.

### GetUploadVersionSummaryOk

`func (o *Upgradestatus) GetUploadVersionSummaryOk() (*string, bool)`

GetUploadVersionSummaryOk returns a tuple with the UploadVersionSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadVersionSummary

`func (o *Upgradestatus) SetUploadVersionSummary(v string)`

SetUploadVersionSummary sets UploadVersionSummary field to given value.

### HasUploadVersionSummary

`func (o *Upgradestatus) HasUploadVersionSummary() bool`

HasUploadVersionSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


