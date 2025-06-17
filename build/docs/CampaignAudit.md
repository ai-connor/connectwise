# CampaignAudit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**EmailsSent** | **NullableInt32** |  | 
**EmailsUnsent** | Pointer to **NullableInt32** |  | [optional] 
**DocumentsCreated** | Pointer to **NullableInt32** |  | [optional] 
**EmailSubject** | Pointer to **string** |  Max length: 1000; | [optional] 
**Group** | Pointer to [**GroupReference**](GroupReference.md) |  | [optional] 
**CampaignId** | Pointer to **NullableInt32** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 

## Methods

### NewCampaignAudit

`func NewCampaignAudit(emailsSent NullableInt32, ) *CampaignAudit`

NewCampaignAudit instantiates a new CampaignAudit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignAuditWithDefaults

`func NewCampaignAuditWithDefaults() *CampaignAudit`

NewCampaignAuditWithDefaults instantiates a new CampaignAudit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignAudit) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignAudit) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignAudit) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CampaignAudit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmailsSent

`func (o *CampaignAudit) GetEmailsSent() int32`

GetEmailsSent returns the EmailsSent field if non-nil, zero value otherwise.

### GetEmailsSentOk

`func (o *CampaignAudit) GetEmailsSentOk() (*int32, bool)`

GetEmailsSentOk returns a tuple with the EmailsSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailsSent

`func (o *CampaignAudit) SetEmailsSent(v int32)`

SetEmailsSent sets EmailsSent field to given value.


### SetEmailsSentNil

`func (o *CampaignAudit) SetEmailsSentNil(b bool)`

 SetEmailsSentNil sets the value for EmailsSent to be an explicit nil

### UnsetEmailsSent
`func (o *CampaignAudit) UnsetEmailsSent()`

UnsetEmailsSent ensures that no value is present for EmailsSent, not even an explicit nil
### GetEmailsUnsent

`func (o *CampaignAudit) GetEmailsUnsent() int32`

GetEmailsUnsent returns the EmailsUnsent field if non-nil, zero value otherwise.

### GetEmailsUnsentOk

`func (o *CampaignAudit) GetEmailsUnsentOk() (*int32, bool)`

GetEmailsUnsentOk returns a tuple with the EmailsUnsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailsUnsent

`func (o *CampaignAudit) SetEmailsUnsent(v int32)`

SetEmailsUnsent sets EmailsUnsent field to given value.

### HasEmailsUnsent

`func (o *CampaignAudit) HasEmailsUnsent() bool`

HasEmailsUnsent returns a boolean if a field has been set.

### SetEmailsUnsentNil

`func (o *CampaignAudit) SetEmailsUnsentNil(b bool)`

 SetEmailsUnsentNil sets the value for EmailsUnsent to be an explicit nil

### UnsetEmailsUnsent
`func (o *CampaignAudit) UnsetEmailsUnsent()`

UnsetEmailsUnsent ensures that no value is present for EmailsUnsent, not even an explicit nil
### GetDocumentsCreated

`func (o *CampaignAudit) GetDocumentsCreated() int32`

GetDocumentsCreated returns the DocumentsCreated field if non-nil, zero value otherwise.

### GetDocumentsCreatedOk

`func (o *CampaignAudit) GetDocumentsCreatedOk() (*int32, bool)`

GetDocumentsCreatedOk returns a tuple with the DocumentsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentsCreated

`func (o *CampaignAudit) SetDocumentsCreated(v int32)`

SetDocumentsCreated sets DocumentsCreated field to given value.

### HasDocumentsCreated

`func (o *CampaignAudit) HasDocumentsCreated() bool`

HasDocumentsCreated returns a boolean if a field has been set.

### SetDocumentsCreatedNil

`func (o *CampaignAudit) SetDocumentsCreatedNil(b bool)`

 SetDocumentsCreatedNil sets the value for DocumentsCreated to be an explicit nil

### UnsetDocumentsCreated
`func (o *CampaignAudit) UnsetDocumentsCreated()`

UnsetDocumentsCreated ensures that no value is present for DocumentsCreated, not even an explicit nil
### GetEmailSubject

`func (o *CampaignAudit) GetEmailSubject() string`

GetEmailSubject returns the EmailSubject field if non-nil, zero value otherwise.

### GetEmailSubjectOk

`func (o *CampaignAudit) GetEmailSubjectOk() (*string, bool)`

GetEmailSubjectOk returns a tuple with the EmailSubject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSubject

`func (o *CampaignAudit) SetEmailSubject(v string)`

SetEmailSubject sets EmailSubject field to given value.

### HasEmailSubject

`func (o *CampaignAudit) HasEmailSubject() bool`

HasEmailSubject returns a boolean if a field has been set.

### GetGroup

`func (o *CampaignAudit) GetGroup() GroupReference`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CampaignAudit) GetGroupOk() (*GroupReference, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CampaignAudit) SetGroup(v GroupReference)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CampaignAudit) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCampaignId

`func (o *CampaignAudit) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CampaignAudit) GetCampaignIdOk() (*int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CampaignAudit) SetCampaignId(v int32)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *CampaignAudit) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignIdNil

`func (o *CampaignAudit) SetCampaignIdNil(b bool)`

 SetCampaignIdNil sets the value for CampaignId to be an explicit nil

### UnsetCampaignId
`func (o *CampaignAudit) UnsetCampaignId()`

UnsetCampaignId ensures that no value is present for CampaignId, not even an explicit nil
### GetCreatedBy

`func (o *CampaignAudit) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CampaignAudit) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CampaignAudit) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *CampaignAudit) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *CampaignAudit) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *CampaignAudit) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *CampaignAudit) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *CampaignAudit) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


