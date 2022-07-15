# TODO: optimize build
CC = gcc
CFLAGS = -Wall -g

cvips: main.c
	${CC} ${CFLAGS} -o cvips main.c `pkg-config vips --cflags --libs`
