![](http://counter.gofiber.io/badge/hi019/better-view-counter)

## Usage
better-view-counter generates a view count badge for your repository's README. It supports a custom label and unique views only. To use it, first create a URL:

https://counter.gofiber.io/badge/YOUR_USERNAME/YOUR_REPO

*Options*
* Only count unique views (by IP): `{URL}?unique=true`


## Why another view counter?
When assessing other view conuters, we found they were capped at a limited number of requests per hour before the badge returned an error. This is not ideal for high-traffic repositories. Through [Fiber](https://gofiber.io), *better-view-counter* can handle about 30-35k requests per second with no cap (but please don't benchmark the main instance!).

## Installing yourself
Head over to the [releases](https://github.com/hi019/better-view-counter/releases) page to get an executable for your platofrm, then simply run it. By default the port is set to `3000`, you can change it by doing `./viewcounter -port 80`.

## Building 
To build,
1. Clone project, cd into directory
2. `go build`
3. Result will be `./viewcounter`