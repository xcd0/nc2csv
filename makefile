all:
	make clean
	git add -A . && git commit -m "update `date +"%Y.%m.%d.%H.%M.%S"`"
	git push

clean:
	rm -rf *.html

