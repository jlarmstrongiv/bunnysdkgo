package storagezone

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428 "github.com/jlarmstrongiv/bunnysdkgo/generated/BunnyApiClient/models/pullzone"
)

type StorageZone struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Custom404FilePath property
    custom404FilePath *string
    // The DateModified property
    dateModified *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Deleted property
    deleted *bool
    // The Discount property
    discount *int32
    // The FilesStored property
    filesStored *int64
    // The Id property
    id *int64
    // The Name property
    name *string
    // The Password property
    password *string
    // The PriceOverride property
    priceOverride *int64
    // The PullZones property
    pullZones []id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428.PullZoneable
    // The ReadOnlyPassword property
    readOnlyPassword *string
    // The Region property
    region *StandardRegions
    // The ReplicationChangeInProgress property
    replicationChangeInProgress *bool
    // The ReplicationRegions property
    replicationRegions []EdgeReplicationRegions
    // The Rewrite404To200 property
    rewrite404To200 *bool
    // The StorageHostname property
    storageHostname *string
    // The StorageUsed property
    storageUsed *int64
    // The UserId property
    userId *string
    // The ZoneTier property
    zoneTier *float64
}
// NewStorageZone instantiates a new StorageZone and sets the default values.
func NewStorageZone()(*StorageZone) {
    m := &StorageZone{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStorageZoneFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStorageZoneFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStorageZone(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StorageZone) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustom404FilePath gets the Custom404FilePath property value. The Custom404FilePath property
// returns a *string when successful
func (m *StorageZone) GetCustom404FilePath()(*string) {
    return m.custom404FilePath
}
// GetDateModified gets the DateModified property value. The DateModified property
// returns a *Time when successful
func (m *StorageZone) GetDateModified()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dateModified
}
// GetDeleted gets the Deleted property value. The Deleted property
// returns a *bool when successful
func (m *StorageZone) GetDeleted()(*bool) {
    return m.deleted
}
// GetDiscount gets the Discount property value. The Discount property
// returns a *int32 when successful
func (m *StorageZone) GetDiscount()(*int32) {
    return m.discount
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StorageZone) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["Custom404FilePath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustom404FilePath(val)
        }
        return nil
    }
    res["DateModified"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateModified(val)
        }
        return nil
    }
    res["Deleted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeleted(val)
        }
        return nil
    }
    res["Discount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscount(val)
        }
        return nil
    }
    res["FilesStored"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilesStored(val)
        }
        return nil
    }
    res["Id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["Name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["Password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    res["PriceOverride"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPriceOverride(val)
        }
        return nil
    }
    res["PullZones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428.CreatePullZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428.PullZoneable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428.PullZoneable)
                }
            }
            m.SetPullZones(res)
        }
        return nil
    }
    res["ReadOnlyPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReadOnlyPassword(val)
        }
        return nil
    }
    res["Region"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseStandardRegions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegion(val.(*StandardRegions))
        }
        return nil
    }
    res["ReplicationChangeInProgress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReplicationChangeInProgress(val)
        }
        return nil
    }
    res["ReplicationRegions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseEdgeReplicationRegions)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EdgeReplicationRegions, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*EdgeReplicationRegions))
                }
            }
            m.SetReplicationRegions(res)
        }
        return nil
    }
    res["Rewrite404To200"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRewrite404To200(val)
        }
        return nil
    }
    res["StorageHostname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageHostname(val)
        }
        return nil
    }
    res["StorageUsed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageUsed(val)
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
    res["ZoneTier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetZoneTier(val)
        }
        return nil
    }
    return res
}
// GetFilesStored gets the FilesStored property value. The FilesStored property
// returns a *int64 when successful
func (m *StorageZone) GetFilesStored()(*int64) {
    return m.filesStored
}
// GetId gets the Id property value. The Id property
// returns a *int64 when successful
func (m *StorageZone) GetId()(*int64) {
    return m.id
}
// GetName gets the Name property value. The Name property
// returns a *string when successful
func (m *StorageZone) GetName()(*string) {
    return m.name
}
// GetPassword gets the Password property value. The Password property
// returns a *string when successful
func (m *StorageZone) GetPassword()(*string) {
    return m.password
}
// GetPriceOverride gets the PriceOverride property value. The PriceOverride property
// returns a *int64 when successful
func (m *StorageZone) GetPriceOverride()(*int64) {
    return m.priceOverride
}
// GetPullZones gets the PullZones property value. The PullZones property
// returns a []PullZoneable when successful
func (m *StorageZone) GetPullZones()([]id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428.PullZoneable) {
    return m.pullZones
}
// GetReadOnlyPassword gets the ReadOnlyPassword property value. The ReadOnlyPassword property
// returns a *string when successful
func (m *StorageZone) GetReadOnlyPassword()(*string) {
    return m.readOnlyPassword
}
// GetRegion gets the Region property value. The Region property
// returns a *StandardRegions when successful
func (m *StorageZone) GetRegion()(*StandardRegions) {
    return m.region
}
// GetReplicationChangeInProgress gets the ReplicationChangeInProgress property value. The ReplicationChangeInProgress property
// returns a *bool when successful
func (m *StorageZone) GetReplicationChangeInProgress()(*bool) {
    return m.replicationChangeInProgress
}
// GetReplicationRegions gets the ReplicationRegions property value. The ReplicationRegions property
// returns a []EdgeReplicationRegions when successful
func (m *StorageZone) GetReplicationRegions()([]EdgeReplicationRegions) {
    return m.replicationRegions
}
// GetRewrite404To200 gets the Rewrite404To200 property value. The Rewrite404To200 property
// returns a *bool when successful
func (m *StorageZone) GetRewrite404To200()(*bool) {
    return m.rewrite404To200
}
// GetStorageHostname gets the StorageHostname property value. The StorageHostname property
// returns a *string when successful
func (m *StorageZone) GetStorageHostname()(*string) {
    return m.storageHostname
}
// GetStorageUsed gets the StorageUsed property value. The StorageUsed property
// returns a *int64 when successful
func (m *StorageZone) GetStorageUsed()(*int64) {
    return m.storageUsed
}
// GetUserId gets the UserId property value. The UserId property
// returns a *string when successful
func (m *StorageZone) GetUserId()(*string) {
    return m.userId
}
// GetZoneTier gets the ZoneTier property value. The ZoneTier property
// returns a *float64 when successful
func (m *StorageZone) GetZoneTier()(*float64) {
    return m.zoneTier
}
// Serialize serializes information the current object
func (m *StorageZone) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("Custom404FilePath", m.GetCustom404FilePath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetRegion() != nil {
        cast := (*m.GetRegion()).String()
        err := writer.WriteStringValue("Region", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetReplicationRegions() != nil {
        err := writer.WriteCollectionOfStringValues("ReplicationRegions", SerializeEdgeReplicationRegions(m.GetReplicationRegions()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("Rewrite404To200", m.GetRewrite404To200())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("ZoneTier", m.GetZoneTier())
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
func (m *StorageZone) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustom404FilePath sets the Custom404FilePath property value. The Custom404FilePath property
func (m *StorageZone) SetCustom404FilePath(value *string)() {
    m.custom404FilePath = value
}
// SetDateModified sets the DateModified property value. The DateModified property
func (m *StorageZone) SetDateModified(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dateModified = value
}
// SetDeleted sets the Deleted property value. The Deleted property
func (m *StorageZone) SetDeleted(value *bool)() {
    m.deleted = value
}
// SetDiscount sets the Discount property value. The Discount property
func (m *StorageZone) SetDiscount(value *int32)() {
    m.discount = value
}
// SetFilesStored sets the FilesStored property value. The FilesStored property
func (m *StorageZone) SetFilesStored(value *int64)() {
    m.filesStored = value
}
// SetId sets the Id property value. The Id property
func (m *StorageZone) SetId(value *int64)() {
    m.id = value
}
// SetName sets the Name property value. The Name property
func (m *StorageZone) SetName(value *string)() {
    m.name = value
}
// SetPassword sets the Password property value. The Password property
func (m *StorageZone) SetPassword(value *string)() {
    m.password = value
}
// SetPriceOverride sets the PriceOverride property value. The PriceOverride property
func (m *StorageZone) SetPriceOverride(value *int64)() {
    m.priceOverride = value
}
// SetPullZones sets the PullZones property value. The PullZones property
func (m *StorageZone) SetPullZones(value []id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428.PullZoneable)() {
    m.pullZones = value
}
// SetReadOnlyPassword sets the ReadOnlyPassword property value. The ReadOnlyPassword property
func (m *StorageZone) SetReadOnlyPassword(value *string)() {
    m.readOnlyPassword = value
}
// SetRegion sets the Region property value. The Region property
func (m *StorageZone) SetRegion(value *StandardRegions)() {
    m.region = value
}
// SetReplicationChangeInProgress sets the ReplicationChangeInProgress property value. The ReplicationChangeInProgress property
func (m *StorageZone) SetReplicationChangeInProgress(value *bool)() {
    m.replicationChangeInProgress = value
}
// SetReplicationRegions sets the ReplicationRegions property value. The ReplicationRegions property
func (m *StorageZone) SetReplicationRegions(value []EdgeReplicationRegions)() {
    m.replicationRegions = value
}
// SetRewrite404To200 sets the Rewrite404To200 property value. The Rewrite404To200 property
func (m *StorageZone) SetRewrite404To200(value *bool)() {
    m.rewrite404To200 = value
}
// SetStorageHostname sets the StorageHostname property value. The StorageHostname property
func (m *StorageZone) SetStorageHostname(value *string)() {
    m.storageHostname = value
}
// SetStorageUsed sets the StorageUsed property value. The StorageUsed property
func (m *StorageZone) SetStorageUsed(value *int64)() {
    m.storageUsed = value
}
// SetUserId sets the UserId property value. The UserId property
func (m *StorageZone) SetUserId(value *string)() {
    m.userId = value
}
// SetZoneTier sets the ZoneTier property value. The ZoneTier property
func (m *StorageZone) SetZoneTier(value *float64)() {
    m.zoneTier = value
}
type StorageZoneable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustom404FilePath()(*string)
    GetDateModified()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeleted()(*bool)
    GetDiscount()(*int32)
    GetFilesStored()(*int64)
    GetId()(*int64)
    GetName()(*string)
    GetPassword()(*string)
    GetPriceOverride()(*int64)
    GetPullZones()([]id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428.PullZoneable)
    GetReadOnlyPassword()(*string)
    GetRegion()(*StandardRegions)
    GetReplicationChangeInProgress()(*bool)
    GetReplicationRegions()([]EdgeReplicationRegions)
    GetRewrite404To200()(*bool)
    GetStorageHostname()(*string)
    GetStorageUsed()(*int64)
    GetUserId()(*string)
    GetZoneTier()(*float64)
    SetCustom404FilePath(value *string)()
    SetDateModified(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeleted(value *bool)()
    SetDiscount(value *int32)()
    SetFilesStored(value *int64)()
    SetId(value *int64)()
    SetName(value *string)()
    SetPassword(value *string)()
    SetPriceOverride(value *int64)()
    SetPullZones(value []id1f978eb657ca81d9e4d7cee675850f471201dd71ed8d22f02520011358b9428.PullZoneable)()
    SetReadOnlyPassword(value *string)()
    SetRegion(value *StandardRegions)()
    SetReplicationChangeInProgress(value *bool)()
    SetReplicationRegions(value []EdgeReplicationRegions)()
    SetRewrite404To200(value *bool)()
    SetStorageHostname(value *string)()
    SetStorageUsed(value *int64)()
    SetUserId(value *string)()
    SetZoneTier(value *float64)()
}
