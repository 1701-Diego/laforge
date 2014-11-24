# Laforge - Engineering Containerized Experiments

Laforge includes a handful of LRP and Task examples

## To Install

```
go get github.com/1701-diego/laforge
```

## To Use

First export the `RECEPTOR` address

For Ketchup:
```
export RECEPTOR=http://username:password@receptor.ketchup.cf-app.com
```

For Diego-Edge:
```
export RECEPTOR=http://receptor.192.168.11.11.xip.io
```

Then, to view a list of experiments:

```
laforge
```

To run an experiment:
```
laforge EXPERIMENT DOMAIN
```

Where DOMAIN is your team's chosen domain.