# EmailToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AddressFlag** | Pointer to **NullableBool** |  | [optional] 
**AgreementFlag** | Pointer to **NullableBool** |  | [optional] 
**CompanyFlag** | Pointer to **NullableBool** |  | [optional] 
**ConfigFlag** | Pointer to **NullableBool** |  | [optional] 
**ContactFlag** | Pointer to **NullableBool** |  | [optional] 
**InvoiceFlag** | Pointer to **NullableBool** |  | [optional] 
**PurchaseOrderFlag** | Pointer to **NullableBool** |  | [optional] 
**PurchaseOrderStatusFlag** | Pointer to **NullableBool** |  | [optional] 
**RmaFlag** | Pointer to **NullableBool** |  | [optional] 
**SalesFlag** | Pointer to **NullableBool** |  | [optional] 
**ServiceFlag** | Pointer to **NullableBool** |  | [optional] 
**TracksFlag** | Pointer to **NullableBool** |  | [optional] 
**WorkflowFlag** | Pointer to **NullableBool** |  | [optional] 
**PortalPasswordFlag** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewEmailToken

`func NewEmailToken() *EmailToken`

NewEmailToken instantiates a new EmailToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailTokenWithDefaults

`func NewEmailTokenWithDefaults() *EmailToken`

NewEmailTokenWithDefaults instantiates a new EmailToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailToken) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailToken) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailToken) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EmailToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetToken

`func (o *EmailToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EmailToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EmailToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *EmailToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetDescription

`func (o *EmailToken) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EmailToken) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EmailToken) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EmailToken) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAddressFlag

`func (o *EmailToken) GetAddressFlag() bool`

GetAddressFlag returns the AddressFlag field if non-nil, zero value otherwise.

### GetAddressFlagOk

`func (o *EmailToken) GetAddressFlagOk() (*bool, bool)`

GetAddressFlagOk returns a tuple with the AddressFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFlag

`func (o *EmailToken) SetAddressFlag(v bool)`

SetAddressFlag sets AddressFlag field to given value.

### HasAddressFlag

`func (o *EmailToken) HasAddressFlag() bool`

HasAddressFlag returns a boolean if a field has been set.

### SetAddressFlagNil

`func (o *EmailToken) SetAddressFlagNil(b bool)`

 SetAddressFlagNil sets the value for AddressFlag to be an explicit nil

### UnsetAddressFlag
`func (o *EmailToken) UnsetAddressFlag()`

UnsetAddressFlag ensures that no value is present for AddressFlag, not even an explicit nil
### GetAgreementFlag

`func (o *EmailToken) GetAgreementFlag() bool`

GetAgreementFlag returns the AgreementFlag field if non-nil, zero value otherwise.

### GetAgreementFlagOk

`func (o *EmailToken) GetAgreementFlagOk() (*bool, bool)`

GetAgreementFlagOk returns a tuple with the AgreementFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementFlag

`func (o *EmailToken) SetAgreementFlag(v bool)`

SetAgreementFlag sets AgreementFlag field to given value.

### HasAgreementFlag

`func (o *EmailToken) HasAgreementFlag() bool`

HasAgreementFlag returns a boolean if a field has been set.

### SetAgreementFlagNil

`func (o *EmailToken) SetAgreementFlagNil(b bool)`

 SetAgreementFlagNil sets the value for AgreementFlag to be an explicit nil

### UnsetAgreementFlag
`func (o *EmailToken) UnsetAgreementFlag()`

UnsetAgreementFlag ensures that no value is present for AgreementFlag, not even an explicit nil
### GetCompanyFlag

`func (o *EmailToken) GetCompanyFlag() bool`

GetCompanyFlag returns the CompanyFlag field if non-nil, zero value otherwise.

### GetCompanyFlagOk

`func (o *EmailToken) GetCompanyFlagOk() (*bool, bool)`

GetCompanyFlagOk returns a tuple with the CompanyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyFlag

`func (o *EmailToken) SetCompanyFlag(v bool)`

SetCompanyFlag sets CompanyFlag field to given value.

### HasCompanyFlag

