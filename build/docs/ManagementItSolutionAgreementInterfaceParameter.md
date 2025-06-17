# ManagementItSolutionAgreementInterfaceParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ManagedDevicesIntegration** | Pointer to [**ManagedDevicesIntegrationReference**](ManagedDevicesIntegrationReference.md) |  | [optional] 
**AgreementType** | [**AgreementTypeReference**](AgreementTypeReference.md) |  | 
**ServerProduct** | Pointer to [**IvItemReference**](IvItemReference.md) |  | [optional] 
**WorkstationProduct** | Pointer to [**IvItemReference**](IvItemReference.md) |  | [optional] 
**SpamStatsProduct** | Pointer to [**IvItemReference**](IvItemReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewManagementItSolutionAgreementInterfaceParameter

`func NewManagementItSolutionAgreementInterfaceParameter(agreementType AgreementTypeReference, ) *ManagementItSolutionAgreementInterfaceParameter`

NewManagementItSolutionAgreementInterfaceParameter instantiates a new ManagementItSolutionAgreementInterfaceParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementItSolutionAgreementInterfaceParameterWithDefaults

`func NewManagementItSolutionAgreementInterfaceParameterWithDefaults() *ManagementItSolutionAgreementInterfaceParameter`

NewManagementItSolutionAgreementInterfaceParameterWithDefaults instantiates a new ManagementItSolutionAgreementInterfaceParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagementItSolutionAgreementInterfaceParameter) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagementItSolutionAgreementInterfaceParameter) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagementItSolutionAgreementInterfaceParameter) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManagementItSolutionAgreementInterfaceParameter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManagedDevicesIntegration

`func (o *ManagementItSolutionAgreementInterfaceParameter) GetManagedDevicesIntegration() ManagedDevicesIntegrationReference`

GetManagedDevicesIntegration returns the ManagedDevicesIntegration field if non-nil, zero value otherwise.

### GetManagedDevicesIntegrationOk

`func (o *ManagementItSolutionAgreementInterfaceParameter) GetManagedDevicesIntegrationOk() (*ManagedDevicesIntegrationReference, bool)`

GetManagedDevicesIntegrationOk returns a tuple with the ManagedDevicesIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedDevicesIntegration

`func (o *ManagementItSolutionAgreementInterfaceParameter) SetManagedDevicesIntegration(v ManagedDevicesIntegrationReference)`

SetManagedDevicesIntegration sets ManagedDevicesIntegration field to given value.

### HasManagedDevicesIntegration

`func (o *ManagementItSolutionAgreementInterfaceParameter) HasManagedDevicesIntegration() bool`

HasManagedDevicesIntegration returns a boolean if a field has been set.

### GetAgreementType

`func (o *ManagementItSolutionAgreementInterfaceParameter) GetAgreementType() AgreementTypeReference`

GetAgreementType returns the AgreementType field if non-nil, zero value otherwise.

### GetAgreementTypeOk

`func (o *ManagementItSolutionAgreementInterfaceParameter) GetAgreementTypeOk() (*AgreementTypeReference, bool)`

GetAgreementTypeOk returns a tuple with the AgreementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementType

`func (o *ManagementItSolutionAgreementInterfaceParameter) SetAgreementType(v AgreementTypeReference)`

SetAgreementType sets AgreementType field to given value.


### GetServerProduct

`func (o *ManagementItSolutionAgreementInterfaceParameter) GetServerProduct() IvItemReference`

GetServerProduct returns the ServerProduct field if non-nil, zero value otherwise.

### GetServerProductOk

`func (o *ManagementItSolutionAgreementInterfaceParameter) GetServerProductOk() (*IvItemReference, bool)`

GetServerProductOk returns a tuple with the ServerProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerProduct

`func (o *ManagementItSolutionAgreementInterfaceParameter) SetServerProduct(v IvItemReference)`

SetServerProduct sets ServerProduct field to given value.

### HasServerProduct

`func (o *ManagementItSolutionAgreementInterfaceParameter) HasServerProduct() bool`

HasServerProduct returns a boolean if a field has been set.

### GetWorkstationProduct

`func (o *ManagementItSolutionAgreementInterfaceParameter) GetWorkstationProduct() IvItemReference`

GetWorkstationProduct returns the WorkstationProduct field if non-nil, zero value otherwise.

### GetWorkstationProductOk

`func (o *ManagementItSolutionAgreementInterfaceParameter) GetWorkstationProductOk() (*IvItemReference, bool)`

GetWorkstationProductOk returns a tuple with the WorkstationProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkstationProduct

`func (o *ManagementItSolutionAgreementInterfaceParameter) SetWorkstationProduct(v IvItemReference)`

SetWorkstationProduct sets WorkstationProduct field to given value.

### HasWorkstationProduct

`func (o *ManagementItSolutionAgreementInterfaceParameter) HasWorkstationProduct() bool`

HasWorkstationProduct returns a boolean if a field has been set.

### GetSpamStatsProduct

`func (o *ManagementItSolutionAgreementInterfaceParameter) GetSpamStatsProduct() IvItemReference`

GetSpamStatsProduct returns the SpamStatsProduct field if non-nil, zero value otherwise.

### GetSpamStatsProductOk

`func (o *ManagementItSolutionAgreementInterfaceParameter) GetSpamStatsProductOk() (*IvItemReference, bool)`

GetSpamStatsProductOk returns a tuple with the SpamStatsProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpamStatsProduct

`func (o *ManagementItSolutionAgreementInterfaceParameter) SetSpamStatsProduct(v IvItemReference)`

SetSpamStatsProduct sets SpamStatsProduct field to given value.

### HasSpamStatsProduct

`func (o *ManagementItSolutionAgreementInterfaceParameter) HasSpamStatsProduct() bool`

HasSpamStatsProduct returns a boolean if a field has been set.

### GetInfo

`func (o *ManagementItSolutionAgreementInterfaceParameter) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ManagementItSolutionAgreementInterfaceParameter) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ManagementItSolutionAgreementInterfaceParameter) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ManagementItSolutionAgreementInterfaceParameter) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


