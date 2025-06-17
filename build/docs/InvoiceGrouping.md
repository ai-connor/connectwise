# InvoiceGrouping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**CustomerDescription** | **string** |  | 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**ShowPriceFlag** | Pointer to **NullableBool** |  | [optional] 
**ShowSubItemsFlag** | Pointer to **NullableBool** |  | [optional] 
**GroupParentChildAdditionsFlag** | Pointer to **bool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewInvoiceGrouping

`func NewInvoiceGrouping(name string, customerDescription string, ) *InvoiceGrouping`

NewInvoiceGrouping instantiates a new InvoiceGrouping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceGroupingWithDefaults

`func NewInvoiceGroupingWithDefaults() *InvoiceGrouping`

NewInvoiceGroupingWithDefaults instantiates a new InvoiceGrouping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceGrouping) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceGrouping) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceGrouping) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceGrouping) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InvoiceGrouping) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvoiceGrouping) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvoiceGrouping) SetName(v string)`

SetName sets Name field to given value.


### GetCustomerDescription

`func (o *InvoiceGrouping) GetCustomerDescription() string`

GetCustomerDescription returns the CustomerDescription field if non-nil, zero value otherwise.

### GetCustomerDescriptionOk

`func (o *InvoiceGrouping) GetCustomerDescriptionOk() (*string, bool)`

GetCustomerDescriptionOk returns a tuple with the CustomerDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDescription

`func (o *InvoiceGrouping) SetCustomerDescription(v string)`

SetCustomerDescription sets CustomerDescription field to given value.


### GetInactiveFlag

`func (o *InvoiceGrouping) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *InvoiceGrouping) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *InvoiceGrouping) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *InvoiceGrouping) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *InvoiceGrouping) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *InvoiceGrouping) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetShowPriceFlag

`func (o *InvoiceGrouping) GetShowPriceFlag() bool`

GetShowPriceFlag returns the ShowPriceFlag field if non-nil, zero value otherwise.

### GetShowPriceFlagOk

`func (o *InvoiceGrouping) GetShowPriceFlagOk() (*bool, bool)`

GetShowPriceFlagOk returns a tuple with the ShowPriceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPriceFlag

`func (o *InvoiceGrouping) SetShowPriceFlag(v bool)`

SetShowPriceFlag sets ShowPriceFlag field to given value.

### HasShowPriceFlag

`func (o *InvoiceGrouping) HasShowPriceFlag() bool`

HasShowPriceFlag returns a boolean if a field has been set.

### SetShowPriceFlagNil

`func (o *InvoiceGrouping) SetShowPriceFlagNil(b bool)`

 SetShowPriceFlagNil sets the value for ShowPriceFlag to be an explicit nil

### UnsetShowPriceFlag
`func (o *InvoiceGrouping) UnsetShowPriceFlag()`

UnsetShowPriceFlag ensures that no value is present for ShowPriceFlag, not even an explicit nil
### GetShowSubItemsFlag

`func (o *InvoiceGrouping) GetShowSubItemsFlag() bool`

GetShowSubItemsFlag returns the ShowSubItemsFlag field if non-nil, zero value otherwise.

### GetShowSubItemsFlagOk

`func (o *InvoiceGrouping) GetShowSubItemsFlagOk() (*bool, bool)`

GetShowSubItemsFlagOk returns a tuple with the ShowSubItemsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSubItemsFlag

`func (o *InvoiceGrouping) SetShowSubItemsFlag(v bool)`

SetShowSubItemsFlag sets ShowSubItemsFlag field to given value.

### HasShowSubItemsFlag

`func (o *InvoiceGrouping) HasShowSubItemsFlag() bool`

HasShowSubItemsFlag returns a boolean if a field has been set.

### SetShowSubItemsFlagNil

`func (o *InvoiceGrouping) SetShowSubItemsFlagNil(b bool)`

 SetShowSubItemsFlagNil sets the value for ShowSubItemsFlag to be an explicit nil

### UnsetShowSubItemsFlag
`func (o *InvoiceGrouping) UnsetShowSubItemsFlag()`

UnsetShowSubItemsFlag ensures that no value is present for ShowSubItemsFlag, not even an explicit nil
### GetGroupParentChildAdditionsFlag

`func (o *InvoiceGrouping) GetGroupParentChildAdditionsFlag() bool`

GetGroupParentChildAdditionsFlag returns the GroupParentChildAdditionsFlag field if non-nil, zero value otherwise.

### GetGroupParentChildAdditionsFlagOk

`func (o *InvoiceGrouping) GetGroupParentChildAdditionsFlagOk() (*bool, bool)`

GetGroupParentChildAdditionsFlagOk returns a tuple with the GroupParentChildAdditionsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupParentChildAdditionsFlag

`func (o *InvoiceGrouping) SetGroupParentChildAdditionsFlag(v bool)`

SetGroupParentChildAdditionsFlag sets GroupParentChildAdditionsFlag field to given value.

### HasGroupParentChildAdditionsFlag

`func (o *InvoiceGrouping) HasGroupParentChildAdditionsFlag() bool`

HasGroupParentChildAdditionsFlag returns a boolean if a field has been set.

### GetInfo

`func (o *InvoiceGrouping) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *InvoiceGrouping) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *InvoiceGrouping) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *InvoiceGrouping) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


