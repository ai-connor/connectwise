# ProcurementSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**StartingPurchaseOrderNum** | **int32** |  | 
**PurchaseOrderPrefix** | Pointer to **string** |  Max length: 5; | [optional] 
**PurchaseOrderSuffix** | Pointer to **string** |  Max length: 5; | [optional] 
**PrefixSuffixType** | Pointer to **NullableString** |  | [optional] 
**DisableCostUpdatesFlag** | Pointer to **NullableBool** |  | [optional] 
**DisableNegativeInventoryFlag** | Pointer to **NullableBool** |  | [optional] 
**CostingMethod** | **NullableString** |  | 
**AutoClosePurchaseOrderFlag** | Pointer to **NullableBool** |  | [optional] 
**AutoClosePurchaseOrderItemFlag** | Pointer to **NullableBool** |  | [optional] 
**AutoApprovePurchaseOrderFlag** | Pointer to **NullableBool** |  | [optional] 
**TaxPurchaseOrderFlag** | Pointer to **NullableBool** |  | [optional] 
**TaxFreightFlag** | Pointer to **NullableBool** |  | [optional] 
**UseVendorTaxCodeFlag** | Pointer to **NullableBool** |  | [optional] 
**NumDecimalPlaces** | Pointer to **NullableInt32** |  | [optional] 
**DisableAutoPickFlag** | Pointer to **NullableBool** |  | [optional] 
**DefaultProductTaxableFlag** | Pointer to **NullableBool** |  | [optional] 
**EoriNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**NotificationForChangesInShippingInfoFlag** | Pointer to **NullableBool** |  | [optional] 
**ShippingInfoNotificationEmail** | Pointer to **string** |  Max length: 250; | [optional] 

## Methods

### NewProcurementSetting

`func NewProcurementSetting(startingPurchaseOrderNum int32, costingMethod NullableString, ) *ProcurementSetting`

NewProcurementSetting instantiates a new ProcurementSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcurementSettingWithDefaults

`func NewProcurementSettingWithDefaults() *ProcurementSetting`

NewProcurementSettingWithDefaults instantiates a new ProcurementSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProcurementSetting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProcurementSetting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProcurementSetting) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProcurementSetting) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartingPurchaseOrderNum

`func (o *ProcurementSetting) GetStartingPurchaseOrderNum() int32`

GetStartingPurchaseOrderNum returns the StartingPurchaseOrderNum field if non-nil, zero value otherwise.

### GetStartingPurchaseOrderNumOk

`func (o *ProcurementSetting) GetStartingPurchaseOrderNumOk() (*int32, bool)`

GetStartingPurchaseOrderNumOk returns a tuple with the StartingPurchaseOrderNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingPurchaseOrderNum

`func (o *ProcurementSetting) SetStartingPurchaseOrderNum(v int32)`

SetStartingPurchaseOrderNum sets StartingPurchaseOrderNum field to given value.


### GetPurchaseOrderPrefix

`func (o *ProcurementSetting) GetPurchaseOrderPrefix() string`

GetPurchaseOrderPrefix returns the PurchaseOrderPrefix field if non-nil, zero value otherwise.

### GetPurchaseOrderPrefixOk

`func (o *ProcurementSetting) GetPurchaseOrderPrefixOk() (*string, bool)`

GetPurchaseOrderPrefixOk returns a tuple with the PurchaseOrderPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderPrefix

`func (o *ProcurementSetting) SetPurchaseOrderPrefix(v string)`

SetPurchaseOrderPrefix sets PurchaseOrderPrefix field to given value.

### HasPurchaseOrderPrefix

`func (o *ProcurementSetting) HasPurchaseOrderPrefix() bool`

HasPurchaseOrderPrefix returns a boolean if a field has been set.

### GetPurchaseOrderSuffix

`func (o *ProcurementSetting) GetPurchaseOrderSuffix() string`

GetPurchaseOrderSuffix returns the PurchaseOrderSuffix field if non-nil, zero value otherwise.

### GetPurchaseOrderSuffixOk

`func (o *ProcurementSetting) GetPurchaseOrderSuffixOk() (*string, bool)`

GetPurchaseOrderSuffixOk returns a tuple with the PurchaseOrderSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderSuffix

`func (o *ProcurementSetting) SetPurchaseOrderSuffix(v string)`

