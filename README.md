# k8s_study

DevOPS task:
1. Ubuntu 16.04 x64 kernel 4.13.0-38-generic
2. Use virtualization.sh to install vagrant and virtual box entities
3. Spin Uping of k8s cluster :git clone https://github.com/kubernetes-incubator/kubespray.git
   
   further instructions and prerequcities can be found in : https://github.com/kubernetes-incubator/kubespray

4. Adding prometheus entities and instructions:
https://devopscube.com/setup-prometheus-monitoring-on-kubernetes/

if in log are problems with RBAC please use:
rbac-setup.yaml - for account prometheus
drbac-setup.yaml - for account default

5. According to https://codefresh.io/docs/docs/deploy-to-kubernetes/access-codefresh-docker-registry-from-kubernetes/
create access to docker registry for default usage

6. According to https://godoc.org/github.com/prometheus/client_golang/prometheus
 created golang file is in attachment.

7. 

8. Docker image build file
