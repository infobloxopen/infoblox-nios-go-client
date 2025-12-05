# Vdiscoverytask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AccountsList** | Pointer to **[]string** | The AWS Account IDs or GCP Project IDs list associated with this task. | [optional] [readonly] 
**AllowUnsecuredConnection** | Pointer to **bool** | Allow unsecured connection over HTTPS and bypass validation of the remote SSL certificate. | [optional] 
**AutoConsolidateCloudEa** | Pointer to **bool** | Whether to insert or update cloud EAs with discovery data. | [optional] 
**AutoConsolidateManagedTenant** | Pointer to **bool** | Whether to replace managed tenant with discovery tenant data. | [optional] 
**AutoConsolidateManagedVm** | Pointer to **bool** | Whether to replace managed virtual machine with discovery vm data. | [optional] 
**AutoCreateDnsHostnameTemplate** | Pointer to **string** | Template string used to generate host name. | [optional] 
**AutoCreateDnsRecord** | Pointer to **bool** | Control whether to create or update DNS record using discovered data. | [optional] 
**AutoCreateDnsRecordType** | Pointer to **string** | Indicates the type of record to create if the auto create DNS record is enabled. | [optional] 
**CdiscoveryFileToken** | Pointer to **string** | The AWS account IDs or GCP Project IDs file&#39;s token. | [optional] 
**Comment** | Pointer to **string** | Comment on the task. | [optional] 
**CredentialsType** | Pointer to **string** | Credentials type used for connecting to the cloud management platform. | [optional] 
**DnsViewPrivateIp** | Pointer to **string** | The DNS view name for private IPs. | [optional] 
**DnsViewPublicIp** | Pointer to **string** | The DNS view name for public IPs. | [optional] 
**DomainName** | Pointer to **string** | The name of the domain to use with keystone v3. | [optional] 
**DriverType** | Pointer to **string** | Type of discovery driver. | [optional] 
**EnableFilter** | Pointer to **bool** | Enable filter for cloud discovery task | [optional] 
**Enabled** | Pointer to **bool** | Whether to enabled the cloud discovery or not. | [optional] 
**FqdnOrIp** | Pointer to **string** | FQDN or IP of the cloud management platform. | [optional] 
**GovcloudEnabled** | Pointer to **bool** | Indicates if gov cloud is enabled or disabled. | [optional] 
**IdentityVersion** | Pointer to **string** | Identity service version. | [optional] 
**LastRun** | Pointer to **int64** | Timestamp of last run. | [optional] [readonly] 
**Member** | Pointer to **string** | Member on which cloud discovery will be run. | [optional] 
**MergeData** | Pointer to **bool** | Whether to replace the old data with new or not. | [optional] 
**MultipleAccountsSyncPolicy** | Pointer to **string** | Discover all child accounts or Upload child account ids to discover.. | [optional] 
**Name** | Pointer to **string** | Name of this cloud discovery task. Uniquely identify a task. | [optional] 
**NetworkFilter** | Pointer to **string** | Options to filter the networks in cdiscovery task. | [optional] 
**NetworkList** | Pointer to **[]string** | List of networks to filter in cdiscovery task. | [optional] 
**Password** | Pointer to **string** | Password used for connecting to the cloud management platform. | [optional] 
**Port** | Pointer to **int64** | Connection port used for connecting to the cloud management platform. | [optional] 
**PrivateNetworkView** | Pointer to **string** | Network view for private IPs. | [optional] 
**PrivateNetworkViewMappingPolicy** | Pointer to **string** | Mapping policy for the network view for private IPs in discovery data. | [optional] 
**Protocol** | Pointer to **string** | Connection protocol used for connecting to the cloud management platform. | [optional] 
**PublicNetworkView** | Pointer to **string** | Network view for public IPs. | [optional] 
**PublicNetworkViewMappingPolicy** | Pointer to **string** | Mapping policy for the network view for public IPs in discovery data. | [optional] 
**RoleArn** | Pointer to **string** | Role ARN for syncing child accounts; maximum 128 characters. | [optional] 
**ScheduledRun** | Pointer to [**VdiscoverytaskScheduledRun**](VdiscoverytaskScheduledRun.md) |  | [optional] 
**SelectedRegions** | Pointer to **string** | String containing selected regions for discovery in comma separated format. | [optional] 
**ServiceAccountFile** | Pointer to **string** | The service_account_file for GCP. | [optional] 
**ServiceAccountFileToken** | Pointer to **string** | Service account file&#39;s token. | [optional] 
**State** | Pointer to **string** | Current state of this task. | [optional] [readonly] 
**StateMsg** | Pointer to **string** | State message of the complete discovery process. | [optional] [readonly] 
**SyncChildAccounts** | Pointer to **bool** | Synchronizing child accounts is enabled or disabled. | [optional] 
**UpdateDnsViewPrivateIp** | Pointer to **bool** | If set to true, the appliance uses a specific DNS view for private IPs. | [optional] 
**UpdateDnsViewPublicIp** | Pointer to **bool** | If set to true, the appliance uses a specific DNS view for public IPs. | [optional] 
**UpdateMetadata** | Pointer to **bool** | Whether to update metadata as a result of this network discovery. | [optional] 
**UseIdentity** | Pointer to **bool** | If set true, all keystone connection will use \&quot;/identity\&quot; endpoint and port value will be ignored. | [optional] 
**Username** | Pointer to **string** | Username used for connecting to the cloud management platform. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewVdiscoverytask

