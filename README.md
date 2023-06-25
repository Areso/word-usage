# word-usage
This project is a microservice for other projects. Takes a word, returns the word frequently of usage.  
  
## build
`export CGO_ENABLED=0` to get rid of `./wordusage: /lib/x86_64-linux-gnu/libc.so.6: version 'GLIBC_2.34' not found (required by ./wordusage)`  
  
## install
wget latest release binary file and dictionary file  
`chmod +x wordusage`  
create a file wordusage.service in /etc/systemd/system directory with a following contents:  
```
[Unit]
Description=WordUsage_en service
After=network.target
StartLimitIntervalSec=0

[Service]
Type=simple
Restart=always
RestartSec=1
ExecStart=/root/go_services/wordusage
Environment=/root/go_services

[Install]
WantedBy=multi-user.target
```
run `systemctl daemon-reload`  
run `systemctl start wordusage`  
run `curl localhost:8090/get_freq?word=big`  
expected output: `330`  
