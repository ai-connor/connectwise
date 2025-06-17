# TaxIntegrationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**EnabledFlag** | Pointer to **bool** |  | [optional] 
**TaxIntegrationType** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTaxIntegrationInfo

`func NewTaxIntegrationInfo() *TaxIntegrationInfo`

NewTaxIntegrationInfo instantiates a new TaxIntegrationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxIntegrationInfoWithDefaults

`func NewTaxIntegrationInfoWithDefaults() *TaxIntegrationInfo`

NewTaxIntegrationInfoWithDefaults instantiates a new TaxIntegrationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxIntegrationInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxIntegrationInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxIntegrationInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TaxIntegrationInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnabledFlag

`func (o *TaxIntegrationInfo) GetEnabledFlag() bool`

GetEnabledFlag returns the EnabledFlag field if non-nil, zero value otherwise.

### GetEnabledFlagOk

`func (o *TaxIntegrationInfo) GetEnabledFlagOk() (*bool, bool)`

GetEnabledFlagOk returns a tuple with the EnabledFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledFlag

`func (o *TaxIntegrationInfo) SetEnabledFlag(v bool)`

SetEnabledFlag sets EnabledFlag field to given value.

### HasEnabledFlag

`func (o *TaxIntegrationInfo) HasEnabledFlag() bool`

HasEnabledFlag returns a boolean if a field has been set.

### GetTaxIntegrationType

`func (o *TaxIntegrationInfo) GetTaxIntegrationType() string`

GetTaxIntegrationType returns the TaxIntegrationType field if non-nil, zero value otherwise.

### GetTaxIntegrationTypeOk

`func (o *TaxIntegrationInfo) GetTaxIntegrationTypeOk() (*string, bool)`

GetTaxIntegrationTypeOk returns a tuple with the TaxIntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxIntegrationType

`func (o *TaxIntegrationInfo) SetTaxIntegrationType(v string)`

SetTaxIntegrationType sets TaxIntegrationType field to given value.

### HasTaxIntegrationType

`func (o *TaxIntegrationInfo) HasTaxIntegrationType() bool`

HasTaxIntegrationType returns a boolean if a field has been set.

### GetInfo

`func (o *TaxIntegrationInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TaxIntegrationInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TaxIntegrationInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TaxIntegrationInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