`func (o *EmailToken) HasCompanyFlag() bool`

HasCompanyFlag returns a boolean if a field has been set.

### SetCompanyFlagNil

`func (o *EmailToken) SetCompanyFlagNil(b bool)`

 SetCompanyFlagNil sets the value for CompanyFlag to be an explicit nil

### UnsetCompanyFlag
`func (o *EmailToken) UnsetCompanyFlag()`

UnsetCompanyFlag ensures that no value is present for CompanyFlag, not even an explicit nil
### GetConfigFlag

`func (o *EmailToken) GetConfigFlag() bool`

GetConfigFlag returns the ConfigFlag field if non-nil, zero value otherwise.

### GetConfigFlagOk

`func (o *EmailToken) GetConfigFlagOk() (*bool, bool)`

GetConfigFlagOk returns a tuple with the ConfigFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigFlag

`func (o *EmailToken) SetConfigFlag(v bool)`

SetConfigFlag sets ConfigFlag field to given value.

### HasConfigFlag

`func (o *EmailToken) HasConfigFlag() bool`

HasConfigFlag returns a boolean if a field has been set.

### SetConfigFlagNil

`func (o *EmailToken) SetConfigFlagNil(b bool)`

 SetConfigFlagNil sets the value for ConfigFlag to be an explicit nil

### UnsetConfigFlag
`func (o *EmailToken) UnsetConfigFlag()`

UnsetConfigFlag ensures that no value is present for ConfigFlag, not even an explicit nil
### GetContactFlag

`func (o *EmailToken) GetContactFlag() bool`

GetContactFlag returns the ContactFlag field if non-nil, zero value otherwise.

### GetContactFlagOk

`func (o *EmailToken) GetContactFlagOk() (*bool, bool)`

GetContactFlagOk returns a tuple with the ContactFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactFlag

`func (o *EmailToken) SetContactFlag(v bool)`

SetContactFlag sets ContactFlag field to given value.

### HasContactFlag

`func (o *EmailToken) HasContactFlag() bool`

HasContactFlag returns a boolean if a field has been set.

### SetContactFlagNil

`func (o *EmailToken) SetContactFlagNil(b bool)`

 SetContactFlagNil sets the value for ContactFlag to be an explicit nil

### UnsetContactFlag
`func (o *EmailToken) UnsetContactFlag()`

UnsetContactFlag ensures that no value is present for ContactFlag, not even an explicit nil
### GetInvoiceFlag

`func (o *EmailToken) GetInvoiceFlag() bool`

GetInvoiceFlag returns the InvoiceFlag field if non-nil, zero value otherwise.

### GetInvoiceFlagOk

`func (o *EmailToken) GetInvoiceFlagOk() (*bool, bool)`

GetInvoiceFlagOk returns a tuple with the InvoiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceFlag

`func (o *EmailToken) SetInvoiceFlag(v bool)`

SetInvoiceFlag sets InvoiceFlag field to given value.

### HasInvoiceFlag

`func (o *EmailToken) HasInvoiceFlag() bool`

HasInvoiceFlag returns a boolean if a field has been set.

### SetInvoiceFlagNil

`func (o *EmailToken) SetInvoiceFlagNil(b bool)`

 SetInvoiceFlagNil sets the value for InvoiceFlag to be an explicit nil

### UnsetInvoiceFlag
`func (o *EmailToken) UnsetInvoiceFlag()`

UnsetInvoiceFlag ensures that no value is present for InvoiceFlag, not even an explicit nil
### GetPurchaseOrderFlag

`func (o *EmailToken) GetPurchaseOrderFlag() bool`

GetPurchaseOrderFlag returns the PurchaseOrderFlag field if non-nil, zero value otherwise.

### GetPurchaseOrderFlagOk

`func (o *EmailToken) GetPurchaseOrderFlagOk() (*bool, bool)`

GetPurchaseOrderFlagOk returns a tuple with the PurchaseOrderFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderFlag

`func (o *EmailToken) SetPurchaseOrderFlag(v bool)`

SetPurchaseOrderFlag sets PurchaseOrderFlag field to given value.

