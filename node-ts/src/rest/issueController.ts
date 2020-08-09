import express from 'express'
import { isUuid } from 'uuidv4';

import {IssueRepository} from '../domain/issueRepository'
import { IssueDto } from './issueDto';

export class InvalidIdException {constructor(readonly id: string){}}

export class IssueController {
    public constructor(private readonly issueRepository: IssueRepository) {
    }

    async getAll(req: express.Request, res: express.Response, next: express.NextFunction): void {
        try {
            const issues = await this.issueRepository.findAll();
            const issueDtos = issues.map(issue => IssueDto.fromIssue(issue));
            res.status(200).json(issueDtos).send();
        } catch (error) {
            next(error);
        }
    }
	
    async getSingle(req: express.Request, res: express.Response, next: express.NextFunction): void {
        try {
            const id = this.getId(req);
            const issue = await this.issueRepository.findSingle(id);
            const issueDto = IssueDto.fromIssue(issue);
            res.status(200).json(issueDto).send();
        } catch (error) {
            next(error);
        }
    }
	
    async create(req: express.Request, res: express.Response, next: express.NextFunction): void {
        try {
            const postBody = IssueDto.fromJson(req.body);
            const newIssue = await this.issueRepository.create(postBody)
            res.status(200).json(newIssue).send();
        } catch (error) {
            next(error);
        }
    }
	
    async update(req: express.Request, res: express.Response, next: express.NextFunction): void {
        try {
            const postBody = IssueDto.fromJson(req.body);
            const newIssue = await this.issueRepository.update(postBody)
            res.status(200).json(newIssue).send();
        } catch (error) {
            next(error);
        }
    }
	
    async partialUpdate(req: express.Request, res: express.Response, next: express.NextFunction): void {
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
	
    async delete(req: express.Request, res: express.Response, next: express.NextFunction): void {
        try {
            const id = this.getId(req);
            await this.issueRepository.delete(id);
            res.sendStatus(200);
        } catch (error) {
            next(error);
        }
    }

    private getId(req: express.Request): string {
        const id = req.params['id'];
        if (!isUuid(id)){
            throw new InvalidIdException(id);
        }
        return id;
    }
}
