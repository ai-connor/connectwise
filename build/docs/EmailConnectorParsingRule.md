# EmailConnectorParsingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ParsingStyle** | Pointer to [**EmailConnectorParsingStyleReference**](EmailConnectorParsingStyleReference.md) |  | [optional] 
**Priority** | **NullableInt32** |  | 
**ParsingVariable** | [**EmailConnectorParsingVariableReference**](EmailConnectorParsingVariableReference.md) |  | 
**SearchTerm** | **string** |  Max length: 250; | 
**ServicePriority** | Pointer to [**PriorityReference**](PriorityReference.md) |  | [optional] 
**ServiceStatus** | Pointer to [**ServiceStatusReference**](ServiceStatusReference.md) |  | [optional] 
**ServiceType** | Pointer to [**ServiceTypeReference**](ServiceTypeReference.md) |  | [optional] 
**ServiceSubType** | Pointer to [**ServiceSubTypeReference**](ServiceSubTypeReference.md) |  | [optional] 
**ServiceItem** | Pointer to [**ServiceItemReference**](ServiceItemReference.md) |  | [optional] 
**ServiceBoard** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewEmailConnectorParsingRule

`func NewEmailConnectorParsingRule(priority NullableInt32, parsingVariable EmailConnectorParsingVariableReference, searchTerm string, ) *EmailConnectorParsingRule`

NewEmailConnectorParsingRule instantiates a new EmailConnectorParsingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailConnectorParsingRuleWithDefaults

`func NewEmailConnectorParsingRuleWithDefaults() *EmailConnectorParsingRule`

NewEmailConnectorParsingRuleWithDefaults instantiates a new EmailConnectorParsingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailConnectorParsingRule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailConnectorParsingRule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailConnectorParsingRule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EmailConnectorParsingRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParsingStyle

`func (o *EmailConnectorParsingRule) GetParsingStyle() EmailConnectorParsingStyleReference`

GetParsingStyle returns the ParsingStyle field if non-nil, zero value otherwise.

### GetParsingStyleOk

`func (o *EmailConnectorParsingRule) GetParsingStyleOk() (*EmailConnectorParsingStyleReference, bool)`

GetParsingStyleOk returns a tuple with the ParsingStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsingStyle

`func (o *EmailConnectorParsingRule) SetParsingStyle(v EmailConnectorParsingStyleReference)`

SetParsingStyle sets ParsingStyle field to given value.

### HasParsingStyle

`func (o *EmailConnectorParsingRule) HasParsingStyle() bool`

HasParsingStyle returns a boolean if a field has been set.

### GetPriority

`func (o *EmailConnectorParsingRule) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *EmailConnectorParsingRule) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *EmailConnectorParsingRule) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### SetPriorityNil

`func (o *EmailConnectorParsingRule) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *EmailConnectorParsingRule) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetParsingVariable

`func (o *EmailConnectorParsingRule) GetParsingVariable() EmailConnectorParsingVariableReference`

GetParsingVariable returns the ParsingVariable field if non-nil, zero value otherwise.

### GetParsingVariableOk

`func (o *EmailConnectorParsingRule) GetParsingVariableOk() (*EmailConnectorParsingVariableReference, bool)`

GetParsingVariableOk returns a tuple with the ParsingVariable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsingVariable

`func (o *EmailConnectorParsingRule) SetParsingVariable(v EmailConnectorParsingVariableReference)`

SetParsingVariable sets ParsingVariable field to given value.


### GetSearchTerm

`func (o *EmailConnectorParsingRule) GetSearchTerm() string`

GetSearchTerm returns the SearchTerm field if non-nil, zero value otherwise.

### GetSearchTermOk

`func (o *EmailConnectorParsingRule) GetSearchTermOk() (*string, bool)`

GetSearchTermOk returns a tuple with the SearchTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchTerm

`func (o *EmailConnectorParsingRule) SetSearchTerm(v string)`

