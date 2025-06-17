# BillableOptionsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**OptionId** | Pointer to **string** |  | [optional] 
**BillableFlag** | Pointer to **bool** |  | [optional] 
**InvoiceFlag** | Pointer to **bool** |  | [optional] 
**TimeFlag** | Pointer to **bool** |  | [optional] 
**ExpenseFlag** | Pointer to **bool** |  | [optional] 
**ProductFlag** | Pointer to **bool** |  | [optional] 
**DefaultFlag** | Pointer to **bool** |  | [optional] 
**IncludeNoDefaultFlag** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**EnumId** | Pointer to **string** |  | [optional] 

## Methods

### NewBillableOptionsInfo

`func NewBillableOptionsInfo() *BillableOptionsInfo`

NewBillableOptionsInfo instantiates a new BillableOptionsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillableOptionsInfoWithDefaults

`func NewBillableOptionsInfoWithDefaults() *BillableOptionsInfo`

NewBillableOptionsInfoWithDefaults instantiates a new BillableOptionsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillableOptionsInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillableOptionsInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillableOptionsInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BillableOptionsInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOptionId

`func (o *BillableOptionsInfo) GetOptionId() string`

GetOptionId returns the OptionId field if non-nil, zero value otherwise.

### GetOptionIdOk

`func (o *BillableOptionsInfo) GetOptionIdOk() (*string, bool)`

GetOptionIdOk returns a tuple with the OptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionId

`func (o *BillableOptionsInfo) SetOptionId(v string)`

SetOptionId sets OptionId field to given value.

### HasOptionId

`func (o *BillableOptionsInfo) HasOptionId() bool`

HasOptionId returns a boolean if a field has been set.

### GetBillableFlag

`func (o *BillableOptionsInfo) GetBillableFlag() bool`

GetBillableFlag returns the BillableFlag field if non-nil, zero value otherwise.

### GetBillableFlagOk

`func (o *BillableOptionsInfo) GetBillableFlagOk() (*bool, bool)`

GetBillableFlagOk returns a tuple with the BillableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableFlag

`func (o *BillableOptionsInfo) SetBillableFlag(v bool)`

SetBillableFlag sets BillableFlag field to given value.

### HasBillableFlag

`func (o *BillableOptionsInfo) HasBillableFlag() bool`

HasBillableFlag returns a boolean if a field has been set.

### GetInvoiceFlag

`func (o *BillableOptionsInfo) GetInvoiceFlag() bool`

GetInvoiceFlag returns the InvoiceFlag field if non-nil, zero value otherwise.

### GetInvoiceFlagOk

`func (o *BillableOptionsInfo) GetInvoiceFlagOk() (*bool, bool)`

GetInvoiceFlagOk returns a tuple with the InvoiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceFlag

`func (o *BillableOptionsInfo) SetInvoiceFlag(v bool)`

SetInvoiceFlag sets InvoiceFlag field to given value.

### HasInvoiceFlag

`func (o *BillableOptionsInfo) HasInvoiceFlag() bool`

HasInvoiceFlag returns a boolean if a field has been set.

### GetTimeFlag

`func (o *BillableOptionsInfo) GetTimeFlag() bool`

GetTimeFlag returns the TimeFlag field if non-nil, zero value otherwise.

### GetTimeFlagOk

`func (o *BillableOptionsInfo) GetTimeFlagOk() (*bool, bool)`

GetTimeFlagOk returns a tuple with the TimeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFlag

`func (o *BillableOptionsInfo) SetTimeFlag(v bool)`

SetTimeFlag sets TimeFlag field to given value.

### HasTimeFlag

`func (o *BillableOptionsInfo) HasTimeFlag() bool`

HasTimeFlag returns a boolean if a field has been set.

### GetExpenseFlag

`func (o *BillableOptionsInfo) GetExpenseFlag() bool`

GetExpenseFlag returns the ExpenseFlag field if non-nil, zero value otherwise.

### GetExpenseFlagOk

`func (o *BillableOptionsInfo) GetExpenseFlagOk() (*bool, bool)`

GetExpenseFlagOk returns a tuple with the ExpenseFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseFlag

`func (o *BillableOptionsInfo) SetExpenseFlag(v bool)`

SetExpenseFlag sets ExpenseFlag field to given value.

### HasExpenseFlag

`func (o *BillableOptionsInfo) HasExpenseFlag() bool`

HasExpenseFlag returns a boolean if a field has been set.

### GetProductFlag

`func (o *BillableOptionsInfo) GetProductFlag() bool`

GetProductFlag returns the ProductFlag field if non-nil, zero value otherwise.

### GetProductFlagOk

`func (o *BillableOptionsInfo) GetProductFlagOk() (*bool, bool)`

GetProductFlagOk returns a tuple with the ProductFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductFlag

`func (o *BillableOptionsInfo) SetProductFlag(v bool)`

SetProductFlag sets ProductFlag field to given value.

### HasProductFlag

`func (o *BillableOptionsInfo) HasProductFlag() bool`

HasProductFlag returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *BillableOptionsInfo) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *BillableOptionsInfo) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *BillableOptionsInfo) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *BillableOptionsInfo) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### GetIncludeNoDefaultFlag

`func (o *BillableOptionsInfo) GetIncludeNoDefaultFlag() bool`

GetIncludeNoDefaultFlag returns the IncludeNoDefaultFlag field if non-nil, zero value otherwise.

### GetIncludeNoDefaultFlagOk

`func (o *BillableOptionsInfo) GetIncludeNoDefaultFlagOk() (*bool, bool)`

GetIncludeNoDefaultFlagOk returns a tuple with the IncludeNoDefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeNoDefaultFlag

`func (o *BillableOptionsInfo) SetIncludeNoDefaultFlag(v bool)`

SetIncludeNoDefaultFlag sets IncludeNoDefaultFlag field to given value.

### HasIncludeNoDefaultFlag

`func (o *BillableOptionsInfo) HasIncludeNoDefaultFlag() bool`

HasIncludeNoDefaultFlag returns a boolean if a field has been set.

### GetName

`func (o *BillableOptionsInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillableOptionsInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillableOptionsInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillableOptionsInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnumId

`func (o *BillableOptionsInfo) GetEnumId() string`

GetEnumId returns the EnumId field if non-nil, zero value otherwise.

### GetEnumIdOk

`func (o *BillableOptionsInfo) GetEnumIdOk() (*string, bool)`

GetEnumIdOk returns a tuple with the EnumId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumId

`func (o *BillableOptionsInfo) SetEnumId(v string)`

SetEnumId sets EnumId field to given value.

### HasEnumId

`func (o *BillableOptionsInfo) HasEnumId() bool`

HasEnumId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


