# AdjustmentType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | **string** |  Max length: 50; | 
**Name** | Pointer to **string** |  Max length: 100; | [optional] 
**AuditTrailFlag** | Pointer to **NullableBool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAdjustmentType

`func NewAdjustmentType(identifier string, ) *AdjustmentType`

NewAdjustmentType instantiates a new AdjustmentType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdjustmentTypeWithDefaults

`func NewAdjustmentTypeWithDefaults() *AdjustmentType`

NewAdjustmentTypeWithDefaults instantiates a new AdjustmentType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdjustmentType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdjustmentType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdjustmentType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AdjustmentType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *AdjustmentType) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *AdjustmentType) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *AdjustmentType) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetName

`func (o *AdjustmentType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdjustmentType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdjustmentType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdjustmentType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAuditTrailFlag

`func (o *AdjustmentType) GetAuditTrailFlag() bool`

GetAuditTrailFlag returns the AuditTrailFlag field if non-nil, zero value otherwise.

### GetAuditTrailFlagOk

`func (o *AdjustmentType) GetAuditTrailFlagOk() (*bool, bool)`

GetAuditTrailFlagOk returns a tuple with the AuditTrailFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditTrailFlag

`func (o *AdjustmentType) SetAuditTrailFlag(v bool)`

SetAuditTrailFlag sets AuditTrailFlag field to given value.

### HasAuditTrailFlag

`func (o *AdjustmentType) HasAuditTrailFlag() bool`

HasAuditTrailFlag returns a boolean if a field has been set.

### SetAuditTrailFlagNil

`func (o *AdjustmentType) SetAuditTrailFlagNil(b bool)`

 SetAuditTrailFlagNil sets the value for AuditTrailFlag to be an explicit nil

### UnsetAuditTrailFlag
`func (o *AdjustmentType) UnsetAuditTrailFlag()`

UnsetAuditTrailFlag ensures that no value is present for AuditTrailFlag, not even an explicit nil
### GetDateCreated

`func (o *AdjustmentType) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AdjustmentType) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AdjustmentType) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AdjustmentType) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AdjustmentType) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AdjustmentType) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AdjustmentType) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AdjustmentType) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetInfo

`func (o *AdjustmentType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AdjustmentType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AdjustmentType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AdjustmentType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


