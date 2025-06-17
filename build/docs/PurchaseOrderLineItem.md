# PurchaseOrderLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**BackorderedFlag** | Pointer to **NullableBool** |  | [optional] 
**CanceledBy** | Pointer to **string** |  | [optional] 
**CanceledFlag** | Pointer to **NullableBool** |  | [optional] 
**CanceledReason** | Pointer to **string** |  Max length: 100; | [optional] 
**ClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**DateCanceled** | Pointer to **time.Time** |  | [optional] 
**DateCanceledUtc** | Pointer to **time.Time** |  | [optional] 
**Description** | **string** |  Max length: 6000; | 
**DisplayInternalNotesFlag** | Pointer to **NullableBool** |  | [optional] 
**ExpectedShipDate** | Pointer to **time.Time** |  | [optional] 
**InternalNotes** | Pointer to **string** |  Max length: 1000; | [optional] 
**LineNumber** | **NullableInt32** |  | 
**PackingSlip** | Pointer to **string** |  Max length: 50; | [optional] 
**Product** | [**IvItemReference**](IvItemReference.md) |  | 
**PurchaseOrderId** | Pointer to **NullableInt32** |  | [optional] 
**PurchaseOrderNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**Quantity** | **NullableFloat64** |  | 
**ReceivedQuantity** | Pointer to **NullableInt32** |  | [optional] 
**SerialNumbers** | Pointer to **string** |  | [optional] 
**ShipDate** | Pointer to **time.Time** |  | [optional] 
**ShipmentMethod** | Pointer to [**ShipmentMethodReference**](ShipmentMethodReference.md) |  | [optional] 
**Tax** | Pointer to **NullableFloat64** |  | [optional] 
**TrackingNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**UnitCost** | Pointer to **NullableFloat64** |  | [optional] 
**UnitOfMeasure** | [**UnitOfMeasureReference**](UnitOfMeasureReference.md) |  | 
**VendorOrderNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**VendorSku** | Pointer to **string** |  Max length: 50; | [optional] 
**Warehouse** | Pointer to [**WarehouseReference**](WarehouseReference.md) |  | [optional] 
**WarehouseBin** | Pointer to [**WarehouseBinReference**](WarehouseBinReference.md) |  | [optional] 
**ShipSet** | Pointer to **string** |  Max length: 10; | [optional] 
**DateReceived** | Pointer to **time.Time** |  | [optional] 
**ReceivedStatus** | Pointer to **NullableString** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**ExtCost** | Pointer to **float64** |  | [optional] 
**ExpectedArrivalDate** | Pointer to **time.Time** |  | [optional] 
**IsDetachAvailable** | Pointer to **bool** |  | [optional] 
**BatchedFlag** | Pointer to **bool** |  | [optional] 
**UnbatchedRecId** | Pointer to **NullableInt32** |  | [optional] 
**SalesOrder** | Pointer to [**[]SalesOrderReference**](SalesOrderReference.md) |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewPurchaseOrderLineItem

`func NewPurchaseOrderLineItem(description string, lineNumber NullableInt32, product IvItemReference, quantity NullableFloat64, unitOfMeasure UnitOfMeasureReference, ) *PurchaseOrderLineItem`

NewPurchaseOrderLineItem instantiates a new PurchaseOrderLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseOrderLineItemWithDefaults

`func NewPurchaseOrderLineItemWithDefaults() *PurchaseOrderLineItem`

NewPurchaseOrderLineItemWithDefaults instantiates a new PurchaseOrderLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PurchaseOrderLineItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PurchaseOrderLineItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PurchaseOrderLineItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PurchaseOrderLineItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBackorderedFlag

`func (o *PurchaseOrderLineItem) GetBackorderedFlag() bool`

GetBackorderedFlag returns the BackorderedFlag field if non-nil, zero value otherwise.

### GetBackorderedFlagOk

`func (o *PurchaseOrderLineItem) GetBackorderedFlagOk() (*bool, bool)`

GetBackorderedFlagOk returns a tuple with the BackorderedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackorderedFlag

