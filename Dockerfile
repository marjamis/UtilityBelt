FROM scratch

COPY ./application/bin/entrypoint /application/entrypoint
COPY ./fonts/ /fonts/
COPY ./js/ /js/
COPY ./stylesheets/ /stylesheets/
COPY ./templates/ /templates/

EXPOSE 8080

ENTRYPOINT [ "/application/entrypoint" ]
