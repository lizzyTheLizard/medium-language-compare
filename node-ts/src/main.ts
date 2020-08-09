import {IssueController} from './rest/issueController'
import {IssueRepository} from './persistence/issueRepository'
import {Engine} from './rest/engine'
import {connect} from './persistence/connect';


connect().then(client => {
    const issueRepository = new IssueRepository(client);
    const issueEndpoint = new IssueController(issueRepository);
    const engine = new Engine(issueEndpoint);
    engine.run();
}).catch(e => console.error('Cannot start application',e))

