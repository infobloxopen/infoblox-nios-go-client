# GetMemberFiledistributionResponse

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
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**MemberFiledistribution**](MemberFiledistribution.md) |  | [optional] 

## Methods

### NewGetMemberFiledistributionResponse

`func NewGetMemberFiledistributionResponse() *GetMemberFiledistributionResponse`

NewGetMemberFiledistributionResponse instantiates a new GetMemberFiledistributionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMemberFiledistributionResponseWithDefaults

`func NewGetMemberFiledistributionResponseWithDefaults() *GetMemberFiledistributionResponse`

NewGetMemberFiledistributionResponseWithDefaults instantiates a new GetMemberFiledistributionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetMemberFiledistributionResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetMemberFiledistributionResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetMemberFiledistributionResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetMemberFiledistributionResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllowUploads

`func (o *GetMemberFiledistributionResponse) GetAllowUploads() bool`

GetAllowUploads returns the AllowUploads field if non-nil, zero value otherwise.

### GetAllowUploadsOk

`func (o *GetMemberFiledistributionResponse) GetAllowUploadsOk() (*bool, bool)`

GetAllowUploadsOk returns a tuple with the AllowUploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUploads

`func (o *GetMemberFiledistributionResponse) SetAllowUploads(v bool)`

SetAllowUploads sets AllowUploads field to given value.

### HasAllowUploads

`func (o *GetMemberFiledistributionResponse) HasAllowUploads() bool`

HasAllowUploads returns a boolean if a field has been set.

### GetComment

