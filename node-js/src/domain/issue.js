class Issue {
    constructor(id, name,description){
        this.id = id;
        this.name = name;
        this.description = description;
    }

    getId() {
        return this.id
    }

    getName() {
        return this.name
    }

    getDescription() {
        return this.description
    }

	update(name, description) {
        return new Issue(
            this.id,
            Issue.isEmpty(name) ? this.name : name,
            Issue.isEmpty(description) ? this.description : description);
    }
    
    static isEmpty(str) {
        return !str || str.length === 0;
    }
}

module.exports = Issue;
