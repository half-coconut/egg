

https://github.com/librant/git-ops-learn

https://hub.docker.com/repository/docker/halfcoconut/hello-egg/general


- commond
```shell
(.venv) chenchen@chenchendeMacBook-Pro-2 egg % history
 1096  kubectl get pods
 1097  while true; do sleep 1; curl http://127.0.0.1; echo -e '\n'$(date);done
 1098  kubectl get pods
 1099  kubectl exec -it hello-egg-6d97d7497f-gf6kz -- bash 
 1100  kubectl get pods
 1101  while true; do sleep 1; curl http://127.0.0.1; echo -e '\n'$(date);done
 1102  kubectl get pods
 1103  kubectl apply -f https://ghproxy.com/https://raw.githubusercontent.com/lyzhang1999/resource/main/metrics/metrics.yaml
 1104  kubectl apply -f https://ghproxy.com/https://raw.githubusercontent.com/lyzhang1999/resource/main/metrics/metrics.yaml
 1105  kubectl wait deployment -n kube-system metrics-server --for condition=Available=True --timeout=90s
 1106  kubectl autoscale deployment hello-egg --cpu-percent=50 --min=2 --max=10
 1107  kubectl get pods
 1108  kubectl patch deployment hello-egg --type='json' -p='[{"op": "add", "path": "/spec/template/spec/containers/0/resources", "value": {"requests": {"memory": "100Mi", "cpu": "100m"}}}]'\ndeployment.apps/hello-world-flask patched
 1109  kubectl patch deployment hello-egg --type='json' -p='[{"op": "add", "path": "/spec/template/spec/containers/0/resources", "value": {"requests": {"memory": "100Mi", "cpu": "100m"}}}]'
 1110  kubectl get pods
 1111  kubectl exec -it hello-egg-5df44f68d-vx58w -- bash
```
