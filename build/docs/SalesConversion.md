# SalesConversion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentType** | Pointer to **string** |  | [optional] 
**ConvertedTo** | Pointer to [**ConversionTypeReference**](ConversionTypeReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSalesConversion

`func NewSalesConversion() *SalesConversion`

NewSalesConversion instantiates a new SalesConversion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesConversionWithDefaults

`func NewSalesConversionWithDefaults() *SalesConversion`

NewSalesConversionWithDefaults instantiates a new SalesConversion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentType

`func (o *SalesConversion) GetParentType() string`

GetParentType returns the ParentType field if non-nil, zero value otherwise.

### GetParentTypeOk

`func (o *SalesConversion) GetParentTypeOk() (*string, bool)`

GetParentTypeOk returns a tuple with the ParentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentType

`func (o *SalesConversion) SetParentType(v string)`

SetParentType sets ParentType field to given value.

### HasParentType

`func (o *SalesConversion) HasParentType() bool`

HasParentType returns a boolean if a field has been set.

### GetConvertedTo

`func (o *SalesConversion) GetConvertedTo() ConversionTypeReference`

GetConvertedTo returns the ConvertedTo field if non-nil, zero value otherwise.

### GetConvertedToOk

`func (o *SalesConversion) GetConvertedToOk() (*ConversionTypeReference, bool)`

GetConvertedToOk returns a tuple with the ConvertedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConvertedTo

`func (o *SalesConversion) SetConvertedTo(v ConversionTypeReference)`

SetConvertedTo sets ConvertedTo field to given value.

### HasConvertedTo

`func (o *SalesConversion) HasConvertedTo() bool`

HasConvertedTo returns a boolean if a field has been set.

### GetInfo

`func (o *SalesConversion) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SalesConversion) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SalesConversion) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SalesConversion) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


