# LdapConfigurationReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Server** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewLdapConfigurationReference

`func NewLdapConfigurationReference() *LdapConfigurationReference`

NewLdapConfigurationReference instantiates a new LdapConfigurationReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapConfigurationReferenceWithDefaults

`func NewLdapConfigurationReferenceWithDefaults() *LdapConfigurationReference`

NewLdapConfigurationReferenceWithDefaults instantiates a new LdapConfigurationReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LdapConfigurationReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LdapConfigurationReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LdapConfigurationReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LdapConfigurationReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *LdapConfigurationReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LdapConfigurationReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *LdapConfigurationReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LdapConfigurationReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LdapConfigurationReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LdapConfigurationReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServer

`func (o *LdapConfigurationReference) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *LdapConfigurationReference) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *LdapConfigurationReference) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *LdapConfigurationReference) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetInfo

`func (o *LdapConfigurationReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *LdapConfigurationReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *LdapConfigurationReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *LdapConfigurationReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


