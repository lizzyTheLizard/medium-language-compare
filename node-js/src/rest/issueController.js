const uuid = require('uuidv4');
const InvalidIdException = require('./invalidIdException')
const IssueDto = require('./issueDto')


class IssueController {
	constructor(issueRepository) {
		this.issueRepository = issueRepository;
	}

	async getAll(req, res, next) {
		try {
			const issues = await this.issueRepository.findAll();
			const issueDtos = issues.map(issue => IssueDto.fromIssue(issue));
			res.status(200).json(issueDtos).send();
		} catch (error) {
			next(error);
		}
	}
	
	async getSingle(req, res, next) {
		try {
			const id = this.getId(req);
			const issue = await this.issueRepository.findSingle(id);
			const issueDto = IssueDto.fromIssue(issue);
			res.status(200).json(issueDto).send();
		} catch (error) {
			next(error);
		}
	}
	
	async create(req, res, next) {
		try {
			const postBody = IssueDto.fromJson(req.body);
			const newIssue = await this.issueRepository.create(postBody)
			res.status(200).json(newIssue).send();
		} catch (error) {
			next(error);
		}
	}
	
	async update(req, res, next) {
		try {
			const postBody = IssueDto.fromJson(req.body);
			const newIssue = await this.issueRepository.update(postBody)
			res.status(200).json(newIssue).send();
		} catch (error) {
			next(error);
		}
	}
	
	async partialUpdate(req, res, next) {
		try {
			const id = this.getId(req);
			const oldIssue = await this.issueRepository.findSingle(id);
			const updatedIssue = oldIssue.update(req.body.name, req.body.description);
			const newIssue = await this.issueRepository.update(updatedIssue)
			res.status(200).json(newIssue).send();
		} catch (error) {
			next(error);
		}
	}
	
	async delete(req, res, next) {
		try {
			const id = this.getId(req);
			await this.issueRepository.delete(id);
			res.sendStatus(200);
		} catch (error) {
			next(error);
		}
	}

	getId(req, res, next) {
		const id = req.params['id'];
		if(!uuid.isUuid(id)){
			throw new InvalidIdException(id);
		}
		return id;
	}
}

module.exports = IssueController
