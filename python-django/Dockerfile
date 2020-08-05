FROM python

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

WORKDIR /usr/src/app
COPY ./requirements.txt .
RUN pip install -r requirements.txt

COPY ./manage.py .
COPY ./issue ./issue/
COPY ./pythonDjango ././pythonDjango/
CMD ["gunicorn", "pythonDjango.wsgi:application","--bind","0.0.0.0:8080","--capture-output"]