`func NewVdiscoverytask() *Vdiscoverytask`

NewVdiscoverytask instantiates a new Vdiscoverytask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVdiscoverytaskWithDefaults

`func NewVdiscoverytaskWithDefaults() *Vdiscoverytask`

NewVdiscoverytaskWithDefaults instantiates a new Vdiscoverytask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Vdiscoverytask) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Vdiscoverytask) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Vdiscoverytask) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Vdiscoverytask) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAccountsList

`func (o *Vdiscoverytask) GetAccountsList() []string`

GetAccountsList returns the AccountsList field if non-nil, zero value otherwise.

### GetAccountsListOk

`func (o *Vdiscoverytask) GetAccountsListOk() (*[]string, bool)`

GetAccountsListOk returns a tuple with the AccountsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsList

`func (o *Vdiscoverytask) SetAccountsList(v []string)`

SetAccountsList sets AccountsList field to given value.

### HasAccountsList

`func (o *Vdiscoverytask) HasAccountsList() bool`

HasAccountsList returns a boolean if a field has been set.

### GetAllowUnsecuredConnection

`func (o *Vdiscoverytask) GetAllowUnsecuredConnection() bool`

GetAllowUnsecuredConnection returns the AllowUnsecuredConnection field if non-nil, zero value otherwise.

### GetAllowUnsecuredConnectionOk

`func (o *Vdiscoverytask) GetAllowUnsecuredConnectionOk() (*bool, bool)`

GetAllowUnsecuredConnectionOk returns a tuple with the AllowUnsecuredConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUnsecuredConnection

`func (o *Vdiscoverytask) SetAllowUnsecuredConnection(v bool)`

SetAllowUnsecuredConnection sets AllowUnsecuredConnection field to given value.

### HasAllowUnsecuredConnection

`func (o *Vdiscoverytask) HasAllowUnsecuredConnection() bool`

HasAllowUnsecuredConnection returns a boolean if a field has been set.

### GetAutoConsolidateCloudEa

`func (o *Vdiscoverytask) GetAutoConsolidateCloudEa() bool`

GetAutoConsolidateCloudEa returns the AutoConsolidateCloudEa field if non-nil, zero value otherwise.

### GetAutoConsolidateCloudEaOk

`func (o *Vdiscoverytask) GetAutoConsolidateCloudEaOk() (*bool, bool)`

GetAutoConsolidateCloudEaOk returns a tuple with the AutoConsolidateCloudEa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConsolidateCloudEa

`func (o *Vdiscoverytask) SetAutoConsolidateCloudEa(v bool)`

SetAutoConsolidateCloudEa sets AutoConsolidateCloudEa field to given value.

### HasAutoConsolidateCloudEa

`func (o *Vdiscoverytask) HasAutoConsolidateCloudEa() bool`

HasAutoConsolidateCloudEa returns a boolean if a field has been set.

### GetAutoConsolidateManagedTenant

`func (o *Vdiscoverytask) GetAutoConsolidateManagedTenant() bool`

GetAutoConsolidateManagedTenant returns the AutoConsolidateManagedTenant field if non-nil, zero value otherwise.

### GetAutoConsolidateManagedTenantOk

`func (o *Vdiscoverytask) GetAutoConsolidateManagedTenantOk() (*bool, bool)`

GetAutoConsolidateManagedTenantOk returns a tuple with the AutoConsolidateManagedTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConsolidateManagedTenant

