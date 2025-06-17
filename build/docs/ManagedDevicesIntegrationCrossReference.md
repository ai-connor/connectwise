# ManagedDevicesIntegrationCrossReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ManagedDevicesIntegration** | Pointer to [**ManagedDevicesIntegrationReference**](ManagedDevicesIntegrationReference.md) |  | [optional] 
**VendorType** | Pointer to **string** |  Max length: 255; | [optional] 
**VendorLevel** | Pointer to **string** |  Max length: 255; | [optional] 
**AgreementType** | Pointer to [**AgreementTypeReference**](AgreementTypeReference.md) |  | [optional] 
**Product** | Pointer to [**IvItemReference**](IvItemReference.md) |  | [optional] 
**ConfigurationType** | Pointer to [**ConfigurationTypeReference**](ConfigurationTypeReference.md) |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewManagedDevicesIntegrationCrossReference

`func NewManagedDevicesIntegrationCrossReference() *ManagedDevicesIntegrationCrossReference`

NewManagedDevicesIntegrationCrossReference instantiates a new ManagedDevicesIntegrationCrossReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedDevicesIntegrationCrossReferenceWithDefaults

`func NewManagedDevicesIntegrationCrossReferenceWithDefaults() *ManagedDevicesIntegrationCrossReference`

NewManagedDevicesIntegrationCrossReferenceWithDefaults instantiates a new ManagedDevicesIntegrationCrossReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagedDevicesIntegrationCrossReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedDevicesIntegrationCrossReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedDevicesIntegrationCrossReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManagedDevicesIntegrationCrossReference) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManagedDevicesIntegration

`func (o *ManagedDevicesIntegrationCrossReference) GetManagedDevicesIntegration() ManagedDevicesIntegrationReference`

GetManagedDevicesIntegration returns the ManagedDevicesIntegration field if non-nil, zero value otherwise.

### GetManagedDevicesIntegrationOk

`func (o *ManagedDevicesIntegrationCrossReference) GetManagedDevicesIntegrationOk() (*ManagedDevicesIntegrationReference, bool)`

GetManagedDevicesIntegrationOk returns a tuple with the ManagedDevicesIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedDevicesIntegration

`func (o *ManagedDevicesIntegrationCrossReference) SetManagedDevicesIntegration(v ManagedDevicesIntegrationReference)`

SetManagedDevicesIntegration sets ManagedDevicesIntegration field to given value.

### HasManagedDevicesIntegration

`func (o *ManagedDevicesIntegrationCrossReference) HasManagedDevicesIntegration() bool`

HasManagedDevicesIntegration returns a boolean if a field has been set.

### GetVendorType

`func (o *ManagedDevicesIntegrationCrossReference) GetVendorType() string`

GetVendorType returns the VendorType field if non-nil, zero value otherwise.

### GetVendorTypeOk

`func (o *ManagedDevicesIntegrationCrossReference) GetVendorTypeOk() (*string, bool)`

GetVendorTypeOk returns a tuple with the VendorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorType

`func (o *ManagedDevicesIntegrationCrossReference) SetVendorType(v string)`

SetVendorType sets VendorType field to given value.

### HasVendorType

`func (o *ManagedDevicesIntegrationCrossReference) HasVendorType() bool`

HasVendorType returns a boolean if a field has been set.

### GetVendorLevel

`func (o *ManagedDevicesIntegrationCrossReference) GetVendorLevel() string`

GetVendorLevel returns the VendorLevel field if non-nil, zero value otherwise.

### GetVendorLevelOk

`func (o *ManagedDevicesIntegrationCrossReference) GetVendorLevelOk() (*string, bool)`

GetVendorLevelOk returns a tuple with the VendorLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorLevel

`func (o *ManagedDevicesIntegrationCrossReference) SetVendorLevel(v string)`

SetVendorLevel sets VendorLevel field to given value.

### HasVendorLevel

`func (o *ManagedDevicesIntegrationCrossReference) HasVendorLevel() bool`

HasVendorLevel returns a boolean if a field has been set.

### GetAgreementType

`func (o *ManagedDevicesIntegrationCrossReference) GetAgreementType() AgreementTypeReference`

GetAgreementType returns the AgreementType field if non-nil, zero value otherwise.

### GetAgreementTypeOk

`func (o *ManagedDevicesIntegrationCrossReference) GetAgreementTypeOk() (*AgreementTypeReference, bool)`

GetAgreementTypeOk returns a tuple with the AgreementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementType

`func (o *ManagedDevicesIntegrationCrossReference) SetAgreementType(v AgreementTypeReference)`

SetAgreementType sets AgreementType field to given value.

### HasAgreementType

`func (o *ManagedDevicesIntegrationCrossReference) HasAgreementType() bool`

HasAgreementType returns a boolean if a field has been set.

### GetProduct

`func (o *ManagedDevicesIntegrationCrossReference) GetProduct() IvItemReference`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ManagedDevicesIntegrationCrossReference) GetProductOk() (*IvItemReference, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ManagedDevicesIntegrationCrossReference) SetProduct(v IvItemReference)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *ManagedDevicesIntegrationCrossReference) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetConfigurationType

`func (o *ManagedDevicesIntegrationCrossReference) GetConfigurationType() ConfigurationTypeReference`

GetConfigurationType returns the ConfigurationType field if non-nil, zero value otherwise.

### GetConfigurationTypeOk

`func (o *ManagedDevicesIntegrationCrossReference) GetConfigurationTypeOk() (*ConfigurationTypeReference, bool)`

GetConfigurationTypeOk returns a tuple with the ConfigurationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationType

`func (o *ManagedDevicesIntegrationCrossReference) SetConfigurationType(v ConfigurationTypeReference)`

SetConfigurationType sets ConfigurationType field to given value.

### HasConfigurationType

`func (o *ManagedDevicesIntegrationCrossReference) HasConfigurationType() bool`

HasConfigurationType returns a boolean if a field has been set.

### GetInactiveFlag

`func (o *ManagedDevicesIntegrationCrossReference) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *ManagedDevicesIntegrationCrossReference) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *ManagedDevicesIntegrationCrossReference) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *ManagedDevicesIntegrationCrossReference) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *ManagedDevicesIntegrationCrossReference) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *ManagedDevicesIntegrationCrossReference) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetInfo

`func (o *ManagedDevicesIntegrationCrossReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ManagedDevicesIntegrationCrossReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ManagedDevicesIntegrationCrossReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ManagedDevicesIntegrationCrossReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


