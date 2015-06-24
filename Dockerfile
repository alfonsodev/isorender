FROM node
RUN git clone http://github.com/alfonsodev/isorender
RUN cd isorender; npm install;
ENTRYPOINT ["/usr/local/bin/node", "/isorender/index"]
EXPOSE 3000
VOLUME /isorender/public
