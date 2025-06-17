# BillingCycle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | **string** |  Max length: 5; | 
**Name** | **string** |  Max length: 50; | 
**DefaultFlag** | Pointer to **bool** |  | [optional] 
**BillingOptions** | **NullableString** |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBillingCycle

`func NewBillingCycle(identifier string, name string, billingOptions NullableString, ) *BillingCycle`

NewBillingCycle instantiates a new BillingCycle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingCycleWithDefaults

`func NewBillingCycleWithDefaults() *BillingCycle`

NewBillingCycleWithDefaults instantiates a new BillingCycle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingCycle) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingCycle) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingCycle) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BillingCycle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *BillingCycle) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *BillingCycle) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *BillingCycle) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetName

`func (o *BillingCycle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingCycle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingCycle) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultFlag

`func (o *BillingCycle) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *BillingCycle) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *BillingCycle) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *BillingCycle) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### GetBillingOptions

`func (o *BillingCycle) GetBillingOptions() string`

GetBillingOptions returns the BillingOptions field if non-nil, zero value otherwise.

### GetBillingOptionsOk

`func (o *BillingCycle) GetBillingOptionsOk() (*string, bool)`

GetBillingOptionsOk returns a tuple with the BillingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingOptions

`func (o *BillingCycle) SetBillingOptions(v string)`

SetBillingOptions sets BillingOptions field to given value.


### SetBillingOptionsNil

`func (o *BillingCycle) SetBillingOptionsNil(b bool)`

 SetBillingOptionsNil sets the value for BillingOptions to be an explicit nil

### UnsetBillingOptions
`func (o *BillingCycle) UnsetBillingOptions()`

UnsetBillingOptions ensures that no value is present for BillingOptions, not even an explicit nil
### GetInfo

`func (o *BillingCycle) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BillingCycle) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BillingCycle) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BillingCycle) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