`func (o *Vdiscoverytask) SetAutoConsolidateManagedTenant(v bool)`

SetAutoConsolidateManagedTenant sets AutoConsolidateManagedTenant field to given value.

### HasAutoConsolidateManagedTenant

`func (o *Vdiscoverytask) HasAutoConsolidateManagedTenant() bool`

HasAutoConsolidateManagedTenant returns a boolean if a field has been set.

### GetAutoConsolidateManagedVm

`func (o *Vdiscoverytask) GetAutoConsolidateManagedVm() bool`

GetAutoConsolidateManagedVm returns the AutoConsolidateManagedVm field if non-nil, zero value otherwise.

### GetAutoConsolidateManagedVmOk

`func (o *Vdiscoverytask) GetAutoConsolidateManagedVmOk() (*bool, bool)`

GetAutoConsolidateManagedVmOk returns a tuple with the AutoConsolidateManagedVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConsolidateManagedVm

`func (o *Vdiscoverytask) SetAutoConsolidateManagedVm(v bool)`

SetAutoConsolidateManagedVm sets AutoConsolidateManagedVm field to given value.

### HasAutoConsolidateManagedVm

`func (o *Vdiscoverytask) HasAutoConsolidateManagedVm() bool`

HasAutoConsolidateManagedVm returns a boolean if a field has been set.

### GetAutoCreateDnsHostnameTemplate

`func (o *Vdiscoverytask) GetAutoCreateDnsHostnameTemplate() string`

GetAutoCreateDnsHostnameTemplate returns the AutoCreateDnsHostnameTemplate field if non-nil, zero value otherwise.

### GetAutoCreateDnsHostnameTemplateOk

`func (o *Vdiscoverytask) GetAutoCreateDnsHostnameTemplateOk() (*string, bool)`

GetAutoCreateDnsHostnameTemplateOk returns a tuple with the AutoCreateDnsHostnameTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateDnsHostnameTemplate

`func (o *Vdiscoverytask) SetAutoCreateDnsHostnameTemplate(v string)`

SetAutoCreateDnsHostnameTemplate sets AutoCreateDnsHostnameTemplate field to given value.

### HasAutoCreateDnsHostnameTemplate

`func (o *Vdiscoverytask) HasAutoCreateDnsHostnameTemplate() bool`

HasAutoCreateDnsHostnameTemplate returns a boolean if a field has been set.

### GetAutoCreateDnsRecord

`func (o *Vdiscoverytask) GetAutoCreateDnsRecord() bool`

GetAutoCreateDnsRecord returns the AutoCreateDnsRecord field if non-nil, zero value otherwise.

### GetAutoCreateDnsRecordOk

`func (o *Vdiscoverytask) GetAutoCreateDnsRecordOk() (*bool, bool)`

GetAutoCreateDnsRecordOk returns a tuple with the AutoCreateDnsRecord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateDnsRecord

`func (o *Vdiscoverytask) SetAutoCreateDnsRecord(v bool)`

SetAutoCreateDnsRecord sets AutoCreateDnsRecord field to given value.

### HasAutoCreateDnsRecord

`func (o *Vdiscoverytask) HasAutoCreateDnsRecord() bool`

HasAutoCreateDnsRecord returns a boolean if a field has been set.

### GetAutoCreateDnsRecordType

`func (o *Vdiscoverytask) GetAutoCreateDnsRecordType() string`

GetAutoCreateDnsRecordType returns the AutoCreateDnsRecordType field if non-nil, zero value otherwise.

### GetAutoCreateDnsRecordTypeOk

`func (o *Vdiscoverytask) GetAutoCreateDnsRecordTypeOk() (*string, bool)`

GetAutoCreateDnsRecordTypeOk returns a tuple with the AutoCreateDnsRecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateDnsRecordType

`func (o *Vdiscoverytask) SetAutoCreateDnsRecordType(v string)`

SetAutoCreateDnsRecordType sets AutoCreateDnsRecordType field to given value.

### HasAutoCreateDnsRecordType

`func (o *Vdiscoverytask) HasAutoCreateDnsRecordType() bool`

HasAutoCreateDnsRecordType returns a boolean if a field has been set.

### GetCdiscoveryFileToken

`func (o *Vdiscoverytask) GetCdiscoveryFileToken() string`

GetCdiscoveryFileToken returns the CdiscoveryFileToken field if non-nil, zero value otherwise.

