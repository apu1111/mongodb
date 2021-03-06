/*
Copyright AppsCode Inc. and Contributors

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
	"fmt"

	"kubedb.dev/apimachinery/apis"
	"kubedb.dev/apimachinery/apis/catalog"
	"kubedb.dev/apimachinery/crds"

	"kmodules.xyz/client-go/apiextensions"
)

func (_ EtcdVersion) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(SchemeGroupVersion.WithResource(ResourcePluralEtcdVersion))
}

var _ apis.ResourceInfo = &EtcdVersion{}

func (e EtcdVersion) ResourceFQN() string {
	return fmt.Sprintf("%s.%s", ResourcePluralEtcdVersion, catalog.GroupName)
}

func (e EtcdVersion) ResourceShortCode() string {
	return ResourceCodeEtcdVersion
}

func (e EtcdVersion) ResourceKind() string {
	return ResourceKindEtcdVersion
}

func (e EtcdVersion) ResourceSingular() string {
	return ResourceSingularEtcdVersion
}

func (e EtcdVersion) ResourcePlural() string {
	return ResourcePluralEtcdVersion
}

func (e EtcdVersion) ValidateSpecs() error {
	if e.Spec.Version == "" ||
		e.Spec.DB.Image == "" ||
		e.Spec.Tools.Image == "" ||
		e.Spec.Exporter.Image == "" {
		return fmt.Errorf(`atleast one of the following specs is not set for etcdVersion "%v":
spec.version,
spec.db.image,
spec.tools.image,
spec.exporter.image.`, e.Name)
	}
	return nil
}
