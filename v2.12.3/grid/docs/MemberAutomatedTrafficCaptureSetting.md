# MemberAutomatedTrafficCaptureSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrafficCaptureEnable** | Pointer to **bool** | Enable automated traffic capture based on monitoring thresholds. | [optional] 
**Destination** | Pointer to **string** | Destination of traffic capture files. Save traffic capture locally or upload to remote server using FTP or SCP. | [optional] 
**Duration** | Pointer to **int64** | The time interval on which traffic will be captured(in sec). | [optional] 
**IncludeSupportBundle** | Pointer to **bool** | Enable automatic download for support bundle. | [optional] 
**KeepLocalCopy** | Pointer to **bool** | Save traffic capture files locally. | [optional] 
**DestinationHost** | Pointer to **string** | IP Address of the destination host. | [optional] 
**TrafficCaptureDirectory** | Pointer to **string** | Directory to store the traffic capture files on the remote server. | [optional] 
**SupportBundleDirectory** | Pointer to **string** | Directory to store the support bundle on the remote server. | [optional] 
**Username** | Pointer to **string** | User name for accessing the FTP/SCP server. | [optional] 
**Password** | Pointer to **string** | Password for accessing the FTP/SCP server. This field is not readable. | [optional] 

## Methods

### NewMemberAutomatedTrafficCaptureSetting

`func NewMemberAutomatedTrafficCaptureSetting() *MemberAutomatedTrafficCaptureSetting`

NewMemberAutomatedTrafficCaptureSetting instantiates a new MemberAutomatedTrafficCaptureSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberAutomatedTrafficCaptureSettingWithDefaults

`func NewMemberAutomatedTrafficCaptureSettingWithDefaults() *MemberAutomatedTrafficCaptureSetting`

NewMemberAutomatedTrafficCaptureSettingWithDefaults instantiates a new MemberAutomatedTrafficCaptureSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrafficCaptureEnable

`func (o *MemberAutomatedTrafficCaptureSetting) GetTrafficCaptureEnable() bool`

GetTrafficCaptureEnable returns the TrafficCaptureEnable field if non-nil, zero value otherwise.

### GetTrafficCaptureEnableOk

`func (o *MemberAutomatedTrafficCaptureSetting) GetTrafficCaptureEnableOk() (*bool, bool)`

GetTrafficCaptureEnableOk returns a tuple with the TrafficCaptureEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficCaptureEnable

`func (o *MemberAutomatedTrafficCaptureSetting) SetTrafficCaptureEnable(v bool)`

SetTrafficCaptureEnable sets TrafficCaptureEnable field to given value.

### HasTrafficCaptureEnable

`func (o *MemberAutomatedTrafficCaptureSetting) HasTrafficCaptureEnable() bool`

HasTrafficCaptureEnable returns a boolean if a field has been set.

### GetDestination

`func (o *MemberAutomatedTrafficCaptureSetting) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *MemberAutomatedTrafficCaptureSetting) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *MemberAutomatedTrafficCaptureSetting) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *MemberAutomatedTrafficCaptureSetting) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDuration

`func (o *MemberAutomatedTrafficCaptureSetting) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *MemberAutomatedTrafficCaptureSetting) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *MemberAutomatedTrafficCaptureSetting) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *MemberAutomatedTrafficCaptureSetting) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetIncludeSupportBundle

`func (o *MemberAutomatedTrafficCaptureSetting) GetIncludeSupportBundle() bool`

GetIncludeSupportBundle returns the IncludeSupportBundle field if non-nil, zero value otherwise.

### GetIncludeSupportBundleOk

`func (o *MemberAutomatedTrafficCaptureSetting) GetIncludeSupportBundleOk() (*bool, bool)`

GetIncludeSupportBundleOk returns a tuple with the IncludeSupportBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSupportBundle

`func (o *MemberAutomatedTrafficCaptureSetting) SetIncludeSupportBundle(v bool)`

SetIncludeSupportBundle sets IncludeSupportBundle field to given value.

### HasIncludeSupportBundle

