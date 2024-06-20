package v1

import (
	v1 "github.com/vshn/appcat-apiserver/apis/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Workaround to make nested defaulting work.
// kubebuilder is unable to set a {} default
//go:generate yq -i e ../../generated/vshn.appcat.vshn.io_vshnmariadbs.yaml --expression "with(.spec.versions[]; .schema.openAPIV3Schema.properties.spec.properties.parameters.default={})"
//go:generate yq -i e ../../generated/vshn.appcat.vshn.io_vshnmariadbs.yaml --expression "with(.spec.versions[]; .schema.openAPIV3Schema.properties.spec.properties.parameters.properties.size.default={})"
//go:generate yq -i e ../../generated/vshn.appcat.vshn.io_vshnmariadbs.yaml --expression "with(.spec.versions[]; .schema.openAPIV3Schema.properties.spec.properties.parameters.properties.service.default={})"
//go:generate yq -i e ../../generated/vshn.appcat.vshn.io_vshnmariadbs.yaml --expression "with(.spec.versions[]; .schema.openAPIV3Schema.properties.spec.properties.parameters.properties.tls.default={})"
//go:generate yq -i e ../../generated/vshn.appcat.vshn.io_vshnmariadbs.yaml --expression "with(.spec.versions[]; .schema.openAPIV3Schema.properties.spec.properties.parameters.properties.backup.default={})"

// +kubebuilder:object:root=true

// VSHNMariaDB is the API for creating MariaDB clusters.
type VSHNMariaDB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of a VSHNMariaDB.
	Spec VSHNMariaDBSpec `json:"spec"`

	// Status reflects the observed state of a VSHNMariaDB.
	Status VSHNMariaDBStatus `json:"status,omitempty"`
}

// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true
type VSHNMariaDBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []VSHNMariaDB `json:"items,omitempty"`
}

// VSHNMariaDBSpec defines the desired state of a VSHNMariaDB.
type VSHNMariaDBSpec struct {
	// Parameters are the configurable fields of a VSHNMariaDB.
	Parameters VSHNMariaDBParameters `json:"parameters,omitempty"`

	// WriteConnectionSecretToRef references a secret to which the connection details will be written.
	WriteConnectionSecretToRef v1.LocalObjectReference `json:"writeConnectionSecretToRef,omitempty"`
}

// VSHNMariaDBParameters are the configurable fields of a VSHNMariaDB.
type VSHNMariaDBParameters struct {
	// Service contains MariaDB DBaaS specific properties
	Service VSHNMariaDBServiceSpec `json:"service,omitempty"`

	// Size contains settings to control the sizing of a service.
	Size VSHNMariaDBSizeSpec `json:"size,omitempty"`

	// Scheduling contains settings to control the scheduling of an instance.
	Scheduling VSHNDBaaSSchedulingSpec `json:"scheduling,omitempty"`

	// TLS contains settings to control tls traffic of a service.
	TLS VSHNMariaDBTLSSpec `json:"tls,omitempty"`

	// Backup contains settings to control how the instance should get backed up.
	Backup K8upBackupSpec `json:"backup,omitempty"`

	// Restore contains settings to control the restore of an instance.
	Restore K8upRestoreSpec `json:"restore,omitempty"`

	// Maintenance contains settings to control the maintenance of an instance.
	Maintenance VSHNDBaaSMaintenanceScheduleSpec `json:"maintenance,omitempty"`
}

// VSHNMariaDBServiceSpec contains MariaDB DBaaS specific properties
type VSHNMariaDBServiceSpec struct {
	// +kubebuilder:validation:Enum="6.2";"7.0"
	// +kubebuilder:default="7.0"

	// Version contains supported version of MariaDB.
	// Multiple versions are supported. The latest version "7.0" is the default version.
	Version string `json:"version,omitempty"`

	// MariaDBSettings contains additional MariaDB settings.
	MariaDBSettings string `json:"MariaDBSettings,omitempty"`
}

// VSHNMariaDBSizeSpec contains settings to control the sizing of a service.
type VSHNMariaDBSizeSpec struct {

	// CPURequests defines the requests amount of Kubernetes CPUs for an instance.
	CPURequests string `json:"cpuRequests,omitempty"`

	// CPULimits defines the limits amount of Kubernetes CPUs for an instance.
	CPULimits string `json:"cpuLimits,omitempty"`

	// MemoryRequests defines the requests amount of memory in units of bytes for an instance.
	MemoryRequests string `json:"memoryRequests,omitempty"`

	// MemoryLimits defines the limits amount of memory in units of bytes for an instance.
	MemoryLimits string `json:"memoryLimits,omitempty"`

	// Disk defines the amount of disk space for an instance.
	Disk string `json:"disk,omitempty"`

	// Plan is the name of the resource plan that defines the compute resources.
	Plan string `json:"plan,omitempty"`
}

// VSHNMariaDBTLSSpec contains settings to control tls traffic of a service.
type VSHNMariaDBTLSSpec struct {
	// +kubebuilder:default=true

	// TLSEnabled enables TLS traffic for the service
	TLSEnabled bool `json:"enabled,omitempty"`

	// +kubebuilder:default=true
	// TLSAuthClients enables client authentication requirement
	TLSAuthClients bool `json:"authClients,omitempty"`
}

// VSHNMariaDBStatus reflects the observed state of a VSHNMariaDB.
type VSHNMariaDBStatus struct {
	// MariaDBConditions contains the status conditions of the backing object.
	NamespaceConditions         []v1.Condition `json:"namespaceConditions,omitempty"`
	SelfSignedIssuerConditions  []v1.Condition `json:"selfSignedIssuerConditions,omitempty"`
	LocalCAConditions           []v1.Condition `json:"localCAConditions,omitempty"`
	CaCertificateConditions     []v1.Condition `json:"caCertificateConditions,omitempty"`
	ServerCertificateConditions []v1.Condition `json:"serverCertificateConditions,omitempty"`
	ClientCertificateConditions []v1.Condition `json:"clientCertificateConditions,omitempty"`
	// InstanceNamespace contains the name of the namespace where the instance resides
	InstanceNamespace string `json:"instanceNamespace,omitempty"`
}
