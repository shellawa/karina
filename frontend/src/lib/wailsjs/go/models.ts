export namespace common {
	
	export interface TestSolveResult {
	    verdict: string;
	    time: number;
	    memory: number;
	    actual_output: string;
	}

}

export namespace models {
	
	export interface Participant {
	    id: string;
	    name: string;
	    organization: string;
	    code?: string;
	}
	export interface ParticipantSolveResult {
	    id: string;
	    results: common.TestSolveResult[];
	}
	export interface Problem {
	    id: string;
	    label: string;
	    description: string;
	    time: number;
	    memory: number;
	    io_mode: number;
	}
	export interface TestCase {
	    Id: string;
	    Input: string;
	    Output: string;
	}

}

