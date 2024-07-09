/**
 * Graph API Error
 */

// TODO: 仮置き
type GraphMainError = {
  code: string;
  message: string;
  target?: string;
  details?: {
    code: string;
    message: string;
    target?: string | null;
  }[];
  innerError?: {
    'request-id'?: string | null;
    'client-request-id'?: string | null;
    date?: string | null;
    '@odata.type': string;
  };
};

export class GraphError extends Error {
  code: string;
  target?: string | null;
  details?: {
    code: string;
    message: string;
    target?: string | null;
  }[];
  innerError?: {
    'request-id'?: string | null;
    'client-request-id'?: string | null;
    date?: string | null;
    '@odata.type': string;
  };

  constructor(error: GraphMainError) {
    super(error.message);

    this.name = 'GraphError';
    this.code = error.code;
    this.target = error.target;
    this.details = error.details;
    this.innerError = error.innerError;

    if (Error.captureStackTrace) {
      Error.captureStackTrace(this, GraphError);
    }
  }
}
