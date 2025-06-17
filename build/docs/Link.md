# Link

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**TableReferenceId** | Pointer to **NullableInt32** |  | [optional] 
**Url** | Pointer to **string** |  Max length: 1000; | [optional] 
**ScreenLink** | Pointer to **NullableString** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewLink

`func NewLink(name string, ) *Link`

NewLink instantiates a new Link object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkWithDefaults

`func NewLinkWithDefaults() *Link`

NewLinkWithDefaults instantiates a new Link object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Link) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Link) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Link) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Link) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Link) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Link) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Link) SetName(v string)`

SetName sets Name field to given value.


### GetTableReferenceId

`func (o *Link) GetTableReferenceId() int32`

GetTableReferenceId returns the TableReferenceId field if non-nil, zero value otherwise.

### GetTableReferenceIdOk

`func (o *Link) GetTableReferenceIdOk() (*int32, bool)`

GetTableReferenceIdOk returns a tuple with the TableReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableReferenceId

`func (o *Link) SetTableReferenceId(v int32)`

SetTableReferenceId sets TableReferenceId field to given value.

### HasTableReferenceId

`func (o *Link) HasTableReferenceId() bool`

HasTableReferenceId returns a boolean if a field has been set.

### SetTableReferenceIdNil

`func (o *Link) SetTableReferenceIdNil(b bool)`

 SetTableReferenceIdNil sets the value for TableReferenceId to be an explicit nil

### UnsetTableReferenceId
`func (o *Link) UnsetTableReferenceId()`

UnsetTableReferenceId ensures that no value is present for TableReferenceId, not even an explicit nil
### GetUrl

`func (o *Link) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Link) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Link) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Link) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetScreenLink

`func (o *Link) GetScreenLink() string`

GetScreenLink returns the ScreenLink field if non-nil, zero value otherwise.

### GetScreenLinkOk

`func (o *Link) GetScreenLinkOk() (*string, bool)`

GetScreenLinkOk returns a tuple with the ScreenLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenLink

`func (o *Link) SetScreenLink(v string)`

SetScreenLink sets ScreenLink field to given value.

### HasScreenLink

`func (o *Link) HasScreenLink() bool`

HasScreenLink returns a boolean if a field has been set.

### SetScreenLinkNil

`func (o *Link) SetScreenLinkNil(b bool)`

 SetScreenLinkNil sets the value for ScreenLink to be an explicit nil

### UnsetScreenLink
`func (o *Link) UnsetScreenLink()`

UnsetScreenLink ensures that no value is present for ScreenLink, not even an explicit nil
### GetInfo

`func (o *Link) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Link) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Link) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Link) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


