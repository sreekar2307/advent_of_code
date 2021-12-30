go = /usr/local/bin/go

three_one = /tmp/advent_of_code/2021/03/01
three_two = /tmp/advent_of_code/2021/03/02
one_sample:
	$(go) build -o $(three_one)/main 01/main.go && \
    cp 01/sample.txt $(three_one)/sample.txt && \
    $(three_one)/main $(three_one)/sample.txt
one_input:
	$(go) build -o $(three_one)/main 01/main.go && \
    cp 01/input.txt $(three_one)/input.txt && \
    $(three_one)/main $(three_one)/input.txt
two_sample:
	$(go) build -o $(three_two)/main 02/main.go && \
    cp 02/sample.txt $(three_two)/sample.txt && \
    $(three_two)/main $(three_two)/sample.txt
two_input:
	$(go) build -o $(three_two)/main 02/main.go && \
    cp 02/input.txt $(three_two)/input.txt && \
    $(three_two)/main $(three_two)/input.txt
clean:
	rm -r /tmp/advent_of_code || exit 0