SetPurchaseOrderSuffix sets PurchaseOrderSuffix field to given value.

### HasPurchaseOrderSuffix

`func (o *ProcurementSetting) HasPurchaseOrderSuffix() bool`

HasPurchaseOrderSuffix returns a boolean if a field has been set.

### GetPrefixSuffixType

`func (o *ProcurementSetting) GetPrefixSuffixType() string`

GetPrefixSuffixType returns the PrefixSuffixType field if non-nil, zero value otherwise.

### GetPrefixSuffixTypeOk

`func (o *ProcurementSetting) GetPrefixSuffixTypeOk() (*string, bool)`

GetPrefixSuffixTypeOk returns a tuple with the PrefixSuffixType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixSuffixType

`func (o *ProcurementSetting) SetPrefixSuffixType(v string)`

SetPrefixSuffixType sets PrefixSuffixType field to given value.

### HasPrefixSuffixType

`func (o *ProcurementSetting) HasPrefixSuffixType() bool`

HasPrefixSuffixType returns a boolean if a field has been set.

### SetPrefixSuffixTypeNil

`func (o *ProcurementSetting) SetPrefixSuffixTypeNil(b bool)`

 SetPrefixSuffixTypeNil sets the value for PrefixSuffixType to be an explicit nil

### UnsetPrefixSuffixType
`func (o *ProcurementSetting) UnsetPrefixSuffixType()`

UnsetPrefixSuffixType ensures that no value is present for PrefixSuffixType, not even an explicit nil
### GetDisableCostUpdatesFlag

`func (o *ProcurementSetting) GetDisableCostUpdatesFlag() bool`

GetDisableCostUpdatesFlag returns the DisableCostUpdatesFlag field if non-nil, zero value otherwise.

### GetDisableCostUpdatesFlagOk

`func (o *ProcurementSetting) GetDisableCostUpdatesFlagOk() (*bool, bool)`

GetDisableCostUpdatesFlagOk returns a tuple with the DisableCostUpdatesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCostUpdatesFlag

`func (o *ProcurementSetting) SetDisableCostUpdatesFlag(v bool)`

SetDisableCostUpdatesFlag sets DisableCostUpdatesFlag field to given value.

### HasDisableCostUpdatesFlag

`func (o *ProcurementSetting) HasDisableCostUpdatesFlag() bool`

HasDisableCostUpdatesFlag returns a boolean if a field has been set.

### SetDisableCostUpdatesFlagNil

`func (o *ProcurementSetting) SetDisableCostUpdatesFlagNil(b bool)`

 SetDisableCostUpdatesFlagNil sets the value for DisableCostUpdatesFlag to be an explicit nil

### UnsetDisableCostUpdatesFlag
`func (o *ProcurementSetting) UnsetDisableCostUpdatesFlag()`

UnsetDisableCostUpdatesFlag ensures that no value is present for DisableCostUpdatesFlag, not even an explicit nil
### GetDisableNegativeInventoryFlag

`func (o *ProcurementSetting) GetDisableNegativeInventoryFlag() bool`

GetDisableNegativeInventoryFlag returns the DisableNegativeInventoryFlag field if non-nil, zero value otherwise.

### GetDisableNegativeInventoryFlagOk

`func (o *ProcurementSetting) GetDisableNegativeInventoryFlagOk() (*bool, bool)`

GetDisableNegativeInventoryFlagOk returns a tuple with the DisableNegativeInventoryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNegativeInventoryFlag

`func (o *ProcurementSetting) SetDisableNegativeInventoryFlag(v bool)`

SetDisableNegativeInventoryFlag sets DisableNegativeInventoryFlag field to given value.

### HasDisableNegativeInventoryFlag

`func (o *ProcurementSetting) HasDisableNegativeInventoryFlag() bool`

HasDisableNegativeInventoryFlag returns a boolean if a field has been set.

### SetDisableNegativeInventoryFlagNil

`func (o *ProcurementSetting) SetDisableNegativeInventoryFlagNil(b bool)`

 SetDisableNegativeInventoryFlagNil sets the value for DisableNegativeInventoryFlag to be an explicit nil

### UnsetDisableNegativeInventoryFlag
`func (o *ProcurementSetting) UnsetDisableNegativeInventoryFlag()`

UnsetDisableNegativeInventoryFlag ensures that no value is present for DisableNegativeInventoryFlag, not even an explicit nil
### GetCostingMethod

