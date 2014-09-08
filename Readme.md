[![GoDoc](https://godoc.org/github.com/communaute-cimi/checker/service?status.svg)](https://godoc.org/github.com/communaute-cimi/checker/service)

```sh
$ checker 
--- SocleSIG DEV ---
2014/09/08 14:29:39 apache 10.226.148.159:80 a live
2014/09/08 14:29:39 memcache 127.0.0.1:11211 down
2014/09/08 14:29:39 postgres 127.0.0.1:5432 down
```

```sh
$ checker -h
Usage of checker:
  -config="checker.json": config file
```

```json
{
    "checkers":[{
	"name":"SocleSIG DEV",
	"services":[
	    {
		"name":"apache",
		"config":{
		    "host":"10.226.148.159",
		    "port":80
		}
	    }, {
		"name":"memcache",
		"config":{
		    "host":"127.0.0.1",
		    "port":11211
		}
	    }, {
		"name":"postgres",
		"config":{
		    "host":"127.0.0.1",
		    "port":5432,
		    "user":"ja",
		    "password":"jo",
		    "db":"dn_ja"
		}
	    }
	]
    }]
}
```
