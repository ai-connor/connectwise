# CompanyTypeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**IsVendor** | Pointer to **bool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCompanyTypeInfo

`func NewCompanyTypeInfo() *CompanyTypeInfo`

NewCompanyTypeInfo instantiates a new CompanyTypeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyTypeInfoWithDefaults

`func NewCompanyTypeInfoWithDefaults() *CompanyTypeInfo`

NewCompanyTypeInfoWithDefaults instantiates a new CompanyTypeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyTypeInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyTypeInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyTypeInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyTypeInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CompanyTypeInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyTypeInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyTypeInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyTypeInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsVendor

`func (o *CompanyTypeInfo) GetIsVendor() bool`

GetIsVendor returns the IsVendor field if non-nil, zero value otherwise.

### GetIsVendorOk

`func (o *CompanyTypeInfo) GetIsVendorOk() (*bool, bool)`

GetIsVendorOk returns a tuple with the IsVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVendor

`func (o *CompanyTypeInfo) SetIsVendor(v bool)`

SetIsVendor sets IsVendor field to given value.

### HasIsVendor

`func (o *CompanyTypeInfo) HasIsVendor() bool`

HasIsVendor returns a boolean if a field has been set.

### GetInfo

`func (o *CompanyTypeInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CompanyTypeInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CompanyTypeInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CompanyTypeInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


