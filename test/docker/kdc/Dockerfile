FROM ubuntu
RUN apt-get update && apt-get install -y --no-install-recommends krb5-kdc krb5-admin-server expect && apt-get clean
ADD krb5.conf /etc/krb5.conf
ADD kdc.conf /etc/krb5kdc/kdc.conf
ADD kadm5.acl /etc/krb5kdc/kadm5.acl
RUN yes "admin" | head -n 2 | kdb5_util create -r EXAMPLE.COM -s
ADD init.sh /init.sh
RUN expect /init.sh
ADD entrypoint.sh /entrypoint.sh
EXPOSE 88 750
ENTRYPOINT /entrypoint.sh