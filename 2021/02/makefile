go = /usr/local/bin/go

two_one = /tmp/advent_of_code/2021/02/01
two_two = /tmp/advent_of_code/2021/02/02
one_sample:
	$(go) build -o $(two_one)/main 01/main.go && \
    cp 01/sample.txt $(two_one)/sample.txt && \
    $(two_one)/main $(two_one)/sample.txt
one_input:
	$(go) build -o $(two_one)/main 01/main.go && \
    cp 01/input.txt $(two_one)/input.txt && \
    $(two_one)/main $(two_one)/input.txt
two_sample:
	$(go) build -o $(two_two)/main 02/main.go && \
    cp 02/sample.txt $(two_two)/sample.txt && \
    $(two_two)/main $(two_two)/sample.txt
two_input:
	$(go) build -o $(two_two)/main 02/main.go && \
    cp 02/input.txt $(two_two)/input.txt && \
    $(two_two)/main $(two_two)/input.txt
clean:
	rm -r /tmp/advent_of_code || exit 0