`func (o *GetMemberFiledistributionResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetMemberFiledistributionResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetMemberFiledistributionResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetMemberFiledistributionResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEnableFtp

`func (o *GetMemberFiledistributionResponse) GetEnableFtp() bool`

GetEnableFtp returns the EnableFtp field if non-nil, zero value otherwise.

### GetEnableFtpOk

`func (o *GetMemberFiledistributionResponse) GetEnableFtpOk() (*bool, bool)`

GetEnableFtpOk returns a tuple with the EnableFtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFtp

`func (o *GetMemberFiledistributionResponse) SetEnableFtp(v bool)`

SetEnableFtp sets EnableFtp field to given value.

### HasEnableFtp

`func (o *GetMemberFiledistributionResponse) HasEnableFtp() bool`

HasEnableFtp returns a boolean if a field has been set.

### GetEnableFtpFilelist

`func (o *GetMemberFiledistributionResponse) GetEnableFtpFilelist() bool`

GetEnableFtpFilelist returns the EnableFtpFilelist field if non-nil, zero value otherwise.

### GetEnableFtpFilelistOk

`func (o *GetMemberFiledistributionResponse) GetEnableFtpFilelistOk() (*bool, bool)`

GetEnableFtpFilelistOk returns a tuple with the EnableFtpFilelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFtpFilelist

`func (o *GetMemberFiledistributionResponse) SetEnableFtpFilelist(v bool)`

SetEnableFtpFilelist sets EnableFtpFilelist field to given value.

### HasEnableFtpFilelist

`func (o *GetMemberFiledistributionResponse) HasEnableFtpFilelist() bool`

HasEnableFtpFilelist returns a boolean if a field has been set.

### GetEnableFtpPassive

`func (o *GetMemberFiledistributionResponse) GetEnableFtpPassive() bool`

GetEnableFtpPassive returns the EnableFtpPassive field if non-nil, zero value otherwise.

### GetEnableFtpPassiveOk

`func (o *GetMemberFiledistributionResponse) GetEnableFtpPassiveOk() (*bool, bool)`

GetEnableFtpPassiveOk returns a tuple with the EnableFtpPassive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFtpPassive

`func (o *GetMemberFiledistributionResponse) SetEnableFtpPassive(v bool)`

SetEnableFtpPassive sets EnableFtpPassive field to given value.

### HasEnableFtpPassive

`func (o *GetMemberFiledistributionResponse) HasEnableFtpPassive() bool`

HasEnableFtpPassive returns a boolean if a field has been set.

### GetEnableHttp

`func (o *GetMemberFiledistributionResponse) GetEnableHttp() bool`

GetEnableHttp returns the EnableHttp field if non-nil, zero value otherwise.

### GetEnableHttpOk

`func (o *GetMemberFiledistributionResponse) GetEnableHttpOk() (*bool, bool)`

GetEnableHttpOk returns a tuple with the EnableHttp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHttp

`func (o *GetMemberFiledistributionResponse) SetEnableHttp(v bool)`

SetEnableHttp sets EnableHttp field to given value.

### HasEnableHttp

`func (o *GetMemberFiledistributionResponse) HasEnableHttp() bool`

HasEnableHttp returns a boolean if a field has been set.

### GetEnableHttpAcl

`func (o *GetMemberFiledistributionResponse) GetEnableHttpAcl() bool`

GetEnableHttpAcl returns the EnableHttpAcl field if non-nil, zero value otherwise.

### GetEnableHttpAclOk

`func (o *GetMemberFiledistributionResponse) GetEnableHttpAclOk() (*bool, bool)`

GetEnableHttpAclOk returns a tuple with the EnableHttpAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHttpAcl

`func (o *GetMemberFiledistributionResponse) SetEnableHttpAcl(v bool)`

SetEnableHttpAcl sets EnableHttpAcl field to given value.

### HasEnableHttpAcl

`func (o *GetMemberFiledistributionResponse) HasEnableHttpAcl() bool`

HasEnableHttpAcl returns a boolean if a field has been set.

### GetEnableTftp

`func (o *GetMemberFiledistributionResponse) GetEnableTftp() bool`

GetEnableTftp returns the EnableTftp field if non-nil, zero value otherwise.

### GetEnableTftpOk

`func (o *GetMemberFiledistributionResponse) GetEnableTftpOk() (*bool, bool)`

GetEnableTftpOk returns a tuple with the EnableTftp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTftp

`func (o *GetMemberFiledistributionResponse) SetEnableTftp(v bool)`

SetEnableTftp sets EnableTftp field to given value.

### HasEnableTftp

`func (o *GetMemberFiledistributionResponse) HasEnableTftp() bool`

HasEnableTftp returns a boolean if a field has been set.

### GetFtpAcls

`func (o *GetMemberFiledistributionResponse) GetFtpAcls() []MemberFiledistributionFtpAcls`

GetFtpAcls returns the FtpAcls field if non-nil, zero value otherwise.

### GetFtpAclsOk

`func (o *GetMemberFiledistributionResponse) GetFtpAclsOk() (*[]MemberFiledistributionFtpAcls, bool)`

GetFtpAclsOk returns a tuple with the FtpAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpAcls

`func (o *GetMemberFiledistributionResponse) SetFtpAcls(v []MemberFiledistributionFtpAcls)`

SetFtpAcls sets FtpAcls field to given value.

### HasFtpAcls

`func (o *GetMemberFiledistributionResponse) HasFtpAcls() bool`

HasFtpAcls returns a boolean if a field has been set.

### GetFtpPort

`func (o *GetMemberFiledistributionResponse) GetFtpPort() int64`

GetFtpPort returns the FtpPort field if non-nil, zero value otherwise.

### GetFtpPortOk

`func (o *GetMemberFiledistributionResponse) GetFtpPortOk() (*int64, bool)`

GetFtpPortOk returns a tuple with the FtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpPort

`func (o *GetMemberFiledistributionResponse) SetFtpPort(v int64)`

SetFtpPort sets FtpPort field to given value.

### HasFtpPort

`func (o *GetMemberFiledistributionResponse) HasFtpPort() bool`

HasFtpPort returns a boolean if a field has been set.

### GetFtpStatus

`func (o *GetMemberFiledistributionResponse) GetFtpStatus() string`

GetFtpStatus returns the FtpStatus field if non-nil, zero value otherwise.

### GetFtpStatusOk

`func (o *GetMemberFiledistributionResponse) GetFtpStatusOk() (*string, bool)`

GetFtpStatusOk returns a tuple with the FtpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpStatus

`func (o *GetMemberFiledistributionResponse) SetFtpStatus(v string)`

SetFtpStatus sets FtpStatus field to given value.

### HasFtpStatus

`func (o *GetMemberFiledistributionResponse) HasFtpStatus() bool`

HasFtpStatus returns a boolean if a field has been set.

### GetHostName

`func (o *GetMemberFiledistributionResponse) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *GetMemberFiledistributionResponse) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *GetMemberFiledistributionResponse) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *GetMemberFiledistributionResponse) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetHttpAcls

`func (o *GetMemberFiledistributionResponse) GetHttpAcls() []MemberFiledistributionHttpAcls`

GetHttpAcls returns the HttpAcls field if non-nil, zero value otherwise.

### GetHttpAclsOk

`func (o *GetMemberFiledistributionResponse) GetHttpAclsOk() (*[]MemberFiledistributionHttpAcls, bool)`

GetHttpAclsOk returns a tuple with the HttpAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpAcls

`func (o *GetMemberFiledistributionResponse) SetHttpAcls(v []MemberFiledistributionHttpAcls)`

SetHttpAcls sets HttpAcls field to given value.

### HasHttpAcls

`func (o *GetMemberFiledistributionResponse) HasHttpAcls() bool`

HasHttpAcls returns a boolean if a field has been set.

### GetHttpStatus

