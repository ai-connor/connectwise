# NotificationRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ExternalFlag** | Pointer to **NullableBool** |  | [optional] 
**ServiceFlag** | Pointer to **NullableBool** |  | [optional] 
**SalesFlag** | Pointer to **NullableBool** |  | [optional] 
**InvoiceFlag** | Pointer to **NullableBool** |  | [optional] 
**AgreementFlag** | Pointer to **NullableBool** |  | [optional] 
**MemberFlag** | Pointer to **NullableBool** |  | [optional] 
**ConfigFlag** | Pointer to **NullableBool** |  | [optional] 
**MspFlag** | Pointer to **NullableBool** |  | [optional] 
**TrackFlag** | Pointer to **NullableBool** |  | [optional] 
**ProjectFlag** | Pointer to **NullableBool** |  | [optional] 
**ProcurementFlag** | Pointer to **NullableBool** |  | [optional] 
**KnowledgeBaseFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewNotificationRecipient

`func NewNotificationRecipient() *NotificationRecipient`

NewNotificationRecipient instantiates a new NotificationRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRecipientWithDefaults

`func NewNotificationRecipientWithDefaults() *NotificationRecipient`

NewNotificationRecipientWithDefaults instantiates a new NotificationRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationRecipient) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationRecipient) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationRecipient) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationRecipient) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *NotificationRecipient) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *NotificationRecipient) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *NotificationRecipient) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *NotificationRecipient) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *NotificationRecipient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationRecipient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationRecipient) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationRecipient) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalFlag

`func (o *NotificationRecipient) GetExternalFlag() bool`

GetExternalFlag returns the ExternalFlag field if non-nil, zero value otherwise.

### GetExternalFlagOk

`func (o *NotificationRecipient) GetExternalFlagOk() (*bool, bool)`

GetExternalFlagOk returns a tuple with the ExternalFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFlag

`func (o *NotificationRecipient) SetExternalFlag(v bool)`

SetExternalFlag sets ExternalFlag field to given value.

### HasExternalFlag

`func (o *NotificationRecipient) HasExternalFlag() bool`

HasExternalFlag returns a boolean if a field has been set.

### SetExternalFlagNil

`func (o *NotificationRecipient) SetExternalFlagNil(b bool)`

 SetExternalFlagNil sets the value for ExternalFlag to be an explicit nil

### UnsetExternalFlag
`func (o *NotificationRecipient) UnsetExternalFlag()`

UnsetExternalFlag ensures that no value is present for ExternalFlag, not even an explicit nil
### GetServiceFlag

`func (o *NotificationRecipient) GetServiceFlag() bool`

GetServiceFlag returns the ServiceFlag field if non-nil, zero value otherwise.

### GetServiceFlagOk

`func (o *NotificationRecipient) GetServiceFlagOk() (*bool, bool)`

GetServiceFlagOk returns a tuple with the ServiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFlag

`func (o *NotificationRecipient) SetServiceFlag(v bool)`

SetServiceFlag sets ServiceFlag field to given value.

### HasServiceFlag

`func (o *NotificationRecipient) HasServiceFlag() bool`

HasServiceFlag returns a boolean if a field has been set.

### SetServiceFlagNil

`func (o *NotificationRecipient) SetServiceFlagNil(b bool)`

 SetServiceFlagNil sets the value for ServiceFlag to be an explicit nil

### UnsetServiceFlag
`func (o *NotificationRecipient) UnsetServiceFlag()`

UnsetServiceFlag ensures that no value is present for ServiceFlag, not even an explicit nil
### GetSalesFlag

`func (o *NotificationRecipient) GetSalesFlag() bool`

GetSalesFlag returns the SalesFlag field if non-nil, zero value otherwise.

### GetSalesFlagOk

`func (o *NotificationRecipient) GetSalesFlagOk() (*bool, bool)`

GetSalesFlagOk returns a tuple with the SalesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesFlag

`func (o *NotificationRecipient) SetSalesFlag(v bool)`

SetSalesFlag sets SalesFlag field to given value.

### HasSalesFlag

`func (o *NotificationRecipient) HasSalesFlag() bool`

HasSalesFlag returns a boolean if a field has been set.

### SetSalesFlagNil

`func (o *NotificationRecipient) SetSalesFlagNil(b bool)`

 SetSalesFlagNil sets the value for SalesFlag to be an explicit nil

### UnsetSalesFlag
`func (o *NotificationRecipient) UnsetSalesFlag()`

UnsetSalesFlag ensures that no value is present for SalesFlag, not even an explicit nil
### GetInvoiceFlag

`func (o *NotificationRecipient) GetInvoiceFlag() bool`

GetInvoiceFlag returns the InvoiceFlag field if non-nil, zero value otherwise.

### GetInvoiceFlagOk

