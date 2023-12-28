1.train:
	@go run $(day)/part1/main.go $(day)/train1.txt

2.train:
	@go run $(day)/part2/main.go $(day)/train2.txt

1.test:
	@go run $(day)/part1/main.go $(day)/test.txt

2.test:
	@go run $(day)/part2/main.go $(day)/test.txt