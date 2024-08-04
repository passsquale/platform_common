package filter

// Filter - фильтр, содержащий условия запроса
type Filter struct {
	Conditions []Condition
}

// Condition - Условие запроса
type Condition struct {
	Key   string
	Value interface{}
}

// MakeFilter - создает фильтр с условиями
func MakeFilter(conditions ...Condition) Filter {
	return Filter{
		Conditions: conditions,
	}
}

// AddCondition - добавляет новое условие в фильтр
func (f *Filter) AddCondition(condition Condition) {
	f.Conditions = append(f.Conditions, condition)
}
