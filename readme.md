# Simple K8s Scheduler Extender

use the config file in config/ to create the scheduler, and run simple-k8s-scheduler-extender.

```
kube-scheduoler --policy-config-file='$(pwd)/config/policy.json'
./simple-k8s-scheduler-extender
```

## Related Bugs in K8s

ref [kubernetes/kubenetes#32652](https://github.com/kubernetes/kubernetes/pull/32652)