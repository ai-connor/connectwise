# TicketMerge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MergeTicketIds** | **[]int32** |  | 
**Status** | [**ServiceStatusReference**](ServiceStatusReference.md) |  | 

## Methods

### NewTicketMerge

`func NewTicketMerge(mergeTicketIds []int32, status ServiceStatusReference, ) *TicketMerge`

NewTicketMerge instantiates a new TicketMerge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketMergeWithDefaults

`func NewTicketMergeWithDefaults() *TicketMerge`

NewTicketMergeWithDefaults instantiates a new TicketMerge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMergeTicketIds

`func (o *TicketMerge) GetMergeTicketIds() []int32`

GetMergeTicketIds returns the MergeTicketIds field if non-nil, zero value otherwise.

### GetMergeTicketIdsOk

`func (o *TicketMerge) GetMergeTicketIdsOk() (*[]int32, bool)`

GetMergeTicketIdsOk returns a tuple with the MergeTicketIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeTicketIds

`func (o *TicketMerge) SetMergeTicketIds(v []int32)`

SetMergeTicketIds sets MergeTicketIds field to given value.


### GetStatus

`func (o *TicketMerge) GetStatus() ServiceStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TicketMerge) GetStatusOk() (*ServiceStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TicketMerge) SetStatus(v ServiceStatusReference)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


