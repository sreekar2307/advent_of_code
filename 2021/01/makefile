go = /usr/local/bin/go

one_one = /tmp/advent_of_code/2021/01/01
one_two = /tmp/advent_of_code/2021/01/02
one_sample:
	$(go) build -o $(one_one)/main 01/main.go && \
    cp 01/sample.txt $(one_one)/sample.txt && \
    $(one_one)/main $(one_one)/sample.txt
one_input:
	$(go) build -o $(one_one)/main 01/main.go && \
    cp 01/input.txt $(one_one)/input.txt && \
    $(one_one)/main $(one_one)/input.txt
two_sample:
	$(go) build -o $(one_two)/main 02/main.go && \
    cp 02/sample.txt $(one_two)/sample.txt && \
    $(one_two)/main $(one_two)/sample.txt
two_input:
	$(go) build -o $(one_two)/main 02/main.go && \
    cp 02/input.txt $(one_two)/input.txt && \
    $(one_two)/main $(one_two)/input.txt
clean:
	rm -r /tmp/advent_of_code || exit 0