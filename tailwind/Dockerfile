FROM node:latest
WORKDIR /tailwind
# RUN type npm
# RUN npm init -y 
# RUN npm install -D tailwindcss@3
# # RUN npm install -g npx
# RUN npx tailwindcss init

RUN npm init -y && \
    npm install tailwindcss@3 && \
    npx tailwindcss init

CMD npx tailwindcss -c /src/tailwind.config.js -i /src/styles.css -o /dst/styles.css --watch --poll
# CMD ["npx", "tailwindcss", "-c", "/src/tailwind.config.js", "-i", "/src/styles.css", "-o", "/dst/styles.css", "--watch", "--poll"]


