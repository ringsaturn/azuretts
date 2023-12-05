
mock:
	mockgen -source=client.go  -destination=mockclient/client.go -package=mockclient

.PHONY:clean
clean:
	find . -name '*.mp3' | xargs rm
