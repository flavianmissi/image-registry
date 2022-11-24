// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// ServiceCAStatusApplyConfiguration represents an declarative configuration of the ServiceCAStatus type for use
// with apply.
type ServiceCAStatusApplyConfiguration struct {
	OperatorStatusApplyConfiguration `json:",inline"`
}

// ServiceCAStatusApplyConfiguration constructs an declarative configuration of the ServiceCAStatus type for use with
// apply.
func ServiceCAStatus() *ServiceCAStatusApplyConfiguration {
	return &ServiceCAStatusApplyConfiguration{}
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *ServiceCAStatusApplyConfiguration) WithObservedGeneration(value int64) *ServiceCAStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *ServiceCAStatusApplyConfiguration) WithConditions(values ...*OperatorConditionApplyConfiguration) *ServiceCAStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *ServiceCAStatusApplyConfiguration) WithVersion(value string) *ServiceCAStatusApplyConfiguration {
	b.Version = &value
	return b
}

// WithReadyReplicas sets the ReadyReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadyReplicas field is set to the value of the last call.
func (b *ServiceCAStatusApplyConfiguration) WithReadyReplicas(value int32) *ServiceCAStatusApplyConfiguration {
	b.ReadyReplicas = &value
	return b
}

// WithGenerations adds the given value to the Generations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Generations field.
func (b *ServiceCAStatusApplyConfiguration) WithGenerations(values ...*GenerationStatusApplyConfiguration) *ServiceCAStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithGenerations")
		}
		b.Generations = append(b.Generations, *values[i])
	}
	return b
}
