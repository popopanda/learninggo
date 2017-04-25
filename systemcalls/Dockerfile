FROM anapsix/alpine-java

ENV KAFKA_HOME /opt/kafka/kafka_2.10-0.10.2.0
ENV PATH ${PATH}:${KAFKA_HOME}/bin

RUN apk add --update curl docker jq go sudo

WORKDIR /opt/kafka

ADD http://mirrors.koehn.com/apache/kafka/0.10.2.0/kafka_2.10-0.10.2.0.tgz ./

RUN tar -zxf kafka_2.10-0.10.2.0.tgz

ADD systemcalls.go /opt

CMD ["/usr/bin/go", "run", "/opt/systemcalls.go"]