`func (o *ProcurementSetting) GetCostingMethod() string`

GetCostingMethod returns the CostingMethod field if non-nil, zero value otherwise.

### GetCostingMethodOk

`func (o *ProcurementSetting) GetCostingMethodOk() (*string, bool)`

GetCostingMethodOk returns a tuple with the CostingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostingMethod

`func (o *ProcurementSetting) SetCostingMethod(v string)`

SetCostingMethod sets CostingMethod field to given value.


### SetCostingMethodNil

`func (o *ProcurementSetting) SetCostingMethodNil(b bool)`

 SetCostingMethodNil sets the value for CostingMethod to be an explicit nil

### UnsetCostingMethod
`func (o *ProcurementSetting) UnsetCostingMethod()`

UnsetCostingMethod ensures that no value is present for CostingMethod, not even an explicit nil
### GetAutoClosePurchaseOrderFlag

`func (o *ProcurementSetting) GetAutoClosePurchaseOrderFlag() bool`

GetAutoClosePurchaseOrderFlag returns the AutoClosePurchaseOrderFlag field if non-nil, zero value otherwise.

### GetAutoClosePurchaseOrderFlagOk

`func (o *ProcurementSetting) GetAutoClosePurchaseOrderFlagOk() (*bool, bool)`

GetAutoClosePurchaseOrderFlagOk returns a tuple with the AutoClosePurchaseOrderFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoClosePurchaseOrderFlag

`func (o *ProcurementSetting) SetAutoClosePurchaseOrderFlag(v bool)`

SetAutoClosePurchaseOrderFlag sets AutoClosePurchaseOrderFlag field to given value.

### HasAutoClosePurchaseOrderFlag

`func (o *ProcurementSetting) HasAutoClosePurchaseOrderFlag() bool`

HasAutoClosePurchaseOrderFlag returns a boolean if a field has been set.

### SetAutoClosePurchaseOrderFlagNil

`func (o *ProcurementSetting) SetAutoClosePurchaseOrderFlagNil(b bool)`

 SetAutoClosePurchaseOrderFlagNil sets the value for AutoClosePurchaseOrderFlag to be an explicit nil

### UnsetAutoClosePurchaseOrderFlag
`func (o *ProcurementSetting) UnsetAutoClosePurchaseOrderFlag()`

UnsetAutoClosePurchaseOrderFlag ensures that no value is present for AutoClosePurchaseOrderFlag, not even an explicit nil
### GetAutoClosePurchaseOrderItemFlag

`func (o *ProcurementSetting) GetAutoClosePurchaseOrderItemFlag() bool`

GetAutoClosePurchaseOrderItemFlag returns the AutoClosePurchaseOrderItemFlag field if non-nil, zero value otherwise.

### GetAutoClosePurchaseOrderItemFlagOk

`func (o *ProcurementSetting) GetAutoClosePurchaseOrderItemFlagOk() (*bool, bool)`

GetAutoClosePurchaseOrderItemFlagOk returns a tuple with the AutoClosePurchaseOrderItemFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoClosePurchaseOrderItemFlag

`func (o *ProcurementSetting) SetAutoClosePurchaseOrderItemFlag(v bool)`

SetAutoClosePurchaseOrderItemFlag sets AutoClosePurchaseOrderItemFlag field to given value.

### HasAutoClosePurchaseOrderItemFlag

`func (o *ProcurementSetting) HasAutoClosePurchaseOrderItemFlag() bool`

HasAutoClosePurchaseOrderItemFlag returns a boolean if a field has been set.

### SetAutoClosePurchaseOrderItemFlagNil

`func (o *ProcurementSetting) SetAutoClosePurchaseOrderItemFlagNil(b bool)`

 SetAutoClosePurchaseOrderItemFlagNil sets the value for AutoClosePurchaseOrderItemFlag to be an explicit nil

### UnsetAutoClosePurchaseOrderItemFlag
`func (o *ProcurementSetting) UnsetAutoClosePurchaseOrderItemFlag()`

UnsetAutoClosePurchaseOrderItemFlag ensures that no value is present for AutoClosePurchaseOrderItemFlag, not even an explicit nil
### GetAutoApprovePurchaseOrderFlag

`func (o *ProcurementSetting) GetAutoApprovePurchaseOrderFlag() bool`

