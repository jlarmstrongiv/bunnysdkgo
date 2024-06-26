package optimizer

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Properties struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The aspect_ratio property
    aspect_ratio *string
    // The auto_optimize property
    auto_optimize *Properties_auto_optimize
    // The blur property
    blur *string
    // The brightness property
    brightness *string
    // The contrast property
    contrast *string
    // The crop property
    crop *string
    // The crop_gravity property
    crop_gravity *Properties_crop_gravity
    // The flip property
    flip *string
    // The flop property
    flop *string
    // The gamma property
    gamma *string
    // The height property
    height *string
    // The hue property
    hue *string
    // The optimizer property
    optimizer *Properties_optimizer
    // The quality property
    quality *string
    // The saturation property
    saturation *string
    // The sharpen property
    sharpen *string
    // The width property
    width *string
}
// NewProperties instantiates a new Properties and sets the default values.
func NewProperties()(*Properties) {
    m := &Properties{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePropertiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePropertiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProperties(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *Properties) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAspectRatio gets the aspect_ratio property value. The aspect_ratio property
// returns a *string when successful
func (m *Properties) GetAspectRatio()(*string) {
    return m.aspect_ratio
}
// GetAutoOptimize gets the auto_optimize property value. The auto_optimize property
// returns a *Properties_auto_optimize when successful
func (m *Properties) GetAutoOptimize()(*Properties_auto_optimize) {
    return m.auto_optimize
}
// GetBlur gets the blur property value. The blur property
// returns a *string when successful
func (m *Properties) GetBlur()(*string) {
    return m.blur
}
// GetBrightness gets the brightness property value. The brightness property
// returns a *string when successful
func (m *Properties) GetBrightness()(*string) {
    return m.brightness
}
// GetContrast gets the contrast property value. The contrast property
// returns a *string when successful
func (m *Properties) GetContrast()(*string) {
    return m.contrast
}
// GetCrop gets the crop property value. The crop property
// returns a *string when successful
func (m *Properties) GetCrop()(*string) {
    return m.crop
}
// GetCropGravity gets the crop_gravity property value. The crop_gravity property
// returns a *Properties_crop_gravity when successful
func (m *Properties) GetCropGravity()(*Properties_crop_gravity) {
    return m.crop_gravity
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *Properties) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["aspect_ratio"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAspectRatio(val)
        }
        return nil
    }
    res["auto_optimize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProperties_auto_optimize)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutoOptimize(val.(*Properties_auto_optimize))
        }
        return nil
    }
    res["blur"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBlur(val)
        }
        return nil
    }
    res["brightness"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrightness(val)
        }
        return nil
    }
    res["contrast"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContrast(val)
        }
        return nil
    }
    res["crop"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCrop(val)
        }
        return nil
    }
    res["crop_gravity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProperties_crop_gravity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCropGravity(val.(*Properties_crop_gravity))
        }
        return nil
    }
    res["flip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFlip(val)
        }
        return nil
    }
    res["flop"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFlop(val)
        }
        return nil
    }
    res["gamma"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGamma(val)
        }
        return nil
    }
    res["height"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeight(val)
        }
        return nil
    }
    res["hue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHue(val)
        }
        return nil
    }
    res["optimizer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProperties_optimizer)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptimizer(val.(*Properties_optimizer))
        }
        return nil
    }
    res["quality"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuality(val)
        }
        return nil
    }
    res["saturation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSaturation(val)
        }
        return nil
    }
    res["sharpen"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharpen(val)
        }
        return nil
    }
    res["width"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
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
// GetFlip gets the flip property value. The flip property
// returns a *string when successful
func (m *Properties) GetFlip()(*string) {
    return m.flip
}
// GetFlop gets the flop property value. The flop property
// returns a *string when successful
func (m *Properties) GetFlop()(*string) {
    return m.flop
}
// GetGamma gets the gamma property value. The gamma property
// returns a *string when successful
func (m *Properties) GetGamma()(*string) {
    return m.gamma
}
// GetHeight gets the height property value. The height property
// returns a *string when successful
func (m *Properties) GetHeight()(*string) {
    return m.height
}
// GetHue gets the hue property value. The hue property
// returns a *string when successful
func (m *Properties) GetHue()(*string) {
    return m.hue
}
// GetOptimizer gets the optimizer property value. The optimizer property
// returns a *Properties_optimizer when successful
func (m *Properties) GetOptimizer()(*Properties_optimizer) {
    return m.optimizer
}
// GetQuality gets the quality property value. The quality property
// returns a *string when successful
func (m *Properties) GetQuality()(*string) {
    return m.quality
}
// GetSaturation gets the saturation property value. The saturation property
// returns a *string when successful
func (m *Properties) GetSaturation()(*string) {
    return m.saturation
}
// GetSharpen gets the sharpen property value. The sharpen property
// returns a *string when successful
func (m *Properties) GetSharpen()(*string) {
    return m.sharpen
}
// GetWidth gets the width property value. The width property
// returns a *string when successful
func (m *Properties) GetWidth()(*string) {
    return m.width
}
// Serialize serializes information the current object
func (m *Properties) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("aspect_ratio", m.GetAspectRatio())
        if err != nil {
            return err
        }
    }
    if m.GetAutoOptimize() != nil {
        cast := (*m.GetAutoOptimize()).String()
        err := writer.WriteStringValue("auto_optimize", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("blur", m.GetBlur())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("brightness", m.GetBrightness())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contrast", m.GetContrast())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("crop", m.GetCrop())
        if err != nil {
            return err
        }
    }
    if m.GetCropGravity() != nil {
        cast := (*m.GetCropGravity()).String()
        err := writer.WriteStringValue("crop_gravity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("flip", m.GetFlip())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("flop", m.GetFlop())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("gamma", m.GetGamma())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("height", m.GetHeight())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hue", m.GetHue())
        if err != nil {
            return err
        }
    }
    if m.GetOptimizer() != nil {
        cast := (*m.GetOptimizer()).String()
        err := writer.WriteStringValue("optimizer", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("quality", m.GetQuality())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("saturation", m.GetSaturation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sharpen", m.GetSharpen())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("width", m.GetWidth())
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
func (m *Properties) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAspectRatio sets the aspect_ratio property value. The aspect_ratio property
func (m *Properties) SetAspectRatio(value *string)() {
    m.aspect_ratio = value
}
// SetAutoOptimize sets the auto_optimize property value. The auto_optimize property
func (m *Properties) SetAutoOptimize(value *Properties_auto_optimize)() {
    m.auto_optimize = value
}
// SetBlur sets the blur property value. The blur property
func (m *Properties) SetBlur(value *string)() {
    m.blur = value
}
// SetBrightness sets the brightness property value. The brightness property
func (m *Properties) SetBrightness(value *string)() {
    m.brightness = value
}
// SetContrast sets the contrast property value. The contrast property
func (m *Properties) SetContrast(value *string)() {
    m.contrast = value
}
// SetCrop sets the crop property value. The crop property
func (m *Properties) SetCrop(value *string)() {
    m.crop = value
}
// SetCropGravity sets the crop_gravity property value. The crop_gravity property
func (m *Properties) SetCropGravity(value *Properties_crop_gravity)() {
    m.crop_gravity = value
}
// SetFlip sets the flip property value. The flip property
func (m *Properties) SetFlip(value *string)() {
    m.flip = value
}
// SetFlop sets the flop property value. The flop property
func (m *Properties) SetFlop(value *string)() {
    m.flop = value
}
// SetGamma sets the gamma property value. The gamma property
func (m *Properties) SetGamma(value *string)() {
    m.gamma = value
}
// SetHeight sets the height property value. The height property
func (m *Properties) SetHeight(value *string)() {
    m.height = value
}
// SetHue sets the hue property value. The hue property
func (m *Properties) SetHue(value *string)() {
    m.hue = value
}
// SetOptimizer sets the optimizer property value. The optimizer property
func (m *Properties) SetOptimizer(value *Properties_optimizer)() {
    m.optimizer = value
}
// SetQuality sets the quality property value. The quality property
func (m *Properties) SetQuality(value *string)() {
    m.quality = value
}
// SetSaturation sets the saturation property value. The saturation property
func (m *Properties) SetSaturation(value *string)() {
    m.saturation = value
}
// SetSharpen sets the sharpen property value. The sharpen property
func (m *Properties) SetSharpen(value *string)() {
    m.sharpen = value
}
// SetWidth sets the width property value. The width property
func (m *Properties) SetWidth(value *string)() {
    m.width = value
}
type Propertiesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAspectRatio()(*string)
    GetAutoOptimize()(*Properties_auto_optimize)
    GetBlur()(*string)
    GetBrightness()(*string)
    GetContrast()(*string)
    GetCrop()(*string)
    GetCropGravity()(*Properties_crop_gravity)
    GetFlip()(*string)
    GetFlop()(*string)
    GetGamma()(*string)
    GetHeight()(*string)
    GetHue()(*string)
    GetOptimizer()(*Properties_optimizer)
    GetQuality()(*string)
    GetSaturation()(*string)
    GetSharpen()(*string)
    GetWidth()(*string)
    SetAspectRatio(value *string)()
    SetAutoOptimize(value *Properties_auto_optimize)()
    SetBlur(value *string)()
    SetBrightness(value *string)()
    SetContrast(value *string)()
    SetCrop(value *string)()
    SetCropGravity(value *Properties_crop_gravity)()
    SetFlip(value *string)()
    SetFlop(value *string)()
    SetGamma(value *string)()
    SetHeight(value *string)()
    SetHue(value *string)()
    SetOptimizer(value *Properties_optimizer)()
    SetQuality(value *string)()
    SetSaturation(value *string)()
    SetSharpen(value *string)()
    SetWidth(value *string)()
}
