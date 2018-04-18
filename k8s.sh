#!/bin/bash

# Kubernetis installation script
# according to https://kubernetes-v1-4.github.io/docs/getting-started-guides/vagrant/
OUTPUT=$(dpkg -s curl | grep Status) 

RES="Status: install ok installed"


#Setting up a cluster is as simple as running:
if [ "$OUTPUT" != "$RES" ] ;
then 
echo "Curl is not installed"
exit 1
#apt-get install curl
else
echo "Curl found"
fi

export KUBERNETES_PROVIDER=vagrant
curl -sS https://get.k8s.io | bash
