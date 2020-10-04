import {IssueRepository as IssueRepositoryInterface, NotFoundException} from '../domain/issueRepository'
import { Issue } from '../domain/issue';
import {Client} from 'pg'

export class IssueRepository implements IssueRepositoryInterface {

    constructor(private readonly client: Client){

    }

    async findSingle(id: string): Promise<Issue> {
        const res = await this.client.query('SELECT id, name, description FROM issue WHERE id = $1', [id]);
        const row = res.rows;
        if (row.length == 0) { 
            throw new NotFoundException(); 
        }
        return this.rowToIssue(row[0]);
    }

    async findAll(): Promise<Issue[]> {
        const res = await this.client
            .query('SELECT id, name, description FROM issue');
        return res.rows.map(i => this.rowToIssue(i));
    }

    async create(issue: Issue): Promise<Issue>  {
        await this.client
            .query('INSERT INTO issue (id, name, description) VALUES($1,$2,$3)', [issue.getId(), issue.getName(), issue.getDescription()]);
        return issue;
    }

    async update(issue: Issue): Promise<Issue> {
        await this.client
            .query('UPDATE issue SET name=$2, description=$3 WHERE id=$1', [issue.getId(), issue.getName(), issue.getDescription()]);
        return issue;
    }

    async delete(id: string): Promise<void> {
        await this.client
            .query('DELETE FROM issue WHERE id=$1', [id])
            .then();
    }

    /* eslint @typescript-eslint/no-explicit-any: "off" */
    /* eslint @typescript-eslint/explicit-module-boundary-types: "off" */
    private rowToIssue(row: any): Issue {
        return new Issue(row.id, row.name, row.description);
    }
}
