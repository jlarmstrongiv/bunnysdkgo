package pullzone

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemAddcertificateAddCertificatePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Base64Url encoded binary data of the certificate file
    certificate *string
    // The Base64Url encoded binary data of the certificate key file
    certificateKey *string
    // The hostname to which the certificate will be added
    hostname *string
}
// NewItemAddcertificateAddCertificatePostRequestBody instantiates a new ItemAddcertificateAddCertificatePostRequestBody and sets the default values.
func NewItemAddcertificateAddCertificatePostRequestBody()(*ItemAddcertificateAddCertificatePostRequestBody) {
    m := &ItemAddcertificateAddCertificatePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemAddcertificateAddCertificatePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemAddcertificateAddCertificatePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemAddcertificateAddCertificatePostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemAddcertificateAddCertificatePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCertificate gets the Certificate property value. The Base64Url encoded binary data of the certificate file
// returns a *string when successful
func (m *ItemAddcertificateAddCertificatePostRequestBody) GetCertificate()(*string) {
    return m.certificate
}
// GetCertificateKey gets the CertificateKey property value. The Base64Url encoded binary data of the certificate key file
// returns a *string when successful
func (m *ItemAddcertificateAddCertificatePostRequestBody) GetCertificateKey()(*string) {
    return m.certificateKey
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemAddcertificateAddCertificatePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["Certificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificate(val)
        }
        return nil
    }
    res["CertificateKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateKey(val)
        }
        return nil
    }
    res["Hostname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHostname(val)
        }
        return nil
    }
    return res
}
// GetHostname gets the Hostname property value. The hostname to which the certificate will be added
// returns a *string when successful
func (m *ItemAddcertificateAddCertificatePostRequestBody) GetHostname()(*string) {
    return m.hostname
}
// Serialize serializes information the current object
func (m *ItemAddcertificateAddCertificatePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("Certificate", m.GetCertificate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("CertificateKey", m.GetCertificateKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("Hostname", m.GetHostname())
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
func (m *ItemAddcertificateAddCertificatePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCertificate sets the Certificate property value. The Base64Url encoded binary data of the certificate file
func (m *ItemAddcertificateAddCertificatePostRequestBody) SetCertificate(value *string)() {
    m.certificate = value
}
// SetCertificateKey sets the CertificateKey property value. The Base64Url encoded binary data of the certificate key file
func (m *ItemAddcertificateAddCertificatePostRequestBody) SetCertificateKey(value *string)() {
    m.certificateKey = value
}
// SetHostname sets the Hostname property value. The hostname to which the certificate will be added
func (m *ItemAddcertificateAddCertificatePostRequestBody) SetHostname(value *string)() {
    m.hostname = value
}
type ItemAddcertificateAddCertificatePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCertificate()(*string)
    GetCertificateKey()(*string)
    GetHostname()(*string)
    SetCertificate(value *string)()
    SetCertificateKey(value *string)()
    SetHostname(value *string)()
}
