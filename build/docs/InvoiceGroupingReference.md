# InvoiceGroupingReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ShowPriceFlag** | Pointer to **bool** |  | [optional] 
**ShowSubItemsFlag** | Pointer to **bool** |  | [optional] 
**GroupParentChildAdditionsFlag** | Pointer to **bool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewInvoiceGroupingReference

`func NewInvoiceGroupingReference() *InvoiceGroupingReference`

NewInvoiceGroupingReference instantiates a new InvoiceGroupingReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceGroupingReferenceWithDefaults

`func NewInvoiceGroupingReferenceWithDefaults() *InvoiceGroupingReference`

NewInvoiceGroupingReferenceWithDefaults instantiates a new InvoiceGroupingReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceGroupingReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceGroupingReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceGroupingReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceGroupingReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *InvoiceGroupingReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InvoiceGroupingReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *InvoiceGroupingReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvoiceGroupingReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvoiceGroupingReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InvoiceGroupingReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InvoiceGroupingReference) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceGroupingReference) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceGroupingReference) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoiceGroupingReference) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetShowPriceFlag

`func (o *InvoiceGroupingReference) GetShowPriceFlag() bool`

GetShowPriceFlag returns the ShowPriceFlag field if non-nil, zero value otherwise.

### GetShowPriceFlagOk

`func (o *InvoiceGroupingReference) GetShowPriceFlagOk() (*bool, bool)`

GetShowPriceFlagOk returns a tuple with the ShowPriceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPriceFlag

`func (o *InvoiceGroupingReference) SetShowPriceFlag(v bool)`

SetShowPriceFlag sets ShowPriceFlag field to given value.

### HasShowPriceFlag

`func (o *InvoiceGroupingReference) HasShowPriceFlag() bool`

HasShowPriceFlag returns a boolean if a field has been set.

### GetShowSubItemsFlag

`func (o *InvoiceGroupingReference) GetShowSubItemsFlag() bool`

GetShowSubItemsFlag returns the ShowSubItemsFlag field if non-nil, zero value otherwise.

### GetShowSubItemsFlagOk

`func (o *InvoiceGroupingReference) GetShowSubItemsFlagOk() (*bool, bool)`

GetShowSubItemsFlagOk returns a tuple with the ShowSubItemsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowSubItemsFlag

`func (o *InvoiceGroupingReference) SetShowSubItemsFlag(v bool)`

SetShowSubItemsFlag sets ShowSubItemsFlag field to given value.

### HasShowSubItemsFlag

`func (o *InvoiceGroupingReference) HasShowSubItemsFlag() bool`

HasShowSubItemsFlag returns a boolean if a field has been set.

### GetGroupParentChildAdditionsFlag

`func (o *InvoiceGroupingReference) GetGroupParentChildAdditionsFlag() bool`

GetGroupParentChildAdditionsFlag returns the GroupParentChildAdditionsFlag field if non-nil, zero value otherwise.

### GetGroupParentChildAdditionsFlagOk

`func (o *InvoiceGroupingReference) GetGroupParentChildAdditionsFlagOk() (*bool, bool)`

GetGroupParentChildAdditionsFlagOk returns a tuple with the GroupParentChildAdditionsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupParentChildAdditionsFlag

`func (o *InvoiceGroupingReference) SetGroupParentChildAdditionsFlag(v bool)`

SetGroupParentChildAdditionsFlag sets GroupParentChildAdditionsFlag field to given value.

### HasGroupParentChildAdditionsFlag

`func (o *InvoiceGroupingReference) HasGroupParentChildAdditionsFlag() bool`

HasGroupParentChildAdditionsFlag returns a boolean if a field has been set.

### GetInfo

`func (o *InvoiceGroupingReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *InvoiceGroupingReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *InvoiceGroupingReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *InvoiceGroupingReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


