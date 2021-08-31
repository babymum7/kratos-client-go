/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * API version: v0.7.3-alpha.2
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PluginDevice PluginDevice plugin device
type PluginDevice struct {
	// description
	Description string `json:"Description"`
	// name
	Name string `json:"Name"`
	// path
	Path string `json:"Path"`
	// settable
	Settable []string `json:"Settable"`
}

// NewPluginDevice instantiates a new PluginDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginDevice(description string, name string, path string, settable []string) *PluginDevice {
	this := PluginDevice{}
	this.Description = description
	this.Name = name
	this.Path = path
	this.Settable = settable
	return &this
}

// NewPluginDeviceWithDefaults instantiates a new PluginDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginDeviceWithDefaults() *PluginDevice {
	this := PluginDevice{}
	return &this
}

// GetDescription returns the Description field value
func (o *PluginDevice) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PluginDevice) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PluginDevice) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *PluginDevice) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PluginDevice) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PluginDevice) SetName(v string) {
	o.Name = v
}

// GetPath returns the Path field value
func (o *PluginDevice) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *PluginDevice) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *PluginDevice) SetPath(v string) {
	o.Path = v
}

// GetSettable returns the Settable field value
func (o *PluginDevice) GetSettable() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Settable
}

// GetSettableOk returns a tuple with the Settable field value
// and a boolean to check if the value has been set.
func (o *PluginDevice) GetSettableOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Settable, true
}

// SetSettable sets field value
func (o *PluginDevice) SetSettable(v []string) {
	o.Settable = v
}

func (o PluginDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Description"] = o.Description
	}
	if true {
		toSerialize["Name"] = o.Name
	}
	if true {
		toSerialize["Path"] = o.Path
	}
	if true {
		toSerialize["Settable"] = o.Settable
	}
	return json.Marshal(toSerialize)
}

type NullablePluginDevice struct {
	value *PluginDevice
	isSet bool
}

func (v NullablePluginDevice) Get() *PluginDevice {
	return v.value
}

func (v *NullablePluginDevice) Set(val *PluginDevice) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginDevice) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginDevice(val *PluginDevice) *NullablePluginDevice {
	return &NullablePluginDevice{value: val, isSet: true}
}

func (v NullablePluginDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


