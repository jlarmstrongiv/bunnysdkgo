package collections
type GetOrderByQueryParameterType int

const (
    DATE_GETORDERBYQUERYPARAMETERTYPE GetOrderByQueryParameterType = iota
    TITLE_GETORDERBYQUERYPARAMETERTYPE
)

func (i GetOrderByQueryParameterType) String() string {
    return []string{"date", "title"}[i]
}
func ParseGetOrderByQueryParameterType(v string) (any, error) {
    result := DATE_GETORDERBYQUERYPARAMETERTYPE
    switch v {
        case "date":
            result = DATE_GETORDERBYQUERYPARAMETERTYPE
        case "title":
            result = TITLE_GETORDERBYQUERYPARAMETERTYPE
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeGetOrderByQueryParameterType(values []GetOrderByQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GetOrderByQueryParameterType) isMultiValue() bool {
    return false
}
