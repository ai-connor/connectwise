# PortalSecurityLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CaptionIdentifier** | Pointer to **string** |  | [optional] 
**IsDefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**Caption** | Pointer to **string** |  Max length: 50; | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPortalSecurityLevel

`func NewPortalSecurityLevel() *PortalSecurityLevel`

NewPortalSecurityLevel instantiates a new PortalSecurityLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalSecurityLevelWithDefaults

`func NewPortalSecurityLevelWithDefaults() *PortalSecurityLevel`

NewPortalSecurityLevelWithDefaults instantiates a new PortalSecurityLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortalSecurityLevel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalSecurityLevel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalSecurityLevel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PortalSecurityLevel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCaptionIdentifier

`func (o *PortalSecurityLevel) GetCaptionIdentifier() string`

GetCaptionIdentifier returns the CaptionIdentifier field if non-nil, zero value otherwise.

### GetCaptionIdentifierOk

`func (o *PortalSecurityLevel) GetCaptionIdentifierOk() (*string, bool)`

GetCaptionIdentifierOk returns a tuple with the CaptionIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptionIdentifier

`func (o *PortalSecurityLevel) SetCaptionIdentifier(v string)`

SetCaptionIdentifier sets CaptionIdentifier field to given value.

### HasCaptionIdentifier

`func (o *PortalSecurityLevel) HasCaptionIdentifier() bool`

HasCaptionIdentifier returns a boolean if a field has been set.

### GetIsDefaultFlag

`func (o *PortalSecurityLevel) GetIsDefaultFlag() bool`

GetIsDefaultFlag returns the IsDefaultFlag field if non-nil, zero value otherwise.

### GetIsDefaultFlagOk

`func (o *PortalSecurityLevel) GetIsDefaultFlagOk() (*bool, bool)`

GetIsDefaultFlagOk returns a tuple with the IsDefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultFlag

`func (o *PortalSecurityLevel) SetIsDefaultFlag(v bool)`

SetIsDefaultFlag sets IsDefaultFlag field to given value.

### HasIsDefaultFlag

`func (o *PortalSecurityLevel) HasIsDefaultFlag() bool`

HasIsDefaultFlag returns a boolean if a field has been set.

### SetIsDefaultFlagNil

`func (o *PortalSecurityLevel) SetIsDefaultFlagNil(b bool)`

 SetIsDefaultFlagNil sets the value for IsDefaultFlag to be an explicit nil

### UnsetIsDefaultFlag
`func (o *PortalSecurityLevel) UnsetIsDefaultFlag()`

UnsetIsDefaultFlag ensures that no value is present for IsDefaultFlag, not even an explicit nil
### GetCaption

`func (o *PortalSecurityLevel) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *PortalSecurityLevel) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *PortalSecurityLevel) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *PortalSecurityLevel) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetInfo

`func (o *PortalSecurityLevel) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PortalSecurityLevel) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PortalSecurityLevel) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PortalSecurityLevel) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


