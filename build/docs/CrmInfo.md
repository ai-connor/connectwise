# CrmInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**AccountManagerRole** | Pointer to [**TeamRoleReference**](TeamRoleReference.md) |  | [optional] 
**TechnicalContactRole** | Pointer to [**TeamRoleReference**](TeamRoleReference.md) |  | [optional] 
**SalesRepRole** | Pointer to [**TeamRoleReference**](TeamRoleReference.md) |  | [optional] 
**Field1Caption** | Pointer to **string** |  | [optional] 
**Field2Caption** | Pointer to **string** |  | [optional] 
**Field3Caption** | Pointer to **string** |  | [optional] 
**Field4Caption** | Pointer to **string** |  | [optional] 
**Field5Caption** | Pointer to **string** |  | [optional] 
**Field6Caption** | Pointer to **string** |  | [optional] 
**Field7Caption** | Pointer to **string** |  | [optional] 
**Field8Caption** | Pointer to **string** |  | [optional] 
**Field9Caption** | Pointer to **string** |  | [optional] 
**Field10Caption** | Pointer to **string** |  | [optional] 
**PrimaryRepCaption** | Pointer to **string** |  | [optional] 
**SecondaryRepCaption** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCrmInfo

`func NewCrmInfo() *CrmInfo`

NewCrmInfo instantiates a new CrmInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrmInfoWithDefaults

`func NewCrmInfoWithDefaults() *CrmInfo`

NewCrmInfoWithDefaults instantiates a new CrmInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CrmInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CrmInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CrmInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CrmInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountManagerRole

`func (o *CrmInfo) GetAccountManagerRole() TeamRoleReference`

GetAccountManagerRole returns the AccountManagerRole field if non-nil, zero value otherwise.

### GetAccountManagerRoleOk

`func (o *CrmInfo) GetAccountManagerRoleOk() (*TeamRoleReference, bool)`

GetAccountManagerRoleOk returns a tuple with the AccountManagerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountManagerRole

`func (o *CrmInfo) SetAccountManagerRole(v TeamRoleReference)`

SetAccountManagerRole sets AccountManagerRole field to given value.

### HasAccountManagerRole

`func (o *CrmInfo) HasAccountManagerRole() bool`

HasAccountManagerRole returns a boolean if a field has been set.

### GetTechnicalContactRole

`func (o *CrmInfo) GetTechnicalContactRole() TeamRoleReference`

GetTechnicalContactRole returns the TechnicalContactRole field if non-nil, zero value otherwise.

### GetTechnicalContactRoleOk

`func (o *CrmInfo) GetTechnicalContactRoleOk() (*TeamRoleReference, bool)`

GetTechnicalContactRoleOk returns a tuple with the TechnicalContactRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalContactRole

`func (o *CrmInfo) SetTechnicalContactRole(v TeamRoleReference)`

SetTechnicalContactRole sets TechnicalContactRole field to given value.

### HasTechnicalContactRole

`func (o *CrmInfo) HasTechnicalContactRole() bool`

HasTechnicalContactRole returns a boolean if a field has been set.

### GetSalesRepRole

`func (o *CrmInfo) GetSalesRepRole() TeamRoleReference`

GetSalesRepRole returns the SalesRepRole field if non-nil, zero value otherwise.

### GetSalesRepRoleOk

`func (o *CrmInfo) GetSalesRepRoleOk() (*TeamRoleReference, bool)`

GetSalesRepRoleOk returns a tuple with the SalesRepRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepRole

`func (o *CrmInfo) SetSalesRepRole(v TeamRoleReference)`

SetSalesRepRole sets SalesRepRole field to given value.

### HasSalesRepRole

`func (o *CrmInfo) HasSalesRepRole() bool`

HasSalesRepRole returns a boolean if a field has been set.

### GetField1Caption

`func (o *CrmInfo) GetField1Caption() string`

GetField1Caption returns the Field1Caption field if non-nil, zero value otherwise.

### GetField1CaptionOk

`func (o *CrmInfo) GetField1CaptionOk() (*string, bool)`

GetField1CaptionOk returns a tuple with the Field1Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField1Caption

`func (o *CrmInfo) SetField1Caption(v string)`

SetField1Caption sets Field1Caption field to given value.

### HasField1Caption

`func (o *CrmInfo) HasField1Caption() bool`

HasField1Caption returns a boolean if a field has been set.

### GetField2Caption

`func (o *CrmInfo) GetField2Caption() string`

GetField2Caption returns the Field2Caption field if non-nil, zero value otherwise.

### GetField2CaptionOk

`func (o *CrmInfo) GetField2CaptionOk() (*string, bool)`

GetField2CaptionOk returns a tuple with the Field2Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField2Caption

`func (o *CrmInfo) SetField2Caption(v string)`

SetField2Caption sets Field2Caption field to given value.

### HasField2Caption

`func (o *CrmInfo) HasField2Caption() bool`

HasField2Caption returns a boolean if a field has been set.

### GetField3Caption

`func (o *CrmInfo) GetField3Caption() string`

GetField3Caption returns the Field3Caption field if non-nil, zero value otherwise.

### GetField3CaptionOk

`func (o *CrmInfo) GetField3CaptionOk() (*string, bool)`

GetField3CaptionOk returns a tuple with the Field3Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField3Caption

`func (o *CrmInfo) SetField3Caption(v string)`

SetField3Caption sets Field3Caption field to given value.

### HasField3Caption

`func (o *CrmInfo) HasField3Caption() bool`

HasField3Caption returns a boolean if a field has been set.

### GetField4Caption

`func (o *CrmInfo) GetField4Caption() string`