### HasPurchaseOrderFlag

`func (o *EmailToken) HasPurchaseOrderFlag() bool`

HasPurchaseOrderFlag returns a boolean if a field has been set.

### SetPurchaseOrderFlagNil

`func (o *EmailToken) SetPurchaseOrderFlagNil(b bool)`

 SetPurchaseOrderFlagNil sets the value for PurchaseOrderFlag to be an explicit nil

### UnsetPurchaseOrderFlag
`func (o *EmailToken) UnsetPurchaseOrderFlag()`

UnsetPurchaseOrderFlag ensures that no value is present for PurchaseOrderFlag, not even an explicit nil
### GetPurchaseOrderStatusFlag

`func (o *EmailToken) GetPurchaseOrderStatusFlag() bool`

GetPurchaseOrderStatusFlag returns the PurchaseOrderStatusFlag field if non-nil, zero value otherwise.

### GetPurchaseOrderStatusFlagOk

`func (o *EmailToken) GetPurchaseOrderStatusFlagOk() (*bool, bool)`

GetPurchaseOrderStatusFlagOk returns a tuple with the PurchaseOrderStatusFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderStatusFlag

`func (o *EmailToken) SetPurchaseOrderStatusFlag(v bool)`

SetPurchaseOrderStatusFlag sets PurchaseOrderStatusFlag field to given value.

### HasPurchaseOrderStatusFlag

`func (o *EmailToken) HasPurchaseOrderStatusFlag() bool`

HasPurchaseOrderStatusFlag returns a boolean if a field has been set.

### SetPurchaseOrderStatusFlagNil

`func (o *EmailToken) SetPurchaseOrderStatusFlagNil(b bool)`

 SetPurchaseOrderStatusFlagNil sets the value for PurchaseOrderStatusFlag to be an explicit nil

### UnsetPurchaseOrderStatusFlag
`func (o *EmailToken) UnsetPurchaseOrderStatusFlag()`

UnsetPurchaseOrderStatusFlag ensures that no value is present for PurchaseOrderStatusFlag, not even an explicit nil
### GetRmaFlag

`func (o *EmailToken) GetRmaFlag() bool`

GetRmaFlag returns the RmaFlag field if non-nil, zero value otherwise.

### GetRmaFlagOk

`func (o *EmailToken) GetRmaFlagOk() (*bool, bool)`

GetRmaFlagOk returns a tuple with the RmaFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmaFlag

`func (o *EmailToken) SetRmaFlag(v bool)`

SetRmaFlag sets RmaFlag field to given value.

### HasRmaFlag

`func (o *EmailToken) HasRmaFlag() bool`

HasRmaFlag returns a boolean if a field has been set.

### SetRmaFlagNil

`func (o *EmailToken) SetRmaFlagNil(b bool)`

 SetRmaFlagNil sets the value for RmaFlag to be an explicit nil

### UnsetRmaFlag
`func (o *EmailToken) UnsetRmaFlag()`

UnsetRmaFlag ensures that no value is present for RmaFlag, not even an explicit nil
### GetSalesFlag

`func (o *EmailToken) GetSalesFlag() bool`

GetSalesFlag returns the SalesFlag field if non-nil, zero value otherwise.

### GetSalesFlagOk

`func (o *EmailToken) GetSalesFlagOk() (*bool, bool)`

GetSalesFlagOk returns a tuple with the SalesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesFlag

`func (o *EmailToken) SetSalesFlag(v bool)`

SetSalesFlag sets SalesFlag field to given value.

### HasSalesFlag

`func (o *EmailToken) HasSalesFlag() bool`

HasSalesFlag returns a boolean if a field has been set.

### SetSalesFlagNil

`func (o *EmailToken) SetSalesFlagNil(b bool)`

 SetSalesFlagNil sets the value for SalesFlag to be an explicit nil

### UnsetSalesFlag
`func (o *EmailToken) UnsetSalesFlag()`

UnsetSalesFlag ensures that no value is present for SalesFlag, not even an explicit nil
### GetServiceFlag

`func (o *EmailToken) GetServiceFlag() bool`

