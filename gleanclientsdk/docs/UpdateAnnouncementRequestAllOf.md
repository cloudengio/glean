# UpdateAnnouncementRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The opaque id of the announcement. | 
**Data** | Pointer to [**AnnouncementCreateOrUpdateData**](AnnouncementCreateOrUpdateData.md) |  | [optional] 

## Methods

### NewUpdateAnnouncementRequestAllOf

`func NewUpdateAnnouncementRequestAllOf(id int32, ) *UpdateAnnouncementRequestAllOf`

NewUpdateAnnouncementRequestAllOf instantiates a new UpdateAnnouncementRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAnnouncementRequestAllOfWithDefaults

`func NewUpdateAnnouncementRequestAllOfWithDefaults() *UpdateAnnouncementRequestAllOf`

NewUpdateAnnouncementRequestAllOfWithDefaults instantiates a new UpdateAnnouncementRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateAnnouncementRequestAllOf) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateAnnouncementRequestAllOf) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateAnnouncementRequestAllOf) SetId(v int32)`

SetId sets Id field to given value.


### GetData

`func (o *UpdateAnnouncementRequestAllOf) GetData() AnnouncementCreateOrUpdateData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateAnnouncementRequestAllOf) GetDataOk() (*AnnouncementCreateOrUpdateData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateAnnouncementRequestAllOf) SetData(v AnnouncementCreateOrUpdateData)`

SetData sets Data field to given value.

### HasData

`func (o *UpdateAnnouncementRequestAllOf) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


