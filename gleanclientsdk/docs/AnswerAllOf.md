# AnswerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CombinedAnswerText** | Pointer to [**StructuredText**](StructuredText.md) |  | [optional] 
**Likes** | Pointer to [**AnswerLikes**](AnswerLikes.md) |  | [optional] 
**UserRole** | Pointer to **string** | DEPRECATED - use roles instead. User&#39;s role on the specific answer. | [optional] 
**Author** | Pointer to [**Person**](Person.md) |  | [optional] 
**CreateTime** | Pointer to **time.Time** | The time the answer was created in ISO format (ISO 8601). | [optional] 
**UpdateTime** | Pointer to **time.Time** | The time the answer was last updated in ISO format (ISO 8601). | [optional] 
**UpdatedBy** | Pointer to [**Person**](Person.md) |  | [optional] 
**Verification** | Pointer to [**Verification**](Verification.md) |  | [optional] 
**Board** | Pointer to [**AnswerBoard**](AnswerBoard.md) |  | [optional] 
**Collections** | Pointer to [**[]Collection**](Collection.md) | The collections to which the answer belongs. | [optional] 
**DocumentCategory** | Pointer to **string** | The document&#39;s document_category(.proto). | [optional] 
**SourceDocument** | Pointer to [**Document**](Document.md) |  | [optional] 

## Methods

### NewAnswerAllOf

`func NewAnswerAllOf() *AnswerAllOf`

NewAnswerAllOf instantiates a new AnswerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnswerAllOfWithDefaults

`func NewAnswerAllOfWithDefaults() *AnswerAllOf`

NewAnswerAllOfWithDefaults instantiates a new AnswerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCombinedAnswerText

`func (o *AnswerAllOf) GetCombinedAnswerText() StructuredText`

GetCombinedAnswerText returns the CombinedAnswerText field if non-nil, zero value otherwise.

### GetCombinedAnswerTextOk

`func (o *AnswerAllOf) GetCombinedAnswerTextOk() (*StructuredText, bool)`

GetCombinedAnswerTextOk returns a tuple with the CombinedAnswerText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedAnswerText

`func (o *AnswerAllOf) SetCombinedAnswerText(v StructuredText)`

SetCombinedAnswerText sets CombinedAnswerText field to given value.

### HasCombinedAnswerText

`func (o *AnswerAllOf) HasCombinedAnswerText() bool`

HasCombinedAnswerText returns a boolean if a field has been set.

### GetLikes

`func (o *AnswerAllOf) GetLikes() AnswerLikes`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *AnswerAllOf) GetLikesOk() (*AnswerLikes, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *AnswerAllOf) SetLikes(v AnswerLikes)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *AnswerAllOf) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetUserRole

`func (o *AnswerAllOf) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *AnswerAllOf) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *AnswerAllOf) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.

### HasUserRole

`func (o *AnswerAllOf) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.

### GetAuthor

`func (o *AnswerAllOf) GetAuthor() Person`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *AnswerAllOf) GetAuthorOk() (*Person, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *AnswerAllOf) SetAuthor(v Person)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *AnswerAllOf) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetCreateTime

`func (o *AnswerAllOf) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *AnswerAllOf) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *AnswerAllOf) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *AnswerAllOf) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *AnswerAllOf) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *AnswerAllOf) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *AnswerAllOf) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *AnswerAllOf) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *AnswerAllOf) GetUpdatedBy() Person`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AnswerAllOf) GetUpdatedByOk() (*Person, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AnswerAllOf) SetUpdatedBy(v Person)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *AnswerAllOf) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetVerification

`func (o *AnswerAllOf) GetVerification() Verification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *AnswerAllOf) GetVerificationOk() (*Verification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *AnswerAllOf) SetVerification(v Verification)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *AnswerAllOf) HasVerification() bool`

HasVerification returns a boolean if a field has been set.

### GetBoard

`func (o *AnswerAllOf) GetBoard() AnswerBoard`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *AnswerAllOf) GetBoardOk() (*AnswerBoard, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *AnswerAllOf) SetBoard(v AnswerBoard)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *AnswerAllOf) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetCollections

`func (o *AnswerAllOf) GetCollections() []Collection`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *AnswerAllOf) GetCollectionsOk() (*[]Collection, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *AnswerAllOf) SetCollections(v []Collection)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *AnswerAllOf) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetDocumentCategory

`func (o *AnswerAllOf) GetDocumentCategory() string`

GetDocumentCategory returns the DocumentCategory field if non-nil, zero value otherwise.

### GetDocumentCategoryOk

`func (o *AnswerAllOf) GetDocumentCategoryOk() (*string, bool)`

GetDocumentCategoryOk returns a tuple with the DocumentCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentCategory

`func (o *AnswerAllOf) SetDocumentCategory(v string)`

SetDocumentCategory sets DocumentCategory field to given value.

### HasDocumentCategory

`func (o *AnswerAllOf) HasDocumentCategory() bool`

HasDocumentCategory returns a boolean if a field has been set.

### GetSourceDocument

`func (o *AnswerAllOf) GetSourceDocument() Document`

GetSourceDocument returns the SourceDocument field if non-nil, zero value otherwise.

### GetSourceDocumentOk

`func (o *AnswerAllOf) GetSourceDocumentOk() (*Document, bool)`

GetSourceDocumentOk returns a tuple with the SourceDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDocument

`func (o *AnswerAllOf) SetSourceDocument(v Document)`

SetSourceDocument sets SourceDocument field to given value.

### HasSourceDocument

`func (o *AnswerAllOf) HasSourceDocument() bool`

HasSourceDocument returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


