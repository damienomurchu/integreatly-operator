// +build !ignore_autogenerated

/*
Copyright 2020 Red Hat, Inc.

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

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/integreatly/v1alpha1/.RHMI":       schema_apis_integreatly_v1alpha1__RHMI(ref),
		"./pkg/apis/integreatly/v1alpha1/.RHMISpec":   schema_apis_integreatly_v1alpha1__RHMISpec(ref),
		"./pkg/apis/integreatly/v1alpha1/.RHMIStatus": schema_apis_integreatly_v1alpha1__RHMIStatus(ref),
	}
}

func schema_apis_integreatly_v1alpha1__RHMI(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RHMI is the Schema for the RHMI API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/integreatly/v1alpha1/.RHMISpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/integreatly/v1alpha1/.RHMIStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/integreatly/v1alpha1/.RHMISpec", "./pkg/apis/integreatly/v1alpha1/.RHMIStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_apis_integreatly_v1alpha1__RHMISpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RHMISpec defines the desired state of Installation",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"routingSubdomain": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"masterURL": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"namespacePrefix": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"selfSignedCerts": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"pullSecret": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/integreatly/v1alpha1/.PullSecretSpec"),
						},
					},
					"useClusterStorage": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"operatorsInProductNamespace": {
						SchemaProps: spec.SchemaProps{
							Description: "OperatorsInProductNamespace is a flag that decides if the product operators should be installed in the product namespace (when set to true) or in standalone namespace (when set to false, default). Standalone namespace will be used only for those operators that support it.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"smtpSecret": {
						SchemaProps: spec.SchemaProps{
							Description: "SMTPSecret is the name of a secret in the installation namespace containing SMTP connection details. The secret must contain the following fields:\n\nhost port tls username password",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"pagerDutySecret": {
						SchemaProps: spec.SchemaProps{
							Description: "PagerDutySecret is the name of a secret in the installation namespace containing PagerDuty account details. The secret must contain the following fields:\n\nserviceKey",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"deadMansSnitchSecret": {
						SchemaProps: spec.SchemaProps{
							Description: "DeadMansSnitchSecret is the name of a secret in the installation namespace containing connection details for Dead Mans Snitch. The secret must contain the following fields:\n\nurl",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"type", "namespacePrefix"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/integreatly/v1alpha1/.PullSecretSpec"},
	}
}

func schema_apis_integreatly_v1alpha1__RHMIStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RHMIStatus defines the observed state of Installation",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"stages": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELDS - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/integreatly/v1alpha1/.RHMIStageStatus"),
									},
								},
							},
						},
					},
					"stage": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"preflightStatus": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"preflightMessage": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"lastError": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"gitHubOAuthEnabled": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"smtpEnabled": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
				},
				Required: []string{"stages", "stage", "lastError"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/integreatly/v1alpha1/.RHMIStageStatus"},
	}
}
