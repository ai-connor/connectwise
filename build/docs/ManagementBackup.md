# ManagementBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | [**AgreementTypeReference**](AgreementTypeReference.md) |  | 
**Item** | [**CatalogItemReference**](CatalogItemReference.md) |  | 
**BillingLevel** | **NullableString** |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewManagementBackup

`func NewManagementBackup(type_ AgreementTypeReference, item CatalogItemReference, billingLevel NullableString, ) *ManagementBackup`

NewManagementBackup instantiates a new ManagementBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementBackupWithDefaults

`func NewManagementBackupWithDefaults() *ManagementBackup`

NewManagementBackupWithDefaults instantiates a new ManagementBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagementBackup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagementBackup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagementBackup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManagementBackup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ManagementBackup) GetType() AgreementTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ManagementBackup) GetTypeOk() (*AgreementTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ManagementBackup) SetType(v AgreementTypeReference)`

SetType sets Type field to given value.


### GetItem

`func (o *ManagementBackup) GetItem() CatalogItemReference`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ManagementBackup) GetItemOk() (*CatalogItemReference, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ManagementBackup) SetItem(v CatalogItemReference)`

SetItem sets Item field to given value.


### GetBillingLevel

`func (o *ManagementBackup) GetBillingLevel() string`

GetBillingLevel returns the BillingLevel field if non-nil, zero value otherwise.

### GetBillingLevelOk

`func (o *ManagementBackup) GetBillingLevelOk() (*string, bool)`

GetBillingLevelOk returns a tuple with the BillingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingLevel

`func (o *ManagementBackup) SetBillingLevel(v string)`

SetBillingLevel sets BillingLevel field to given value.


### SetBillingLevelNil

`func (o *ManagementBackup) SetBillingLevelNil(b bool)`

 SetBillingLevelNil sets the value for BillingLevel to be an explicit nil

### UnsetBillingLevel
`func (o *ManagementBackup) UnsetBillingLevel()`

UnsetBillingLevel ensures that no value is present for BillingLevel, not even an explicit nil
### GetInfo

`func (o *ManagementBackup) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ManagementBackup) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ManagementBackup) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ManagementBackup) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


