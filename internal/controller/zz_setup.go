/*
Copyright 2021 The Crossplane Authors.

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

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	key "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/api/key"
	keyapplication "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/application/key"
	organization "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/child/organization"
	json "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/dashboard/json"
	list "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/dashboard/list"
	dashboard "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/datadog/dashboard"
	downtime "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/datadog/downtime"
	monitor "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/datadog/monitor"
	role "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/datadog/role"
	user "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/datadog/user"
	webhook "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/datadog/webhook"
	aws "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/integration/aws"
	awslambdaarn "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/integration/awslambdaarn"
	awslogcollection "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/integration/awslogcollection"
	awstagfilter "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/integration/awstagfilter"
	azure "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/integration/azure"
	gcp "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/integration/gcp"
	pagerduty "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/integration/pagerduty"
	pagerdutyserviceobject "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/integration/pagerdutyserviceobject"
	slackchannel "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/integration/slackchannel"
	archive "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/logs/archive"
	archiveorder "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/logs/archiveorder"
	custompipeline "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/logs/custompipeline"
	index "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/logs/index"
	indexorder "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/logs/indexorder"
	integrationpipeline "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/logs/integrationpipeline"
	metric "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/logs/metric"
	pipelineorder "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/logs/pipelineorder"
	metadata "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/metric/metadata"
	tagconfiguration "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/metric/tagconfiguration"
	jsonmonitor "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/monitor/json"
	settings "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/organization/settings"
	providerconfig "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/providerconfig"
	monitoringdefaultrule "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/security/monitoringdefaultrule"
	monitoringfilter "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/security/monitoringfilter"
	monitoringrule "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/security/monitoringrule"
	levelobjective "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/service/levelobjective"
	correction "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/slo/correction"
	globalvariable "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/synthetics/globalvariable"
	privatelocation "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/synthetics/privatelocation"
	test "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/synthetics/test"
	customvariable "github.com/crossplane-contrib/provider-jet-datadog/internal/controller/webhook/customvariable"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		key.Setup,
		keyapplication.Setup,
		organization.Setup,
		json.Setup,
		list.Setup,
		dashboard.Setup,
		downtime.Setup,
		monitor.Setup,
		role.Setup,
		user.Setup,
		webhook.Setup,
		aws.Setup,
		awslambdaarn.Setup,
		awslogcollection.Setup,
		awstagfilter.Setup,
		azure.Setup,
		gcp.Setup,
		pagerduty.Setup,
		pagerdutyserviceobject.Setup,
		slackchannel.Setup,
		archive.Setup,
		archiveorder.Setup,
		custompipeline.Setup,
		index.Setup,
		indexorder.Setup,
		integrationpipeline.Setup,
		metric.Setup,
		pipelineorder.Setup,
		metadata.Setup,
		tagconfiguration.Setup,
		jsonmonitor.Setup,
		settings.Setup,
		providerconfig.Setup,
		monitoringdefaultrule.Setup,
		monitoringfilter.Setup,
		monitoringrule.Setup,
		levelobjective.Setup,
		correction.Setup,
		globalvariable.Setup,
		privatelocation.Setup,
		test.Setup,
		customvariable.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
