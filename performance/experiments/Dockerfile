FROM python:3.7.0

RUN pip install -U chaostoolkit
RUN chaos --version
RUN pip install chaostoolkit-kubernetes
RUN chaos discover chaostoolkit-kubernetes

ENTRYPOINT ["/usr/local/bin/chaos"]
CMD ["--help"]