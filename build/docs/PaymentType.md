# PaymentType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 30; | 
**Classification** | [**ClassificationReference**](ClassificationReference.md) |  | 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**CompanyFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPaymentType

`func NewPaymentType(name string, classification ClassificationReference, ) *PaymentType`

NewPaymentType instantiates a new PaymentType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentTypeWithDefaults

`func NewPaymentTypeWithDefaults() *PaymentType`

NewPaymentTypeWithDefaults instantiates a new PaymentType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PaymentType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentType) SetName(v string)`

SetName sets Name field to given value.


### GetClassification

`func (o *PaymentType) GetClassification() ClassificationReference`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *PaymentType) GetClassificationOk() (*ClassificationReference, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *PaymentType) SetClassification(v ClassificationReference)`

SetClassification sets Classification field to given value.


### GetDefaultFlag

`func (o *PaymentType) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *PaymentType) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *PaymentType) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *PaymentType) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *PaymentType) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *PaymentType) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetCompanyFlag

`func (o *PaymentType) GetCompanyFlag() bool`

GetCompanyFlag returns the CompanyFlag field if non-nil, zero value otherwise.

### GetCompanyFlagOk

`func (o *PaymentType) GetCompanyFlagOk() (*bool, bool)`

GetCompanyFlagOk returns a tuple with the CompanyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyFlag

`func (o *PaymentType) SetCompanyFlag(v bool)`

SetCompanyFlag sets CompanyFlag field to given value.

### HasCompanyFlag

`func (o *PaymentType) HasCompanyFlag() bool`

HasCompanyFlag returns a boolean if a field has been set.

### SetCompanyFlagNil

`func (o *PaymentType) SetCompanyFlagNil(b bool)`

 SetCompanyFlagNil sets the value for CompanyFlag to be an explicit nil

### UnsetCompanyFlag
`func (o *PaymentType) UnsetCompanyFlag()`

UnsetCompanyFlag ensures that no value is present for CompanyFlag, not even an explicit nil
### GetInfo

`func (o *PaymentType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PaymentType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PaymentType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PaymentType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


