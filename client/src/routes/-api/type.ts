import type { Unit } from '@/entity/Unit';

export type BelongUnitsResponse =
  | {
      status: 'error';
      message: 'string';
    }
  | {
      status: 'success';
      unit_count: number;
      units: Unit[];
    };

export type Auth = {
  userId: string | undefined;
  idToken: string | undefined;
};

export type Info = {
  version: string;
  created: Date;
  modified: Date;
};

// interface ErrorOptions {
//   cause?: Error;
//   details?: unknown;
// }

// export class APIError extends Error {
//   options?: ErrorOptions;
//   constructor(message?: string, options?: ErrorOptions) {
//     super(message);
//     this.name = 'APIError';
//     this.options = options;
//   }
// }
