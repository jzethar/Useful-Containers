FROM ubuntu:22.04 as ubuntu

RUN apt update && DEBIAN_FRONTEND=noninteractive apt-get install -y git screen wget vim openssh-server \ 
    curl net-tools iproute2 iputils-ping mtr traceroute && \
    mkdir /var/run/sshd && \
    echo "PasswordAuthentication no" >> /etc/ssh/sshd_config

COPY ./sshKeys /root/.ssh

CMD ["/usr/sbin/sshd", "-D"]
