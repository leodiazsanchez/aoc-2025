run:
	go run cmd/aoc/main.go $(day)

get:
	go run cmd/aoc/main.go $(day) -get

create:
	@DAY=$$(printf "%02d" $(day)); \
	mkdir -p internal/day$$DAY && \
	touch internal/day$$DAY/main.go && \
	touch inputs/day$$DAY.txt