`func (o *MemberAutomatedTrafficCaptureSetting) HasIncludeSupportBundle() bool`

HasIncludeSupportBundle returns a boolean if a field has been set.

### GetKeepLocalCopy

`func (o *MemberAutomatedTrafficCaptureSetting) GetKeepLocalCopy() bool`

GetKeepLocalCopy returns the KeepLocalCopy field if non-nil, zero value otherwise.

### GetKeepLocalCopyOk

`func (o *MemberAutomatedTrafficCaptureSetting) GetKeepLocalCopyOk() (*bool, bool)`

GetKeepLocalCopyOk returns a tuple with the KeepLocalCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepLocalCopy

`func (o *MemberAutomatedTrafficCaptureSetting) SetKeepLocalCopy(v bool)`

SetKeepLocalCopy sets KeepLocalCopy field to given value.

### HasKeepLocalCopy

`func (o *MemberAutomatedTrafficCaptureSetting) HasKeepLocalCopy() bool`

HasKeepLocalCopy returns a boolean if a field has been set.

### GetDestinationHost

`func (o *MemberAutomatedTrafficCaptureSetting) GetDestinationHost() string`

GetDestinationHost returns the DestinationHost field if non-nil, zero value otherwise.

### GetDestinationHostOk

`func (o *MemberAutomatedTrafficCaptureSetting) GetDestinationHostOk() (*string, bool)`

GetDestinationHostOk returns a tuple with the DestinationHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationHost

`func (o *MemberAutomatedTrafficCaptureSetting) SetDestinationHost(v string)`

SetDestinationHost sets DestinationHost field to given value.

### HasDestinationHost

`func (o *MemberAutomatedTrafficCaptureSetting) HasDestinationHost() bool`

HasDestinationHost returns a boolean if a field has been set.

### GetTrafficCaptureDirectory

`func (o *MemberAutomatedTrafficCaptureSetting) GetTrafficCaptureDirectory() string`

GetTrafficCaptureDirectory returns the TrafficCaptureDirectory field if non-nil, zero value otherwise.

### GetTrafficCaptureDirectoryOk

`func (o *MemberAutomatedTrafficCaptureSetting) GetTrafficCaptureDirectoryOk() (*string, bool)`

GetTrafficCaptureDirectoryOk returns a tuple with the TrafficCaptureDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficCaptureDirectory

`func (o *MemberAutomatedTrafficCaptureSetting) SetTrafficCaptureDirectory(v string)`

SetTrafficCaptureDirectory sets TrafficCaptureDirectory field to given value.

### HasTrafficCaptureDirectory

`func (o *MemberAutomatedTrafficCaptureSetting) HasTrafficCaptureDirectory() bool`

HasTrafficCaptureDirectory returns a boolean if a field has been set.

### GetSupportBundleDirectory

`func (o *MemberAutomatedTrafficCaptureSetting) GetSupportBundleDirectory() string`

GetSupportBundleDirectory returns the SupportBundleDirectory field if non-nil, zero value otherwise.

### GetSupportBundleDirectoryOk

`func (o *MemberAutomatedTrafficCaptureSetting) GetSupportBundleDirectoryOk() (*string, bool)`

GetSupportBundleDirectoryOk returns a tuple with the SupportBundleDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportBundleDirectory

`func (o *MemberAutomatedTrafficCaptureSetting) SetSupportBundleDirectory(v string)`

SetSupportBundleDirectory sets SupportBundleDirectory field to given value.

### HasSupportBundleDirectory

`func (o *MemberAutomatedTrafficCaptureSetting) HasSupportBundleDirectory() bool`

HasSupportBundleDirectory returns a boolean if a field has been set.

### GetUsername

`func (o *MemberAutomatedTrafficCaptureSetting) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MemberAutomatedTrafficCaptureSetting) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MemberAutomatedTrafficCaptureSetting) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *MemberAutomatedTrafficCaptureSetting) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *MemberAutomatedTrafficCaptureSetting) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MemberAutomatedTrafficCaptureSetting) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MemberAutomatedTrafficCaptureSetting) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MemberAutomatedTrafficCaptureSetting) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


