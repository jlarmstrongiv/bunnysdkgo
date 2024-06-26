package pullzone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BunnyAiImageBlueprint struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Name property
    name *string
    // The Properties property
    properties BunnyAiImageBlueprint_Propertiesable
}
// NewBunnyAiImageBlueprint instantiates a new BunnyAiImageBlueprint and sets the default values.
func NewBunnyAiImageBlueprint()(*BunnyAiImageBlueprint) {
    m := &BunnyAiImageBlueprint{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBunnyAiImageBlueprintFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBunnyAiImageBlueprintFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBunnyAiImageBlueprint(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BunnyAiImageBlueprint) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BunnyAiImageBlueprint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["Properties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBunnyAiImageBlueprint_PropertiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProperties(val.(BunnyAiImageBlueprint_Propertiesable))
        }
        return nil
    }
    return res
}
// GetName gets the Name property value. The Name property
// returns a *string when successful
func (m *BunnyAiImageBlueprint) GetName()(*string) {
    return m.name
}
// GetProperties gets the Properties property value. The Properties property
// returns a BunnyAiImageBlueprint_Propertiesable when successful
func (m *BunnyAiImageBlueprint) GetProperties()(BunnyAiImageBlueprint_Propertiesable) {
    return m.properties
}
// Serialize serializes information the current object
func (m *BunnyAiImageBlueprint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("Name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("Properties", m.GetProperties())
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
func (m *BunnyAiImageBlueprint) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the Name property value. The Name property
func (m *BunnyAiImageBlueprint) SetName(value *string)() {
    m.name = value
}
// SetProperties sets the Properties property value. The Properties property
func (m *BunnyAiImageBlueprint) SetProperties(value BunnyAiImageBlueprint_Propertiesable)() {
    m.properties = value
}
type BunnyAiImageBlueprintable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetProperties()(BunnyAiImageBlueprint_Propertiesable)
    SetName(value *string)()
    SetProperties(value BunnyAiImageBlueprint_Propertiesable)()
}
