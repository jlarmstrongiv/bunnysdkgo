package cleanupunconfiguredresolutions

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CleanupUnconfiguredResolutionsData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The availableResolutionsAfter property
    availableResolutionsAfter []string
    // The deletedOriginal property
    deletedOriginal *bool
    // The deletedStorageObjects property
    deletedStorageObjects *bool
    // The modifiedPlaylist property
    modifiedPlaylist *bool
    // The resolutionsToDelete property
    resolutionsToDelete []string
    // The resolutionsToDeleteFromMp4 property
    resolutionsToDeleteFromMp4 []string
    // The resolutionsToDeleteFromPlaylist property
    resolutionsToDeleteFromPlaylist []string
    // The resolutionsToDeleteFromStorage property
    resolutionsToDeleteFromStorage []string
    // The storageObjectsToDelete property
    storageObjectsToDelete []string
    // The updatedAvailableResolutions property
    updatedAvailableResolutions *bool
    // The videoId property
    videoId *string
    // The videoLibraryId property
    videoLibraryId *int64
}
// NewCleanupUnconfiguredResolutionsData instantiates a new CleanupUnconfiguredResolutionsData and sets the default values.
func NewCleanupUnconfiguredResolutionsData()(*CleanupUnconfiguredResolutionsData) {
    m := &CleanupUnconfiguredResolutionsData{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCleanupUnconfiguredResolutionsDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCleanupUnconfiguredResolutionsDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCleanupUnconfiguredResolutionsData(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CleanupUnconfiguredResolutionsData) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvailableResolutionsAfter gets the availableResolutionsAfter property value. The availableResolutionsAfter property
// returns a []string when successful
func (m *CleanupUnconfiguredResolutionsData) GetAvailableResolutionsAfter()([]string) {
    return m.availableResolutionsAfter
}
// GetDeletedOriginal gets the deletedOriginal property value. The deletedOriginal property
// returns a *bool when successful
func (m *CleanupUnconfiguredResolutionsData) GetDeletedOriginal()(*bool) {
    return m.deletedOriginal
}
// GetDeletedStorageObjects gets the deletedStorageObjects property value. The deletedStorageObjects property
// returns a *bool when successful
func (m *CleanupUnconfiguredResolutionsData) GetDeletedStorageObjects()(*bool) {
    return m.deletedStorageObjects
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CleanupUnconfiguredResolutionsData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["availableResolutionsAfter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetAvailableResolutionsAfter(res)
        }
        return nil
    }
    res["deletedOriginal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletedOriginal(val)
        }
        return nil
    }
    res["deletedStorageObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletedStorageObjects(val)
        }
        return nil
    }
    res["modifiedPlaylist"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedPlaylist(val)
        }
        return nil
    }
    res["resolutionsToDelete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetResolutionsToDelete(res)
        }
        return nil
    }
    res["resolutionsToDeleteFromMp4"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetResolutionsToDeleteFromMp4(res)
        }
        return nil
    }
    res["resolutionsToDeleteFromPlaylist"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetResolutionsToDeleteFromPlaylist(res)
        }
        return nil
    }
    res["resolutionsToDeleteFromStorage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetResolutionsToDeleteFromStorage(res)
        }
        return nil
    }
    res["storageObjectsToDelete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetStorageObjectsToDelete(res)
        }
        return nil
    }
    res["updatedAvailableResolutions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAvailableResolutions(val)
        }
        return nil
    }
    res["videoId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideoId(val)
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
// GetModifiedPlaylist gets the modifiedPlaylist property value. The modifiedPlaylist property
// returns a *bool when successful
func (m *CleanupUnconfiguredResolutionsData) GetModifiedPlaylist()(*bool) {
    return m.modifiedPlaylist
}
// GetResolutionsToDelete gets the resolutionsToDelete property value. The resolutionsToDelete property
// returns a []string when successful
func (m *CleanupUnconfiguredResolutionsData) GetResolutionsToDelete()([]string) {
    return m.resolutionsToDelete
}
// GetResolutionsToDeleteFromMp4 gets the resolutionsToDeleteFromMp4 property value. The resolutionsToDeleteFromMp4 property
// returns a []string when successful
func (m *CleanupUnconfiguredResolutionsData) GetResolutionsToDeleteFromMp4()([]string) {
    return m.resolutionsToDeleteFromMp4
}
// GetResolutionsToDeleteFromPlaylist gets the resolutionsToDeleteFromPlaylist property value. The resolutionsToDeleteFromPlaylist property
// returns a []string when successful
func (m *CleanupUnconfiguredResolutionsData) GetResolutionsToDeleteFromPlaylist()([]string) {
    return m.resolutionsToDeleteFromPlaylist
}
// GetResolutionsToDeleteFromStorage gets the resolutionsToDeleteFromStorage property value. The resolutionsToDeleteFromStorage property
// returns a []string when successful
func (m *CleanupUnconfiguredResolutionsData) GetResolutionsToDeleteFromStorage()([]string) {
    return m.resolutionsToDeleteFromStorage
}
// GetStorageObjectsToDelete gets the storageObjectsToDelete property value. The storageObjectsToDelete property
// returns a []string when successful
func (m *CleanupUnconfiguredResolutionsData) GetStorageObjectsToDelete()([]string) {
    return m.storageObjectsToDelete
}
// GetUpdatedAvailableResolutions gets the updatedAvailableResolutions property value. The updatedAvailableResolutions property
// returns a *bool when successful
func (m *CleanupUnconfiguredResolutionsData) GetUpdatedAvailableResolutions()(*bool) {
    return m.updatedAvailableResolutions
}
// GetVideoId gets the videoId property value. The videoId property
// returns a *string when successful
func (m *CleanupUnconfiguredResolutionsData) GetVideoId()(*string) {
    return m.videoId
}
// GetVideoLibraryId gets the videoLibraryId property value. The videoLibraryId property
// returns a *int64 when successful
func (m *CleanupUnconfiguredResolutionsData) GetVideoLibraryId()(*int64) {
    return m.videoLibraryId
}
// Serialize serializes information the current object
func (m *CleanupUnconfiguredResolutionsData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAvailableResolutionsAfter() != nil {
        err := writer.WriteCollectionOfStringValues("availableResolutionsAfter", m.GetAvailableResolutionsAfter())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("deletedOriginal", m.GetDeletedOriginal())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("deletedStorageObjects", m.GetDeletedStorageObjects())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("modifiedPlaylist", m.GetModifiedPlaylist())
        if err != nil {
            return err
        }
    }
    if m.GetResolutionsToDelete() != nil {
        err := writer.WriteCollectionOfStringValues("resolutionsToDelete", m.GetResolutionsToDelete())
        if err != nil {
            return err
        }
    }
    if m.GetResolutionsToDeleteFromMp4() != nil {
        err := writer.WriteCollectionOfStringValues("resolutionsToDeleteFromMp4", m.GetResolutionsToDeleteFromMp4())
        if err != nil {
            return err
        }
    }
    if m.GetResolutionsToDeleteFromPlaylist() != nil {
        err := writer.WriteCollectionOfStringValues("resolutionsToDeleteFromPlaylist", m.GetResolutionsToDeleteFromPlaylist())
        if err != nil {
            return err
        }
    }
    if m.GetResolutionsToDeleteFromStorage() != nil {
        err := writer.WriteCollectionOfStringValues("resolutionsToDeleteFromStorage", m.GetResolutionsToDeleteFromStorage())
        if err != nil {
            return err
        }
    }
    if m.GetStorageObjectsToDelete() != nil {
        err := writer.WriteCollectionOfStringValues("storageObjectsToDelete", m.GetStorageObjectsToDelete())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("updatedAvailableResolutions", m.GetUpdatedAvailableResolutions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("videoId", m.GetVideoId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("videoLibraryId", m.GetVideoLibraryId())
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
func (m *CleanupUnconfiguredResolutionsData) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvailableResolutionsAfter sets the availableResolutionsAfter property value. The availableResolutionsAfter property
func (m *CleanupUnconfiguredResolutionsData) SetAvailableResolutionsAfter(value []string)() {
    m.availableResolutionsAfter = value
}
// SetDeletedOriginal sets the deletedOriginal property value. The deletedOriginal property
func (m *CleanupUnconfiguredResolutionsData) SetDeletedOriginal(value *bool)() {
    m.deletedOriginal = value
}
// SetDeletedStorageObjects sets the deletedStorageObjects property value. The deletedStorageObjects property
func (m *CleanupUnconfiguredResolutionsData) SetDeletedStorageObjects(value *bool)() {
    m.deletedStorageObjects = value
}
// SetModifiedPlaylist sets the modifiedPlaylist property value. The modifiedPlaylist property
func (m *CleanupUnconfiguredResolutionsData) SetModifiedPlaylist(value *bool)() {
    m.modifiedPlaylist = value
}
// SetResolutionsToDelete sets the resolutionsToDelete property value. The resolutionsToDelete property
func (m *CleanupUnconfiguredResolutionsData) SetResolutionsToDelete(value []string)() {
    m.resolutionsToDelete = value
}
// SetResolutionsToDeleteFromMp4 sets the resolutionsToDeleteFromMp4 property value. The resolutionsToDeleteFromMp4 property
func (m *CleanupUnconfiguredResolutionsData) SetResolutionsToDeleteFromMp4(value []string)() {
    m.resolutionsToDeleteFromMp4 = value
}
// SetResolutionsToDeleteFromPlaylist sets the resolutionsToDeleteFromPlaylist property value. The resolutionsToDeleteFromPlaylist property
func (m *CleanupUnconfiguredResolutionsData) SetResolutionsToDeleteFromPlaylist(value []string)() {
    m.resolutionsToDeleteFromPlaylist = value
}
// SetResolutionsToDeleteFromStorage sets the resolutionsToDeleteFromStorage property value. The resolutionsToDeleteFromStorage property
func (m *CleanupUnconfiguredResolutionsData) SetResolutionsToDeleteFromStorage(value []string)() {
    m.resolutionsToDeleteFromStorage = value
}
// SetStorageObjectsToDelete sets the storageObjectsToDelete property value. The storageObjectsToDelete property
func (m *CleanupUnconfiguredResolutionsData) SetStorageObjectsToDelete(value []string)() {
    m.storageObjectsToDelete = value
}
// SetUpdatedAvailableResolutions sets the updatedAvailableResolutions property value. The updatedAvailableResolutions property
func (m *CleanupUnconfiguredResolutionsData) SetUpdatedAvailableResolutions(value *bool)() {
    m.updatedAvailableResolutions = value
}
// SetVideoId sets the videoId property value. The videoId property
func (m *CleanupUnconfiguredResolutionsData) SetVideoId(value *string)() {
    m.videoId = value
}
// SetVideoLibraryId sets the videoLibraryId property value. The videoLibraryId property
func (m *CleanupUnconfiguredResolutionsData) SetVideoLibraryId(value *int64)() {
    m.videoLibraryId = value
}
type CleanupUnconfiguredResolutionsDataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAvailableResolutionsAfter()([]string)
    GetDeletedOriginal()(*bool)
    GetDeletedStorageObjects()(*bool)
    GetModifiedPlaylist()(*bool)
    GetResolutionsToDelete()([]string)
    GetResolutionsToDeleteFromMp4()([]string)
    GetResolutionsToDeleteFromPlaylist()([]string)
    GetResolutionsToDeleteFromStorage()([]string)
    GetStorageObjectsToDelete()([]string)
    GetUpdatedAvailableResolutions()(*bool)
    GetVideoId()(*string)
    GetVideoLibraryId()(*int64)
    SetAvailableResolutionsAfter(value []string)()
    SetDeletedOriginal(value *bool)()
    SetDeletedStorageObjects(value *bool)()
    SetModifiedPlaylist(value *bool)()
    SetResolutionsToDelete(value []string)()
    SetResolutionsToDeleteFromMp4(value []string)()
    SetResolutionsToDeleteFromPlaylist(value []string)()
    SetResolutionsToDeleteFromStorage(value []string)()
    SetStorageObjectsToDelete(value []string)()
    SetUpdatedAvailableResolutions(value *bool)()
    SetVideoId(value *string)()
    SetVideoLibraryId(value *int64)()
}
