FROM scratch

COPY ./application/bin/entrypoint /application/entrypoint
COPY ./fonts/ /fonts/
COPY ./js/ /js/
COPY ./stylesheets/ /stylesheets/
COPY ./templates/ /templates/

EXPOSE 8080

# Use the User directive to ensure not run as root as no need
ENTRYPOINT [ "/application/entrypoint" ]