GetAutoApprovePurchaseOrderFlag returns the AutoApprovePurchaseOrderFlag field if non-nil, zero value otherwise.

### GetAutoApprovePurchaseOrderFlagOk

`func (o *ProcurementSetting) GetAutoApprovePurchaseOrderFlagOk() (*bool, bool)`

GetAutoApprovePurchaseOrderFlagOk returns a tuple with the AutoApprovePurchaseOrderFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoApprovePurchaseOrderFlag

`func (o *ProcurementSetting) SetAutoApprovePurchaseOrderFlag(v bool)`

SetAutoApprovePurchaseOrderFlag sets AutoApprovePurchaseOrderFlag field to given value.

### HasAutoApprovePurchaseOrderFlag

`func (o *ProcurementSetting) HasAutoApprovePurchaseOrderFlag() bool`

HasAutoApprovePurchaseOrderFlag returns a boolean if a field has been set.

### SetAutoApprovePurchaseOrderFlagNil

`func (o *ProcurementSetting) SetAutoApprovePurchaseOrderFlagNil(b bool)`

 SetAutoApprovePurchaseOrderFlagNil sets the value for AutoApprovePurchaseOrderFlag to be an explicit nil

### UnsetAutoApprovePurchaseOrderFlag
`func (o *ProcurementSetting) UnsetAutoApprovePurchaseOrderFlag()`

UnsetAutoApprovePurchaseOrderFlag ensures that no value is present for AutoApprovePurchaseOrderFlag, not even an explicit nil
### GetTaxPurchaseOrderFlag

`func (o *ProcurementSetting) GetTaxPurchaseOrderFlag() bool`

GetTaxPurchaseOrderFlag returns the TaxPurchaseOrderFlag field if non-nil, zero value otherwise.

### GetTaxPurchaseOrderFlagOk

`func (o *ProcurementSetting) GetTaxPurchaseOrderFlagOk() (*bool, bool)`

GetTaxPurchaseOrderFlagOk returns a tuple with the TaxPurchaseOrderFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPurchaseOrderFlag

`func (o *ProcurementSetting) SetTaxPurchaseOrderFlag(v bool)`

SetTaxPurchaseOrderFlag sets TaxPurchaseOrderFlag field to given value.

### HasTaxPurchaseOrderFlag

`func (o *ProcurementSetting) HasTaxPurchaseOrderFlag() bool`

HasTaxPurchaseOrderFlag returns a boolean if a field has been set.

### SetTaxPurchaseOrderFlagNil

`func (o *ProcurementSetting) SetTaxPurchaseOrderFlagNil(b bool)`

 SetTaxPurchaseOrderFlagNil sets the value for TaxPurchaseOrderFlag to be an explicit nil

### UnsetTaxPurchaseOrderFlag
`func (o *ProcurementSetting) UnsetTaxPurchaseOrderFlag()`

UnsetTaxPurchaseOrderFlag ensures that no value is present for TaxPurchaseOrderFlag, not even an explicit nil
### GetTaxFreightFlag

`func (o *ProcurementSetting) GetTaxFreightFlag() bool`

GetTaxFreightFlag returns the TaxFreightFlag field if non-nil, zero value otherwise.

### GetTaxFreightFlagOk

`func (o *ProcurementSetting) GetTaxFreightFlagOk() (*bool, bool)`

GetTaxFreightFlagOk returns a tuple with the TaxFreightFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxFreightFlag

`func (o *ProcurementSetting) SetTaxFreightFlag(v bool)`

SetTaxFreightFlag sets TaxFreightFlag field to given value.

### HasTaxFreightFlag

`func (o *ProcurementSetting) HasTaxFreightFlag() bool`

HasTaxFreightFlag returns a boolean if a field has been set.

### SetTaxFreightFlagNil

`func (o *ProcurementSetting) SetTaxFreightFlagNil(b bool)`

 SetTaxFreightFlagNil sets the value for TaxFreightFlag to be an explicit nil

### UnsetTaxFreightFlag
`func (o *ProcurementSetting) UnsetTaxFreightFlag()`

UnsetTaxFreightFlag ensures that no value is present for TaxFreightFlag, not even an explicit nil
### GetUseVendorTaxCodeFlag

`func (o *ProcurementSetting) GetUseVendorTaxCodeFlag() bool`

