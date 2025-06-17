# TimeExpense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Tier1ApprovalFlag** | Pointer to **NullableBool** |  | [optional] 
**Tier2ApprovalFlag** | Pointer to **NullableBool** |  | [optional] 
**DisableTimeEntryFlag** | Pointer to **NullableBool** |  | [optional] 
**RequireTimeNoteFlag** | Pointer to **NullableBool** |  | [optional] 
**RequireExpenseNoteFlag** | Pointer to **NullableBool** |  | [optional] 
**RoundingFactor** | Pointer to **NullableFloat64** |  | [optional] 
**InvoiceStart** | Pointer to **NullableInt32** |  | [optional] 
**DefaultSpecialInvoiceType** | Pointer to **NullableString** |  | [optional] 
**InternalCompany** | [**CompanyReference**](CompanyReference.md) |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTimeExpense

`func NewTimeExpense(internalCompany CompanyReference, ) *TimeExpense`

NewTimeExpense instantiates a new TimeExpense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeExpenseWithDefaults

`func NewTimeExpenseWithDefaults() *TimeExpense`

NewTimeExpenseWithDefaults instantiates a new TimeExpense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimeExpense) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimeExpense) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimeExpense) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TimeExpense) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTier1ApprovalFlag

`func (o *TimeExpense) GetTier1ApprovalFlag() bool`

GetTier1ApprovalFlag returns the Tier1ApprovalFlag field if non-nil, zero value otherwise.

### GetTier1ApprovalFlagOk

`func (o *TimeExpense) GetTier1ApprovalFlagOk() (*bool, bool)`

GetTier1ApprovalFlagOk returns a tuple with the Tier1ApprovalFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier1ApprovalFlag

`func (o *TimeExpense) SetTier1ApprovalFlag(v bool)`

SetTier1ApprovalFlag sets Tier1ApprovalFlag field to given value.

### HasTier1ApprovalFlag

`func (o *TimeExpense) HasTier1ApprovalFlag() bool`

HasTier1ApprovalFlag returns a boolean if a field has been set.

### SetTier1ApprovalFlagNil

`func (o *TimeExpense) SetTier1ApprovalFlagNil(b bool)`

 SetTier1ApprovalFlagNil sets the value for Tier1ApprovalFlag to be an explicit nil

### UnsetTier1ApprovalFlag
`func (o *TimeExpense) UnsetTier1ApprovalFlag()`

UnsetTier1ApprovalFlag ensures that no value is present for Tier1ApprovalFlag, not even an explicit nil
### GetTier2ApprovalFlag

`func (o *TimeExpense) GetTier2ApprovalFlag() bool`

GetTier2ApprovalFlag returns the Tier2ApprovalFlag field if non-nil, zero value otherwise.

### GetTier2ApprovalFlagOk

`func (o *TimeExpense) GetTier2ApprovalFlagOk() (*bool, bool)`

GetTier2ApprovalFlagOk returns a tuple with the Tier2ApprovalFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier2ApprovalFlag

`func (o *TimeExpense) SetTier2ApprovalFlag(v bool)`

SetTier2ApprovalFlag sets Tier2ApprovalFlag field to given value.

### HasTier2ApprovalFlag

`func (o *TimeExpense) HasTier2ApprovalFlag() bool`

HasTier2ApprovalFlag returns a boolean if a field has been set.

### SetTier2ApprovalFlagNil

`func (o *TimeExpense) SetTier2ApprovalFlagNil(b bool)`

 SetTier2ApprovalFlagNil sets the value for Tier2ApprovalFlag to be an explicit nil

### UnsetTier2ApprovalFlag
`func (o *TimeExpense) UnsetTier2ApprovalFlag()`

UnsetTier2ApprovalFlag ensures that no value is present for Tier2ApprovalFlag, not even an explicit nil
### GetDisableTimeEntryFlag

`func (o *TimeExpense) GetDisableTimeEntryFlag() bool`

GetDisableTimeEntryFlag returns the DisableTimeEntryFlag field if non-nil, zero value otherwise.

