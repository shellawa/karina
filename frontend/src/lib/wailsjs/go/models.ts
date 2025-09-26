export namespace models {
	
	export interface Participant {
	    id: string;
	    name: string;
	    organization: string;
	}
	export interface Problem {
	    id: string;
	    label: string;
	    description: string;
	    time: number;
	    memory: number;
	}
	export interface TestCase {
	    Id: string;
	    Input: string;
	    Output: string;
	}

}

