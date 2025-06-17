# AgreementTypeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**ApplicationUnits** | Pointer to **NullableString** |  | [optional] 
**BillingTerms** | Pointer to [**BillingTermsReference**](BillingTermsReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAgreementTypeInfo

`func NewAgreementTypeInfo() *AgreementTypeInfo`

NewAgreementTypeInfo instantiates a new AgreementTypeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementTypeInfoWithDefaults

`func NewAgreementTypeInfoWithDefaults() *AgreementTypeInfo`

NewAgreementTypeInfoWithDefaults instantiates a new AgreementTypeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgreementTypeInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgreementTypeInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgreementTypeInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AgreementTypeInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AgreementTypeInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgreementTypeInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgreementTypeInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgreementTypeInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInactiveFlag

`func (o *AgreementTypeInfo) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *AgreementTypeInfo) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *AgreementTypeInfo) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *AgreementTypeInfo) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *AgreementTypeInfo) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *AgreementTypeInfo) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetApplicationUnits

`func (o *AgreementTypeInfo) GetApplicationUnits() string`

GetApplicationUnits returns the ApplicationUnits field if non-nil, zero value otherwise.

### GetApplicationUnitsOk

`func (o *AgreementTypeInfo) GetApplicationUnitsOk() (*string, bool)`

GetApplicationUnitsOk returns a tuple with the ApplicationUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUnits

`func (o *AgreementTypeInfo) SetApplicationUnits(v string)`

SetApplicationUnits sets ApplicationUnits field to given value.

### HasApplicationUnits

`func (o *AgreementTypeInfo) HasApplicationUnits() bool`

HasApplicationUnits returns a boolean if a field has been set.

### SetApplicationUnitsNil

`func (o *AgreementTypeInfo) SetApplicationUnitsNil(b bool)`

 SetApplicationUnitsNil sets the value for ApplicationUnits to be an explicit nil

### UnsetApplicationUnits
`func (o *AgreementTypeInfo) UnsetApplicationUnits()`

UnsetApplicationUnits ensures that no value is present for ApplicationUnits, not even an explicit nil
### GetBillingTerms

`func (o *AgreementTypeInfo) GetBillingTerms() BillingTermsReference`

GetBillingTerms returns the BillingTerms field if non-nil, zero value otherwise.

### GetBillingTermsOk

`func (o *AgreementTypeInfo) GetBillingTermsOk() (*BillingTermsReference, bool)`

GetBillingTermsOk returns a tuple with the BillingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTerms

`func (o *AgreementTypeInfo) SetBillingTerms(v BillingTermsReference)`

SetBillingTerms sets BillingTerms field to given value.

### HasBillingTerms

`func (o *AgreementTypeInfo) HasBillingTerms() bool`

HasBillingTerms returns a boolean if a field has been set.

### GetInfo

`func (o *AgreementTypeInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AgreementTypeInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AgreementTypeInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AgreementTypeInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


