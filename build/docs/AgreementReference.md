# AgreementReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ChargeFirmFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAgreementReference

`func NewAgreementReference() *AgreementReference`

NewAgreementReference instantiates a new AgreementReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementReferenceWithDefaults

`func NewAgreementReferenceWithDefaults() *AgreementReference`

NewAgreementReferenceWithDefaults instantiates a new AgreementReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgreementReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgreementReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgreementReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AgreementReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AgreementReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AgreementReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *AgreementReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgreementReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgreementReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgreementReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AgreementReference) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgreementReference) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgreementReference) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AgreementReference) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChargeFirmFlag

`func (o *AgreementReference) GetChargeFirmFlag() bool`

GetChargeFirmFlag returns the ChargeFirmFlag field if non-nil, zero value otherwise.

### GetChargeFirmFlagOk

`func (o *AgreementReference) GetChargeFirmFlagOk() (*bool, bool)`

GetChargeFirmFlagOk returns a tuple with the ChargeFirmFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeFirmFlag

`func (o *AgreementReference) SetChargeFirmFlag(v bool)`

SetChargeFirmFlag sets ChargeFirmFlag field to given value.

### HasChargeFirmFlag

`func (o *AgreementReference) HasChargeFirmFlag() bool`

HasChargeFirmFlag returns a boolean if a field has been set.

### SetChargeFirmFlagNil

`func (o *AgreementReference) SetChargeFirmFlagNil(b bool)`

 SetChargeFirmFlagNil sets the value for ChargeFirmFlag to be an explicit nil

### UnsetChargeFirmFlag
`func (o *AgreementReference) UnsetChargeFirmFlag()`

UnsetChargeFirmFlag ensures that no value is present for ChargeFirmFlag, not even an explicit nil
### GetInfo

`func (o *AgreementReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AgreementReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AgreementReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AgreementReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


