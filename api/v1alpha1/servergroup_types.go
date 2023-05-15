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

// ServerGroupSpec defines the desired state of ServerGroup
type ServerGroupSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	c.ServerGroupSelfSpec `json:",inline" yaml:",inline"`
	Servers               ServerInServerGroup `json:"servers" yaml:"servers"`
}

type ServerInServerGroup struct {
	Endpoints []EndpointsServerConf `json:"endpoints" yaml:"endpoints"`
}

type EndpointsServerConf struct {
	Name string `json:"name" yaml:"name"`
	Port int    `json:"port" yaml:"port"`

	//+optional
	Weight *int `json:"weight,omitempty" yaml:"weight,omitempty"`
}

// ServerGroupStatus defines the observed state of ServerGroup
type ServerGroupStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	c.ServerGroupStatus `json:",inline" yaml:",inline"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ServerGroup is the Schema for the servergroups API
type ServerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServerGroupSpec   `json:"spec,omitempty"`
	Status ServerGroupStatus `json:"status,omitempty"`

	//+optional
	SyncId int `json:"syncId"`
}

//+kubebuilder:object:root=true

// ServerGroupList contains a list of ServerGroup
type ServerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServerGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServerGroup{}, &ServerGroupList{})
}
