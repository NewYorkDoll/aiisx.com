


node-fetch:
	command -v pnpm >/dev/null >&2 || npm install \
		--no-audit \
		--no-fund \
		--quiet \
		--global pnpm
	cd src/public && \
		yarn install


node-debug:
	cd src/public && \
		yarn dev