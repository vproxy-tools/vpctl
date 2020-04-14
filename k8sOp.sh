#!/bin/bash

op="$1"

if [ -z "$op" ]
then
	echo "First argument should be the operation you want to perform: get|apply|stop|log-ctl|log-vp|exec-ctl|exec-vp"
	exit 1
fi

if [ ! -z "$2" ]
then
	echo "Too many arguments"
	exit 1
fi

if [ "$op" == "get" ]
then
	kubectl -n vproxy-system get pod -o wide | grep vproxy-gateway
elif [ "$op" == "apply" ]
then
	kubectl apply -f misc/k8s-vproxy.yaml
elif [ "$op" == "stop" ]
then
	kubectl -n vproxy-system delete deployment vproxy-gateway
elif [ "$op" == "log-ctl" ] || [ "$op" == "log-vp" ] || [ "$op" == "exec-ctl" ] || [ "$op" == "exec-vp" ]
then
	pod=`kubectl -n vproxy-system get pod | grep vproxy-gateway | grep Running | awk '{print $1}'`
	c=`echo "$pod" | wc -l`
	if [ "$c" == 0 ]
	then
		echo "No running vproxy-gateway yet"
		exit 1
	elif [ "$c" -gt "1" ]
	then
		echo "More than one running vproxy-gateway"
		exit 1
	fi
	if [ "$op" == "log-ctl" ]
	then
		kubectl -n vproxy-system logs --tail=50 -f "$pod" controller
	elif [ "$op" == "log-vp" ]
	then
		kubectl -n vproxy-system logs --tail=50 -f "$pod" vproxy
	elif [ "$op" == "exec-ctl" ]
	then
		kubectl -n vproxy-system exec -it "$pod" --container=controller /bin/bash
	elif [ "$op" == "exec-vp" ]
	then
		kubectl -n vproxy-system exec -it "$pod" --container=vproxy sh
	else
		echo "BUG"
		exit 1
	fi
else
	echo "No such operation: $op"
	exit 1
fi