### GetCdiscoveryFileTokenOk

`func (o *Vdiscoverytask) GetCdiscoveryFileTokenOk() (*string, bool)`

GetCdiscoveryFileTokenOk returns a tuple with the CdiscoveryFileToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdiscoveryFileToken

`func (o *Vdiscoverytask) SetCdiscoveryFileToken(v string)`

SetCdiscoveryFileToken sets CdiscoveryFileToken field to given value.

### HasCdiscoveryFileToken

`func (o *Vdiscoverytask) HasCdiscoveryFileToken() bool`

HasCdiscoveryFileToken returns a boolean if a field has been set.

### GetComment

`func (o *Vdiscoverytask) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Vdiscoverytask) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Vdiscoverytask) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Vdiscoverytask) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCredentialsType

`func (o *Vdiscoverytask) GetCredentialsType() string`

GetCredentialsType returns the CredentialsType field if non-nil, zero value otherwise.

### GetCredentialsTypeOk

`func (o *Vdiscoverytask) GetCredentialsTypeOk() (*string, bool)`

GetCredentialsTypeOk returns a tuple with the CredentialsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsType

`func (o *Vdiscoverytask) SetCredentialsType(v string)`

SetCredentialsType sets CredentialsType field to given value.

### HasCredentialsType

`func (o *Vdiscoverytask) HasCredentialsType() bool`

HasCredentialsType returns a boolean if a field has been set.

### GetDnsViewPrivateIp

`func (o *Vdiscoverytask) GetDnsViewPrivateIp() string`

GetDnsViewPrivateIp returns the DnsViewPrivateIp field if non-nil, zero value otherwise.

### GetDnsViewPrivateIpOk

`func (o *Vdiscoverytask) GetDnsViewPrivateIpOk() (*string, bool)`

GetDnsViewPrivateIpOk returns a tuple with the DnsViewPrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsViewPrivateIp

`func (o *Vdiscoverytask) SetDnsViewPrivateIp(v string)`

SetDnsViewPrivateIp sets DnsViewPrivateIp field to given value.

### HasDnsViewPrivateIp

`func (o *Vdiscoverytask) HasDnsViewPrivateIp() bool`

HasDnsViewPrivateIp returns a boolean if a field has been set.

### GetDnsViewPublicIp

`func (o *Vdiscoverytask) GetDnsViewPublicIp() string`

GetDnsViewPublicIp returns the DnsViewPublicIp field if non-nil, zero value otherwise.

### GetDnsViewPublicIpOk

`func (o *Vdiscoverytask) GetDnsViewPublicIpOk() (*string, bool)`

GetDnsViewPublicIpOk returns a tuple with the DnsViewPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsViewPublicIp

`func (o *Vdiscoverytask) SetDnsViewPublicIp(v string)`

SetDnsViewPublicIp sets DnsViewPublicIp field to given value.

### HasDnsViewPublicIp

`func (o *Vdiscoverytask) HasDnsViewPublicIp() bool`

HasDnsViewPublicIp returns a boolean if a field has been set.

### GetDomainName

`func (o *Vdiscoverytask) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *Vdiscoverytask) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *Vdiscoverytask) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *Vdiscoverytask) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDriverType

`func (o *Vdiscoverytask) GetDriverType() string`

GetDriverType returns the DriverType field if non-nil, zero value otherwise.

### GetDriverTypeOk

`func (o *Vdiscoverytask) GetDriverTypeOk() (*string, bool)`

GetDriverTypeOk returns a tuple with the DriverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverType

`func (o *Vdiscoverytask) SetDriverType(v string)`

SetDriverType sets DriverType field to given value.

### HasDriverType

`func (o *Vdiscoverytask) HasDriverType() bool`

HasDriverType returns a boolean if a field has been set.

### GetEnableFilter

`func (o *Vdiscoverytask) GetEnableFilter() bool`

GetEnableFilter returns the EnableFilter field if non-nil, zero value otherwise.

### GetEnableFilterOk

`func (o *Vdiscoverytask) GetEnableFilterOk() (*bool, bool)`

GetEnableFilterOk returns a tuple with the EnableFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFilter

`func (o *Vdiscoverytask) SetEnableFilter(v bool)`

SetEnableFilter sets EnableFilter field to given value.

### HasEnableFilter

`func (o *Vdiscoverytask) HasEnableFilter() bool`

HasEnableFilter returns a boolean if a field has been set.

### GetEnabled

