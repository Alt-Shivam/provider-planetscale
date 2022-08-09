/*
Copyright 2022 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// DatabaseParameters are the configurable fields of a Database.
type DatabaseParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// DatabaseObservation are the observable fields of a Database.
type DatabaseObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A DatabaseSpec defines the desired state of a Database.
type DatabaseSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DatabaseParameters `json:"forProvider"`
}

// A DatabaseStatus represents the observed state of a Database.
type DatabaseStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Database is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,planetscale}
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatabaseSpec   `json:"spec"`
	Status DatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseList contains a list of Database
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Database `json:"items"`
}

// Database type metadata.
var (
	DatabaseKind             = reflect.TypeOf(Database{}).Name()
	DatabaseGroupKind        = schema.GroupKind{Group: Group, Kind: DatabaseKind}.String()
	DatabaseKindAPIVersion   = DatabaseKind + "." + SchemeGroupVersion.String()
	DatabaseGroupVersionKind = SchemeGroupVersion.WithKind(DatabaseKind)
)

func init() {
	SchemeBuilder.Register(&Database{}, &DatabaseList{})
}
