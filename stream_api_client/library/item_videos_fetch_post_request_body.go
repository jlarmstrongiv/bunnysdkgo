package library

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ItemVideosFetchPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The headers that will be sent along with the fetch request.
    headers ItemVideosFetchPostRequestBody_headersable
    // The title that will be set to video.
    title *string
    // The URL from which the video will be fetched from.
    url *string
}
// NewItemVideosFetchPostRequestBody instantiates a new ItemVideosFetchPostRequestBody and sets the default values.
func NewItemVideosFetchPostRequestBody()(*ItemVideosFetchPostRequestBody) {
    m := &ItemVideosFetchPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemVideosFetchPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemVideosFetchPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemVideosFetchPostRequestBody(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemVideosFetchPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemVideosFetchPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["headers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemVideosFetchPostRequestBody_headersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeaders(val.(ItemVideosFetchPostRequestBody_headersable))
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
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    return res
}
// GetHeaders gets the headers property value. The headers that will be sent along with the fetch request.
// returns a ItemVideosFetchPostRequestBody_headersable when successful
func (m *ItemVideosFetchPostRequestBody) GetHeaders()(ItemVideosFetchPostRequestBody_headersable) {
    return m.headers
}
// GetTitle gets the title property value. The title that will be set to video.
// returns a *string when successful
func (m *ItemVideosFetchPostRequestBody) GetTitle()(*string) {
    return m.title
}
// GetUrl gets the url property value. The URL from which the video will be fetched from.
// returns a *string when successful
func (m *ItemVideosFetchPostRequestBody) GetUrl()(*string) {
    return m.url
}
// Serialize serializes information the current object
func (m *ItemVideosFetchPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("headers", m.GetHeaders())
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
        err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *ItemVideosFetchPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHeaders sets the headers property value. The headers that will be sent along with the fetch request.
func (m *ItemVideosFetchPostRequestBody) SetHeaders(value ItemVideosFetchPostRequestBody_headersable)() {
    m.headers = value
}
// SetTitle sets the title property value. The title that will be set to video.
func (m *ItemVideosFetchPostRequestBody) SetTitle(value *string)() {
    m.title = value
}
// SetUrl sets the url property value. The URL from which the video will be fetched from.
func (m *ItemVideosFetchPostRequestBody) SetUrl(value *string)() {
    m.url = value
}
type ItemVideosFetchPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHeaders()(ItemVideosFetchPostRequestBody_headersable)
    GetTitle()(*string)
    GetUrl()(*string)
    SetHeaders(value ItemVideosFetchPostRequestBody_headersable)()
    SetTitle(value *string)()
    SetUrl(value *string)()
}
