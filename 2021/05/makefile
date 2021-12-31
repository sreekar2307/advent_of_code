go = /usr/local/bin/go

five_one = /tmp/advent_of_code/2021/05/01
five_two = /tmp/advent_of_code/2021/05/02
one_sample:
	$(go) build -o $(five_one)/main 01/main.go && \
    cp 01/sample.txt $(five_one)/sample.txt && \
    $(five_one)/main $(five_one)/sample.txt
one_input:
	$(go) build -o $(five_one)/main 01/main.go && \
    cp 01/input.txt $(five_one)/input.txt && \
    $(five_one)/main $(five_one)/input.txt
two_sample:
	$(go) build -o $(five_two)/main 02/main.go && \
    cp 02/sample.txt $(five_two)/sample.txt && \
    $(five_two)/main $(five_two)/sample.txt
two_input:
	$(go) build -o $(five_two)/main 02/main.go && \
    cp 02/input.txt $(five_two)/input.txt && \
    $(five_two)/main $(five_two)/input.txt
clean:
	rm -r /tmp/advent_of_code || exit 0