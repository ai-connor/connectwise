# InvoiceReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**BillingType** | Pointer to **string** |  | [optional] 
**ApplyToType** | Pointer to **string** |  | [optional] 
**InvoiceDate** | Pointer to **string** |  | [optional] 
**ChargeFirmFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewInvoiceReference

`func NewInvoiceReference() *InvoiceReference`

NewInvoiceReference instantiates a new InvoiceReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceReferenceWithDefaults

`func NewInvoiceReferenceWithDefaults() *InvoiceReference`

NewInvoiceReferenceWithDefaults instantiates a new InvoiceReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *InvoiceReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *InvoiceReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIdentifier

`func (o *InvoiceReference) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *InvoiceReference) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *InvoiceReference) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *InvoiceReference) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetBillingType

`func (o *InvoiceReference) GetBillingType() string`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *InvoiceReference) GetBillingTypeOk() (*string, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *InvoiceReference) SetBillingType(v string)`

SetBillingType sets BillingType field to given value.

### HasBillingType

`func (o *InvoiceReference) HasBillingType() bool`

HasBillingType returns a boolean if a field has been set.

### GetApplyToType

`func (o *InvoiceReference) GetApplyToType() string`

GetApplyToType returns the ApplyToType field if non-nil, zero value otherwise.

### GetApplyToTypeOk

`func (o *InvoiceReference) GetApplyToTypeOk() (*string, bool)`

GetApplyToTypeOk returns a tuple with the ApplyToType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToType

`func (o *InvoiceReference) SetApplyToType(v string)`

SetApplyToType sets ApplyToType field to given value.

### HasApplyToType

`func (o *InvoiceReference) HasApplyToType() bool`

HasApplyToType returns a boolean if a field has been set.

### GetInvoiceDate

`func (o *InvoiceReference) GetInvoiceDate() string`

GetInvoiceDate returns the InvoiceDate field if non-nil, zero value otherwise.

### GetInvoiceDateOk

`func (o *InvoiceReference) GetInvoiceDateOk() (*string, bool)`

GetInvoiceDateOk returns a tuple with the InvoiceDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDate

`func (o *InvoiceReference) SetInvoiceDate(v string)`

SetInvoiceDate sets InvoiceDate field to given value.

### HasInvoiceDate

`func (o *InvoiceReference) HasInvoiceDate() bool`

HasInvoiceDate returns a boolean if a field has been set.

### GetChargeFirmFlag

`func (o *InvoiceReference) GetChargeFirmFlag() bool`

GetChargeFirmFlag returns the ChargeFirmFlag field if non-nil, zero value otherwise.

### GetChargeFirmFlagOk

`func (o *InvoiceReference) GetChargeFirmFlagOk() (*bool, bool)`

GetChargeFirmFlagOk returns a tuple with the ChargeFirmFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeFirmFlag

`func (o *InvoiceReference) SetChargeFirmFlag(v bool)`

SetChargeFirmFlag sets ChargeFirmFlag field to given value.

### HasChargeFirmFlag

`func (o *InvoiceReference) HasChargeFirmFlag() bool`

HasChargeFirmFlag returns a boolean if a field has been set.

### SetChargeFirmFlagNil

`func (o *InvoiceReference) SetChargeFirmFlagNil(b bool)`

 SetChargeFirmFlagNil sets the value for ChargeFirmFlag to be an explicit nil

### UnsetChargeFirmFlag
`func (o *InvoiceReference) UnsetChargeFirmFlag()`

UnsetChargeFirmFlag ensures that no value is present for ChargeFirmFlag, not even an explicit nil
### GetInfo

`func (o *InvoiceReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *InvoiceReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *InvoiceReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *InvoiceReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