### GetDisableTimeEntryFlagOk

`func (o *TimeExpense) GetDisableTimeEntryFlagOk() (*bool, bool)`

GetDisableTimeEntryFlagOk returns a tuple with the DisableTimeEntryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableTimeEntryFlag

`func (o *TimeExpense) SetDisableTimeEntryFlag(v bool)`

SetDisableTimeEntryFlag sets DisableTimeEntryFlag field to given value.

### HasDisableTimeEntryFlag

`func (o *TimeExpense) HasDisableTimeEntryFlag() bool`

HasDisableTimeEntryFlag returns a boolean if a field has been set.

### SetDisableTimeEntryFlagNil

`func (o *TimeExpense) SetDisableTimeEntryFlagNil(b bool)`

 SetDisableTimeEntryFlagNil sets the value for DisableTimeEntryFlag to be an explicit nil

### UnsetDisableTimeEntryFlag
`func (o *TimeExpense) UnsetDisableTimeEntryFlag()`

UnsetDisableTimeEntryFlag ensures that no value is present for DisableTimeEntryFlag, not even an explicit nil
### GetRequireTimeNoteFlag

`func (o *TimeExpense) GetRequireTimeNoteFlag() bool`

GetRequireTimeNoteFlag returns the RequireTimeNoteFlag field if non-nil, zero value otherwise.

### GetRequireTimeNoteFlagOk

`func (o *TimeExpense) GetRequireTimeNoteFlagOk() (*bool, bool)`

GetRequireTimeNoteFlagOk returns a tuple with the RequireTimeNoteFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTimeNoteFlag

`func (o *TimeExpense) SetRequireTimeNoteFlag(v bool)`

SetRequireTimeNoteFlag sets RequireTimeNoteFlag field to given value.

### HasRequireTimeNoteFlag

`func (o *TimeExpense) HasRequireTimeNoteFlag() bool`

HasRequireTimeNoteFlag returns a boolean if a field has been set.

### SetRequireTimeNoteFlagNil

`func (o *TimeExpense) SetRequireTimeNoteFlagNil(b bool)`

 SetRequireTimeNoteFlagNil sets the value for RequireTimeNoteFlag to be an explicit nil

### UnsetRequireTimeNoteFlag
`func (o *TimeExpense) UnsetRequireTimeNoteFlag()`

UnsetRequireTimeNoteFlag ensures that no value is present for RequireTimeNoteFlag, not even an explicit nil
### GetRequireExpenseNoteFlag

`func (o *TimeExpense) GetRequireExpenseNoteFlag() bool`

GetRequireExpenseNoteFlag returns the RequireExpenseNoteFlag field if non-nil, zero value otherwise.

### GetRequireExpenseNoteFlagOk

`func (o *TimeExpense) GetRequireExpenseNoteFlagOk() (*bool, bool)`

GetRequireExpenseNoteFlagOk returns a tuple with the RequireExpenseNoteFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireExpenseNoteFlag

`func (o *TimeExpense) SetRequireExpenseNoteFlag(v bool)`

SetRequireExpenseNoteFlag sets RequireExpenseNoteFlag field to given value.

### HasRequireExpenseNoteFlag

`func (o *TimeExpense) HasRequireExpenseNoteFlag() bool`

HasRequireExpenseNoteFlag returns a boolean if a field has been set.

### SetRequireExpenseNoteFlagNil

`func (o *TimeExpense) SetRequireExpenseNoteFlagNil(b bool)`

 SetRequireExpenseNoteFlagNil sets the value for RequireExpenseNoteFlag to be an explicit nil

### UnsetRequireExpenseNoteFlag
`func (o *TimeExpense) UnsetRequireExpenseNoteFlag()`

UnsetRequireExpenseNoteFlag ensures that no value is present for RequireExpenseNoteFlag, not even an explicit nil
### GetRoundingFactor

`func (o *TimeExpense) GetRoundingFactor() float64`

GetRoundingFactor returns the RoundingFactor field if non-nil, zero value otherwise.

