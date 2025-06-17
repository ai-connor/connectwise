# ImapInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**EmailConnector** | Pointer to [**EmailConnectorReference**](EmailConnectorReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewImapInfo

`func NewImapInfo() *ImapInfo`

NewImapInfo instantiates a new ImapInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImapInfoWithDefaults

`func NewImapInfoWithDefaults() *ImapInfo`

NewImapInfoWithDefaults instantiates a new ImapInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImapInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImapInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImapInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ImapInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ImapInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImapInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImapInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImapInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmailConnector

`func (o *ImapInfo) GetEmailConnector() EmailConnectorReference`

GetEmailConnector returns the EmailConnector field if non-nil, zero value otherwise.

### GetEmailConnectorOk

`func (o *ImapInfo) GetEmailConnectorOk() (*EmailConnectorReference, bool)`

GetEmailConnectorOk returns a tuple with the EmailConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConnector

`func (o *ImapInfo) SetEmailConnector(v EmailConnectorReference)`

SetEmailConnector sets EmailConnector field to given value.

### HasEmailConnector

`func (o *ImapInfo) HasEmailConnector() bool`

HasEmailConnector returns a boolean if a field has been set.

### GetInfo

`func (o *ImapInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ImapInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ImapInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ImapInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


