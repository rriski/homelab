# Homelab

> ⚠️ WORK IN PROGRESS

## Hardware

![Hardware](https://user-images.githubusercontent.com/27996771/98970963-25137200-2543-11eb-8f2d-f9a2d45756ef.JPG)

- 4 nodes of NEC SFF PC (Japanese version of the ThinkCentre M700)
  - CPU: Intel Core i5-6600T
  - RAM: 16GB
  - SSD: 128GB
- TP-Link TL-SG108 switch

## Technology stack

<table>
  <tr>
    <td align="center"><a><img src="https://simpleicons.org/icons/ansible.svg" width="100px;"/><br/>Ansible</td>
    <td align="center"><a><img src="https://simpleicons.org/icons/cloudflare.svg" width="100px;"/><br/>Cloudflare</td>
    <td align="center"><a><img src="https://simpleicons.org/icons/docker.svg" width="100px;"/><br/>Docker</td>
    <td align="center"><a><img src="https://simpleicons.org/icons/fedora.svg" width="100px;"/><br/>Fedora</td>
    <td align="center"><a><img src="https://simpleicons.org/icons/gitea.svg" width="100px;"/><br/>Gitea</td>
    <td align="center"><a><img src="https://simpleicons.org/icons/helm.svg" width="100px;"/><br/>Helm</td>
  </tr>
  <tr>
    <td align="center"><a><img src="https://simpleicons.org/icons/kubernetes.svg" width="100px;"/><br/>Kubernetes</td>
    <td align="center"><a><img src="https://simpleicons.org/icons/prometheus.svg" width="100px;"/><br/>Prometheus</td>
    <td align="center"><a><img src="https://simpleicons.org/icons/rancher.svg" width="100px;"/><br/>Rancher</td>
    <td align="center"><a><img src="https://simpleicons.org/icons/terraform.svg" width="100px;"/><br/>Terraform</td>
    <td align="center"><a><img src="https://simpleicons.org/icons/vault.svg" width="100px;"/><br/>Vault</td>
    <td align="center"><a><img src="https://simpleicons.org/icons/wireguard.svg" width="100px;"/><br/>Wireguard</td>
  </tr>
  <tr>
  </tr>
</table>

## Architecture

| Layer | Name                   | Description                                             | Provisioner         |
|-------|------------------------|---------------------------------------------------------|---------------------|
| 0     | [metal](./metal)       | Bare metal OS installation, Terraform state backend,... | Ansible, PXE server |
| 1     | [infra](./infra)       | Kubernetes clusters                                     | Terraform, Helm     |
| 2     | [apps](./apps)         | Gitea, Vault and more in the future                     | Argo                |

## Usage

### Prerequisite

For the controller (to run Ansible, stateless PXE server, Terraform...):

- SSH keys in `~/.ssh/{id_ed25519,id_ed25519.pub}` (you can generate it with `ssh-keygen -t ed25519`)
- Docker with `host` networking driver (which means [only Docker on Linux hosts](https://docs.docker.com/network/host/), you can use a Linux virtual machine with bridged networking if you're on macOS or Windows)

For bare metal nodes:

- PXE IPv4 enabled
- Wake-on-LAN enabled
- Secure boot disabled (optional, depending on the OS)

### Configurations

- [Bare metal nodes settings](./metal/hosts.ini) (IP, MAC...)
- [OS settings](./metal/group_vars/all.yml) (PXE, network...)

### Building

Open the tools container:

```sh
make tools
```

Then build the homelab:

```sh
make
```
