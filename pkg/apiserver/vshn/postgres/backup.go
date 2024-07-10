package postgres

import (
	appcatv1 "github.com/vshn/appcat-apiserver/apis/appcat/v1"
	vshnv1 "github.com/vshn/appcat-apiserver/apis/vshn/v1"
	"github.com/vshn/appcat-apiserver/pkg/apiserver"
	"github.com/vshn/appcat-apiserver/pkg/apiserver/noop"
	"k8s.io/apimachinery/pkg/runtime"
	genericregistry "k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/client-go/dynamic"
	restbuilder "sigs.k8s.io/apiserver-runtime/pkg/builder/rest"
	"sigs.k8s.io/apiserver-runtime/pkg/util/loopback"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// New returns a new storage provider for VSHNPostgresBackup
func New() restbuilder.ResourceHandlerProvider {
	return func(s *runtime.Scheme, gasdf genericregistry.RESTOptionsGetter) (rest.Storage, error) {
		c, err := client.New(loopback.GetLoopbackMasterClientConfig(), client.Options{})
		if err != nil {
			return nil, err
		}

		err = vshnv1.AddToScheme(c.Scheme())
		if err != nil {
			return nil, err
		}

		noopImplementation := noop.New(s, &appcatv1.VSHNPostgresBackup{}, &appcatv1.VSHNPostgresBackupList{})

		if !apiserver.IsTypeAvailable(vshnv1.GroupVersion.String(), "XVSHNPostgreSQL") {
			return noopImplementation, nil
		}

		dc, err := dynamic.NewForConfig(loopback.GetLoopbackMasterClientConfig())
		if err != nil {
			return nil, err
		}
		return &vshnPostgresBackupStorage{
			sgbackups: &kubeSGBackupProvider{
				DynamicClient: dc.Resource(sgbackupGroupVersionResource),
			},
			vshnpostgresql: &kubeXVSHNPostgresqlProvider{
				Client: c,
			},
			Noop: *noopImplementation,
		}, nil
	}
}

type vshnPostgresBackupStorage struct {
	sgbackups      sgbackupProvider
	vshnpostgresql vshnPostgresqlProvider
	noop.Noop
}

// GetSingularName is needed for the OpenAPI Registartion
func (in *vshnPostgresBackupStorage) GetSingularName() string {
	return "vshnpostgresbackup"
}

func (v vshnPostgresBackupStorage) New() runtime.Object {
	return &appcatv1.VSHNPostgresBackup{}
}

func (v vshnPostgresBackupStorage) Destroy() {}

var _ rest.Scoper = &vshnPostgresBackupStorage{}
var _ rest.Storage = &vshnPostgresBackupStorage{}

func (v *vshnPostgresBackupStorage) NamespaceScoped() bool {
	return true
}
