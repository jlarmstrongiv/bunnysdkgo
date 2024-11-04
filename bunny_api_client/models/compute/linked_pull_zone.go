package compute

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LinkedPullZone struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The DefaultHostname property
    defaultHostname *string
    // The Id property
    id *int64
    // The PullZoneName property
    pullZoneName *string
}
// NewLinkedPullZone instantiates a new LinkedPullZone and sets the default values.
func NewLinkedPullZone()(*LinkedPullZone) {
    m := &LinkedPullZone{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLinkedPullZoneFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLinkedPullZoneFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLinkedPullZone(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LinkedPullZone) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDefaultHostname gets the DefaultHostname property value. The DefaultHostname property
// returns a *string when successful
func (m *LinkedPullZone) GetDefaultHostname()(*string) {
    return m.defaultHostname
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LinkedPullZone) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["DefaultHostname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultHostname(val)
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
    return res
}
// GetId gets the Id property value. The Id property
// returns a *int64 when successful
func (m *LinkedPullZone) GetId()(*int64) {
    return m.id
}
// GetPullZoneName gets the PullZoneName property value. The PullZoneName property
// returns a *string when successful
func (m *LinkedPullZone) GetPullZoneName()(*string) {
    return m.pullZoneName
}
// Serialize serializes information the current object
func (m *LinkedPullZone) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("DefaultHostname", m.GetDefaultHostname())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("Id", m.GetId())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *LinkedPullZone) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDefaultHostname sets the DefaultHostname property value. The DefaultHostname property
func (m *LinkedPullZone) SetDefaultHostname(value *string)() {
    m.defaultHostname = value
}
// SetId sets the Id property value. The Id property
func (m *LinkedPullZone) SetId(value *int64)() {
    m.id = value
}
// SetPullZoneName sets the PullZoneName property value. The PullZoneName property
func (m *LinkedPullZone) SetPullZoneName(value *string)() {
    m.pullZoneName = value
}
type LinkedPullZoneable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultHostname()(*string)
    GetId()(*int64)
    GetPullZoneName()(*string)
    SetDefaultHostname(value *string)()
    SetId(value *int64)()
    SetPullZoneName(value *string)()
}
