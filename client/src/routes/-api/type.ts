import type { Unit } from '@/entity/Unit';

export type BelongUnits = {
  units: Unit[];
};

export type Info = {
  version: string;
  updated: Date;
  modified: Date;
};
