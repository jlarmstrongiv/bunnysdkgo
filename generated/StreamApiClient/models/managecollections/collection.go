package managecollections

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Collection struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The unique ID of the collection
    guid *string
    // The name of the collection
    name *string
    // The URLs of preview images of videos in the collection
    previewImageUrls *string
    // The IDs of videos to be used as preview icons
    previewVideoIds *string
    // The total storage size of the collection
    totalSize *int64
    // The number of videos that the collection contains
    videoCount *int64
    // The video library ID that contains the collection
    videoLibraryId *int64
}
// NewCollection instantiates a new Collection and sets the default values.
func NewCollection()(*Collection) {
    m := &Collection{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCollectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCollectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCollection(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Collection) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Collection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["guid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGuid(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["previewImageUrls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewImageUrls(val)
        }
        return nil
    }
    res["previewVideoIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewVideoIds(val)
        }
        return nil
    }
    res["totalSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalSize(val)
        }
        return nil
    }
    res["videoCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideoCount(val)
        }
        return nil
    }
    res["videoLibraryId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideoLibraryId(val)
        }
        return nil
    }
    return res
}
// GetGuid gets the guid property value. The unique ID of the collection
// returns a *string when successful
func (m *Collection) GetGuid()(*string) {
    return m.guid
}
// GetName gets the name property value. The name of the collection
// returns a *string when successful
func (m *Collection) GetName()(*string) {
    return m.name
}
// GetPreviewImageUrls gets the previewImageUrls property value. The URLs of preview images of videos in the collection
// returns a *string when successful
func (m *Collection) GetPreviewImageUrls()(*string) {
    return m.previewImageUrls
}
// GetPreviewVideoIds gets the previewVideoIds property value. The IDs of videos to be used as preview icons
// returns a *string when successful
func (m *Collection) GetPreviewVideoIds()(*string) {
    return m.previewVideoIds
}
// GetTotalSize gets the totalSize property value. The total storage size of the collection
// returns a *int64 when successful
func (m *Collection) GetTotalSize()(*int64) {
    return m.totalSize
}
// GetVideoCount gets the videoCount property value. The number of videos that the collection contains
// returns a *int64 when successful
func (m *Collection) GetVideoCount()(*int64) {
    return m.videoCount
}
// GetVideoLibraryId gets the videoLibraryId property value. The video library ID that contains the collection
// returns a *int64 when successful
func (m *Collection) GetVideoLibraryId()(*int64) {
    return m.videoLibraryId
}
// Serialize serializes information the current object
func (m *Collection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Collection) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGuid sets the guid property value. The unique ID of the collection
func (m *Collection) SetGuid(value *string)() {
    m.guid = value
}
// SetName sets the name property value. The name of the collection
func (m *Collection) SetName(value *string)() {
    m.name = value
}
// SetPreviewImageUrls sets the previewImageUrls property value. The URLs of preview images of videos in the collection
func (m *Collection) SetPreviewImageUrls(value *string)() {
    m.previewImageUrls = value
}
// SetPreviewVideoIds sets the previewVideoIds property value. The IDs of videos to be used as preview icons
func (m *Collection) SetPreviewVideoIds(value *string)() {
    m.previewVideoIds = value
}
// SetTotalSize sets the totalSize property value. The total storage size of the collection
func (m *Collection) SetTotalSize(value *int64)() {
    m.totalSize = value
}
// SetVideoCount sets the videoCount property value. The number of videos that the collection contains
func (m *Collection) SetVideoCount(value *int64)() {
    m.videoCount = value
}
// SetVideoLibraryId sets the videoLibraryId property value. The video library ID that contains the collection
func (m *Collection) SetVideoLibraryId(value *int64)() {
    m.videoLibraryId = value
}
type Collectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGuid()(*string)
    GetName()(*string)
    GetPreviewImageUrls()(*string)
    GetPreviewVideoIds()(*string)
    GetTotalSize()(*int64)
    GetVideoCount()(*int64)
    GetVideoLibraryId()(*int64)
    SetGuid(value *string)()
    SetName(value *string)()
    SetPreviewImageUrls(value *string)()
    SetPreviewVideoIds(value *string)()
    SetTotalSize(value *int64)()
    SetVideoCount(value *int64)()
    SetVideoLibraryId(value *int64)()
}
