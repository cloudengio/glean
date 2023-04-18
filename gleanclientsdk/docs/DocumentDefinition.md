# DocumentDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Document title, in plain text, if present. | [optional] 
**Container** | Pointer to **string** | The container name for the content (Folder for example for file content). | [optional] 
**ContainerDatasourceId** | Pointer to **string** | This represents the datasource sepcific id of the container. | [optional] 
**ContainerObjectType** | Pointer to **string** | This represents the object type of the container. It cannot have spaces or _ | [optional] 
**Datasource** | **string** |  | 
**ObjectType** | Pointer to **string** | The type of the document (Case, KnowledgeArticle for Salesforce for example). It cannot have spaces or _ | [optional] 
**ViewURL** | Pointer to **string** | The permalink for viewing the document. | [optional] 
**Id** | Pointer to **string** | The datasource specific id for the document. This field is case insensitive and should not be more than 200 characters in length. | [optional] 
**Summary** | Pointer to [**ContentDefinition**](ContentDefinition.md) |  | [optional] 
**Body** | Pointer to [**ContentDefinition**](ContentDefinition.md) |  | [optional] 
**Author** | Pointer to [**UserReferenceDefinition**](UserReferenceDefinition.md) |  | [optional] 
**Owner** | Pointer to [**UserReferenceDefinition**](UserReferenceDefinition.md) |  | [optional] 
**Permissions** | Pointer to [**DocumentPermissionsDefinition**](DocumentPermissionsDefinition.md) |  | [optional] 
**CreatedAt** | Pointer to **int64** | The creation time, in epoch seconds. | [optional] 
**UpdatedAt** | Pointer to **int64** | The last update time, in epoch seconds. | [optional] 
**UpdatedBy** | Pointer to [**UserReferenceDefinition**](UserReferenceDefinition.md) |  | [optional] 
**Tags** | Pointer to **[]string** | Labels associated with the document. | [optional] 
**Interactions** | Pointer to [**DocumentInteractionsDefinition**](DocumentInteractionsDefinition.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**AdditionalUrls** | Pointer to **[]string** | Additional variations of the url that this document points to. | [optional] 
**CustomProperties** | Pointer to [**[]CustomProperty**](CustomProperty.md) | Additional metadata properties of the document. | [optional] 

## Methods

### NewDocumentDefinition

`func NewDocumentDefinition(datasource string, ) *DocumentDefinition`

NewDocumentDefinition instantiates a new DocumentDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentDefinitionWithDefaults

`func NewDocumentDefinitionWithDefaults() *DocumentDefinition`

NewDocumentDefinitionWithDefaults instantiates a new DocumentDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *DocumentDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DocumentDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DocumentDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *DocumentDefinition) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContainer

`func (o *DocumentDefinition) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *DocumentDefinition) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *DocumentDefinition) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *DocumentDefinition) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetContainerDatasourceId

`func (o *DocumentDefinition) GetContainerDatasourceId() string`

GetContainerDatasourceId returns the ContainerDatasourceId field if non-nil, zero value otherwise.

### GetContainerDatasourceIdOk

`func (o *DocumentDefinition) GetContainerDatasourceIdOk() (*string, bool)`

GetContainerDatasourceIdOk returns a tuple with the ContainerDatasourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerDatasourceId

`func (o *DocumentDefinition) SetContainerDatasourceId(v string)`

SetContainerDatasourceId sets ContainerDatasourceId field to given value.

### HasContainerDatasourceId

`func (o *DocumentDefinition) HasContainerDatasourceId() bool`

HasContainerDatasourceId returns a boolean if a field has been set.

### GetContainerObjectType

`func (o *DocumentDefinition) GetContainerObjectType() string`

GetContainerObjectType returns the ContainerObjectType field if non-nil, zero value otherwise.

### GetContainerObjectTypeOk

`func (o *DocumentDefinition) GetContainerObjectTypeOk() (*string, bool)`

GetContainerObjectTypeOk returns a tuple with the ContainerObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerObjectType

`func (o *DocumentDefinition) SetContainerObjectType(v string)`

SetContainerObjectType sets ContainerObjectType field to given value.

### HasContainerObjectType

`func (o *DocumentDefinition) HasContainerObjectType() bool`

HasContainerObjectType returns a boolean if a field has been set.

### GetDatasource

`func (o *DocumentDefinition) GetDatasource() string`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *DocumentDefinition) GetDatasourceOk() (*string, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *DocumentDefinition) SetDatasource(v string)`

SetDatasource sets Datasource field to given value.


### GetObjectType

