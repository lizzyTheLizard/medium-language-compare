const Issue = require('../domain/issue');
const NotFoundException = require('../domain/notFoundException')
class IssueRepository {

    constructor(client){
        this.client = client;
    }

    async findSingle(id) {
        const res = await this.client.query('SELECT id, name, description FROM issue WHERE id = $1', [id]);
        const row = res.rows;
        if (row.length == 0) { 
            throw new NotFoundException(); 
        }
        return this.rowToIssue(row[0]);
    }

    async findAll() {
        const res = await this.client
            .query('SELECT id, name, description FROM issue');
        return res.rows.map(i => this.rowToIssue(i));
    }

    async create(issue)  {
        await this.client
            .query('INSERT INTO issue (id, name, description) VALUES($1,$2,$3)', [issue.getId(), issue.getName(), issue.getDescription()]);
        return issue;
    }

    async update(issue) {
        await this.client
            .query('UPDATE issue SET name=$2, description=$3 WHERE id=$1', [issue.getId(), issue.getName(), issue.getDescription()]);
        return issue;
    }

    async delete(id) {
        await this.client
            .query('DELETE FROM issue WHERE id=$1', [id])
            .then();
    }

    rowToIssue(row) {
        return new Issue(row.id, row.name, row.description);
    }
}

module.exports = IssueRepository;