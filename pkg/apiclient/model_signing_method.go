/*
Daytona Server API

Daytona Server API

API version: v0.0.0-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"encoding/json"
	"fmt"
)

// SigningMethod the model 'SigningMethod'
type SigningMethod string

// List of SigningMethod
const (
	SigningMethodSSH SigningMethod = "ssh"
	SigningMethodGPG SigningMethod = "gpg"
)

// All allowed values of SigningMethod enum
var AllowedSigningMethodEnumValues = []SigningMethod{
	"ssh",
	"gpg",
}

func (v *SigningMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SigningMethod(value)
	for _, existing := range AllowedSigningMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SigningMethod", value)
}

// NewSigningMethodFromValue returns a pointer to a valid SigningMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSigningMethodFromValue(v string) (*SigningMethod, error) {
	ev := SigningMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SigningMethod: valid values are %v", v, AllowedSigningMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SigningMethod) IsValid() bool {
	for _, existing := range AllowedSigningMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SigningMethod value
func (v SigningMethod) Ptr() *SigningMethod {
	return &v
}

type NullableSigningMethod struct {
	value *SigningMethod
	isSet bool
}

func (v NullableSigningMethod) Get() *SigningMethod {
	return v.value
}

func (v *NullableSigningMethod) Set(val *SigningMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableSigningMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableSigningMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSigningMethod(val *SigningMethod) *NullableSigningMethod {
	return &NullableSigningMethod{value: val, isSet: true}
}

func (v NullableSigningMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSigningMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}