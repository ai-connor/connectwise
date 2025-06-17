# PurchaseOrderStatusEmailTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**PurchaseOrderStatusReference**](PurchaseOrderStatusReference.md) |  | [optional] 
**UseSenderFlag** | Pointer to **NullableBool** |  | [optional] 
**FirstName** | Pointer to **string** |  Max length: 100; | [optional] 
**LastName** | Pointer to **string** |  Max length: 100; | [optional] 
**EmailAddress** | Pointer to **string** |  Max length: 100; | [optional] 
**Subject** | **string** |  Max length: 200; | 
**Body** | Pointer to **string** |  | [optional] 
**CopySenderFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPurchaseOrderStatusEmailTemplate

`func NewPurchaseOrderStatusEmailTemplate(subject string, ) *PurchaseOrderStatusEmailTemplate`

NewPurchaseOrderStatusEmailTemplate instantiates a new PurchaseOrderStatusEmailTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseOrderStatusEmailTemplateWithDefaults

`func NewPurchaseOrderStatusEmailTemplateWithDefaults() *PurchaseOrderStatusEmailTemplate`

NewPurchaseOrderStatusEmailTemplateWithDefaults instantiates a new PurchaseOrderStatusEmailTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PurchaseOrderStatusEmailTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PurchaseOrderStatusEmailTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PurchaseOrderStatusEmailTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PurchaseOrderStatusEmailTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *PurchaseOrderStatusEmailTemplate) GetStatus() PurchaseOrderStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PurchaseOrderStatusEmailTemplate) GetStatusOk() (*PurchaseOrderStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PurchaseOrderStatusEmailTemplate) SetStatus(v PurchaseOrderStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PurchaseOrderStatusEmailTemplate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUseSenderFlag

`func (o *PurchaseOrderStatusEmailTemplate) GetUseSenderFlag() bool`

GetUseSenderFlag returns the UseSenderFlag field if non-nil, zero value otherwise.

### GetUseSenderFlagOk

`func (o *PurchaseOrderStatusEmailTemplate) GetUseSenderFlagOk() (*bool, bool)`

GetUseSenderFlagOk returns a tuple with the UseSenderFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSenderFlag

`func (o *PurchaseOrderStatusEmailTemplate) SetUseSenderFlag(v bool)`

SetUseSenderFlag sets UseSenderFlag field to given value.

### HasUseSenderFlag

`func (o *PurchaseOrderStatusEmailTemplate) HasUseSenderFlag() bool`

HasUseSenderFlag returns a boolean if a field has been set.

### SetUseSenderFlagNil

`func (o *PurchaseOrderStatusEmailTemplate) SetUseSenderFlagNil(b bool)`

 SetUseSenderFlagNil sets the value for UseSenderFlag to be an explicit nil

### UnsetUseSenderFlag
`func (o *PurchaseOrderStatusEmailTemplate) UnsetUseSenderFlag()`

UnsetUseSenderFlag ensures that no value is present for UseSenderFlag, not even an explicit nil
### GetFirstName

`func (o *PurchaseOrderStatusEmailTemplate) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PurchaseOrderStatusEmailTemplate) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PurchaseOrderStatusEmailTemplate) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PurchaseOrderStatusEmailTemplate) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *PurchaseOrderStatusEmailTemplate) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PurchaseOrderStatusEmailTemplate) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PurchaseOrderStatusEmailTemplate) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PurchaseOrderStatusEmailTemplate) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *PurchaseOrderStatusEmailTemplate) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *PurchaseOrderStatusEmailTemplate) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *PurchaseOrderStatusEmailTemplate) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *PurchaseOrderStatusEmailTemplate) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetSubject

`func (o *PurchaseOrderStatusEmailTemplate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *PurchaseOrderStatusEmailTemplate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *PurchaseOrderStatusEmailTemplate) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetBody

`func (o *PurchaseOrderStatusEmailTemplate) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PurchaseOrderStatusEmailTemplate) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PurchaseOrderStatusEmailTemplate) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *PurchaseOrderStatusEmailTemplate) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCopySenderFlag

`func (o *PurchaseOrderStatusEmailTemplate) GetCopySenderFlag() bool`

GetCopySenderFlag returns the CopySenderFlag field if non-nil, zero value otherwise.

### GetCopySenderFlagOk

`func (o *PurchaseOrderStatusEmailTemplate) GetCopySenderFlagOk() (*bool, bool)`

GetCopySenderFlagOk returns a tuple with the CopySenderFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopySenderFlag

`func (o *PurchaseOrderStatusEmailTemplate) SetCopySenderFlag(v bool)`

SetCopySenderFlag sets CopySenderFlag field to given value.

### HasCopySenderFlag

`func (o *PurchaseOrderStatusEmailTemplate) HasCopySenderFlag() bool`

HasCopySenderFlag returns a boolean if a field has been set.

### SetCopySenderFlagNil

`func (o *PurchaseOrderStatusEmailTemplate) SetCopySenderFlagNil(b bool)`

 SetCopySenderFlagNil sets the value for CopySenderFlag to be an explicit nil

### UnsetCopySenderFlag
`func (o *PurchaseOrderStatusEmailTemplate) UnsetCopySenderFlag()`

UnsetCopySenderFlag ensures that no value is present for CopySenderFlag, not even an explicit nil
### GetInfo

`func (o *PurchaseOrderStatusEmailTemplate) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PurchaseOrderStatusEmailTemplate) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PurchaseOrderStatusEmailTemplate) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PurchaseOrderStatusEmailTemplate) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