`func (o *NotificationRecipient) GetInvoiceFlagOk() (*bool, bool)`

GetInvoiceFlagOk returns a tuple with the InvoiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceFlag

`func (o *NotificationRecipient) SetInvoiceFlag(v bool)`

SetInvoiceFlag sets InvoiceFlag field to given value.

### HasInvoiceFlag

`func (o *NotificationRecipient) HasInvoiceFlag() bool`

HasInvoiceFlag returns a boolean if a field has been set.

### SetInvoiceFlagNil

`func (o *NotificationRecipient) SetInvoiceFlagNil(b bool)`

 SetInvoiceFlagNil sets the value for InvoiceFlag to be an explicit nil

### UnsetInvoiceFlag
`func (o *NotificationRecipient) UnsetInvoiceFlag()`

UnsetInvoiceFlag ensures that no value is present for InvoiceFlag, not even an explicit nil
### GetAgreementFlag

`func (o *NotificationRecipient) GetAgreementFlag() bool`

GetAgreementFlag returns the AgreementFlag field if non-nil, zero value otherwise.

### GetAgreementFlagOk

`func (o *NotificationRecipient) GetAgreementFlagOk() (*bool, bool)`

GetAgreementFlagOk returns a tuple with the AgreementFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementFlag

`func (o *NotificationRecipient) SetAgreementFlag(v bool)`

SetAgreementFlag sets AgreementFlag field to given value.

### HasAgreementFlag

`func (o *NotificationRecipient) HasAgreementFlag() bool`

HasAgreementFlag returns a boolean if a field has been set.

### SetAgreementFlagNil

`func (o *NotificationRecipient) SetAgreementFlagNil(b bool)`

 SetAgreementFlagNil sets the value for AgreementFlag to be an explicit nil

### UnsetAgreementFlag
`func (o *NotificationRecipient) UnsetAgreementFlag()`

UnsetAgreementFlag ensures that no value is present for AgreementFlag, not even an explicit nil
### GetMemberFlag

`func (o *NotificationRecipient) GetMemberFlag() bool`

GetMemberFlag returns the MemberFlag field if non-nil, zero value otherwise.

### GetMemberFlagOk

`func (o *NotificationRecipient) GetMemberFlagOk() (*bool, bool)`

GetMemberFlagOk returns a tuple with the MemberFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberFlag

`func (o *NotificationRecipient) SetMemberFlag(v bool)`

SetMemberFlag sets MemberFlag field to given value.

### HasMemberFlag

`func (o *NotificationRecipient) HasMemberFlag() bool`

HasMemberFlag returns a boolean if a field has been set.

### SetMemberFlagNil

`func (o *NotificationRecipient) SetMemberFlagNil(b bool)`

 SetMemberFlagNil sets the value for MemberFlag to be an explicit nil

### UnsetMemberFlag
`func (o *NotificationRecipient) UnsetMemberFlag()`

UnsetMemberFlag ensures that no value is present for MemberFlag, not even an explicit nil
### GetConfigFlag

`func (o *NotificationRecipient) GetConfigFlag() bool`

GetConfigFlag returns the ConfigFlag field if non-nil, zero value otherwise.

### GetConfigFlagOk

`func (o *NotificationRecipient) GetConfigFlagOk() (*bool, bool)`

GetConfigFlagOk returns a tuple with the ConfigFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigFlag

`func (o *NotificationRecipient) SetConfigFlag(v bool)`

SetConfigFlag sets ConfigFlag field to given value.

### HasConfigFlag

`func (o *NotificationRecipient) HasConfigFlag() bool`

HasConfigFlag returns a boolean if a field has been set.

### SetConfigFlagNil

`func (o *NotificationRecipient) SetConfigFlagNil(b bool)`

 SetConfigFlagNil sets the value for ConfigFlag to be an explicit nil

### UnsetConfigFlag
`func (o *NotificationRecipient) UnsetConfigFlag()`

UnsetConfigFlag ensures that no value is present for ConfigFlag, not even an explicit nil
### GetMspFlag

`func (o *NotificationRecipient) GetMspFlag() bool`

GetMspFlag returns the MspFlag field if non-nil, zero value otherwise.

### GetMspFlagOk

`func (o *NotificationRecipient) GetMspFlagOk() (*bool, bool)`

GetMspFlagOk returns a tuple with the MspFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMspFlag

`func (o *NotificationRecipient) SetMspFlag(v bool)`

SetMspFlag sets MspFlag field to given value.

### HasMspFlag

`func (o *NotificationRecipient) HasMspFlag() bool`

HasMspFlag returns a boolean if a field has been set.

### SetMspFlagNil

`func (o *NotificationRecipient) SetMspFlagNil(b bool)`

 SetMspFlagNil sets the value for MspFlag to be an explicit nil

