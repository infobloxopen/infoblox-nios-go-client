# MemberFiledistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AllowUploads** | Pointer to **bool** | Determines whether uploads to the Grid member are allowed. | [optional] 
**Comment** | Pointer to **string** | The Grid member descriptive comment. | [optional] [readonly] 
**EnableFtp** | Pointer to **bool** | Determines whether the FTP prtocol is enabled for file distribution. | [optional] 
**EnableFtpFilelist** | Pointer to **bool** | Determines whether the LIST command for FTP is enabled. | [optional] 
**EnableFtpPassive** | Pointer to **bool** | Determines whether the passive mode for FTP is enabled. | [optional] 
**EnableHttp** | Pointer to **bool** | Determines whether the HTTP prtocol is enabled for file distribution. | [optional] 
**EnableHttpAcl** | Pointer to **bool** | Determines whether the HTTP prtocol access control (AC) settings are enabled. | [optional] 
**EnableTftp** | Pointer to **bool** | Determines whether the TFTP prtocol is enabled for file distribution. | [optional] 
**FtpAcls** | Pointer to [**[]MemberFiledistributionFtpAcls**](MemberFiledistributionFtpAcls.md) | Access control (AC) settings for the FTP protocol. | [optional] 
**FtpPort** | Pointer to **int64** | The network port used by the FTP protocol. | [optional] 
**FtpStatus** | Pointer to **string** | The FTP protocol status. | [optional] [readonly] 
**HostName** | Pointer to **string** | The Grid member host name. | [optional] [readonly] 
**HttpAcls** | Pointer to [**[]MemberFiledistributionHttpAcls**](MemberFiledistributionHttpAcls.md) | Access control (AC) settings for the HTTP protocol. | [optional] 
**HttpStatus** | Pointer to **string** | The HTTP protocol status. | [optional] [readonly] 
**Ipv4Address** | Pointer to **string** | The IPv4 address of the Grid member. | [optional] [readonly] 
**Ipv6Address** | Pointer to **string** | The IPv6 address of the Grid member. | [optional] [readonly] 
**Status** | Pointer to **string** | The Grid member file distribution status. | [optional] [readonly] 
**TftpAcls** | Pointer to [**[]MemberFiledistributionTftpAcls**](MemberFiledistributionTftpAcls.md) | The access control (AC) settings for the TFTP protocol. | [optional] 
**TftpPort** | Pointer to **int64** | The network port used by the TFTP protocol. | [optional] 
**TftpStatus** | Pointer to **string** | The TFTP protocol status. | [optional] [readonly] 
**UseAllowUploads** | Pointer to **bool** | Use flag for: allow_uploads | [optional] 

## Methods

### NewMemberFiledistribution

`func NewMemberFiledistribution() *MemberFiledistribution`

NewMemberFiledistribution instantiates a new MemberFiledistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberFiledistributionWithDefaults

`func NewMemberFiledistributionWithDefaults() *MemberFiledistribution`

NewMemberFiledistributionWithDefaults instantiates a new MemberFiledistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *MemberFiledistribution) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *MemberFiledistribution) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *MemberFiledistribution) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *MemberFiledistribution) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllowUploads

`func (o *MemberFiledistribution) GetAllowUploads() bool`

GetAllowUploads returns the AllowUploads field if non-nil, zero value otherwise.

### GetAllowUploadsOk

`func (o *MemberFiledistribution) GetAllowUploadsOk() (*bool, bool)`

GetAllowUploadsOk returns a tuple with the AllowUploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUploads

`func (o *MemberFiledistribution) SetAllowUploads(v bool)`

SetAllowUploads sets AllowUploads field to given value.

### HasAllowUploads

`func (o *MemberFiledistribution) HasAllowUploads() bool`

HasAllowUploads returns a boolean if a field has been set.

### GetComment

