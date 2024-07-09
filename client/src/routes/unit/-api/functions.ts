import type { RepassRequest, RepassResponse } from './type';

export const getUnitData = async (unitId: number, IdToken: string | undefined) => {
  const res = await fetch(`/api/subunit/${unitId}`, {
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

// TODO: 戻り値の型はまだAPI側で定義も実装もない
export const repass = async (data: RepassRequest): Promise<RepassResponse> => {
  const body = {
    current_user: data.current_user,
    target_users: data.target_user,
  };
  const res = await fetch(`/api/units/${data.unitId}/repass`, {
    method: 'POST',
    mode: 'cors',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${data.auth.idToken}`,
    },
    body: JSON.stringify(body),
  });

  if (!res.ok) {
    throw new Error(`${res.status} ${res.statusText}`);
  }
  return await res.json();
};
