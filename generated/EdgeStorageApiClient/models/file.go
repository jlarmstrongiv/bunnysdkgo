package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type File struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ArrayNumber property
    arrayNumber *int64
    // The Checksum property
    checksum *string
    // The ContentType property
    contentType *string
    // The DateCreated property
    dateCreated *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Guid property
    guid *string
    // The IsDirectory property
    isDirectory *bool
    // The LastChanged property
    lastChanged *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Length property
    length *int64
    // The ObjectName property
    objectName *string
    // The Path property
    path *string
    // The ReplicatedZones property
    replicatedZones *string
    // The ServerId property
    serverId *int64
    // The StorageZoneId property
    storageZoneId *int64
    // The StorageZoneName property
    storageZoneName *string
    // The UserId property
    userId *string
}
// NewFile instantiates a new File and sets the default values.
func NewFile()(*File) {
    m := &File{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFile(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *File) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArrayNumber gets the ArrayNumber property value. The ArrayNumber property
// returns a *int64 when successful
func (m *File) GetArrayNumber()(*int64) {
    return m.arrayNumber
}
// GetChecksum gets the Checksum property value. The Checksum property
// returns a *string when successful
func (m *File) GetChecksum()(*string) {
    return m.checksum
}
// GetContentType gets the ContentType property value. The ContentType property
// returns a *string when successful
func (m *File) GetContentType()(*string) {
    return m.contentType
}
// GetDateCreated gets the DateCreated property value. The DateCreated property
// returns a *Time when successful
func (m *File) GetDateCreated()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dateCreated
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *File) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ArrayNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArrayNumber(val)
        }
        return nil
    }
    res["Checksum"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChecksum(val)
        }
        return nil
    }
    res["ContentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentType(val)
        }
        return nil
    }
    res["DateCreated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateCreated(val)
        }
        return nil
    }
    res["Guid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGuid(val)
        }
        return nil
    }
    res["IsDirectory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDirectory(val)
        }
        return nil
    }
    res["LastChanged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastChanged(val)
        }
        return nil
    }
    res["Length"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLength(val)
        }
        return nil
    }
    res["ObjectName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectName(val)
        }
        return nil
    }
    res["Path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
        }
        return nil
    }
    res["ReplicatedZones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReplicatedZones(val)
        }
        return nil
    }
    res["ServerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServerId(val)
        }
        return nil
    }
    res["StorageZoneId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageZoneId(val)
        }
        return nil
    }
    res["StorageZoneName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageZoneName(val)
        }
        return nil
    }
    res["UserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetGuid gets the Guid property value. The Guid property
// returns a *string when successful
func (m *File) GetGuid()(*string) {
    return m.guid
}
// GetIsDirectory gets the IsDirectory property value. The IsDirectory property
// returns a *bool when successful
func (m *File) GetIsDirectory()(*bool) {
    return m.isDirectory
}
// GetLastChanged gets the LastChanged property value. The LastChanged property
// returns a *Time when successful
func (m *File) GetLastChanged()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastChanged
}
// GetLength gets the Length property value. The Length property
// returns a *int64 when successful
func (m *File) GetLength()(*int64) {
    return m.length
}
// GetObjectName gets the ObjectName property value. The ObjectName property
// returns a *string when successful
func (m *File) GetObjectName()(*string) {
    return m.objectName
}
// GetPath gets the Path property value. The Path property
// returns a *string when successful
func (m *File) GetPath()(*string) {
    return m.path
}
// GetReplicatedZones gets the ReplicatedZones property value. The ReplicatedZones property
// returns a *string when successful
func (m *File) GetReplicatedZones()(*string) {
    return m.replicatedZones
}
// GetServerId gets the ServerId property value. The ServerId property
// returns a *int64 when successful
func (m *File) GetServerId()(*int64) {
    return m.serverId
}
// GetStorageZoneId gets the StorageZoneId property value. The StorageZoneId property
// returns a *int64 when successful
func (m *File) GetStorageZoneId()(*int64) {
    return m.storageZoneId
}
// GetStorageZoneName gets the StorageZoneName property value. The StorageZoneName property
// returns a *string when successful
func (m *File) GetStorageZoneName()(*string) {
    return m.storageZoneName
}
// GetUserId gets the UserId property value. The UserId property
// returns a *string when successful
func (m *File) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *File) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt64Value("ArrayNumber", m.GetArrayNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Checksum", m.GetChecksum())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ContentType", m.GetContentType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("DateCreated", m.GetDateCreated())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Guid", m.GetGuid())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("IsDirectory", m.GetIsDirectory())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("LastChanged", m.GetLastChanged())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("Length", m.GetLength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ObjectName", m.GetObjectName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Path", m.GetPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ReplicatedZones", m.GetReplicatedZones())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("ServerId", m.GetServerId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("StorageZoneId", m.GetStorageZoneId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("StorageZoneName", m.GetStorageZoneName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("UserId", m.GetUserId())
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
func (m *File) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArrayNumber sets the ArrayNumber property value. The ArrayNumber property
func (m *File) SetArrayNumber(value *int64)() {
    m.arrayNumber = value
}
// SetChecksum sets the Checksum property value. The Checksum property
func (m *File) SetChecksum(value *string)() {
    m.checksum = value
}
// SetContentType sets the ContentType property value. The ContentType property
func (m *File) SetContentType(value *string)() {
    m.contentType = value
}
// SetDateCreated sets the DateCreated property value. The DateCreated property
func (m *File) SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dateCreated = value
}
// SetGuid sets the Guid property value. The Guid property
func (m *File) SetGuid(value *string)() {
    m.guid = value
}
// SetIsDirectory sets the IsDirectory property value. The IsDirectory property
func (m *File) SetIsDirectory(value *bool)() {
    m.isDirectory = value
}
// SetLastChanged sets the LastChanged property value. The LastChanged property
func (m *File) SetLastChanged(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastChanged = value
}
// SetLength sets the Length property value. The Length property
func (m *File) SetLength(value *int64)() {
    m.length = value
}
// SetObjectName sets the ObjectName property value. The ObjectName property
func (m *File) SetObjectName(value *string)() {
    m.objectName = value
}
// SetPath sets the Path property value. The Path property
func (m *File) SetPath(value *string)() {
    m.path = value
}
// SetReplicatedZones sets the ReplicatedZones property value. The ReplicatedZones property
func (m *File) SetReplicatedZones(value *string)() {
    m.replicatedZones = value
}
// SetServerId sets the ServerId property value. The ServerId property
func (m *File) SetServerId(value *int64)() {
    m.serverId = value
}
// SetStorageZoneId sets the StorageZoneId property value. The StorageZoneId property
func (m *File) SetStorageZoneId(value *int64)() {
    m.storageZoneId = value
}
// SetStorageZoneName sets the StorageZoneName property value. The StorageZoneName property
func (m *File) SetStorageZoneName(value *string)() {
    m.storageZoneName = value
}
// SetUserId sets the UserId property value. The UserId property
func (m *File) SetUserId(value *string)() {
    m.userId = value
}
type Fileable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArrayNumber()(*int64)
    GetChecksum()(*string)
    GetContentType()(*string)
    GetDateCreated()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetGuid()(*string)
    GetIsDirectory()(*bool)
    GetLastChanged()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLength()(*int64)
    GetObjectName()(*string)
    GetPath()(*string)
    GetReplicatedZones()(*string)
    GetServerId()(*int64)
    GetStorageZoneId()(*int64)
    GetStorageZoneName()(*string)
    GetUserId()(*string)
    SetArrayNumber(value *int64)()
    SetChecksum(value *string)()
    SetContentType(value *string)()
    SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetGuid(value *string)()
    SetIsDirectory(value *bool)()
    SetLastChanged(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLength(value *int64)()
    SetObjectName(value *string)()
    SetPath(value *string)()
    SetReplicatedZones(value *string)()
    SetServerId(value *int64)()
    SetStorageZoneId(value *int64)()
    SetStorageZoneName(value *string)()
    SetUserId(value *string)()
}
