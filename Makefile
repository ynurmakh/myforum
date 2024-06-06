run:
	go mod vendor 
	docker run -it --rm --name forum -p 9999:9999 -v $(PWD):/APP/. ynurmakh/forum:latest

rund:
	go mod vendor 
	docker run -d --rm --name forum -p 9999:9999 -v $(PWD):/APP/. ynurmakh/forum:latest

restart:
	-docker kill forum
	sleep 1
	make rund 

push:
	git add .
	git commit -m "$(m)"
	git push 