FROM anapsix/alpine-java

RUN apk add --update unzip wget curl docker jq go

WORKDIR /opt/kafka

ADD http://apache.claz.org/kafka/0.10.2.0/kafka-0.10.2.0-src.tgz ./

RUN tar -zxf kafka-0.10.2.0-src.tgz && \
  cp -r /opt/kafka/kafka-0.10.2.0-src/bin/* /usr/local/bin

ADD systemcalls.go /opt
