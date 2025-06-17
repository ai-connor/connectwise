# Info

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** |  | [optional] 
**IsCloud** | Pointer to **bool** |  | [optional] 
**ServerTimeZone** | Pointer to **string** |  | [optional] 
**LicenseBits** | Pointer to [**[]LicenseBit**](LicenseBit.md) |  | [optional] 
**CloudRegion** | Pointer to **string** |  | [optional] 
**MaxWorkFlowRecordsAllowed** | Pointer to **int32** |  | [optional] 

## Methods

### NewInfo

`func NewInfo() *Info`

NewInfo instantiates a new Info object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfoWithDefaults

`func NewInfoWithDefaults() *Info`

NewInfoWithDefaults instantiates a new Info object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *Info) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Info) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Info) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Info) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIsCloud

`func (o *Info) GetIsCloud() bool`

GetIsCloud returns the IsCloud field if non-nil, zero value otherwise.

### GetIsCloudOk

`func (o *Info) GetIsCloudOk() (*bool, bool)`

GetIsCloudOk returns a tuple with the IsCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloud

`func (o *Info) SetIsCloud(v bool)`

SetIsCloud sets IsCloud field to given value.

### HasIsCloud

`func (o *Info) HasIsCloud() bool`

HasIsCloud returns a boolean if a field has been set.

### GetServerTimeZone

`func (o *Info) GetServerTimeZone() string`

GetServerTimeZone returns the ServerTimeZone field if non-nil, zero value otherwise.

### GetServerTimeZoneOk

`func (o *Info) GetServerTimeZoneOk() (*string, bool)`

GetServerTimeZoneOk returns a tuple with the ServerTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTimeZone

`func (o *Info) SetServerTimeZone(v string)`

SetServerTimeZone sets ServerTimeZone field to given value.

### HasServerTimeZone

`func (o *Info) HasServerTimeZone() bool`

HasServerTimeZone returns a boolean if a field has been set.

### GetLicenseBits

`func (o *Info) GetLicenseBits() []LicenseBit`

GetLicenseBits returns the LicenseBits field if non-nil, zero value otherwise.

### GetLicenseBitsOk

`func (o *Info) GetLicenseBitsOk() (*[]LicenseBit, bool)`

GetLicenseBitsOk returns a tuple with the LicenseBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseBits

`func (o *Info) SetLicenseBits(v []LicenseBit)`

SetLicenseBits sets LicenseBits field to given value.

### HasLicenseBits

`func (o *Info) HasLicenseBits() bool`

HasLicenseBits returns a boolean if a field has been set.

### GetCloudRegion

`func (o *Info) GetCloudRegion() string`

GetCloudRegion returns the CloudRegion field if non-nil, zero value otherwise.

### GetCloudRegionOk

`func (o *Info) GetCloudRegionOk() (*string, bool)`

GetCloudRegionOk returns a tuple with the CloudRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudRegion

`func (o *Info) SetCloudRegion(v string)`

SetCloudRegion sets CloudRegion field to given value.

### HasCloudRegion

`func (o *Info) HasCloudRegion() bool`

HasCloudRegion returns a boolean if a field has been set.

### GetMaxWorkFlowRecordsAllowed

`func (o *Info) GetMaxWorkFlowRecordsAllowed() int32`

GetMaxWorkFlowRecordsAllowed returns the MaxWorkFlowRecordsAllowed field if non-nil, zero value otherwise.

### GetMaxWorkFlowRecordsAllowedOk

`func (o *Info) GetMaxWorkFlowRecordsAllowedOk() (*int32, bool)`

GetMaxWorkFlowRecordsAllowedOk returns a tuple with the MaxWorkFlowRecordsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWorkFlowRecordsAllowed

`func (o *Info) SetMaxWorkFlowRecordsAllowed(v int32)`

SetMaxWorkFlowRecordsAllowed sets MaxWorkFlowRecordsAllowed field to given value.

### HasMaxWorkFlowRecordsAllowed

`func (o *Info) HasMaxWorkFlowRecordsAllowed() bool`

HasMaxWorkFlowRecordsAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


