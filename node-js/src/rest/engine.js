const express = require('express')
const NotFoundException = require('../domain/notFoundException')
const InvalidIdException = require('./invalidIdException')

class Engine {    

    constructor(issueController){
        this.app = express();
        this.app.use(express.json());
		this.app.get("/issue/", (req,res,next) => issueController.getAll(req,res,next))
		this.app.get("/issue/:id", (req,res,next) => issueController.getSingle(req,res,next))
		this.app.post("/issue", (req, res, next) => issueController.create(req,res,next))
		this.app.post("/issue", (req, res, next) => issueController.create(req,res,next))
		this.app.put("/issue/:id", (req, res, next) => issueController.update(req,res,next))
		this.app.patch("/issue/:id", (req, res, next) => issueController.partialUpdate(req,res,next))
        this.app.delete("/issue/:id", (req,res,next) => issueController.delete(req,res,next))
        this.app.use(this.errorHandler);
    }

    errorHandler(err, request, res, next){    
        if(err instanceof NotFoundException) {
            res.status(404).json({"error":"Issue not found"})
            return;
        }
        if(err instanceof InvalidIdException) {
            console.log('Invalid ID',err.id)
            res.status(400).json({"error":"Invalid ID given"})
            return;
        }
        console.warn("Error while process request", err)
        res.status(500).send('Something broke!');
    }

    run(){
        this.app.listen(8080, () => {
            console.log("Server running on port 8080");
        });
    }
}
module.exports = Engine