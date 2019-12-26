all:
	make clean
	git add -A . && git commit -m "update `dt`"
	git push

clean:
	rm -rf *.html