GetServiceFlag returns the ServiceFlag field if non-nil, zero value otherwise.

### GetServiceFlagOk

`func (o *EmailToken) GetServiceFlagOk() (*bool, bool)`

GetServiceFlagOk returns a tuple with the ServiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFlag

`func (o *EmailToken) SetServiceFlag(v bool)`

SetServiceFlag sets ServiceFlag field to given value.

### HasServiceFlag

`func (o *EmailToken) HasServiceFlag() bool`

HasServiceFlag returns a boolean if a field has been set.

### SetServiceFlagNil

`func (o *EmailToken) SetServiceFlagNil(b bool)`

 SetServiceFlagNil sets the value for ServiceFlag to be an explicit nil

### UnsetServiceFlag
`func (o *EmailToken) UnsetServiceFlag()`

UnsetServiceFlag ensures that no value is present for ServiceFlag, not even an explicit nil
### GetTracksFlag

`func (o *EmailToken) GetTracksFlag() bool`

GetTracksFlag returns the TracksFlag field if non-nil, zero value otherwise.

### GetTracksFlagOk

`func (o *EmailToken) GetTracksFlagOk() (*bool, bool)`

GetTracksFlagOk returns a tuple with the TracksFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracksFlag

`func (o *EmailToken) SetTracksFlag(v bool)`

SetTracksFlag sets TracksFlag field to given value.

### HasTracksFlag

`func (o *EmailToken) HasTracksFlag() bool`

HasTracksFlag returns a boolean if a field has been set.

### SetTracksFlagNil

`func (o *EmailToken) SetTracksFlagNil(b bool)`

 SetTracksFlagNil sets the value for TracksFlag to be an explicit nil

### UnsetTracksFlag
`func (o *EmailToken) UnsetTracksFlag()`

UnsetTracksFlag ensures that no value is present for TracksFlag, not even an explicit nil
### GetWorkflowFlag

`func (o *EmailToken) GetWorkflowFlag() bool`

GetWorkflowFlag returns the WorkflowFlag field if non-nil, zero value otherwise.

### GetWorkflowFlagOk

`func (o *EmailToken) GetWorkflowFlagOk() (*bool, bool)`

GetWorkflowFlagOk returns a tuple with the WorkflowFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowFlag

`func (o *EmailToken) SetWorkflowFlag(v bool)`

SetWorkflowFlag sets WorkflowFlag field to given value.

### HasWorkflowFlag

`func (o *EmailToken) HasWorkflowFlag() bool`

HasWorkflowFlag returns a boolean if a field has been set.

### SetWorkflowFlagNil

`func (o *EmailToken) SetWorkflowFlagNil(b bool)`

 SetWorkflowFlagNil sets the value for WorkflowFlag to be an explicit nil

### UnsetWorkflowFlag
`func (o *EmailToken) UnsetWorkflowFlag()`

UnsetWorkflowFlag ensures that no value is present for WorkflowFlag, not even an explicit nil
### GetPortalPasswordFlag

`func (o *EmailToken) GetPortalPasswordFlag() bool`

GetPortalPasswordFlag returns the PortalPasswordFlag field if non-nil, zero value otherwise.

### GetPortalPasswordFlagOk

`func (o *EmailToken) GetPortalPasswordFlagOk() (*bool, bool)`

GetPortalPasswordFlagOk returns a tuple with the PortalPasswordFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalPasswordFlag

`func (o *EmailToken) SetPortalPasswordFlag(v bool)`

SetPortalPasswordFlag sets PortalPasswordFlag field to given value.

### HasPortalPasswordFlag

`func (o *EmailToken) HasPortalPasswordFlag() bool`

HasPortalPasswordFlag returns a boolean if a field has been set.

### SetPortalPasswordFlagNil

`func (o *EmailToken) SetPortalPasswordFlagNil(b bool)`

 SetPortalPasswordFlagNil sets the value for PortalPasswordFlag to be an explicit nil

### UnsetPortalPasswordFlag
`func (o *EmailToken) UnsetPortalPasswordFlag()`

UnsetPortalPasswordFlag ensures that no value is present for PortalPasswordFlag, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


