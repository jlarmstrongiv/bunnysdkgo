package dnszone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPostRequestBody the template for adding optional properties.
type ItemPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The CustomNameserversEnabled property
    customNameserversEnabled *bool
    // The LogAnonymizationType property
    logAnonymizationType *float64
    // The LoggingEnabled property
    loggingEnabled *bool
    // Determines if the log anonymization should be enabled
    loggingIPAnonymizationEnabled *bool
    // The Nameserver1 property
    nameserver1 *string
    // The Nameserver2 property
    nameserver2 *string
    // The SoaEmail property
    soaEmail *string
}
// NewItemPostRequestBody instantiates a new ItemPostRequestBody and sets the default values.
func NewItemPostRequestBody()(*ItemPostRequestBody) {
    m := &ItemPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCustomNameserversEnabled gets the CustomNameserversEnabled property value. The CustomNameserversEnabled property
// returns a *bool when successful
func (m *ItemPostRequestBody) GetCustomNameserversEnabled()(*bool) {
    return m.customNameserversEnabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["CustomNameserversEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomNameserversEnabled(val)
        }
        return nil
    }
    res["LogAnonymizationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogAnonymizationType(val)
        }
        return nil
    }
    res["LoggingEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoggingEnabled(val)
        }
        return nil
    }
    res["LoggingIPAnonymizationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLoggingIPAnonymizationEnabled(val)
        }
        return nil
    }
    res["Nameserver1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNameserver1(val)
        }
        return nil
    }
    res["Nameserver2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNameserver2(val)
        }
        return nil
    }
    res["SoaEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoaEmail(val)
        }
        return nil
    }
    return res
}
// GetLogAnonymizationType gets the LogAnonymizationType property value. The LogAnonymizationType property
// returns a *float64 when successful
func (m *ItemPostRequestBody) GetLogAnonymizationType()(*float64) {
    return m.logAnonymizationType
}
// GetLoggingEnabled gets the LoggingEnabled property value. The LoggingEnabled property
// returns a *bool when successful
func (m *ItemPostRequestBody) GetLoggingEnabled()(*bool) {
    return m.loggingEnabled
}
// GetLoggingIPAnonymizationEnabled gets the LoggingIPAnonymizationEnabled property value. Determines if the log anonymization should be enabled
// returns a *bool when successful
func (m *ItemPostRequestBody) GetLoggingIPAnonymizationEnabled()(*bool) {
    return m.loggingIPAnonymizationEnabled
}
// GetNameserver1 gets the Nameserver1 property value. The Nameserver1 property
// returns a *string when successful
func (m *ItemPostRequestBody) GetNameserver1()(*string) {
    return m.nameserver1
}
// GetNameserver2 gets the Nameserver2 property value. The Nameserver2 property
// returns a *string when successful
func (m *ItemPostRequestBody) GetNameserver2()(*string) {
    return m.nameserver2
}
// GetSoaEmail gets the SoaEmail property value. The SoaEmail property
// returns a *string when successful
func (m *ItemPostRequestBody) GetSoaEmail()(*string) {
    return m.soaEmail
}
// Serialize serializes information the current object
func (m *ItemPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("CustomNameserversEnabled", m.GetCustomNameserversEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("LogAnonymizationType", m.GetLogAnonymizationType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("LoggingEnabled", m.GetLoggingEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("LoggingIPAnonymizationEnabled", m.GetLoggingIPAnonymizationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Nameserver1", m.GetNameserver1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Nameserver2", m.GetNameserver2())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("SoaEmail", m.GetSoaEmail())
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
func (m *ItemPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCustomNameserversEnabled sets the CustomNameserversEnabled property value. The CustomNameserversEnabled property
func (m *ItemPostRequestBody) SetCustomNameserversEnabled(value *bool)() {
    m.customNameserversEnabled = value
}
// SetLogAnonymizationType sets the LogAnonymizationType property value. The LogAnonymizationType property
func (m *ItemPostRequestBody) SetLogAnonymizationType(value *float64)() {
    m.logAnonymizationType = value
}
// SetLoggingEnabled sets the LoggingEnabled property value. The LoggingEnabled property
func (m *ItemPostRequestBody) SetLoggingEnabled(value *bool)() {
    m.loggingEnabled = value
}
// SetLoggingIPAnonymizationEnabled sets the LoggingIPAnonymizationEnabled property value. Determines if the log anonymization should be enabled
func (m *ItemPostRequestBody) SetLoggingIPAnonymizationEnabled(value *bool)() {
    m.loggingIPAnonymizationEnabled = value
}
// SetNameserver1 sets the Nameserver1 property value. The Nameserver1 property
func (m *ItemPostRequestBody) SetNameserver1(value *string)() {
    m.nameserver1 = value
}
// SetNameserver2 sets the Nameserver2 property value. The Nameserver2 property
func (m *ItemPostRequestBody) SetNameserver2(value *string)() {
    m.nameserver2 = value
}
// SetSoaEmail sets the SoaEmail property value. The SoaEmail property
func (m *ItemPostRequestBody) SetSoaEmail(value *string)() {
    m.soaEmail = value
}
type ItemPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCustomNameserversEnabled()(*bool)
    GetLogAnonymizationType()(*float64)
    GetLoggingEnabled()(*bool)
    GetLoggingIPAnonymizationEnabled()(*bool)
    GetNameserver1()(*string)
    GetNameserver2()(*string)
    GetSoaEmail()(*string)
    SetCustomNameserversEnabled(value *bool)()
    SetLogAnonymizationType(value *float64)()
    SetLoggingEnabled(value *bool)()
    SetLoggingIPAnonymizationEnabled(value *bool)()
    SetNameserver1(value *string)()
    SetNameserver2(value *string)()
    SetSoaEmail(value *string)()
}
