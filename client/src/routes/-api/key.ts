export const belongsUnitKeys = {
  units: () => ['units'] as const,
  unit: (id: number) => [...belongsUnitKeys.units(), id] as const,
} as const;

export const info = {
  data: ['info'] as const,
} as const;
