myhttp
======

`myhttp` is a tool that can make parallel GET requests to the given URLs. It will
return the MD5 format of the response body along with the URL. If a request
fails, it will provide a message to inform the user about the failure. `myhttp`
quarentees that the output is consistent for every request.

Installation
------------

Visit the [Releases](https://github.com/huseyin/myhttp/releases) page for pre-built
binaries.

If you prefer to build it manually, please follow the below instructions.

```sh
make build
```

To remove the leftovers which are generated after build step, please run the
below command.

Note that this will remove the binary file as well.

```sh
make clean
```

Usage
-----

Please have a look at the help text of `myhttp` for more information.

```sh
➜ myhttp -help
Usage of myhttp:
  -parallel int
        The number of parallel requests. (default 10)
```

To make a request without additional parameters:

```sh
myhttp https://example.com
```

`myhttp` accepts multiple URLs.

```sh
myhttp https://example.com https://golang.org
```

`myhttp` also allows you to make concurrent requests, `10` by default. It also
handles signals sent at runtime, and shuts down the connections gracefully.

You can indicate the concurrency level by this:

```sh
myhttp -parallel=5 https://example.com
```

License
-------

Not licensed.