GetUseVendorTaxCodeFlag returns the UseVendorTaxCodeFlag field if non-nil, zero value otherwise.

### GetUseVendorTaxCodeFlagOk

`func (o *ProcurementSetting) GetUseVendorTaxCodeFlagOk() (*bool, bool)`

GetUseVendorTaxCodeFlagOk returns a tuple with the UseVendorTaxCodeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseVendorTaxCodeFlag

`func (o *ProcurementSetting) SetUseVendorTaxCodeFlag(v bool)`

SetUseVendorTaxCodeFlag sets UseVendorTaxCodeFlag field to given value.

### HasUseVendorTaxCodeFlag

`func (o *ProcurementSetting) HasUseVendorTaxCodeFlag() bool`

HasUseVendorTaxCodeFlag returns a boolean if a field has been set.

### SetUseVendorTaxCodeFlagNil

`func (o *ProcurementSetting) SetUseVendorTaxCodeFlagNil(b bool)`

 SetUseVendorTaxCodeFlagNil sets the value for UseVendorTaxCodeFlag to be an explicit nil

### UnsetUseVendorTaxCodeFlag
`func (o *ProcurementSetting) UnsetUseVendorTaxCodeFlag()`

UnsetUseVendorTaxCodeFlag ensures that no value is present for UseVendorTaxCodeFlag, not even an explicit nil
### GetNumDecimalPlaces

`func (o *ProcurementSetting) GetNumDecimalPlaces() int32`

GetNumDecimalPlaces returns the NumDecimalPlaces field if non-nil, zero value otherwise.

### GetNumDecimalPlacesOk

`func (o *ProcurementSetting) GetNumDecimalPlacesOk() (*int32, bool)`

GetNumDecimalPlacesOk returns a tuple with the NumDecimalPlaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDecimalPlaces

`func (o *ProcurementSetting) SetNumDecimalPlaces(v int32)`

SetNumDecimalPlaces sets NumDecimalPlaces field to given value.

### HasNumDecimalPlaces

`func (o *ProcurementSetting) HasNumDecimalPlaces() bool`

HasNumDecimalPlaces returns a boolean if a field has been set.

### SetNumDecimalPlacesNil

`func (o *ProcurementSetting) SetNumDecimalPlacesNil(b bool)`

 SetNumDecimalPlacesNil sets the value for NumDecimalPlaces to be an explicit nil

### UnsetNumDecimalPlaces
`func (o *ProcurementSetting) UnsetNumDecimalPlaces()`

UnsetNumDecimalPlaces ensures that no value is present for NumDecimalPlaces, not even an explicit nil
### GetDisableAutoPickFlag

`func (o *ProcurementSetting) GetDisableAutoPickFlag() bool`

GetDisableAutoPickFlag returns the DisableAutoPickFlag field if non-nil, zero value otherwise.

### GetDisableAutoPickFlagOk

`func (o *ProcurementSetting) GetDisableAutoPickFlagOk() (*bool, bool)`

GetDisableAutoPickFlagOk returns a tuple with the DisableAutoPickFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAutoPickFlag

`func (o *ProcurementSetting) SetDisableAutoPickFlag(v bool)`

SetDisableAutoPickFlag sets DisableAutoPickFlag field to given value.

### HasDisableAutoPickFlag

`func (o *ProcurementSetting) HasDisableAutoPickFlag() bool`

HasDisableAutoPickFlag returns a boolean if a field has been set.

### SetDisableAutoPickFlagNil

`func (o *ProcurementSetting) SetDisableAutoPickFlagNil(b bool)`

 SetDisableAutoPickFlagNil sets the value for DisableAutoPickFlag to be an explicit nil

### UnsetDisableAutoPickFlag
`func (o *ProcurementSetting) UnsetDisableAutoPickFlag()`

UnsetDisableAutoPickFlag ensures that no value is present for DisableAutoPickFlag, not even an explicit nil
### GetDefaultProductTaxableFlag

`func (o *ProcurementSetting) GetDefaultProductTaxableFlag() bool`

GetDefaultProductTaxableFlag returns the DefaultProductTaxableFlag field if non-nil, zero value otherwise.

### GetDefaultProductTaxableFlagOk

`func (o *ProcurementSetting) GetDefaultProductTaxableFlagOk() (*bool, bool)`