`func (o *Vdiscoverytask) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Vdiscoverytask) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Vdiscoverytask) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Vdiscoverytask) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFqdnOrIp

`func (o *Vdiscoverytask) GetFqdnOrIp() string`

GetFqdnOrIp returns the FqdnOrIp field if non-nil, zero value otherwise.

### GetFqdnOrIpOk

`func (o *Vdiscoverytask) GetFqdnOrIpOk() (*string, bool)`

GetFqdnOrIpOk returns a tuple with the FqdnOrIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdnOrIp

`func (o *Vdiscoverytask) SetFqdnOrIp(v string)`

SetFqdnOrIp sets FqdnOrIp field to given value.

### HasFqdnOrIp

`func (o *Vdiscoverytask) HasFqdnOrIp() bool`

HasFqdnOrIp returns a boolean if a field has been set.

### GetGovcloudEnabled

`func (o *Vdiscoverytask) GetGovcloudEnabled() bool`

GetGovcloudEnabled returns the GovcloudEnabled field if non-nil, zero value otherwise.

### GetGovcloudEnabledOk

`func (o *Vdiscoverytask) GetGovcloudEnabledOk() (*bool, bool)`

GetGovcloudEnabledOk returns a tuple with the GovcloudEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGovcloudEnabled

`func (o *Vdiscoverytask) SetGovcloudEnabled(v bool)`

SetGovcloudEnabled sets GovcloudEnabled field to given value.

### HasGovcloudEnabled

`func (o *Vdiscoverytask) HasGovcloudEnabled() bool`

HasGovcloudEnabled returns a boolean if a field has been set.

### GetIdentityVersion

`func (o *Vdiscoverytask) GetIdentityVersion() string`

GetIdentityVersion returns the IdentityVersion field if non-nil, zero value otherwise.

### GetIdentityVersionOk

`func (o *Vdiscoverytask) GetIdentityVersionOk() (*string, bool)`

GetIdentityVersionOk returns a tuple with the IdentityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityVersion

`func (o *Vdiscoverytask) SetIdentityVersion(v string)`

SetIdentityVersion sets IdentityVersion field to given value.

### HasIdentityVersion

`func (o *Vdiscoverytask) HasIdentityVersion() bool`

HasIdentityVersion returns a boolean if a field has been set.

### GetLastRun

`func (o *Vdiscoverytask) GetLastRun() int64`

GetLastRun returns the LastRun field if non-nil, zero value otherwise.

### GetLastRunOk

`func (o *Vdiscoverytask) GetLastRunOk() (*int64, bool)`

GetLastRunOk returns a tuple with the LastRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRun

`func (o *Vdiscoverytask) SetLastRun(v int64)`

SetLastRun sets LastRun field to given value.

### HasLastRun

`func (o *Vdiscoverytask) HasLastRun() bool`

HasLastRun returns a boolean if a field has been set.

### GetMember

`func (o *Vdiscoverytask) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *Vdiscoverytask) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *Vdiscoverytask) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *Vdiscoverytask) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetMergeData

`func (o *Vdiscoverytask) GetMergeData() bool`

GetMergeData returns the MergeData field if non-nil, zero value otherwise.

### GetMergeDataOk

`func (o *Vdiscoverytask) GetMergeDataOk() (*bool, bool)`

GetMergeDataOk returns a tuple with the MergeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeData

`func (o *Vdiscoverytask) SetMergeData(v bool)`

SetMergeData sets MergeData field to given value.

### HasMergeData

`func (o *Vdiscoverytask) HasMergeData() bool`

HasMergeData returns a boolean if a field has been set.

### GetMultipleAccountsSyncPolicy

`func (o *Vdiscoverytask) GetMultipleAccountsSyncPolicy() string`

GetMultipleAccountsSyncPolicy returns the MultipleAccountsSyncPolicy field if non-nil, zero value otherwise.

### GetMultipleAccountsSyncPolicyOk

`func (o *Vdiscoverytask) GetMultipleAccountsSyncPolicyOk() (*string, bool)`

GetMultipleAccountsSyncPolicyOk returns a tuple with the MultipleAccountsSyncPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleAccountsSyncPolicy

`func (o *Vdiscoverytask) SetMultipleAccountsSyncPolicy(v string)`

SetMultipleAccountsSyncPolicy sets MultipleAccountsSyncPolicy field to given value.

### HasMultipleAccountsSyncPolicy

`func (o *Vdiscoverytask) HasMultipleAccountsSyncPolicy() bool`

