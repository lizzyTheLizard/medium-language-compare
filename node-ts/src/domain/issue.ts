export class Issue {
    constructor(private readonly id: string, 
        private readonly name : string, 
        private readonly description: string){

    }

    getId(): string {
        return this.id
    }

    getName(): string {
        return this.name
    }

    getDescription(): string {
        return this.description
    }

    update(name: string, description: string) : Issue {
        return new Issue(
            this.id,
            Issue.isEmpty(name) ? this.name : name,
            Issue.isEmpty(description) ? this.description : description);
    }
    
    private static isEmpty(str: string): boolean {
        return !str || str.length === 0;
    }
}