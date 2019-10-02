// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/wildfly/v1alpha1.PodStatus":               schema_pkg_apis_wildfly_v1alpha1_PodStatus(ref),
		"./pkg/apis/wildfly/v1alpha1.StandaloneConfigMapSpec": schema_pkg_apis_wildfly_v1alpha1_StandaloneConfigMapSpec(ref),
		"./pkg/apis/wildfly/v1alpha1.StorageSpec":             schema_pkg_apis_wildfly_v1alpha1_StorageSpec(ref),
		"./pkg/apis/wildfly/v1alpha1.WildFlyServer":           schema_pkg_apis_wildfly_v1alpha1_WildFlyServer(ref),
		"./pkg/apis/wildfly/v1alpha1.WildFlyServerSpec":       schema_pkg_apis_wildfly_v1alpha1_WildFlyServerSpec(ref),
		"./pkg/apis/wildfly/v1alpha1.WildFlyServerStatus":     schema_pkg_apis_wildfly_v1alpha1_WildFlyServerStatus(ref),
	}
}

func schema_pkg_apis_wildfly_v1alpha1_PodStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PodStatus defines the observed state of pods running the WildFlyServer application",
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"podIP": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"state": {
						SchemaProps: spec.SchemaProps{
							Description: "Represent the state of the Pod, it's used especially during scale down the expected values are represented by the PodState* constants\n\nRead-only.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"name", "podIP", "state"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_wildfly_v1alpha1_StandaloneConfigMapSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StandaloneConfigMapSpec defines the desired configMap configuration to obtain the standalone configuration for WildFlyServer",
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"key": {
						SchemaProps: spec.SchemaProps{
							Description: "Key of the config map whose value is the standalone XML configuration file (\"standalone.xml\" if omitted)",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"name"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_wildfly_v1alpha1_StorageSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageSpec defines the desired storage for WildFlyServer",
				Properties: map[string]spec.Schema{
					"emptyDir": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.EmptyDirVolumeSource"),
						},
					},
					"volumeClaimTemplate": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/core/v1.PersistentVolumeClaim"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.EmptyDirVolumeSource", "k8s.io/api/core/v1.PersistentVolumeClaim"},
	}
}

func schema_pkg_apis_wildfly_v1alpha1_WildFlyServer(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WildFlyServer is the Schema for the wildflyservers API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
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
							Ref: ref("./pkg/apis/wildfly/v1alpha1.WildFlyServerSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/wildfly/v1alpha1.WildFlyServerStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/wildfly/v1alpha1.WildFlyServerSpec", "./pkg/apis/wildfly/v1alpha1.WildFlyServerStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_wildfly_v1alpha1_WildFlyServerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WildFlyServerSpec defines the desired state of WildFlyServer",
				Properties: map[string]spec.Schema{
					"applicationImage": {
						SchemaProps: spec.SchemaProps{
							Description: "ApplicationImage is the name of the application image to be deployed",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Replicas is the desired number of replicas for the application",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"sessionAffinity": {
						SchemaProps: spec.SchemaProps{
							Description: "SessionAffinity defines if connections from the same client ip are passed to the same WildFlyServer instance/pod each time (false if omitted)",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"disableHTTPRoute": {
						SchemaProps: spec.SchemaProps{
							Description: "DisableHTTPRoute disables the creation a route to the HTTP port of the application service (false if omitted)",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"standaloneConfigMap": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/wildfly/v1alpha1.StandaloneConfigMapSpec"),
						},
					},
					"storage": {
						SchemaProps: spec.SchemaProps{
							Description: "StorageSpec defines specific storage required for the server own data directory. If omitted, an EmptyDir is used (that will not persist data across pod restart).",
							Ref:         ref("./pkg/apis/wildfly/v1alpha1.StorageSpec"),
						},
					},
					"serviceAccountName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"envFrom": {
						SchemaProps: spec.SchemaProps{
							Description: "EnvFrom contains environment variables from a source such as a ConfigMap or a Secret",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.EnvFromSource"),
									},
								},
							},
						},
					},
					"env": {
						SchemaProps: spec.SchemaProps{
							Description: "Env contains environment variables for the containers running the WildFlyServer application",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.EnvVar"),
									},
								},
							},
						},
					},
					"secrets": {
						SchemaProps: spec.SchemaProps{
							Description: "Secrets is a list of Secrets in the same namespace as the WildFlyServer object, which shall be mounted into the WildFlyServer Pods. The Secrets are mounted into /etc/secrets/<secret-name>.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"configMaps": {
						SchemaProps: spec.SchemaProps{
							Description: "ConfigMaps is a list of ConfigMaps in the same namespace as the WildFlyServer object, which shall be mounted into the WildFlyServer Pods. The ConfigMaps are mounted into /etc/configmaps/<configmap-name>.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"applicationImage", "replicas"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/wildfly/v1alpha1.StandaloneConfigMapSpec", "./pkg/apis/wildfly/v1alpha1.StorageSpec", "k8s.io/api/core/v1.EnvFromSource", "k8s.io/api/core/v1.EnvVar"},
	}
}

func schema_pkg_apis_wildfly_v1alpha1_WildFlyServerStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WildFlyServerStatus defines the observed state of WildFlyServer",
				Properties: map[string]spec.Schema{
					"replicas": {
						SchemaProps: spec.SchemaProps{
							Description: "Replicas is the actual number of replicas for the application",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"pods": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/wildfly/v1alpha1.PodStatus"),
									},
								},
							},
						},
					},
					"hosts": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"scalingdownPods": {
						SchemaProps: spec.SchemaProps{
							Description: "Represents the number of pods which are in scaledown process what particular pod is scaling down can be verified by PodStatus\n\nRead-only.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
				Required: []string{"replicas", "scalingdownPods"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/wildfly/v1alpha1.PodStatus"},
	}
}
