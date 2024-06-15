package optimizer
type Properties_optimizer int

const (
    NONE_PROPERTIES_OPTIMIZER Properties_optimizer = iota
    IMAGE_PROPERTIES_OPTIMIZER
)

func (i Properties_optimizer) String() string {
    return []string{"none", "image"}[i]
}
func ParseProperties_optimizer(v string) (any, error) {
    result := NONE_PROPERTIES_OPTIMIZER
    switch v {
        case "none":
            result = NONE_PROPERTIES_OPTIMIZER
        case "image":
            result = IMAGE_PROPERTIES_OPTIMIZER
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeProperties_optimizer(values []Properties_optimizer) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Properties_optimizer) isMultiValue() bool {
    return false
}
