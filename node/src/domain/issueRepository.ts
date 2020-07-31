import {Issue} from './issue'

export interface IssueRepository {
	findSingle(id: string) : Promise<Issue>;
	findAll(): Promise<Issue[]>;
	create(issue: Issue): Promise<Issue>
	update(issue: Issue): Promise<Issue>
	delete(id: string): Promise<void>
}

export class NotFoundException{

}