/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * API version: v0.6.0-alpha.6
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// SubmitSelfServiceSettingsFlow - struct for SubmitSelfServiceSettingsFlow
type SubmitSelfServiceSettingsFlow struct {
	SubmitSelfServiceSettingsFlowWithPasswordMethod *SubmitSelfServiceSettingsFlowWithPasswordMethod
	SubmitSelfServiceSettingsFlowWithProfileMethod *SubmitSelfServiceSettingsFlowWithProfileMethod
}

// SubmitSelfServiceSettingsFlowWithPasswordMethodAsSubmitSelfServiceSettingsFlow is a convenience function that returns SubmitSelfServiceSettingsFlowWithPasswordMethod wrapped in SubmitSelfServiceSettingsFlow
func SubmitSelfServiceSettingsFlowWithPasswordMethodAsSubmitSelfServiceSettingsFlow(v *SubmitSelfServiceSettingsFlowWithPasswordMethod) SubmitSelfServiceSettingsFlow {
	return SubmitSelfServiceSettingsFlow{
		SubmitSelfServiceSettingsFlowWithPasswordMethod: v,
	}
}

// SubmitSelfServiceSettingsFlowWithProfileMethodAsSubmitSelfServiceSettingsFlow is a convenience function that returns SubmitSelfServiceSettingsFlowWithProfileMethod wrapped in SubmitSelfServiceSettingsFlow
func SubmitSelfServiceSettingsFlowWithProfileMethodAsSubmitSelfServiceSettingsFlow(v *SubmitSelfServiceSettingsFlowWithProfileMethod) SubmitSelfServiceSettingsFlow {
	return SubmitSelfServiceSettingsFlow{
		SubmitSelfServiceSettingsFlowWithProfileMethod: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SubmitSelfServiceSettingsFlow) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SubmitSelfServiceSettingsFlowWithPasswordMethod
	err = newStrictDecoder(data).Decode(&dst.SubmitSelfServiceSettingsFlowWithPasswordMethod)
	if err == nil {
		jsonSubmitSelfServiceSettingsFlowWithPasswordMethod, _ := json.Marshal(dst.SubmitSelfServiceSettingsFlowWithPasswordMethod)
		if string(jsonSubmitSelfServiceSettingsFlowWithPasswordMethod) == "{}" { // empty struct
			dst.SubmitSelfServiceSettingsFlowWithPasswordMethod = nil
		} else {
			match++
		}
	} else {
		dst.SubmitSelfServiceSettingsFlowWithPasswordMethod = nil
	}

	// try to unmarshal data into SubmitSelfServiceSettingsFlowWithProfileMethod
	err = newStrictDecoder(data).Decode(&dst.SubmitSelfServiceSettingsFlowWithProfileMethod)
	if err == nil {
		jsonSubmitSelfServiceSettingsFlowWithProfileMethod, _ := json.Marshal(dst.SubmitSelfServiceSettingsFlowWithProfileMethod)
		if string(jsonSubmitSelfServiceSettingsFlowWithProfileMethod) == "{}" { // empty struct
			dst.SubmitSelfServiceSettingsFlowWithProfileMethod = nil
		} else {
			match++
		}
	} else {
		dst.SubmitSelfServiceSettingsFlowWithProfileMethod = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SubmitSelfServiceSettingsFlowWithPasswordMethod = nil
		dst.SubmitSelfServiceSettingsFlowWithProfileMethod = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(SubmitSelfServiceSettingsFlow)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(SubmitSelfServiceSettingsFlow)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SubmitSelfServiceSettingsFlow) MarshalJSON() ([]byte, error) {
	if src.SubmitSelfServiceSettingsFlowWithPasswordMethod != nil {
		return json.Marshal(&src.SubmitSelfServiceSettingsFlowWithPasswordMethod)
	}

	if src.SubmitSelfServiceSettingsFlowWithProfileMethod != nil {
		return json.Marshal(&src.SubmitSelfServiceSettingsFlowWithProfileMethod)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SubmitSelfServiceSettingsFlow) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SubmitSelfServiceSettingsFlowWithPasswordMethod != nil {
		return obj.SubmitSelfServiceSettingsFlowWithPasswordMethod
	}

	if obj.SubmitSelfServiceSettingsFlowWithProfileMethod != nil {
		return obj.SubmitSelfServiceSettingsFlowWithProfileMethod
	}

	// all schemas are nil
	return nil
}

type NullableSubmitSelfServiceSettingsFlow struct {
	value *SubmitSelfServiceSettingsFlow
	isSet bool
}

func (v NullableSubmitSelfServiceSettingsFlow) Get() *SubmitSelfServiceSettingsFlow {
	return v.value
}

func (v *NullableSubmitSelfServiceSettingsFlow) Set(val *SubmitSelfServiceSettingsFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceSettingsFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceSettingsFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceSettingsFlow(val *SubmitSelfServiceSettingsFlow) *NullableSubmitSelfServiceSettingsFlow {
	return &NullableSubmitSelfServiceSettingsFlow{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceSettingsFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceSettingsFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


