import {Issue} from '../domain/issue'
import { isUuid } from 'uuidv4';
import { InvalidIdException } from './issueController';

export class IssueDto {
    public id  = '';
    public name  = '';
    public description = '';

    static fromIssue(issue : Issue) : IssueDto{
        const issueDto = new IssueDto();
        issueDto.id = issue.getId();
        issueDto.name = issue.getName();
        issueDto.description = issue.getDescription();
        return issueDto;
    }

    /* eslint @typescript-eslint/no-explicit-any: "off" */
    /* eslint @typescript-eslint/explicit-module-boundary-types: "off" */
    static fromJson(json: any): Issue {
        console.log(json)
        if (!isUuid(json.id)){
            throw new InvalidIdException(json.id);
        }
        return new Issue(json.id, json.name, json.description);
    }
}
