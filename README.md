# yesserver
A simple http echo server that accept any http URI, this is specially useful when testing URI manipulations with a LB
It will send back the container/os name + all http headers as text

GET on /metric will return http statistics in prometheus format


## local build and run
```
docker build . -t yesserv:latest
docker run -ti --rm -p 8000:8000 yesserv
```

## run from dockerhub
```
docker pull aca2328/yesserv:latest
docker run -ti --rm -p 8000:8000 aca2328/yesserv
```

## helm deployment

the chart will install one pod replica, one svc of type cluster-ip and one ingress annotated with `kubernetes.io/ingress.class: avi-lb` ( will instruct AKO to program the ingress on AVI dataplane ) and tagged with label `app: gslb` ( this will instruct AMKO to add this ingress to a gslb group)
with this setup it is possible to deploy this app on multiple cluster to demonstrate GSLB inter cluster load-balancing.

```
helm install yeschart yesserv/chart/ --dry-run
helm install yeschart yesserv/chart/
```

## References

https://github.com/vmware/load-balancer-and-ingress-services-for-kubernetes
https://github.com/avinetworks/avi-helm-charts

