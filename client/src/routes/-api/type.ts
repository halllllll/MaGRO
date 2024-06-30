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

export type Info = {
  version: string;
  updated: Date;
  modified: Date;
};

export type Auth = {
  userId: string | undefined;
  idToken: string | undefined;
};
