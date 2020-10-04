import falcon

from ..domain.notFoundException import NotFoundException


class Engine(object):
    def __init__(self, issueController):
        self.__api = falcon.API()
        self.__api.add_route('/issue/', issueController)
        self.__api.add_route('/issue/{id:uuid}/', issueController, suffix="single")
        self.__api.add_error_handler(NotFoundException, self.__notFoundHandler)
        self.__api.req_options.strip_url_path_trailing_slash = True

    def getApplication(self):
        return self.__api

    def __notFoundHandler(self, ex, req, resp, params):
        raise falcon.HTTPNotFound(description=ex.getMessage())
