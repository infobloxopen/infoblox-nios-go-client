# Captiveportal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AuthnServerGroup** | Pointer to **string** | The authentication server group assigned to this captive portal. | [optional] 
**CompanyName** | Pointer to **string** | The company name that appears in the guest registration page. | [optional] 
**EnableSyslogAuthFailure** | Pointer to **bool** | Determines if authentication failures are logged to syslog or not. | [optional] 
**EnableSyslogAuthSuccess** | Pointer to **bool** | Determines if successful authentications are logged to syslog or not. | [optional] 
**EnableUserType** | Pointer to **string** | The type of user to be enabled for the captive portal. | [optional] 
**Encryption** | Pointer to **string** | The encryption the captive portal uses. | [optional] 
**Files** | Pointer to [**[]CaptiveportalFiles**](CaptiveportalFiles.md) | The list of files associated with the captive portal. | [optional] 
**GuestCustomField1Name** | Pointer to **string** | The name of the custom field that you are adding to the guest registration page. | [optional] 
**GuestCustomField1Required** | Pointer to **bool** | Determines if the custom field is required or not. | [optional] 
**GuestCustomField2Name** | Pointer to **string** | The name of the custom field that you are adding to the guest registration page. | [optional] 
**GuestCustomField2Required** | Pointer to **bool** | Determines if the custom field is required or not. | [optional] 
**GuestCustomField3Name** | Pointer to **string** | The name of the custom field that you are adding to the guest registration page. | [optional] 
**GuestCustomField3Required** | Pointer to **bool** | Determines if the custom field is required or not. | [optional] 
**GuestCustomField4Name** | Pointer to **string** | The name of the custom field that you are adding to the guest registration page. | [optional] 
**GuestCustomField4Required** | Pointer to **bool** | Determines if the custom field is required or not. | [optional] 
**GuestEmailRequired** | Pointer to **bool** | Determines if the email address of the guest is required or not. | [optional] 
**GuestFirstNameRequired** | Pointer to **bool** | Determines if the first name of the guest is required or not. | [optional] 
**GuestLastNameRequired** | Pointer to **bool** | Determines if the last name of the guest is required or not. | [optional] 
**GuestMiddleNameRequired** | Pointer to **bool** | Determines if the middle name of the guest is required or not. | [optional] 
**GuestPhoneRequired** | Pointer to **bool** | Determines if the phone number of the guest is required or not. | [optional] 
**HelpdeskMessage** | Pointer to **string** | The helpdesk message that appears in the guest registration page. | [optional] 
**ListenAddressIp** | Pointer to **string** | Determines the IP address on which the captive portal listens. Valid if listen address type is &#39;IP&#39;. | [optional] 
**ListenAddressType** | Pointer to **string** | Determines the type of the IP address on which the captive portal listens. | [optional] 
**Name** | Pointer to **string** | The hostname of the Grid member that hosts the captive portal. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The network view of the captive portal. | [optional] 
**Port** | Pointer to **int64** | The TCP port used by the Captive Portal service. The port is required when the Captive Portal service is enabled. Valid values are between 1 and 63999. Please note that setting the port number to 80 or 443 might impact performance. | [optional] 
**ServiceEnabled** | Pointer to **bool** | Determines if the captive portal service is enabled or not. | [optional] 
**SyslogAuthFailureLevel** | Pointer to **string** | The syslog level at which authentication failures are logged. | [optional] 
**SyslogAuthSuccessLevel** | Pointer to **string** | The syslog level at which successful authentications are logged. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**WelcomeMessage** | Pointer to **string** | The welcome message that appears in the guest registration page. | [optional] 

## Methods

### NewCaptiveportal

`func NewCaptiveportal() *Captiveportal`

NewCaptiveportal instantiates a new Captiveportal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaptiveportalWithDefaults

`func NewCaptiveportalWithDefaults() *Captiveportal`

