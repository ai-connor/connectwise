# GLPath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Path** | Pointer to **string** |  Max length: 255; | [optional] 
**SqlServerName** | Pointer to **string** |  Max length: 255; | [optional] 
**DatabaseName** | Pointer to **string** |  Max length: 100; | [optional] 
**LastPaymentSync** | Pointer to **time.Time** |  | [optional] 
**LastPaymentSyncBy** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewGLPath

`func NewGLPath() *GLPath`

NewGLPath instantiates a new GLPath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGLPathWithDefaults

`func NewGLPathWithDefaults() *GLPath`

NewGLPathWithDefaults instantiates a new GLPath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GLPath) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GLPath) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GLPath) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GLPath) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *GLPath) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GLPath) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GLPath) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GLPath) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetPath

`func (o *GLPath) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GLPath) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GLPath) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GLPath) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSqlServerName

`func (o *GLPath) GetSqlServerName() string`

GetSqlServerName returns the SqlServerName field if non-nil, zero value otherwise.

### GetSqlServerNameOk

`func (o *GLPath) GetSqlServerNameOk() (*string, bool)`

GetSqlServerNameOk returns a tuple with the SqlServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlServerName

`func (o *GLPath) SetSqlServerName(v string)`

SetSqlServerName sets SqlServerName field to given value.

### HasSqlServerName

`func (o *GLPath) HasSqlServerName() bool`

HasSqlServerName returns a boolean if a field has been set.

### GetDatabaseName

`func (o *GLPath) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *GLPath) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *GLPath) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *GLPath) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### GetLastPaymentSync

`func (o *GLPath) GetLastPaymentSync() time.Time`

GetLastPaymentSync returns the LastPaymentSync field if non-nil, zero value otherwise.

### GetLastPaymentSyncOk

`func (o *GLPath) GetLastPaymentSyncOk() (*time.Time, bool)`

GetLastPaymentSyncOk returns a tuple with the LastPaymentSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentSync

`func (o *GLPath) SetLastPaymentSync(v time.Time)`

SetLastPaymentSync sets LastPaymentSync field to given value.

### HasLastPaymentSync

`func (o *GLPath) HasLastPaymentSync() bool`

HasLastPaymentSync returns a boolean if a field has been set.

### GetLastPaymentSyncBy

`func (o *GLPath) GetLastPaymentSyncBy() MemberReference`

GetLastPaymentSyncBy returns the LastPaymentSyncBy field if non-nil, zero value otherwise.

### GetLastPaymentSyncByOk

`func (o *GLPath) GetLastPaymentSyncByOk() (*MemberReference, bool)`

GetLastPaymentSyncByOk returns a tuple with the LastPaymentSyncBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentSyncBy

`func (o *GLPath) SetLastPaymentSyncBy(v MemberReference)`

SetLastPaymentSyncBy sets LastPaymentSyncBy field to given value.

### HasLastPaymentSyncBy

`func (o *GLPath) HasLastPaymentSyncBy() bool`

HasLastPaymentSyncBy returns a boolean if a field has been set.

### GetInfo

`func (o *GLPath) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *GLPath) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *GLPath) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *GLPath) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