`func (o *GetMemberFiledistributionResponse) GetHttpStatus() string`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *GetMemberFiledistributionResponse) GetHttpStatusOk() (*string, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *GetMemberFiledistributionResponse) SetHttpStatus(v string)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *GetMemberFiledistributionResponse) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.

### GetIpv4Address

`func (o *GetMemberFiledistributionResponse) GetIpv4Address() string`

GetIpv4Address returns the Ipv4Address field if non-nil, zero value otherwise.

### GetIpv4AddressOk

`func (o *GetMemberFiledistributionResponse) GetIpv4AddressOk() (*string, bool)`

GetIpv4AddressOk returns a tuple with the Ipv4Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Address

`func (o *GetMemberFiledistributionResponse) SetIpv4Address(v string)`

SetIpv4Address sets Ipv4Address field to given value.

### HasIpv4Address

`func (o *GetMemberFiledistributionResponse) HasIpv4Address() bool`

HasIpv4Address returns a boolean if a field has been set.

### GetIpv6Address

`func (o *GetMemberFiledistributionResponse) GetIpv6Address() string`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *GetMemberFiledistributionResponse) GetIpv6AddressOk() (*string, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *GetMemberFiledistributionResponse) SetIpv6Address(v string)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *GetMemberFiledistributionResponse) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### GetStatus

`func (o *GetMemberFiledistributionResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMemberFiledistributionResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMemberFiledistributionResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetMemberFiledistributionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTftpAcls

`func (o *GetMemberFiledistributionResponse) GetTftpAcls() []MemberFiledistributionTftpAcls`

GetTftpAcls returns the TftpAcls field if non-nil, zero value otherwise.

### GetTftpAclsOk

`func (o *GetMemberFiledistributionResponse) GetTftpAclsOk() (*[]MemberFiledistributionTftpAcls, bool)`

GetTftpAclsOk returns a tuple with the TftpAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTftpAcls

`func (o *GetMemberFiledistributionResponse) SetTftpAcls(v []MemberFiledistributionTftpAcls)`

SetTftpAcls sets TftpAcls field to given value.

### HasTftpAcls

`func (o *GetMemberFiledistributionResponse) HasTftpAcls() bool`

HasTftpAcls returns a boolean if a field has been set.

### GetTftpPort

`func (o *GetMemberFiledistributionResponse) GetTftpPort() int64`

GetTftpPort returns the TftpPort field if non-nil, zero value otherwise.

### GetTftpPortOk

`func (o *GetMemberFiledistributionResponse) GetTftpPortOk() (*int64, bool)`

GetTftpPortOk returns a tuple with the TftpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTftpPort

`func (o *GetMemberFiledistributionResponse) SetTftpPort(v int64)`

SetTftpPort sets TftpPort field to given value.

### HasTftpPort

`func (o *GetMemberFiledistributionResponse) HasTftpPort() bool`

HasTftpPort returns a boolean if a field has been set.

### GetTftpStatus

`func (o *GetMemberFiledistributionResponse) GetTftpStatus() string`

GetTftpStatus returns the TftpStatus field if non-nil, zero value otherwise.

### GetTftpStatusOk

`func (o *GetMemberFiledistributionResponse) GetTftpStatusOk() (*string, bool)`

GetTftpStatusOk returns a tuple with the TftpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTftpStatus

`func (o *GetMemberFiledistributionResponse) SetTftpStatus(v string)`

SetTftpStatus sets TftpStatus field to given value.

### HasTftpStatus

`func (o *GetMemberFiledistributionResponse) HasTftpStatus() bool`

HasTftpStatus returns a boolean if a field has been set.

### GetUseAllowUploads

`func (o *GetMemberFiledistributionResponse) GetUseAllowUploads() bool`

GetUseAllowUploads returns the UseAllowUploads field if non-nil, zero value otherwise.

### GetUseAllowUploadsOk

`func (o *GetMemberFiledistributionResponse) GetUseAllowUploadsOk() (*bool, bool)`

GetUseAllowUploadsOk returns a tuple with the UseAllowUploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowUploads

`func (o *GetMemberFiledistributionResponse) SetUseAllowUploads(v bool)`

SetUseAllowUploads sets UseAllowUploads field to given value.

### HasUseAllowUploads

`func (o *GetMemberFiledistributionResponse) HasUseAllowUploads() bool`

HasUseAllowUploads returns a boolean if a field has been set.

### GetUuid

`func (o *GetMemberFiledistributionResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetMemberFiledistributionResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetMemberFiledistributionResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetMemberFiledistributionResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetMemberFiledistributionResponse) GetResult() MemberFiledistribution`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMemberFiledistributionResponse) GetResultOk() (*MemberFiledistribution, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMemberFiledistributionResponse) SetResult(v MemberFiledistribution)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetMemberFiledistributionResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