NewCaptiveportalWithDefaults instantiates a new Captiveportal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Captiveportal) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Captiveportal) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Captiveportal) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Captiveportal) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAuthnServerGroup

`func (o *Captiveportal) GetAuthnServerGroup() string`

GetAuthnServerGroup returns the AuthnServerGroup field if non-nil, zero value otherwise.

### GetAuthnServerGroupOk

`func (o *Captiveportal) GetAuthnServerGroupOk() (*string, bool)`

GetAuthnServerGroupOk returns a tuple with the AuthnServerGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnServerGroup

`func (o *Captiveportal) SetAuthnServerGroup(v string)`

SetAuthnServerGroup sets AuthnServerGroup field to given value.

### HasAuthnServerGroup

`func (o *Captiveportal) HasAuthnServerGroup() bool`

HasAuthnServerGroup returns a boolean if a field has been set.

### GetCompanyName

`func (o *Captiveportal) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *Captiveportal) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *Captiveportal) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *Captiveportal) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetEnableSyslogAuthFailure

`func (o *Captiveportal) GetEnableSyslogAuthFailure() bool`

GetEnableSyslogAuthFailure returns the EnableSyslogAuthFailure field if non-nil, zero value otherwise.

### GetEnableSyslogAuthFailureOk

`func (o *Captiveportal) GetEnableSyslogAuthFailureOk() (*bool, bool)`

GetEnableSyslogAuthFailureOk returns a tuple with the EnableSyslogAuthFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSyslogAuthFailure

`func (o *Captiveportal) SetEnableSyslogAuthFailure(v bool)`

SetEnableSyslogAuthFailure sets EnableSyslogAuthFailure field to given value.

### HasEnableSyslogAuthFailure

`func (o *Captiveportal) HasEnableSyslogAuthFailure() bool`

HasEnableSyslogAuthFailure returns a boolean if a field has been set.

### GetEnableSyslogAuthSuccess

`func (o *Captiveportal) GetEnableSyslogAuthSuccess() bool`

GetEnableSyslogAuthSuccess returns the EnableSyslogAuthSuccess field if non-nil, zero value otherwise.

### GetEnableSyslogAuthSuccessOk

`func (o *Captiveportal) GetEnableSyslogAuthSuccessOk() (*bool, bool)`

GetEnableSyslogAuthSuccessOk returns a tuple with the EnableSyslogAuthSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSyslogAuthSuccess

`func (o *Captiveportal) SetEnableSyslogAuthSuccess(v bool)`

SetEnableSyslogAuthSuccess sets EnableSyslogAuthSuccess field to given value.

### HasEnableSyslogAuthSuccess

`func (o *Captiveportal) HasEnableSyslogAuthSuccess() bool`

HasEnableSyslogAuthSuccess returns a boolean if a field has been set.

### GetEnableUserType

`func (o *Captiveportal) GetEnableUserType() string`

GetEnableUserType returns the EnableUserType field if non-nil, zero value otherwise.

### GetEnableUserTypeOk

`func (o *Captiveportal) GetEnableUserTypeOk() (*string, bool)`

GetEnableUserTypeOk returns a tuple with the EnableUserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableUserType

`func (o *Captiveportal) SetEnableUserType(v string)`

SetEnableUserType sets EnableUserType field to given value.

### HasEnableUserType

`func (o *Captiveportal) HasEnableUserType() bool`

HasEnableUserType returns a boolean if a field has been set.

### GetEncryption

`func (o *Captiveportal) GetEncryption() string`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *Captiveportal) GetEncryptionOk() (*string, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *Captiveportal) SetEncryption(v string)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *Captiveportal) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetFiles

`func (o *Captiveportal) GetFiles() []CaptiveportalFiles`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *Captiveportal) GetFilesOk() (*[]CaptiveportalFiles, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *Captiveportal) SetFiles(v []CaptiveportalFiles)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *Captiveportal) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetGuestCustomField1Name

`func (o *Captiveportal) GetGuestCustomField1Name() string`

GetGuestCustomField1Name returns the GuestCustomField1Name field if non-nil, zero value otherwise.

### GetGuestCustomField1NameOk

`func (o *Captiveportal) GetGuestCustomField1NameOk() (*string, bool)`

GetGuestCustomField1NameOk returns a tuple with the GuestCustomField1Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCustomField1Name

`func (o *Captiveportal) SetGuestCustomField1Name(v string)`

SetGuestCustomField1Name sets GuestCustomField1Name field to given value.

### HasGuestCustomField1Name

`func (o *Captiveportal) HasGuestCustomField1Name() bool`

HasGuestCustomField1Name returns a boolean if a field has been set.

### GetGuestCustomField1Required

`func (o *Captiveportal) GetGuestCustomField1Required() bool`

GetGuestCustomField1Required returns the GuestCustomField1Required field if non-nil, zero value otherwise.

### GetGuestCustomField1RequiredOk

`func (o *Captiveportal) GetGuestCustomField1RequiredOk() (*bool, bool)`

GetGuestCustomField1RequiredOk returns a tuple with the GuestCustomField1Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCustomField1Required

`func (o *Captiveportal) SetGuestCustomField1Required(v bool)`

SetGuestCustomField1Required sets GuestCustomField1Required field to given value.

### HasGuestCustomField1Required

`func (o *Captiveportal) HasGuestCustomField1Required() bool`

HasGuestCustomField1Required returns a boolean if a field has been set.

### GetGuestCustomField2Name

`func (o *Captiveportal) GetGuestCustomField2Name() string`

GetGuestCustomField2Name returns the GuestCustomField2Name field if non-nil, zero value otherwise.

### GetGuestCustomField2NameOk

`func (o *Captiveportal) GetGuestCustomField2NameOk() (*string, bool)`

GetGuestCustomField2NameOk returns a tuple with the GuestCustomField2Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCustomField2Name

`func (o *Captiveportal) SetGuestCustomField2Name(v string)`

SetGuestCustomField2Name sets GuestCustomField2Name field to given value.

### HasGuestCustomField2Name

`func (o *Captiveportal) HasGuestCustomField2Name() bool`

HasGuestCustomField2Name returns a boolean if a field has been set.

### GetGuestCustomField2Required

`func (o *Captiveportal) GetGuestCustomField2Required() bool`

GetGuestCustomField2Required returns the GuestCustomField2Required field if non-nil, zero value otherwise.

### GetGuestCustomField2RequiredOk

`func (o *Captiveportal) GetGuestCustomField2RequiredOk() (*bool, bool)`

GetGuestCustomField2RequiredOk returns a tuple with the GuestCustomField2Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCustomField2Required

`func (o *Captiveportal) SetGuestCustomField2Required(v bool)`

SetGuestCustomField2Required sets GuestCustomField2Required field to given value.

### HasGuestCustomField2Required

`func (o *Captiveportal) HasGuestCustomField2Required() bool`

HasGuestCustomField2Required returns a boolean if a field has been set.

### GetGuestCustomField3Name

`func (o *Captiveportal) GetGuestCustomField3Name() string`

GetGuestCustomField3Name returns the GuestCustomField3Name field if non-nil, zero value otherwise.

### GetGuestCustomField3NameOk

`func (o *Captiveportal) GetGuestCustomField3NameOk() (*string, bool)`

GetGuestCustomField3NameOk returns a tuple with the GuestCustomField3Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCustomField3Name

`func (o *Captiveportal) SetGuestCustomField3Name(v string)`

SetGuestCustomField3Name sets GuestCustomField3Name field to given value.

### HasGuestCustomField3Name

`func (o *Captiveportal) HasGuestCustomField3Name() bool`

HasGuestCustomField3Name returns a boolean if a field has been set.

### GetGuestCustomField3Required

`func (o *Captiveportal) GetGuestCustomField3Required() bool`

GetGuestCustomField3Required returns the GuestCustomField3Required field if non-nil, zero value otherwise.

### GetGuestCustomField3RequiredOk

`func (o *Captiveportal) GetGuestCustomField3RequiredOk() (*bool, bool)`

GetGuestCustomField3RequiredOk returns a tuple with the GuestCustomField3Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCustomField3Required

`func (o *Captiveportal) SetGuestCustomField3Required(v bool)`

SetGuestCustomField3Required sets GuestCustomField3Required field to given value.

### HasGuestCustomField3Required

`func (o *Captiveportal) HasGuestCustomField3Required() bool`

HasGuestCustomField3Required returns a boolean if a field has been set.

### GetGuestCustomField4Name

`func (o *Captiveportal) GetGuestCustomField4Name() string`

GetGuestCustomField4Name returns the GuestCustomField4Name field if non-nil, zero value otherwise.

### GetGuestCustomField4NameOk

`func (o *Captiveportal) GetGuestCustomField4NameOk() (*string, bool)`

GetGuestCustomField4NameOk returns a tuple with the GuestCustomField4Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCustomField4Name

`func (o *Captiveportal) SetGuestCustomField4Name(v string)`

SetGuestCustomField4Name sets GuestCustomField4Name field to given value.

### HasGuestCustomField4Name

`func (o *Captiveportal) HasGuestCustomField4Name() bool`

HasGuestCustomField4Name returns a boolean if a field has been set.

### GetGuestCustomField4Required

`func (o *Captiveportal) GetGuestCustomField4Required() bool`

GetGuestCustomField4Required returns the GuestCustomField4Required field if non-nil, zero value otherwise.

### GetGuestCustomField4RequiredOk

`func (o *Captiveportal) GetGuestCustomField4RequiredOk() (*bool, bool)`

GetGuestCustomField4RequiredOk returns a tuple with the GuestCustomField4Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCustomField4Required

`func (o *Captiveportal) SetGuestCustomField4Required(v bool)`

SetGuestCustomField4Required sets GuestCustomField4Required field to given value.

### HasGuestCustomField4Required

`func (o *Captiveportal) HasGuestCustomField4Required() bool`

HasGuestCustomField4Required returns a boolean if a field has been set.

### GetGuestEmailRequired

`func (o *Captiveportal) GetGuestEmailRequired() bool`

GetGuestEmailRequired returns the GuestEmailRequired field if non-nil, zero value otherwise.

### GetGuestEmailRequiredOk

`func (o *Captiveportal) GetGuestEmailRequiredOk() (*bool, bool)`

GetGuestEmailRequiredOk returns a tuple with the GuestEmailRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestEmailRequired

`func (o *Captiveportal) SetGuestEmailRequired(v bool)`

SetGuestEmailRequired sets GuestEmailRequired field to given value.

### HasGuestEmailRequired

`func (o *Captiveportal) HasGuestEmailRequired() bool`

HasGuestEmailRequired returns a boolean if a field has been set.

### GetGuestFirstNameRequired

`func (o *Captiveportal) GetGuestFirstNameRequired() bool`

GetGuestFirstNameRequired returns the GuestFirstNameRequired field if non-nil, zero value otherwise.

### GetGuestFirstNameRequiredOk

`func (o *Captiveportal) GetGuestFirstNameRequiredOk() (*bool, bool)`

GetGuestFirstNameRequiredOk returns a tuple with the GuestFirstNameRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestFirstNameRequired

`func (o *Captiveportal) SetGuestFirstNameRequired(v bool)`

SetGuestFirstNameRequired sets GuestFirstNameRequired field to given value.

### HasGuestFirstNameRequired

`func (o *Captiveportal) HasGuestFirstNameRequired() bool`

HasGuestFirstNameRequired returns a boolean if a field has been set.

### GetGuestLastNameRequired

`func (o *Captiveportal) GetGuestLastNameRequired() bool`

GetGuestLastNameRequired returns the GuestLastNameRequired field if non-nil, zero value otherwise.

### GetGuestLastNameRequiredOk

`func (o *Captiveportal) GetGuestLastNameRequiredOk() (*bool, bool)`

GetGuestLastNameRequiredOk returns a tuple with the GuestLastNameRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestLastNameRequired

`func (o *Captiveportal) SetGuestLastNameRequired(v bool)`

SetGuestLastNameRequired sets GuestLastNameRequired field to given value.

### HasGuestLastNameRequired

`func (o *Captiveportal) HasGuestLastNameRequired() bool`

HasGuestLastNameRequired returns a boolean if a field has been set.

### GetGuestMiddleNameRequired

`func (o *Captiveportal) GetGuestMiddleNameRequired() bool`

GetGuestMiddleNameRequired returns the GuestMiddleNameRequired field if non-nil, zero value otherwise.

### GetGuestMiddleNameRequiredOk

`func (o *Captiveportal) GetGuestMiddleNameRequiredOk() (*bool, bool)`

GetGuestMiddleNameRequiredOk returns a tuple with the GuestMiddleNameRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestMiddleNameRequired

`func (o *Captiveportal) SetGuestMiddleNameRequired(v bool)`

SetGuestMiddleNameRequired sets GuestMiddleNameRequired field to given value.

### HasGuestMiddleNameRequired

`func (o *Captiveportal) HasGuestMiddleNameRequired() bool`

HasGuestMiddleNameRequired returns a boolean if a field has been set.

### GetGuestPhoneRequired

`func (o *Captiveportal) GetGuestPhoneRequired() bool`

GetGuestPhoneRequired returns the GuestPhoneRequired field if non-nil, zero value otherwise.

### GetGuestPhoneRequiredOk

`func (o *Captiveportal) GetGuestPhoneRequiredOk() (*bool, bool)`

GetGuestPhoneRequiredOk returns a tuple with the GuestPhoneRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestPhoneRequired

`func (o *Captiveportal) SetGuestPhoneRequired(v bool)`

SetGuestPhoneRequired sets GuestPhoneRequired field to given value.

### HasGuestPhoneRequired

`func (o *Captiveportal) HasGuestPhoneRequired() bool`

HasGuestPhoneRequired returns a boolean if a field has been set.

### GetHelpdeskMessage

`func (o *Captiveportal) GetHelpdeskMessage() string`

GetHelpdeskMessage returns the HelpdeskMessage field if non-nil, zero value otherwise.

### GetHelpdeskMessageOk

`func (o *Captiveportal) GetHelpdeskMessageOk() (*string, bool)`

GetHelpdeskMessageOk returns a tuple with the HelpdeskMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpdeskMessage

`func (o *Captiveportal) SetHelpdeskMessage(v string)`

SetHelpdeskMessage sets HelpdeskMessage field to given value.

### HasHelpdeskMessage

`func (o *Captiveportal) HasHelpdeskMessage() bool`

HasHelpdeskMessage returns a boolean if a field has been set.

### GetListenAddressIp

`func (o *Captiveportal) GetListenAddressIp() string`

GetListenAddressIp returns the ListenAddressIp field if non-nil, zero value otherwise.

### GetListenAddressIpOk

`func (o *Captiveportal) GetListenAddressIpOk() (*string, bool)`

GetListenAddressIpOk returns a tuple with the ListenAddressIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenAddressIp

`func (o *Captiveportal) SetListenAddressIp(v string)`

SetListenAddressIp sets ListenAddressIp field to given value.

### HasListenAddressIp

`func (o *Captiveportal) HasListenAddressIp() bool`

HasListenAddressIp returns a boolean if a field has been set.

### GetListenAddressType

`func (o *Captiveportal) GetListenAddressType() string`

GetListenAddressType returns the ListenAddressType field if non-nil, zero value otherwise.

### GetListenAddressTypeOk

`func (o *Captiveportal) GetListenAddressTypeOk() (*string, bool)`

GetListenAddressTypeOk returns a tuple with the ListenAddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenAddressType

`func (o *Captiveportal) SetListenAddressType(v string)`

SetListenAddressType sets ListenAddressType field to given value.

### HasListenAddressType

`func (o *Captiveportal) HasListenAddressType() bool`

HasListenAddressType returns a boolean if a field has been set.

### GetName

`func (o *Captiveportal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Captiveportal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Captiveportal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Captiveportal) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkView

