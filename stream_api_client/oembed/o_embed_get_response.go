package oembed

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OEmbedGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The height property
    height *int32
    // The html property
    html *string
    // The providerName property
    providerName *string
    // The providerUrl property
    providerUrl *string
    // The thumbnailUrl property
    thumbnailUrl *string
    // The title property
    title *string
    // The type property
    typeEscaped *string
    // The version property
    version *string
    // The width property
    width *int32
}
// NewOEmbedGetResponse instantiates a new OEmbedGetResponse and sets the default values.
func NewOEmbedGetResponse()(*OEmbedGetResponse) {
    m := &OEmbedGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOEmbedGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOEmbedGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOEmbedGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OEmbedGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OEmbedGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["height"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeight(val)
        }
        return nil
    }
    res["html"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHtml(val)
        }
        return nil
    }
    res["providerName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderName(val)
        }
        return nil
    }
    res["providerUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProviderUrl(val)
        }
        return nil
    }
    res["thumbnailUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThumbnailUrl(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    res["width"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWidth(val)
        }
        return nil
    }
    return res
}
// GetHeight gets the height property value. The height property
// returns a *int32 when successful
func (m *OEmbedGetResponse) GetHeight()(*int32) {
    return m.height
}
// GetHtml gets the html property value. The html property
// returns a *string when successful
func (m *OEmbedGetResponse) GetHtml()(*string) {
    return m.html
}
// GetProviderName gets the providerName property value. The providerName property
// returns a *string when successful
func (m *OEmbedGetResponse) GetProviderName()(*string) {
    return m.providerName
}
// GetProviderUrl gets the providerUrl property value. The providerUrl property
// returns a *string when successful
func (m *OEmbedGetResponse) GetProviderUrl()(*string) {
    return m.providerUrl
}
// GetThumbnailUrl gets the thumbnailUrl property value. The thumbnailUrl property
// returns a *string when successful
func (m *OEmbedGetResponse) GetThumbnailUrl()(*string) {
    return m.thumbnailUrl
}
// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *OEmbedGetResponse) GetTitle()(*string) {
    return m.title
}
// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *OEmbedGetResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetVersion gets the version property value. The version property
// returns a *string when successful
func (m *OEmbedGetResponse) GetVersion()(*string) {
    return m.version
}
// GetWidth gets the width property value. The width property
// returns a *int32 when successful
func (m *OEmbedGetResponse) GetWidth()(*int32) {
    return m.width
}
// Serialize serializes information the current object
func (m *OEmbedGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("height", m.GetHeight())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("html", m.GetHtml())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("providerName", m.GetProviderName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("providerUrl", m.GetProviderUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("thumbnailUrl", m.GetThumbnailUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("width", m.GetWidth())
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
func (m *OEmbedGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHeight sets the height property value. The height property
func (m *OEmbedGetResponse) SetHeight(value *int32)() {
    m.height = value
}
// SetHtml sets the html property value. The html property
func (m *OEmbedGetResponse) SetHtml(value *string)() {
    m.html = value
}
// SetProviderName sets the providerName property value. The providerName property
func (m *OEmbedGetResponse) SetProviderName(value *string)() {
    m.providerName = value
}
// SetProviderUrl sets the providerUrl property value. The providerUrl property
func (m *OEmbedGetResponse) SetProviderUrl(value *string)() {
    m.providerUrl = value
}
// SetThumbnailUrl sets the thumbnailUrl property value. The thumbnailUrl property
func (m *OEmbedGetResponse) SetThumbnailUrl(value *string)() {
    m.thumbnailUrl = value
}
// SetTitle sets the title property value. The title property
func (m *OEmbedGetResponse) SetTitle(value *string)() {
    m.title = value
}
// SetTypeEscaped sets the type property value. The type property
func (m *OEmbedGetResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetVersion sets the version property value. The version property
func (m *OEmbedGetResponse) SetVersion(value *string)() {
    m.version = value
}
// SetWidth sets the width property value. The width property
func (m *OEmbedGetResponse) SetWidth(value *int32)() {
    m.width = value
}
type OEmbedGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHeight()(*int32)
    GetHtml()(*string)
    GetProviderName()(*string)
    GetProviderUrl()(*string)
    GetThumbnailUrl()(*string)
    GetTitle()(*string)
    GetTypeEscaped()(*string)
    GetVersion()(*string)
    GetWidth()(*int32)
    SetHeight(value *int32)()
    SetHtml(value *string)()
    SetProviderName(value *string)()
    SetProviderUrl(value *string)()
    SetThumbnailUrl(value *string)()
    SetTitle(value *string)()
    SetTypeEscaped(value *string)()
    SetVersion(value *string)()
    SetWidth(value *int32)()
}