HasMultipleAccountsSyncPolicy returns a boolean if a field has been set.

### GetName

`func (o *Vdiscoverytask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Vdiscoverytask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Vdiscoverytask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Vdiscoverytask) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkFilter

`func (o *Vdiscoverytask) GetNetworkFilter() string`

GetNetworkFilter returns the NetworkFilter field if non-nil, zero value otherwise.

### GetNetworkFilterOk

`func (o *Vdiscoverytask) GetNetworkFilterOk() (*string, bool)`

GetNetworkFilterOk returns a tuple with the NetworkFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFilter

`func (o *Vdiscoverytask) SetNetworkFilter(v string)`

SetNetworkFilter sets NetworkFilter field to given value.

### HasNetworkFilter

`func (o *Vdiscoverytask) HasNetworkFilter() bool`

HasNetworkFilter returns a boolean if a field has been set.

### GetNetworkList

`func (o *Vdiscoverytask) GetNetworkList() []string`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *Vdiscoverytask) GetNetworkListOk() (*[]string, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *Vdiscoverytask) SetNetworkList(v []string)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *Vdiscoverytask) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.

### GetPassword

`func (o *Vdiscoverytask) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Vdiscoverytask) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Vdiscoverytask) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Vdiscoverytask) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *Vdiscoverytask) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Vdiscoverytask) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Vdiscoverytask) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *Vdiscoverytask) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPrivateNetworkView

`func (o *Vdiscoverytask) GetPrivateNetworkView() string`

GetPrivateNetworkView returns the PrivateNetworkView field if non-nil, zero value otherwise.

### GetPrivateNetworkViewOk

`func (o *Vdiscoverytask) GetPrivateNetworkViewOk() (*string, bool)`

GetPrivateNetworkViewOk returns a tuple with the PrivateNetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkView

`func (o *Vdiscoverytask) SetPrivateNetworkView(v string)`

SetPrivateNetworkView sets PrivateNetworkView field to given value.

### HasPrivateNetworkView

`func (o *Vdiscoverytask) HasPrivateNetworkView() bool`

HasPrivateNetworkView returns a boolean if a field has been set.

### GetPrivateNetworkViewMappingPolicy

`func (o *Vdiscoverytask) GetPrivateNetworkViewMappingPolicy() string`

GetPrivateNetworkViewMappingPolicy returns the PrivateNetworkViewMappingPolicy field if non-nil, zero value otherwise.

### GetPrivateNetworkViewMappingPolicyOk

`func (o *Vdiscoverytask) GetPrivateNetworkViewMappingPolicyOk() (*string, bool)`

GetPrivateNetworkViewMappingPolicyOk returns a tuple with the PrivateNetworkViewMappingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNetworkViewMappingPolicy

`func (o *Vdiscoverytask) SetPrivateNetworkViewMappingPolicy(v string)`

SetPrivateNetworkViewMappingPolicy sets PrivateNetworkViewMappingPolicy field to given value.

### HasPrivateNetworkViewMappingPolicy

`func (o *Vdiscoverytask) HasPrivateNetworkViewMappingPolicy() bool`

HasPrivateNetworkViewMappingPolicy returns a boolean if a field has been set.

### GetProtocol

`func (o *Vdiscoverytask) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Vdiscoverytask) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Vdiscoverytask) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *Vdiscoverytask) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetPublicNetworkView

`func (o *Vdiscoverytask) GetPublicNetworkView() string`

GetPublicNetworkView returns the PublicNetworkView field if non-nil, zero value otherwise.

### GetPublicNetworkViewOk

`func (o *Vdiscoverytask) GetPublicNetworkViewOk() (*string, bool)`

GetPublicNetworkViewOk returns a tuple with the PublicNetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetworkView

`func (o *Vdiscoverytask) SetPublicNetworkView(v string)`

SetPublicNetworkView sets PublicNetworkView field to given value.

### HasPublicNetworkView

`func (o *Vdiscoverytask) HasPublicNetworkView() bool`

HasPublicNetworkView returns a boolean if a field has been set.

### GetPublicNetworkViewMappingPolicy

`func (o *Vdiscoverytask) GetPublicNetworkViewMappingPolicy() string`

GetPublicNetworkViewMappingPolicy returns the PublicNetworkViewMappingPolicy field if non-nil, zero value otherwise.

### GetPublicNetworkViewMappingPolicyOk

`func (o *Vdiscoverytask) GetPublicNetworkViewMappingPolicyOk() (*string, bool)`

GetPublicNetworkViewMappingPolicyOk returns a tuple with the PublicNetworkViewMappingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNetworkViewMappingPolicy

`func (o *Vdiscoverytask) SetPublicNetworkViewMappingPolicy(v string)`

SetPublicNetworkViewMappingPolicy sets PublicNetworkViewMappingPolicy field to given value.

### HasPublicNetworkViewMappingPolicy

`func (o *Vdiscoverytask) HasPublicNetworkViewMappingPolicy() bool`

HasPublicNetworkViewMappingPolicy returns a boolean if a field has been set.

### GetRoleArn

`func (o *Vdiscoverytask) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *Vdiscoverytask) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *Vdiscoverytask) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *Vdiscoverytask) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetScheduledRun

