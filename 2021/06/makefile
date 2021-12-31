go = /usr/local/bin/go

six_one = /tmp/advent_of_code/2021/06/01
six_two = /tmp/advent_of_code/2021/06/02
one_sample:
	$(go) build -o $(six_one)/main 01/main.go && \
    cp 01/sample.txt $(six_one)/sample.txt && \
    $(six_one)/main $(six_one)/sample.txt
one_input:
	$(go) build -o $(six_one)/main 01/main.go && \
    cp 01/input.txt $(six_one)/input.txt && \
    $(six_one)/main $(six_one)/input.txt
two_sample:
	$(go) build -o $(six_two)/main 02/main.go && \
    cp 02/sample.txt $(six_two)/sample.txt && \
    $(six_two)/main $(six_two)/sample.txt
two_input:
	$(go) build -o $(six_two)/main 02/main.go && \
    cp 02/input.txt $(six_two)/input.txt && \
    $(six_two)/main $(six_two)/input.txt
clean:
	rm -r /tmp/advent_of_code || exit 0