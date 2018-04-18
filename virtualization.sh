#!/bin/bash
#simplification of vagrant installation

#will using sudo user

if [[ $EUID != 0 ]];
then 
echo "This script needs sudo privelegies... Now exiting"
sleep 5
exit 1
fi

echo " Used sudo command continue vagrant setup"

apt-get install -y virtualbox
if [ $? != 0 ];
then
echo "Cannot install virtualbox infrasture"
exit 1
fi

apt-get install -y vagrant

if [ $? != 0 ];
then
echo "Cannot install vagrant infrasture"
exit 1
fi

apt-get install -y virtualbox-dkms

if [ $? != 0 ];
then
echo "Cannot install virtualbox-dkms tools"
exit 1
fi
