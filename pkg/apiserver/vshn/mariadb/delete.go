package mariadb

import (
	"context"

	v1 "github.com/vshn/appcat-apiserver/apis/appcat/v1"
	metainternalversion "k8s.io/apimachinery/pkg/apis/meta/internalversion"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/rest"
)

var _ rest.GracefulDeleter = &vshnMariaDBBackupStorage{}
var _ rest.CollectionDeleter = &vshnMariaDBBackupStorage{}

func (v vshnMariaDBBackupStorage) Delete(_ context.Context, name string, _ rest.ValidateObjectFunc, _ *metav1.DeleteOptions) (runtime.Object, bool, error) {
	return &v1.VSHNMariaDBBackup{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}, false, nil
}

func (v *vshnMariaDBBackupStorage) DeleteCollection(ctx context.Context, _ rest.ValidateObjectFunc, _ *metav1.DeleteOptions, _ *metainternalversion.ListOptions) (runtime.Object, error) {
	return &v1.VSHNMariaDBBackupList{
		Items: []v1.VSHNMariaDBBackup{},
	}, nil
}
