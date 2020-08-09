import json
import falcon
from ..domain.issue import Issue
from ..domain.notFoundException import NotFoundException

#TODO Implement the Database

class IssueRepository(object):
    def findSingle(self, id):
        raise NotFoundException(id)
    def findAll(self):
        return [Issue("123e4567-e89b-12d3-a456-426614174000", "Test1", "First Test"), Issue("123e4567-e89b-12d3-a456-426614174010", "Test10", "Another Test")]
    def create(self, issue):
        return issue
    def update(self, issue):
        return issue
    def delete(self, id):
        pass