# CampaignSubTypeCampaignSubType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | [**CampaignTypeReference**](CampaignTypeReference.md) |  | 
**Name** | **string** |  Max length: 100; | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCampaignSubTypeCampaignSubType

`func NewCampaignSubTypeCampaignSubType(type_ CampaignTypeReference, name string, ) *CampaignSubTypeCampaignSubType`

NewCampaignSubTypeCampaignSubType instantiates a new CampaignSubTypeCampaignSubType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignSubTypeCampaignSubTypeWithDefaults

`func NewCampaignSubTypeCampaignSubTypeWithDefaults() *CampaignSubTypeCampaignSubType`

NewCampaignSubTypeCampaignSubTypeWithDefaults instantiates a new CampaignSubTypeCampaignSubType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignSubTypeCampaignSubType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignSubTypeCampaignSubType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignSubTypeCampaignSubType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CampaignSubTypeCampaignSubType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CampaignSubTypeCampaignSubType) GetType() CampaignTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CampaignSubTypeCampaignSubType) GetTypeOk() (*CampaignTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CampaignSubTypeCampaignSubType) SetType(v CampaignTypeReference)`

SetType sets Type field to given value.


### GetName

`func (o *CampaignSubTypeCampaignSubType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignSubTypeCampaignSubType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignSubTypeCampaignSubType) SetName(v string)`

SetName sets Name field to given value.


### GetInfo

`func (o *CampaignSubTypeCampaignSubType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CampaignSubTypeCampaignSubType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CampaignSubTypeCampaignSubType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CampaignSubTypeCampaignSubType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


