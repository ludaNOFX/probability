.PHONY: build clean run

build:
	go build -o bin/computer_simulation/cmps ./cmd/computer_simulation
	go build -o bin/random_variable_model/rvm ./cmd/random_variable_model

run-computer:
	./bin/computer_simulation/cmps

run-random: 
	./bin/random_variable_model/rvm

clean:
	rm -rf bin/