`func (o *MemberFiledistribution) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MemberFiledistribution) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MemberFiledistribution) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MemberFiledistribution) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEnableFtp

`func (o *MemberFiledistribution) GetEnableFtp() bool`

GetEnableFtp returns the EnableFtp field if non-nil, zero value otherwise.

### GetEnableFtpOk

`func (o *MemberFiledistribution) GetEnableFtpOk() (*bool, bool)`

GetEnableFtpOk returns a tuple with the EnableFtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFtp

`func (o *MemberFiledistribution) SetEnableFtp(v bool)`

SetEnableFtp sets EnableFtp field to given value.

### HasEnableFtp

`func (o *MemberFiledistribution) HasEnableFtp() bool`

HasEnableFtp returns a boolean if a field has been set.

### GetEnableFtpFilelist

`func (o *MemberFiledistribution) GetEnableFtpFilelist() bool`

GetEnableFtpFilelist returns the EnableFtpFilelist field if non-nil, zero value otherwise.

### GetEnableFtpFilelistOk

`func (o *MemberFiledistribution) GetEnableFtpFilelistOk() (*bool, bool)`

GetEnableFtpFilelistOk returns a tuple with the EnableFtpFilelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFtpFilelist

`func (o *MemberFiledistribution) SetEnableFtpFilelist(v bool)`

SetEnableFtpFilelist sets EnableFtpFilelist field to given value.

### HasEnableFtpFilelist

`func (o *MemberFiledistribution) HasEnableFtpFilelist() bool`

HasEnableFtpFilelist returns a boolean if a field has been set.

### GetEnableFtpPassive

`func (o *MemberFiledistribution) GetEnableFtpPassive() bool`

GetEnableFtpPassive returns the EnableFtpPassive field if non-nil, zero value otherwise.

### GetEnableFtpPassiveOk

`func (o *MemberFiledistribution) GetEnableFtpPassiveOk() (*bool, bool)`

GetEnableFtpPassiveOk returns a tuple with the EnableFtpPassive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFtpPassive

`func (o *MemberFiledistribution) SetEnableFtpPassive(v bool)`

SetEnableFtpPassive sets EnableFtpPassive field to given value.

### HasEnableFtpPassive

`func (o *MemberFiledistribution) HasEnableFtpPassive() bool`

HasEnableFtpPassive returns a boolean if a field has been set.

### GetEnableHttp

`func (o *MemberFiledistribution) GetEnableHttp() bool`

GetEnableHttp returns the EnableHttp field if non-nil, zero value otherwise.

### GetEnableHttpOk

`func (o *MemberFiledistribution) GetEnableHttpOk() (*bool, bool)`

GetEnableHttpOk returns a tuple with the EnableHttp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHttp

`func (o *MemberFiledistribution) SetEnableHttp(v bool)`

SetEnableHttp sets EnableHttp field to given value.

### HasEnableHttp

`func (o *MemberFiledistribution) HasEnableHttp() bool`

HasEnableHttp returns a boolean if a field has been set.

### GetEnableHttpAcl

`func (o *MemberFiledistribution) GetEnableHttpAcl() bool`

GetEnableHttpAcl returns the EnableHttpAcl field if non-nil, zero value otherwise.

### GetEnableHttpAclOk

`func (o *MemberFiledistribution) GetEnableHttpAclOk() (*bool, bool)`

GetEnableHttpAclOk returns a tuple with the EnableHttpAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHttpAcl

`func (o *MemberFiledistribution) SetEnableHttpAcl(v bool)`

SetEnableHttpAcl sets EnableHttpAcl field to given value.

### HasEnableHttpAcl

`func (o *MemberFiledistribution) HasEnableHttpAcl() bool`

HasEnableHttpAcl returns a boolean if a field has been set.

### GetEnableTftp

`func (o *MemberFiledistribution) GetEnableTftp() bool`

GetEnableTftp returns the EnableTftp field if non-nil, zero value otherwise.

### GetEnableTftpOk

`func (o *MemberFiledistribution) GetEnableTftpOk() (*bool, bool)`

GetEnableTftpOk returns a tuple with the EnableTftp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTftp

`func (o *MemberFiledistribution) SetEnableTftp(v bool)`

SetEnableTftp sets EnableTftp field to given value.

### HasEnableTftp

`func (o *MemberFiledistribution) HasEnableTftp() bool`

HasEnableTftp returns a boolean if a field has been set.

### GetFtpAcls

`func (o *MemberFiledistribution) GetFtpAcls() []MemberFiledistributionFtpAcls`

GetFtpAcls returns the FtpAcls field if non-nil, zero value otherwise.

### GetFtpAclsOk

`func (o *MemberFiledistribution) GetFtpAclsOk() (*[]MemberFiledistributionFtpAcls, bool)`

GetFtpAclsOk returns a tuple with the FtpAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpAcls

`func (o *MemberFiledistribution) SetFtpAcls(v []MemberFiledistributionFtpAcls)`

SetFtpAcls sets FtpAcls field to given value.

### HasFtpAcls

`func (o *MemberFiledistribution) HasFtpAcls() bool`

HasFtpAcls returns a boolean if a field has been set.

### GetFtpPort

`func (o *MemberFiledistribution) GetFtpPort() int64`

GetFtpPort returns the FtpPort field if non-nil, zero value otherwise.

### GetFtpPortOk

`func (o *MemberFiledistribution) GetFtpPortOk() (*int64, bool)`

GetFtpPortOk returns a tuple with the FtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpPort

`func (o *MemberFiledistribution) SetFtpPort(v int64)`

SetFtpPort sets FtpPort field to given value.

### HasFtpPort

`func (o *MemberFiledistribution) HasFtpPort() bool`

HasFtpPort returns a boolean if a field has been set.

### GetFtpStatus

`func (o *MemberFiledistribution) GetFtpStatus() string`

GetFtpStatus returns the FtpStatus field if non-nil, zero value otherwise.

### GetFtpStatusOk

`func (o *MemberFiledistribution) GetFtpStatusOk() (*string, bool)`

GetFtpStatusOk returns a tuple with the FtpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpStatus

`func (o *MemberFiledistribution) SetFtpStatus(v string)`

SetFtpStatus sets FtpStatus field to given value.

### HasFtpStatus

`func (o *MemberFiledistribution) HasFtpStatus() bool`

HasFtpStatus returns a boolean if a field has been set.

### GetHostName

`func (o *MemberFiledistribution) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *MemberFiledistribution) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *MemberFiledistribution) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *MemberFiledistribution) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetHttpAcls

