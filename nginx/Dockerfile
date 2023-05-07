FROM ubuntu:lastest

RUN api-get update && apt-get install --yes nginx

EXPOSE 80

CMD [ "nginx", "-g", "demon off;" ]