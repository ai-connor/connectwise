# InvoiceEmailTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**ServiceSurvey** | Pointer to [**ServiceSurveyReference**](ServiceSurveyReference.md) |  | [optional] 
**UseSenderFlag** | Pointer to **NullableBool** |  | [optional] 
**FirstName** | Pointer to **string** | From fields (first name, last name, email address) are required if useSenderFlag is false. Max length: 100; | [optional] 
**LastName** | Pointer to **string** | From fields (first name, last name, email address) are required if useSenderFlag is false. Max length: 100; | [optional] 
**EmailAddress** | Pointer to **string** | From fields (first name, last name, email address) are required if useSenderFlag is false. Max length: 100; | [optional] 
**Subject** | **string** |  Max length: 200; | 
**Body** | Pointer to **string** |  | [optional] 
**CopySenderFlag** | Pointer to **NullableBool** |  | [optional] 
**InvoiceStatus** | Pointer to [**BillingStatusReference**](BillingStatusReference.md) |  | [optional] 
**AttachInvoiceFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewInvoiceEmailTemplate

`func NewInvoiceEmailTemplate(name string, subject string, ) *InvoiceEmailTemplate`

NewInvoiceEmailTemplate instantiates a new InvoiceEmailTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceEmailTemplateWithDefaults

`func NewInvoiceEmailTemplateWithDefaults() *InvoiceEmailTemplate`

NewInvoiceEmailTemplateWithDefaults instantiates a new InvoiceEmailTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvoiceEmailTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoiceEmailTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoiceEmailTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InvoiceEmailTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *InvoiceEmailTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvoiceEmailTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvoiceEmailTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetServiceSurvey

`func (o *InvoiceEmailTemplate) GetServiceSurvey() ServiceSurveyReference`

GetServiceSurvey returns the ServiceSurvey field if non-nil, zero value otherwise.

### GetServiceSurveyOk

`func (o *InvoiceEmailTemplate) GetServiceSurveyOk() (*ServiceSurveyReference, bool)`

GetServiceSurveyOk returns a tuple with the ServiceSurvey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSurvey

`func (o *InvoiceEmailTemplate) SetServiceSurvey(v ServiceSurveyReference)`

SetServiceSurvey sets ServiceSurvey field to given value.

### HasServiceSurvey

`func (o *InvoiceEmailTemplate) HasServiceSurvey() bool`

HasServiceSurvey returns a boolean if a field has been set.

### GetUseSenderFlag

`func (o *InvoiceEmailTemplate) GetUseSenderFlag() bool`

GetUseSenderFlag returns the UseSenderFlag field if non-nil, zero value otherwise.

### GetUseSenderFlagOk

`func (o *InvoiceEmailTemplate) GetUseSenderFlagOk() (*bool, bool)`

GetUseSenderFlagOk returns a tuple with the UseSenderFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSenderFlag

`func (o *InvoiceEmailTemplate) SetUseSenderFlag(v bool)`

SetUseSenderFlag sets UseSenderFlag field to given value.

### HasUseSenderFlag

`func (o *InvoiceEmailTemplate) HasUseSenderFlag() bool`

HasUseSenderFlag returns a boolean if a field has been set.

### SetUseSenderFlagNil

`func (o *InvoiceEmailTemplate) SetUseSenderFlagNil(b bool)`

 SetUseSenderFlagNil sets the value for UseSenderFlag to be an explicit nil

### UnsetUseSenderFlag
`func (o *InvoiceEmailTemplate) UnsetUseSenderFlag()`

UnsetUseSenderFlag ensures that no value is present for UseSenderFlag, not even an explicit nil
### GetFirstName

`func (o *InvoiceEmailTemplate) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *InvoiceEmailTemplate) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *InvoiceEmailTemplate) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *InvoiceEmailTemplate) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *InvoiceEmailTemplate) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *InvoiceEmailTemplate) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *InvoiceEmailTemplate) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *InvoiceEmailTemplate) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *InvoiceEmailTemplate) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *InvoiceEmailTemplate) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *InvoiceEmailTemplate) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *InvoiceEmailTemplate) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetSubject

`func (o *InvoiceEmailTemplate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *InvoiceEmailTemplate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *InvoiceEmailTemplate) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetBody

