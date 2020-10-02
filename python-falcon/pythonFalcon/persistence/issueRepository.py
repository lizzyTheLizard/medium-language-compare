import json
import falcon
from ..domain.issue import Issue
from ..domain.notFoundException import NotFoundException
import psycopg2

class IssueRepository(object):
    def __init__(self):
        self._db = psycopg2.connect("host=postgres dbname=postgres user=postgres password=postgres")
    def findSingle(self, id):
        cur = self._db.cursor()
        cur.execute('SELECT id, name, description FROM issue WHERE id = %s', (str(id),))
        row = cur.fetchone()
        cur.close()
        if row is None:
            raise NotFoundException(id)
        return Issue(row[0], row[1], row[2])
    def findAll(self):
        cur = self._db.cursor()
        cur.execute('SELECT id, name, description FROM issue')
        result = [];
        row = cur.fetchone()
        while row is not None:
            result.append(Issue(row[0], row[1], row[2]))
            row = cur.fetchone()
        cur.close()
        return result;
    def create(self, issue):
        cur = self._db.cursor()
        cur.execute('INSERT INTO issue (id, name, description) VALUES(%s,%s,%s)', (str(issue.getId()), issue.getName(), issue.getDescription()))
        cur.close()
        return issue
    def update(self, issue):
        cur = self._db.cursor()
        cur.execute('UPDATE issue SET name=%s, description=%s WHERE id=%s', (issue.getName(), issue.getDescription(),str(issue.getId())))
        cur.close()
        return issue
    def delete(self, id):
        cur = self._db.cursor()
        cur.execute('DELETE FROM issue WHERE id = %s', (str(id),))
        cur.close()