`func (o *PurchaseOrderLineItem) SetBackorderedFlag(v bool)`

SetBackorderedFlag sets BackorderedFlag field to given value.

### HasBackorderedFlag

`func (o *PurchaseOrderLineItem) HasBackorderedFlag() bool`

HasBackorderedFlag returns a boolean if a field has been set.

### SetBackorderedFlagNil

`func (o *PurchaseOrderLineItem) SetBackorderedFlagNil(b bool)`

 SetBackorderedFlagNil sets the value for BackorderedFlag to be an explicit nil

### UnsetBackorderedFlag
`func (o *PurchaseOrderLineItem) UnsetBackorderedFlag()`

UnsetBackorderedFlag ensures that no value is present for BackorderedFlag, not even an explicit nil
### GetCanceledBy

`func (o *PurchaseOrderLineItem) GetCanceledBy() string`

GetCanceledBy returns the CanceledBy field if non-nil, zero value otherwise.

### GetCanceledByOk

`func (o *PurchaseOrderLineItem) GetCanceledByOk() (*string, bool)`

GetCanceledByOk returns a tuple with the CanceledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledBy

`func (o *PurchaseOrderLineItem) SetCanceledBy(v string)`

SetCanceledBy sets CanceledBy field to given value.

### HasCanceledBy

`func (o *PurchaseOrderLineItem) HasCanceledBy() bool`

HasCanceledBy returns a boolean if a field has been set.

### GetCanceledFlag

`func (o *PurchaseOrderLineItem) GetCanceledFlag() bool`

GetCanceledFlag returns the CanceledFlag field if non-nil, zero value otherwise.

### GetCanceledFlagOk

`func (o *PurchaseOrderLineItem) GetCanceledFlagOk() (*bool, bool)`

GetCanceledFlagOk returns a tuple with the CanceledFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledFlag

`func (o *PurchaseOrderLineItem) SetCanceledFlag(v bool)`

SetCanceledFlag sets CanceledFlag field to given value.

### HasCanceledFlag

`func (o *PurchaseOrderLineItem) HasCanceledFlag() bool`

HasCanceledFlag returns a boolean if a field has been set.

### SetCanceledFlagNil

`func (o *PurchaseOrderLineItem) SetCanceledFlagNil(b bool)`

 SetCanceledFlagNil sets the value for CanceledFlag to be an explicit nil

### UnsetCanceledFlag
`func (o *PurchaseOrderLineItem) UnsetCanceledFlag()`

UnsetCanceledFlag ensures that no value is present for CanceledFlag, not even an explicit nil
### GetCanceledReason

`func (o *PurchaseOrderLineItem) GetCanceledReason() string`

GetCanceledReason returns the CanceledReason field if non-nil, zero value otherwise.

### GetCanceledReasonOk

`func (o *PurchaseOrderLineItem) GetCanceledReasonOk() (*string, bool)`

GetCanceledReasonOk returns a tuple with the CanceledReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledReason

`func (o *PurchaseOrderLineItem) SetCanceledReason(v string)`

SetCanceledReason sets CanceledReason field to given value.

### HasCanceledReason

`func (o *PurchaseOrderLineItem) HasCanceledReason() bool`

HasCanceledReason returns a boolean if a field has been set.

### GetClosedFlag

`func (o *PurchaseOrderLineItem) GetClosedFlag() bool`

GetClosedFlag returns the ClosedFlag field if non-nil, zero value otherwise.

### GetClosedFlagOk

`func (o *PurchaseOrderLineItem) GetClosedFlagOk() (*bool, bool)`

GetClosedFlagOk returns a tuple with the ClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedFlag

`func (o *PurchaseOrderLineItem) SetClosedFlag(v bool)`

SetClosedFlag sets ClosedFlag field to given value.

### HasClosedFlag

`func (o *PurchaseOrderLineItem) HasClosedFlag() bool`

HasClosedFlag returns a boolean if a field has been set.

### SetClosedFlagNil

