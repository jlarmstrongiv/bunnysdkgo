package getstoragezoneconnections

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Connection struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ConnectionType property
    connectionType *float64
    // The main custom connected CDN domain
    mainCustomDomain *string
    // The total amount of bandwidth served by this zone this month
    monthlyBandwidthUsed *float64
    // The total monthly charges incurred by this zone
    monthlyCharges *float64
    // The ID of the connected pull zone
    pullZoneId *int64
    // The name of the connected pull zone
    pullZoneName *string
    // The Tier property
    tier *float64
}
// NewConnection instantiates a new Connection and sets the default values.
func NewConnection()(*Connection) {
    m := &Connection{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConnectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateConnectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConnection(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Connection) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConnectionType gets the ConnectionType property value. The ConnectionType property
// returns a *float64 when successful
func (m *Connection) GetConnectionType()(*float64) {
    return m.connectionType
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Connection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ConnectionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionType(val)
        }
        return nil
    }
    res["MainCustomDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMainCustomDomain(val)
        }
        return nil
    }
    res["MonthlyBandwidthUsed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonthlyBandwidthUsed(val)
        }
        return nil
    }
    res["MonthlyCharges"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonthlyCharges(val)
        }
        return nil
    }
    res["PullZoneId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullZoneId(val)
        }
        return nil
    }
    res["PullZoneName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPullZoneName(val)
        }
        return nil
    }
    res["Tier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTier(val)
        }
        return nil
    }
    return res
}
// GetMainCustomDomain gets the MainCustomDomain property value. The main custom connected CDN domain
// returns a *string when successful
func (m *Connection) GetMainCustomDomain()(*string) {
    return m.mainCustomDomain
}
// GetMonthlyBandwidthUsed gets the MonthlyBandwidthUsed property value. The total amount of bandwidth served by this zone this month
// returns a *float64 when successful
func (m *Connection) GetMonthlyBandwidthUsed()(*float64) {
    return m.monthlyBandwidthUsed
}
// GetMonthlyCharges gets the MonthlyCharges property value. The total monthly charges incurred by this zone
// returns a *float64 when successful
func (m *Connection) GetMonthlyCharges()(*float64) {
    return m.monthlyCharges
}
// GetPullZoneId gets the PullZoneId property value. The ID of the connected pull zone
// returns a *int64 when successful
func (m *Connection) GetPullZoneId()(*int64) {
    return m.pullZoneId
}
// GetPullZoneName gets the PullZoneName property value. The name of the connected pull zone
// returns a *string when successful
func (m *Connection) GetPullZoneName()(*string) {
    return m.pullZoneName
}
// GetTier gets the Tier property value. The Tier property
// returns a *float64 when successful
func (m *Connection) GetTier()(*float64) {
    return m.tier
}
// Serialize serializes information the current object
func (m *Connection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("ConnectionType", m.GetConnectionType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("MainCustomDomain", m.GetMainCustomDomain())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("MonthlyBandwidthUsed", m.GetMonthlyBandwidthUsed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("MonthlyCharges", m.GetMonthlyCharges())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("PullZoneId", m.GetPullZoneId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("PullZoneName", m.GetPullZoneName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("Tier", m.GetTier())
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
func (m *Connection) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConnectionType sets the ConnectionType property value. The ConnectionType property
func (m *Connection) SetConnectionType(value *float64)() {
    m.connectionType = value
}
// SetMainCustomDomain sets the MainCustomDomain property value. The main custom connected CDN domain
func (m *Connection) SetMainCustomDomain(value *string)() {
    m.mainCustomDomain = value
}
// SetMonthlyBandwidthUsed sets the MonthlyBandwidthUsed property value. The total amount of bandwidth served by this zone this month
func (m *Connection) SetMonthlyBandwidthUsed(value *float64)() {
    m.monthlyBandwidthUsed = value
}
// SetMonthlyCharges sets the MonthlyCharges property value. The total monthly charges incurred by this zone
func (m *Connection) SetMonthlyCharges(value *float64)() {
    m.monthlyCharges = value
}
// SetPullZoneId sets the PullZoneId property value. The ID of the connected pull zone
func (m *Connection) SetPullZoneId(value *int64)() {
    m.pullZoneId = value
}
// SetPullZoneName sets the PullZoneName property value. The name of the connected pull zone
func (m *Connection) SetPullZoneName(value *string)() {
    m.pullZoneName = value
}
// SetTier sets the Tier property value. The Tier property
func (m *Connection) SetTier(value *float64)() {
    m.tier = value
}
type Connectionable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConnectionType()(*float64)
    GetMainCustomDomain()(*string)
    GetMonthlyBandwidthUsed()(*float64)
    GetMonthlyCharges()(*float64)
    GetPullZoneId()(*int64)
    GetPullZoneName()(*string)
    GetTier()(*float64)
    SetConnectionType(value *float64)()
    SetMainCustomDomain(value *string)()
    SetMonthlyBandwidthUsed(value *float64)()
    SetMonthlyCharges(value *float64)()
    SetPullZoneId(value *int64)()
    SetPullZoneName(value *string)()
    SetTier(value *float64)()
}
