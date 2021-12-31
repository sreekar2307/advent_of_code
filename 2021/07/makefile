go = /usr/local/bin/go

seven_one = /tmp/advent_of_code/2021/07/01
seven_two = /tmp/advent_of_code/2021/07/02
one_sample:
	$(go) build -o $(seven_one)/main 01/main.go && \
    cp 01/sample.txt $(seven_one)/sample.txt && \
    $(seven_one)/main $(seven_one)/sample.txt
one_input:
	$(go) build -o $(seven_one)/main 01/main.go && \
    cp 01/input.txt $(seven_one)/input.txt && \
    $(seven_one)/main $(seven_one)/input.txt
two_sample:
	$(go) build -o $(seven_two)/main 02/main.go && \
    cp 02/sample.txt $(seven_two)/sample.txt && \
    $(seven_two)/main $(seven_two)/sample.txt
two_input:
	$(go) build -o $(seven_two)/main 02/main.go && \
    cp 02/input.txt $(seven_two)/input.txt && \
    $(seven_two)/main $(seven_two)/input.txt
clean:
	rm -r /tmp/advent_of_code || exit 0