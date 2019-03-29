# secret-mangement-operator

### Prerequisites
[go 1.10.8 or later](https://golang.org/dl/)

[operator sdk 0.6.0](https://github.com/operator-framework/operator-sdk/tree/v0.6.x)

[minishift](https://github.com/minishift/minishift)

### Start Developing Locally

```
dep ensure
minishift start
oc login ...
export OPERATOR_NAME=secret-operator
operator-sdk up local
```
