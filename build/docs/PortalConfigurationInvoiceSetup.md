# PortalConfigurationInvoiceSetup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**PortalConfiguration** | Pointer to [**PortalConfigurationReference**](PortalConfigurationReference.md) |  | [optional] 
**DisplayInvPmtFlag** | Pointer to **NullableBool** |  | [optional] 
**AllowInvPmtFlag** | Pointer to **NullableBool** |  | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**PaymentProcessor** | Pointer to [**PortalConfigurationPaymentProcessorReference**](PortalConfigurationPaymentProcessorReference.md) |  | [optional] 
**Login** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**UrlOverride** | Pointer to **string** |  | [optional] 
**BillingStatusIds** | Pointer to **[]int32** |  | [optional] 
**AddAllStatuses** | Pointer to **NullableBool** |  | [optional] 
**RemoveAllStatuses** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPortalConfigurationInvoiceSetup

`func NewPortalConfigurationInvoiceSetup() *PortalConfigurationInvoiceSetup`

NewPortalConfigurationInvoiceSetup instantiates a new PortalConfigurationInvoiceSetup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalConfigurationInvoiceSetupWithDefaults

`func NewPortalConfigurationInvoiceSetupWithDefaults() *PortalConfigurationInvoiceSetup`

NewPortalConfigurationInvoiceSetupWithDefaults instantiates a new PortalConfigurationInvoiceSetup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortalConfigurationInvoiceSetup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalConfigurationInvoiceSetup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalConfigurationInvoiceSetup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PortalConfigurationInvoiceSetup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPortalConfiguration

`func (o *PortalConfigurationInvoiceSetup) GetPortalConfiguration() PortalConfigurationReference`

GetPortalConfiguration returns the PortalConfiguration field if non-nil, zero value otherwise.

### GetPortalConfigurationOk

`func (o *PortalConfigurationInvoiceSetup) GetPortalConfigurationOk() (*PortalConfigurationReference, bool)`

GetPortalConfigurationOk returns a tuple with the PortalConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalConfiguration

`func (o *PortalConfigurationInvoiceSetup) SetPortalConfiguration(v PortalConfigurationReference)`

SetPortalConfiguration sets PortalConfiguration field to given value.

### HasPortalConfiguration

`func (o *PortalConfigurationInvoiceSetup) HasPortalConfiguration() bool`

HasPortalConfiguration returns a boolean if a field has been set.

### GetDisplayInvPmtFlag

`func (o *PortalConfigurationInvoiceSetup) GetDisplayInvPmtFlag() bool`

GetDisplayInvPmtFlag returns the DisplayInvPmtFlag field if non-nil, zero value otherwise.

### GetDisplayInvPmtFlagOk

`func (o *PortalConfigurationInvoiceSetup) GetDisplayInvPmtFlagOk() (*bool, bool)`

GetDisplayInvPmtFlagOk returns a tuple with the DisplayInvPmtFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInvPmtFlag

`func (o *PortalConfigurationInvoiceSetup) SetDisplayInvPmtFlag(v bool)`

SetDisplayInvPmtFlag sets DisplayInvPmtFlag field to given value.

### HasDisplayInvPmtFlag

`func (o *PortalConfigurationInvoiceSetup) HasDisplayInvPmtFlag() bool`

HasDisplayInvPmtFlag returns a boolean if a field has been set.

### SetDisplayInvPmtFlagNil

`func (o *PortalConfigurationInvoiceSetup) SetDisplayInvPmtFlagNil(b bool)`

 SetDisplayInvPmtFlagNil sets the value for DisplayInvPmtFlag to be an explicit nil

### UnsetDisplayInvPmtFlag
`func (o *PortalConfigurationInvoiceSetup) UnsetDisplayInvPmtFlag()`

UnsetDisplayInvPmtFlag ensures that no value is present for DisplayInvPmtFlag, not even an explicit nil
### GetAllowInvPmtFlag

`func (o *PortalConfigurationInvoiceSetup) GetAllowInvPmtFlag() bool`

GetAllowInvPmtFlag returns the AllowInvPmtFlag field if non-nil, zero value otherwise.

### GetAllowInvPmtFlagOk

`func (o *PortalConfigurationInvoiceSetup) GetAllowInvPmtFlagOk() (*bool, bool)`

GetAllowInvPmtFlagOk returns a tuple with the AllowInvPmtFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowInvPmtFlag

`func (o *PortalConfigurationInvoiceSetup) SetAllowInvPmtFlag(v bool)`

SetAllowInvPmtFlag sets AllowInvPmtFlag field to given value.

### HasAllowInvPmtFlag

`func (o *PortalConfigurationInvoiceSetup) HasAllowInvPmtFlag() bool`

HasAllowInvPmtFlag returns a boolean if a field has been set.

### SetAllowInvPmtFlagNil

`func (o *PortalConfigurationInvoiceSetup) SetAllowInvPmtFlagNil(b bool)`

 SetAllowInvPmtFlagNil sets the value for AllowInvPmtFlag to be an explicit nil

### UnsetAllowInvPmtFlag
`func (o *PortalConfigurationInvoiceSetup) UnsetAllowInvPmtFlag()`

UnsetAllowInvPmtFlag ensures that no value is present for AllowInvPmtFlag, not even an explicit nil
### GetLocation

