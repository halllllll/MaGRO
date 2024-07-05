export const belongsSubunitKeys = {
  subunit: () => ['subunit'] as const,
  here: (unitid: number) => [...belongsSubunitKeys.subunit(), unitid] as const,
} as const;
