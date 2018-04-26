# k8s_study

System requirements:
1. Ubuntu 16.04 x64 kernel 4.13.0-38-generic
2. git clone https://github.com/kubernetes-incubator/kubespray.git
   
   #2.1  kubespray dependecies:

      Ansible v2.4 (or newer)

      Jinja 2.9 (or newer)

      python-netaddr installed on the machine that running Ansible commands
   
      Target servers must have access to the Internet in order to pull docker images
      
      Target servers are configured to allow IPv4 forwarding
      
      Target servers have SSH connectivity ( tcp/22 ) directly to your nodes or through a bastion host/ssh jump box
      Target servers have a privileged user
      Your SSH key must be copied to all the servers that are part of your inventory
      Firewall rules configured properly to allow Ansible and Kubernetes components to communicate
      If using a cloud provider, you must have the appropriate credentials available and exported as environment variables
   
      By default, Vagrant uses Ubuntu 16.04 box to provision a local cluster. You may use an alternative supported operating     system for your local cluster.

Adding prometheus:

wget https://raw.githubusercontent.com/bibinwilson/kubernetes-prometheus/master/config-map.yaml
kubectl create -f ./config-map.yaml -n monitoring
wget #prometheus-deployment.yaml
kubectl create  -f prometheus-deployment.yaml --namespace=monitoring
kubectl get deployments --namespace=monitoring
kubectl get pods --namespace=monitoring
kubectl port-forward prometheus-monitoring-XXXXXX-YYYYY 8080:9090

      To access the Prometheus dashboard over a IP or a DNS name, you need to expose it as kubernetes service.

      Create a file named prometheus-service.yaml and copy the following contents. We will expose Prometheus on all kubernetes            node IP’s on port 30000.
     apiVersion: v1
         kind: Service
         metadata:
       name: prometheus-service
       spec:
      selector: 
       app: prometheus-server
      type: NodePort
      ports:
       - port: 8080
         targetPort: 9090 
         nodePort: 30000
        
    kubectl create -f prometheus-service.yaml --namespace=monitoring
