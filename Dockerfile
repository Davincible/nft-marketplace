FROM 	node:18-alpine

ENV 	APP_ROOT /app

RUN 	mkdir ${APP_ROOT}
WORKDIR ${APP_ROOT}
ADD 	. ${APP_ROOT}

RUN 	apk add git
RUN 	npm install

EXPOSE 	8545

CMD 	[ "yarn", "run", "dev" ]
