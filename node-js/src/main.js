const connect = require('./persistence/connect');
const Engine = require('./rest/engine');
const IssueController = require('./rest/issueController');
const IssueRepository = require('./persistence/issueRepository');

connect.connect().then(client => {
    const issueRepository = new IssueRepository(client);
    const issueController = new IssueController(issueRepository);
    const engine = new Engine(issueController);
    engine.run();
}).catch(e => console.error('Cannot start application',e));
