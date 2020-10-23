package app

import (
	corev1 "k8s.io/api/core/v1"

	api "github.com/percona/percona-xtradb-cluster-operator/pkg/apis/pxc/v1"
)

func PMMClient(spec *api.PMMSpec, secrets string, v120OrGreater bool, v170OrGreater bool) corev1.Container {
	ports := []corev1.ContainerPort{{ContainerPort: 7777}}

	for i := 30100; i <= 30105; i++ {
		ports = append(ports, corev1.ContainerPort{ContainerPort: int32(i)})
	}

	pmmEnvs := []corev1.EnvVar{
		{
			Name:  "PMM_SERVER",
			Value: spec.ServerHost,
		},
	}

	clientEnvs := []corev1.EnvVar{
		{
			Name:  "CLIENT_PORT_LISTEN",
			Value: "7777",
		},
		{
			Name:  "CLIENT_PORT_MIN",
			Value: "30100",
		},
		{
			Name:  "CLIENT_PORT_MAX",
			Value: "30105",
		},
	}

	if spec.ServerUser != "" {
		pmmEnvs = append(pmmEnvs, pmmEnvServerUser(spec.ServerUser, secrets)...)
	}

	container := corev1.Container{
		Name:            "pmm-client",
		Image:           spec.Image,
		ImagePullPolicy: spec.ImagePullPolicy,
		Env:             pmmEnvs,
		SecurityContext: spec.ContainerSecurityContext,
	}

	if v120OrGreater {
		container.Env = append(container.Env, clientEnvs...)
		container.Ports = ports
	}

	if v170OrGreater {
		container.Env = append(container.Env, pmmAgentEnvs(spec.ServerHost, spec.ServerUser, secrets)...)
	}

	return container
}

func pmmAgentEnvs(pmmServerHost, pmmServerUser, secrets string) []corev1.EnvVar {
	return []corev1.EnvVar{
		{
			Name:  "PMM_AGENT_SERVER_ADDRESS",
			Value: pmmServerHost,
		},
		{
			Name:  "PMM_AGENT_LISTEN_PORT",
			Value: "7777",
		},
		{
			Name:  "PMM_AGENT_PORTS_MIN",
			Value: "30100",
		},
		{
			Name:  "PMM_AGENT_PORTS_MAX",
			Value: "30105",
		},
		{
			Name: "PMM_AGENT_SETUP_NODE_ADDRESS",
			ValueFrom: &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					FieldPath: "status.podIP",
				},
			},
		},
		{
			Name: "PMM_AGENT_SETUP_CONTAINER_ID",
			ValueFrom: &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					FieldPath: "metadata.name",
				},
			},
		},
		{
			Name: "PMM_AGENT_SETUP_NODE_NAME",
			ValueFrom: &corev1.EnvVarSource{
				FieldRef: &corev1.ObjectFieldSelector{
					FieldPath: "metadata.name",
				},
			},
		},
		{
			Name:  "PMM_AGENT_SETUP",
			Value: "1",
		},
		{
			Name:  "PMM_AGENT_SETUP_FORCE",
			Value: "1",
		},
		{
			Name:  "PMM_AGENT_SERVER_INSECURE_TLS",
			Value: "1",
		},
		{
			Name:  "PMM_AGENT_CONFIG_FILE",
			Value: "/usr/local/percona/pmm2/config/pmm-agent.yaml",
		},
		{
			Name:  "PMM_AGENT_PRERUN_FILE",
			Value: "/var/lib/mysql/setup-pmm-agent.sh",
		},
		{
			Name:  "PMM_AGENT_SERVER_USERNAME",
			Value: pmmServerUser,
		},
		{
			Name: "PMM_AGENT_SERVER_PASSWORD",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: SecretKeySelector(secrets, "pmmserver"),
			},
		},
	}
}

func pmmEnvServerUser(user, secrets string) []corev1.EnvVar {
	return []corev1.EnvVar{
		{
			Name:  "PMM_USER",
			Value: user,
		},
		{
			Name: "PMM_PASSWORD",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: SecretKeySelector(secrets, "pmmserver"),
			},
		},
	}
}
