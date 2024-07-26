package microlesson

//	Type elements specify which types can be assigned to a type parameter and which operators are supported.
//	They list concrete types separated by I. The allowed operators are the ones that are valid for all of the listed types.
//	Be aware that interfaces with type elements are valid only as type constraints.
//	It is a compile-time error to use them as the type for a variable, field, return value, or parameter.

//	Interface Type Term:
//		Use this when you want to constrain the type to a specific set of types or to enforce certain methods.
//	Generic Type Parameter T:
//		Use this when you want flexibility to work with any type, while still maintaining type safety.

// Use the ~ modifier to allow a type to be valid for a type term that has that type as its underlying type

type AcceptableTypeTerms interface {
	int | ~string
}

type StackV2[T AcceptableTypeTerms] struct {
	values []T
}

func (s *StackV2[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *StackV2[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StackV2[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}