`func (o *Vdiscoverytask) GetScheduledRun() VdiscoverytaskScheduledRun`

GetScheduledRun returns the ScheduledRun field if non-nil, zero value otherwise.

### GetScheduledRunOk

`func (o *Vdiscoverytask) GetScheduledRunOk() (*VdiscoverytaskScheduledRun, bool)`

GetScheduledRunOk returns a tuple with the ScheduledRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledRun

`func (o *Vdiscoverytask) SetScheduledRun(v VdiscoverytaskScheduledRun)`

SetScheduledRun sets ScheduledRun field to given value.

### HasScheduledRun

`func (o *Vdiscoverytask) HasScheduledRun() bool`

HasScheduledRun returns a boolean if a field has been set.

### GetSelectedRegions

`func (o *Vdiscoverytask) GetSelectedRegions() string`

GetSelectedRegions returns the SelectedRegions field if non-nil, zero value otherwise.

### GetSelectedRegionsOk

`func (o *Vdiscoverytask) GetSelectedRegionsOk() (*string, bool)`

GetSelectedRegionsOk returns a tuple with the SelectedRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRegions

`func (o *Vdiscoverytask) SetSelectedRegions(v string)`

SetSelectedRegions sets SelectedRegions field to given value.

### HasSelectedRegions

`func (o *Vdiscoverytask) HasSelectedRegions() bool`

HasSelectedRegions returns a boolean if a field has been set.

### GetServiceAccountFile

`func (o *Vdiscoverytask) GetServiceAccountFile() string`

GetServiceAccountFile returns the ServiceAccountFile field if non-nil, zero value otherwise.

### GetServiceAccountFileOk

`func (o *Vdiscoverytask) GetServiceAccountFileOk() (*string, bool)`

GetServiceAccountFileOk returns a tuple with the ServiceAccountFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountFile

`func (o *Vdiscoverytask) SetServiceAccountFile(v string)`

SetServiceAccountFile sets ServiceAccountFile field to given value.

### HasServiceAccountFile

`func (o *Vdiscoverytask) HasServiceAccountFile() bool`

HasServiceAccountFile returns a boolean if a field has been set.

### GetServiceAccountFileToken

`func (o *Vdiscoverytask) GetServiceAccountFileToken() string`

GetServiceAccountFileToken returns the ServiceAccountFileToken field if non-nil, zero value otherwise.

### GetServiceAccountFileTokenOk

`func (o *Vdiscoverytask) GetServiceAccountFileTokenOk() (*string, bool)`

GetServiceAccountFileTokenOk returns a tuple with the ServiceAccountFileToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountFileToken

`func (o *Vdiscoverytask) SetServiceAccountFileToken(v string)`

SetServiceAccountFileToken sets ServiceAccountFileToken field to given value.

### HasServiceAccountFileToken

`func (o *Vdiscoverytask) HasServiceAccountFileToken() bool`

HasServiceAccountFileToken returns a boolean if a field has been set.

### GetState

