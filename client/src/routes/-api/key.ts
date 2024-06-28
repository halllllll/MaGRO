export const belongsUnitKeys = {
  units: () => ['units'] as const,
  me: (id: string | undefined) => [...belongsUnitKeys.units(), id] as const,
} as const;

export const info = {
  data: () => ['info'] as const,
  me: (id: string | undefined) => [...info.data(), id],
} as const;
