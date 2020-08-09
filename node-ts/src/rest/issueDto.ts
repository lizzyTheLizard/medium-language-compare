import {Issue} from '../domain/issue'
import { isUuid } from 'uuidv4';
import { InvalidIdException } from './issueController';

export class IssueDto {
    public id : string = "";
    public name : string = "";
    public description: string = "";

    static fromIssue(issue : Issue) : IssueDto{
        const issueDto = new IssueDto();
        issueDto.id = issue.getId();
        issueDto.name = issue.getName();
        issueDto.description = issue.getDescription();
        return issueDto;
    }

    static fromJson(json: any): Issue {
        const issueDto = new IssueDto();
        console.log(json)
		if(!isUuid(json.id)){
			throw new InvalidIdException(json.id);
		}
        return new Issue(json.id, json.name, json.description);
    }
}