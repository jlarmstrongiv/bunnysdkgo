package videolibrary

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemVideolibraryPostRequestBody_AppleFairPlayDrm struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The CertificateExpirationDate property
    certificateExpirationDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The CertificateId property
    certificateId *int64
    // The Enabled property
    enabled *bool
    // The Provider property
    provider *string
}
// NewItemVideolibraryPostRequestBody_AppleFairPlayDrm instantiates a new ItemVideolibraryPostRequestBody_AppleFairPlayDrm and sets the default values.
func NewItemVideolibraryPostRequestBody_AppleFairPlayDrm()(*ItemVideolibraryPostRequestBody_AppleFairPlayDrm) {
    m := &ItemVideolibraryPostRequestBody_AppleFairPlayDrm{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemVideolibraryPostRequestBody_AppleFairPlayDrmFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemVideolibraryPostRequestBody_AppleFairPlayDrmFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemVideolibraryPostRequestBody_AppleFairPlayDrm(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemVideolibraryPostRequestBody_AppleFairPlayDrm) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCertificateExpirationDate gets the CertificateExpirationDate property value. The CertificateExpirationDate property
// returns a *Time when successful
func (m *ItemVideolibraryPostRequestBody_AppleFairPlayDrm) GetCertificateExpirationDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.certificateExpirationDate
}
// GetCertificateId gets the CertificateId property value. The CertificateId property
// returns a *int64 when successful
func (m *ItemVideolibraryPostRequestBody_AppleFairPlayDrm) GetCertificateId()(*int64) {
    return m.certificateId
}
// GetEnabled gets the Enabled property value. The Enabled property
// returns a *bool when successful
func (m *ItemVideolibraryPostRequestBody_AppleFairPlayDrm) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemVideolibraryPostRequestBody_AppleFairPlayDrm) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetProvider gets the Provider property value. The Provider property
// returns a *string when successful
func (m *ItemVideolibraryPostRequestBody_AppleFairPlayDrm) GetProvider()(*string) {
    return m.provider
}
// Serialize serializes information the current object
func (m *ItemVideolibraryPostRequestBody_AppleFairPlayDrm) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("Provider", m.GetProvider())
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
func (m *ItemVideolibraryPostRequestBody_AppleFairPlayDrm) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCertificateExpirationDate sets the CertificateExpirationDate property value. The CertificateExpirationDate property
func (m *ItemVideolibraryPostRequestBody_AppleFairPlayDrm) SetCertificateExpirationDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.certificateExpirationDate = value
}
// SetCertificateId sets the CertificateId property value. The CertificateId property
func (m *ItemVideolibraryPostRequestBody_AppleFairPlayDrm) SetCertificateId(value *int64)() {
    m.certificateId = value
}
// SetEnabled sets the Enabled property value. The Enabled property
func (m *ItemVideolibraryPostRequestBody_AppleFairPlayDrm) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetProvider sets the Provider property value. The Provider property
func (m *ItemVideolibraryPostRequestBody_AppleFairPlayDrm) SetProvider(value *string)() {
    m.provider = value
}
type ItemVideolibraryPostRequestBody_AppleFairPlayDrmable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificateExpirationDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCertificateId()(*int64)
    GetEnabled()(*bool)
    GetProvider()(*string)
    SetCertificateExpirationDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCertificateId(value *int64)()
    SetEnabled(value *bool)()
    SetProvider(value *string)()
}