`func (o *InvoiceEmailTemplate) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *InvoiceEmailTemplate) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *InvoiceEmailTemplate) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *InvoiceEmailTemplate) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCopySenderFlag

`func (o *InvoiceEmailTemplate) GetCopySenderFlag() bool`

GetCopySenderFlag returns the CopySenderFlag field if non-nil, zero value otherwise.

### GetCopySenderFlagOk

`func (o *InvoiceEmailTemplate) GetCopySenderFlagOk() (*bool, bool)`

GetCopySenderFlagOk returns a tuple with the CopySenderFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopySenderFlag

`func (o *InvoiceEmailTemplate) SetCopySenderFlag(v bool)`

SetCopySenderFlag sets CopySenderFlag field to given value.

### HasCopySenderFlag

`func (o *InvoiceEmailTemplate) HasCopySenderFlag() bool`

HasCopySenderFlag returns a boolean if a field has been set.

### SetCopySenderFlagNil

`func (o *InvoiceEmailTemplate) SetCopySenderFlagNil(b bool)`

 SetCopySenderFlagNil sets the value for CopySenderFlag to be an explicit nil

### UnsetCopySenderFlag
`func (o *InvoiceEmailTemplate) UnsetCopySenderFlag()`

UnsetCopySenderFlag ensures that no value is present for CopySenderFlag, not even an explicit nil
### GetInvoiceStatus

`func (o *InvoiceEmailTemplate) GetInvoiceStatus() BillingStatusReference`

GetInvoiceStatus returns the InvoiceStatus field if non-nil, zero value otherwise.

### GetInvoiceStatusOk

`func (o *InvoiceEmailTemplate) GetInvoiceStatusOk() (*BillingStatusReference, bool)`

GetInvoiceStatusOk returns a tuple with the InvoiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceStatus

`func (o *InvoiceEmailTemplate) SetInvoiceStatus(v BillingStatusReference)`

SetInvoiceStatus sets InvoiceStatus field to given value.

### HasInvoiceStatus

`func (o *InvoiceEmailTemplate) HasInvoiceStatus() bool`

HasInvoiceStatus returns a boolean if a field has been set.

### GetAttachInvoiceFlag

`func (o *InvoiceEmailTemplate) GetAttachInvoiceFlag() bool`

GetAttachInvoiceFlag returns the AttachInvoiceFlag field if non-nil, zero value otherwise.

### GetAttachInvoiceFlagOk

`func (o *InvoiceEmailTemplate) GetAttachInvoiceFlagOk() (*bool, bool)`

GetAttachInvoiceFlagOk returns a tuple with the AttachInvoiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachInvoiceFlag

`func (o *InvoiceEmailTemplate) SetAttachInvoiceFlag(v bool)`

SetAttachInvoiceFlag sets AttachInvoiceFlag field to given value.

### HasAttachInvoiceFlag

`func (o *InvoiceEmailTemplate) HasAttachInvoiceFlag() bool`

HasAttachInvoiceFlag returns a boolean if a field has been set.

### SetAttachInvoiceFlagNil

`func (o *InvoiceEmailTemplate) SetAttachInvoiceFlagNil(b bool)`

 SetAttachInvoiceFlagNil sets the value for AttachInvoiceFlag to be an explicit nil

### UnsetAttachInvoiceFlag
`func (o *InvoiceEmailTemplate) UnsetAttachInvoiceFlag()`

UnsetAttachInvoiceFlag ensures that no value is present for AttachInvoiceFlag, not even an explicit nil
### GetInfo

`func (o *InvoiceEmailTemplate) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *InvoiceEmailTemplate) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *InvoiceEmailTemplate) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *InvoiceEmailTemplate) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


