# AllowedFileType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**FileType** | **string** |  | 
**SizeLimit** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAllowedFileType

`func NewAllowedFileType(fileType string, ) *AllowedFileType`

NewAllowedFileType instantiates a new AllowedFileType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowedFileTypeWithDefaults

`func NewAllowedFileTypeWithDefaults() *AllowedFileType`

NewAllowedFileTypeWithDefaults instantiates a new AllowedFileType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AllowedFileType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AllowedFileType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AllowedFileType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AllowedFileType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFileType

`func (o *AllowedFileType) GetFileType() string`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *AllowedFileType) GetFileTypeOk() (*string, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *AllowedFileType) SetFileType(v string)`

SetFileType sets FileType field to given value.


### GetSizeLimit

`func (o *AllowedFileType) GetSizeLimit() int32`

GetSizeLimit returns the SizeLimit field if non-nil, zero value otherwise.

### GetSizeLimitOk

`func (o *AllowedFileType) GetSizeLimitOk() (*int32, bool)`

GetSizeLimitOk returns a tuple with the SizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeLimit

`func (o *AllowedFileType) SetSizeLimit(v int32)`

SetSizeLimit sets SizeLimit field to given value.

### HasSizeLimit

`func (o *AllowedFileType) HasSizeLimit() bool`

HasSizeLimit returns a boolean if a field has been set.

### SetSizeLimitNil

`func (o *AllowedFileType) SetSizeLimitNil(b bool)`

 SetSizeLimitNil sets the value for SizeLimit to be an explicit nil

### UnsetSizeLimit
`func (o *AllowedFileType) UnsetSizeLimit()`

UnsetSizeLimit ensures that no value is present for SizeLimit, not even an explicit nil
### GetInfo

`func (o *AllowedFileType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AllowedFileType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AllowedFileType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AllowedFileType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


