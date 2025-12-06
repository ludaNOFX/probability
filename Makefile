.PHONY: build clean run

build:
	go build -o bin/computer_simulation ./cmd/computer_simulation
	go build -o bin/random_variable_model ./cmd/random_variable_model

run-computer:
	./bin/computer_simulation

run-random: 
	./bin/random_variable_model

clean:
	rm -rf bin/