### UnsetMspFlag
`func (o *NotificationRecipient) UnsetMspFlag()`

UnsetMspFlag ensures that no value is present for MspFlag, not even an explicit nil
### GetTrackFlag

`func (o *NotificationRecipient) GetTrackFlag() bool`

GetTrackFlag returns the TrackFlag field if non-nil, zero value otherwise.

### GetTrackFlagOk

`func (o *NotificationRecipient) GetTrackFlagOk() (*bool, bool)`

GetTrackFlagOk returns a tuple with the TrackFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackFlag

`func (o *NotificationRecipient) SetTrackFlag(v bool)`

SetTrackFlag sets TrackFlag field to given value.

### HasTrackFlag

`func (o *NotificationRecipient) HasTrackFlag() bool`

HasTrackFlag returns a boolean if a field has been set.

### SetTrackFlagNil

`func (o *NotificationRecipient) SetTrackFlagNil(b bool)`

 SetTrackFlagNil sets the value for TrackFlag to be an explicit nil

### UnsetTrackFlag
`func (o *NotificationRecipient) UnsetTrackFlag()`

UnsetTrackFlag ensures that no value is present for TrackFlag, not even an explicit nil
### GetProjectFlag

`func (o *NotificationRecipient) GetProjectFlag() bool`

GetProjectFlag returns the ProjectFlag field if non-nil, zero value otherwise.

### GetProjectFlagOk

`func (o *NotificationRecipient) GetProjectFlagOk() (*bool, bool)`

GetProjectFlagOk returns a tuple with the ProjectFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectFlag

`func (o *NotificationRecipient) SetProjectFlag(v bool)`

SetProjectFlag sets ProjectFlag field to given value.

### HasProjectFlag

`func (o *NotificationRecipient) HasProjectFlag() bool`

HasProjectFlag returns a boolean if a field has been set.

### SetProjectFlagNil

`func (o *NotificationRecipient) SetProjectFlagNil(b bool)`

 SetProjectFlagNil sets the value for ProjectFlag to be an explicit nil

### UnsetProjectFlag
`func (o *NotificationRecipient) UnsetProjectFlag()`

UnsetProjectFlag ensures that no value is present for ProjectFlag, not even an explicit nil
### GetProcurementFlag

`func (o *NotificationRecipient) GetProcurementFlag() bool`

GetProcurementFlag returns the ProcurementFlag field if non-nil, zero value otherwise.

### GetProcurementFlagOk

`func (o *NotificationRecipient) GetProcurementFlagOk() (*bool, bool)`

GetProcurementFlagOk returns a tuple with the ProcurementFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcurementFlag

`func (o *NotificationRecipient) SetProcurementFlag(v bool)`

SetProcurementFlag sets ProcurementFlag field to given value.

### HasProcurementFlag

`func (o *NotificationRecipient) HasProcurementFlag() bool`

HasProcurementFlag returns a boolean if a field has been set.

### SetProcurementFlagNil

`func (o *NotificationRecipient) SetProcurementFlagNil(b bool)`

 SetProcurementFlagNil sets the value for ProcurementFlag to be an explicit nil

### UnsetProcurementFlag
`func (o *NotificationRecipient) UnsetProcurementFlag()`

UnsetProcurementFlag ensures that no value is present for ProcurementFlag, not even an explicit nil
### GetKnowledgeBaseFlag

`func (o *NotificationRecipient) GetKnowledgeBaseFlag() bool`

GetKnowledgeBaseFlag returns the KnowledgeBaseFlag field if non-nil, zero value otherwise.

### GetKnowledgeBaseFlagOk

`func (o *NotificationRecipient) GetKnowledgeBaseFlagOk() (*bool, bool)`

GetKnowledgeBaseFlagOk returns a tuple with the KnowledgeBaseFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeBaseFlag

`func (o *NotificationRecipient) SetKnowledgeBaseFlag(v bool)`

SetKnowledgeBaseFlag sets KnowledgeBaseFlag field to given value.

### HasKnowledgeBaseFlag

`func (o *NotificationRecipient) HasKnowledgeBaseFlag() bool`

HasKnowledgeBaseFlag returns a boolean if a field has been set.

### SetKnowledgeBaseFlagNil

`func (o *NotificationRecipient) SetKnowledgeBaseFlagNil(b bool)`

 SetKnowledgeBaseFlagNil sets the value for KnowledgeBaseFlag to be an explicit nil

### UnsetKnowledgeBaseFlag
`func (o *NotificationRecipient) UnsetKnowledgeBaseFlag()`

UnsetKnowledgeBaseFlag ensures that no value is present for KnowledgeBaseFlag, not even an explicit nil
### GetInfo

`func (o *NotificationRecipient) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *NotificationRecipient) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *NotificationRecipient) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *NotificationRecipient) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


