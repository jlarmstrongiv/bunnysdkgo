package pullzone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BunnyAiImageBlueprint_Properties struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Cfg property
    cfg *string
    // The NegativePrompt property
    negativePrompt *string
    // The PostPrompt property
    postPrompt *string
    // The PrePrompt property
    prePrompt *string
    // The Steps property
    steps *string
}
// NewBunnyAiImageBlueprint_Properties instantiates a new BunnyAiImageBlueprint_Properties and sets the default values.
func NewBunnyAiImageBlueprint_Properties()(*BunnyAiImageBlueprint_Properties) {
    m := &BunnyAiImageBlueprint_Properties{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBunnyAiImageBlueprint_PropertiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBunnyAiImageBlueprint_PropertiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBunnyAiImageBlueprint_Properties(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BunnyAiImageBlueprint_Properties) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCfg gets the Cfg property value. The Cfg property
// returns a *string when successful
func (m *BunnyAiImageBlueprint_Properties) GetCfg()(*string) {
    return m.cfg
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BunnyAiImageBlueprint_Properties) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["Cfg"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCfg(val)
        }
        return nil
    }
    res["NegativePrompt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNegativePrompt(val)
        }
        return nil
    }
    res["PostPrompt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostPrompt(val)
        }
        return nil
    }
    res["PrePrompt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrePrompt(val)
        }
        return nil
    }
    res["Steps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSteps(val)
        }
        return nil
    }
    return res
}
// GetNegativePrompt gets the NegativePrompt property value. The NegativePrompt property
// returns a *string when successful
func (m *BunnyAiImageBlueprint_Properties) GetNegativePrompt()(*string) {
    return m.negativePrompt
}
// GetPostPrompt gets the PostPrompt property value. The PostPrompt property
// returns a *string when successful
func (m *BunnyAiImageBlueprint_Properties) GetPostPrompt()(*string) {
    return m.postPrompt
}
// GetPrePrompt gets the PrePrompt property value. The PrePrompt property
// returns a *string when successful
func (m *BunnyAiImageBlueprint_Properties) GetPrePrompt()(*string) {
    return m.prePrompt
}
// GetSteps gets the Steps property value. The Steps property
// returns a *string when successful
func (m *BunnyAiImageBlueprint_Properties) GetSteps()(*string) {
    return m.steps
}
// Serialize serializes information the current object
func (m *BunnyAiImageBlueprint_Properties) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("Cfg", m.GetCfg())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("NegativePrompt", m.GetNegativePrompt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("PostPrompt", m.GetPostPrompt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("PrePrompt", m.GetPrePrompt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Steps", m.GetSteps())
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
func (m *BunnyAiImageBlueprint_Properties) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCfg sets the Cfg property value. The Cfg property
func (m *BunnyAiImageBlueprint_Properties) SetCfg(value *string)() {
    m.cfg = value
}
// SetNegativePrompt sets the NegativePrompt property value. The NegativePrompt property
func (m *BunnyAiImageBlueprint_Properties) SetNegativePrompt(value *string)() {
    m.negativePrompt = value
}
// SetPostPrompt sets the PostPrompt property value. The PostPrompt property
func (m *BunnyAiImageBlueprint_Properties) SetPostPrompt(value *string)() {
    m.postPrompt = value
}
// SetPrePrompt sets the PrePrompt property value. The PrePrompt property
func (m *BunnyAiImageBlueprint_Properties) SetPrePrompt(value *string)() {
    m.prePrompt = value
}
// SetSteps sets the Steps property value. The Steps property
func (m *BunnyAiImageBlueprint_Properties) SetSteps(value *string)() {
    m.steps = value
}
type BunnyAiImageBlueprint_Propertiesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCfg()(*string)
    GetNegativePrompt()(*string)
    GetPostPrompt()(*string)
    GetPrePrompt()(*string)
    GetSteps()(*string)
    SetCfg(value *string)()
    SetNegativePrompt(value *string)()
    SetPostPrompt(value *string)()
    SetPrePrompt(value *string)()
    SetSteps(value *string)()
}
