const Unit_ID: string = 'MAGRO_SESSION_USER_UNIT_ID';

export const SetUnitID = (unitId: string) => {
  sessionStorage.setItem(Unit_ID, unitId);
};

export const GetUnitID = () => {
  return sessionStorage.getItem(Unit_ID);
};

export const RemoveUnitID = () => {
  sessionStorage.removeItem(Unit_ID);
};
