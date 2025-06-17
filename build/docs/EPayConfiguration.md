# EPayConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Location** | [**SystemLocationReference**](SystemLocationReference.md) |  | 
**Currency** | [**CurrencyReference**](CurrencyReference.md) |  | 
**Url** | **string** |  Max length: 400; | 
**StoreIdentifier** | **string** |  Max length: 500; | 
**EncryptionKey** | Pointer to **string** |  Max length: 500; | [optional] 
**InitializationVector** | Pointer to **string** |  Max length: 500; | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewEPayConfiguration

`func NewEPayConfiguration(location SystemLocationReference, currency CurrencyReference, url string, storeIdentifier string, ) *EPayConfiguration`

NewEPayConfiguration instantiates a new EPayConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEPayConfigurationWithDefaults

`func NewEPayConfigurationWithDefaults() *EPayConfiguration`

NewEPayConfigurationWithDefaults instantiates a new EPayConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EPayConfiguration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EPayConfiguration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EPayConfiguration) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EPayConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *EPayConfiguration) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *EPayConfiguration) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *EPayConfiguration) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.


### GetCurrency

`func (o *EPayConfiguration) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *EPayConfiguration) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *EPayConfiguration) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.


### GetUrl

`func (o *EPayConfiguration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EPayConfiguration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EPayConfiguration) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetStoreIdentifier

`func (o *EPayConfiguration) GetStoreIdentifier() string`

GetStoreIdentifier returns the StoreIdentifier field if non-nil, zero value otherwise.

### GetStoreIdentifierOk

`func (o *EPayConfiguration) GetStoreIdentifierOk() (*string, bool)`

GetStoreIdentifierOk returns a tuple with the StoreIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreIdentifier

`func (o *EPayConfiguration) SetStoreIdentifier(v string)`

SetStoreIdentifier sets StoreIdentifier field to given value.


### GetEncryptionKey

`func (o *EPayConfiguration) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *EPayConfiguration) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *EPayConfiguration) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *EPayConfiguration) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetInitializationVector

`func (o *EPayConfiguration) GetInitializationVector() string`

GetInitializationVector returns the InitializationVector field if non-nil, zero value otherwise.

### GetInitializationVectorOk

`func (o *EPayConfiguration) GetInitializationVectorOk() (*string, bool)`

GetInitializationVectorOk returns a tuple with the InitializationVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitializationVector

`func (o *EPayConfiguration) SetInitializationVector(v string)`

SetInitializationVector sets InitializationVector field to given value.

### HasInitializationVector

`func (o *EPayConfiguration) HasInitializationVector() bool`

HasInitializationVector returns a boolean if a field has been set.

### GetInfo

`func (o *EPayConfiguration) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *EPayConfiguration) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *EPayConfiguration) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *EPayConfiguration) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


