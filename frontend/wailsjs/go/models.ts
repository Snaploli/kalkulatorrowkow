export namespace main {
	
	export class CalculationResult {
	    p: number;
	    q: number;
	    external: number;
	    grooveType: string;
	    suggestedName: string;
	    suggestedP: number;
	    suggestedW: number;
	    suggestedGw: number;
	    errMsg: string;
	
	    static createFrom(source: any = {}) {
	        return new CalculationResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.p = source["p"];
	        this.q = source["q"];
	        this.external = source["external"];
	        this.grooveType = source["grooveType"];
	        this.suggestedName = source["suggestedName"];
	        this.suggestedP = source["suggestedP"];
	        this.suggestedW = source["suggestedW"];
	        this.suggestedGw = source["suggestedGw"];
	        this.errMsg = source["errMsg"];
	    }
	}

}