`func (o *Captiveportal) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *Captiveportal) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *Captiveportal) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *Captiveportal) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetPort

`func (o *Captiveportal) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Captiveportal) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Captiveportal) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *Captiveportal) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetServiceEnabled

`func (o *Captiveportal) GetServiceEnabled() bool`

GetServiceEnabled returns the ServiceEnabled field if non-nil, zero value otherwise.

### GetServiceEnabledOk

`func (o *Captiveportal) GetServiceEnabledOk() (*bool, bool)`

GetServiceEnabledOk returns a tuple with the ServiceEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceEnabled

`func (o *Captiveportal) SetServiceEnabled(v bool)`

SetServiceEnabled sets ServiceEnabled field to given value.

### HasServiceEnabled

`func (o *Captiveportal) HasServiceEnabled() bool`

HasServiceEnabled returns a boolean if a field has been set.

### GetSyslogAuthFailureLevel

`func (o *Captiveportal) GetSyslogAuthFailureLevel() string`

GetSyslogAuthFailureLevel returns the SyslogAuthFailureLevel field if non-nil, zero value otherwise.

### GetSyslogAuthFailureLevelOk

`func (o *Captiveportal) GetSyslogAuthFailureLevelOk() (*string, bool)`

GetSyslogAuthFailureLevelOk returns a tuple with the SyslogAuthFailureLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogAuthFailureLevel

