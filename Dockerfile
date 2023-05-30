FROM ubuntu:22.04

# update everything
RUN apt update -y && apt upgrade -y && apt install git -y && apt install wget

# install go
RUN wget https://dl.google.com/go/go1.20.4.linux-amd64.tar.gz && tar -C /opt -xzf go1.20.4.linux-amd64.tar.gz

# clone some codedocker system prune -a
RUN git clone https://github.com/Gurman520/battle_of_psychics

# Setup
RUN cd ./battle_of_psychics && export PATH=$PATH:/opt/go/bin && go build ./main.go

RUN cp -r ./battle_of_psychics/. .

CMD export PATH=$PATH:/opt/go/bin && go run ./main.go