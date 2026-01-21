# GetParentalcontrolSubscriberrecordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AccountingSessionId** | Pointer to **string** | accounting_session_id | [optional] 
**AltIpAddr** | Pointer to **string** | alt_ip_addr | [optional] 
**Ans0** | Pointer to **string** | ans0 | [optional] 
**Ans1** | Pointer to **string** | ans1 | [optional] 
**Ans2** | Pointer to **string** | ans2 | [optional] 
**Ans3** | Pointer to **string** | ans3 | [optional] 
**Ans4** | Pointer to **string** | ans4 | [optional] 
**BlackList** | Pointer to **string** | black_list | [optional] 
**Bwflag** | Pointer to **bool** | bwflag | [optional] 
**DynamicCategoryPolicy** | Pointer to **bool** | dynamic_category_policy | [optional] 
**Flags** | Pointer to **string** | flags | [optional] 
**IpAddr** | Pointer to **string** | ip_addr | [optional] 
**Ipsd** | Pointer to **string** | ipsd | [optional] 
**Localid** | Pointer to **string** | localid | [optional] 
**NasContextual** | Pointer to **string** | nas_contextual | [optional] 
**OpCode** | Pointer to **string** | op_code | [optional] 
**ParentalControlPolicy** | Pointer to **string** | parental_control_policy | [optional] 
**Prefix** | Pointer to **int64** | prefix | [optional] 
**ProxyAll** | Pointer to **bool** | proxy_all | [optional] 
**Site** | Pointer to **string** | site | [optional] 
**SubscriberId** | Pointer to **string** | subscriber_id | [optional] 
**SubscriberSecurePolicy** | Pointer to **string** | subscriber_secure_policy | [optional] 
**UnknownCategoryPolicy** | Pointer to **bool** | unknown_category_policy | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**WhiteList** | Pointer to **string** | white_list | [optional] 
**WpcCategoryPolicy** | Pointer to **string** | wpc_category_policy | [optional] 
**Result** | Pointer to [**ParentalcontrolSubscriberrecord**](ParentalcontrolSubscriberrecord.md) |  | [optional] 

## Methods

### NewGetParentalcontrolSubscriberrecordResponse

`func NewGetParentalcontrolSubscriberrecordResponse() *GetParentalcontrolSubscriberrecordResponse`

NewGetParentalcontrolSubscriberrecordResponse instantiates a new GetParentalcontrolSubscriberrecordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetParentalcontrolSubscriberrecordResponseWithDefaults

`func NewGetParentalcontrolSubscriberrecordResponseWithDefaults() *GetParentalcontrolSubscriberrecordResponse`

NewGetParentalcontrolSubscriberrecordResponseWithDefaults instantiates a new GetParentalcontrolSubscriberrecordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetParentalcontrolSubscriberrecordResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetParentalcontrolSubscriberrecordResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetParentalcontrolSubscriberrecordResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAccountingSessionId

`func (o *GetParentalcontrolSubscriberrecordResponse) GetAccountingSessionId() string`

GetAccountingSessionId returns the AccountingSessionId field if non-nil, zero value otherwise.

### GetAccountingSessionIdOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetAccountingSessionIdOk() (*string, bool)`

GetAccountingSessionIdOk returns a tuple with the AccountingSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountingSessionId

`func (o *GetParentalcontrolSubscriberrecordResponse) SetAccountingSessionId(v string)`

SetAccountingSessionId sets AccountingSessionId field to given value.

### HasAccountingSessionId

`func (o *GetParentalcontrolSubscriberrecordResponse) HasAccountingSessionId() bool`

HasAccountingSessionId returns a boolean if a field has been set.

### GetAltIpAddr

`func (o *GetParentalcontrolSubscriberrecordResponse) GetAltIpAddr() string`

GetAltIpAddr returns the AltIpAddr field if non-nil, zero value otherwise.

### GetAltIpAddrOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetAltIpAddrOk() (*string, bool)`

GetAltIpAddrOk returns a tuple with the AltIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltIpAddr

`func (o *GetParentalcontrolSubscriberrecordResponse) SetAltIpAddr(v string)`

SetAltIpAddr sets AltIpAddr field to given value.

### HasAltIpAddr

`func (o *GetParentalcontrolSubscriberrecordResponse) HasAltIpAddr() bool`

HasAltIpAddr returns a boolean if a field has been set.

### GetAns0

`func (o *GetParentalcontrolSubscriberrecordResponse) GetAns0() string`

GetAns0 returns the Ans0 field if non-nil, zero value otherwise.

### GetAns0Ok

`func (o *GetParentalcontrolSubscriberrecordResponse) GetAns0Ok() (*string, bool)`

GetAns0Ok returns a tuple with the Ans0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAns0

`func (o *GetParentalcontrolSubscriberrecordResponse) SetAns0(v string)`

SetAns0 sets Ans0 field to given value.

### HasAns0

`func (o *GetParentalcontrolSubscriberrecordResponse) HasAns0() bool`

HasAns0 returns a boolean if a field has been set.

### GetAns1

`func (o *GetParentalcontrolSubscriberrecordResponse) GetAns1() string`

GetAns1 returns the Ans1 field if non-nil, zero value otherwise.

### GetAns1Ok

`func (o *GetParentalcontrolSubscriberrecordResponse) GetAns1Ok() (*string, bool)`

GetAns1Ok returns a tuple with the Ans1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAns1

`func (o *GetParentalcontrolSubscriberrecordResponse) SetAns1(v string)`

SetAns1 sets Ans1 field to given value.

### HasAns1

`func (o *GetParentalcontrolSubscriberrecordResponse) HasAns1() bool`

HasAns1 returns a boolean if a field has been set.

### GetAns2

`func (o *GetParentalcontrolSubscriberrecordResponse) GetAns2() string`

GetAns2 returns the Ans2 field if non-nil, zero value otherwise.

### GetAns2Ok

`func (o *GetParentalcontrolSubscriberrecordResponse) GetAns2Ok() (*string, bool)`

GetAns2Ok returns a tuple with the Ans2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAns2

`func (o *GetParentalcontrolSubscriberrecordResponse) SetAns2(v string)`

SetAns2 sets Ans2 field to given value.

### HasAns2

`func (o *GetParentalcontrolSubscriberrecordResponse) HasAns2() bool`

HasAns2 returns a boolean if a field has been set.

### GetAns3

`func (o *GetParentalcontrolSubscriberrecordResponse) GetAns3() string`

GetAns3 returns the Ans3 field if non-nil, zero value otherwise.

### GetAns3Ok

`func (o *GetParentalcontrolSubscriberrecordResponse) GetAns3Ok() (*string, bool)`

GetAns3Ok returns a tuple with the Ans3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAns3

`func (o *GetParentalcontrolSubscriberrecordResponse) SetAns3(v string)`

SetAns3 sets Ans3 field to given value.

### HasAns3

`func (o *GetParentalcontrolSubscriberrecordResponse) HasAns3() bool`

HasAns3 returns a boolean if a field has been set.

### GetAns4

`func (o *GetParentalcontrolSubscriberrecordResponse) GetAns4() string`

GetAns4 returns the Ans4 field if non-nil, zero value otherwise.

### GetAns4Ok

`func (o *GetParentalcontrolSubscriberrecordResponse) GetAns4Ok() (*string, bool)`

GetAns4Ok returns a tuple with the Ans4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAns4

`func (o *GetParentalcontrolSubscriberrecordResponse) SetAns4(v string)`

SetAns4 sets Ans4 field to given value.

### HasAns4

`func (o *GetParentalcontrolSubscriberrecordResponse) HasAns4() bool`

HasAns4 returns a boolean if a field has been set.

### GetBlackList

`func (o *GetParentalcontrolSubscriberrecordResponse) GetBlackList() string`

GetBlackList returns the BlackList field if non-nil, zero value otherwise.

### GetBlackListOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetBlackListOk() (*string, bool)`

GetBlackListOk returns a tuple with the BlackList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlackList

`func (o *GetParentalcontrolSubscriberrecordResponse) SetBlackList(v string)`

SetBlackList sets BlackList field to given value.

### HasBlackList

`func (o *GetParentalcontrolSubscriberrecordResponse) HasBlackList() bool`

HasBlackList returns a boolean if a field has been set.

### GetBwflag

`func (o *GetParentalcontrolSubscriberrecordResponse) GetBwflag() bool`

GetBwflag returns the Bwflag field if non-nil, zero value otherwise.

### GetBwflagOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetBwflagOk() (*bool, bool)`

GetBwflagOk returns a tuple with the Bwflag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwflag

`func (o *GetParentalcontrolSubscriberrecordResponse) SetBwflag(v bool)`

SetBwflag sets Bwflag field to given value.

### HasBwflag

`func (o *GetParentalcontrolSubscriberrecordResponse) HasBwflag() bool`

HasBwflag returns a boolean if a field has been set.

### GetDynamicCategoryPolicy

`func (o *GetParentalcontrolSubscriberrecordResponse) GetDynamicCategoryPolicy() bool`

GetDynamicCategoryPolicy returns the DynamicCategoryPolicy field if non-nil, zero value otherwise.

### GetDynamicCategoryPolicyOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetDynamicCategoryPolicyOk() (*bool, bool)`

GetDynamicCategoryPolicyOk returns a tuple with the DynamicCategoryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicCategoryPolicy

`func (o *GetParentalcontrolSubscriberrecordResponse) SetDynamicCategoryPolicy(v bool)`

SetDynamicCategoryPolicy sets DynamicCategoryPolicy field to given value.

### HasDynamicCategoryPolicy

`func (o *GetParentalcontrolSubscriberrecordResponse) HasDynamicCategoryPolicy() bool`

HasDynamicCategoryPolicy returns a boolean if a field has been set.

### GetFlags

`func (o *GetParentalcontrolSubscriberrecordResponse) GetFlags() string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetFlagsOk() (*string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *GetParentalcontrolSubscriberrecordResponse) SetFlags(v string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *GetParentalcontrolSubscriberrecordResponse) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetIpAddr

`func (o *GetParentalcontrolSubscriberrecordResponse) GetIpAddr() string`

GetIpAddr returns the IpAddr field if non-nil, zero value otherwise.

### GetIpAddrOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetIpAddrOk() (*string, bool)`

GetIpAddrOk returns a tuple with the IpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddr

`func (o *GetParentalcontrolSubscriberrecordResponse) SetIpAddr(v string)`

SetIpAddr sets IpAddr field to given value.

### HasIpAddr

`func (o *GetParentalcontrolSubscriberrecordResponse) HasIpAddr() bool`

HasIpAddr returns a boolean if a field has been set.

### GetIpsd

`func (o *GetParentalcontrolSubscriberrecordResponse) GetIpsd() string`

GetIpsd returns the Ipsd field if non-nil, zero value otherwise.

### GetIpsdOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetIpsdOk() (*string, bool)`

GetIpsdOk returns a tuple with the Ipsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsd

`func (o *GetParentalcontrolSubscriberrecordResponse) SetIpsd(v string)`

SetIpsd sets Ipsd field to given value.

### HasIpsd

`func (o *GetParentalcontrolSubscriberrecordResponse) HasIpsd() bool`

HasIpsd returns a boolean if a field has been set.

### GetLocalid

`func (o *GetParentalcontrolSubscriberrecordResponse) GetLocalid() string`

GetLocalid returns the Localid field if non-nil, zero value otherwise.

### GetLocalidOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetLocalidOk() (*string, bool)`

GetLocalidOk returns a tuple with the Localid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalid

`func (o *GetParentalcontrolSubscriberrecordResponse) SetLocalid(v string)`

SetLocalid sets Localid field to given value.

### HasLocalid