`func (o *PortalConfigurationInvoiceSetup) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PortalConfigurationInvoiceSetup) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PortalConfigurationInvoiceSetup) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PortalConfigurationInvoiceSetup) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPaymentProcessor

`func (o *PortalConfigurationInvoiceSetup) GetPaymentProcessor() PortalConfigurationPaymentProcessorReference`

GetPaymentProcessor returns the PaymentProcessor field if non-nil, zero value otherwise.

### GetPaymentProcessorOk

`func (o *PortalConfigurationInvoiceSetup) GetPaymentProcessorOk() (*PortalConfigurationPaymentProcessorReference, bool)`

GetPaymentProcessorOk returns a tuple with the PaymentProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentProcessor

`func (o *PortalConfigurationInvoiceSetup) SetPaymentProcessor(v PortalConfigurationPaymentProcessorReference)`

SetPaymentProcessor sets PaymentProcessor field to given value.

### HasPaymentProcessor

`func (o *PortalConfigurationInvoiceSetup) HasPaymentProcessor() bool`

HasPaymentProcessor returns a boolean if a field has been set.

### GetLogin

`func (o *PortalConfigurationInvoiceSetup) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *PortalConfigurationInvoiceSetup) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *PortalConfigurationInvoiceSetup) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *PortalConfigurationInvoiceSetup) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetPassword

`func (o *PortalConfigurationInvoiceSetup) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PortalConfigurationInvoiceSetup) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PortalConfigurationInvoiceSetup) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PortalConfigurationInvoiceSetup) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUrlOverride

`func (o *PortalConfigurationInvoiceSetup) GetUrlOverride() string`

GetUrlOverride returns the UrlOverride field if non-nil, zero value otherwise.

### GetUrlOverrideOk

`func (o *PortalConfigurationInvoiceSetup) GetUrlOverrideOk() (*string, bool)`

GetUrlOverrideOk returns a tuple with the UrlOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlOverride

`func (o *PortalConfigurationInvoiceSetup) SetUrlOverride(v string)`

SetUrlOverride sets UrlOverride field to given value.

### HasUrlOverride

`func (o *PortalConfigurationInvoiceSetup) HasUrlOverride() bool`

HasUrlOverride returns a boolean if a field has been set.

### GetBillingStatusIds

`func (o *PortalConfigurationInvoiceSetup) GetBillingStatusIds() []int32`

GetBillingStatusIds returns the BillingStatusIds field if non-nil, zero value otherwise.

### GetBillingStatusIdsOk

`func (o *PortalConfigurationInvoiceSetup) GetBillingStatusIdsOk() (*[]int32, bool)`

GetBillingStatusIdsOk returns a tuple with the BillingStatusIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStatusIds

`func (o *PortalConfigurationInvoiceSetup) SetBillingStatusIds(v []int32)`

SetBillingStatusIds sets BillingStatusIds field to given value.

### HasBillingStatusIds

`func (o *PortalConfigurationInvoiceSetup) HasBillingStatusIds() bool`

HasBillingStatusIds returns a boolean if a field has been set.

### GetAddAllStatuses

`func (o *PortalConfigurationInvoiceSetup) GetAddAllStatuses() bool`

GetAddAllStatuses returns the AddAllStatuses field if non-nil, zero value otherwise.

### GetAddAllStatusesOk

`func (o *PortalConfigurationInvoiceSetup) GetAddAllStatusesOk() (*bool, bool)`

GetAddAllStatusesOk returns a tuple with the AddAllStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAllStatuses

`func (o *PortalConfigurationInvoiceSetup) SetAddAllStatuses(v bool)`

SetAddAllStatuses sets AddAllStatuses field to given value.

### HasAddAllStatuses

`func (o *PortalConfigurationInvoiceSetup) HasAddAllStatuses() bool`

HasAddAllStatuses returns a boolean if a field has been set.

### SetAddAllStatusesNil

`func (o *PortalConfigurationInvoiceSetup) SetAddAllStatusesNil(b bool)`

 SetAddAllStatusesNil sets the value for AddAllStatuses to be an explicit nil

### UnsetAddAllStatuses
`func (o *PortalConfigurationInvoiceSetup) UnsetAddAllStatuses()`

UnsetAddAllStatuses ensures that no value is present for AddAllStatuses, not even an explicit nil
### GetRemoveAllStatuses

`func (o *PortalConfigurationInvoiceSetup) GetRemoveAllStatuses() bool`

GetRemoveAllStatuses returns the RemoveAllStatuses field if non-nil, zero value otherwise.

### GetRemoveAllStatusesOk

`func (o *PortalConfigurationInvoiceSetup) GetRemoveAllStatusesOk() (*bool, bool)`

GetRemoveAllStatusesOk returns a tuple with the RemoveAllStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAllStatuses

`func (o *PortalConfigurationInvoiceSetup) SetRemoveAllStatuses(v bool)`

SetRemoveAllStatuses sets RemoveAllStatuses field to given value.

### HasRemoveAllStatuses

`func (o *PortalConfigurationInvoiceSetup) HasRemoveAllStatuses() bool`

HasRemoveAllStatuses returns a boolean if a field has been set.

### SetRemoveAllStatusesNil

`func (o *PortalConfigurationInvoiceSetup) SetRemoveAllStatusesNil(b bool)`

 SetRemoveAllStatusesNil sets the value for RemoveAllStatuses to be an explicit nil

### UnsetRemoveAllStatuses
`func (o *PortalConfigurationInvoiceSetup) UnsetRemoveAllStatuses()`

UnsetRemoveAllStatuses ensures that no value is present for RemoveAllStatuses, not even an explicit nil
### GetInfo

`func (o *PortalConfigurationInvoiceSetup) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PortalConfigurationInvoiceSetup) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PortalConfigurationInvoiceSetup) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PortalConfigurationInvoiceSetup) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