GetField4Caption returns the Field4Caption field if non-nil, zero value otherwise.

### GetField4CaptionOk

`func (o *CrmInfo) GetField4CaptionOk() (*string, bool)`

GetField4CaptionOk returns a tuple with the Field4Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField4Caption

`func (o *CrmInfo) SetField4Caption(v string)`

SetField4Caption sets Field4Caption field to given value.

### HasField4Caption

`func (o *CrmInfo) HasField4Caption() bool`

HasField4Caption returns a boolean if a field has been set.

### GetField5Caption

`func (o *CrmInfo) GetField5Caption() string`

GetField5Caption returns the Field5Caption field if non-nil, zero value otherwise.

### GetField5CaptionOk

`func (o *CrmInfo) GetField5CaptionOk() (*string, bool)`

GetField5CaptionOk returns a tuple with the Field5Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField5Caption

`func (o *CrmInfo) SetField5Caption(v string)`

SetField5Caption sets Field5Caption field to given value.

### HasField5Caption

`func (o *CrmInfo) HasField5Caption() bool`

HasField5Caption returns a boolean if a field has been set.

### GetField6Caption

`func (o *CrmInfo) GetField6Caption() string`

GetField6Caption returns the Field6Caption field if non-nil, zero value otherwise.

### GetField6CaptionOk

`func (o *CrmInfo) GetField6CaptionOk() (*string, bool)`

GetField6CaptionOk returns a tuple with the Field6Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField6Caption

`func (o *CrmInfo) SetField6Caption(v string)`

SetField6Caption sets Field6Caption field to given value.

### HasField6Caption

`func (o *CrmInfo) HasField6Caption() bool`

HasField6Caption returns a boolean if a field has been set.

### GetField7Caption

`func (o *CrmInfo) GetField7Caption() string`

GetField7Caption returns the Field7Caption field if non-nil, zero value otherwise.

### GetField7CaptionOk

`func (o *CrmInfo) GetField7CaptionOk() (*string, bool)`

GetField7CaptionOk returns a tuple with the Field7Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField7Caption

`func (o *CrmInfo) SetField7Caption(v string)`

SetField7Caption sets Field7Caption field to given value.

### HasField7Caption

`func (o *CrmInfo) HasField7Caption() bool`

HasField7Caption returns a boolean if a field has been set.

### GetField8Caption

`func (o *CrmInfo) GetField8Caption() string`

GetField8Caption returns the Field8Caption field if non-nil, zero value otherwise.

### GetField8CaptionOk

`func (o *CrmInfo) GetField8CaptionOk() (*string, bool)`

GetField8CaptionOk returns a tuple with the Field8Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField8Caption

`func (o *CrmInfo) SetField8Caption(v string)`

SetField8Caption sets Field8Caption field to given value.

### HasField8Caption

`func (o *CrmInfo) HasField8Caption() bool`

HasField8Caption returns a boolean if a field has been set.

### GetField9Caption

`func (o *CrmInfo) GetField9Caption() string`

GetField9Caption returns the Field9Caption field if non-nil, zero value otherwise.

### GetField9CaptionOk

`func (o *CrmInfo) GetField9CaptionOk() (*string, bool)`

GetField9CaptionOk returns a tuple with the Field9Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField9Caption

`func (o *CrmInfo) SetField9Caption(v string)`

SetField9Caption sets Field9Caption field to given value.

### HasField9Caption

`func (o *CrmInfo) HasField9Caption() bool`

HasField9Caption returns a boolean if a field has been set.

### GetField10Caption

`func (o *CrmInfo) GetField10Caption() string`

GetField10Caption returns the Field10Caption field if non-nil, zero value otherwise.

### GetField10CaptionOk

`func (o *CrmInfo) GetField10CaptionOk() (*string, bool)`

GetField10CaptionOk returns a tuple with the Field10Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField10Caption

`func (o *CrmInfo) SetField10Caption(v string)`

SetField10Caption sets Field10Caption field to given value.

### HasField10Caption

`func (o *CrmInfo) HasField10Caption() bool`

HasField10Caption returns a boolean if a field has been set.

### GetPrimaryRepCaption

`func (o *CrmInfo) GetPrimaryRepCaption() string`

GetPrimaryRepCaption returns the PrimaryRepCaption field if non-nil, zero value otherwise.

### GetPrimaryRepCaptionOk

`func (o *CrmInfo) GetPrimaryRepCaptionOk() (*string, bool)`

GetPrimaryRepCaptionOk returns a tuple with the PrimaryRepCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryRepCaption

`func (o *CrmInfo) SetPrimaryRepCaption(v string)`

SetPrimaryRepCaption sets PrimaryRepCaption field to given value.

### HasPrimaryRepCaption

`func (o *CrmInfo) HasPrimaryRepCaption() bool`

HasPrimaryRepCaption returns a boolean if a field has been set.

### GetSecondaryRepCaption

`func (o *CrmInfo) GetSecondaryRepCaption() string`

GetSecondaryRepCaption returns the SecondaryRepCaption field if non-nil, zero value otherwise.

### GetSecondaryRepCaptionOk

`func (o *CrmInfo) GetSecondaryRepCaptionOk() (*string, bool)`

GetSecondaryRepCaptionOk returns a tuple with the SecondaryRepCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryRepCaption

`func (o *CrmInfo) SetSecondaryRepCaption(v string)`

SetSecondaryRepCaption sets SecondaryRepCaption field to given value.

### HasSecondaryRepCaption

`func (o *CrmInfo) HasSecondaryRepCaption() bool`

HasSecondaryRepCaption returns a boolean if a field has been set.

### GetInfo

`func (o *CrmInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CrmInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CrmInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CrmInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