`func (o *PurchaseOrderLineItem) SetClosedFlagNil(b bool)`

 SetClosedFlagNil sets the value for ClosedFlag to be an explicit nil

### UnsetClosedFlag
`func (o *PurchaseOrderLineItem) UnsetClosedFlag()`

UnsetClosedFlag ensures that no value is present for ClosedFlag, not even an explicit nil
### GetDateCanceled

`func (o *PurchaseOrderLineItem) GetDateCanceled() time.Time`

GetDateCanceled returns the DateCanceled field if non-nil, zero value otherwise.

### GetDateCanceledOk

`func (o *PurchaseOrderLineItem) GetDateCanceledOk() (*time.Time, bool)`

GetDateCanceledOk returns a tuple with the DateCanceled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCanceled

`func (o *PurchaseOrderLineItem) SetDateCanceled(v time.Time)`

SetDateCanceled sets DateCanceled field to given value.

### HasDateCanceled

`func (o *PurchaseOrderLineItem) HasDateCanceled() bool`

HasDateCanceled returns a boolean if a field has been set.

### GetDateCanceledUtc

`func (o *PurchaseOrderLineItem) GetDateCanceledUtc() time.Time`

GetDateCanceledUtc returns the DateCanceledUtc field if non-nil, zero value otherwise.

### GetDateCanceledUtcOk

`func (o *PurchaseOrderLineItem) GetDateCanceledUtcOk() (*time.Time, bool)`

GetDateCanceledUtcOk returns a tuple with the DateCanceledUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCanceledUtc

`func (o *PurchaseOrderLineItem) SetDateCanceledUtc(v time.Time)`

SetDateCanceledUtc sets DateCanceledUtc field to given value.

### HasDateCanceledUtc

`func (o *PurchaseOrderLineItem) HasDateCanceledUtc() bool`

HasDateCanceledUtc returns a boolean if a field has been set.

### GetDescription

`func (o *PurchaseOrderLineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PurchaseOrderLineItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PurchaseOrderLineItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisplayInternalNotesFlag

`func (o *PurchaseOrderLineItem) GetDisplayInternalNotesFlag() bool`

GetDisplayInternalNotesFlag returns the DisplayInternalNotesFlag field if non-nil, zero value otherwise.

### GetDisplayInternalNotesFlagOk

`func (o *PurchaseOrderLineItem) GetDisplayInternalNotesFlagOk() (*bool, bool)`

GetDisplayInternalNotesFlagOk returns a tuple with the DisplayInternalNotesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayInternalNotesFlag

`func (o *PurchaseOrderLineItem) SetDisplayInternalNotesFlag(v bool)`

SetDisplayInternalNotesFlag sets DisplayInternalNotesFlag field to given value.

### HasDisplayInternalNotesFlag

`func (o *PurchaseOrderLineItem) HasDisplayInternalNotesFlag() bool`

HasDisplayInternalNotesFlag returns a boolean if a field has been set.

### SetDisplayInternalNotesFlagNil

`func (o *PurchaseOrderLineItem) SetDisplayInternalNotesFlagNil(b bool)`

 SetDisplayInternalNotesFlagNil sets the value for DisplayInternalNotesFlag to be an explicit nil

### UnsetDisplayInternalNotesFlag
`func (o *PurchaseOrderLineItem) UnsetDisplayInternalNotesFlag()`

UnsetDisplayInternalNotesFlag ensures that no value is present for DisplayInternalNotesFlag, not even an explicit nil
### GetExpectedShipDate

`func (o *PurchaseOrderLineItem) GetExpectedShipDate() time.Time`

GetExpectedShipDate returns the ExpectedShipDate field if non-nil, zero value otherwise.

### GetExpectedShipDateOk

`func (o *PurchaseOrderLineItem) GetExpectedShipDateOk() (*time.Time, bool)`

GetExpectedShipDateOk returns a tuple with the ExpectedShipDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedShipDate

`func (o *PurchaseOrderLineItem) SetExpectedShipDate(v time.Time)`

SetExpectedShipDate sets ExpectedShipDate field to given value.

### HasExpectedShipDate

`func (o *PurchaseOrderLineItem) HasExpectedShipDate() bool`

HasExpectedShipDate returns a boolean if a field has been set.

### GetInternalNotes

`func (o *PurchaseOrderLineItem) GetInternalNotes() string`

GetInternalNotes returns the InternalNotes field if non-nil, zero value otherwise.

### GetInternalNotesOk

`func (o *PurchaseOrderLineItem) GetInternalNotesOk() (*string, bool)`

GetInternalNotesOk returns a tuple with the InternalNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNotes

`func (o *PurchaseOrderLineItem) SetInternalNotes(v string)`

SetInternalNotes sets InternalNotes field to given value.

### HasInternalNotes

`func (o *PurchaseOrderLineItem) HasInternalNotes() bool`

HasInternalNotes returns a boolean if a field has been set.

### GetLineNumber

`func (o *PurchaseOrderLineItem) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *PurchaseOrderLineItem) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *PurchaseOrderLineItem) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.