`func (o *Vdiscoverytask) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Vdiscoverytask) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Vdiscoverytask) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Vdiscoverytask) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateMsg

`func (o *Vdiscoverytask) GetStateMsg() string`

GetStateMsg returns the StateMsg field if non-nil, zero value otherwise.

### GetStateMsgOk

`func (o *Vdiscoverytask) GetStateMsgOk() (*string, bool)`

GetStateMsgOk returns a tuple with the StateMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMsg

`func (o *Vdiscoverytask) SetStateMsg(v string)`

SetStateMsg sets StateMsg field to given value.

### HasStateMsg

`func (o *Vdiscoverytask) HasStateMsg() bool`

HasStateMsg returns a boolean if a field has been set.

### GetSyncChildAccounts

`func (o *Vdiscoverytask) GetSyncChildAccounts() bool`

GetSyncChildAccounts returns the SyncChildAccounts field if non-nil, zero value otherwise.

### GetSyncChildAccountsOk

`func (o *Vdiscoverytask) GetSyncChildAccountsOk() (*bool, bool)`

GetSyncChildAccountsOk returns a tuple with the SyncChildAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncChildAccounts

`func (o *Vdiscoverytask) SetSyncChildAccounts(v bool)`

SetSyncChildAccounts sets SyncChildAccounts field to given value.

### HasSyncChildAccounts

`func (o *Vdiscoverytask) HasSyncChildAccounts() bool`

HasSyncChildAccounts returns a boolean if a field has been set.

### GetUpdateDnsViewPrivateIp

`func (o *Vdiscoverytask) GetUpdateDnsViewPrivateIp() bool`

GetUpdateDnsViewPrivateIp returns the UpdateDnsViewPrivateIp field if non-nil, zero value otherwise.

### GetUpdateDnsViewPrivateIpOk

`func (o *Vdiscoverytask) GetUpdateDnsViewPrivateIpOk() (*bool, bool)`

GetUpdateDnsViewPrivateIpOk returns a tuple with the UpdateDnsViewPrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDnsViewPrivateIp

`func (o *Vdiscoverytask) SetUpdateDnsViewPrivateIp(v bool)`

SetUpdateDnsViewPrivateIp sets UpdateDnsViewPrivateIp field to given value.

### HasUpdateDnsViewPrivateIp

`func (o *Vdiscoverytask) HasUpdateDnsViewPrivateIp() bool`

HasUpdateDnsViewPrivateIp returns a boolean if a field has been set.

### GetUpdateDnsViewPublicIp

`func (o *Vdiscoverytask) GetUpdateDnsViewPublicIp() bool`

GetUpdateDnsViewPublicIp returns the UpdateDnsViewPublicIp field if non-nil, zero value otherwise.

### GetUpdateDnsViewPublicIpOk

`func (o *Vdiscoverytask) GetUpdateDnsViewPublicIpOk() (*bool, bool)`

GetUpdateDnsViewPublicIpOk returns a tuple with the UpdateDnsViewPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDnsViewPublicIp

`func (o *Vdiscoverytask) SetUpdateDnsViewPublicIp(v bool)`

SetUpdateDnsViewPublicIp sets UpdateDnsViewPublicIp field to given value.

### HasUpdateDnsViewPublicIp

`func (o *Vdiscoverytask) HasUpdateDnsViewPublicIp() bool`

HasUpdateDnsViewPublicIp returns a boolean if a field has been set.

### GetUpdateMetadata

`func (o *Vdiscoverytask) GetUpdateMetadata() bool`

GetUpdateMetadata returns the UpdateMetadata field if non-nil, zero value otherwise.

### GetUpdateMetadataOk

`func (o *Vdiscoverytask) GetUpdateMetadataOk() (*bool, bool)`

GetUpdateMetadataOk returns a tuple with the UpdateMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateMetadata

`func (o *Vdiscoverytask) SetUpdateMetadata(v bool)`

SetUpdateMetadata sets UpdateMetadata field to given value.

### HasUpdateMetadata

`func (o *Vdiscoverytask) HasUpdateMetadata() bool`

HasUpdateMetadata returns a boolean if a field has been set.

### GetUseIdentity

`func (o *Vdiscoverytask) GetUseIdentity() bool`

GetUseIdentity returns the UseIdentity field if non-nil, zero value otherwise.

### GetUseIdentityOk

`func (o *Vdiscoverytask) GetUseIdentityOk() (*bool, bool)`

GetUseIdentityOk returns a tuple with the UseIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseIdentity

`func (o *Vdiscoverytask) SetUseIdentity(v bool)`

SetUseIdentity sets UseIdentity field to given value.

### HasUseIdentity

`func (o *Vdiscoverytask) HasUseIdentity() bool`

HasUseIdentity returns a boolean if a field has been set.

### GetUsername

`func (o *Vdiscoverytask) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Vdiscoverytask) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Vdiscoverytask) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Vdiscoverytask) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUuid

`func (o *Vdiscoverytask) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Vdiscoverytask) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Vdiscoverytask) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Vdiscoverytask) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


