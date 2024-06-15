package storagezone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type StorageZoneCreate struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Custom404FilePath property
    custom404FilePath *string
    // The Name property
    name *string
    // The OriginUrl property
    originUrl *string
    // The Region property
    region *StandardRegions
    // The ReplicationRegions property
    replicationRegions []EdgeReplicationRegions
    // The Rewrite404To200 property
    rewrite404To200 *bool
    // The ZoneTier property
    zoneTier *float64
}
// NewStorageZoneCreate instantiates a new StorageZoneCreate and sets the default values.
func NewStorageZoneCreate()(*StorageZoneCreate) {
    m := &StorageZoneCreate{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateStorageZoneCreateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateStorageZoneCreateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewStorageZoneCreate(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *StorageZoneCreate) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustom404FilePath gets the Custom404FilePath property value. The Custom404FilePath property
// returns a *string when successful
func (m *StorageZoneCreate) GetCustom404FilePath()(*string) {
    return m.custom404FilePath
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *StorageZoneCreate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["OriginUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginUrl(val)
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
// GetName gets the Name property value. The Name property
// returns a *string when successful
func (m *StorageZoneCreate) GetName()(*string) {
    return m.name
}
// GetOriginUrl gets the OriginUrl property value. The OriginUrl property
// returns a *string when successful
func (m *StorageZoneCreate) GetOriginUrl()(*string) {
    return m.originUrl
}
// GetRegion gets the Region property value. The Region property
// returns a *StandardRegions when successful
func (m *StorageZoneCreate) GetRegion()(*StandardRegions) {
    return m.region
}
// GetReplicationRegions gets the ReplicationRegions property value. The ReplicationRegions property
// returns a []EdgeReplicationRegions when successful
func (m *StorageZoneCreate) GetReplicationRegions()([]EdgeReplicationRegions) {
    return m.replicationRegions
}
// GetRewrite404To200 gets the Rewrite404To200 property value. The Rewrite404To200 property
// returns a *bool when successful
func (m *StorageZoneCreate) GetRewrite404To200()(*bool) {
    return m.rewrite404To200
}
// GetZoneTier gets the ZoneTier property value. The ZoneTier property
// returns a *float64 when successful
func (m *StorageZoneCreate) GetZoneTier()(*float64) {
    return m.zoneTier
}
// Serialize serializes information the current object
func (m *StorageZoneCreate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err := writer.WriteStringValue("OriginUrl", m.GetOriginUrl())
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
func (m *StorageZoneCreate) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustom404FilePath sets the Custom404FilePath property value. The Custom404FilePath property
func (m *StorageZoneCreate) SetCustom404FilePath(value *string)() {
    m.custom404FilePath = value
}
// SetName sets the Name property value. The Name property
func (m *StorageZoneCreate) SetName(value *string)() {
    m.name = value
}
// SetOriginUrl sets the OriginUrl property value. The OriginUrl property
func (m *StorageZoneCreate) SetOriginUrl(value *string)() {
    m.originUrl = value
}
// SetRegion sets the Region property value. The Region property
func (m *StorageZoneCreate) SetRegion(value *StandardRegions)() {
    m.region = value
}
// SetReplicationRegions sets the ReplicationRegions property value. The ReplicationRegions property
func (m *StorageZoneCreate) SetReplicationRegions(value []EdgeReplicationRegions)() {
    m.replicationRegions = value
}
// SetRewrite404To200 sets the Rewrite404To200 property value. The Rewrite404To200 property
func (m *StorageZoneCreate) SetRewrite404To200(value *bool)() {
    m.rewrite404To200 = value
}
// SetZoneTier sets the ZoneTier property value. The ZoneTier property
func (m *StorageZoneCreate) SetZoneTier(value *float64)() {
    m.zoneTier = value
}
type StorageZoneCreateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustom404FilePath()(*string)
    GetName()(*string)
    GetOriginUrl()(*string)
    GetRegion()(*StandardRegions)
    GetReplicationRegions()([]EdgeReplicationRegions)
    GetRewrite404To200()(*bool)
    GetZoneTier()(*float64)
    SetCustom404FilePath(value *string)()
    SetName(value *string)()
    SetOriginUrl(value *string)()
    SetRegion(value *StandardRegions)()
    SetReplicationRegions(value []EdgeReplicationRegions)()
    SetRewrite404To200(value *bool)()
    SetZoneTier(value *float64)()
}
