FROM golang:1.13 as build_revtrvp
ADD . /go/src/github.com/NEU-SNS/revtrvp
WORKDIR /go/src/github.com/NEU-SNS/revtrvp
RUN go build -o revtrvp .
RUN chmod -R a+rx /go/src/github.com/NEU-SNS/revtrvp/revtrvp

FROM ubuntu:18.04 as build_scamper

RUN apt-get update && \
    apt-get install -y make coreutils autoconf libtool git build-essential wget && \
    apt-get clean && \
    rm -rf /var/lib/opt/lists/*

RUN ls -l
RUN mkdir -p scamper-src && cd scamper-src && \
    wget http://www.ccs.neu.edu/home/rhansen2/scamper.tar.gz && \
    tar xzf scamper.tar.gz && cd scamper-cvs-20150901
WORKDIR /scamper-src/scamper-cvs-20150901/
RUN ./configure && make install


#WORKDIR /plvp
#COPY . /plvp

#RUN useradd -ms /bin/bash plvp
FROM ubuntu:18.04

RUN dpkg --add-architecture i386

RUN apt-get update && apt-get install -y \
    wget \
    build-essential \
    libc6:i386 \
    libncurses5:i386 \
    libstdc++6:i386 \
    iputils-ping \
    inetutils-traceroute \
    init-system-helpers \
&&  apt-get clean \
&&  rm -rf /var/lib/apt/lists/*

COPY --from=build_revtrvp /go/src/github.com/NEU-SNS/revtrvp/revtrvp /
COPY --from=build_revtrvp /go/src/github.com/NEU-SNS/revtrvp/root.crt /
COPY --from=build_revtrvp /go/src/github.com/NEU-SNS/revtrvp/plvp.config /
COPY --from=build_scamper /usr/local/bin/scamper /usr/local/bin

RUN ldconfig
RUN which scamper
WORKDIR /

ENTRYPOINT ["/revtrvp"]
CMD ["/root.crt", "plvp.config", "-loglevel", "error"]

EXPOSE 4381
