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

package main

import (
	"fmt"
	"log"
	"os"

	appcatv1 "github.com/vshn/appcat-apiserver/apis/appcat/v1"
	"github.com/vshn/appcat-apiserver/apiserver/appcat"
	"k8s.io/klog"
	"sigs.k8s.io/apiserver-runtime/pkg/builder"
)

func main() {

	var appcatEnabled bool = false
	// var vshnEnabled bool = false

	if os.Getenv("APPCAT_ENABLED") == "1" {
		appcatEnabled = true
	}
	// if os.Getenv("VSHN_ENABLED") == "1" {
	// 	vshnEnabled = true
	// }
	builder := builder.APIServer

	if appcatEnabled {
		builder.WithResourceAndHandler(&appcatv1.AppCat{}, appcat.New())
		fmt.Println("Enabling appcat")
	}

	// if vshnEnabled {
	// 	builder.WithResourceAndHandler(&appcatv1.VSHNPostgresBackup{}, postgres.New())
	// 	fmt.Println("Enabling vshnpostgresql")
	// }
	builder.WithoutEtcd().
		ExposeLoopbackAuthorizer().
		ExposeLoopbackMasterClientConfig()

	cmd, err := builder.Build()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Execute(); err != nil {
		klog.Fatal(err)
	}
}
