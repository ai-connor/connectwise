# CorporateStructureInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**LocationCaption** | Pointer to **string** |  | [optional] 
**GroupCaption** | Pointer to **string** |  | [optional] 
**BaseCurrency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCorporateStructureInfo

`func NewCorporateStructureInfo() *CorporateStructureInfo`

NewCorporateStructureInfo instantiates a new CorporateStructureInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporateStructureInfoWithDefaults

`func NewCorporateStructureInfoWithDefaults() *CorporateStructureInfo`

NewCorporateStructureInfoWithDefaults instantiates a new CorporateStructureInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CorporateStructureInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CorporateStructureInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CorporateStructureInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CorporateStructureInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocationCaption

`func (o *CorporateStructureInfo) GetLocationCaption() string`

GetLocationCaption returns the LocationCaption field if non-nil, zero value otherwise.

### GetLocationCaptionOk

`func (o *CorporateStructureInfo) GetLocationCaptionOk() (*string, bool)`

GetLocationCaptionOk returns a tuple with the LocationCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationCaption

`func (o *CorporateStructureInfo) SetLocationCaption(v string)`

SetLocationCaption sets LocationCaption field to given value.

### HasLocationCaption

`func (o *CorporateStructureInfo) HasLocationCaption() bool`

HasLocationCaption returns a boolean if a field has been set.

### GetGroupCaption

`func (o *CorporateStructureInfo) GetGroupCaption() string`

GetGroupCaption returns the GroupCaption field if non-nil, zero value otherwise.

### GetGroupCaptionOk

`func (o *CorporateStructureInfo) GetGroupCaptionOk() (*string, bool)`

GetGroupCaptionOk returns a tuple with the GroupCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCaption

`func (o *CorporateStructureInfo) SetGroupCaption(v string)`

SetGroupCaption sets GroupCaption field to given value.

### HasGroupCaption

`func (o *CorporateStructureInfo) HasGroupCaption() bool`

HasGroupCaption returns a boolean if a field has been set.

### GetBaseCurrency

`func (o *CorporateStructureInfo) GetBaseCurrency() CurrencyReference`

GetBaseCurrency returns the BaseCurrency field if non-nil, zero value otherwise.

### GetBaseCurrencyOk

`func (o *CorporateStructureInfo) GetBaseCurrencyOk() (*CurrencyReference, bool)`

GetBaseCurrencyOk returns a tuple with the BaseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrency

`func (o *CorporateStructureInfo) SetBaseCurrency(v CurrencyReference)`

SetBaseCurrency sets BaseCurrency field to given value.

### HasBaseCurrency

`func (o *CorporateStructureInfo) HasBaseCurrency() bool`

HasBaseCurrency returns a boolean if a field has been set.

### GetInfo

`func (o *CorporateStructureInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CorporateStructureInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CorporateStructureInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CorporateStructureInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


