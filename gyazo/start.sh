#!/bin/bash

chown daemon:daemon /usr/local/apache2/htdocs/data/
chown daemon:daemon /usr/local/apache2/htdocs/db/

(echo -n "seigo2016:realm:" && echo -n "seigo2016:realm:$password" | md5sum - | cut -d' ' -f1) > /.htdigest

chown daemon:daemon /.htdigest
chmod 400 /.htdigest

httpd -D FOREGROUND
