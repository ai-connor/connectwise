# M365Contact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**UserPrincipalName** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ContactRecId** | Pointer to **int32** |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 
**M365ContactId** | Pointer to **string** |  | [optional] 
**Department** | Pointer to **string** |  | [optional] 
**EmployeeType** | Pointer to **string** |  | [optional] 
**ManagerId** | Pointer to **string** |  | [optional] 
**ProxyAddresses** | Pointer to **string** |  | [optional] 
**ProxyAddressesPlain** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to **string** |  | [optional] 
**DirectoryRoles** | Pointer to **string** |  | [optional] 
**AssignedLicenses** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewM365Contact

`func NewM365Contact() *M365Contact`

NewM365Contact instantiates a new M365Contact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewM365ContactWithDefaults

`func NewM365ContactWithDefaults() *M365Contact`

NewM365ContactWithDefaults instantiates a new M365Contact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *M365Contact) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *M365Contact) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *M365Contact) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *M365Contact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserPrincipalName

`func (o *M365Contact) GetUserPrincipalName() string`

GetUserPrincipalName returns the UserPrincipalName field if non-nil, zero value otherwise.

### GetUserPrincipalNameOk

`func (o *M365Contact) GetUserPrincipalNameOk() (*string, bool)`

GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrincipalName

`func (o *M365Contact) SetUserPrincipalName(v string)`

SetUserPrincipalName sets UserPrincipalName field to given value.

### HasUserPrincipalName

`func (o *M365Contact) HasUserPrincipalName() bool`

HasUserPrincipalName returns a boolean if a field has been set.

### GetDisplayName

`func (o *M365Contact) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *M365Contact) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *M365Contact) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *M365Contact) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetContactRecId

`func (o *M365Contact) GetContactRecId() int32`

GetContactRecId returns the ContactRecId field if non-nil, zero value otherwise.

### GetContactRecIdOk

`func (o *M365Contact) GetContactRecIdOk() (*int32, bool)`

GetContactRecIdOk returns a tuple with the ContactRecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactRecId

`func (o *M365Contact) SetContactRecId(v int32)`

SetContactRecId sets ContactRecId field to given value.

### HasContactRecId

`func (o *M365Contact) HasContactRecId() bool`

HasContactRecId returns a boolean if a field has been set.

### GetTenantId

`func (o *M365Contact) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *M365Contact) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *M365Contact) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *M365Contact) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetM365ContactId

`func (o *M365Contact) GetM365ContactId() string`

GetM365ContactId returns the M365ContactId field if non-nil, zero value otherwise.

### GetM365ContactIdOk

`func (o *M365Contact) GetM365ContactIdOk() (*string, bool)`

GetM365ContactIdOk returns a tuple with the M365ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetM365ContactId

`func (o *M365Contact) SetM365ContactId(v string)`

SetM365ContactId sets M365ContactId field to given value.

### HasM365ContactId

`func (o *M365Contact) HasM365ContactId() bool`

HasM365ContactId returns a boolean if a field has been set.

### GetDepartment

`func (o *M365Contact) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *M365Contact) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *M365Contact) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *M365Contact) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetEmployeeType

`func (o *M365Contact) GetEmployeeType() string`

GetEmployeeType returns the EmployeeType field if non-nil, zero value otherwise.

### GetEmployeeTypeOk

`func (o *M365Contact) GetEmployeeTypeOk() (*string, bool)`

GetEmployeeTypeOk returns a tuple with the EmployeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeType

`func (o *M365Contact) SetEmployeeType(v string)`

SetEmployeeType sets EmployeeType field to given value.

### HasEmployeeType

`func (o *M365Contact) HasEmployeeType() bool`

HasEmployeeType returns a boolean if a field has been set.

### GetManagerId

`func (o *M365Contact) GetManagerId() string`

GetManagerId returns the ManagerId field if non-nil, zero value otherwise.

### GetManagerIdOk

`func (o *M365Contact) GetManagerIdOk() (*string, bool)`

GetManagerIdOk returns a tuple with the ManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerId

`func (o *M365Contact) SetManagerId(v string)`

SetManagerId sets ManagerId field to given value.

### HasManagerId

`func (o *M365Contact) HasManagerId() bool`

HasManagerId returns a boolean if a field has been set.

### GetProxyAddresses

`func (o *M365Contact) GetProxyAddresses() string`

GetProxyAddresses returns the ProxyAddresses field if non-nil, zero value otherwise.

### GetProxyAddressesOk

`func (o *M365Contact) GetProxyAddressesOk() (*string, bool)`

GetProxyAddressesOk returns a tuple with the ProxyAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAddresses

`func (o *M365Contact) SetProxyAddresses(v string)`

SetProxyAddresses sets ProxyAddresses field to given value.

### HasProxyAddresses

`func (o *M365Contact) HasProxyAddresses() bool`

HasProxyAddresses returns a boolean if a field has been set.

### GetProxyAddressesPlain

`func (o *M365Contact) GetProxyAddressesPlain() string`

GetProxyAddressesPlain returns the ProxyAddressesPlain field if non-nil, zero value otherwise.

### GetProxyAddressesPlainOk

`func (o *M365Contact) GetProxyAddressesPlainOk() (*string, bool)`

GetProxyAddressesPlainOk returns a tuple with the ProxyAddressesPlain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAddressesPlain

`func (o *M365Contact) SetProxyAddressesPlain(v string)`

SetProxyAddressesPlain sets ProxyAddressesPlain field to given value.

### HasProxyAddressesPlain

`func (o *M365Contact) HasProxyAddressesPlain() bool`

HasProxyAddressesPlain returns a boolean if a field has been set.

### GetGroups

`func (o *M365Contact) GetGroups() string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *M365Contact) GetGroupsOk() (*string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *M365Contact) SetGroups(v string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *M365Contact) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetDirectoryRoles

`func (o *M365Contact) GetDirectoryRoles() string`

GetDirectoryRoles returns the DirectoryRoles field if non-nil, zero value otherwise.

### GetDirectoryRolesOk

`func (o *M365Contact) GetDirectoryRolesOk() (*string, bool)`

GetDirectoryRolesOk returns a tuple with the DirectoryRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryRoles

`func (o *M365Contact) SetDirectoryRoles(v string)`

SetDirectoryRoles sets DirectoryRoles field to given value.

### HasDirectoryRoles

`func (o *M365Contact) HasDirectoryRoles() bool`

HasDirectoryRoles returns a boolean if a field has been set.

### GetAssignedLicenses

`func (o *M365Contact) GetAssignedLicenses() string`

GetAssignedLicenses returns the AssignedLicenses field if non-nil, zero value otherwise.

### GetAssignedLicensesOk

`func (o *M365Contact) GetAssignedLicensesOk() (*string, bool)`

GetAssignedLicensesOk returns a tuple with the AssignedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedLicenses

`func (o *M365Contact) SetAssignedLicenses(v string)`

SetAssignedLicenses sets AssignedLicenses field to given value.

### HasAssignedLicenses

`func (o *M365Contact) HasAssignedLicenses() bool`

HasAssignedLicenses returns a boolean if a field has been set.

### GetInfo

`func (o *M365Contact) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *M365Contact) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *M365Contact) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *M365Contact) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


