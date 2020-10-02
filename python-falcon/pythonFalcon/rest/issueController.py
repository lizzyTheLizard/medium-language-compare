import json
import falcon
from ..domain.issue import Issue

class IssueController(object):
    def __init__(self, issueRepository):
        self.__repository = issueRepository

    def on_get(self, req, resp):
        result = [self.__issueToDict(i) for i in self.__repository.findAll()]
        resp.body = json.dumps(result)
        resp.status = falcon.HTTP_200

    def on_get_single(self, req, resp, id):
        result = self.__repository.findSingle(id)
        resp.body = json.dumps(self.__issueToDict(result))
        resp.status = falcon.HTTP_200

    def on_post(self, req, resp):
        input = self.__bodyToIssue(req)
        output = self.__repository.create(input)
        resp.body = json.dumps(self.__issueToDict(output))
        resp.status = falcon.HTTP_200

    def on_put_single(self, req, resp, id):
        input = self.__bodyToIssue(req)
        output = self.__repository.update(input)
        resp.body = json.dumps(self.__issueToDict(output))
        resp.status = falcon.HTTP_200

    def on_patch_single(self, req, resp, id):
        issue = self.__repository.findSingle(id)
        issue = issue.update(req.media.get("name", None), req.media.get('description', None))
        output = self.__repository.update(issue)
        resp.body = json.dumps(self.__issueToDict(output))
        resp.status = falcon.HTTP_200

    def on_delete_single(self, req, resp, id):
        self.__repository.delete(id)
        resp.status = falcon.HTTP_200

    def __issueToDict(self, issue):
        return {"id":issue.getId(), "name":issue.getName(), "description":issue.getDescription()}

    def __bodyToIssue(self, req):
        return Issue(req.media['id'], req.media['name'], req.media['description'])