### SetLineNumberNil

`func (o *PurchaseOrderLineItem) SetLineNumberNil(b bool)`

 SetLineNumberNil sets the value for LineNumber to be an explicit nil

### UnsetLineNumber
`func (o *PurchaseOrderLineItem) UnsetLineNumber()`

UnsetLineNumber ensures that no value is present for LineNumber, not even an explicit nil
### GetPackingSlip

`func (o *PurchaseOrderLineItem) GetPackingSlip() string`

GetPackingSlip returns the PackingSlip field if non-nil, zero value otherwise.

### GetPackingSlipOk

`func (o *PurchaseOrderLineItem) GetPackingSlipOk() (*string, bool)`

GetPackingSlipOk returns a tuple with the PackingSlip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackingSlip

`func (o *PurchaseOrderLineItem) SetPackingSlip(v string)`

SetPackingSlip sets PackingSlip field to given value.

### HasPackingSlip

`func (o *PurchaseOrderLineItem) HasPackingSlip() bool`

HasPackingSlip returns a boolean if a field has been set.

### GetProduct

`func (o *PurchaseOrderLineItem) GetProduct() IvItemReference`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PurchaseOrderLineItem) GetProductOk() (*IvItemReference, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PurchaseOrderLineItem) SetProduct(v IvItemReference)`

SetProduct sets Product field to given value.


### GetPurchaseOrderId

`func (o *PurchaseOrderLineItem) GetPurchaseOrderId() int32`

GetPurchaseOrderId returns the PurchaseOrderId field if non-nil, zero value otherwise.

### GetPurchaseOrderIdOk

`func (o *PurchaseOrderLineItem) GetPurchaseOrderIdOk() (*int32, bool)`

GetPurchaseOrderIdOk returns a tuple with the PurchaseOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderId

`func (o *PurchaseOrderLineItem) SetPurchaseOrderId(v int32)`

SetPurchaseOrderId sets PurchaseOrderId field to given value.

### HasPurchaseOrderId

`func (o *PurchaseOrderLineItem) HasPurchaseOrderId() bool`

HasPurchaseOrderId returns a boolean if a field has been set.

### SetPurchaseOrderIdNil

`func (o *PurchaseOrderLineItem) SetPurchaseOrderIdNil(b bool)`

 SetPurchaseOrderIdNil sets the value for PurchaseOrderId to be an explicit nil

### UnsetPurchaseOrderId
`func (o *PurchaseOrderLineItem) UnsetPurchaseOrderId()`

UnsetPurchaseOrderId ensures that no value is present for PurchaseOrderId, not even an explicit nil
### GetPurchaseOrderNumber

`func (o *PurchaseOrderLineItem) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *PurchaseOrderLineItem) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *PurchaseOrderLineItem) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *PurchaseOrderLineItem) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetQuantity

`func (o *PurchaseOrderLineItem) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *PurchaseOrderLineItem) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *PurchaseOrderLineItem) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.


### SetQuantityNil

`func (o *PurchaseOrderLineItem) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *PurchaseOrderLineItem) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetReceivedQuantity

`func (o *PurchaseOrderLineItem) GetReceivedQuantity() int32`

GetReceivedQuantity returns the ReceivedQuantity field if non-nil, zero value otherwise.

### GetReceivedQuantityOk

`func (o *PurchaseOrderLineItem) GetReceivedQuantityOk() (*int32, bool)`

GetReceivedQuantityOk returns a tuple with the ReceivedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedQuantity

`func (o *PurchaseOrderLineItem) SetReceivedQuantity(v int32)`

SetReceivedQuantity sets ReceivedQuantity field to given value.

### HasReceivedQuantity

`func (o *PurchaseOrderLineItem) HasReceivedQuantity() bool`

HasReceivedQuantity returns a boolean if a field has been set.

### SetReceivedQuantityNil

`func (o *PurchaseOrderLineItem) SetReceivedQuantityNil(b bool)`

 SetReceivedQuantityNil sets the value for ReceivedQuantity to be an explicit nil

### UnsetReceivedQuantity
`func (o *PurchaseOrderLineItem) UnsetReceivedQuantity()`

UnsetReceivedQuantity ensures that no value is present for ReceivedQuantity, not even an explicit nil
### GetSerialNumbers

`func (o *PurchaseOrderLineItem) GetSerialNumbers() string`

GetSerialNumbers returns the SerialNumbers field if non-nil, zero value otherwise.

### GetSerialNumbersOk

`func (o *PurchaseOrderLineItem) GetSerialNumbersOk() (*string, bool)`

GetSerialNumbersOk returns a tuple with the SerialNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumbers

`func (o *PurchaseOrderLineItem) SetSerialNumbers(v string)`

SetSerialNumbers sets SerialNumbers field to given value.

### HasSerialNumbers

`func (o *PurchaseOrderLineItem) HasSerialNumbers() bool`

HasSerialNumbers returns a boolean if a field has been set.

### GetShipDate

`func (o *PurchaseOrderLineItem) GetShipDate() time.Time`

GetShipDate returns the ShipDate field if non-nil, zero value otherwise.

### GetShipDateOk

`func (o *PurchaseOrderLineItem) GetShipDateOk() (*time.Time, bool)`

GetShipDateOk returns a tuple with the ShipDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipDate

`func (o *PurchaseOrderLineItem) SetShipDate(v time.Time)`

SetShipDate sets ShipDate field to given value.

### HasShipDate

`func (o *PurchaseOrderLineItem) HasShipDate() bool`

HasShipDate returns a boolean if a field has been set.

### GetShipmentMethod

`func (o *PurchaseOrderLineItem) GetShipmentMethod() ShipmentMethodReference`

GetShipmentMethod returns the ShipmentMethod field if non-nil, zero value otherwise.

### GetShipmentMethodOk

`func (o *PurchaseOrderLineItem) GetShipmentMethodOk() (*ShipmentMethodReference, bool)`

GetShipmentMethodOk returns a tuple with the ShipmentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentMethod

`func (o *PurchaseOrderLineItem) SetShipmentMethod(v ShipmentMethodReference)`

SetShipmentMethod sets ShipmentMethod field to given value.

### HasShipmentMethod

`func (o *PurchaseOrderLineItem) HasShipmentMethod() bool`

HasShipmentMethod returns a boolean if a field has been set.

### GetTax

`func (o *PurchaseOrderLineItem) GetTax() float64`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *PurchaseOrderLineItem) GetTaxOk() (*float64, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *PurchaseOrderLineItem) SetTax(v float64)`

SetTax sets Tax field to given value.

### HasTax

`func (o *PurchaseOrderLineItem) HasTax() bool`

HasTax returns a boolean if a field has been set.

### SetTaxNil

`func (o *PurchaseOrderLineItem) SetTaxNil(b bool)`

 SetTaxNil sets the value for Tax to be an explicit nil

### UnsetTax
`func (o *PurchaseOrderLineItem) UnsetTax()`

UnsetTax ensures that no value is present for Tax, not even an explicit nil
### GetTrackingNumber

`func (o *PurchaseOrderLineItem) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *PurchaseOrderLineItem) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *PurchaseOrderLineItem) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *PurchaseOrderLineItem) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetUnitCost

`func (o *PurchaseOrderLineItem) GetUnitCost() float64`

GetUnitCost returns the UnitCost field if non-nil, zero value otherwise.

### GetUnitCostOk

`func (o *PurchaseOrderLineItem) GetUnitCostOk() (*float64, bool)`

GetUnitCostOk returns a tuple with the UnitCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCost

`func (o *PurchaseOrderLineItem) SetUnitCost(v float64)`

SetUnitCost sets UnitCost field to given value.

### HasUnitCost

`func (o *PurchaseOrderLineItem) HasUnitCost() bool`

HasUnitCost returns a boolean if a field has been set.

### SetUnitCostNil

`func (o *PurchaseOrderLineItem) SetUnitCostNil(b bool)`

 SetUnitCostNil sets the value for UnitCost to be an explicit nil

### UnsetUnitCost
`func (o *PurchaseOrderLineItem) UnsetUnitCost()`

UnsetUnitCost ensures that no value is present for UnitCost, not even an explicit nil
### GetUnitOfMeasure

`func (o *PurchaseOrderLineItem) GetUnitOfMeasure() UnitOfMeasureReference`

GetUnitOfMeasure returns the UnitOfMeasure field if non-nil, zero value otherwise.

### GetUnitOfMeasureOk

`func (o *PurchaseOrderLineItem) GetUnitOfMeasureOk() (*UnitOfMeasureReference, bool)`

GetUnitOfMeasureOk returns a tuple with the UnitOfMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitOfMeasure

`func (o *PurchaseOrderLineItem) SetUnitOfMeasure(v UnitOfMeasureReference)`

SetUnitOfMeasure sets UnitOfMeasure field to given value.


### GetVendorOrderNumber

`func (o *PurchaseOrderLineItem) GetVendorOrderNumber() string`

GetVendorOrderNumber returns the VendorOrderNumber field if non-nil, zero value otherwise.

### GetVendorOrderNumberOk

`func (o *PurchaseOrderLineItem) GetVendorOrderNumberOk() (*string, bool)`

GetVendorOrderNumberOk returns a tuple with the VendorOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorOrderNumber

`func (o *PurchaseOrderLineItem) SetVendorOrderNumber(v string)`

SetVendorOrderNumber sets VendorOrderNumber field to given value.

### HasVendorOrderNumber

`func (o *PurchaseOrderLineItem) HasVendorOrderNumber() bool`

HasVendorOrderNumber returns a boolean if a field has been set.

### GetVendorSku

`func (o *PurchaseOrderLineItem) GetVendorSku() string`

GetVendorSku returns the VendorSku field if non-nil, zero value otherwise.

### GetVendorSkuOk

`func (o *PurchaseOrderLineItem) GetVendorSkuOk() (*string, bool)`

GetVendorSkuOk returns a tuple with the VendorSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSku

`func (o *PurchaseOrderLineItem) SetVendorSku(v string)`

SetVendorSku sets VendorSku field to given value.

### HasVendorSku

`func (o *PurchaseOrderLineItem) HasVendorSku() bool`

HasVendorSku returns a boolean if a field has been set.

### GetWarehouse

`func (o *PurchaseOrderLineItem) GetWarehouse() WarehouseReference`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *PurchaseOrderLineItem) GetWarehouseOk() (*WarehouseReference, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *PurchaseOrderLineItem) SetWarehouse(v WarehouseReference)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *PurchaseOrderLineItem) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### GetWarehouseBin

`func (o *PurchaseOrderLineItem) GetWarehouseBin() WarehouseBinReference`

GetWarehouseBin returns the WarehouseBin field if non-nil, zero value otherwise.

### GetWarehouseBinOk

`func (o *PurchaseOrderLineItem) GetWarehouseBinOk() (*WarehouseBinReference, bool)`

GetWarehouseBinOk returns a tuple with the WarehouseBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseBin

