![](http://counter.gofiber.io/badge/hi019/better-view-counter)

## Usage
better-view-counter generates a view count badge for your repository's README. It supports a custom label and unique views only. To use it, first create a URL:

https://counter.gofiber.io/badge/YOUR_USERNAME/YOUR_REPO

*Options*
* Only count unique views (by IP): `{URL}?unique=true`

Then in your README, embed it in an svg: `![](https://counter.gofiber.io/badge/YOUR_USERNAME/YOUR_REPO)`


## Why another view counter?
When assessing other view conuters, we found they were capped at a limited number of requests per hour before the badge returned an error. This is not ideal for high-traffic repositories. Through [Fiber](https://gofiber.io), *better-view-counter* can handle about 30-40k requests per second with no cap on a 1 core vps (but please don't benchmark the main instance!). Memory usage is also great, staying under 25mb during heavy benchmarks.

## Installing yourself
Head over to the [releases](https://github.com/hi019/better-view-counter/releases) page to get an executable for your platofrm, then simply run it. By default the port is set to `3000`, you can change it by doing `./viewcounter -port 80`.

## Building 
To build,
1. Clone project, cd into directory
2. `go build`
3. Result will be `./viewcounter`

## Benchmarks
On a 1 core VPS:

```
./bombardier -c 750 -n 1000000 http://127.0.0.1:3000/badge/demo/demo
Statistics        Avg      Stdev        Max
  Reqs/sec     38462.53    4243.36   47245.98
  Latency       19.50ms   105.75ms      7.57s
  HTTP codes:
    1xx - 0, 2xx - 1000000, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:    52.80MB/s
```

Memory usage under 25mb.
