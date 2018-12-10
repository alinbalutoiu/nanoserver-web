# nanoserver-web
Simple web server on top of nanoserver container.

## How to build?

From the root directory of this repo run the command which applies to the operating system you are using.

### For Windows Server 2016 LTSC

```
docker build . -f sac2016/Dockerfile -t nanoserver-web:sac2016
```

### For Windows Server 2016 version 1709

```
docker build . -f 1709/Dockerfile -t nanoserver-web:1709
```

### For Windows Server 2016 version 1803

```
docker build . -f 1803/Dockerfile -t nanoserver-web:1803
```

### For Windows Server 2019 version 1809

```
docker build . -f 1809/Dockerfile -t nanoserver-web:1809
```