`func (o *PurchaseOrderLineItem) SetWarehouseBin(v WarehouseBinReference)`

SetWarehouseBin sets WarehouseBin field to given value.

### HasWarehouseBin

`func (o *PurchaseOrderLineItem) HasWarehouseBin() bool`

HasWarehouseBin returns a boolean if a field has been set.

### GetShipSet

`func (o *PurchaseOrderLineItem) GetShipSet() string`

GetShipSet returns the ShipSet field if non-nil, zero value otherwise.

### GetShipSetOk

`func (o *PurchaseOrderLineItem) GetShipSetOk() (*string, bool)`

GetShipSetOk returns a tuple with the ShipSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipSet

`func (o *PurchaseOrderLineItem) SetShipSet(v string)`

SetShipSet sets ShipSet field to given value.

### HasShipSet

`func (o *PurchaseOrderLineItem) HasShipSet() bool`

HasShipSet returns a boolean if a field has been set.

### GetDateReceived

`func (o *PurchaseOrderLineItem) GetDateReceived() time.Time`

GetDateReceived returns the DateReceived field if non-nil, zero value otherwise.

### GetDateReceivedOk

`func (o *PurchaseOrderLineItem) GetDateReceivedOk() (*time.Time, bool)`

GetDateReceivedOk returns a tuple with the DateReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateReceived

`func (o *PurchaseOrderLineItem) SetDateReceived(v time.Time)`

SetDateReceived sets DateReceived field to given value.

### HasDateReceived

`func (o *PurchaseOrderLineItem) HasDateReceived() bool`

HasDateReceived returns a boolean if a field has been set.

### GetReceivedStatus

`func (o *PurchaseOrderLineItem) GetReceivedStatus() string`

GetReceivedStatus returns the ReceivedStatus field if non-nil, zero value otherwise.

### GetReceivedStatusOk

`func (o *PurchaseOrderLineItem) GetReceivedStatusOk() (*string, bool)`

GetReceivedStatusOk returns a tuple with the ReceivedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedStatus

`func (o *PurchaseOrderLineItem) SetReceivedStatus(v string)`

SetReceivedStatus sets ReceivedStatus field to given value.

### HasReceivedStatus

`func (o *PurchaseOrderLineItem) HasReceivedStatus() bool`

HasReceivedStatus returns a boolean if a field has been set.

### SetReceivedStatusNil

`func (o *PurchaseOrderLineItem) SetReceivedStatusNil(b bool)`

 SetReceivedStatusNil sets the value for ReceivedStatus to be an explicit nil

### UnsetReceivedStatus
`func (o *PurchaseOrderLineItem) UnsetReceivedStatus()`

UnsetReceivedStatus ensures that no value is present for ReceivedStatus, not even an explicit nil
### GetInfo

