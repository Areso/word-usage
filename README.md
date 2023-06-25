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
  
### exposing
in case of apache2 installed:  
add the following block to /etc/apache2/defaul-ssl.conf:  
```
<VirtualHost freq.english.areso.pro:8099>
    ProxyPreserveHost On
    ProxyPass / http://127.0.0.1:8090/
    ProxyPassReverse / http://127.0.0.1:8090/
    <Directory /usr/lib/cgi-bin>
        SSLOptions +StdEnvVars
    </Directory>
    SSLCertificateFile /etc/letsencrypt/live/freq.english.areso.pro/fullchain.pem
    SSLCertificateKeyFile /etc/letsencrypt/live/freq.english.areso.pro/privkey.pem
    Include /etc/letsencrypt/options-ssl-apache.conf
    ErrorLog ${APACHE_LOG_DIR}/freq_english_error.log
    CustomLog ${APACHE_LOG_DIR}/freq_english_access.log combined
</VirtualHost>
```
after that create file `/etc/apache2/freq.english.areso.pro.conf` with the same content  
run `sudo a2ensite freq.english.areso.pro`  
run `sudo certbot --apache --renew-by-default --no-redirect -d freq.english.areso.pro`  
add recurrent task `crontab -e` add `0 1 1 */2 * sudo certbot --apache --renew-by-default --no-redirect -d freq.english.areso.pro >> /var/log/cert/freq.english.areso.pro 2>&1`  

