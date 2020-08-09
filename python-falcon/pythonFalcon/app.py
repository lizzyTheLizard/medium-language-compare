from .rest.issueController import IssueController
from .persistence.issueRepository import IssueRepository
from .rest.engine import Engine

issueRepository = IssueRepository()
issueController = IssueController(issueRepository)
engine = Engine(issueController)
application = engine.getApplication()