`func (o *DocumentDefinition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *DocumentDefinition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *DocumentDefinition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *DocumentDefinition) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetViewURL

`func (o *DocumentDefinition) GetViewURL() string`

GetViewURL returns the ViewURL field if non-nil, zero value otherwise.

### GetViewURLOk

`func (o *DocumentDefinition) GetViewURLOk() (*string, bool)`

GetViewURLOk returns a tuple with the ViewURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewURL

`func (o *DocumentDefinition) SetViewURL(v string)`

SetViewURL sets ViewURL field to given value.

### HasViewURL

`func (o *DocumentDefinition) HasViewURL() bool`

HasViewURL returns a boolean if a field has been set.

### GetId

`func (o *DocumentDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocumentDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocumentDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocumentDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSummary

`func (o *DocumentDefinition) GetSummary() ContentDefinition`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *DocumentDefinition) GetSummaryOk() (*ContentDefinition, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *DocumentDefinition) SetSummary(v ContentDefinition)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *DocumentDefinition) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetBody

`func (o *DocumentDefinition) GetBody() ContentDefinition`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *DocumentDefinition) GetBodyOk() (*ContentDefinition, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *DocumentDefinition) SetBody(v ContentDefinition)`

SetBody sets Body field to given value.

### HasBody

`func (o *DocumentDefinition) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetAuthor

`func (o *DocumentDefinition) GetAuthor() UserReferenceDefinition`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *DocumentDefinition) GetAuthorOk() (*UserReferenceDefinition, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *DocumentDefinition) SetAuthor(v UserReferenceDefinition)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *DocumentDefinition) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetOwner

`func (o *DocumentDefinition) GetOwner() UserReferenceDefinition`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DocumentDefinition) GetOwnerOk() (*UserReferenceDefinition, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DocumentDefinition) SetOwner(v UserReferenceDefinition)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DocumentDefinition) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPermissions

`func (o *DocumentDefinition) GetPermissions() DocumentPermissionsDefinition`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *DocumentDefinition) GetPermissionsOk() (*DocumentPermissionsDefinition, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *DocumentDefinition) SetPermissions(v DocumentPermissionsDefinition)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *DocumentDefinition) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DocumentDefinition) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocumentDefinition) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocumentDefinition) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DocumentDefinition) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DocumentDefinition) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DocumentDefinition) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DocumentDefinition) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DocumentDefinition) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DocumentDefinition) GetUpdatedBy() UserReferenceDefinition`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DocumentDefinition) GetUpdatedByOk() (*UserReferenceDefinition, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DocumentDefinition) SetUpdatedBy(v UserReferenceDefinition)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DocumentDefinition) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetTags

`func (o *DocumentDefinition) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DocumentDefinition) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DocumentDefinition) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DocumentDefinition) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetInteractions

`func (o *DocumentDefinition) GetInteractions() DocumentInteractionsDefinition`

GetInteractions returns the Interactions field if non-nil, zero value otherwise.

### GetInteractionsOk

`func (o *DocumentDefinition) GetInteractionsOk() (*DocumentInteractionsDefinition, bool)`

GetInteractionsOk returns a tuple with the Interactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractions

`func (o *DocumentDefinition) SetInteractions(v DocumentInteractionsDefinition)`

SetInteractions sets Interactions field to given value.

### HasInteractions

`func (o *DocumentDefinition) HasInteractions() bool`

HasInteractions returns a boolean if a field has been set.

### GetStatus

`func (o *DocumentDefinition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DocumentDefinition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DocumentDefinition) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DocumentDefinition) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAdditionalUrls

`func (o *DocumentDefinition) GetAdditionalUrls() []string`

GetAdditionalUrls returns the AdditionalUrls field if non-nil, zero value otherwise.

### GetAdditionalUrlsOk

`func (o *DocumentDefinition) GetAdditionalUrlsOk() (*[]string, bool)`

GetAdditionalUrlsOk returns a tuple with the AdditionalUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalUrls

`func (o *DocumentDefinition) SetAdditionalUrls(v []string)`

SetAdditionalUrls sets AdditionalUrls field to given value.

### HasAdditionalUrls

`func (o *DocumentDefinition) HasAdditionalUrls() bool`

HasAdditionalUrls returns a boolean if a field has been set.

### GetCustomProperties

`func (o *DocumentDefinition) GetCustomProperties() []CustomProperty`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *DocumentDefinition) GetCustomPropertiesOk() (*[]CustomProperty, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *DocumentDefinition) SetCustomProperties(v []CustomProperty)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *DocumentDefinition) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