`func (o *PurchaseOrderLineItem) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PurchaseOrderLineItem) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PurchaseOrderLineItem) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PurchaseOrderLineItem) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetExtCost

`func (o *PurchaseOrderLineItem) GetExtCost() float64`

GetExtCost returns the ExtCost field if non-nil, zero value otherwise.

### GetExtCostOk

`func (o *PurchaseOrderLineItem) GetExtCostOk() (*float64, bool)`

GetExtCostOk returns a tuple with the ExtCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtCost

`func (o *PurchaseOrderLineItem) SetExtCost(v float64)`

SetExtCost sets ExtCost field to given value.

### HasExtCost

`func (o *PurchaseOrderLineItem) HasExtCost() bool`

HasExtCost returns a boolean if a field has been set.

### GetExpectedArrivalDate

`func (o *PurchaseOrderLineItem) GetExpectedArrivalDate() time.Time`

GetExpectedArrivalDate returns the ExpectedArrivalDate field if non-nil, zero value otherwise.

### GetExpectedArrivalDateOk

`func (o *PurchaseOrderLineItem) GetExpectedArrivalDateOk() (*time.Time, bool)`

GetExpectedArrivalDateOk returns a tuple with the ExpectedArrivalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedArrivalDate

`func (o *PurchaseOrderLineItem) SetExpectedArrivalDate(v time.Time)`

SetExpectedArrivalDate sets ExpectedArrivalDate field to given value.

### HasExpectedArrivalDate

`func (o *PurchaseOrderLineItem) HasExpectedArrivalDate() bool`

HasExpectedArrivalDate returns a boolean if a field has been set.

### GetIsDetachAvailable

`func (o *PurchaseOrderLineItem) GetIsDetachAvailable() bool`

GetIsDetachAvailable returns the IsDetachAvailable field if non-nil, zero value otherwise.

### GetIsDetachAvailableOk

`func (o *PurchaseOrderLineItem) GetIsDetachAvailableOk() (*bool, bool)`

GetIsDetachAvailableOk returns a tuple with the IsDetachAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDetachAvailable

`func (o *PurchaseOrderLineItem) SetIsDetachAvailable(v bool)`

SetIsDetachAvailable sets IsDetachAvailable field to given value.

### HasIsDetachAvailable

`func (o *PurchaseOrderLineItem) HasIsDetachAvailable() bool`

HasIsDetachAvailable returns a boolean if a field has been set.

### GetBatchedFlag

`func (o *PurchaseOrderLineItem) GetBatchedFlag() bool`

GetBatchedFlag returns the BatchedFlag field if non-nil, zero value otherwise.

### GetBatchedFlagOk

`func (o *PurchaseOrderLineItem) GetBatchedFlagOk() (*bool, bool)`

GetBatchedFlagOk returns a tuple with the BatchedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchedFlag

`func (o *PurchaseOrderLineItem) SetBatchedFlag(v bool)`

SetBatchedFlag sets BatchedFlag field to given value.

### HasBatchedFlag

`func (o *PurchaseOrderLineItem) HasBatchedFlag() bool`

HasBatchedFlag returns a boolean if a field has been set.

### GetUnbatchedRecId

`func (o *PurchaseOrderLineItem) GetUnbatchedRecId() int32`

GetUnbatchedRecId returns the UnbatchedRecId field if non-nil, zero value otherwise.

### GetUnbatchedRecIdOk

`func (o *PurchaseOrderLineItem) GetUnbatchedRecIdOk() (*int32, bool)`

GetUnbatchedRecIdOk returns a tuple with the UnbatchedRecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnbatchedRecId

`func (o *PurchaseOrderLineItem) SetUnbatchedRecId(v int32)`

SetUnbatchedRecId sets UnbatchedRecId field to given value.

### HasUnbatchedRecId

`func (o *PurchaseOrderLineItem) HasUnbatchedRecId() bool`

HasUnbatchedRecId returns a boolean if a field has been set.

### SetUnbatchedRecIdNil

`func (o *PurchaseOrderLineItem) SetUnbatchedRecIdNil(b bool)`

 SetUnbatchedRecIdNil sets the value for UnbatchedRecId to be an explicit nil

### UnsetUnbatchedRecId
`func (o *PurchaseOrderLineItem) UnsetUnbatchedRecId()`

UnsetUnbatchedRecId ensures that no value is present for UnbatchedRecId, not even an explicit nil
### GetSalesOrder

`func (o *PurchaseOrderLineItem) GetSalesOrder() []SalesOrderReference`

GetSalesOrder returns the SalesOrder field if non-nil, zero value otherwise.

### GetSalesOrderOk

`func (o *PurchaseOrderLineItem) GetSalesOrderOk() (*[]SalesOrderReference, bool)`

GetSalesOrderOk returns a tuple with the SalesOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrder

`func (o *PurchaseOrderLineItem) SetSalesOrder(v []SalesOrderReference)`

SetSalesOrder sets SalesOrder field to given value.

### HasSalesOrder

`func (o *PurchaseOrderLineItem) HasSalesOrder() bool`

HasSalesOrder returns a boolean if a field has been set.

### GetCustomFields

`func (o *PurchaseOrderLineItem) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PurchaseOrderLineItem) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PurchaseOrderLineItem) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PurchaseOrderLineItem) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


