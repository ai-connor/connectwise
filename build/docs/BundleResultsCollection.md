# BundleResultsCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]BundleResult**](BundleResult.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBundleResultsCollection

`func NewBundleResultsCollection() *BundleResultsCollection`

NewBundleResultsCollection instantiates a new BundleResultsCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleResultsCollectionWithDefaults

`func NewBundleResultsCollectionWithDefaults() *BundleResultsCollection`

NewBundleResultsCollectionWithDefaults instantiates a new BundleResultsCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *BundleResultsCollection) GetResults() []BundleResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BundleResultsCollection) GetResultsOk() (*[]BundleResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BundleResultsCollection) SetResults(v []BundleResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *BundleResultsCollection) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetInfo

`func (o *BundleResultsCollection) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BundleResultsCollection) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BundleResultsCollection) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BundleResultsCollection) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


