package videolibrary

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemVideolibraryPostRequestBody_GoogleWidevineDrm struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The CertificateExpirationDate property
    certificateExpirationDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The CertificateId property
    certificateId *int64
    // The Enabled property
    enabled *bool
    // The MinClientSecurityLevel property
    minClientSecurityLevel *float64
    // The Provider property
    provider *string
    // The SdOnlyForL3 property
    sdOnlyForL3 *bool
}
// NewItemVideolibraryPostRequestBody_GoogleWidevineDrm instantiates a new ItemVideolibraryPostRequestBody_GoogleWidevineDrm and sets the default values.
func NewItemVideolibraryPostRequestBody_GoogleWidevineDrm()(*ItemVideolibraryPostRequestBody_GoogleWidevineDrm) {
    m := &ItemVideolibraryPostRequestBody_GoogleWidevineDrm{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemVideolibraryPostRequestBody_GoogleWidevineDrmFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemVideolibraryPostRequestBody_GoogleWidevineDrmFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemVideolibraryPostRequestBody_GoogleWidevineDrm(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemVideolibraryPostRequestBody_GoogleWidevineDrm) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCertificateExpirationDate gets the CertificateExpirationDate property value. The CertificateExpirationDate property
// returns a *Time when successful
func (m *ItemVideolibraryPostRequestBody_GoogleWidevineDrm) GetCertificateExpirationDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.certificateExpirationDate
}
// GetCertificateId gets the CertificateId property value. The CertificateId property
// returns a *int64 when successful
func (m *ItemVideolibraryPostRequestBody_GoogleWidevineDrm) GetCertificateId()(*int64) {
    return m.certificateId
}
// GetEnabled gets the Enabled property value. The Enabled property
// returns a *bool when successful
func (m *ItemVideolibraryPostRequestBody_GoogleWidevineDrm) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemVideolibraryPostRequestBody_GoogleWidevineDrm) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["CertificateExpirationDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateExpirationDate(val)
        }
        return nil
    }
    res["CertificateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateId(val)
        }
        return nil
    }
    res["Enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["MinClientSecurityLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinClientSecurityLevel(val)
        }
        return nil
    }
    res["Provider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvider(val)
        }
        return nil
    }
    res["SdOnlyForL3"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSdOnlyForL3(val)
        }
        return nil
    }
    return res
}
// GetMinClientSecurityLevel gets the MinClientSecurityLevel property value. The MinClientSecurityLevel property
// returns a *float64 when successful
func (m *ItemVideolibraryPostRequestBody_GoogleWidevineDrm) GetMinClientSecurityLevel()(*float64) {
    return m.minClientSecurityLevel
}
// GetProvider gets the Provider property value. The Provider property
// returns a *string when successful
func (m *ItemVideolibraryPostRequestBody_GoogleWidevineDrm) GetProvider()(*string) {
    return m.provider
}
// GetSdOnlyForL3 gets the SdOnlyForL3 property value. The SdOnlyForL3 property
// returns a *bool when successful
func (m *ItemVideolibraryPostRequestBody_GoogleWidevineDrm) GetSdOnlyForL3()(*bool) {
    return m.sdOnlyForL3
}
// Serialize serializes information the current object
func (m *ItemVideolibraryPostRequestBody_GoogleWidevineDrm) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("CertificateExpirationDate", m.GetCertificateExpirationDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("CertificateId", m.GetCertificateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("Enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("MinClientSecurityLevel", m.GetMinClientSecurityLevel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Provider", m.GetProvider())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("SdOnlyForL3", m.GetSdOnlyForL3())
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
func (m *ItemVideolibraryPostRequestBody_GoogleWidevineDrm) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCertificateExpirationDate sets the CertificateExpirationDate property value. The CertificateExpirationDate property
func (m *ItemVideolibraryPostRequestBody_GoogleWidevineDrm) SetCertificateExpirationDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.certificateExpirationDate = value
}
// SetCertificateId sets the CertificateId property value. The CertificateId property
func (m *ItemVideolibraryPostRequestBody_GoogleWidevineDrm) SetCertificateId(value *int64)() {
    m.certificateId = value
}
// SetEnabled sets the Enabled property value. The Enabled property
func (m *ItemVideolibraryPostRequestBody_GoogleWidevineDrm) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetMinClientSecurityLevel sets the MinClientSecurityLevel property value. The MinClientSecurityLevel property
func (m *ItemVideolibraryPostRequestBody_GoogleWidevineDrm) SetMinClientSecurityLevel(value *float64)() {
    m.minClientSecurityLevel = value
}
// SetProvider sets the Provider property value. The Provider property
func (m *ItemVideolibraryPostRequestBody_GoogleWidevineDrm) SetProvider(value *string)() {
    m.provider = value
}
// SetSdOnlyForL3 sets the SdOnlyForL3 property value. The SdOnlyForL3 property
func (m *ItemVideolibraryPostRequestBody_GoogleWidevineDrm) SetSdOnlyForL3(value *bool)() {
    m.sdOnlyForL3 = value
}
type ItemVideolibraryPostRequestBody_GoogleWidevineDrmable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateExpirationDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCertificateId()(*int64)
    GetEnabled()(*bool)
    GetMinClientSecurityLevel()(*float64)
    GetProvider()(*string)
    GetSdOnlyForL3()(*bool)
    SetCertificateExpirationDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCertificateId(value *int64)()
    SetEnabled(value *bool)()
    SetMinClientSecurityLevel(value *float64)()
    SetProvider(value *string)()
    SetSdOnlyForL3(value *bool)()
}
