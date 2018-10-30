FROM alpine:3.7

# Use aliyun source
RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main" > /etc/apk/repositories
RUN echo "http://mirrors.aliyun.com/alpine/v3.7/community" >> /etc/apk/repositories

RUN apk update && apk add -t .base \
  curl \ 
  bash \ 
  tzdata \ 
  ca-certificates \
  keepalived \
  ipvsadm \
  iptables \
  openssh-server \
  && cp -r -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
  && rm -rf /var/lib/apk/*

RUN sed -i 's/UsePAM yes/UsePAM no/g' /etc/ssh/sshd_config

RUN echo "root:admin" | chpasswd
RUN echo "admin   ALL=(ALL)       ALL" >> /etc/sudoers

RUN ssh-keygen -t dsa -f /etc/ssh/ssh_host_dsa_key
RUN ssh-keygen -t rsa -f /etc/ssh/ssh_host_rsa_key

RUN mkdir /var/run/sshd
EXPOSE 22

CMD ["/usr/sbin/sshd", "-D"]
