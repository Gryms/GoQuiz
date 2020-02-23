SRC		=	$(addprefix src/, quiz.go)

NAME	=	bin/main

build:
	go build -o $(NAME) $(SRC)

run:
	go run $(SRC)

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o $(NAME)-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o $(NAME)-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o $(NAME)-freebsd-386 main.go

clean:
	rm -rf $(NAME)

all:	build

re:		clean all

.PHONY: all run build compile clean fclean re