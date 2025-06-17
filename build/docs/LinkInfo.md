# LinkInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ScreenLink** | Pointer to **NullableString** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewLinkInfo

`func NewLinkInfo() *LinkInfo`

NewLinkInfo instantiates a new LinkInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkInfoWithDefaults

`func NewLinkInfoWithDefaults() *LinkInfo`

NewLinkInfoWithDefaults instantiates a new LinkInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinkInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LinkInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LinkInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LinkInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LinkInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LinkInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScreenLink

`func (o *LinkInfo) GetScreenLink() string`

GetScreenLink returns the ScreenLink field if non-nil, zero value otherwise.

### GetScreenLinkOk

`func (o *LinkInfo) GetScreenLinkOk() (*string, bool)`

GetScreenLinkOk returns a tuple with the ScreenLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenLink

`func (o *LinkInfo) SetScreenLink(v string)`

SetScreenLink sets ScreenLink field to given value.

### HasScreenLink

`func (o *LinkInfo) HasScreenLink() bool`

HasScreenLink returns a boolean if a field has been set.

### SetScreenLinkNil

`func (o *LinkInfo) SetScreenLinkNil(b bool)`

 SetScreenLinkNil sets the value for ScreenLink to be an explicit nil

### UnsetScreenLink
`func (o *LinkInfo) UnsetScreenLink()`

UnsetScreenLink ensures that no value is present for ScreenLink, not even an explicit nil
### GetInfo

`func (o *LinkInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *LinkInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *LinkInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *LinkInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