GetDefaultProductTaxableFlagOk returns a tuple with the DefaultProductTaxableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProductTaxableFlag

`func (o *ProcurementSetting) SetDefaultProductTaxableFlag(v bool)`

SetDefaultProductTaxableFlag sets DefaultProductTaxableFlag field to given value.

### HasDefaultProductTaxableFlag

`func (o *ProcurementSetting) HasDefaultProductTaxableFlag() bool`

HasDefaultProductTaxableFlag returns a boolean if a field has been set.

### SetDefaultProductTaxableFlagNil

`func (o *ProcurementSetting) SetDefaultProductTaxableFlagNil(b bool)`

 SetDefaultProductTaxableFlagNil sets the value for DefaultProductTaxableFlag to be an explicit nil

### UnsetDefaultProductTaxableFlag
`func (o *ProcurementSetting) UnsetDefaultProductTaxableFlag()`

UnsetDefaultProductTaxableFlag ensures that no value is present for DefaultProductTaxableFlag, not even an explicit nil
### GetEoriNumber

`func (o *ProcurementSetting) GetEoriNumber() string`

GetEoriNumber returns the EoriNumber field if non-nil, zero value otherwise.

### GetEoriNumberOk

`func (o *ProcurementSetting) GetEoriNumberOk() (*string, bool)`

GetEoriNumberOk returns a tuple with the EoriNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEoriNumber

`func (o *ProcurementSetting) SetEoriNumber(v string)`

SetEoriNumber sets EoriNumber field to given value.

### HasEoriNumber

`func (o *ProcurementSetting) HasEoriNumber() bool`

HasEoriNumber returns a boolean if a field has been set.

### GetInfo

`func (o *ProcurementSetting) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProcurementSetting) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProcurementSetting) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProcurementSetting) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetNotificationForChangesInShippingInfoFlag

`func (o *ProcurementSetting) GetNotificationForChangesInShippingInfoFlag() bool`

GetNotificationForChangesInShippingInfoFlag returns the NotificationForChangesInShippingInfoFlag field if non-nil, zero value otherwise.

### GetNotificationForChangesInShippingInfoFlagOk

`func (o *ProcurementSetting) GetNotificationForChangesInShippingInfoFlagOk() (*bool, bool)`

GetNotificationForChangesInShippingInfoFlagOk returns a tuple with the NotificationForChangesInShippingInfoFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationForChangesInShippingInfoFlag

`func (o *ProcurementSetting) SetNotificationForChangesInShippingInfoFlag(v bool)`

SetNotificationForChangesInShippingInfoFlag sets NotificationForChangesInShippingInfoFlag field to given value.

### HasNotificationForChangesInShippingInfoFlag

`func (o *ProcurementSetting) HasNotificationForChangesInShippingInfoFlag() bool`

HasNotificationForChangesInShippingInfoFlag returns a boolean if a field has been set.

### SetNotificationForChangesInShippingInfoFlagNil

`func (o *ProcurementSetting) SetNotificationForChangesInShippingInfoFlagNil(b bool)`

 SetNotificationForChangesInShippingInfoFlagNil sets the value for NotificationForChangesInShippingInfoFlag to be an explicit nil

### UnsetNotificationForChangesInShippingInfoFlag
`func (o *ProcurementSetting) UnsetNotificationForChangesInShippingInfoFlag()`

UnsetNotificationForChangesInShippingInfoFlag ensures that no value is present for NotificationForChangesInShippingInfoFlag, not even an explicit nil
### GetShippingInfoNotificationEmail

`func (o *ProcurementSetting) GetShippingInfoNotificationEmail() string`

GetShippingInfoNotificationEmail returns the ShippingInfoNotificationEmail field if non-nil, zero value otherwise.

### GetShippingInfoNotificationEmailOk

`func (o *ProcurementSetting) GetShippingInfoNotificationEmailOk() (*string, bool)`

GetShippingInfoNotificationEmailOk returns a tuple with the ShippingInfoNotificationEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingInfoNotificationEmail

`func (o *ProcurementSetting) SetShippingInfoNotificationEmail(v string)`

SetShippingInfoNotificationEmail sets ShippingInfoNotificationEmail field to given value.

### HasShippingInfoNotificationEmail

`func (o *ProcurementSetting) HasShippingInfoNotificationEmail() bool`

HasShippingInfoNotificationEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


