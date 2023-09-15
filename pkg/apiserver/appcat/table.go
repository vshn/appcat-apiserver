package appcat

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/vshn/appcat-apiserver/apis/appcat/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/duration"
	"k8s.io/apiserver/pkg/registry/rest"
)

var _ rest.TableConvertor = &appcatStorage{}

var (
	appCatDisplayname = "displayname"
	appCatZone        = "zone"
	appCatDocs        = "endUserDocsUrl"
)

// ConvertToTable translates the given object to a table for kubectl printing
func (s *appcatStorage) ConvertToTable(_ context.Context, obj runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	var table metav1.Table

	appcats := []v1.AppCat{}
	if meta.IsListType(obj) {
		appcatList, ok := obj.(*v1.AppCatList)
		if !ok {
			return nil, fmt.Errorf("not an appcat: %#v", obj)
		}
		appcats = appcatList.Items
	} else {
		appcat, ok := obj.(*v1.AppCat)
		if !ok {
			return nil, fmt.Errorf("not an appcat: %#v", obj)
		}
		appcats = append(appcats, *appcat)
	}

	for _, appcat := range appcats {
		table.Rows = append(table.Rows, appcatToTableRow(&appcat))
	}

	if opt, ok := tableOptions.(*metav1.TableOptions); !ok || !opt.NoHeaders {
		desc := metav1.ObjectMeta{}.SwaggerDoc()
		table.ColumnDefinitions = []metav1.TableColumnDefinition{
			{Name: "AppCat Name", Type: "string", Format: "name", Description: desc["name"]},
			{Name: "AppCat Display Name", Type: "string", Format: "name", Description: "The display name of the service"},
			{Name: "Service Zone", Type: "string", Description: "Available zones of the service"},
			{Name: "User Docs", Type: "string", Description: "The user documentation of the service"},
			{Name: "Age", Type: "date", Description: desc["creationTimestamp"]},
		}
	}
	return &table, nil
}

func appcatToTableRow(appcat *v1.AppCat) metav1.TableRow {
	return metav1.TableRow{
		Cells: []interface{}{
			appcat.GetName(),
			appcat.Details[appCatDisplayname],
			appcat.Details[appCatZone],
			appcat.Details[appCatDocs],
			duration.HumanDuration(time.Since(appcat.GetCreationTimestamp().Time))},
		Object: runtime.RawExtension{Object: appcat},
	}
}
