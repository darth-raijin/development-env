# Laboratory for Temporal
Sample Go application that connects to a Temporal server and processes a payment flow with
[Nets](https://www.nets.eu/dk-da)

## How to use
Idk

## Starting temporal server
However you do is up to you. 
* Docker image
* Service in cluster
* Temporal CLI

Simplest way is CLI

```
brew install temporal 
```

Followed by:

```
temporal server start-dev
```

## Environment configuration
In a .env have the following, alternatively a configmap in the .yaml

```
NETS_API_KEY=
NETS_SECRET_KEY=
```

Fetch the tokens from: https://portal.dibspayment.eu/integration


## Additional Resources

For more information on Temporal, refer to the official documentation:

- [Temporal Documentation](https://docs.temporal.io/)
- [Temporal Server Repository](https://github.com/temporalio/temporal)
- [Temporal Go SDK](https://github.com/temporalio/sdk-go)