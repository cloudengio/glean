# AnnouncementAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The opaque id of the announcement. | [optional] 
**Author** | Pointer to [**Person**](Person.md) |  | [optional] 
**CreateTimestamp** | Pointer to **int32** | Server Unix timestamp of the creation time (in seconds since epoch UTC). | [optional] 
**LastUpdateTimestamp** | Pointer to **int32** | Server Unix timestamp of the last update time (in seconds since epoch UTC). | [optional] 
**UpdatedBy** | Pointer to [**Person**](Person.md) |  | [optional] 
**ViewerInfo** | Pointer to [**AnnouncementAllOfViewerInfo**](AnnouncementAllOfViewerInfo.md) |  | [optional] 
**SourceDocument** | Pointer to [**Document**](Document.md) |  | [optional] 
**IsPublished** | Pointer to **bool** | Whether or not the announcement is published. | [optional] 

## Methods

### NewAnnouncementAllOf

`func NewAnnouncementAllOf() *AnnouncementAllOf`

NewAnnouncementAllOf instantiates a new AnnouncementAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnouncementAllOfWithDefaults

`func NewAnnouncementAllOfWithDefaults() *AnnouncementAllOf`

NewAnnouncementAllOfWithDefaults instantiates a new AnnouncementAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AnnouncementAllOf) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnnouncementAllOf) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnnouncementAllOf) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AnnouncementAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthor

`func (o *AnnouncementAllOf) GetAuthor() Person`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *AnnouncementAllOf) GetAuthorOk() (*Person, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *AnnouncementAllOf) SetAuthor(v Person)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *AnnouncementAllOf) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCreateTimestamp

`func (o *AnnouncementAllOf) GetCreateTimestamp() int32`

GetCreateTimestamp returns the CreateTimestamp field if non-nil, zero value otherwise.

### GetCreateTimestampOk

`func (o *AnnouncementAllOf) GetCreateTimestampOk() (*int32, bool)`

GetCreateTimestampOk returns a tuple with the CreateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimestamp

`func (o *AnnouncementAllOf) SetCreateTimestamp(v int32)`

SetCreateTimestamp sets CreateTimestamp field to given value.

### HasCreateTimestamp

`func (o *AnnouncementAllOf) HasCreateTimestamp() bool`

HasCreateTimestamp returns a boolean if a field has been set.

### GetLastUpdateTimestamp

`func (o *AnnouncementAllOf) GetLastUpdateTimestamp() int32`

GetLastUpdateTimestamp returns the LastUpdateTimestamp field if non-nil, zero value otherwise.

### GetLastUpdateTimestampOk

`func (o *AnnouncementAllOf) GetLastUpdateTimestampOk() (*int32, bool)`

GetLastUpdateTimestampOk returns a tuple with the LastUpdateTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimestamp

`func (o *AnnouncementAllOf) SetLastUpdateTimestamp(v int32)`

SetLastUpdateTimestamp sets LastUpdateTimestamp field to given value.

### HasLastUpdateTimestamp

`func (o *AnnouncementAllOf) HasLastUpdateTimestamp() bool`

HasLastUpdateTimestamp returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *AnnouncementAllOf) GetUpdatedBy() Person`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AnnouncementAllOf) GetUpdatedByOk() (*Person, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AnnouncementAllOf) SetUpdatedBy(v Person)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *AnnouncementAllOf) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetViewerInfo

`func (o *AnnouncementAllOf) GetViewerInfo() AnnouncementAllOfViewerInfo`

GetViewerInfo returns the ViewerInfo field if non-nil, zero value otherwise.

### GetViewerInfoOk

`func (o *AnnouncementAllOf) GetViewerInfoOk() (*AnnouncementAllOfViewerInfo, bool)`

GetViewerInfoOk returns a tuple with the ViewerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerInfo

`func (o *AnnouncementAllOf) SetViewerInfo(v AnnouncementAllOfViewerInfo)`

SetViewerInfo sets ViewerInfo field to given value.

### HasViewerInfo

`func (o *AnnouncementAllOf) HasViewerInfo() bool`

HasViewerInfo returns a boolean if a field has been set.

### GetSourceDocument

`func (o *AnnouncementAllOf) GetSourceDocument() Document`

GetSourceDocument returns the SourceDocument field if non-nil, zero value otherwise.

### GetSourceDocumentOk

`func (o *AnnouncementAllOf) GetSourceDocumentOk() (*Document, bool)`

GetSourceDocumentOk returns a tuple with the SourceDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocument

`func (o *AnnouncementAllOf) SetSourceDocument(v Document)`

SetSourceDocument sets SourceDocument field to given value.

### HasSourceDocument

`func (o *AnnouncementAllOf) HasSourceDocument() bool`

HasSourceDocument returns a boolean if a field has been set.

### GetIsPublished

`func (o *AnnouncementAllOf) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *AnnouncementAllOf) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *AnnouncementAllOf) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *AnnouncementAllOf) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


