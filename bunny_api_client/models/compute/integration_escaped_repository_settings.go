package compute

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Integration_RepositorySettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Id property
    id *int64
    // The Name property
    name *string
    // The Private property
    private *bool
    // The TemplateUrl property
    templateUrl *string
}
// NewIntegration_RepositorySettings instantiates a new Integration_RepositorySettings and sets the default values.
func NewIntegration_RepositorySettings()(*Integration_RepositorySettings) {
    m := &Integration_RepositorySettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIntegration_RepositorySettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateIntegration_RepositorySettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIntegration_RepositorySettings(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Integration_RepositorySettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Integration_RepositorySettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["Private"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivate(val)
        }
        return nil
    }
    res["TemplateUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateUrl(val)
        }
        return nil
    }
    return res
}
// GetId gets the Id property value. The Id property
// returns a *int64 when successful
func (m *Integration_RepositorySettings) GetId()(*int64) {
    return m.id
}
// GetName gets the Name property value. The Name property
// returns a *string when successful
func (m *Integration_RepositorySettings) GetName()(*string) {
    return m.name
}
// GetPrivate gets the Private property value. The Private property
// returns a *bool when successful
func (m *Integration_RepositorySettings) GetPrivate()(*bool) {
    return m.private
}
// GetTemplateUrl gets the TemplateUrl property value. The TemplateUrl property
// returns a *string when successful
func (m *Integration_RepositorySettings) GetTemplateUrl()(*string) {
    return m.templateUrl
}
// Serialize serializes information the current object
func (m *Integration_RepositorySettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("Name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("Private", m.GetPrivate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("TemplateUrl", m.GetTemplateUrl())
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
func (m *Integration_RepositorySettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetId sets the Id property value. The Id property
func (m *Integration_RepositorySettings) SetId(value *int64)() {
    m.id = value
}
// SetName sets the Name property value. The Name property
func (m *Integration_RepositorySettings) SetName(value *string)() {
    m.name = value
}
// SetPrivate sets the Private property value. The Private property
func (m *Integration_RepositorySettings) SetPrivate(value *bool)() {
    m.private = value
}
// SetTemplateUrl sets the TemplateUrl property value. The TemplateUrl property
func (m *Integration_RepositorySettings) SetTemplateUrl(value *string)() {
    m.templateUrl = value
}
type Integration_RepositorySettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetId()(*int64)
    GetName()(*string)
    GetPrivate()(*bool)
    GetTemplateUrl()(*string)
    SetId(value *int64)()
    SetName(value *string)()
    SetPrivate(value *bool)()
    SetTemplateUrl(value *string)()
}
