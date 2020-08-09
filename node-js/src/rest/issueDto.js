const Issue = require('../domain/issue');
const uuid = require('uuidv4');
const InvalidIdException = require('./invalidIdException')

class IssueDto {
    static fromIssue(issue){
        const issueDto = new IssueDto();
        issueDto.id = issue.getId();
        issueDto.name = issue.getName();
        issueDto.description = issue.getDescription();
        return issueDto;
    }

    static fromJson(json) {
		if(!uuid.isUuid(json.id)){
			throw new InvalidIdException(json.id);
		}
        return new Issue(json.id, json.name, json.description);
    }
}

module.exports = IssueDto