FROM ubuntu
RUN apt-get update &&apt-get -y install \
        gcc \
        ssh \
        krb5-user \
        libgssapi-krb5-2 \
        libkrb5-dev \
        libsasl2-modules-gssapi-mit \
        wget \
        expect && apt-get clean

RUN cd /tmp && wget https://dl.google.com/go/go1.12.1.linux-amd64.tar.gz && tar xvf go1.12.1.linux-amd64.tar.gz && mv go/ /opt
ENV GOROOT="/opt/go"
ADD krb5.conf /etc/krb5.conf
ADD init.sh /init.sh
ADD entrypoint.sh /entrypoint.sh
RUN chmod +x entrypoint.sh
ENTRYPOINT /entrypoint.sh