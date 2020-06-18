# CI for all branches except Master
cloudbuild.yaml

Insert on GCP the following regex for exclude master branch from CI
```bash
^(([^m]|m($|[^a]|a($|[^s]|s($|[^t]|t($|[^e]|e($|[^r]))))))|master.+)
```


# CI+CD only for master branch
cloudbuild.prod.yaml

