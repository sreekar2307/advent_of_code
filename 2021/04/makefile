go = /usr/local/bin/go

four_one = /tmp/advent_of_code/2021/04/01
four_two = /tmp/advent_of_code/2021/04/02
one_sample:
	$(go) build -o $(four_one)/main 01/main.go && \
    cp 01/sample.txt $(four_one)/sample.txt && \
    $(four_one)/main $(four_one)/sample.txt
one_input:
	$(go) build -o $(four_one)/main 01/main.go && \
    cp 01/input.txt $(four_one)/input.txt && \
    $(four_one)/main $(four_one)/input.txt
two_sample:
	$(go) build -o $(four_two)/main 02/main.go && \
    cp 02/sample.txt $(four_two)/sample.txt && \
    $(four_two)/main $(four_two)/sample.txt
two_input:
	$(go) build -o $(four_two)/main 02/main.go && \
    cp 02/input.txt $(four_two)/input.txt && \
    $(four_two)/main $(four_two)/input.txt
clean:
	rm -r /tmp/advent_of_code || exit 0