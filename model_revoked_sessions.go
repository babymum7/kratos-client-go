/*
Ory Kratos API

Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 

API version: latest
Contact: hi@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// RevokedSessions struct for RevokedSessions
type RevokedSessions struct {
	// The number of sessions that were revoked.
	Count *int64 `json:"count,omitempty"`
}

// NewRevokedSessions instantiates a new RevokedSessions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevokedSessions() *RevokedSessions {
	this := RevokedSessions{}
	return &this
}

// NewRevokedSessionsWithDefaults instantiates a new RevokedSessions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevokedSessionsWithDefaults() *RevokedSessions {
	this := RevokedSessions{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *RevokedSessions) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevokedSessions) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *RevokedSessions) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *RevokedSessions) SetCount(v int64) {
	o.Count = &v
}

func (o RevokedSessions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableRevokedSessions struct {
	value *RevokedSessions
	isSet bool
}

func (v NullableRevokedSessions) Get() *RevokedSessions {
	return v.value
}

func (v *NullableRevokedSessions) Set(val *RevokedSessions) {
	v.value = val
	v.isSet = true
}

func (v NullableRevokedSessions) IsSet() bool {
	return v.isSet
}

func (v *NullableRevokedSessions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevokedSessions(val *RevokedSessions) *NullableRevokedSessions {
	return &NullableRevokedSessions{value: val, isSet: true}
}

func (v NullableRevokedSessions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevokedSessions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


