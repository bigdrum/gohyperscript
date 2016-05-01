publish:
	cd js && \
	gopherjs build -m && \
	uglifyjs js.js -m -c -o ../../js.min.js && \
	rm js.js && \
	cd .. && cd .. && \
	git checkout gh-pages && \
	mv js.min.js js.js && \
	git commit -a -m "Update js." && \
	git checkout master