SetSearchTerm sets SearchTerm field to given value.


### GetServicePriority

`func (o *EmailConnectorParsingRule) GetServicePriority() PriorityReference`

GetServicePriority returns the ServicePriority field if non-nil, zero value otherwise.

### GetServicePriorityOk

`func (o *EmailConnectorParsingRule) GetServicePriorityOk() (*PriorityReference, bool)`

GetServicePriorityOk returns a tuple with the ServicePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePriority

`func (o *EmailConnectorParsingRule) SetServicePriority(v PriorityReference)`

SetServicePriority sets ServicePriority field to given value.

### HasServicePriority

`func (o *EmailConnectorParsingRule) HasServicePriority() bool`

HasServicePriority returns a boolean if a field has been set.

### GetServiceStatus

`func (o *EmailConnectorParsingRule) GetServiceStatus() ServiceStatusReference`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *EmailConnectorParsingRule) GetServiceStatusOk() (*ServiceStatusReference, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *EmailConnectorParsingRule) SetServiceStatus(v ServiceStatusReference)`

SetServiceStatus sets ServiceStatus field to given value.

### HasServiceStatus

`func (o *EmailConnectorParsingRule) HasServiceStatus() bool`

HasServiceStatus returns a boolean if a field has been set.

### GetServiceType

`func (o *EmailConnectorParsingRule) GetServiceType() ServiceTypeReference`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *EmailConnectorParsingRule) GetServiceTypeOk() (*ServiceTypeReference, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *EmailConnectorParsingRule) SetServiceType(v ServiceTypeReference)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *EmailConnectorParsingRule) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetServiceSubType

`func (o *EmailConnectorParsingRule) GetServiceSubType() ServiceSubTypeReference`

GetServiceSubType returns the ServiceSubType field if non-nil, zero value otherwise.

### GetServiceSubTypeOk

`func (o *EmailConnectorParsingRule) GetServiceSubTypeOk() (*ServiceSubTypeReference, bool)`

GetServiceSubTypeOk returns a tuple with the ServiceSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSubType

`func (o *EmailConnectorParsingRule) SetServiceSubType(v ServiceSubTypeReference)`

SetServiceSubType sets ServiceSubType field to given value.

### HasServiceSubType

`func (o *EmailConnectorParsingRule) HasServiceSubType() bool`

HasServiceSubType returns a boolean if a field has been set.

### GetServiceItem

`func (o *EmailConnectorParsingRule) GetServiceItem() ServiceItemReference`

GetServiceItem returns the ServiceItem field if non-nil, zero value otherwise.

### GetServiceItemOk

`func (o *EmailConnectorParsingRule) GetServiceItemOk() (*ServiceItemReference, bool)`

GetServiceItemOk returns a tuple with the ServiceItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItem

`func (o *EmailConnectorParsingRule) SetServiceItem(v ServiceItemReference)`

SetServiceItem sets ServiceItem field to given value.

### HasServiceItem

`func (o *EmailConnectorParsingRule) HasServiceItem() bool`

HasServiceItem returns a boolean if a field has been set.

### GetServiceBoard

`func (o *EmailConnectorParsingRule) GetServiceBoard() BoardReference`

GetServiceBoard returns the ServiceBoard field if non-nil, zero value otherwise.

### GetServiceBoardOk

`func (o *EmailConnectorParsingRule) GetServiceBoardOk() (*BoardReference, bool)`

GetServiceBoardOk returns a tuple with the ServiceBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceBoard

`func (o *EmailConnectorParsingRule) SetServiceBoard(v BoardReference)`

SetServiceBoard sets ServiceBoard field to given value.

### HasServiceBoard

`func (o *EmailConnectorParsingRule) HasServiceBoard() bool`

HasServiceBoard returns a boolean if a field has been set.

### GetInfo

`func (o *EmailConnectorParsingRule) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *EmailConnectorParsingRule) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *EmailConnectorParsingRule) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *EmailConnectorParsingRule) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


