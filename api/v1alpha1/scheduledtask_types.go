/*
Copyright 2026.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:validation:Enum=Ignore;Pause
type FailurePolicy string

const (
	FailurePolicyIgnore FailurePolicy = "Ignore"
	FailurePolicyPause  FailurePolicy = "Pause"
)

// ScheduledTaskSpec defines the desired state of ScheduledTask
type ScheduledTaskSpec struct {

	// +required
	Schedule string `json:"schedule"`

	// +required
	Image string `json:"image"`
	
	// +optional
	RetryLimit *int32 `json:"retryLimit"`

	// +optional
	// +kubebuilder:default=Ignore
	FailurePolicy FailurePolicy `json:"failurePolicy"`


}

// +kubebuilder:validation:Enum=Succeeded;Failed
type LastStatus string

const (
	LastStatusSucceeded LastStatus = "Succeeded"
	LastStatusFailed LastStatus = "Failed"
)

// ScheduledTaskStatus defines the observed state of ScheduledTask.
type ScheduledTaskStatus struct {

	// The status of each condition is one of True, False, or Unknown.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// +optional
	LastRun *metav1.Time `json:"lastRun,omitempty"`
	
	// +optional
	NextRun *metav1.Time `json:"nextRun,omitempty"`

	// +optional
	LastStatus LastStatus `json:"lastStatus,omitempty"`
	
	RunCount int32 `json:"runCount"`
	SuccessCount int32 `json:"successCount"`
	FailureCount int32 `json:"failureCount"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ScheduledTask is the Schema for the scheduledtasks API
type ScheduledTask struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of ScheduledTask
	// +required
	Spec ScheduledTaskSpec `json:"spec"`

	// status defines the observed state of ScheduledTask
	// +optional
	Status ScheduledTaskStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// ScheduledTaskList contains a list of ScheduledTask
type ScheduledTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []ScheduledTask `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ScheduledTask{}, &ScheduledTaskList{})
}
