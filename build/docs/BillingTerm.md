# BillingTerm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**DueDays** | **NullableInt32** |  | 
**TermsXref** | Pointer to **string** |  Max length: 50; | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBillingTerm

`func NewBillingTerm(name string, dueDays NullableInt32, ) *BillingTerm`

NewBillingTerm instantiates a new BillingTerm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingTermWithDefaults

`func NewBillingTermWithDefaults() *BillingTerm`

NewBillingTermWithDefaults instantiates a new BillingTerm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingTerm) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingTerm) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingTerm) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BillingTerm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BillingTerm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingTerm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingTerm) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultFlag

`func (o *BillingTerm) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *BillingTerm) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *BillingTerm) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *BillingTerm) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *BillingTerm) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *BillingTerm) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetDueDays

`func (o *BillingTerm) GetDueDays() int32`

GetDueDays returns the DueDays field if non-nil, zero value otherwise.

### GetDueDaysOk

`func (o *BillingTerm) GetDueDaysOk() (*int32, bool)`

GetDueDaysOk returns a tuple with the DueDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDays

`func (o *BillingTerm) SetDueDays(v int32)`

SetDueDays sets DueDays field to given value.


### SetDueDaysNil

`func (o *BillingTerm) SetDueDaysNil(b bool)`

 SetDueDaysNil sets the value for DueDays to be an explicit nil

### UnsetDueDays
`func (o *BillingTerm) UnsetDueDays()`

UnsetDueDays ensures that no value is present for DueDays, not even an explicit nil
### GetTermsXref

`func (o *BillingTerm) GetTermsXref() string`

GetTermsXref returns the TermsXref field if non-nil, zero value otherwise.

### GetTermsXrefOk

`func (o *BillingTerm) GetTermsXrefOk() (*string, bool)`

GetTermsXrefOk returns a tuple with the TermsXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsXref

`func (o *BillingTerm) SetTermsXref(v string)`

SetTermsXref sets TermsXref field to given value.

### HasTermsXref

`func (o *BillingTerm) HasTermsXref() bool`

HasTermsXref returns a boolean if a field has been set.

### GetInfo

`func (o *BillingTerm) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BillingTerm) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BillingTerm) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BillingTerm) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


