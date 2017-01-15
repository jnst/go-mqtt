FROM ubuntu:16.04

MAINTAINER jnst

RUN apt-get update && apt-get install -y \
    build-essential \
    curl \
    git \
    openssh-server \
    python-software-properties \
    software-properties-common \
    unzip \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/*

RUN locale-gen ja_JP

RUN cd /mnt && curl -L -o emqttd.zip http://emqtt.io/downloads/latest/ubuntu16_04 && unzip emqttd.zip && rm emqttd.zip

RUN wget https://bootstrap.pypa.io/ez_setup.py -O - | python;
RUN wget https://bootstrap.pypa.io/get-pip.py;python get-pip.py;
RUN pip install supervisor

ADD config/supervisord.conf /etc/supervisord.conf

EXPOSE 22 1883 18083
CMD ["/usr/local/bin/supervisord"]
