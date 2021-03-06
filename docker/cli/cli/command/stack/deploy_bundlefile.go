package stack

import (
	"golang.org/x/net/context"

	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/swarm"
	"github.com/appcelerator/amp/docker/cli/cli/command"
	"github.com/appcelerator/amp/docker/cli/cli/compose/convert"
)

func deployBundle(ctx context.Context, dockerCli command.Cli, opts DeployOptions) error {
	bundle, err := loadBundlefile(dockerCli.Err(), opts.Namespace, opts.Bundlefile)
	if err != nil {
		return err
	}

	if err := checkDaemonIsSwarmManager(ctx, dockerCli); err != nil {
		return err
	}

	namespace := convert.NewNamespace(opts.Namespace)

	if opts.Prune {
		services := map[string]struct{}{}
		for service := range bundle.Services {
			services[service] = struct{}{}
		}
		pruneServices(ctx, dockerCli, namespace, services)
	}

	networks := make(map[string]types.NetworkCreate)
	for _, service := range bundle.Services {
		for _, networkName := range service.Networks {
			networks[networkName] = types.NetworkCreate{
				Labels: convert.AddStackLabel(namespace, nil),
			}
		}
	}

	services := make(map[string]swarm.ServiceSpec)
	for internalName, service := range bundle.Services {
		name := namespace.Scope(internalName)

		var ports []swarm.PortConfig
		for _, portSpec := range service.Ports {
			ports = append(ports, swarm.PortConfig{
				Protocol:   swarm.PortConfigProtocol(portSpec.Protocol),
				TargetPort: portSpec.Port,
			})
		}

		nets := []swarm.NetworkAttachmentConfig{}
		for _, networkName := range service.Networks {
			nets = append(nets, swarm.NetworkAttachmentConfig{
				Target:  namespace.Scope(networkName),
				Aliases: []string{internalName},
			})
		}

		serviceSpec := swarm.ServiceSpec{
			Annotations: swarm.Annotations{
				Name:   name,
				Labels: convert.AddStackLabel(namespace, service.Labels),
			},
			TaskTemplate: swarm.TaskSpec{
				ContainerSpec: &swarm.ContainerSpec{
					Image:   service.Image,
					Command: service.Command,
					Args:    service.Args,
					Env:     service.Env,
					// Service Labels will not be copied to Containers
					// automatically during the deployment so we apply
					// it here.
					Labels: convert.AddStackLabel(namespace, nil),
				},
			},
			EndpointSpec: &swarm.EndpointSpec{
				Ports: ports,
			},
			Networks: nets,
		}

		services[internalName] = serviceSpec
	}

	if err := createNetworks(ctx, dockerCli, namespace, networks); err != nil {
		return err
	}
	return deployServices(ctx, dockerCli, services, namespace, opts.SendRegistryAuth, opts.ResolveImage)
}
