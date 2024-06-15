package getlanguages

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Language struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Name property
    name *string
    // The ShortCode property
    shortCode *string
    // The SupportPlayerTranslation property
    supportPlayerTranslation *bool
    // The SupportTranscribing property
    supportTranscribing *bool
    // The TranscribingAccuracy property
    transcribingAccuracy *int32
}
// NewLanguage instantiates a new Language and sets the default values.
func NewLanguage()(*Language) {
    m := &Language{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLanguageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLanguageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLanguage(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Language) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Language) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["ShortCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShortCode(val)
        }
        return nil
    }
    res["SupportPlayerTranslation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportPlayerTranslation(val)
        }
        return nil
    }
    res["SupportTranscribing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportTranscribing(val)
        }
        return nil
    }
    res["TranscribingAccuracy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTranscribingAccuracy(val)
        }
        return nil
    }
    return res
}
// GetName gets the Name property value. The Name property
// returns a *string when successful
func (m *Language) GetName()(*string) {
    return m.name
}
// GetShortCode gets the ShortCode property value. The ShortCode property
// returns a *string when successful
func (m *Language) GetShortCode()(*string) {
    return m.shortCode
}
// GetSupportPlayerTranslation gets the SupportPlayerTranslation property value. The SupportPlayerTranslation property
// returns a *bool when successful
func (m *Language) GetSupportPlayerTranslation()(*bool) {
    return m.supportPlayerTranslation
}
// GetSupportTranscribing gets the SupportTranscribing property value. The SupportTranscribing property
// returns a *bool when successful
func (m *Language) GetSupportTranscribing()(*bool) {
    return m.supportTranscribing
}
// GetTranscribingAccuracy gets the TranscribingAccuracy property value. The TranscribingAccuracy property
// returns a *int32 when successful
func (m *Language) GetTranscribingAccuracy()(*int32) {
    return m.transcribingAccuracy
}
// Serialize serializes information the current object
func (m *Language) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("Name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ShortCode", m.GetShortCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("SupportPlayerTranslation", m.GetSupportPlayerTranslation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("SupportTranscribing", m.GetSupportTranscribing())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("TranscribingAccuracy", m.GetTranscribingAccuracy())
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
func (m *Language) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the Name property value. The Name property
func (m *Language) SetName(value *string)() {
    m.name = value
}
// SetShortCode sets the ShortCode property value. The ShortCode property
func (m *Language) SetShortCode(value *string)() {
    m.shortCode = value
}
// SetSupportPlayerTranslation sets the SupportPlayerTranslation property value. The SupportPlayerTranslation property
func (m *Language) SetSupportPlayerTranslation(value *bool)() {
    m.supportPlayerTranslation = value
}
// SetSupportTranscribing sets the SupportTranscribing property value. The SupportTranscribing property
func (m *Language) SetSupportTranscribing(value *bool)() {
    m.supportTranscribing = value
}
// SetTranscribingAccuracy sets the TranscribingAccuracy property value. The TranscribingAccuracy property
func (m *Language) SetTranscribingAccuracy(value *int32)() {
    m.transcribingAccuracy = value
}
type Languageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetShortCode()(*string)
    GetSupportPlayerTranslation()(*bool)
    GetSupportTranscribing()(*bool)
    GetTranscribingAccuracy()(*int32)
    SetName(value *string)()
    SetShortCode(value *string)()
    SetSupportPlayerTranslation(value *bool)()
    SetSupportTranscribing(value *bool)()
    SetTranscribingAccuracy(value *int32)()
}
