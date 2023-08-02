FROM golang:1.16-alpine

LABEL maintainer="alifsnitt@gmail.com"

WORKDIR /usr/src/app

# Update package
RUN apk add --update --no-cache --virtual .build-dev build-base git

COPY . .

RUN make install \
  && make build

# Expose port
EXPOSE 9000

# Run application
CMD ["make", "start"]

