FROM httpd:2.4

RUN apt-get update
RUN apt-get -y install ruby wget git

#RUN install -d -m 644 -g daemon -o daemon /usr/local/apache2/htdocs/data/
#RUN install -d -m 644 -g daemon -o daemon /usr/local/apache2/htdocs/db/

COPY ./upload.cgi /usr/local/apache2/cgi-bin/upload.cgi
RUN chown daemon:daemon /usr/local/apache2/cgi-bin/upload.cgi
RUN chmod +x /usr/local/apache2/cgi-bin/upload.cgi


COPY ./apache2.conf /usr/local/apache2/conf/httpd.conf

RUN apt-get clean && rm -rf /var/cache/apt/archives/* /var/lib/apt/lists/*

WORKDIR /

COPY start.sh /start.sh

RUN chown daemon:daemon /start.sh
RUN chmod +x /start.sh

CMD ["/start.sh"]