`func (o *MemberFiledistribution) GetHttpAcls() []MemberFiledistributionHttpAcls`

GetHttpAcls returns the HttpAcls field if non-nil, zero value otherwise.

### GetHttpAclsOk

`func (o *MemberFiledistribution) GetHttpAclsOk() (*[]MemberFiledistributionHttpAcls, bool)`

GetHttpAclsOk returns a tuple with the HttpAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpAcls

`func (o *MemberFiledistribution) SetHttpAcls(v []MemberFiledistributionHttpAcls)`

SetHttpAcls sets HttpAcls field to given value.

### HasHttpAcls

`func (o *MemberFiledistribution) HasHttpAcls() bool`

HasHttpAcls returns a boolean if a field has been set.

### GetHttpStatus

`func (o *MemberFiledistribution) GetHttpStatus() string`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *MemberFiledistribution) GetHttpStatusOk() (*string, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *MemberFiledistribution) SetHttpStatus(v string)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *MemberFiledistribution) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.

### GetIpv4Address

`func (o *MemberFiledistribution) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *MemberFiledistribution) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *MemberFiledistribution) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *MemberFiledistribution) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetIpv6Address

`func (o *MemberFiledistribution) GetIpv6Address() string`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *MemberFiledistribution) GetIpv6AddressOk() (*string, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *MemberFiledistribution) SetIpv6Address(v string)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *MemberFiledistribution) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### GetStatus

`func (o *MemberFiledistribution) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MemberFiledistribution) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MemberFiledistribution) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MemberFiledistribution) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTftpAcls

`func (o *MemberFiledistribution) GetTftpAcls() []MemberFiledistributionTftpAcls`

GetTftpAcls returns the TftpAcls field if non-nil, zero value otherwise.

### GetTftpAclsOk

`func (o *MemberFiledistribution) GetTftpAclsOk() (*[]MemberFiledistributionTftpAcls, bool)`

GetTftpAclsOk returns a tuple with the TftpAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTftpAcls

`func (o *MemberFiledistribution) SetTftpAcls(v []MemberFiledistributionTftpAcls)`

SetTftpAcls sets TftpAcls field to given value.

### HasTftpAcls

`func (o *MemberFiledistribution) HasTftpAcls() bool`

HasTftpAcls returns a boolean if a field has been set.

### GetTftpPort

`func (o *MemberFiledistribution) GetTftpPort() int64`

GetTftpPort returns the TftpPort field if non-nil, zero value otherwise.

### GetTftpPortOk

`func (o *MemberFiledistribution) GetTftpPortOk() (*int64, bool)`

GetTftpPortOk returns a tuple with the TftpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTftpPort

`func (o *MemberFiledistribution) SetTftpPort(v int64)`

SetTftpPort sets TftpPort field to given value.

### HasTftpPort

`func (o *MemberFiledistribution) HasTftpPort() bool`

HasTftpPort returns a boolean if a field has been set.

### GetTftpStatus

`func (o *MemberFiledistribution) GetTftpStatus() string`

GetTftpStatus returns the TftpStatus field if non-nil, zero value otherwise.

### GetTftpStatusOk

`func (o *MemberFiledistribution) GetTftpStatusOk() (*string, bool)`

GetTftpStatusOk returns a tuple with the TftpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTftpStatus

`func (o *MemberFiledistribution) SetTftpStatus(v string)`

SetTftpStatus sets TftpStatus field to given value.

### HasTftpStatus

`func (o *MemberFiledistribution) HasTftpStatus() bool`

HasTftpStatus returns a boolean if a field has been set.

### GetUseAllowUploads

`func (o *MemberFiledistribution) GetUseAllowUploads() bool`

GetUseAllowUploads returns the UseAllowUploads field if non-nil, zero value otherwise.

### GetUseAllowUploadsOk

`func (o *MemberFiledistribution) GetUseAllowUploadsOk() (*bool, bool)`

GetUseAllowUploadsOk returns a tuple with the UseAllowUploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowUploads

`func (o *MemberFiledistribution) SetUseAllowUploads(v bool)`

SetUseAllowUploads sets UseAllowUploads field to given value.

### HasUseAllowUploads

`func (o *MemberFiledistribution) HasUseAllowUploads() bool`

HasUseAllowUploads returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


