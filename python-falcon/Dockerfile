FROM python

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

WORKDIR /usr/src/app
COPY ./requirements.txt .
RUN pip install -r requirements.txt

COPY ./pythonFalcon ././pythonFalcon/
CMD ["gunicorn", "pythonFalcon.app","--bind","0.0.0.0:8080","--capture-output"]
