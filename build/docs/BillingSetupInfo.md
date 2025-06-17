# BillingSetupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**RemitName** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBillingSetupInfo

`func NewBillingSetupInfo() *BillingSetupInfo`

NewBillingSetupInfo instantiates a new BillingSetupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingSetupInfoWithDefaults

`func NewBillingSetupInfoWithDefaults() *BillingSetupInfo`

NewBillingSetupInfoWithDefaults instantiates a new BillingSetupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingSetupInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingSetupInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingSetupInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BillingSetupInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRemitName

`func (o *BillingSetupInfo) GetRemitName() string`

GetRemitName returns the RemitName field if non-nil, zero value otherwise.

### GetRemitNameOk

`func (o *BillingSetupInfo) GetRemitNameOk() (*string, bool)`

GetRemitNameOk returns a tuple with the RemitName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemitName

`func (o *BillingSetupInfo) SetRemitName(v string)`

SetRemitName sets RemitName field to given value.

### HasRemitName

`func (o *BillingSetupInfo) HasRemitName() bool`

HasRemitName returns a boolean if a field has been set.

### GetLocation

`func (o *BillingSetupInfo) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BillingSetupInfo) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BillingSetupInfo) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BillingSetupInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCurrency

`func (o *BillingSetupInfo) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingSetupInfo) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingSetupInfo) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingSetupInfo) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInfo

`func (o *BillingSetupInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BillingSetupInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BillingSetupInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BillingSetupInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


