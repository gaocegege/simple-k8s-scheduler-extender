# Simple K8s Scheduler Extender

use the config file in config/ to create the scheduler, and run simple-k8s-scheduler-extender.

```
kube-scheduoler --policy-config-file='$(pwd)/config/policy.json'
cd $GOPATH/src/github.com/gaocegege/simple-k8s-scheduler-extender
godep go build .
./simple-k8s-scheduler-extender
```

## NOTICE

The algorithm is a demo.