`func (o *Captiveportal) SetSyslogAuthFailureLevel(v string)`

SetSyslogAuthFailureLevel sets SyslogAuthFailureLevel field to given value.

### HasSyslogAuthFailureLevel

`func (o *Captiveportal) HasSyslogAuthFailureLevel() bool`

HasSyslogAuthFailureLevel returns a boolean if a field has been set.

### GetSyslogAuthSuccessLevel

`func (o *Captiveportal) GetSyslogAuthSuccessLevel() string`

GetSyslogAuthSuccessLevel returns the SyslogAuthSuccessLevel field if non-nil, zero value otherwise.

### GetSyslogAuthSuccessLevelOk

`func (o *Captiveportal) GetSyslogAuthSuccessLevelOk() (*string, bool)`

GetSyslogAuthSuccessLevelOk returns a tuple with the SyslogAuthSuccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogAuthSuccessLevel

`func (o *Captiveportal) SetSyslogAuthSuccessLevel(v string)`

SetSyslogAuthSuccessLevel sets SyslogAuthSuccessLevel field to given value.

### HasSyslogAuthSuccessLevel

`func (o *Captiveportal) HasSyslogAuthSuccessLevel() bool`

HasSyslogAuthSuccessLevel returns a boolean if a field has been set.

### GetUuid

`func (o *Captiveportal) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Captiveportal) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Captiveportal) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Captiveportal) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWelcomeMessage

`func (o *Captiveportal) GetWelcomeMessage() string`

GetWelcomeMessage returns the WelcomeMessage field if non-nil, zero value otherwise.

### GetWelcomeMessageOk

`func (o *Captiveportal) GetWelcomeMessageOk() (*string, bool)`

GetWelcomeMessageOk returns a tuple with the WelcomeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeMessage

`func (o *Captiveportal) SetWelcomeMessage(v string)`

SetWelcomeMessage sets WelcomeMessage field to given value.

### HasWelcomeMessage

`func (o *Captiveportal) HasWelcomeMessage() bool`

HasWelcomeMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


