.PHONY: docs docs-no-pull docs-test docs-local-static

IMAGE = grafarg/grafarg-docs-dev:latest
CONTENT_PATH = /hugo/content/docs/grafarg/latest
LOCAL_STATIC_PATH = ../../website/static
PORT = 3002:3002

docs:
	docker pull $(IMAGE)
	docker run -v $(shell pwd)/sources:$(CONTENT_PATH):Z -p $(PORT) --rm -it $(IMAGE)

docs-no-pull:
	docker run -v $(shell pwd)/sources:$(CONTENT_PATH):Z -p $(PORT) --rm -it $(IMAGE)

docs-test:
	docker pull $(IMAGE)
	docker run -v $(shell pwd)/sources:$(CONTENT_PATH):Z --rm -it $(IMAGE) /bin/bash -c 'make prod'

# expects that you have grafarg/website checked out in same path as the grafarg repo.
docs-local-static:
	docker pull $(IMAGE)
	if [ ! -d "$(LOCAL_STATIC_PATH)" ]; then echo "local path (website project) $(LOCAL_STATIC_PATH) not found"]; exit 1; fi
	docker run -v $(shell pwd)/sources:$(CONTENT_PATH):Z \
		-v $(shell pwd)/$(LOCAL_STATIC_PATH):/hugo/static:Z -p $(PORT) --rm -it $(IMAGE)
