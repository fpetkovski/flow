# Flow Query Language

> **Note**: This is an experimental proof-of-concept project exploring alternative syntax for PromQL queries.
>
> **Try it out**: [Interactive Playground](https://github.com/fpetkovski/flow)

Flow is a query language that transpiles to PromQL (Prometheus Query Language). When writing PromQL queries, users typically start by exploring a metric and its labels, then add transformations like rate calculations and aggregations. However, PromQL's nested function syntax forces you to jump back and forth in the query—adding a new transformation means wrapping the entire expression in yet another function. Flow solves this by providing a left-to-right pipeline syntax using the `|` operator, allowing you to write queries in the same order you think about them: start with a metric, then apply transformations step by step.

The language supports all PromQL features including metric selection, label matching, aggregations, range functions, and binary operations. It also introduces `let` expressions for defining reusable query fragments, reducing duplication and improving maintainability. Flow is available both as a Go library and as a WebAssembly module for client-side transpilation.

## Examples

### Basic metric selection and aggregation
```
# Flow
http_requests_total{method="POST"} | rate 5m | sum by (endpoint)

# PromQL
sum by (endpoint) (rate(http_requests_total{method="POST"}[5m]))
```

### Let expressions for reusable queries
```
# Flow
let
  req_get = http_requests_total{method="GET"} | rate 5m | sum by (endpoint),
  req_total = http_requests_total | rate 5m | sum by (endpoint)
in req_get / req_total

# PromQL
sum by (endpoint) (rate(http_requests_total{method="GET"}[5m]))
/
sum by (endpoint) (rate(http_requests_total[5m]))
```

### Complex calculations with binary operations
```
# Flow
let
  req_get = http_requests_total{method="GET"} | rate 5m | sum by (endpoint),
  req_total = http_requests_total | rate 5m | sum by (endpoint)
in 100 * (req_total - req_get) / req_total

# PromQL
100 * (
  sum by (endpoint) (rate(http_requests_total[5m]))
  -
  sum by (endpoint) (rate(http_requests_total{method="GET"}[5m]))
)
/
sum by (endpoint) (rate(http_requests_total[5m]))
```

### Functions and histogram operations
```
# Flow
let
  avail = kube_deployment_status_replicas_available{namespace="monitoring"},
  loop_time = looping_time{group_name="realtime"} | rate 2m | sum by () | histogram_sum,
  sleep_time = sleeping_time{group_name="realtime"} | rate 2m | sum by () | histogram_sum
in (avail > 1) * loop_time / (loop_time + sleep_time)

# PromQL
(kube_deployment_status_replicas_available{namespace="monitoring"} > 1)
*
histogram_sum(sum(rate(looping_time{group_name="realtime"}[2m])))
/
(
  histogram_sum(sum(rate(looping_time{group_name="realtime"}[2m])))
  +
  histogram_sum(sum(rate(sleeping_time{group_name="realtime"}[2m])))
)
```
