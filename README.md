rubygems-proxy
==============

an example caching proxy server for rubygems using Go's standard http library.

This example code is based on [Trivial HTTP Proxy in Go — IOTTMCO](http://bytbox.net/blog/2013/01/trivial-http-proxy-golang.html).

Run the proxy.

```
./rubygems-proxy -root cache
```

Run gem install in another terminal.

```
gem install rails -p http://localhost:8080
```

The proxy log.

```
2013/09/23 22:08:32 method:HEAD url:http://rubygems.org/latest_specs.4.8.gz
2013/09/23 22:08:39 forward proxy to remote: http://rubygems.org/latest_specs.4.8.gz
2013/09/23 22:08:40 method:GET  url:http://rubygems.org/latest_specs.4.8.gz
2013/09/23 22:08:44 saved to local file: cache/latest_specs.4.8.gz
2013/09/23 22:08:44 method:HEAD url:http://rubygems.org/specs.4.8.gz
2013/09/23 22:08:49 forward proxy to remote: http://rubygems.org/specs.4.8.gz
2013/09/23 22:08:49 method:GET  url:http://rubygems.org/specs.4.8.gz
2013/09/23 22:08:57 saved to local file: cache/specs.4.8.gz
2013/09/23 22:09:01 method:GET  url:http://rubygems.org/gems/rails-4.0.0.gem
2013/09/23 22:09:12 saved to local file: cache/gems/rails-4.0.0.gem
```


Uninstall and reinstall

```
gem uninstall rails 
gem install rails -p http://localhost:8080
```

This time all responses are served from local cache.

```
2013/09/23 22:09:27 method:HEAD url:http://rubygems.org/latest_specs.4.8.gz
2013/09/23 22:09:27 served from local file: cache/latest_specs.4.8.gz
2013/09/23 22:09:27 method:GET  url:http://rubygems.org/latest_specs.4.8.gz
2013/09/23 22:09:27 served from local file: cache/latest_specs.4.8.gz
2013/09/23 22:09:27 method:HEAD url:http://rubygems.org/specs.4.8.gz
2013/09/23 22:09:27 served from local file: cache/specs.4.8.gz
2013/09/23 22:09:27 method:GET  url:http://rubygems.org/specs.4.8.gz
2013/09/23 22:09:27 served from local file: cache/specs.4.8.gz
2013/09/23 22:09:32 method:GET  url:http://rubygems.org/gems/rails-4.0.0.gem
2013/09/23 22:09:32 served from local file: cache/gems/rails-4.0.0.gem
```
