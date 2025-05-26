.PHONY: buffered

buffered:
	go run channels/buffered/buffered.go

.PHONY: range-close

range-close:
	go run channels/range-close/fibonacci.go

.PHONY: select

select:
	go run channels/select/select.go

.PHONY: channels

channels:
	go run channels/main.go

.PHONY: fanin

fanin:
	go run fan-in/main.go

.PHONY: generators-handles

generators:
	go run generators-handles/main.go

.PHONY: intro

intro:
	go run intro/main.go

.PHONY: mutex

mutex:
	go run mutex/main.go