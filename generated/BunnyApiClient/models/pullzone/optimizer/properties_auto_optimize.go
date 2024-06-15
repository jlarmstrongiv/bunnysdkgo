package optimizer
type Properties_auto_optimize int

const (
    LOW_PROPERTIES_AUTO_OPTIMIZE Properties_auto_optimize = iota
    MEDIUM_PROPERTIES_AUTO_OPTIMIZE
    HIGH_PROPERTIES_AUTO_OPTIMIZE
)

func (i Properties_auto_optimize) String() string {
    return []string{"low", "medium", "high"}[i]
}
func ParseProperties_auto_optimize(v string) (any, error) {
    result := LOW_PROPERTIES_AUTO_OPTIMIZE
    switch v {
        case "low":
            result = LOW_PROPERTIES_AUTO_OPTIMIZE
        case "medium":
            result = MEDIUM_PROPERTIES_AUTO_OPTIMIZE
        case "high":
            result = HIGH_PROPERTIES_AUTO_OPTIMIZE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeProperties_auto_optimize(values []Properties_auto_optimize) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i Properties_auto_optimize) isMultiValue() bool {
    return false
}
