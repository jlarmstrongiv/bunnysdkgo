package shield

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UpdateShieldZoneRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The shieldZone property
    shieldZone ShieldZoneRequestable
    // The shieldZoneId property
    shieldZoneId *int64
}
// NewUpdateShieldZoneRequest instantiates a new UpdateShieldZoneRequest and sets the default values.
func NewUpdateShieldZoneRequest()(*UpdateShieldZoneRequest) {
    m := &UpdateShieldZoneRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUpdateShieldZoneRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUpdateShieldZoneRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUpdateShieldZoneRequest(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *UpdateShieldZoneRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UpdateShieldZoneRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["shieldZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateShieldZoneRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShieldZone(val.(ShieldZoneRequestable))
        }
        return nil
    }
    res["shieldZoneId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShieldZoneId(val)
        }
        return nil
    }
    return res
}
// GetShieldZone gets the shieldZone property value. The shieldZone property
// returns a ShieldZoneRequestable when successful
func (m *UpdateShieldZoneRequest) GetShieldZone()(ShieldZoneRequestable) {
    return m.shieldZone
}
// GetShieldZoneId gets the shieldZoneId property value. The shieldZoneId property
// returns a *int64 when successful
func (m *UpdateShieldZoneRequest) GetShieldZoneId()(*int64) {
    return m.shieldZoneId
}
// Serialize serializes information the current object
func (m *UpdateShieldZoneRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("shieldZone", m.GetShieldZone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("shieldZoneId", m.GetShieldZoneId())
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
func (m *UpdateShieldZoneRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetShieldZone sets the shieldZone property value. The shieldZone property
func (m *UpdateShieldZoneRequest) SetShieldZone(value ShieldZoneRequestable)() {
    m.shieldZone = value
}
// SetShieldZoneId sets the shieldZoneId property value. The shieldZoneId property
func (m *UpdateShieldZoneRequest) SetShieldZoneId(value *int64)() {
    m.shieldZoneId = value
}
type UpdateShieldZoneRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetShieldZone()(ShieldZoneRequestable)
    GetShieldZoneId()(*int64)
    SetShieldZone(value ShieldZoneRequestable)()
    SetShieldZoneId(value *int64)()
}
