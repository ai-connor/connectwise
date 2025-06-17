# Imap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 200; | 
**ImapName** | **string** |  Max length: 40; | 
**ProcessedName** | **string** |  Max length: 40; | 
**FailedFolder** | **string** |  Max length: 40; | 
**Server** | **string** |  Max length: 200; | 
**UserName** | **string** |  Max length: 80; | 
**Password** | Pointer to **string** |  Max length: 80; | [optional] 
**Port** | **NullableInt32** |  | 
**SslFlag** | Pointer to **NullableBool** |  | [optional] 
**EmailConnector** | Pointer to [**EmailConnectorReference**](EmailConnectorReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewImap

`func NewImap(name string, imapName string, processedName string, failedFolder string, server string, userName string, port NullableInt32, ) *Imap`

NewImap instantiates a new Imap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImapWithDefaults

`func NewImapWithDefaults() *Imap`

NewImapWithDefaults instantiates a new Imap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Imap) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Imap) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Imap) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Imap) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Imap) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Imap) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Imap) SetName(v string)`

SetName sets Name field to given value.


### GetImapName

`func (o *Imap) GetImapName() string`

GetImapName returns the ImapName field if non-nil, zero value otherwise.

### GetImapNameOk

`func (o *Imap) GetImapNameOk() (*string, bool)`

GetImapNameOk returns a tuple with the ImapName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapName

`func (o *Imap) SetImapName(v string)`

SetImapName sets ImapName field to given value.


### GetProcessedName

`func (o *Imap) GetProcessedName() string`

GetProcessedName returns the ProcessedName field if non-nil, zero value otherwise.

### GetProcessedNameOk

`func (o *Imap) GetProcessedNameOk() (*string, bool)`

GetProcessedNameOk returns a tuple with the ProcessedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedName

`func (o *Imap) SetProcessedName(v string)`

SetProcessedName sets ProcessedName field to given value.


### GetFailedFolder

`func (o *Imap) GetFailedFolder() string`

GetFailedFolder returns the FailedFolder field if non-nil, zero value otherwise.

### GetFailedFolderOk

`func (o *Imap) GetFailedFolderOk() (*string, bool)`

GetFailedFolderOk returns a tuple with the FailedFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedFolder

`func (o *Imap) SetFailedFolder(v string)`

SetFailedFolder sets FailedFolder field to given value.


### GetServer

`func (o *Imap) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *Imap) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *Imap) SetServer(v string)`

SetServer sets Server field to given value.


### GetUserName

`func (o *Imap) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *Imap) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *Imap) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetPassword

`func (o *Imap) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Imap) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Imap) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *Imap) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *Imap) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Imap) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Imap) SetPort(v int32)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *Imap) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *Imap) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetSslFlag

`func (o *Imap) GetSslFlag() bool`

GetSslFlag returns the SslFlag field if non-nil, zero value otherwise.

### GetSslFlagOk

`func (o *Imap) GetSslFlagOk() (*bool, bool)`

GetSslFlagOk returns a tuple with the SslFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslFlag

`func (o *Imap) SetSslFlag(v bool)`

SetSslFlag sets SslFlag field to given value.

### HasSslFlag

`func (o *Imap) HasSslFlag() bool`

HasSslFlag returns a boolean if a field has been set.

### SetSslFlagNil

`func (o *Imap) SetSslFlagNil(b bool)`

 SetSslFlagNil sets the value for SslFlag to be an explicit nil

### UnsetSslFlag
`func (o *Imap) UnsetSslFlag()`

UnsetSslFlag ensures that no value is present for SslFlag, not even an explicit nil
### GetEmailConnector

`func (o *Imap) GetEmailConnector() EmailConnectorReference`

GetEmailConnector returns the EmailConnector field if non-nil, zero value otherwise.

### GetEmailConnectorOk

`func (o *Imap) GetEmailConnectorOk() (*EmailConnectorReference, bool)`

GetEmailConnectorOk returns a tuple with the EmailConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConnector

`func (o *Imap) SetEmailConnector(v EmailConnectorReference)`

SetEmailConnector sets EmailConnector field to given value.

### HasEmailConnector

`func (o *Imap) HasEmailConnector() bool`

HasEmailConnector returns a boolean if a field has been set.

### GetInfo

`func (o *Imap) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Imap) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Imap) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Imap) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


