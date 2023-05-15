#!/bin/bash

op="$1"

if [ -z "$op" ]
then
	echo "First argument should be the operation you want to perform: get|apply|stop|log|exec"
	exit 1
fi

target="$2"

if [ "$op" == "log" ] || [ "$op" == "exec" ]
then
	if [ -z "$target" ]
	then
		echo "Second argument should be the operating target: ctl|vp"
		exit 1
	fi
	if [ "$target" != "ctl" ] && [ "$target" != "vp" ]
	then
		echo "No such operating target"
		exit 1
	fi
	shift 2
else
	shift 1
fi

if [ "$op" == "get" ]
then
	kubectl -n vproxy-system get pod "$@" -o wide | grep vproxy-gateway
elif [ "$op" == "apply" ]
then
	kubectl apply -f misc/k8s-vproxy.yaml
elif [ "$op" == "stop" ]
then
	kubectl -n vproxy-system delete deployment vproxy-gateway
elif [ "$op" == "log" ] || [ "$op" == "exec" ]
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

	if [ "$target" == "ctl" ]
	then
		container="controller"
	elif [ "$target" == "vp" ]
	then
		container="vproxy"
	else
		echo "BUG 2"
		exit 1
	fi

	if [ "$op" == "log" ]
	then
		extra_log_cmd="--tail=50 -f"
		if [ ! -z "$1" ]
		then
			extra_log_cmd="$@"
		fi
		kubectl -n vproxy-system logs $extra_log_cmd "$pod" "$container"
	elif [ "$op" == "exec" ]
	then
		if [ "$target" == "ctl" ]
		then
			exec_cmd="/bin/bash"
		elif [ "$target" == "vp" ]
		then
			exec_cmd="sh"
		else
			echo "BUG 3"
			exit 1
		fi
		if [ ! -z "$1" ]
		then
			exec_cmd="$@"
		fi
		kubectl -n vproxy-system exec -it "$pod" --container="$container" -- $exec_cmd
	else
		echo "BUG 1"
		exit 1
	fi
else
	echo "No such operation: $op"
	exit 1
fi