### GetRoundingFactorOk

`func (o *TimeExpense) GetRoundingFactorOk() (*float64, bool)`

GetRoundingFactorOk returns a tuple with the RoundingFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundingFactor

`func (o *TimeExpense) SetRoundingFactor(v float64)`

SetRoundingFactor sets RoundingFactor field to given value.

### HasRoundingFactor

`func (o *TimeExpense) HasRoundingFactor() bool`

HasRoundingFactor returns a boolean if a field has been set.

### SetRoundingFactorNil

`func (o *TimeExpense) SetRoundingFactorNil(b bool)`

 SetRoundingFactorNil sets the value for RoundingFactor to be an explicit nil

### UnsetRoundingFactor
`func (o *TimeExpense) UnsetRoundingFactor()`

UnsetRoundingFactor ensures that no value is present for RoundingFactor, not even an explicit nil
### GetInvoiceStart

`func (o *TimeExpense) GetInvoiceStart() int32`

GetInvoiceStart returns the InvoiceStart field if non-nil, zero value otherwise.

### GetInvoiceStartOk

`func (o *TimeExpense) GetInvoiceStartOk() (*int32, bool)`

GetInvoiceStartOk returns a tuple with the InvoiceStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStart

`func (o *TimeExpense) SetInvoiceStart(v int32)`

SetInvoiceStart sets InvoiceStart field to given value.

### HasInvoiceStart

`func (o *TimeExpense) HasInvoiceStart() bool`

HasInvoiceStart returns a boolean if a field has been set.

### SetInvoiceStartNil

`func (o *TimeExpense) SetInvoiceStartNil(b bool)`

 SetInvoiceStartNil sets the value for InvoiceStart to be an explicit nil

### UnsetInvoiceStart
`func (o *TimeExpense) UnsetInvoiceStart()`

UnsetInvoiceStart ensures that no value is present for InvoiceStart, not even an explicit nil
### GetDefaultSpecialInvoiceType

`func (o *TimeExpense) GetDefaultSpecialInvoiceType() string`

GetDefaultSpecialInvoiceType returns the DefaultSpecialInvoiceType field if non-nil, zero value otherwise.

### GetDefaultSpecialInvoiceTypeOk

`func (o *TimeExpense) GetDefaultSpecialInvoiceTypeOk() (*string, bool)`

GetDefaultSpecialInvoiceTypeOk returns a tuple with the DefaultSpecialInvoiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSpecialInvoiceType

`func (o *TimeExpense) SetDefaultSpecialInvoiceType(v string)`

SetDefaultSpecialInvoiceType sets DefaultSpecialInvoiceType field to given value.

### HasDefaultSpecialInvoiceType

`func (o *TimeExpense) HasDefaultSpecialInvoiceType() bool`

HasDefaultSpecialInvoiceType returns a boolean if a field has been set.

### SetDefaultSpecialInvoiceTypeNil

`func (o *TimeExpense) SetDefaultSpecialInvoiceTypeNil(b bool)`

 SetDefaultSpecialInvoiceTypeNil sets the value for DefaultSpecialInvoiceType to be an explicit nil

### UnsetDefaultSpecialInvoiceType
`func (o *TimeExpense) UnsetDefaultSpecialInvoiceType()`

UnsetDefaultSpecialInvoiceType ensures that no value is present for DefaultSpecialInvoiceType, not even an explicit nil
### GetInternalCompany

`func (o *TimeExpense) GetInternalCompany() CompanyReference`

GetInternalCompany returns the InternalCompany field if non-nil, zero value otherwise.

### GetInternalCompanyOk

`func (o *TimeExpense) GetInternalCompanyOk() (*CompanyReference, bool)`

GetInternalCompanyOk returns a tuple with the InternalCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalCompany

`func (o *TimeExpense) SetInternalCompany(v CompanyReference)`

SetInternalCompany sets InternalCompany field to given value.


### GetInfo

`func (o *TimeExpense) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TimeExpense) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TimeExpense) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TimeExpense) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