`func (o *GetParentalcontrolSubscriberrecordResponse) HasLocalid() bool`

HasLocalid returns a boolean if a field has been set.

### GetNasContextual

`func (o *GetParentalcontrolSubscriberrecordResponse) GetNasContextual() string`

GetNasContextual returns the NasContextual field if non-nil, zero value otherwise.

### GetNasContextualOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetNasContextualOk() (*string, bool)`

GetNasContextualOk returns a tuple with the NasContextual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNasContextual

`func (o *GetParentalcontrolSubscriberrecordResponse) SetNasContextual(v string)`

SetNasContextual sets NasContextual field to given value.

### HasNasContextual

`func (o *GetParentalcontrolSubscriberrecordResponse) HasNasContextual() bool`

HasNasContextual returns a boolean if a field has been set.

### GetOpCode

`func (o *GetParentalcontrolSubscriberrecordResponse) GetOpCode() string`

GetOpCode returns the OpCode field if non-nil, zero value otherwise.

### GetOpCodeOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetOpCodeOk() (*string, bool)`

GetOpCodeOk returns a tuple with the OpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpCode

`func (o *GetParentalcontrolSubscriberrecordResponse) SetOpCode(v string)`

SetOpCode sets OpCode field to given value.

### HasOpCode

`func (o *GetParentalcontrolSubscriberrecordResponse) HasOpCode() bool`

HasOpCode returns a boolean if a field has been set.

### GetParentalControlPolicy

`func (o *GetParentalcontrolSubscriberrecordResponse) GetParentalControlPolicy() string`

GetParentalControlPolicy returns the ParentalControlPolicy field if non-nil, zero value otherwise.

### GetParentalControlPolicyOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetParentalControlPolicyOk() (*string, bool)`

GetParentalControlPolicyOk returns a tuple with the ParentalControlPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentalControlPolicy

`func (o *GetParentalcontrolSubscriberrecordResponse) SetParentalControlPolicy(v string)`

SetParentalControlPolicy sets ParentalControlPolicy field to given value.

### HasParentalControlPolicy

`func (o *GetParentalcontrolSubscriberrecordResponse) HasParentalControlPolicy() bool`

HasParentalControlPolicy returns a boolean if a field has been set.

### GetPrefix

`func (o *GetParentalcontrolSubscriberrecordResponse) GetPrefix() int64`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetPrefixOk() (*int64, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GetParentalcontrolSubscriberrecordResponse) SetPrefix(v int64)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *GetParentalcontrolSubscriberrecordResponse) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetProxyAll

`func (o *GetParentalcontrolSubscriberrecordResponse) GetProxyAll() bool`

GetProxyAll returns the ProxyAll field if non-nil, zero value otherwise.

### GetProxyAllOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetProxyAllOk() (*bool, bool)`

GetProxyAllOk returns a tuple with the ProxyAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAll

`func (o *GetParentalcontrolSubscriberrecordResponse) SetProxyAll(v bool)`

SetProxyAll sets ProxyAll field to given value.

### HasProxyAll

`func (o *GetParentalcontrolSubscriberrecordResponse) HasProxyAll() bool`

HasProxyAll returns a boolean if a field has been set.

### GetSite

`func (o *GetParentalcontrolSubscriberrecordResponse) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *GetParentalcontrolSubscriberrecordResponse) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *GetParentalcontrolSubscriberrecordResponse) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSubscriberId

`func (o *GetParentalcontrolSubscriberrecordResponse) GetSubscriberId() string`

GetSubscriberId returns the SubscriberId field if non-nil, zero value otherwise.

### GetSubscriberIdOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetSubscriberIdOk() (*string, bool)`

GetSubscriberIdOk returns a tuple with the SubscriberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberId

`func (o *GetParentalcontrolSubscriberrecordResponse) SetSubscriberId(v string)`

SetSubscriberId sets SubscriberId field to given value.

### HasSubscriberId

`func (o *GetParentalcontrolSubscriberrecordResponse) HasSubscriberId() bool`

