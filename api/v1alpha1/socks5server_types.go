/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	c "github.com/vproxy-tools/vpctl/pkg/vproxy_config"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// Socks5ServerSpec defines the desired state of Socks5Server
type Socks5ServerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	c.Socks5ServerSpec `json:",inline" yaml:",inline"`
}

// Socks5ServerStatus defines the observed state of Socks5Server
type Socks5ServerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:categories=all,shortName=socks5

// Socks5Server is the Schema for the socks5servers API
type Socks5Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Socks5ServerSpec   `json:"spec,omitempty"`
	Status Socks5ServerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// Socks5ServerList contains a list of Socks5Server
type Socks5ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Socks5Server `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Socks5Server{}, &Socks5ServerList{})
}
