# GLExportAdjustmentTransactionDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlClass** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Memo** | Pointer to **string** |  | [optional] 
**Item** | Pointer to [**IvItemReference**](IvItemReference.md) |  | [optional] 
**Quantity** | Pointer to **NullableInt32** |  | [optional] 
**Total** | Pointer to **NullableFloat64** |  | [optional] 
**Cost** | Pointer to **NullableFloat64** |  | [optional] 
**CostAccountNumber** | Pointer to **string** |  | [optional] 
**InventoryAccountNumber** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**ProductAccountNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewGLExportAdjustmentTransactionDetail

`func NewGLExportAdjustmentTransactionDetail() *GLExportAdjustmentTransactionDetail`

NewGLExportAdjustmentTransactionDetail instantiates a new GLExportAdjustmentTransactionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGLExportAdjustmentTransactionDetailWithDefaults

`func NewGLExportAdjustmentTransactionDetailWithDefaults() *GLExportAdjustmentTransactionDetail`

NewGLExportAdjustmentTransactionDetailWithDefaults instantiates a new GLExportAdjustmentTransactionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlClass

`func (o *GLExportAdjustmentTransactionDetail) GetGlClass() string`

GetGlClass returns the GlClass field if non-nil, zero value otherwise.

### GetGlClassOk

`func (o *GLExportAdjustmentTransactionDetail) GetGlClassOk() (*string, bool)`

GetGlClassOk returns a tuple with the GlClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlClass

`func (o *GLExportAdjustmentTransactionDetail) SetGlClass(v string)`

SetGlClass sets GlClass field to given value.

### HasGlClass

`func (o *GLExportAdjustmentTransactionDetail) HasGlClass() bool`

HasGlClass returns a boolean if a field has been set.

### GetDescription

`func (o *GLExportAdjustmentTransactionDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GLExportAdjustmentTransactionDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GLExportAdjustmentTransactionDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GLExportAdjustmentTransactionDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMemo

`func (o *GLExportAdjustmentTransactionDetail) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *GLExportAdjustmentTransactionDetail) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *GLExportAdjustmentTransactionDetail) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *GLExportAdjustmentTransactionDetail) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### GetItem

`func (o *GLExportAdjustmentTransactionDetail) GetItem() IvItemReference`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *GLExportAdjustmentTransactionDetail) GetItemOk() (*IvItemReference, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *GLExportAdjustmentTransactionDetail) SetItem(v IvItemReference)`

SetItem sets Item field to given value.

### HasItem

`func (o *GLExportAdjustmentTransactionDetail) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetQuantity

`func (o *GLExportAdjustmentTransactionDetail) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GLExportAdjustmentTransactionDetail) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GLExportAdjustmentTransactionDetail) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *GLExportAdjustmentTransactionDetail) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *GLExportAdjustmentTransactionDetail) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *GLExportAdjustmentTransactionDetail) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetTotal

`func (o *GLExportAdjustmentTransactionDetail) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GLExportAdjustmentTransactionDetail) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GLExportAdjustmentTransactionDetail) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GLExportAdjustmentTransactionDetail) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *GLExportAdjustmentTransactionDetail) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *GLExportAdjustmentTransactionDetail) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetCost

`func (o *GLExportAdjustmentTransactionDetail) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GLExportAdjustmentTransactionDetail) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GLExportAdjustmentTransactionDetail) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GLExportAdjustmentTransactionDetail) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *GLExportAdjustmentTransactionDetail) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *GLExportAdjustmentTransactionDetail) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCostAccountNumber

`func (o *GLExportAdjustmentTransactionDetail) GetCostAccountNumber() string`

GetCostAccountNumber returns the CostAccountNumber field if non-nil, zero value otherwise.

### GetCostAccountNumberOk

`func (o *GLExportAdjustmentTransactionDetail) GetCostAccountNumberOk() (*string, bool)`

GetCostAccountNumberOk returns a tuple with the CostAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostAccountNumber

`func (o *GLExportAdjustmentTransactionDetail) SetCostAccountNumber(v string)`

SetCostAccountNumber sets CostAccountNumber field to given value.

### HasCostAccountNumber

`func (o *GLExportAdjustmentTransactionDetail) HasCostAccountNumber() bool`

HasCostAccountNumber returns a boolean if a field has been set.

### GetInventoryAccountNumber

`func (o *GLExportAdjustmentTransactionDetail) GetInventoryAccountNumber() string`

GetInventoryAccountNumber returns the InventoryAccountNumber field if non-nil, zero value otherwise.

### GetInventoryAccountNumberOk

`func (o *GLExportAdjustmentTransactionDetail) GetInventoryAccountNumberOk() (*string, bool)`

GetInventoryAccountNumberOk returns a tuple with the InventoryAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryAccountNumber

`func (o *GLExportAdjustmentTransactionDetail) SetInventoryAccountNumber(v string)`

SetInventoryAccountNumber sets InventoryAccountNumber field to given value.

### HasInventoryAccountNumber

`func (o *GLExportAdjustmentTransactionDetail) HasInventoryAccountNumber() bool`

HasInventoryAccountNumber returns a boolean if a field has been set.

### GetAccountNumber

`func (o *GLExportAdjustmentTransactionDetail) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *GLExportAdjustmentTransactionDetail) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *GLExportAdjustmentTransactionDetail) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *GLExportAdjustmentTransactionDetail) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetProductAccountNumber

`func (o *GLExportAdjustmentTransactionDetail) GetProductAccountNumber() string`

GetProductAccountNumber returns the ProductAccountNumber field if non-nil, zero value otherwise.

### GetProductAccountNumberOk

`func (o *GLExportAdjustmentTransactionDetail) GetProductAccountNumberOk() (*string, bool)`

GetProductAccountNumberOk returns a tuple with the ProductAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductAccountNumber

`func (o *GLExportAdjustmentTransactionDetail) SetProductAccountNumber(v string)`

SetProductAccountNumber sets ProductAccountNumber field to given value.

### HasProductAccountNumber

`func (o *GLExportAdjustmentTransactionDetail) HasProductAccountNumber() bool`

HasProductAccountNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