HasSubscriberId returns a boolean if a field has been set.

### GetSubscriberSecurePolicy

`func (o *GetParentalcontrolSubscriberrecordResponse) GetSubscriberSecurePolicy() string`

GetSubscriberSecurePolicy returns the SubscriberSecurePolicy field if non-nil, zero value otherwise.

### GetSubscriberSecurePolicyOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetSubscriberSecurePolicyOk() (*string, bool)`

GetSubscriberSecurePolicyOk returns a tuple with the SubscriberSecurePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberSecurePolicy

`func (o *GetParentalcontrolSubscriberrecordResponse) SetSubscriberSecurePolicy(v string)`

SetSubscriberSecurePolicy sets SubscriberSecurePolicy field to given value.

### HasSubscriberSecurePolicy

`func (o *GetParentalcontrolSubscriberrecordResponse) HasSubscriberSecurePolicy() bool`

HasSubscriberSecurePolicy returns a boolean if a field has been set.

### GetUnknownCategoryPolicy

`func (o *GetParentalcontrolSubscriberrecordResponse) GetUnknownCategoryPolicy() bool`

GetUnknownCategoryPolicy returns the UnknownCategoryPolicy field if non-nil, zero value otherwise.

### GetUnknownCategoryPolicyOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetUnknownCategoryPolicyOk() (*bool, bool)`

GetUnknownCategoryPolicyOk returns a tuple with the UnknownCategoryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownCategoryPolicy

`func (o *GetParentalcontrolSubscriberrecordResponse) SetUnknownCategoryPolicy(v bool)`

SetUnknownCategoryPolicy sets UnknownCategoryPolicy field to given value.

### HasUnknownCategoryPolicy

`func (o *GetParentalcontrolSubscriberrecordResponse) HasUnknownCategoryPolicy() bool`

HasUnknownCategoryPolicy returns a boolean if a field has been set.

### GetUuid

`func (o *GetParentalcontrolSubscriberrecordResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetParentalcontrolSubscriberrecordResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetParentalcontrolSubscriberrecordResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWhiteList

`func (o *GetParentalcontrolSubscriberrecordResponse) GetWhiteList() string`

GetWhiteList returns the WhiteList field if non-nil, zero value otherwise.

### GetWhiteListOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetWhiteListOk() (*string, bool)`

GetWhiteListOk returns a tuple with the WhiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteList

`func (o *GetParentalcontrolSubscriberrecordResponse) SetWhiteList(v string)`

SetWhiteList sets WhiteList field to given value.

### HasWhiteList

`func (o *GetParentalcontrolSubscriberrecordResponse) HasWhiteList() bool`

HasWhiteList returns a boolean if a field has been set.

### GetWpcCategoryPolicy

`func (o *GetParentalcontrolSubscriberrecordResponse) GetWpcCategoryPolicy() string`

GetWpcCategoryPolicy returns the WpcCategoryPolicy field if non-nil, zero value otherwise.

### GetWpcCategoryPolicyOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetWpcCategoryPolicyOk() (*string, bool)`

GetWpcCategoryPolicyOk returns a tuple with the WpcCategoryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWpcCategoryPolicy

`func (o *GetParentalcontrolSubscriberrecordResponse) SetWpcCategoryPolicy(v string)`

SetWpcCategoryPolicy sets WpcCategoryPolicy field to given value.

### HasWpcCategoryPolicy

`func (o *GetParentalcontrolSubscriberrecordResponse) HasWpcCategoryPolicy() bool`

HasWpcCategoryPolicy returns a boolean if a field has been set.

### GetResult

`func (o *GetParentalcontrolSubscriberrecordResponse) GetResult() ParentalcontrolSubscriberrecord`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetParentalcontrolSubscriberrecordResponse) GetResultOk() (*ParentalcontrolSubscriberrecord, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetParentalcontrolSubscriberrecordResponse) SetResult(v ParentalcontrolSubscriberrecord)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetParentalcontrolSubscriberrecordResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


