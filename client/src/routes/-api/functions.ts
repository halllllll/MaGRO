import type { BelongUnitsResponse, Info } from './type';

export const getMaGROInfo = async (IdToken: string | undefined): Promise<Info> => {
  const res = await fetch('/api/info', {
    method: 'GET',
    mode: 'cors',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${IdToken}`,
    },
  });
  if (!res.ok) {
    throw new Error(`${res.status} ${res.statusText}`);
  }
  return await res.json();
};

export const getUnitsByUser = async (IdToken: string | undefined): Promise<BelongUnitsResponse> => {
  const res = await fetch('/api/units', {
    method: 'GET',
    mode: 'cors',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${IdToken}`,
    },
  });

  if (!res.ok) {
    throw new Error(`${res.status} ${res.statusText}`);
  }

  return await res.json();
};
