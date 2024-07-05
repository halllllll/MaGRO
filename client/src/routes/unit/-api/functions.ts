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
